package service

import (
	"context"
	"crypto/ecdsa"
	"dhb/app/app/internal/pkg/middleware/auth"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	jwt2 "github.com/golang-jwt/jwt/v4"
	"io"
	"math/big"
	"net/url"
	"strconv"
	"strings"

	v1 "dhb/app/app/api"
	"dhb/app/app/internal/biz"
	"dhb/app/app/internal/conf"
	"io/ioutil"
	"net/http"
	"time"
)

// AppService service.
type AppService struct {
	v1.UnimplementedAppServer

	uuc *biz.UserUseCase
	ruc *biz.RecordUseCase
	log *log.Helper
	ca  *conf.Auth
}

// NewAppService new a service.
func NewAppService(uuc *biz.UserUseCase, ruc *biz.RecordUseCase, logger log.Logger, ca *conf.Auth) *AppService {
	return &AppService{uuc: uuc, ruc: ruc, log: log.NewHelper(logger), ca: ca}
}

// EthAuthorize ethAuthorize.
func (a *AppService) EthAuthorize(ctx context.Context, req *v1.EthAuthorizeRequest) (*v1.EthAuthorizeReply, error) {
	// TODO 有效的参数验证
	userAddress := req.SendBody.Address // 以太坊账户
	if "" == userAddress || 20 > len(userAddress) ||
		strings.EqualFold("0x000000000000000000000000000000000000dead", userAddress) ||
		strings.EqualFold("0xc1a19468996CB1bf108D00ec3d7fDfcaAbab883f", userAddress) ||
		strings.EqualFold("0x0a8B0dC827E237Ca26AC83e56BabE7D43102F798", userAddress) ||
		strings.EqualFold("0x5af02e8205427a690256b458001fbf7f679b1e15", userAddress) {
		return nil, errors.New(500, "AUTHORIZE_ERROR", "账户地址参数错误")
	}

	// TODO 验证签名

	// 根据地址查询用户，不存在时则创建
	user, err := a.uuc.GetExistUserByAddressOrCreate(ctx, &biz.User{
		Address: userAddress,
	}, req)
	if err != nil {
		return nil, err
	}

	claims := auth.CustomClaims{
		UserId:   user.ID,
		UserType: "user",
		StandardClaims: jwt2.StandardClaims{
			NotBefore: time.Now().Unix(),              // 签名的生效时间
			ExpiresAt: time.Now().Unix() + 60*60*24*7, // 7天过期
			Issuer:    "DHB",
		},
	}
	token, err := auth.CreateToken(claims, a.ca.JwtKey)
	if err != nil {
		return nil, errors.New(500, "AUTHORIZE_ERROR", "生成token失败")
	}

	userInfoRsp := v1.EthAuthorizeReply{
		Token: token,
	}
	return &userInfoRsp, nil
}

