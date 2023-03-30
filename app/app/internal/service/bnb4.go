// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package service

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// Bnb4MetaData contains all meta data concerning the Bnb4 contract.
var Bnb4MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_time\",\"type\":\"uint256\"}],\"name\":\"Launch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REWARD_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_buyMarketRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_buyPoolRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_buyRewardRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_buyRewardRateLevel1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_buyRewardRateLevel2\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_buyRewardRateLevel3\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_buyRewardRateLevel4\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_buyRewardRateLevel5\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_sellMarketRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_sellPoolRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_sellRewardRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_sellRewardRateLevel1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_sellRewardRateLevel2\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_sellRewardRateLevel3\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_sellRewardRateLevel4\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_sellRewardRateLevel5\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_usersReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"ammPairs\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bnb\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUserLow\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUserLowBalanceSum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUserLowNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUserReward\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUserTop\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isBlack\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isLaunch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isWhite\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"removeLiquidityStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reward\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardBuyAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardSellAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"res\",\"type\":\"bool\"}],\"name\":\"setBlack\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bnb\",\"type\":\"address\"}],\"name\":\"setBnb\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"res\",\"type\":\"uint256\"}],\"name\":\"setBuyRewardRateLevel1\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"res\",\"type\":\"uint256\"}],\"name\":\"setBuyRewardRateLevel2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"res\",\"type\":\"uint256\"}],\"name\":\"setBuyRewardRateLevel3\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"res\",\"type\":\"uint256\"}],\"name\":\"setBuyRewardRateLevel4\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"res\",\"type\":\"uint256\"}],\"name\":\"setBuyRewardRateLevel5\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recommendAccount\",\"type\":\"address\"}],\"name\":\"setExistsUserByAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"launch_\",\"type\":\"bool\"}],\"name\":\"setLaunch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"setRemoveLiquidityStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"res\",\"type\":\"uint256\"}],\"name\":\"setSellMarketRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"res\",\"type\":\"uint256\"}],\"name\":\"setSellPooltRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"res\",\"type\":\"uint256\"}],\"name\":\"setSellRewardRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"res\",\"type\":\"uint256\"}],\"name\":\"setSellRewardRateLevel1\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"res\",\"type\":\"uint256\"}],\"name\":\"setSellRewardRateLevel2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"res\",\"type\":\"uint256\"}],\"name\":\"setSellRewardRateLevel3\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"res\",\"type\":\"uint256\"}],\"name\":\"setSellRewardRateLevel4\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"res\",\"type\":\"uint256\"}],\"name\":\"setSellRewardRateLevel5\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"setUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"account\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"accountRecommend\",\"type\":\"address[]\"}],\"name\":\"setUsers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"setUsersFirst\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"res\",\"type\":\"bool\"}],\"name\":\"setWhite\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"res\",\"type\":\"uint256\"}],\"name\":\"setbuyMarketRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"res\",\"type\":\"uint256\"}],\"name\":\"setbuyPoolRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"res\",\"type\":\"uint256\"}],\"name\":\"setbuyRewardRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uniswapV2Pair\",\"outputs\":[{\"internalType\":\"contractIUniswapV2Pair\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uniswapV2Router\",\"outputs\":[{\"internalType\":\"contractIUniswapV2Router02\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// Bnb4ABI is the input ABI used to generate the binding from.
// Deprecated: Use Bnb4MetaData.ABI instead.
var Bnb4ABI = Bnb4MetaData.ABI

// Bnb4 is an auto generated Go binding around an Ethereum contract.
type Bnb4 struct {
	Bnb4Caller     // Read-only binding to the contract
	Bnb4Transactor // Write-only binding to the contract
	Bnb4Filterer   // Log filterer for contract events
}

// Bnb4Caller is an auto generated read-only Go binding around an Ethereum contract.
type Bnb4Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Bnb4Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Bnb4Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Bnb4Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Bnb4Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Bnb4Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Bnb4Session struct {
	Contract     *Bnb4             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Bnb4CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Bnb4CallerSession struct {
	Contract *Bnb4Caller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// Bnb4TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Bnb4TransactorSession struct {
	Contract     *Bnb4Transactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Bnb4Raw is an auto generated low-level Go binding around an Ethereum contract.
type Bnb4Raw struct {
	Contract *Bnb4 // Generic contract binding to access the raw methods on
}

// Bnb4CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Bnb4CallerRaw struct {
	Contract *Bnb4Caller // Generic read-only contract binding to access the raw methods on
}

// Bnb4TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Bnb4TransactorRaw struct {
	Contract *Bnb4Transactor // Generic write-only contract binding to access the raw methods on
}

