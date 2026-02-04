// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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
	_ = abi.ConvertType
)

// AuctionMetaData contains all meta data concerning the Auction contract.
var AuctionMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"}],\"name\":\"AuctionCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountUsd\",\"type\":\"uint256\"}],\"name\":\"AuctionEnded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"usdValue\",\"type\":\"uint256\"}],\"name\":\"BidPlaced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"auctionCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"auctions\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"highestBidder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"highestBidEth\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"highestBidERC20\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"highestBidToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"highestBidUsd\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ended\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"}],\"name\":\"bid\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"bidWithERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"createAuction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"}],\"name\":\"endAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"erc20PriceFeeds\",\"outputs\":[{\"internalType\":\"contractAggregatorV3Interface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"erc20ToUsd\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ethPriceFeed\",\"outputs\":[{\"internalType\":\"contractAggregatorV3Interface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ethAmount\",\"type\":\"uint256\"}],\"name\":\"ethToUsd\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ethPriceFeed\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minBidUsd\",\"type\":\"uint256\"}],\"name\":\"initializeV2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minBidUsd\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pendingERC20Returns\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pendingReturns\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feed\",\"type\":\"address\"}],\"name\":\"setERC20PriceFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minBidUsd\",\"type\":\"uint256\"}],\"name\":\"setMinBidUsd\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"withdrawERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AuctionABI is the input ABI used to generate the binding from.
// Deprecated: Use AuctionMetaData.ABI instead.
var AuctionABI = AuctionMetaData.ABI

// Auction is an auto generated Go binding around an Ethereum contract.
type Auction struct {
	AuctionCaller     // Read-only binding to the contract
	AuctionTransactor // Write-only binding to the contract
	AuctionFilterer   // Log filterer for contract events
}