// Deposit deposit.
func (a *AppService) Deposit(ctx context.Context, req *v1.DepositRequest) (*v1.DepositReply, error) {

	var (
		depositUsdtResult map[string]*eth
		//depositDhbResult      map[string]*eth
		//tmpDepositDhbResult   map[string]*eth
		//userDepositDhbResult  map[string]map[string]*eth
		notExistDepositResult []*biz.EthUserRecord
		existEthUserRecords   map[string]*biz.EthUserRecord
		depositUsers          map[string]*biz.User
		fromAccount           []string
		hashKeys              []string
		//lock                  bool
		err error
		//configs               []*biz.Config
		//level1Dhb             string
		//level2Dhb             string
		//level3Dhb             string
	)

	// 配置
	//configs, err = a.uuc.GetDhbConfig(ctx)
	//if nil != configs {
	//	for _, vConfig := range configs {
	//		if "level1Dhb" == vConfig.KeyName {
	//			level1Dhb = vConfig.Value + "0000000000000000"
	//		} else if "level2Dhb" == vConfig.KeyName {
	//			level2Dhb = vConfig.Value + "0000000000000000"
	//		} else if "level3Dhb" == vConfig.KeyName {
	//			level3Dhb = vConfig.Value + "0000000000000000"
	//		}
	//	}
	//}

	//if lock, _ = a.ruc.LockEthUserRecordHandle(ctx); !lock { // 上全局锁简单，防止资源更新抢占
	//	return &v1.DepositReply{}, nil
	//}

	// 每次一共最多查2000条，所以注意好外层调用的定时查询的时间设置，当然都可以重新定义，
	// 在功能上调用者查询两种币的交易记录，每次都要把数据覆盖查询，是一个较大范围的查找防止遗漏数据，范围最起码要大于实际这段时间的入单量，不能边界查询容易掉单，这样的实现是因为简单
	for i := 1; i <= 10; i++ {

		depositUsdtResult, err = requestEthDepositResult(200, int64(i), "0x55d398326f99059fF775485246999027B3197955")
		// 辅助查询
		//depositDhbResult, err = requestEthDepositResult(200, int64(i), "0x96BD81715c69eE013405B4005Ba97eA1f420fd87")
		//tmpDepositDhbResult, err = requestEthDepositResult(100, int64(i+1), "0x96BD81715c69eE013405B4005Ba97eA1f420fd87")
		//for kTmpDepositDhbResult, v := range tmpDepositDhbResult {
		//	if _, ok := tmpDepositDhbResult[kTmpDepositDhbResult]; !ok {
		//		depositDhbResult[kTmpDepositDhbResult] = v
		//	}
		//}

		if 0 >= len(depositUsdtResult) {
			break
		}

		for hashKey, vDepositResult := range depositUsdtResult { // 主查询
			hashKeys = append(hashKeys, hashKey)
			fromAccount = append(fromAccount, vDepositResult.From)
		}
		//userDepositDhbResult = make(map[string]map[string]*eth, 0) // 辅助数据
		//for k, v := range depositDhbResult {
		//	hashKeys = append(hashKeys, k)
		//	fromAccount = append(fromAccount, v.From)
		//	if _, ok := userDepositDhbResult[v.From]; !ok {
		//		userDepositDhbResult[v.From] = make(map[string]*eth, 0)
		//	}
		//	userDepositDhbResult[v.From][k] = v
		//}

		depositUsers, err = a.uuc.GetUserByAddress(ctx, fromAccount...)
		if nil != err || nil == depositUsers {
			continue
		}
		existEthUserRecords, err = a.ruc.GetEthUserRecordByTxHash(ctx, hashKeys...)
		// 统计开始
		notExistDepositResult = make([]*biz.EthUserRecord, 0)
		for _, vDepositUsdtResult := range depositUsdtResult { // 主查usdt
			if _, ok := existEthUserRecords[vDepositUsdtResult.Hash]; ok { // 记录已存在
				continue
			}
			if _, ok := depositUsers[vDepositUsdtResult.From]; !ok { // 用户不存在
				continue
			}
			//if _, ok := userDepositDhbResult[vDepositUsdtResult.From]; !ok { // 没有dhb的充值记录
			//	continue
			//}
			//var (
			//	tmpDhbHash, tmpDhbHashValue string
			//)

			//tmpPass := false
			//for _, vUserDepositDhbResult := range userDepositDhbResult[vDepositUsdtResult.From] { // 充值数额类型匹配
			//	if _, ok := existEthUserRecords[vUserDepositDhbResult.Hash]; ok { // 记录已存在
			//		continue
			//	}
			//
			//	if "10000000000000000" == vDepositUsdtResult.Value {
			//		tmpPass = true
			//	} else if "20000000000000000" == vDepositUsdtResult.Value {
			//		tmpPass = true
			//	} else if "50000000000000000" == vDepositUsdtResult.Value {
			//		tmpPass = true
			//	} else {
			//		continue
			//	}
			//
			//	tmpDhbHash = vUserDepositDhbResult.Hash
			//	tmpDhbHashValue = vUserDepositDhbResult.Value
			//}
			//if !tmpPass {
			//	continue
			//}

			if "100000000000000000000" == vDepositUsdtResult.Value {

			} else if "200000000000000000000" == vDepositUsdtResult.Value {

			} else if "500000000000000000000" == vDepositUsdtResult.Value {

			} else {
				continue
			}

			notExistDepositResult = append(notExistDepositResult, &biz.EthUserRecord{ // 两种币的记录
				UserId:   depositUsers[vDepositUsdtResult.From].ID,
				Hash:     vDepositUsdtResult.Hash,
				Status:   "success",
				Type:     "deposit",
				Amount:   vDepositUsdtResult.Value,
				CoinType: "USDT",
			})
			//&biz.EthUserRecord{
			//	UserId:   depositUsers[vDepositUsdtResult.From].ID,
			//	Hash:     tmpDhbHash,
			//	Status:   "success",
			//	Type:     "deposit",
			//	Amount:   tmpDhbHashValue,
			//	CoinType: "DHB",
			//}
		}

		_, err = a.ruc.EthUserRecordHandle(ctx, notExistDepositResult...)
		if nil != err {
			fmt.Println(err)
		}

		//time.Sleep(2 * time.Second)
	}

	//_, _ = a.ruc.UnLockEthUserRecordHandle(ctx)
	return &v1.DepositReply{}, nil
}