// NewBnb4 creates a new instance of Bnb4, bound to a specific deployed contract.
func NewBnb4(address common.Address, backend bind.ContractBackend) (*Bnb4, error) {
	contract, err := bindBnb4(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bnb4{Bnb4Caller: Bnb4Caller{contract: contract}, Bnb4Transactor: Bnb4Transactor{contract: contract}, Bnb4Filterer: Bnb4Filterer{contract: contract}}, nil
}

// NewBnb4Caller creates a new read-only instance of Bnb4, bound to a specific deployed contract.
func NewBnb4Caller(address common.Address, caller bind.ContractCaller) (*Bnb4Caller, error) {
	contract, err := bindBnb4(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Bnb4Caller{contract: contract}, nil
}

// NewBnb4Transactor creates a new write-only instance of Bnb4, bound to a specific deployed contract.
func NewBnb4Transactor(address common.Address, transactor bind.ContractTransactor) (*Bnb4Transactor, error) {
	contract, err := bindBnb4(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Bnb4Transactor{contract: contract}, nil
}

// NewBnb4Filterer creates a new log filterer instance of Bnb4, bound to a specific deployed contract.
func NewBnb4Filterer(address common.Address, filterer bind.ContractFilterer) (*Bnb4Filterer, error) {
	contract, err := bindBnb4(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Bnb4Filterer{contract: contract}, nil
}

// bindBnb4 binds a generic wrapper to an already deployed contract.
func bindBnb4(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Bnb4ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bnb4 *Bnb4Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bnb4.Contract.Bnb4Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bnb4 *Bnb4Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bnb4.Contract.Bnb4Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bnb4 *Bnb4Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bnb4.Contract.Bnb4Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bnb4 *Bnb4CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bnb4.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bnb4 *Bnb4TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bnb4.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bnb4 *Bnb4TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bnb4.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Bnb4 *Bnb4Caller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bnb4.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Bnb4 *Bnb4Session) DEFAULTADMINROLE() ([32]byte, error) {
	return _Bnb4.Contract.DEFAULTADMINROLE(&_Bnb4.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Bnb4 *Bnb4CallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Bnb4.Contract.DEFAULTADMINROLE(&_Bnb4.CallOpts)
}

// REWARDROLE is a free data retrieval call binding the contract method 0xe8884f2b.
//
// Solidity: function REWARD_ROLE() view returns(bytes32)
func (_Bnb4 *Bnb4Caller) REWARDROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bnb4.contract.Call(opts, &out, "REWARD_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// REWARDROLE is a free data retrieval call binding the contract method 0xe8884f2b.
//
// Solidity: function REWARD_ROLE() view returns(bytes32)
func (_Bnb4 *Bnb4Session) REWARDROLE() ([32]byte, error) {
	return _Bnb4.Contract.REWARDROLE(&_Bnb4.CallOpts)
}

// REWARDROLE is a free data retrieval call binding the contract method 0xe8884f2b.
//
// Solidity: function REWARD_ROLE() view returns(bytes32)
func (_Bnb4 *Bnb4CallerSession) REWARDROLE() ([32]byte, error) {
	return _Bnb4.Contract.REWARDROLE(&_Bnb4.CallOpts)
}

// BuyMarketRate is a free data retrieval call binding the contract method 0xcab6cb32.
//
// Solidity: function _buyMarketRate() view returns(uint256)
func (_Bnb4 *Bnb4Caller) BuyMarketRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bnb4.contract.Call(opts, &out, "_buyMarketRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BuyMarketRate is a free data retrieval call binding the contract method 0xcab6cb32.
//
// Solidity: function _buyMarketRate() view returns(uint256)
func (_Bnb4 *Bnb4Session) BuyMarketRate() (*big.Int, error) {
	return _Bnb4.Contract.BuyMarketRate(&_Bnb4.CallOpts)
}

// BuyMarketRate is a free data retrieval call binding the contract method 0xcab6cb32.
//
// Solidity: function _buyMarketRate() view returns(uint256)
func (_Bnb4 *Bnb4CallerSession) BuyMarketRate() (*big.Int, error) {
	return _Bnb4.Contract.BuyMarketRate(&_Bnb4.CallOpts)
}

// BuyPoolRate is a free data retrieval call binding the contract method 0xfaf2437e.
//
// Solidity: function _buyPoolRate() view returns(uint256)
func (_Bnb4 *Bnb4Caller) BuyPoolRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bnb4.contract.Call(opts, &out, "_buyPoolRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BuyPoolRate is a free data retrieval call binding the contract method 0xfaf2437e.
//
// Solidity: function _buyPoolRate() view returns(uint256)
func (_Bnb4 *Bnb4Session) BuyPoolRate() (*big.Int, error) {
	return _Bnb4.Contract.BuyPoolRate(&_Bnb4.CallOpts)
}

// BuyPoolRate is a free data retrieval call binding the contract method 0xfaf2437e.
//
// Solidity: function _buyPoolRate() view returns(uint256)
func (_Bnb4 *Bnb4CallerSession) BuyPoolRate() (*big.Int, error) {
	return _Bnb4.Contract.BuyPoolRate(&_Bnb4.CallOpts)
}

// BuyRewardRate is a free data retrieval call binding the contract method 0x0765fbc8.
//
// Solidity: function _buyRewardRate() view returns(uint256)
func (_Bnb4 *Bnb4Caller) BuyRewardRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bnb4.contract.Call(opts, &out, "_buyRewardRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BuyRewardRate is a free data retrieval call binding the contract method 0x0765fbc8.
//
// Solidity: function _buyRewardRate() view returns(uint256)
func (_Bnb4 *Bnb4Session) BuyRewardRate() (*big.Int, error) {
	return _Bnb4.Contract.BuyRewardRate(&_Bnb4.CallOpts)
}

// BuyRewardRate is a free data retrieval call binding the contract method 0x0765fbc8.
//
// Solidity: function _buyRewardRate() view returns(uint256)
func (_Bnb4 *Bnb4CallerSession) BuyRewardRate() (*big.Int, error) {
	return _Bnb4.Contract.BuyRewardRate(&_Bnb4.CallOpts)
}

// BuyRewardRateLevel1 is a free data retrieval call binding the contract method 0xa5e45c8c.
//
// Solidity: function _buyRewardRateLevel1() view returns(uint256)
func (_Bnb4 *Bnb4Caller) BuyRewardRateLevel1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bnb4.contract.Call(opts, &out, "_buyRewardRateLevel1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BuyRewardRateLevel1 is a free data retrieval call binding the contract method 0xa5e45c8c.
//
// Solidity: function _buyRewardRateLevel1() view returns(uint256)
func (_Bnb4 *Bnb4Session) BuyRewardRateLevel1() (*big.Int, error) {
	return _Bnb4.Contract.BuyRewardRateLevel1(&_Bnb4.CallOpts)
}

// BuyRewardRateLevel1 is a free data retrieval call binding the contract method 0xa5e45c8c.
//
// Solidity: function _buyRewardRateLevel1() view returns(uint256)
func (_Bnb4 *Bnb4CallerSession) BuyRewardRateLevel1() (*big.Int, error) {
	return _Bnb4.Contract.BuyRewardRateLevel1(&_Bnb4.CallOpts)
}

// BuyRewardRateLevel2 is a free data retrieval call binding the contract method 0x8469146f.
//
// Solidity: function _buyRewardRateLevel2() view returns(uint256)
func (_Bnb4 *Bnb4Caller) BuyRewardRateLevel2(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bnb4.contract.Call(opts, &out, "_buyRewardRateLevel2")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BuyRewardRateLevel2 is a free data retrieval call binding the contract method 0x8469146f.
//
// Solidity: function _buyRewardRateLevel2() view returns(uint256)
func (_Bnb4 *Bnb4Session) BuyRewardRateLevel2() (*big.Int, error) {
	return _Bnb4.Contract.BuyRewardRateLevel2(&_Bnb4.CallOpts)
}

// BuyRewardRateLevel2 is a free data retrieval call binding the contract method 0x8469146f.
//
// Solidity: function _buyRewardRateLevel2() view returns(uint256)
func (_Bnb4 *Bnb4CallerSession) BuyRewardRateLevel2() (*big.Int, error) {
	return _Bnb4.Contract.BuyRewardRateLevel2(&_Bnb4.CallOpts)
}

// BuyRewardRateLevel3 is a free data retrieval call binding the contract method 0x70c04159.
//
// Solidity: function _buyRewardRateLevel3() view returns(uint256)
func (_Bnb4 *Bnb4Caller) BuyRewardRateLevel3(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bnb4.contract.Call(opts, &out, "_buyRewardRateLevel3")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BuyRewardRateLevel3 is a free data retrieval call binding the contract method 0x70c04159.
//
// Solidity: function _buyRewardRateLevel3() view returns(uint256)
func (_Bnb4 *Bnb4Session) BuyRewardRateLevel3() (*big.Int, error) {
	return _Bnb4.Contract.BuyRewardRateLevel3(&_Bnb4.CallOpts)
}

// BuyRewardRateLevel3 is a free data retrieval call binding the contract method 0x70c04159.
//
// Solidity: function _buyRewardRateLevel3() view returns(uint256)
func (_Bnb4 *Bnb4CallerSession) BuyRewardRateLevel3() (*big.Int, error) {
	return _Bnb4.Contract.BuyRewardRateLevel3(&_Bnb4.CallOpts)
}

// BuyRewardRateLevel4 is a free data retrieval call binding the contract method 0x4478a1c0.
//
// Solidity: function _buyRewardRateLevel4() view returns(uint256)
func (_Bnb4 *Bnb4Caller) BuyRewardRateLevel4(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bnb4.contract.Call(opts, &out, "_buyRewardRateLevel4")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BuyRewardRateLevel4 is a free data retrieval call binding the contract method 0x4478a1c0.
//
// Solidity: function _buyRewardRateLevel4() view returns(uint256)
func (_Bnb4 *Bnb4Session) BuyRewardRateLevel4() (*big.Int, error) {
	return _Bnb4.Contract.BuyRewardRateLevel4(&_Bnb4.CallOpts)
}

// BuyRewardRateLevel4 is a free data retrieval call binding the contract method 0x4478a1c0.
//
// Solidity: function _buyRewardRateLevel4() view returns(uint256)
func (_Bnb4 *Bnb4CallerSession) BuyRewardRateLevel4() (*big.Int, error) {
	return _Bnb4.Contract.BuyRewardRateLevel4(&_Bnb4.CallOpts)
}

// BuyRewardRateLevel5 is a free data retrieval call binding the contract method 0x5be07577.
//
// Solidity: function _buyRewardRateLevel5() view returns(uint256)
func (_Bnb4 *Bnb4Caller) BuyRewardRateLevel5(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bnb4.contract.Call(opts, &out, "_buyRewardRateLevel5")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BuyRewardRateLevel5 is a free data retrieval call binding the contract method 0x5be07577.
//
// Solidity: function _buyRewardRateLevel5() view returns(uint256)
func (_Bnb4 *Bnb4Session) BuyRewardRateLevel5() (*big.Int, error) {
	return _Bnb4.Contract.BuyRewardRateLevel5(&_Bnb4.CallOpts)
}

// BuyRewardRateLevel5 is a free data retrieval call binding the contract method 0x5be07577.
//
// Solidity: function _buyRewardRateLevel5() view returns(uint256)
func (_Bnb4 *Bnb4CallerSession) BuyRewardRateLevel5() (*big.Int, error) {
	return _Bnb4.Contract.BuyRewardRateLevel5(&_Bnb4.CallOpts)
}

// SellMarketRate is a free data retrieval call binding the contract method 0xbfb07d8b.
//
// Solidity: function _sellMarketRate() view returns(uint256)
func (_Bnb4 *Bnb4Caller) SellMarketRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bnb4.contract.Call(opts, &out, "_sellMarketRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SellMarketRate is a free data retrieval call binding the contract method 0xbfb07d8b.
//
// Solidity: function _sellMarketRate() view returns(uint256)
func (_Bnb4 *Bnb4Session) SellMarketRate() (*big.Int, error) {
	return _Bnb4.Contract.SellMarketRate(&_Bnb4.CallOpts)
}

// SellMarketRate is a free data retrieval call binding the contract method 0xbfb07d8b.
//
// Solidity: function _sellMarketRate() view returns(uint256)
func (_Bnb4 *Bnb4CallerSession) SellMarketRate() (*big.Int, error) {
	return _Bnb4.Contract.SellMarketRate(&_Bnb4.CallOpts)
}

// SellPoolRate is a free data retrieval call binding the contract method 0x6e986dcc.
//
// Solidity: function _sellPoolRate() view returns(uint256)
func (_Bnb4 *Bnb4Caller) SellPoolRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bnb4.contract.Call(opts, &out, "_sellPoolRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SellPoolRate is a free data retrieval call binding the contract method 0x6e986dcc.
//
// Solidity: function _sellPoolRate() view returns(uint256)
func (_Bnb4 *Bnb4Session) SellPoolRate() (*big.Int, error) {
	return _Bnb4.Contract.SellPoolRate(&_Bnb4.CallOpts)
}

// SellPoolRate is a free data retrieval call binding the contract method 0x6e986dcc.
//
// Solidity: function _sellPoolRate() view returns(uint256)
func (_Bnb4 *Bnb4CallerSession) SellPoolRate() (*big.Int, error) {
	return _Bnb4.Contract.SellPoolRate(&_Bnb4.CallOpts)
}

// SellRewardRate is a free data retrieval call binding the contract method 0xa0e84f29.
//
// Solidity: function _sellRewardRate() view returns(uint256)
func (_Bnb4 *Bnb4Caller) SellRewardRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bnb4.contract.Call(opts, &out, "_sellRewardRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SellRewardRate is a free data retrieval call binding the contract method 0xa0e84f29.
//
// Solidity: function _sellRewardRate() view returns(uint256)
func (_Bnb4 *Bnb4Session) SellRewardRate() (*big.Int, error) {
	return _Bnb4.Contract.SellRewardRate(&_Bnb4.CallOpts)
}

// SellRewardRate is a free data retrieval call binding the contract method 0xa0e84f29.
//
// Solidity: function _sellRewardRate() view returns(uint256)
func (_Bnb4 *Bnb4CallerSession) SellRewardRate() (*big.Int, error) {
	return _Bnb4.Contract.SellRewardRate(&_Bnb4.CallOpts)
}

// SellRewardRateLevel1 is a free data retrieval call binding the contract method 0x0a401d3f.
//
// Solidity: function _sellRewardRateLevel1() view returns(uint256)
func (_Bnb4 *Bnb4Caller) SellRewardRateLevel1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bnb4.contract.Call(opts, &out, "_sellRewardRateLevel1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SellRewardRateLevel1 is a free data retrieval call binding the contract method 0x0a401d3f.
//
// Solidity: function _sellRewardRateLevel1() view returns(uint256)
func (_Bnb4 *Bnb4Session) SellRewardRateLevel1() (*big.Int, error) {
	return _Bnb4.Contract.SellRewardRateLevel1(&_Bnb4.CallOpts)
}

// SellRewardRateLevel1 is a free data retrieval call binding the contract method 0x0a401d3f.
//
// Solidity: function _sellRewardRateLevel1() view returns(uint256)
func (_Bnb4 *Bnb4CallerSession) SellRewardRateLevel1() (*big.Int, error) {
	return _Bnb4.Contract.SellRewardRateLevel1(&_Bnb4.CallOpts)
}

// SellRewardRateLevel2 is a free data retrieval call binding the contract method 0xbb574487.
//
// Solidity: function _sellRewardRateLevel2() view returns(uint256)
func (_Bnb4 *Bnb4Caller) SellRewardRateLevel2(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bnb4.contract.Call(opts, &out, "_sellRewardRateLevel2")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SellRewardRateLevel2 is a free data retrieval call binding the contract method 0xbb574487.
//
// Solidity: function _sellRewardRateLevel2() view returns(uint256)
func (_Bnb4 *Bnb4Session) SellRewardRateLevel2() (*big.Int, error) {
	return _Bnb4.Contract.SellRewardRateLevel2(&_Bnb4.CallOpts)
}

// SellRewardRateLevel2 is a free data retrieval call binding the contract method 0xbb574487.
//
// Solidity: function _sellRewardRateLevel2() view returns(uint256)
func (_Bnb4 *Bnb4CallerSession) SellRewardRateLevel2() (*big.Int, error) {
	return _Bnb4.Contract.SellRewardRateLevel2(&_Bnb4.CallOpts)
}

// SellRewardRateLevel3 is a free data retrieval call binding the contract method 0x7489aab4.
//
// Solidity: function _sellRewardRateLevel3() view returns(uint256)
func (_Bnb4 *Bnb4Caller) SellRewardRateLevel3(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bnb4.contract.Call(opts, &out, "_sellRewardRateLevel3")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SellRewardRateLevel3 is a free data retrieval call binding the contract method 0x7489aab4.
//
// Solidity: function _sellRewardRateLevel3() view returns(uint256)
func (_Bnb4 *Bnb4Session) SellRewardRateLevel3() (*big.Int, error) {
	return _Bnb4.Contract.SellRewardRateLevel3(&_Bnb4.CallOpts)
}

// SellRewardRateLevel3 is a free data retrieval call binding the contract method 0x7489aab4.
//
// Solidity: function _sellRewardRateLevel3() view returns(uint256)
func (_Bnb4 *Bnb4CallerSession) SellRewardRateLevel3() (*big.Int, error) {
	return _Bnb4.Contract.SellRewardRateLevel3(&_Bnb4.CallOpts)
}

// SellRewardRateLevel4 is a free data retrieval call binding the contract method 0x6a4212bc.
//
// Solidity: function _sellRewardRateLevel4() view returns(uint256)
func (_Bnb4 *Bnb4Caller) SellRewardRateLevel4(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bnb4.contract.Call(opts, &out, "_sellRewardRateLevel4")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SellRewardRateLevel4 is a free data retrieval call binding the contract method 0x6a4212bc.
//
// Solidity: function _sellRewardRateLevel4() view returns(uint256)
func (_Bnb4 *Bnb4Session) SellRewardRateLevel4() (*big.Int, error) {
	return _Bnb4.Contract.SellRewardRateLevel4(&_Bnb4.CallOpts)
}

// SellRewardRateLevel4 is a free data retrieval call binding the contract method 0x6a4212bc.
//
// Solidity: function _sellRewardRateLevel4() view returns(uint256)
func (_Bnb4 *Bnb4CallerSession) SellRewardRateLevel4() (*big.Int, error) {
	return _Bnb4.Contract.SellRewardRateLevel4(&_Bnb4.CallOpts)
}

// SellRewardRateLevel5 is a free data retrieval call binding the contract method 0x3298322a.
//
// Solidity: function _sellRewardRateLevel5() view returns(uint256)
func (_Bnb4 *Bnb4Caller) SellRewardRateLevel5(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bnb4.contract.Call(opts, &out, "_sellRewardRateLevel5")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SellRewardRateLevel5 is a free data retrieval call binding the contract method 0x3298322a.
//
// Solidity: function _sellRewardRateLevel5() view returns(uint256)
func (_Bnb4 *Bnb4Session) SellRewardRateLevel5() (*big.Int, error) {
	return _Bnb4.Contract.SellRewardRateLevel5(&_Bnb4.CallOpts)
}

// SellRewardRateLevel5 is a free data retrieval call binding the contract method 0x3298322a.
//
// Solidity: function _sellRewardRateLevel5() view returns(uint256)
func (_Bnb4 *Bnb4CallerSession) SellRewardRateLevel5() (*big.Int, error) {
	return _Bnb4.Contract.SellRewardRateLevel5(&_Bnb4.CallOpts)
}

// UsersReward is a free data retrieval call binding the contract method 0x7e916821.
//
// Solidity: function _usersReward(address , uint256 ) view returns(uint256)
func (_Bnb4 *Bnb4Caller) UsersReward(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Bnb4.contract.Call(opts, &out, "_usersReward", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UsersReward is a free data retrieval call binding the contract method 0x7e916821.
//
// Solidity: function _usersReward(address , uint256 ) view returns(uint256)
func (_Bnb4 *Bnb4Session) UsersReward(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Bnb4.Contract.UsersReward(&_Bnb4.CallOpts, arg0, arg1)
}

// UsersReward is a free data retrieval call binding the contract method 0x7e916821.
//
// Solidity: function _usersReward(address , uint256 ) view returns(uint256)
func (_Bnb4 *Bnb4CallerSession) UsersReward(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Bnb4.Contract.UsersReward(&_Bnb4.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Bnb4 *Bnb4Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bnb4.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Bnb4 *Bnb4Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Bnb4.Contract.Allowance(&_Bnb4.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Bnb4 *Bnb4CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Bnb4.Contract.Allowance(&_Bnb4.CallOpts, owner, spender)
}

// AmmPairs is a free data retrieval call binding the contract method 0xa72905a2.
//
// Solidity: function ammPairs(address ) view returns(bool)
func (_Bnb4 *Bnb4Caller) AmmPairs(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Bnb4.contract.Call(opts, &out, "ammPairs", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AmmPairs is a free data retrieval call binding the contract method 0xa72905a2.
//
// Solidity: function ammPairs(address ) view returns(bool)
func (_Bnb4 *Bnb4Session) AmmPairs(arg0 common.Address) (bool, error) {
	return _Bnb4.Contract.AmmPairs(&_Bnb4.CallOpts, arg0)
}

// AmmPairs is a free data retrieval call binding the contract method 0xa72905a2.
//
// Solidity: function ammPairs(address ) view returns(bool)
func (_Bnb4 *Bnb4CallerSession) AmmPairs(arg0 common.Address) (bool, error) {
	return _Bnb4.Contract.AmmPairs(&_Bnb4.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Bnb4 *Bnb4Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bnb4.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Bnb4 *Bnb4Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _Bnb4.Contract.BalanceOf(&_Bnb4.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Bnb4 *Bnb4CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Bnb4.Contract.BalanceOf(&_Bnb4.CallOpts, account)
}

// Bnb is a free data retrieval call binding the contract method 0x49a876e4.
//
// Solidity: function bnb() view returns(address)
func (_Bnb4 *Bnb4Caller) Bnb(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bnb4.contract.Call(opts, &out, "bnb")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bnb is a free data retrieval call binding the contract method 0x49a876e4.
//
// Solidity: function bnb() view returns(address)
func (_Bnb4 *Bnb4Session) Bnb() (common.Address, error) {
	return _Bnb4.Contract.Bnb(&_Bnb4.CallOpts)
}

// Bnb is a free data retrieval call binding the contract method 0x49a876e4.
//
// Solidity: function bnb() view returns(address)
func (_Bnb4 *Bnb4CallerSession) Bnb() (common.Address, error) {
	return _Bnb4.Contract.Bnb(&_Bnb4.CallOpts)
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() view returns(uint256)
func (_Bnb4 *Bnb4Caller) Cap(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bnb4.contract.Call(opts, &out, "cap")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() view returns(uint256)
func (_Bnb4 *Bnb4Session) Cap() (*big.Int, error) {
	return _Bnb4.Contract.Cap(&_Bnb4.CallOpts)
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() view returns(uint256)
func (_Bnb4 *Bnb4CallerSession) Cap() (*big.Int, error) {
	return _Bnb4.Contract.Cap(&_Bnb4.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Bnb4 *Bnb4Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Bnb4.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Bnb4 *Bnb4Session) Decimals() (uint8, error) {
	return _Bnb4.Contract.Decimals(&_Bnb4.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Bnb4 *Bnb4CallerSession) Decimals() (uint8, error) {
	return _Bnb4.Contract.Decimals(&_Bnb4.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Bnb4 *Bnb4Caller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Bnb4.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Bnb4 *Bnb4Session) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Bnb4.Contract.GetRoleAdmin(&_Bnb4.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Bnb4 *Bnb4CallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Bnb4.Contract.GetRoleAdmin(&_Bnb4.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Bnb4 *Bnb4Caller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Bnb4.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Bnb4 *Bnb4Session) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Bnb4.Contract.GetRoleMember(&_Bnb4.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Bnb4 *Bnb4CallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Bnb4.Contract.GetRoleMember(&_Bnb4.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Bnb4 *Bnb4Caller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Bnb4.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Bnb4 *Bnb4Session) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Bnb4.Contract.GetRoleMemberCount(&_Bnb4.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Bnb4 *Bnb4CallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Bnb4.Contract.GetRoleMemberCount(&_Bnb4.CallOpts, role)
}

// GetUserLow is a free data retrieval call binding the contract method 0x23f2d822.
//
// Solidity: function getUserLow() view returns(address[])
func (_Bnb4 *Bnb4Caller) GetUserLow(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Bnb4.contract.Call(opts, &out, "getUserLow")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetUserLow is a free data retrieval call binding the contract method 0x23f2d822.
//
// Solidity: function getUserLow() view returns(address[])
func (_Bnb4 *Bnb4Session) GetUserLow() ([]common.Address, error) {
	return _Bnb4.Contract.GetUserLow(&_Bnb4.CallOpts)
}

// GetUserLow is a free data retrieval call binding the contract method 0x23f2d822.
//
// Solidity: function getUserLow() view returns(address[])
func (_Bnb4 *Bnb4CallerSession) GetUserLow() ([]common.Address, error) {
	return _Bnb4.Contract.GetUserLow(&_Bnb4.CallOpts)
}

// GetUserLowBalanceSum is a free data retrieval call binding the contract method 0x03d9c08f.
//
// Solidity: function getUserLowBalanceSum() view returns(uint256)
func (_Bnb4 *Bnb4Caller) GetUserLowBalanceSum(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bnb4.contract.Call(opts, &out, "getUserLowBalanceSum")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserLowBalanceSum is a free data retrieval call binding the contract method 0x03d9c08f.
//
// Solidity: function getUserLowBalanceSum() view returns(uint256)
func (_Bnb4 *Bnb4Session) GetUserLowBalanceSum() (*big.Int, error) {
	return _Bnb4.Contract.GetUserLowBalanceSum(&_Bnb4.CallOpts)
}

// GetUserLowBalanceSum is a free data retrieval call binding the contract method 0x03d9c08f.
//
// Solidity: function getUserLowBalanceSum() view returns(uint256)
func (_Bnb4 *Bnb4CallerSession) GetUserLowBalanceSum() (*big.Int, error) {
	return _Bnb4.Contract.GetUserLowBalanceSum(&_Bnb4.CallOpts)
}

// GetUserLowNum is a free data retrieval call binding the contract method 0xf58de136.
//
// Solidity: function getUserLowNum() view returns(uint256)
func (_Bnb4 *Bnb4Caller) GetUserLowNum(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bnb4.contract.Call(opts, &out, "getUserLowNum")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserLowNum is a free data retrieval call binding the contract method 0xf58de136.
//
// Solidity: function getUserLowNum() view returns(uint256)
func (_Bnb4 *Bnb4Session) GetUserLowNum() (*big.Int, error) {
	return _Bnb4.Contract.GetUserLowNum(&_Bnb4.CallOpts)
}

// GetUserLowNum is a free data retrieval call binding the contract method 0xf58de136.
//
// Solidity: function getUserLowNum() view returns(uint256)
func (_Bnb4 *Bnb4CallerSession) GetUserLowNum() (*big.Int, error) {
	return _Bnb4.Contract.GetUserLowNum(&_Bnb4.CallOpts)
}

// GetUserReward is a free data retrieval call binding the contract method 0xd793d83e.
//
// Solidity: function getUserReward() view returns(uint256[])
func (_Bnb4 *Bnb4Caller) GetUserReward(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _Bnb4.contract.Call(opts, &out, "getUserReward")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetUserReward is a free data retrieval call binding the contract method 0xd793d83e.
//
// Solidity: function getUserReward() view returns(uint256[])
func (_Bnb4 *Bnb4Session) GetUserReward() ([]*big.Int, error) {
	return _Bnb4.Contract.GetUserReward(&_Bnb4.CallOpts)
}

// GetUserReward is a free data retrieval call binding the contract method 0xd793d83e.
//
// Solidity: function getUserReward() view returns(uint256[])
func (_Bnb4 *Bnb4CallerSession) GetUserReward() ([]*big.Int, error) {
	return _Bnb4.Contract.GetUserReward(&_Bnb4.CallOpts)
}

// GetUserTop is a free data retrieval call binding the contract method 0x35835618.
//
// Solidity: function getUserTop() view returns(address)
func (_Bnb4 *Bnb4Caller) GetUserTop(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bnb4.contract.Call(opts, &out, "getUserTop")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetUserTop is a free data retrieval call binding the contract method 0x35835618.
//
// Solidity: function getUserTop() view returns(address)
func (_Bnb4 *Bnb4Session) GetUserTop() (common.Address, error) {
	return _Bnb4.Contract.GetUserTop(&_Bnb4.CallOpts)
}

// GetUserTop is a free data retrieval call binding the contract method 0x35835618.
//
// Solidity: function getUserTop() view returns(address)
func (_Bnb4 *Bnb4CallerSession) GetUserTop() (common.Address, error) {
	return _Bnb4.Contract.GetUserTop(&_Bnb4.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Bnb4 *Bnb4Caller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Bnb4.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Bnb4 *Bnb4Session) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Bnb4.Contract.HasRole(&_Bnb4.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Bnb4 *Bnb4CallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Bnb4.Contract.HasRole(&_Bnb4.CallOpts, role, account)
}

// IsBlack is a free data retrieval call binding the contract method 0x332daccf.
//
// Solidity: function isBlack(address account) view returns(bool)
func (_Bnb4 *Bnb4Caller) IsBlack(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _Bnb4.contract.Call(opts, &out, "isBlack", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBlack is a free data retrieval call binding the contract method 0x332daccf.
//
// Solidity: function isBlack(address account) view returns(bool)
func (_Bnb4 *Bnb4Session) IsBlack(account common.Address) (bool, error) {
	return _Bnb4.Contract.IsBlack(&_Bnb4.CallOpts, account)
}

// IsBlack is a free data retrieval call binding the contract method 0x332daccf.
//
// Solidity: function isBlack(address account) view returns(bool)
func (_Bnb4 *Bnb4CallerSession) IsBlack(account common.Address) (bool, error) {
	return _Bnb4.Contract.IsBlack(&_Bnb4.CallOpts, account)
}

// IsLaunch is a free data retrieval call binding the contract method 0xf5d588f5.
//
// Solidity: function isLaunch() view returns(bool)
func (_Bnb4 *Bnb4Caller) IsLaunch(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Bnb4.contract.Call(opts, &out, "isLaunch")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLaunch is a free data retrieval call binding the contract method 0xf5d588f5.
//
// Solidity: function isLaunch() view returns(bool)
func (_Bnb4 *Bnb4Session) IsLaunch() (bool, error) {
	return _Bnb4.Contract.IsLaunch(&_Bnb4.CallOpts)
}

// IsLaunch is a free data retrieval call binding the contract method 0xf5d588f5.
//
// Solidity: function isLaunch() view returns(bool)
func (_Bnb4 *Bnb4CallerSession) IsLaunch() (bool, error) {
	return _Bnb4.Contract.IsLaunch(&_Bnb4.CallOpts)
}

// IsWhite is a free data retrieval call binding the contract method 0xa348c289.
//
// Solidity: function isWhite(address account) view returns(bool)
func (_Bnb4 *Bnb4Caller) IsWhite(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _Bnb4.contract.Call(opts, &out, "isWhite", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWhite is a free data retrieval call binding the contract method 0xa348c289.
//
// Solidity: function isWhite(address account) view returns(bool)
func (_Bnb4 *Bnb4Session) IsWhite(account common.Address) (bool, error) {
	return _Bnb4.Contract.IsWhite(&_Bnb4.CallOpts, account)
}

// IsWhite is a free data retrieval call binding the contract method 0xa348c289.
//
// Solidity: function isWhite(address account) view returns(bool)
func (_Bnb4 *Bnb4CallerSession) IsWhite(account common.Address) (bool, error) {
	return _Bnb4.Contract.IsWhite(&_Bnb4.CallOpts, account)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Bnb4 *Bnb4Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Bnb4.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Bnb4 *Bnb4Session) Name() (string, error) {
	return _Bnb4.Contract.Name(&_Bnb4.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Bnb4 *Bnb4CallerSession) Name() (string, error) {
	return _Bnb4.Contract.Name(&_Bnb4.CallOpts)
}

// RemoveLiquidityStatus is a free data retrieval call binding the contract method 0xf030c51b.
//
// Solidity: function removeLiquidityStatus() view returns(bool)
func (_Bnb4 *Bnb4Caller) RemoveLiquidityStatus(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Bnb4.contract.Call(opts, &out, "removeLiquidityStatus")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RemoveLiquidityStatus is a free data retrieval call binding the contract method 0xf030c51b.
//
// Solidity: function removeLiquidityStatus() view returns(bool)
func (_Bnb4 *Bnb4Session) RemoveLiquidityStatus() (bool, error) {
	return _Bnb4.Contract.RemoveLiquidityStatus(&_Bnb4.CallOpts)
}

// RemoveLiquidityStatus is a free data retrieval call binding the contract method 0xf030c51b.
//
// Solidity: function removeLiquidityStatus() view returns(bool)
func (_Bnb4 *Bnb4CallerSession) RemoveLiquidityStatus() (bool, error) {
	return _Bnb4.Contract.RemoveLiquidityStatus(&_Bnb4.CallOpts)
}

// RewardBuyAmount is a free data retrieval call binding the contract method 0x10da207d.
//
// Solidity: function rewardBuyAmount() view returns(uint256)
func (_Bnb4 *Bnb4Caller) RewardBuyAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bnb4.contract.Call(opts, &out, "rewardBuyAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardBuyAmount is a free data retrieval call binding the contract method 0x10da207d.
//
// Solidity: function rewardBuyAmount() view returns(uint256)
func (_Bnb4 *Bnb4Session) RewardBuyAmount() (*big.Int, error) {
	return _Bnb4.Contract.RewardBuyAmount(&_Bnb4.CallOpts)
}

// RewardBuyAmount is a free data retrieval call binding the contract method 0x10da207d.
//
// Solidity: function rewardBuyAmount() view returns(uint256)
func (_Bnb4 *Bnb4CallerSession) RewardBuyAmount() (*big.Int, error) {
	return _Bnb4.Contract.RewardBuyAmount(&_Bnb4.CallOpts)
}

// RewardSellAmount is a free data retrieval call binding the contract method 0x0cc6cac8.
//
// Solidity: function rewardSellAmount() view returns(uint256)
func (_Bnb4 *Bnb4Caller) RewardSellAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bnb4.contract.Call(opts, &out, "rewardSellAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardSellAmount is a free data retrieval call binding the contract method 0x0cc6cac8.
//
// Solidity: function rewardSellAmount() view returns(uint256)
func (_Bnb4 *Bnb4Session) RewardSellAmount() (*big.Int, error) {
	return _Bnb4.Contract.RewardSellAmount(&_Bnb4.CallOpts)
}

// RewardSellAmount is a free data retrieval call binding the contract method 0x0cc6cac8.
//
// Solidity: function rewardSellAmount() view returns(uint256)
func (_Bnb4 *Bnb4CallerSession) RewardSellAmount() (*big.Int, error) {
	return _Bnb4.Contract.RewardSellAmount(&_Bnb4.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_Bnb4 *Bnb4Caller) StartTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bnb4.contract.Call(opts, &out, "startTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_Bnb4 *Bnb4Session) StartTime() (*big.Int, error) {
	return _Bnb4.Contract.StartTime(&_Bnb4.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_Bnb4 *Bnb4CallerSession) StartTime() (*big.Int, error) {
	return _Bnb4.Contract.StartTime(&_Bnb4.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Bnb4 *Bnb4Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Bnb4.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Bnb4 *Bnb4Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Bnb4.Contract.SupportsInterface(&_Bnb4.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Bnb4 *Bnb4CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Bnb4.Contract.SupportsInterface(&_Bnb4.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Bnb4 *Bnb4Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Bnb4.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Bnb4 *Bnb4Session) Symbol() (string, error) {
	return _Bnb4.Contract.Symbol(&_Bnb4.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Bnb4 *Bnb4CallerSession) Symbol() (string, error) {
	return _Bnb4.Contract.Symbol(&_Bnb4.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Bnb4 *Bnb4Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bnb4.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Bnb4 *Bnb4Session) TotalSupply() (*big.Int, error) {
	return _Bnb4.Contract.TotalSupply(&_Bnb4.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Bnb4 *Bnb4CallerSession) TotalSupply() (*big.Int, error) {
	return _Bnb4.Contract.TotalSupply(&_Bnb4.CallOpts)
}

// UniswapV2Pair is a free data retrieval call binding the contract method 0x49bd5a5e.
//
// Solidity: function uniswapV2Pair() view returns(address)
func (_Bnb4 *Bnb4Caller) UniswapV2Pair(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bnb4.contract.Call(opts, &out, "uniswapV2Pair")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UniswapV2Pair is a free data retrieval call binding the contract method 0x49bd5a5e.
//
// Solidity: function uniswapV2Pair() view returns(address)
func (_Bnb4 *Bnb4Session) UniswapV2Pair() (common.Address, error) {
	return _Bnb4.Contract.UniswapV2Pair(&_Bnb4.CallOpts)
}

// UniswapV2Pair is a free data retrieval call binding the contract method 0x49bd5a5e.
//
// Solidity: function uniswapV2Pair() view returns(address)
func (_Bnb4 *Bnb4CallerSession) UniswapV2Pair() (common.Address, error) {
	return _Bnb4.Contract.UniswapV2Pair(&_Bnb4.CallOpts)
}

// UniswapV2Router is a free data retrieval call binding the contract method 0x1694505e.
//
// Solidity: function uniswapV2Router() view returns(address)
func (_Bnb4 *Bnb4Caller) UniswapV2Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bnb4.contract.Call(opts, &out, "uniswapV2Router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UniswapV2Router is a free data retrieval call binding the contract method 0x1694505e.
//
// Solidity: function uniswapV2Router() view returns(address)
func (_Bnb4 *Bnb4Session) UniswapV2Router() (common.Address, error) {
	return _Bnb4.Contract.UniswapV2Router(&_Bnb4.CallOpts)
}

// UniswapV2Router is a free data retrieval call binding the contract method 0x1694505e.
//
// Solidity: function uniswapV2Router() view returns(address)
func (_Bnb4 *Bnb4CallerSession) UniswapV2Router() (common.Address, error) {
	return _Bnb4.Contract.UniswapV2Router(&_Bnb4.CallOpts)
}

// Launch is a paid mutator transaction binding the contract method 0xa72e5e15.
//
// Solidity: function Launch(uint256 _time) returns()
func (_Bnb4 *Bnb4Transactor) Launch(opts *bind.TransactOpts, _time *big.Int) (*types.Transaction, error) {
	return _Bnb4.contract.Transact(opts, "Launch", _time)
}

// Launch is a paid mutator transaction binding the contract method 0xa72e5e15.
//
// Solidity: function Launch(uint256 _time) returns()
func (_Bnb4 *Bnb4Session) Launch(_time *big.Int) (*types.Transaction, error) {
	return _Bnb4.Contract.Launch(&_Bnb4.TransactOpts, _time)
}

// Launch is a paid mutator transaction binding the contract method 0xa72e5e15.
//
// Solidity: function Launch(uint256 _time) returns()
func (_Bnb4 *Bnb4TransactorSession) Launch(_time *big.Int) (*types.Transaction, error) {
	return _Bnb4.Contract.Launch(&_Bnb4.TransactOpts, _time)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Bnb4 *Bnb4Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bnb4.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Bnb4 *Bnb4Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bnb4.Contract.Approve(&_Bnb4.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Bnb4 *Bnb4TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bnb4.Contract.Approve(&_Bnb4.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Bnb4 *Bnb4Transactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Bnb4.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Bnb4 *Bnb4Session) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Bnb4.Contract.DecreaseAllowance(&_Bnb4.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Bnb4 *Bnb4TransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Bnb4.Contract.DecreaseAllowance(&_Bnb4.TransactOpts, spender, subtractedValue)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Bnb4 *Bnb4Transactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bnb4.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Bnb4 *Bnb4Session) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bnb4.Contract.GrantRole(&_Bnb4.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Bnb4 *Bnb4TransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bnb4.Contract.GrantRole(&_Bnb4.TransactOpts, role, account)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Bnb4 *Bnb4Transactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Bnb4.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Bnb4 *Bnb4Session) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Bnb4.Contract.IncreaseAllowance(&_Bnb4.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Bnb4 *Bnb4TransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Bnb4.Contract.IncreaseAllowance(&_Bnb4.TransactOpts, spender, addedValue)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Bnb4 *Bnb4Transactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bnb4.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Bnb4 *Bnb4Session) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bnb4.Contract.RenounceRole(&_Bnb4.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Bnb4 *Bnb4TransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bnb4.Contract.RenounceRole(&_Bnb4.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Bnb4 *Bnb4Transactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bnb4.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Bnb4 *Bnb4Session) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bnb4.Contract.RevokeRole(&_Bnb4.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Bnb4 *Bnb4TransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bnb4.Contract.RevokeRole(&_Bnb4.TransactOpts, role, account)
}

// Reward is a paid mutator transaction binding the contract method 0x228cb733.
//
// Solidity: function reward() payable returns()
func (_Bnb4 *Bnb4Transactor) Reward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bnb4.contract.Transact(opts, "reward")
}

// Reward is a paid mutator transaction binding the contract method 0x228cb733.
//
// Solidity: function reward() payable returns()
func (_Bnb4 *Bnb4Session) Reward() (*types.Transaction, error) {
	return _Bnb4.Contract.Reward(&_Bnb4.TransactOpts)
}

// Reward is a paid mutator transaction binding the contract method 0x228cb733.
//
// Solidity: function reward() payable returns()
func (_Bnb4 *Bnb4TransactorSession) Reward() (*types.Transaction, error) {
	return _Bnb4.Contract.Reward(&_Bnb4.TransactOpts)
}

// SetBlack is a paid mutator transaction binding the contract method 0xcb23bf08.
//
// Solidity: function setBlack(address account, bool res) returns()
func (_Bnb4 *Bnb4Transactor) SetBlack(opts *bind.TransactOpts, account common.Address, res bool) (*types.Transaction, error) {
	return _Bnb4.contract.Transact(opts, "setBlack", account, res)
}

// SetBlack is a paid mutator transaction binding the contract method 0xcb23bf08.
//
// Solidity: function setBlack(address account, bool res) returns()
func (_Bnb4 *Bnb4Session) SetBlack(account common.Address, res bool) (*types.Transaction, error) {
	return _Bnb4.Contract.SetBlack(&_Bnb4.TransactOpts, account, res)
}

// SetBlack is a paid mutator transaction binding the contract method 0xcb23bf08.
//
// Solidity: function setBlack(address account, bool res) returns()
func (_Bnb4 *Bnb4TransactorSession) SetBlack(account common.Address, res bool) (*types.Transaction, error) {
	return _Bnb4.Contract.SetBlack(&_Bnb4.TransactOpts, account, res)
}

// SetBnb is a paid mutator transaction binding the contract method 0x337e9527.
//
// Solidity: function setBnb(address _bnb) returns()
func (_Bnb4 *Bnb4Transactor) SetBnb(opts *bind.TransactOpts, _bnb common.Address) (*types.Transaction, error) {
	return _Bnb4.contract.Transact(opts, "setBnb", _bnb)
}

// SetBnb is a paid mutator transaction binding the contract method 0x337e9527.
//
// Solidity: function setBnb(address _bnb) returns()
func (_Bnb4 *Bnb4Session) SetBnb(_bnb common.Address) (*types.Transaction, error) {
	return _Bnb4.Contract.SetBnb(&_Bnb4.TransactOpts, _bnb)
}

// SetBnb is a paid mutator transaction binding the contract method 0x337e9527.
//
// Solidity: function setBnb(address _bnb) returns()
func (_Bnb4 *Bnb4TransactorSession) SetBnb(_bnb common.Address) (*types.Transaction, error) {
	return _Bnb4.Contract.SetBnb(&_Bnb4.TransactOpts, _bnb)
}

// SetBuyRewardRateLevel1 is a paid mutator transaction binding the contract method 0x65507591.
//
// Solidity: function setBuyRewardRateLevel1(uint256 res) returns()
func (_Bnb4 *Bnb4Transactor) SetBuyRewardRateLevel1(opts *bind.TransactOpts, res *big.Int) (*types.Transaction, error) {
	return _Bnb4.contract.Transact(opts, "setBuyRewardRateLevel1", res)
}

// SetBuyRewardRateLevel1 is a paid mutator transaction binding the contract method 0x65507591.
//
// Solidity: function setBuyRewardRateLevel1(uint256 res) returns()
func (_Bnb4 *Bnb4Session) SetBuyRewardRateLevel1(res *big.Int) (*types.Transaction, error) {
	return _Bnb4.Contract.SetBuyRewardRateLevel1(&_Bnb4.TransactOpts, res)
}

// SetBuyRewardRateLevel1 is a paid mutator transaction binding the contract method 0x65507591.
//
// Solidity: function setBuyRewardRateLevel1(uint256 res) returns()
func (_Bnb4 *Bnb4TransactorSession) SetBuyRewardRateLevel1(res *big.Int) (*types.Transaction, error) {
	return _Bnb4.Contract.SetBuyRewardRateLevel1(&_Bnb4.TransactOpts, res)
}

// SetBuyRewardRateLevel2 is a paid mutator transaction binding the contract method 0xaa214b0f.
//
// Solidity: function setBuyRewardRateLevel2(uint256 res) returns()
func (_Bnb4 *Bnb4Transactor) SetBuyRewardRateLevel2(opts *bind.TransactOpts, res *big.Int) (*types.Transaction, error) {
	return _Bnb4.contract.Transact(opts, "setBuyRewardRateLevel2", res)
}

// SetBuyRewardRateLevel2 is a paid mutator transaction binding the contract method 0xaa214b0f.
//
// Solidity: function setBuyRewardRateLevel2(uint256 res) returns()
func (_Bnb4 *Bnb4Session) SetBuyRewardRateLevel2(res *big.Int) (*types.Transaction, error) {
	return _Bnb4.Contract.SetBuyRewardRateLevel2(&_Bnb4.TransactOpts, res)
}

// SetBuyRewardRateLevel2 is a paid mutator transaction binding the contract method 0xaa214b0f.
//
// Solidity: function setBuyRewardRateLevel2(uint256 res) returns()
func (_Bnb4 *Bnb4TransactorSession) SetBuyRewardRateLevel2(res *big.Int) (*types.Transaction, error) {
	return _Bnb4.Contract.SetBuyRewardRateLevel2(&_Bnb4.TransactOpts, res)
}

// SetBuyRewardRateLevel3 is a paid mutator transaction binding the contract method 0xf8dd41ec.
//
// Solidity: function setBuyRewardRateLevel3(uint256 res) returns()
func (_Bnb4 *Bnb4Transactor) SetBuyRewardRateLevel3(opts *bind.TransactOpts, res *big.Int) (*types.Transaction, error) {
	return _Bnb4.contract.Transact(opts, "setBuyRewardRateLevel3", res)
}

// SetBuyRewardRateLevel3 is a paid mutator transaction binding the contract method 0xf8dd41ec.
//
// Solidity: function setBuyRewardRateLevel3(uint256 res) returns()
func (_Bnb4 *Bnb4Session) SetBuyRewardRateLevel3(res *big.Int) (*types.Transaction, error) {
	return _Bnb4.Contract.SetBuyRewardRateLevel3(&_Bnb4.TransactOpts, res)
}

// SetBuyRewardRateLevel3 is a paid mutator transaction binding the contract method 0xf8dd41ec.
//
// Solidity: function setBuyRewardRateLevel3(uint256 res) returns()
func (_Bnb4 *Bnb4TransactorSession) SetBuyRewardRateLevel3(res *big.Int) (*types.Transaction, error) {
	return _Bnb4.Contract.SetBuyRewardRateLevel3(&_Bnb4.TransactOpts, res)
}

// SetBuyRewardRateLevel4 is a paid mutator transaction binding the contract method 0xc1fdd98c.
//
// Solidity: function setBuyRewardRateLevel4(uint256 res) returns()
func (_Bnb4 *Bnb4Transactor) SetBuyRewardRateLevel4(opts *bind.TransactOpts, res *big.Int) (*types.Transaction, error) {
	return _Bnb4.contract.Transact(opts, "setBuyRewardRateLevel4", res)
}

// SetBuyRewardRateLevel4 is a paid mutator transaction binding the contract method 0xc1fdd98c.
//
// Solidity: function setBuyRewardRateLevel4(uint256 res) returns()
func (_Bnb4 *Bnb4Session) SetBuyRewardRateLevel4(res *big.Int) (*types.Transaction, error) {
	return _Bnb4.Contract.SetBuyRewardRateLevel4(&_Bnb4.TransactOpts, res)
}

// SetBuyRewardRateLevel4 is a paid mutator transaction binding the contract method 0xc1fdd98c.
//
// Solidity: function setBuyRewardRateLevel4(uint256 res) returns()
func (_Bnb4 *Bnb4TransactorSession) SetBuyRewardRateLevel4(res *big.Int) (*types.Transaction, error) {
	return _Bnb4.Contract.SetBuyRewardRateLevel4(&_Bnb4.TransactOpts, res)
}

// SetBuyRewardRateLevel5 is a paid mutator transaction binding the contract method 0x38909f0e.
//
// Solidity: function setBuyRewardRateLevel5(uint256 res) returns()
func (_Bnb4 *Bnb4Transactor) SetBuyRewardRateLevel5(opts *bind.TransactOpts, res *big.Int) (*types.Transaction, error) {
	return _Bnb4.contract.Transact(opts, "setBuyRewardRateLevel5", res)
}

// SetBuyRewardRateLevel5 is a paid mutator transaction binding the contract method 0x38909f0e.
//
// Solidity: function setBuyRewardRateLevel5(uint256 res) returns()
func (_Bnb4 *Bnb4Session) SetBuyRewardRateLevel5(res *big.Int) (*types.Transaction, error) {
	return _Bnb4.Contract.SetBuyRewardRateLevel5(&_Bnb4.TransactOpts, res)
}

// SetBuyRewardRateLevel5 is a paid mutator transaction binding the contract method 0x38909f0e.
//
// Solidity: function setBuyRewardRateLevel5(uint256 res) returns()
func (_Bnb4 *Bnb4TransactorSession) SetBuyRewardRateLevel5(res *big.Int) (*types.Transaction, error) {
	return _Bnb4.Contract.SetBuyRewardRateLevel5(&_Bnb4.TransactOpts, res)
}

// SetExistsUserByAdmin is a paid mutator transaction binding the contract method 0x41d8e35d.
//
// Solidity: function setExistsUserByAdmin(address account, address recommendAccount) returns()
func (_Bnb4 *Bnb4Transactor) SetExistsUserByAdmin(opts *bind.TransactOpts, account common.Address, recommendAccount common.Address) (*types.Transaction, error) {
	return _Bnb4.contract.Transact(opts, "setExistsUserByAdmin", account, recommendAccount)
}

// SetExistsUserByAdmin is a paid mutator transaction binding the contract method 0x41d8e35d.
//
// Solidity: function setExistsUserByAdmin(address account, address recommendAccount) returns()
func (_Bnb4 *Bnb4Session) SetExistsUserByAdmin(account common.Address, recommendAccount common.Address) (*types.Transaction, error) {
	return _Bnb4.Contract.SetExistsUserByAdmin(&_Bnb4.TransactOpts, account, recommendAccount)
}

// SetExistsUserByAdmin is a paid mutator transaction binding the contract method 0x41d8e35d.
//
// Solidity: function setExistsUserByAdmin(address account, address recommendAccount) returns()
func (_Bnb4 *Bnb4TransactorSession) SetExistsUserByAdmin(account common.Address, recommendAccount common.Address) (*types.Transaction, error) {
	return _Bnb4.Contract.SetExistsUserByAdmin(&_Bnb4.TransactOpts, account, recommendAccount)
}

// SetLaunch is a paid mutator transaction binding the contract method 0xc222740d.
//
// Solidity: function setLaunch(bool launch_) returns()
func (_Bnb4 *Bnb4Transactor) SetLaunch(opts *bind.TransactOpts, launch_ bool) (*types.Transaction, error) {
	return _Bnb4.contract.Transact(opts, "setLaunch", launch_)
}

// SetLaunch is a paid mutator transaction binding the contract method 0xc222740d.
//
// Solidity: function setLaunch(bool launch_) returns()
func (_Bnb4 *Bnb4Session) SetLaunch(launch_ bool) (*types.Transaction, error) {
	return _Bnb4.Contract.SetLaunch(&_Bnb4.TransactOpts, launch_)
}

// SetLaunch is a paid mutator transaction binding the contract method 0xc222740d.
//
// Solidity: function setLaunch(bool launch_) returns()
func (_Bnb4 *Bnb4TransactorSession) SetLaunch(launch_ bool) (*types.Transaction, error) {
	return _Bnb4.Contract.SetLaunch(&_Bnb4.TransactOpts, launch_)
}

// SetRemoveLiquidityStatus is a paid mutator transaction binding the contract method 0xc9fa56a5.
//
// Solidity: function setRemoveLiquidityStatus(bool status) returns()
func (_Bnb4 *Bnb4Transactor) SetRemoveLiquidityStatus(opts *bind.TransactOpts, status bool) (*types.Transaction, error) {
	return _Bnb4.contract.Transact(opts, "setRemoveLiquidityStatus", status)
}

// SetRemoveLiquidityStatus is a paid mutator transaction binding the contract method 0xc9fa56a5.
//
// Solidity: function setRemoveLiquidityStatus(bool status) returns()
func (_Bnb4 *Bnb4Session) SetRemoveLiquidityStatus(status bool) (*types.Transaction, error) {
	return _Bnb4.Contract.SetRemoveLiquidityStatus(&_Bnb4.TransactOpts, status)
}

// SetRemoveLiquidityStatus is a paid mutator transaction binding the contract method 0xc9fa56a5.
//
// Solidity: function setRemoveLiquidityStatus(bool status) returns()
func (_Bnb4 *Bnb4TransactorSession) SetRemoveLiquidityStatus(status bool) (*types.Transaction, error) {
	return _Bnb4.Contract.SetRemoveLiquidityStatus(&_Bnb4.TransactOpts, status)
}

// SetSellMarketRate is a paid mutator transaction binding the contract method 0x681cb554.
//
// Solidity: function setSellMarketRate(uint256 res) returns()
func (_Bnb4 *Bnb4Transactor) SetSellMarketRate(opts *bind.TransactOpts, res *big.Int) (*types.Transaction, error) {
	return _Bnb4.contract.Transact(opts, "setSellMarketRate", res)
}

// SetSellMarketRate is a paid mutator transaction binding the contract method 0x681cb554.
//
// Solidity: function setSellMarketRate(uint256 res) returns()
func (_Bnb4 *Bnb4Session) SetSellMarketRate(res *big.Int) (*types.Transaction, error) {
	return _Bnb4.Contract.SetSellMarketRate(&_Bnb4.TransactOpts, res)
}

// SetSellMarketRate is a paid mutator transaction binding the contract method 0x681cb554.
//
// Solidity: function setSellMarketRate(uint256 res) returns()
func (_Bnb4 *Bnb4TransactorSession) SetSellMarketRate(res *big.Int) (*types.Transaction, error) {
	return _Bnb4.Contract.SetSellMarketRate(&_Bnb4.TransactOpts, res)
}

// SetSellPooltRate is a paid mutator transaction binding the contract method 0x72f82a40.
//
// Solidity: function setSellPooltRate(uint256 res) returns()
func (_Bnb4 *Bnb4Transactor) SetSellPooltRate(opts *bind.TransactOpts, res *big.Int) (*types.Transaction, error) {
	return _Bnb4.contract.Transact(opts, "setSellPooltRate", res)
}

// SetSellPooltRate is a paid mutator transaction binding the contract method 0x72f82a40.
//
// Solidity: function setSellPooltRate(uint256 res) returns()
func (_Bnb4 *Bnb4Session) SetSellPooltRate(res *big.Int) (*types.Transaction, error) {
	return _Bnb4.Contract.SetSellPooltRate(&_Bnb4.TransactOpts, res)
}

// SetSellPooltRate is a paid mutator transaction binding the contract method 0x72f82a40.
//
// Solidity: function setSellPooltRate(uint256 res) returns()
func (_Bnb4 *Bnb4TransactorSession) SetSellPooltRate(res *big.Int) (*types.Transaction, error) {
	return _Bnb4.Contract.SetSellPooltRate(&_Bnb4.TransactOpts, res)
}

// SetSellRewardRate is a paid mutator transaction binding the contract method 0xa7c25ba6.
//
// Solidity: function setSellRewardRate(uint256 res) returns()
func (_Bnb4 *Bnb4Transactor) SetSellRewardRate(opts *bind.TransactOpts, res *big.Int) (*types.Transaction, error) {
	return _Bnb4.contract.Transact(opts, "setSellRewardRate", res)
}

// SetSellRewardRate is a paid mutator transaction binding the contract method 0xa7c25ba6.
//
// Solidity: function setSellRewardRate(uint256 res) returns()
func (_Bnb4 *Bnb4Session) SetSellRewardRate(res *big.Int) (*types.Transaction, error) {
	return _Bnb4.Contract.SetSellRewardRate(&_Bnb4.TransactOpts, res)
}

// SetSellRewardRate is a paid mutator transaction binding the contract method 0xa7c25ba6.
//
// Solidity: function setSellRewardRate(uint256 res) returns()
func (_Bnb4 *Bnb4TransactorSession) SetSellRewardRate(res *big.Int) (*types.Transaction, error) {
	return _Bnb4.Contract.SetSellRewardRate(&_Bnb4.TransactOpts, res)
}

// SetSellRewardRateLevel1 is a paid mutator transaction binding the contract method 0xb7d36f02.
//
// Solidity: function setSellRewardRateLevel1(uint256 res) returns()
func (_Bnb4 *Bnb4Transactor) SetSellRewardRateLevel1(opts *bind.TransactOpts, res *big.Int) (*types.Transaction, error) {
	return _Bnb4.contract.Transact(opts, "setSellRewardRateLevel1", res)
}

// SetSellRewardRateLevel1 is a paid mutator transaction binding the contract method 0xb7d36f02.
//
// Solidity: function setSellRewardRateLevel1(uint256 res) returns()
func (_Bnb4 *Bnb4Session) SetSellRewardRateLevel1(res *big.Int) (*types.Transaction, error) {
	return _Bnb4.Contract.SetSellRewardRateLevel1(&_Bnb4.TransactOpts, res)
}

// SetSellRewardRateLevel1 is a paid mutator transaction binding the contract method 0xb7d36f02.
//
// Solidity: function setSellRewardRateLevel1(uint256 res) returns()
func (_Bnb4 *Bnb4TransactorSession) SetSellRewardRateLevel1(res *big.Int) (*types.Transaction, error) {
	return _Bnb4.Contract.SetSellRewardRateLevel1(&_Bnb4.TransactOpts, res)
}

// SetSellRewardRateLevel2 is a paid mutator transaction binding the contract method 0xc284fdbc.
//
// Solidity: function setSellRewardRateLevel2(uint256 res) returns()
func (_Bnb4 *Bnb4Transactor) SetSellRewardRateLevel2(opts *bind.TransactOpts, res *big.Int) (*types.Transaction, error) {
	return _Bnb4.contract.Transact(opts, "setSellRewardRateLevel2", res)
}

// SetSellRewardRateLevel2 is a paid mutator transaction binding the contract method 0xc284fdbc.
//
// Solidity: function setSellRewardRateLevel2(uint256 res) returns()
func (_Bnb4 *Bnb4Session) SetSellRewardRateLevel2(res *big.Int) (*types.Transaction, error) {
	return _Bnb4.Contract.SetSellRewardRateLevel2(&_Bnb4.TransactOpts, res)
}

// SetSellRewardRateLevel2 is a paid mutator transaction binding the contract method 0xc284fdbc.
//
// Solidity: function setSellRewardRateLevel2(uint256 res) returns()
func (_Bnb4 *Bnb4TransactorSession) SetSellRewardRateLevel2(res *big.Int) (*types.Transaction, error) {
	return _Bnb4.Contract.SetSellRewardRateLevel2(&_Bnb4.TransactOpts, res)
}

// SetSellRewardRateLevel3 is a paid mutator transaction binding the contract method 0x0411dc7c.
//
// Solidity: function setSellRewardRateLevel3(uint256 res) returns()
func (_Bnb4 *Bnb4Transactor) SetSellRewardRateLevel3(opts *bind.TransactOpts, res *big.Int) (*types.Transaction, error) {
	return _Bnb4.contract.Transact(opts, "setSellRewardRateLevel3", res)
}

// SetSellRewardRateLevel3 is a paid mutator transaction binding the contract method 0x0411dc7c.
//
// Solidity: function setSellRewardRateLevel3(uint256 res) returns()
func (_Bnb4 *Bnb4Session) SetSellRewardRateLevel3(res *big.Int) (*types.Transaction, error) {
	return _Bnb4.Contract.SetSellRewardRateLevel3(&_Bnb4.TransactOpts, res)
}

// SetSellRewardRateLevel3 is a paid mutator transaction binding the contract method 0x0411dc7c.
//
// Solidity: function setSellRewardRateLevel3(uint256 res) returns()
func (_Bnb4 *Bnb4TransactorSession) SetSellRewardRateLevel3(res *big.Int) (*types.Transaction, error) {
	return _Bnb4.Contract.SetSellRewardRateLevel3(&_Bnb4.TransactOpts, res)
}

// SetSellRewardRateLevel4 is a paid mutator transaction binding the contract method 0xd8042dad.
//
// Solidity: function setSellRewardRateLevel4(uint256 res) returns()
func (_Bnb4 *Bnb4Transactor) SetSellRewardRateLevel4(opts *bind.TransactOpts, res *big.Int) (*types.Transaction, error) {
	return _Bnb4.contract.Transact(opts, "setSellRewardRateLevel4", res)
}

// SetSellRewardRateLevel4 is a paid mutator transaction binding the contract method 0xd8042dad.
//
// Solidity: function setSellRewardRateLevel4(uint256 res) returns()
func (_Bnb4 *Bnb4Session) SetSellRewardRateLevel4(res *big.Int) (*types.Transaction, error) {
	return _Bnb4.Contract.SetSellRewardRateLevel4(&_Bnb4.TransactOpts, res)
}

// SetSellRewardRateLevel4 is a paid mutator transaction binding the contract method 0xd8042dad.
//
// Solidity: function setSellRewardRateLevel4(uint256 res) returns()
func (_Bnb4 *Bnb4TransactorSession) SetSellRewardRateLevel4(res *big.Int) (*types.Transaction, error) {
	return _Bnb4.Contract.SetSellRewardRateLevel4(&_Bnb4.TransactOpts, res)
}

// SetSellRewardRateLevel5 is a paid mutator transaction binding the contract method 0x9a513f45.
//
// Solidity: function setSellRewardRateLevel5(uint256 res) returns()
func (_Bnb4 *Bnb4Transactor) SetSellRewardRateLevel5(opts *bind.TransactOpts, res *big.Int) (*types.Transaction, error) {
	return _Bnb4.contract.Transact(opts, "setSellRewardRateLevel5", res)
}

// SetSellRewardRateLevel5 is a paid mutator transaction binding the contract method 0x9a513f45.
//
// Solidity: function setSellRewardRateLevel5(uint256 res) returns()
func (_Bnb4 *Bnb4Session) SetSellRewardRateLevel5(res *big.Int) (*types.Transaction, error) {
	return _Bnb4.Contract.SetSellRewardRateLevel5(&_Bnb4.TransactOpts, res)
}

// SetSellRewardRateLevel5 is a paid mutator transaction binding the contract method 0x9a513f45.
//
// Solidity: function setSellRewardRateLevel5(uint256 res) returns()
func (_Bnb4 *Bnb4TransactorSession) SetSellRewardRateLevel5(res *big.Int) (*types.Transaction, error) {
	return _Bnb4.Contract.SetSellRewardRateLevel5(&_Bnb4.TransactOpts, res)
}

// SetUser is a paid mutator transaction binding the contract method 0x858ced35.
//
// Solidity: function setUser(address account) returns()
func (_Bnb4 *Bnb4Transactor) SetUser(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Bnb4.contract.Transact(opts, "setUser", account)
}

// SetUser is a paid mutator transaction binding the contract method 0x858ced35.
//
// Solidity: function setUser(address account) returns()
func (_Bnb4 *Bnb4Session) SetUser(account common.Address) (*types.Transaction, error) {
	return _Bnb4.Contract.SetUser(&_Bnb4.TransactOpts, account)
}

// SetUser is a paid mutator transaction binding the contract method 0x858ced35.
//
// Solidity: function setUser(address account) returns()
func (_Bnb4 *Bnb4TransactorSession) SetUser(account common.Address) (*types.Transaction, error) {
	return _Bnb4.Contract.SetUser(&_Bnb4.TransactOpts, account)
}

// SetUsers is a paid mutator transaction binding the contract method 0x0a1b312f.
//
// Solidity: function setUsers(address[] account, address[] accountRecommend) returns()
func (_Bnb4 *Bnb4Transactor) SetUsers(opts *bind.TransactOpts, account []common.Address, accountRecommend []common.Address) (*types.Transaction, error) {
	return _Bnb4.contract.Transact(opts, "setUsers", account, accountRecommend)
}

// SetUsers is a paid mutator transaction binding the contract method 0x0a1b312f.
//
// Solidity: function setUsers(address[] account, address[] accountRecommend) returns()
func (_Bnb4 *Bnb4Session) SetUsers(account []common.Address, accountRecommend []common.Address) (*types.Transaction, error) {
	return _Bnb4.Contract.SetUsers(&_Bnb4.TransactOpts, account, accountRecommend)
}

// SetUsers is a paid mutator transaction binding the contract method 0x0a1b312f.
//
// Solidity: function setUsers(address[] account, address[] accountRecommend) returns()
func (_Bnb4 *Bnb4TransactorSession) SetUsers(account []common.Address, accountRecommend []common.Address) (*types.Transaction, error) {
	return _Bnb4.Contract.SetUsers(&_Bnb4.TransactOpts, account, accountRecommend)
}

// SetUsersFirst is a paid mutator transaction binding the contract method 0xebca5f05.
//
// Solidity: function setUsersFirst(address account) returns()
func (_Bnb4 *Bnb4Transactor) SetUsersFirst(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Bnb4.contract.Transact(opts, "setUsersFirst", account)
}

// SetUsersFirst is a paid mutator transaction binding the contract method 0xebca5f05.
//
// Solidity: function setUsersFirst(address account) returns()
func (_Bnb4 *Bnb4Session) SetUsersFirst(account common.Address) (*types.Transaction, error) {
	return _Bnb4.Contract.SetUsersFirst(&_Bnb4.TransactOpts, account)
}

// SetUsersFirst is a paid mutator transaction binding the contract method 0xebca5f05.
//
// Solidity: function setUsersFirst(address account) returns()
func (_Bnb4 *Bnb4TransactorSession) SetUsersFirst(account common.Address) (*types.Transaction, error) {
	return _Bnb4.Contract.SetUsersFirst(&_Bnb4.TransactOpts, account)
}

// SetWhite is a paid mutator transaction binding the contract method 0x47240874.
//
// Solidity: function setWhite(address account, bool res) returns()
func (_Bnb4 *Bnb4Transactor) SetWhite(opts *bind.TransactOpts, account common.Address, res bool) (*types.Transaction, error) {
	return _Bnb4.contract.Transact(opts, "setWhite", account, res)
}

// SetWhite is a paid mutator transaction binding the contract method 0x47240874.
//
// Solidity: function setWhite(address account, bool res) returns()
func (_Bnb4 *Bnb4Session) SetWhite(account common.Address, res bool) (*types.Transaction, error) {
	return _Bnb4.Contract.SetWhite(&_Bnb4.TransactOpts, account, res)
}

// SetWhite is a paid mutator transaction binding the contract method 0x47240874.
//
// Solidity: function setWhite(address account, bool res) returns()
func (_Bnb4 *Bnb4TransactorSession) SetWhite(account common.Address, res bool) (*types.Transaction, error) {
	return _Bnb4.Contract.SetWhite(&_Bnb4.TransactOpts, account, res)
}

// SetbuyMarketRate is a paid mutator transaction binding the contract method 0xc355c4ca.
//
// Solidity: function setbuyMarketRate(uint256 res) returns()
func (_Bnb4 *Bnb4Transactor) SetbuyMarketRate(opts *bind.TransactOpts, res *big.Int) (*types.Transaction, error) {
	return _Bnb4.contract.Transact(opts, "setbuyMarketRate", res)
}

// SetbuyMarketRate is a paid mutator transaction binding the contract method 0xc355c4ca.
//
// Solidity: function setbuyMarketRate(uint256 res) returns()
func (_Bnb4 *Bnb4Session) SetbuyMarketRate(res *big.Int) (*types.Transaction, error) {
	return _Bnb4.Contract.SetbuyMarketRate(&_Bnb4.TransactOpts, res)
}

// SetbuyMarketRate is a paid mutator transaction binding the contract method 0xc355c4ca.
//
// Solidity: function setbuyMarketRate(uint256 res) returns()
func (_Bnb4 *Bnb4TransactorSession) SetbuyMarketRate(res *big.Int) (*types.Transaction, error) {
	return _Bnb4.Contract.SetbuyMarketRate(&_Bnb4.TransactOpts, res)
}

// SetbuyPoolRate is a paid mutator transaction binding the contract method 0x442d9b92.
//
// Solidity: function setbuyPoolRate(uint256 res) returns()
func (_Bnb4 *Bnb4Transactor) SetbuyPoolRate(opts *bind.TransactOpts, res *big.Int) (*types.Transaction, error) {
	return _Bnb4.contract.Transact(opts, "setbuyPoolRate", res)
}

// SetbuyPoolRate is a paid mutator transaction binding the contract method 0x442d9b92.
//
// Solidity: function setbuyPoolRate(uint256 res) returns()
func (_Bnb4 *Bnb4Session) SetbuyPoolRate(res *big.Int) (*types.Transaction, error) {
	return _Bnb4.Contract.SetbuyPoolRate(&_Bnb4.TransactOpts, res)
}

// SetbuyPoolRate is a paid mutator transaction binding the contract method 0x442d9b92.
//
// Solidity: function setbuyPoolRate(uint256 res) returns()
func (_Bnb4 *Bnb4TransactorSession) SetbuyPoolRate(res *big.Int) (*types.Transaction, error) {
	return _Bnb4.Contract.SetbuyPoolRate(&_Bnb4.TransactOpts, res)
}

// SetbuyRewardRate is a paid mutator transaction binding the contract method 0x2c7d413f.
//
// Solidity: function setbuyRewardRate(uint256 res) returns()
func (_Bnb4 *Bnb4Transactor) SetbuyRewardRate(opts *bind.TransactOpts, res *big.Int) (*types.Transaction, error) {
	return _Bnb4.contract.Transact(opts, "setbuyRewardRate", res)
}

// SetbuyRewardRate is a paid mutator transaction binding the contract method 0x2c7d413f.
//
// Solidity: function setbuyRewardRate(uint256 res) returns()
func (_Bnb4 *Bnb4Session) SetbuyRewardRate(res *big.Int) (*types.Transaction, error) {
	return _Bnb4.Contract.SetbuyRewardRate(&_Bnb4.TransactOpts, res)
}

// SetbuyRewardRate is a paid mutator transaction binding the contract method 0x2c7d413f.
//
// Solidity: function setbuyRewardRate(uint256 res) returns()
func (_Bnb4 *Bnb4TransactorSession) SetbuyRewardRate(res *big.Int) (*types.Transaction, error) {
	return _Bnb4.Contract.SetbuyRewardRate(&_Bnb4.TransactOpts, res)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Bnb4 *Bnb4Transactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bnb4.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Bnb4 *Bnb4Session) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bnb4.Contract.Transfer(&_Bnb4.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Bnb4 *Bnb4TransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bnb4.Contract.Transfer(&_Bnb4.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Bnb4 *Bnb4Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bnb4.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Bnb4 *Bnb4Session) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bnb4.Contract.TransferFrom(&_Bnb4.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Bnb4 *Bnb4TransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bnb4.Contract.TransferFrom(&_Bnb4.TransactOpts, from, to, amount)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_Bnb4 *Bnb4Transactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Bnb4.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_Bnb4 *Bnb4Session) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Bnb4.Contract.Fallback(&_Bnb4.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_Bnb4 *Bnb4TransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Bnb4.Contract.Fallback(&_Bnb4.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Bnb4 *Bnb4Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bnb4.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Bnb4 *Bnb4Session) Receive() (*types.Transaction, error) {
	return _Bnb4.Contract.Receive(&_Bnb4.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Bnb4 *Bnb4TransactorSession) Receive() (*types.Transaction, error) {
	return _Bnb4.Contract.Receive(&_Bnb4.TransactOpts)
}

// Bnb4ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Bnb4 contract.
type Bnb4ApprovalIterator struct {
	Event *Bnb4Approval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Bnb4ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bnb4Approval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Bnb4Approval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Bnb4ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bnb4ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bnb4Approval represents a Approval event raised by the Bnb4 contract.
type Bnb4Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Bnb4 *Bnb4Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*Bnb4ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Bnb4.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &Bnb4ApprovalIterator{contract: _Bnb4.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Bnb4 *Bnb4Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *Bnb4Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Bnb4.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bnb4Approval)
				if err := _Bnb4.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Bnb4 *Bnb4Filterer) ParseApproval(log types.Log) (*Bnb4Approval, error) {
	event := new(Bnb4Approval)
	if err := _Bnb4.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bnb4RoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Bnb4 contract.
type Bnb4RoleAdminChangedIterator struct {
	Event *Bnb4RoleAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Bnb4RoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bnb4RoleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Bnb4RoleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Bnb4RoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bnb4RoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bnb4RoleAdminChanged represents a RoleAdminChanged event raised by the Bnb4 contract.
type Bnb4RoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Bnb4 *Bnb4Filterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*Bnb4RoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Bnb4.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &Bnb4RoleAdminChangedIterator{contract: _Bnb4.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Bnb4 *Bnb4Filterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *Bnb4RoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Bnb4.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bnb4RoleAdminChanged)
				if err := _Bnb4.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Bnb4 *Bnb4Filterer) ParseRoleAdminChanged(log types.Log) (*Bnb4RoleAdminChanged, error) {
	event := new(Bnb4RoleAdminChanged)
	if err := _Bnb4.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bnb4RoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Bnb4 contract.
type Bnb4RoleGrantedIterator struct {
	Event *Bnb4RoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Bnb4RoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bnb4RoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Bnb4RoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Bnb4RoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bnb4RoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bnb4RoleGranted represents a RoleGranted event raised by the Bnb4 contract.
type Bnb4RoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Bnb4 *Bnb4Filterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*Bnb4RoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Bnb4.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &Bnb4RoleGrantedIterator{contract: _Bnb4.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Bnb4 *Bnb4Filterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *Bnb4RoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Bnb4.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bnb4RoleGranted)
				if err := _Bnb4.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Bnb4 *Bnb4Filterer) ParseRoleGranted(log types.Log) (*Bnb4RoleGranted, error) {
	event := new(Bnb4RoleGranted)
	if err := _Bnb4.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bnb4RoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Bnb4 contract.
type Bnb4RoleRevokedIterator struct {
	Event *Bnb4RoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Bnb4RoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bnb4RoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Bnb4RoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Bnb4RoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bnb4RoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bnb4RoleRevoked represents a RoleRevoked event raised by the Bnb4 contract.
type Bnb4RoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Bnb4 *Bnb4Filterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*Bnb4RoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Bnb4.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &Bnb4RoleRevokedIterator{contract: _Bnb4.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Bnb4 *Bnb4Filterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *Bnb4RoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Bnb4.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bnb4RoleRevoked)
				if err := _Bnb4.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Bnb4 *Bnb4Filterer) ParseRoleRevoked(log types.Log) (*Bnb4RoleRevoked, error) {
	event := new(Bnb4RoleRevoked)
	if err := _Bnb4.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bnb4TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Bnb4 contract.
type Bnb4TransferIterator struct {
	Event *Bnb4Transfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Bnb4TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bnb4Transfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Bnb4Transfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Bnb4TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bnb4TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bnb4Transfer represents a Transfer event raised by the Bnb4 contract.
type Bnb4Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Bnb4 *Bnb4Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*Bnb4TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Bnb4.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Bnb4TransferIterator{contract: _Bnb4.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Bnb4 *Bnb4Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *Bnb4Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Bnb4.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bnb4Transfer)
				if err := _Bnb4.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Bnb4 *Bnb4Filterer) ParseTransfer(log types.Log) (*Bnb4Transfer, error) {
	event := new(Bnb4Transfer)
	if err := _Bnb4.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