// AuctionCaller is an auto generated read-only Go binding around an Ethereum contract.
type AuctionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuctionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AuctionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuctionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AuctionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuctionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AuctionSession struct {
	Contract     *Auction          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AuctionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AuctionCallerSession struct {
	Contract *AuctionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// AuctionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AuctionTransactorSession struct {
	Contract     *AuctionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AuctionRaw is an auto generated low-level Go binding around an Ethereum contract.
type AuctionRaw struct {
	Contract *Auction // Generic contract binding to access the raw methods on
}

// AuctionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AuctionCallerRaw struct {
	Contract *AuctionCaller // Generic read-only contract binding to access the raw methods on
}

// AuctionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AuctionTransactorRaw struct {
	Contract *AuctionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAuction creates a new instance of Auction, bound to a specific deployed contract.
func NewAuction(address common.Address, backend bind.ContractBackend) (*Auction, error) {
	contract, err := bindAuction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Auction{AuctionCaller: AuctionCaller{contract: contract}, AuctionTransactor: AuctionTransactor{contract: contract}, AuctionFilterer: AuctionFilterer{contract: contract}}, nil
}

// NewAuctionCaller creates a new read-only instance of Auction, bound to a specific deployed contract.
func NewAuctionCaller(address common.Address, caller bind.ContractCaller) (*AuctionCaller, error) {
	contract, err := bindAuction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AuctionCaller{contract: contract}, nil
}

// NewAuctionTransactor creates a new write-only instance of Auction, bound to a specific deployed contract.
func NewAuctionTransactor(address common.Address, transactor bind.ContractTransactor) (*AuctionTransactor, error) {
	contract, err := bindAuction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AuctionTransactor{contract: contract}, nil
}

// NewAuctionFilterer creates a new log filterer instance of Auction, bound to a specific deployed contract.
func NewAuctionFilterer(address common.Address, filterer bind.ContractFilterer) (*AuctionFilterer, error) {
	contract, err := bindAuction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AuctionFilterer{contract: contract}, nil
}

// bindAuction binds a generic wrapper to an already deployed contract.
func bindAuction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AuctionMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Auction *AuctionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Auction.Contract.AuctionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Auction *AuctionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Auction.Contract.AuctionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Auction *AuctionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Auction.Contract.AuctionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Auction *AuctionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Auction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Auction *AuctionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Auction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Auction *AuctionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Auction.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Auction *AuctionCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Auction.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Auction *AuctionSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Auction.Contract.UPGRADEINTERFACEVERSION(&_Auction.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Auction *AuctionCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Auction.Contract.UPGRADEINTERFACEVERSION(&_Auction.CallOpts)
}

// AuctionCount is a free data retrieval call binding the contract method 0x2ad71573.
//
// Solidity: function auctionCount() view returns(uint256)
func (_Auction *AuctionCaller) AuctionCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Auction.contract.Call(opts, &out, "auctionCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AuctionCount is a free data retrieval call binding the contract method 0x2ad71573.
//
// Solidity: function auctionCount() view returns(uint256)
func (_Auction *AuctionSession) AuctionCount() (*big.Int, error) {
	return _Auction.Contract.AuctionCount(&_Auction.CallOpts)
}

// AuctionCount is a free data retrieval call binding the contract method 0x2ad71573.
//
// Solidity: function auctionCount() view returns(uint256)
func (_Auction *AuctionCallerSession) AuctionCount() (*big.Int, error) {
	return _Auction.Contract.AuctionCount(&_Auction.CallOpts)
}

// Auctions is a free data retrieval call binding the contract method 0x571a26a0.
//
// Solidity: function auctions(uint256 ) view returns(address seller, address nft, uint256 tokenId, uint256 endTime, address highestBidder, uint256 highestBidEth, uint256 highestBidERC20, address highestBidToken, uint256 highestBidUsd, bool ended)
func (_Auction *AuctionCaller) Auctions(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Seller          common.Address
	Nft             common.Address
	TokenId         *big.Int
	EndTime         *big.Int
	HighestBidder   common.Address
	HighestBidEth   *big.Int
	HighestBidERC20 *big.Int
	HighestBidToken common.Address
	HighestBidUsd   *big.Int
	Ended           bool
}, error) {
	var out []interface{}
	err := _Auction.contract.Call(opts, &out, "auctions", arg0)

	outstruct := new(struct {
		Seller          common.Address
		Nft             common.Address
		TokenId         *big.Int
		EndTime         *big.Int
		HighestBidder   common.Address
		HighestBidEth   *big.Int
		HighestBidERC20 *big.Int
		HighestBidToken common.Address
		HighestBidUsd   *big.Int
		Ended           bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Seller = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Nft = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.TokenId = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.EndTime = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.HighestBidder = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.HighestBidEth = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.HighestBidERC20 = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.HighestBidToken = *abi.ConvertType(out[7], new(common.Address)).(*common.Address)
	outstruct.HighestBidUsd = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.Ended = *abi.ConvertType(out[9], new(bool)).(*bool)

	return *outstruct, err

}

// Auctions is a free data retrieval call binding the contract method 0x571a26a0.
//
// Solidity: function auctions(uint256 ) view returns(address seller, address nft, uint256 tokenId, uint256 endTime, address highestBidder, uint256 highestBidEth, uint256 highestBidERC20, address highestBidToken, uint256 highestBidUsd, bool ended)
func (_Auction *AuctionSession) Auctions(arg0 *big.Int) (struct {
	Seller          common.Address
	Nft             common.Address
	TokenId         *big.Int
	EndTime         *big.Int
	HighestBidder   common.Address
	HighestBidEth   *big.Int
	HighestBidERC20 *big.Int
	HighestBidToken common.Address
	HighestBidUsd   *big.Int
	Ended           bool
}, error) {
	return _Auction.Contract.Auctions(&_Auction.CallOpts, arg0)
}

// Auctions is a free data retrieval call binding the contract method 0x571a26a0.
//
// Solidity: function auctions(uint256 ) view returns(address seller, address nft, uint256 tokenId, uint256 endTime, address highestBidder, uint256 highestBidEth, uint256 highestBidERC20, address highestBidToken, uint256 highestBidUsd, bool ended)
func (_Auction *AuctionCallerSession) Auctions(arg0 *big.Int) (struct {
	Seller          common.Address
	Nft             common.Address
	TokenId         *big.Int
	EndTime         *big.Int
	HighestBidder   common.Address
	HighestBidEth   *big.Int
	HighestBidERC20 *big.Int
	HighestBidToken common.Address
	HighestBidUsd   *big.Int
	Ended           bool
}, error) {
	return _Auction.Contract.Auctions(&_Auction.CallOpts, arg0)
}

// Erc20PriceFeeds is a free data retrieval call binding the contract method 0x1b20095d.
//
// Solidity: function erc20PriceFeeds(address ) view returns(address)
func (_Auction *AuctionCaller) Erc20PriceFeeds(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Auction.contract.Call(opts, &out, "erc20PriceFeeds", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Erc20PriceFeeds is a free data retrieval call binding the contract method 0x1b20095d.
//
// Solidity: function erc20PriceFeeds(address ) view returns(address)
func (_Auction *AuctionSession) Erc20PriceFeeds(arg0 common.Address) (common.Address, error) {
	return _Auction.Contract.Erc20PriceFeeds(&_Auction.CallOpts, arg0)
}

// Erc20PriceFeeds is a free data retrieval call binding the contract method 0x1b20095d.
//
// Solidity: function erc20PriceFeeds(address ) view returns(address)
func (_Auction *AuctionCallerSession) Erc20PriceFeeds(arg0 common.Address) (common.Address, error) {
	return _Auction.Contract.Erc20PriceFeeds(&_Auction.CallOpts, arg0)
}

// Erc20ToUsd is a free data retrieval call binding the contract method 0x2644d269.
//
// Solidity: function erc20ToUsd(address token, uint256 amount) view returns(uint256)
func (_Auction *AuctionCaller) Erc20ToUsd(opts *bind.CallOpts, token common.Address, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Auction.contract.Call(opts, &out, "erc20ToUsd", token, amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Erc20ToUsd is a free data retrieval call binding the contract method 0x2644d269.
//
// Solidity: function erc20ToUsd(address token, uint256 amount) view returns(uint256)
func (_Auction *AuctionSession) Erc20ToUsd(token common.Address, amount *big.Int) (*big.Int, error) {
	return _Auction.Contract.Erc20ToUsd(&_Auction.CallOpts, token, amount)
}

// Erc20ToUsd is a free data retrieval call binding the contract method 0x2644d269.
//
// Solidity: function erc20ToUsd(address token, uint256 amount) view returns(uint256)
func (_Auction *AuctionCallerSession) Erc20ToUsd(token common.Address, amount *big.Int) (*big.Int, error) {
	return _Auction.Contract.Erc20ToUsd(&_Auction.CallOpts, token, amount)
}

// EthPriceFeed is a free data retrieval call binding the contract method 0xaf7665ce.
//
// Solidity: function ethPriceFeed() view returns(address)
func (_Auction *AuctionCaller) EthPriceFeed(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Auction.contract.Call(opts, &out, "ethPriceFeed")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EthPriceFeed is a free data retrieval call binding the contract method 0xaf7665ce.
//
// Solidity: function ethPriceFeed() view returns(address)
func (_Auction *AuctionSession) EthPriceFeed() (common.Address, error) {
	return _Auction.Contract.EthPriceFeed(&_Auction.CallOpts)
}

// EthPriceFeed is a free data retrieval call binding the contract method 0xaf7665ce.
//
// Solidity: function ethPriceFeed() view returns(address)
func (_Auction *AuctionCallerSession) EthPriceFeed() (common.Address, error) {
	return _Auction.Contract.EthPriceFeed(&_Auction.CallOpts)
}

// EthToUsd is a free data retrieval call binding the contract method 0x946d1480.
//
// Solidity: function ethToUsd(uint256 ethAmount) view returns(uint256)
func (_Auction *AuctionCaller) EthToUsd(opts *bind.CallOpts, ethAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Auction.contract.Call(opts, &out, "ethToUsd", ethAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EthToUsd is a free data retrieval call binding the contract method 0x946d1480.
//
// Solidity: function ethToUsd(uint256 ethAmount) view returns(uint256)
func (_Auction *AuctionSession) EthToUsd(ethAmount *big.Int) (*big.Int, error) {
	return _Auction.Contract.EthToUsd(&_Auction.CallOpts, ethAmount)
}

// EthToUsd is a free data retrieval call binding the contract method 0x946d1480.
//
// Solidity: function ethToUsd(uint256 ethAmount) view returns(uint256)
func (_Auction *AuctionCallerSession) EthToUsd(ethAmount *big.Int) (*big.Int, error) {
	return _Auction.Contract.EthToUsd(&_Auction.CallOpts, ethAmount)
}

// MinBidUsd is a free data retrieval call binding the contract method 0x6c3c4c3b.
//
// Solidity: function minBidUsd() view returns(uint256)
func (_Auction *AuctionCaller) MinBidUsd(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Auction.contract.Call(opts, &out, "minBidUsd")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinBidUsd is a free data retrieval call binding the contract method 0x6c3c4c3b.
//
// Solidity: function minBidUsd() view returns(uint256)
func (_Auction *AuctionSession) MinBidUsd() (*big.Int, error) {
	return _Auction.Contract.MinBidUsd(&_Auction.CallOpts)
}

// MinBidUsd is a free data retrieval call binding the contract method 0x6c3c4c3b.
//
// Solidity: function minBidUsd() view returns(uint256)
func (_Auction *AuctionCallerSession) MinBidUsd() (*big.Int, error) {
	return _Auction.Contract.MinBidUsd(&_Auction.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Auction *AuctionCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Auction.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Auction *AuctionSession) Owner() (common.Address, error) {
	return _Auction.Contract.Owner(&_Auction.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Auction *AuctionCallerSession) Owner() (common.Address, error) {
	return _Auction.Contract.Owner(&_Auction.CallOpts)
}

// PendingERC20Returns is a free data retrieval call binding the contract method 0x95fdc17a.
//
// Solidity: function pendingERC20Returns(address , address ) view returns(uint256)
func (_Auction *AuctionCaller) PendingERC20Returns(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Auction.contract.Call(opts, &out, "pendingERC20Returns", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingERC20Returns is a free data retrieval call binding the contract method 0x95fdc17a.
//
// Solidity: function pendingERC20Returns(address , address ) view returns(uint256)
func (_Auction *AuctionSession) PendingERC20Returns(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Auction.Contract.PendingERC20Returns(&_Auction.CallOpts, arg0, arg1)
}

// PendingERC20Returns is a free data retrieval call binding the contract method 0x95fdc17a.
//
// Solidity: function pendingERC20Returns(address , address ) view returns(uint256)
func (_Auction *AuctionCallerSession) PendingERC20Returns(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Auction.Contract.PendingERC20Returns(&_Auction.CallOpts, arg0, arg1)
}

// PendingReturns is a free data retrieval call binding the contract method 0x26b387bb.
//
// Solidity: function pendingReturns(address ) view returns(uint256)
func (_Auction *AuctionCaller) PendingReturns(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Auction.contract.Call(opts, &out, "pendingReturns", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingReturns is a free data retrieval call binding the contract method 0x26b387bb.
//
// Solidity: function pendingReturns(address ) view returns(uint256)
func (_Auction *AuctionSession) PendingReturns(arg0 common.Address) (*big.Int, error) {
	return _Auction.Contract.PendingReturns(&_Auction.CallOpts, arg0)
}

// PendingReturns is a free data retrieval call binding the contract method 0x26b387bb.
//
// Solidity: function pendingReturns(address ) view returns(uint256)
func (_Auction *AuctionCallerSession) PendingReturns(arg0 common.Address) (*big.Int, error) {
	return _Auction.Contract.PendingReturns(&_Auction.CallOpts, arg0)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Auction *AuctionCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Auction.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Auction *AuctionSession) ProxiableUUID() ([32]byte, error) {
	return _Auction.Contract.ProxiableUUID(&_Auction.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Auction *AuctionCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Auction.Contract.ProxiableUUID(&_Auction.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Auction *AuctionCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Auction.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Auction *AuctionSession) Version() (string, error) {
	return _Auction.Contract.Version(&_Auction.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Auction *AuctionCallerSession) Version() (string, error) {
	return _Auction.Contract.Version(&_Auction.CallOpts)
}

// Bid is a paid mutator transaction binding the contract method 0x454a2ab3.
//
// Solidity: function bid(uint256 auctionId) payable returns()
func (_Auction *AuctionTransactor) Bid(opts *bind.TransactOpts, auctionId *big.Int) (*types.Transaction, error) {
	return _Auction.contract.Transact(opts, "bid", auctionId)
}

// Bid is a paid mutator transaction binding the contract method 0x454a2ab3.
//
// Solidity: function bid(uint256 auctionId) payable returns()
func (_Auction *AuctionSession) Bid(auctionId *big.Int) (*types.Transaction, error) {
	return _Auction.Contract.Bid(&_Auction.TransactOpts, auctionId)
}

// Bid is a paid mutator transaction binding the contract method 0x454a2ab3.
//
// Solidity: function bid(uint256 auctionId) payable returns()
func (_Auction *AuctionTransactorSession) Bid(auctionId *big.Int) (*types.Transaction, error) {
	return _Auction.Contract.Bid(&_Auction.TransactOpts, auctionId)
}

// BidWithERC20 is a paid mutator transaction binding the contract method 0x29b0af0a.
//
// Solidity: function bidWithERC20(uint256 auctionId, address token, uint256 amount) returns()
func (_Auction *AuctionTransactor) BidWithERC20(opts *bind.TransactOpts, auctionId *big.Int, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Auction.contract.Transact(opts, "bidWithERC20", auctionId, token, amount)
}

// BidWithERC20 is a paid mutator transaction binding the contract method 0x29b0af0a.
//
// Solidity: function bidWithERC20(uint256 auctionId, address token, uint256 amount) returns()
func (_Auction *AuctionSession) BidWithERC20(auctionId *big.Int, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Auction.Contract.BidWithERC20(&_Auction.TransactOpts, auctionId, token, amount)
}

// BidWithERC20 is a paid mutator transaction binding the contract method 0x29b0af0a.
//
// Solidity: function bidWithERC20(uint256 auctionId, address token, uint256 amount) returns()
func (_Auction *AuctionTransactorSession) BidWithERC20(auctionId *big.Int, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Auction.Contract.BidWithERC20(&_Auction.TransactOpts, auctionId, token, amount)
}

// CreateAuction is a paid mutator transaction binding the contract method 0xabb94051.
//
// Solidity: function createAuction(address nft, uint256 tokenId, uint256 duration) returns(uint256 auctionId)
func (_Auction *AuctionTransactor) CreateAuction(opts *bind.TransactOpts, nft common.Address, tokenId *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _Auction.contract.Transact(opts, "createAuction", nft, tokenId, duration)
}

// CreateAuction is a paid mutator transaction binding the contract method 0xabb94051.
//
// Solidity: function createAuction(address nft, uint256 tokenId, uint256 duration) returns(uint256 auctionId)
func (_Auction *AuctionSession) CreateAuction(nft common.Address, tokenId *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _Auction.Contract.CreateAuction(&_Auction.TransactOpts, nft, tokenId, duration)
}

// CreateAuction is a paid mutator transaction binding the contract method 0xabb94051.
//
// Solidity: function createAuction(address nft, uint256 tokenId, uint256 duration) returns(uint256 auctionId)
func (_Auction *AuctionTransactorSession) CreateAuction(nft common.Address, tokenId *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _Auction.Contract.CreateAuction(&_Auction.TransactOpts, nft, tokenId, duration)
}

// EndAuction is a paid mutator transaction binding the contract method 0xb9a2de3a.
//
// Solidity: function endAuction(uint256 auctionId) returns()
func (_Auction *AuctionTransactor) EndAuction(opts *bind.TransactOpts, auctionId *big.Int) (*types.Transaction, error) {
	return _Auction.contract.Transact(opts, "endAuction", auctionId)
}

// EndAuction is a paid mutator transaction binding the contract method 0xb9a2de3a.
//
// Solidity: function endAuction(uint256 auctionId) returns()
func (_Auction *AuctionSession) EndAuction(auctionId *big.Int) (*types.Transaction, error) {
	return _Auction.Contract.EndAuction(&_Auction.TransactOpts, auctionId)
}

// EndAuction is a paid mutator transaction binding the contract method 0xb9a2de3a.
//
// Solidity: function endAuction(uint256 auctionId) returns()
func (_Auction *AuctionTransactorSession) EndAuction(auctionId *big.Int) (*types.Transaction, error) {
	return _Auction.Contract.EndAuction(&_Auction.TransactOpts, auctionId)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _ethPriceFeed) returns()
func (_Auction *AuctionTransactor) Initialize(opts *bind.TransactOpts, _ethPriceFeed common.Address) (*types.Transaction, error) {
	return _Auction.contract.Transact(opts, "initialize", _ethPriceFeed)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _ethPriceFeed) returns()
func (_Auction *AuctionSession) Initialize(_ethPriceFeed common.Address) (*types.Transaction, error) {
	return _Auction.Contract.Initialize(&_Auction.TransactOpts, _ethPriceFeed)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _ethPriceFeed) returns()
func (_Auction *AuctionTransactorSession) Initialize(_ethPriceFeed common.Address) (*types.Transaction, error) {
	return _Auction.Contract.Initialize(&_Auction.TransactOpts, _ethPriceFeed)
}

// InitializeV2 is a paid mutator transaction binding the contract method 0x5b4e128c.
//
// Solidity: function initializeV2(uint256 _minBidUsd) returns()
func (_Auction *AuctionTransactor) InitializeV2(opts *bind.TransactOpts, _minBidUsd *big.Int) (*types.Transaction, error) {
	return _Auction.contract.Transact(opts, "initializeV2", _minBidUsd)
}

// InitializeV2 is a paid mutator transaction binding the contract method 0x5b4e128c.
//
// Solidity: function initializeV2(uint256 _minBidUsd) returns()
func (_Auction *AuctionSession) InitializeV2(_minBidUsd *big.Int) (*types.Transaction, error) {
	return _Auction.Contract.InitializeV2(&_Auction.TransactOpts, _minBidUsd)
}

// InitializeV2 is a paid mutator transaction binding the contract method 0x5b4e128c.
//
// Solidity: function initializeV2(uint256 _minBidUsd) returns()
func (_Auction *AuctionTransactorSession) InitializeV2(_minBidUsd *big.Int) (*types.Transaction, error) {
	return _Auction.Contract.InitializeV2(&_Auction.TransactOpts, _minBidUsd)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Auction *AuctionTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Auction.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Auction *AuctionSession) RenounceOwnership() (*types.Transaction, error) {
	return _Auction.Contract.RenounceOwnership(&_Auction.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Auction *AuctionTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Auction.Contract.RenounceOwnership(&_Auction.TransactOpts)
}

// SetERC20PriceFeed is a paid mutator transaction binding the contract method 0x12cb4a9c.
//
// Solidity: function setERC20PriceFeed(address token, address feed) returns()
func (_Auction *AuctionTransactor) SetERC20PriceFeed(opts *bind.TransactOpts, token common.Address, feed common.Address) (*types.Transaction, error) {
	return _Auction.contract.Transact(opts, "setERC20PriceFeed", token, feed)
}

// SetERC20PriceFeed is a paid mutator transaction binding the contract method 0x12cb4a9c.
//
// Solidity: function setERC20PriceFeed(address token, address feed) returns()
func (_Auction *AuctionSession) SetERC20PriceFeed(token common.Address, feed common.Address) (*types.Transaction, error) {
	return _Auction.Contract.SetERC20PriceFeed(&_Auction.TransactOpts, token, feed)
}

// SetERC20PriceFeed is a paid mutator transaction binding the contract method 0x12cb4a9c.
//
// Solidity: function setERC20PriceFeed(address token, address feed) returns()
func (_Auction *AuctionTransactorSession) SetERC20PriceFeed(token common.Address, feed common.Address) (*types.Transaction, error) {
	return _Auction.Contract.SetERC20PriceFeed(&_Auction.TransactOpts, token, feed)
}

// SetMinBidUsd is a paid mutator transaction binding the contract method 0x71ef768d.
//
// Solidity: function setMinBidUsd(uint256 _minBidUsd) returns()
func (_Auction *AuctionTransactor) SetMinBidUsd(opts *bind.TransactOpts, _minBidUsd *big.Int) (*types.Transaction, error) {
	return _Auction.contract.Transact(opts, "setMinBidUsd", _minBidUsd)
}

// SetMinBidUsd is a paid mutator transaction binding the contract method 0x71ef768d.
//
// Solidity: function setMinBidUsd(uint256 _minBidUsd) returns()
func (_Auction *AuctionSession) SetMinBidUsd(_minBidUsd *big.Int) (*types.Transaction, error) {
	return _Auction.Contract.SetMinBidUsd(&_Auction.TransactOpts, _minBidUsd)
}

// SetMinBidUsd is a paid mutator transaction binding the contract method 0x71ef768d.
//
// Solidity: function setMinBidUsd(uint256 _minBidUsd) returns()
func (_Auction *AuctionTransactorSession) SetMinBidUsd(_minBidUsd *big.Int) (*types.Transaction, error) {
	return _Auction.Contract.SetMinBidUsd(&_Auction.TransactOpts, _minBidUsd)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Auction *AuctionTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Auction.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Auction *AuctionSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Auction.Contract.TransferOwnership(&_Auction.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Auction *AuctionTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Auction.Contract.TransferOwnership(&_Auction.TransactOpts, newOwner)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Auction *AuctionTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Auction.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Auction *AuctionSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Auction.Contract.UpgradeToAndCall(&_Auction.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Auction *AuctionTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Auction.Contract.UpgradeToAndCall(&_Auction.TransactOpts, newImplementation, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Auction *AuctionTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Auction.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Auction *AuctionSession) Withdraw() (*types.Transaction, error) {
	return _Auction.Contract.Withdraw(&_Auction.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Auction *AuctionTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Auction.Contract.Withdraw(&_Auction.TransactOpts)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0xf4f3b200.
//
// Solidity: function withdrawERC20(address token) returns()
func (_Auction *AuctionTransactor) WithdrawERC20(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Auction.contract.Transact(opts, "withdrawERC20", token)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0xf4f3b200.
//
// Solidity: function withdrawERC20(address token) returns()
func (_Auction *AuctionSession) WithdrawERC20(token common.Address) (*types.Transaction, error) {
	return _Auction.Contract.WithdrawERC20(&_Auction.TransactOpts, token)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0xf4f3b200.
//
// Solidity: function withdrawERC20(address token) returns()
func (_Auction *AuctionTransactorSession) WithdrawERC20(token common.Address) (*types.Transaction, error) {
	return _Auction.Contract.WithdrawERC20(&_Auction.TransactOpts, token)
}

// AuctionAuctionCreatedIterator is returned from FilterAuctionCreated and is used to iterate over the raw logs and unpacked data for AuctionCreated events raised by the Auction contract.
type AuctionAuctionCreatedIterator struct {
	Event *AuctionAuctionCreated // Event containing the contract specifics and raw log

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
func (it *AuctionAuctionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionAuctionCreated)
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
		it.Event = new(AuctionAuctionCreated)
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
func (it *AuctionAuctionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionAuctionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionAuctionCreated represents a AuctionCreated event raised by the Auction contract.
type AuctionAuctionCreated struct {
	AuctionId  *big.Int
	Seller     common.Address
	NftAddress common.Address
	TokenId    *big.Int
	EndTime    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAuctionCreated is a free log retrieval operation binding the contract event 0x1e9fc4128626f087a2c46b8f819387ca9e130c81ced063176def67bc205b3b00.
//
// Solidity: event AuctionCreated(uint256 indexed auctionId, address indexed seller, address indexed nftAddress, uint256 tokenId, uint256 endTime)
func (_Auction *AuctionFilterer) FilterAuctionCreated(opts *bind.FilterOpts, auctionId []*big.Int, seller []common.Address, nftAddress []common.Address) (*AuctionAuctionCreatedIterator, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}

	logs, sub, err := _Auction.contract.FilterLogs(opts, "AuctionCreated", auctionIdRule, sellerRule, nftAddressRule)
	if err != nil {
		return nil, err
	}
	return &AuctionAuctionCreatedIterator{contract: _Auction.contract, event: "AuctionCreated", logs: logs, sub: sub}, nil
}

// WatchAuctionCreated is a free log subscription operation binding the contract event 0x1e9fc4128626f087a2c46b8f819387ca9e130c81ced063176def67bc205b3b00.
//
// Solidity: event AuctionCreated(uint256 indexed auctionId, address indexed seller, address indexed nftAddress, uint256 tokenId, uint256 endTime)
func (_Auction *AuctionFilterer) WatchAuctionCreated(opts *bind.WatchOpts, sink chan<- *AuctionAuctionCreated, auctionId []*big.Int, seller []common.Address, nftAddress []common.Address) (event.Subscription, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}

	logs, sub, err := _Auction.contract.WatchLogs(opts, "AuctionCreated", auctionIdRule, sellerRule, nftAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionAuctionCreated)
				if err := _Auction.contract.UnpackLog(event, "AuctionCreated", log); err != nil {
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

// ParseAuctionCreated is a log parse operation binding the contract event 0x1e9fc4128626f087a2c46b8f819387ca9e130c81ced063176def67bc205b3b00.
//
// Solidity: event AuctionCreated(uint256 indexed auctionId, address indexed seller, address indexed nftAddress, uint256 tokenId, uint256 endTime)
func (_Auction *AuctionFilterer) ParseAuctionCreated(log types.Log) (*AuctionAuctionCreated, error) {
	event := new(AuctionAuctionCreated)
	if err := _Auction.contract.UnpackLog(event, "AuctionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuctionAuctionEndedIterator is returned from FilterAuctionEnded and is used to iterate over the raw logs and unpacked data for AuctionEnded events raised by the Auction contract.
type AuctionAuctionEndedIterator struct {
	Event *AuctionAuctionEnded // Event containing the contract specifics and raw log

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
func (it *AuctionAuctionEndedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionAuctionEnded)
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
		it.Event = new(AuctionAuctionEnded)
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
func (it *AuctionAuctionEndedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionAuctionEndedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionAuctionEnded represents a AuctionEnded event raised by the Auction contract.
type AuctionAuctionEnded struct {
	AuctionId *big.Int
	Winner    common.Address
	AmountUsd *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAuctionEnded is a free log retrieval operation binding the contract event 0xd2aa34a4fdbbc6dff6a3e56f46e0f3ae2a31d7785ff3487aa5c95c642acea501.
//
// Solidity: event AuctionEnded(uint256 indexed auctionId, address winner, uint256 amountUsd)
func (_Auction *AuctionFilterer) FilterAuctionEnded(opts *bind.FilterOpts, auctionId []*big.Int) (*AuctionAuctionEndedIterator, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}

	logs, sub, err := _Auction.contract.FilterLogs(opts, "AuctionEnded", auctionIdRule)
	if err != nil {
		return nil, err
	}
	return &AuctionAuctionEndedIterator{contract: _Auction.contract, event: "AuctionEnded", logs: logs, sub: sub}, nil
}

// WatchAuctionEnded is a free log subscription operation binding the contract event 0xd2aa34a4fdbbc6dff6a3e56f46e0f3ae2a31d7785ff3487aa5c95c642acea501.
//
// Solidity: event AuctionEnded(uint256 indexed auctionId, address winner, uint256 amountUsd)
func (_Auction *AuctionFilterer) WatchAuctionEnded(opts *bind.WatchOpts, sink chan<- *AuctionAuctionEnded, auctionId []*big.Int) (event.Subscription, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}

	logs, sub, err := _Auction.contract.WatchLogs(opts, "AuctionEnded", auctionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionAuctionEnded)
				if err := _Auction.contract.UnpackLog(event, "AuctionEnded", log); err != nil {
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

// ParseAuctionEnded is a log parse operation binding the contract event 0xd2aa34a4fdbbc6dff6a3e56f46e0f3ae2a31d7785ff3487aa5c95c642acea501.
//
// Solidity: event AuctionEnded(uint256 indexed auctionId, address winner, uint256 amountUsd)
func (_Auction *AuctionFilterer) ParseAuctionEnded(log types.Log) (*AuctionAuctionEnded, error) {
	event := new(AuctionAuctionEnded)
	if err := _Auction.contract.UnpackLog(event, "AuctionEnded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuctionBidPlacedIterator is returned from FilterBidPlaced and is used to iterate over the raw logs and unpacked data for BidPlaced events raised by the Auction contract.
type AuctionBidPlacedIterator struct {
	Event *AuctionBidPlaced // Event containing the contract specifics and raw log

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
func (it *AuctionBidPlacedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionBidPlaced)
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
		it.Event = new(AuctionBidPlaced)
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
func (it *AuctionBidPlacedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionBidPlacedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionBidPlaced represents a BidPlaced event raised by the Auction contract.
type AuctionBidPlaced struct {
	AuctionId *big.Int
	Bidder    common.Address
	UsdValue  *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBidPlaced is a free log retrieval operation binding the contract event 0x0e54eff26401bf69b81b26f60bd85ef47f5d85275c1d268d84f68d6897431c47.
//
// Solidity: event BidPlaced(uint256 indexed auctionId, address indexed bidder, uint256 usdValue)
func (_Auction *AuctionFilterer) FilterBidPlaced(opts *bind.FilterOpts, auctionId []*big.Int, bidder []common.Address) (*AuctionBidPlacedIterator, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}
	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _Auction.contract.FilterLogs(opts, "BidPlaced", auctionIdRule, bidderRule)
	if err != nil {
		return nil, err
	}
	return &AuctionBidPlacedIterator{contract: _Auction.contract, event: "BidPlaced", logs: logs, sub: sub}, nil
}

// WatchBidPlaced is a free log subscription operation binding the contract event 0x0e54eff26401bf69b81b26f60bd85ef47f5d85275c1d268d84f68d6897431c47.
//
// Solidity: event BidPlaced(uint256 indexed auctionId, address indexed bidder, uint256 usdValue)
func (_Auction *AuctionFilterer) WatchBidPlaced(opts *bind.WatchOpts, sink chan<- *AuctionBidPlaced, auctionId []*big.Int, bidder []common.Address) (event.Subscription, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}
	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _Auction.contract.WatchLogs(opts, "BidPlaced", auctionIdRule, bidderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionBidPlaced)
				if err := _Auction.contract.UnpackLog(event, "BidPlaced", log); err != nil {
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

// ParseBidPlaced is a log parse operation binding the contract event 0x0e54eff26401bf69b81b26f60bd85ef47f5d85275c1d268d84f68d6897431c47.
//
// Solidity: event BidPlaced(uint256 indexed auctionId, address indexed bidder, uint256 usdValue)
func (_Auction *AuctionFilterer) ParseBidPlaced(log types.Log) (*AuctionBidPlaced, error) {
	event := new(AuctionBidPlaced)
	if err := _Auction.contract.UnpackLog(event, "BidPlaced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuctionInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Auction contract.
type AuctionInitializedIterator struct {
	Event *AuctionInitialized // Event containing the contract specifics and raw log

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
func (it *AuctionInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionInitialized)
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
		it.Event = new(AuctionInitialized)
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
func (it *AuctionInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionInitialized represents a Initialized event raised by the Auction contract.
type AuctionInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Auction *AuctionFilterer) FilterInitialized(opts *bind.FilterOpts) (*AuctionInitializedIterator, error) {

	logs, sub, err := _Auction.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AuctionInitializedIterator{contract: _Auction.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Auction *AuctionFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AuctionInitialized) (event.Subscription, error) {

	logs, sub, err := _Auction.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionInitialized)
				if err := _Auction.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Auction *AuctionFilterer) ParseInitialized(log types.Log) (*AuctionInitialized, error) {
	event := new(AuctionInitialized)
	if err := _Auction.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuctionOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Auction contract.
type AuctionOwnershipTransferredIterator struct {
	Event *AuctionOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AuctionOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionOwnershipTransferred)
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
		it.Event = new(AuctionOwnershipTransferred)
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
func (it *AuctionOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionOwnershipTransferred represents a OwnershipTransferred event raised by the Auction contract.
type AuctionOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Auction *AuctionFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AuctionOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Auction.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AuctionOwnershipTransferredIterator{contract: _Auction.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Auction *AuctionFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AuctionOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Auction.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionOwnershipTransferred)
				if err := _Auction.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Auction *AuctionFilterer) ParseOwnershipTransferred(log types.Log) (*AuctionOwnershipTransferred, error) {
	event := new(AuctionOwnershipTransferred)
	if err := _Auction.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuctionUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Auction contract.
type AuctionUpgradedIterator struct {
	Event *AuctionUpgraded // Event containing the contract specifics and raw log

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
func (it *AuctionUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionUpgraded)
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
		it.Event = new(AuctionUpgraded)
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
func (it *AuctionUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionUpgraded represents a Upgraded event raised by the Auction contract.
type AuctionUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Auction *AuctionFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*AuctionUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Auction.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &AuctionUpgradedIterator{contract: _Auction.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Auction *AuctionFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *AuctionUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Auction.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionUpgraded)
				if err := _Auction.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Auction *AuctionFilterer) ParseUpgraded(log types.Log) (*AuctionUpgraded, error) {
	event := new(AuctionUpgraded)
	if err := _Auction.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuctionWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Auction contract.
type AuctionWithdrawIterator struct {
	Event *AuctionWithdraw // Event containing the contract specifics and raw log

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
func (it *AuctionWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionWithdraw)
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
		it.Event = new(AuctionWithdraw)
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
func (it *AuctionWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionWithdraw represents a Withdraw event raised by the Auction contract.
type AuctionWithdraw struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed user, uint256 amount)
func (_Auction *AuctionFilterer) FilterWithdraw(opts *bind.FilterOpts, user []common.Address) (*AuctionWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Auction.contract.FilterLogs(opts, "Withdraw", userRule)
	if err != nil {
		return nil, err
	}
	return &AuctionWithdrawIterator{contract: _Auction.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed user, uint256 amount)
func (_Auction *AuctionFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *AuctionWithdraw, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Auction.contract.WatchLogs(opts, "Withdraw", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionWithdraw)
				if err := _Auction.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed user, uint256 amount)
func (_Auction *AuctionFilterer) ParseWithdraw(log types.Log) (*AuctionWithdraw, error) {
	event := new(AuctionWithdraw)
	if err := _Auction.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