type eth struct {
	Value       string
	Hash        string
	TokenSymbol string
	From        string
	To          string
}

func requestEthDepositResult(offset int64, page int64, contractAddress string) (map[string]*eth, error) {
	//apiUrl := "https://api-testnet.bscscan.com/api"
	apiUrl := "https://api.bscscan.com/api"
	// URL param
	data := url.Values{}
	data.Set("module", "account")
	data.Set("action", "tokentx")
	data.Set("contractaddress", contractAddress)
	data.Set("address", "0x636F2deAAb4C9A8F3c808D23F16f456009C4e9Fd")
	data.Set("sort", "desc")
	data.Set("offset", strconv.FormatInt(offset, 10))
	data.Set("page", strconv.FormatInt(page, 10))

	u, err := url.ParseRequestURI(apiUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = data.Encode() // URL encode
	client := http.Client{
		Timeout: 10 * time.Second,
	}
	fmt.Println(u.String())

	resp, err := client.Get(u.String())
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var i struct {
		Message string `json:"message"`
		Result  []*eth `json:"Result"`
	}
	err = json.Unmarshal(b, &i)
	if err != nil {
		return nil, err
	}

	res := make(map[string]*eth, 0)
	for _, v := range i.Result {
		if "0x636F2deAAb4C9A8F3c808D23F16f456009C4e9Fd" == v.To { // 接收者
			res[v.Hash] = v
		}
	}

	return res, err
}

// UserInfo userInfo.
func (a *AppService) UserInfo(ctx context.Context, req *v1.UserInfoRequest) (*v1.UserInfoReply, error) {
	// 在上下文 context 中取出 claims 对象
	var userId int64
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["UserId"] == nil {
			return nil, errors.New(500, "ERROR_TOKEN", "无效TOKEN")
		}
		userId = int64(c["UserId"].(float64))
	}

	return a.uuc.UserInfo(ctx, &biz.User{
		ID: userId,
	})
}

// RecommendUpdate recommendUpdate.
func (a *AppService) RecommendUpdate(ctx context.Context, req *v1.RecommendUpdateRequest) (*v1.RecommendUpdateReply, error) {
	// 在上下文 context 中取出 claims 对象
	var userId int64
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["UserId"] == nil {
			return nil, errors.New(500, "ERROR_TOKEN", "无效TOKEN")
		}
		userId = int64(c["UserId"].(float64))
	}

	return a.uuc.UpdateUserRecommend(ctx, &biz.User{
		ID: userId,
	}, req)
}

// RewardList rewardList.
func (a *AppService) RewardList(ctx context.Context, req *v1.RewardListRequest) (*v1.RewardListReply, error) {
	// 在上下文 context 中取出 claims 对象
	var userId int64
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["UserId"] == nil {
			return nil, errors.New(500, "ERROR_TOKEN", "无效TOKEN")
		}
		userId = int64(c["UserId"].(float64))
	}

	return a.uuc.RewardList(ctx, req, &biz.User{
		ID: userId,
	})
}

func (a *AppService) RecommendRewardList(ctx context.Context, req *v1.RecommendRewardListRequest) (*v1.RecommendRewardListReply, error) {
	// 在上下文 context 中取出 claims 对象
	var userId int64
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["UserId"] == nil {
			return nil, errors.New(500, "ERROR_TOKEN", "无效TOKEN")
		}
		userId = int64(c["UserId"].(float64))
	}

	return a.uuc.RecommendRewardList(ctx, &biz.User{
		ID: userId,
	})
}

func (a *AppService) FeeRewardList(ctx context.Context, req *v1.FeeRewardListRequest) (*v1.FeeRewardListReply, error) {
	// 在上下文 context 中取出 claims 对象
	var userId int64
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["UserId"] == nil {
			return nil, errors.New(500, "ERROR_TOKEN", "无效TOKEN")
		}
		userId = int64(c["UserId"].(float64))
	}

	return a.uuc.FeeRewardList(ctx, &biz.User{
		ID: userId,
	})
}

func (a *AppService) WithdrawList(ctx context.Context, req *v1.WithdrawListRequest) (*v1.WithdrawListReply, error) {
	// 在上下文 context 中取出 claims 对象
	var userId int64
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["UserId"] == nil {
			return nil, errors.New(500, "ERROR_TOKEN", "无效TOKEN")
		}
		userId = int64(c["UserId"].(float64))
	}

	return a.uuc.WithdrawList(ctx, &biz.User{
		ID: userId,
	})
}

// Withdraw withdraw.
func (a *AppService) Withdraw(ctx context.Context, req *v1.WithdrawRequest) (*v1.WithdrawReply, error) {
	// 在上下文 context 中取出 claims 对象
	var userId int64
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["UserId"] == nil {
			return nil, errors.New(500, "ERROR_TOKEN", "无效TOKEN")
		}
		userId = int64(c["UserId"].(float64))
	}

	return a.uuc.Withdraw(ctx, req, &biz.User{
		ID: userId,
	})
}

func (a *AppService) AdminRewardList(ctx context.Context, req *v1.AdminRewardListRequest) (*v1.AdminRewardListReply, error) {
	return a.uuc.AdminRewardList(ctx, req)
}

func (a *AppService) AdminUserList(ctx context.Context, req *v1.AdminUserListRequest) (*v1.AdminUserListReply, error) {
	return a.uuc.AdminUserList(ctx, req)
}

func (a *AppService) AdminLocationList(ctx context.Context, req *v1.AdminLocationListRequest) (*v1.AdminLocationListReply, error) {
	return a.uuc.AdminLocationList(ctx, req)
}

func (a *AppService) AdminWithdrawList(ctx context.Context, req *v1.AdminWithdrawListRequest) (*v1.AdminWithdrawListReply, error) {
	return a.uuc.AdminWithdrawList(ctx, req)
}

func (a *AppService) AdminWithdraw(ctx context.Context, req *v1.AdminWithdrawRequest) (*v1.AdminWithdrawReply, error) {
	return a.uuc.AdminWithdraw(ctx, req)
}

func (a *AppService) AdminFee(ctx context.Context, req *v1.AdminFeeRequest) (*v1.AdminFeeReply, error) {
	return a.uuc.AdminFee(ctx, req)
}

func (a *AppService) AdminAll(ctx context.Context, req *v1.AdminAllRequest) (*v1.AdminAllReply, error) {
	return a.uuc.AdminAll(ctx, req)
}

func (a *AppService) AdminUserRecommend(ctx context.Context, req *v1.AdminUserRecommendRequest) (*v1.AdminUserRecommendReply, error) {
	return a.uuc.AdminRecommendList(ctx, req)
}

func (a *AppService) AdminMonthRecommend(ctx context.Context, req *v1.AdminMonthRecommendRequest) (*v1.AdminMonthRecommendReply, error) {
	return a.uuc.AdminMonthRecommend(ctx, req)
}

func (a *AppService) AdminConfig(ctx context.Context, req *v1.AdminConfigRequest) (*v1.AdminConfigReply, error) {
	return a.uuc.AdminConfig(ctx, req)
}

func (a *AppService) AdminConfigUpdate(ctx context.Context, req *v1.AdminConfigUpdateRequest) (*v1.AdminConfigUpdateReply, error) {
	return a.uuc.AdminConfigUpdate(ctx, req)
}

func (a *AppService) AdminWithdrawEth(ctx context.Context, req *v1.AdminWithdrawEthRequest) (*v1.AdminWithdrawEthReply, error) {
	return &v1.AdminWithdrawEthReply{}, nil
}

func (a *AppService) SetAllUserBnbBalance(ctx context.Context, req *v1.SetAllUserBnbBalanceRequest) (*v1.SetAllUserBnbBalanceReply, error) {
	return a.uuc.SetAllUserBnbBalance(ctx, req)
}

func (a *AppService) RewardAllUserBnbBalance(ctx context.Context, req *v1.RewardAllUserBnbBalanceRequest) (*v1.RewardAllUserBnbBalanceReply, error) {
	var (
		users            []*biz.User
		err              error
		buyRewardAmount  float64
		sellRewardAmount float64
		configs          []*biz.Config
	)

	//end := time.Now().Add(3 * time.Minute)
	//for {
	//	if time.Now().After(end) {
	//		return nil, errors.New(500, "reward time error", "分红时间超时")
	//	}
	//
	//	tmpBal := ""
	//	tmpBal, err = balanceAtEth("0x1aDC73f617C7E326BBCe87E99AD35c03f6009861")
	//	if nil != err {
	//		fmt.Println("reward request eth err")
	//		time.Sleep(10 * time.Second)
	//		continue
	//	}
	//	lenBalance := len(tmpBal)
	//	if lenBalance > 8 {
	//		if buyRewardAmount, err = strconv.ParseFloat(tmpBal[:lenBalance-8], 64); err != nil {
	//			return nil, errors.New(500, "reward time error", "分红计算错误")
	//		}
	//
	//		buyRewardAmount /= 10000000000
	//	} else {
	//		buyRewardAmount = 0
	//	}
	//
	//	tmpBal, err = balanceAtEth("0xDf66818ca9FE3a9776B9360a1E6fBCF4e9D82d82")
	//	if nil != err {
	//		fmt.Println("reward request eth err")
	//		time.Sleep(10 * time.Second)
	//		continue
	//	}
	//	lenBalance = len(tmpBal)
	//	if lenBalance > 8 {
	//		if sellRewardAmount, err = strconv.ParseFloat(tmpBal[:lenBalance-8], 64); err != nil {
	//			return nil, errors.New(500, "reward math error", "分红计算错误")
	//		}
	//
	//		sellRewardAmount /= 10000000000
	//	} else {
	//		sellRewardAmount = 0
	//	}
	//
	//	break
	//}

	// 配置
	configs, err = a.uuc.GetDhbConfig(ctx)
	if nil != configs {
		for _, vConfig := range configs {
			if "reward_buy_amount" == vConfig.KeyName {
				buyRewardAmount, _ = strconv.ParseFloat(vConfig.Value, 64)
				buyRewardAmount /= 100
			} else if "reward_sell_amount" == vConfig.KeyName {
				sellRewardAmount, _ = strconv.ParseFloat(vConfig.Value, 64)
				sellRewardAmount /= 100
			}
		}
	}

	if buyRewardAmount <= 0 || sellRewardAmount <= 0 {
		return nil, errors.New(500, "ERROR_TOKEN", "分红金额错误")
	}

	fmt.Println(buyRewardAmount, sellRewardAmount)

	users, err = a.uuc.SelectUsers(ctx)
	if nil != err {
		return nil, err
	}

	usersMap := make(map[int64]string, 0)
	for _, vUsers := range users {
		usersMap[vUsers.ID] = vUsers.Address
	}

	userReward := make(map[int64][]int64, 0)

	for _, vUsers := range users {
		var (
			tmpUserRecommend           *biz.UserRecommend
			tmpUserRecommendLow        []*biz.UserRecommend
			tmpUserRecommendLowUserIds []int64
			tmpBalanceAll              float64
		)
		tmpUserRecommend, err = a.uuc.GetUserRecommend(ctx, vUsers.ID)
		if nil != err {
			return nil, err
		}

		tmpUserRecommendLow, err = a.uuc.GetUserLow(ctx, tmpUserRecommend.RecommendCode+"D"+strconv.FormatInt(vUsers.ID, 10))
		if nil != err {
			return nil, err
		}

		for _, vTmpUserRecommendLow := range tmpUserRecommendLow {
			tmpUserRecommendLowUserIds = append(tmpUserRecommendLowUserIds, vTmpUserRecommendLow.UserId)
		}

		tmpUserRecommendLowUserIds = append(tmpUserRecommendLowUserIds, vUsers.ID)
		tmpBalanceAll, err = a.uuc.GetUserBnbBalance(ctx, tmpUserRecommendLowUserIds)
		if nil != err {
			return nil, err
		}

		tmpBalanceAll, err = strconv.ParseFloat(fmt.Sprintf("%.5f", tmpBalanceAll), 64)
		if nil != err {
			return nil, err
		}

		if tmpBalanceAll > 1000 {
			if _, ok := userReward[1]; !ok {
				userReward[1] = make([]int64, 0)
			}
			userReward[1] = append(userReward[1], vUsers.ID)
		}

		if tmpBalanceAll >= 10001 {
			if _, ok := userReward[2]; !ok {
				userReward[2] = make([]int64, 0)
			}
			userReward[2] = append(userReward[2], vUsers.ID)
		}

		if tmpBalanceAll >= 50001 {
			if _, ok := userReward[3]; !ok {
				userReward[3] = make([]int64, 0)
			}
			userReward[3] = append(userReward[3], vUsers.ID)
		}

		if tmpBalanceAll >= 100001 {
			if _, ok := userReward[4]; !ok {
				userReward[4] = make([]int64, 0)
			}
			userReward[4] = append(userReward[4], vUsers.ID)
		}

		if tmpBalanceAll >= 150001 {
			if _, ok := userReward[5]; !ok {
				userReward[5] = make([]int64, 0)
			}
			userReward[5] = append(userReward[5], vUsers.ID)
		}
	}

	userRewardMap := make(map[int64]float64, 0)
	for k, vUserReward := range userReward {
		var tmpBuyAmount float64
		var tmpSellAmount float64
		if 1 == k {
			//fmt.Println(len(vUserReward))
			tmpBuyAmount = buyRewardAmount / 100 * 10 / float64(len(vUserReward))
			tmpSellAmount = sellRewardAmount / 100 * 20 / float64(len(vUserReward))
		} else if 2 == k {
			//fmt.Println(len(vUserReward))
			tmpBuyAmount = buyRewardAmount / 100 * 15 / float64(len(vUserReward))
			tmpSellAmount = sellRewardAmount / 100 * 20 / float64(len(vUserReward))
		} else if 3 == k {
			//fmt.Println(len(vUserReward))
			tmpBuyAmount = buyRewardAmount / 100 * 20 / float64(len(vUserReward))
			tmpSellAmount = sellRewardAmount / 100 * 20 / float64(len(vUserReward))
		} else if 4 == k {
			//fmt.Println(len(vUserReward))
			tmpBuyAmount = buyRewardAmount / 100 * 25 / float64(len(vUserReward))
			tmpSellAmount = sellRewardAmount / 100 * 20 / float64(len(vUserReward))
		} else if 5 == k {
			//fmt.Println(len(vUserReward))
			tmpBuyAmount = buyRewardAmount / 100 * 30 / float64(len(vUserReward))
			tmpSellAmount = sellRewardAmount / 100 * 20 / float64(len(vUserReward))
		}

		tmpBuyAmount, err = strconv.ParseFloat(fmt.Sprintf("%.7f", tmpBuyAmount), 64)
		if nil != err {
			return nil, err
		}

		tmpSellAmount, err = strconv.ParseFloat(fmt.Sprintf("%.7f", tmpSellAmount), 64)
		if nil != err {
			return nil, err
		}

		for _, vVUserReward := range vUserReward {
			if _, ok := userRewardMap[vVUserReward]; !ok {
				userRewardMap[vVUserReward] = float64(0)
			}

			userRewardMap[vVUserReward] += tmpSellAmount
			userRewardMap[vVUserReward] += tmpBuyAmount
		}
	}

	err = a.uuc.AddUserBnbAmount(ctx, userRewardMap, usersMap)
	if nil != err {
		return nil, err
	}

	return &v1.RewardAllUserBnbBalanceReply{}, nil
}

func (a *AppService) UpdateUserBnbBalance(ctx context.Context, req *v1.UpdateUserBnbBalanceRequest) (*v1.UpdateUserBnbBalanceReply, error) {
	var (
		users []*biz.User
		err   error
	)

	users, err = a.uuc.SelectUsers(ctx)
	if nil != err {
		return nil, err
	}
	tmpErrNum := 0
	tmpSuccessNum := 0
	num := int64(time.Now().Minute())
	//fmt.Println(num, num%4)
	for _, vUsers := range users {
		if vUsers.ID%5 != num%5 {
			continue
		}

		tmpBal := ""
		tmpBal, err = balanceAtEth(vUsers.Address)
		if nil != err {
			//fmt.Println(vUsers.ID, "request bnb4 balance err")
			tmpErrNum++
			continue
		}
		//fmt.Println(num, num%4, tmpBal)

		var currentBalance float64
		lenBalance := len(tmpBal)
		if lenBalance > 14 {
			if currentBalance, err = strconv.ParseFloat(tmpBal[:lenBalance-14], 64); err != nil {
				continue
			}

			currentBalance /= 10000
		} else {
			currentBalance = 0
		}
		_, err = a.uuc.UpdateUserBnbBalance(ctx, vUsers.ID, currentBalance)
		if nil != err {
			fmt.Println(err)
			continue
		}
		tmpSuccessNum++
	}
	if tmpErrNum > 0 {
		fmt.Println(1, tmpErrNum, num, num%5)
	}

	if tmpSuccessNum > 0 {
		fmt.Println(2, tmpSuccessNum, num, num%5)
	}

	return &v1.UpdateUserBnbBalanceReply{}, nil
}

func (a *AppService) UploadRecommendUser(ctx context.Context, req *v1.UploadRecommendUserRequest) (*v1.UploadRecommendUserReply, error) {
	userAddressSlice := make([]string, 0)
	userAddressRecommendSlice := make([]string, 0)

	var err error
	userAddressSlice, userAddressRecommendSlice, err = a.uuc.UploadRecommendUser(ctx, req)
	if err != nil {
		return nil, err
	}
	fmt.Println(userAddressSlice, userAddressRecommendSlice)
	for i := 0; i < len(userAddressSlice); i += 50 {
		_, err = uploadRecommendUserHandle(userAddressSlice[i:i+50], userAddressRecommendSlice[i:i+50])
		if i > 200 {
			break
		}
		if err != nil {
			return nil, err
		}
	}

	return &v1.UploadRecommendUserReply{}, nil
}

func uploadRecommendUserHandle(userAddressSlice, userAddressRecommendSlice []string) (bool, error) {

	client, err := ethclient.Dial("https://data-seed-prebsc-1-s3.binance.org:8545/")
	//client, err := ethclient.Dial("https://bsc-dataseed.binance.org/")
	if err != nil {
		return false, err
	}

	tokenAddress := common.HexToAddress("0xbd3711F867888c9D30E9c67D00F61b33Ce582acB")
	instance, err := NewBnb4(tokenAddress, client)
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	userAddress := make([]common.Address, 0)
	userRecommendAddress := make([]common.Address, 0)

	for _, vUserAddressSlice := range userAddressSlice {
		userAddress = append(userAddress, common.HexToAddress(vUserAddressSlice))
	}
	for _, vUserAddressRecommendSlice := range userAddressRecommendSlice {
		userRecommendAddress = append(userRecommendAddress, common.HexToAddress(vUserAddressRecommendSlice))
	}

	var authUser *bind.TransactOpts
	var chainId = int64(56)

	var privateKey *ecdsa.PrivateKey
	privateKey, err = crypto.HexToECDSA("ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80")
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
		return false, err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
		return false, err
	}

	authUser, err = bind.NewKeyedTransactorWithChainID(privateKey, new(big.Int).SetInt64(chainId))
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	var res *types.Transaction
	res, err = instance.SetUsers(&bind.TransactOpts{
		From:     authUser.From,
		Signer:   authUser.Signer,
		GasPrice: gasPrice,
		GasLimit: 30000000,
	}, userAddress, userRecommendAddress)
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	fmt.Println(res.Hash())

	return true, nil
}

func balanceAtEth(address string) (string, error) {
	var (
		err error
	)
	client, err := ethclient.Dial("https://bsc-dataseed.binance.org/")
	if err != nil {
		return "", err
	}

	tokenAddress := common.HexToAddress("0x0f97F5da8C4715D017F597314DCCd00E0D605Ed8")
	instance, err := NewBnb4(tokenAddress, client)
	if err != nil {
		return "", err
	}

	bal, err := instance.BalanceOf(&bind.CallOpts{}, common.HexToAddress(address))
	if err != nil {
		return "", err
	}

	return bal.String(), nil
}

func GetAmountOut(strAmount string) (string, error) {
	var (
		err error
	)
	client, err := ethclient.Dial("https://bsc-dataseed.binance.org/")
	if err != nil {
		return "", err
	}

	tokenAddress := common.HexToAddress("0x10ED43C718714eb63d5aA57B78B54704E256024E")
	instance, err := NewPancakerouterv2(tokenAddress, client)
	if err != nil {
		return "", err
	}

	addresses := make([]common.Address, 0)
	addresses = append(addresses, common.HexToAddress("0x0f97F5da8C4715D017F597314DCCd00E0D605Ed8"), common.HexToAddress("0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c"))
	amount, _ := new(big.Int).SetString(strAmount, 10)

	bals, err := instance.GetAmountsOut(&bind.CallOpts{}, amount, addresses)
	if err != nil {
		return "", err
	}

	fmt.Println(bals)

	return bals[1].String(), nil
}
