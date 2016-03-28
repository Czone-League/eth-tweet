// This file is an automatically generated Go binding. Do not modify as any
// change will likely be lost upon the next re-generation!

package main

import (
	"math/big"
	"strings"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi"
)

// TweetAccountABI is the input ABI used to generate the binding from.
const TweetAccountABI = `[{"constant":true,"inputs":[],"name":"getOwnerAddress","outputs":[{"name":"adminAddress","type":"address"}],"type":"function"},{"constant":false,"inputs":[],"name":"adminDeleteAccount","outputs":[],"type":"function"},{"constant":false,"inputs":[{"name":"receiver","type":"address"}],"name":"adminRetrieveDonations","outputs":[],"type":"function"},{"constant":true,"inputs":[],"name":"getLatestTweet","outputs":[{"name":"tweetString","type":"string"},{"name":"timestamp","type":"uint256"},{"name":"numberOfTweets","type":"uint256"}],"type":"function"},{"constant":true,"inputs":[],"name":"isAdmin","outputs":[{"name":"isAdmin","type":"bool"}],"type":"function"},{"constant":true,"inputs":[{"name":"tweetId","type":"uint256"}],"name":"getTweet","outputs":[{"name":"tweetString","type":"string"},{"name":"timestamp","type":"uint256"}],"type":"function"},{"constant":true,"inputs":[],"name":"getNumberOfTweets","outputs":[{"name":"numberOfTweets","type":"uint256"}],"type":"function"},{"constant":false,"inputs":[{"name":"tweetString","type":"string"}],"name":"tweet","outputs":[{"name":"result","type":"int256"}],"type":"function"},{"inputs":[],"type":"constructor"}]`

// TweetAccountBin is the compiled bytecode used for deploying new contracts.
const TweetAccountBin = `0x6060604052600060015560028054600160a060020a031916331790556104f4806100296000396000f36060604052361561006c5760e060020a60003504630c4f65bd811461006e5780633e450fff146100815780635c3e426c1461008c578063ae978f081461009a578063b6db75a014610118578063c3ad5ecb14610134578063ca7dc5b1146101b0578063fb46d4c5146101bb575b005b61020f600254600160a060020a03165b90565b61006c6104df61011c565b61006c6004356104b361011c565b604080516020818101835260008083526001805460001990810183528284528551868420830180546002948116156101000290930190921692909204601f810185900485028301850190965285825261022c95929384939083018282801561048b5780601f106104605761010080835404028352916020019161048b565b6102a85b600254600160a060020a03908116339091161461007e565b6040805160208181018352600080835260043580825281835284518583206001908101805460029281161561010002600019011691909104601f81018690048602830186019097528682526102ba96929594919290918301828280156104445780601f1061041957610100808354040283529160200191610444565b6102a860015461007e565b6102a86004808035906020019082018035906020019191908080601f01602080910402602001604051908101604052809392919081815260200183838082843750949650505050505050600061032f61011c565b60408051600160a060020a03929092168252519081900360200190f35b60405180806020018481526020018381526020018281038252858181518152602001915080519060200190808383829060006004602084601f0104600302600f01f150905090810190601f1680156102985780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b60408051918252519081900360200190f35b60405180806020018381526020018281038252848181518152602001915080519060200190808383829060006004602084601f0104600302600f01f150905090810190601f1680156103205780820380516001836020036101000a031916815260200191505b50935050505060405180910390f35b151561033e57506000196103e0565b60a08251111561035157506001196103e0565b6001805460009081526020818152604080832042905583548352822085519084018054818552938390209094600290851615610100026000190190941693909304601f90810183900484019391928701908390106103e557805160ff19168380011785555b506103d29291505b8082111561041557600081556001016103be565b505060018054810190555060005b919050565b828001600101855582156103b6579182015b828111156103b65782518260005055916020019190600101906103f7565b5090565b820191906000526020600020905b81548152906001019060200180831161042757829003601f168201915b5050506000958652505060208490526040909320549293915050565b820191906000526020600020905b81548152906001019060200180831161046e57829003601f168201915b5050600154600019810160009081526020819052604090205494989497509550929350505050565b156104dc57604051600160a060020a03828116916000913016319082818181858883f150505050505b50565b156104f257600254600160a060020a0316ff5b56`

// DeployTweetAccount deploys a new Ethereum contract, binding an instance of TweetAccount to it.
func DeployTweetAccount(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TweetAccount, error) {
	parsed, err := abi.JSON(strings.NewReader(TweetAccountABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TweetAccountBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TweetAccount{TweetAccountCaller: TweetAccountCaller{contract: contract}, TweetAccountTransactor: TweetAccountTransactor{contract: contract}}, nil
}

// TweetAccount is an auto generated Go binding around an Ethereum contract.
type TweetAccount struct {
	TweetAccountCaller     // Read-only binding to the contract
	TweetAccountTransactor // Write-only binding to the contract
}

// TweetAccountCaller is an auto generated read-only Go binding around an Ethereum contract.
type TweetAccountCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TweetAccountTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TweetAccountTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TweetAccountSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TweetAccountSession struct {
	Contract     *TweetAccount     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TweetAccountCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TweetAccountCallerSession struct {
	Contract *TweetAccountCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// TweetAccountTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TweetAccountTransactorSession struct {
	Contract     *TweetAccountTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// NewTweetAccount creates a new instance of TweetAccount, bound to a specific deployed contract.
func NewTweetAccount(address common.Address, backend bind.ContractBackend) (*TweetAccount, error) {
	contract, err := bindTweetAccount(address, backend.(bind.ContractCaller), backend.(bind.ContractTransactor))
	if err != nil {
		return nil, err
	}
	return &TweetAccount{TweetAccountCaller: TweetAccountCaller{contract: contract}, TweetAccountTransactor: TweetAccountTransactor{contract: contract}}, nil
}

// NewTweetAccountCaller creates a new read-only instance of TweetAccount, bound to a specific deployed contract.
func NewTweetAccountCaller(address common.Address, caller bind.ContractCaller) (*TweetAccountCaller, error) {
	contract, err := bindTweetAccount(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &TweetAccountCaller{contract: contract}, nil
}

// NewTweetAccountTransactor creates a new write-only instance of TweetAccount, bound to a specific deployed contract.
func NewTweetAccountTransactor(address common.Address, transactor bind.ContractTransactor) (*TweetAccountTransactor, error) {
	contract, err := bindTweetAccount(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &TweetAccountTransactor{contract: contract}, nil
}

// bindTweetAccount binds a generic wrapper to an already deployed contract.
func bindTweetAccount(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TweetAccountABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// GetLatestTweetResult is the result of the GetLatestTweet invocation."
type GetLatestTweetResult struct {
	TweetString    string
	Timestamp      *big.Int
	NumberOfTweets *big.Int
}

// GetLatestTweet is a free data retrieval call binding the contract method 0xae978f08.
//
// Solidity: function getLatestTweet() constant returns(tweetString string, timestamp uint256, numberOfTweets uint256)
func (_TweetAccount *TweetAccountCaller) GetLatestTweet(opts *bind.CallOpts) (GetLatestTweetResult, error) {
	var (
		ret = new(GetLatestTweetResult)
	)
	out := ret
	err := _TweetAccount.contract.Call(opts, out, "getLatestTweet")
	return *ret, err
}

// GetLatestTweet is a free data retrieval call binding the contract method 0xae978f08.
//
// Solidity: function getLatestTweet() constant returns(tweetString string, timestamp uint256, numberOfTweets uint256)
func (_TweetAccount *TweetAccountSession) GetLatestTweet() (GetLatestTweetResult, error) {
	return _TweetAccount.Contract.GetLatestTweet(&_TweetAccount.CallOpts)
}

// GetLatestTweet is a free data retrieval call binding the contract method 0xae978f08.
//
// Solidity: function getLatestTweet() constant returns(tweetString string, timestamp uint256, numberOfTweets uint256)
func (_TweetAccount *TweetAccountCallerSession) GetLatestTweet() (GetLatestTweetResult, error) {
	return _TweetAccount.Contract.GetLatestTweet(&_TweetAccount.CallOpts)
}

// GetNumberOfTweets is a free data retrieval call binding the contract method 0xca7dc5b1.
//
// Solidity: function getNumberOfTweets() constant returns(numberOfTweets uint256)
func (_TweetAccount *TweetAccountCaller) GetNumberOfTweets(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TweetAccount.contract.Call(opts, out, "getNumberOfTweets")
	return *ret0, err
}

// GetNumberOfTweets is a free data retrieval call binding the contract method 0xca7dc5b1.
//
// Solidity: function getNumberOfTweets() constant returns(numberOfTweets uint256)
func (_TweetAccount *TweetAccountSession) GetNumberOfTweets() (*big.Int, error) {
	return _TweetAccount.Contract.GetNumberOfTweets(&_TweetAccount.CallOpts)
}

// GetNumberOfTweets is a free data retrieval call binding the contract method 0xca7dc5b1.
//
// Solidity: function getNumberOfTweets() constant returns(numberOfTweets uint256)
func (_TweetAccount *TweetAccountCallerSession) GetNumberOfTweets() (*big.Int, error) {
	return _TweetAccount.Contract.GetNumberOfTweets(&_TweetAccount.CallOpts)
}

// GetOwnerAddress is a free data retrieval call binding the contract method 0x0c4f65bd.
//
// Solidity: function getOwnerAddress() constant returns(adminAddress address)
func (_TweetAccount *TweetAccountCaller) GetOwnerAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TweetAccount.contract.Call(opts, out, "getOwnerAddress")
	return *ret0, err
}

// GetOwnerAddress is a free data retrieval call binding the contract method 0x0c4f65bd.
//
// Solidity: function getOwnerAddress() constant returns(adminAddress address)
func (_TweetAccount *TweetAccountSession) GetOwnerAddress() (common.Address, error) {
	return _TweetAccount.Contract.GetOwnerAddress(&_TweetAccount.CallOpts)
}

// GetOwnerAddress is a free data retrieval call binding the contract method 0x0c4f65bd.
//
// Solidity: function getOwnerAddress() constant returns(adminAddress address)
func (_TweetAccount *TweetAccountCallerSession) GetOwnerAddress() (common.Address, error) {
	return _TweetAccount.Contract.GetOwnerAddress(&_TweetAccount.CallOpts)
}

// GetTweetResult is the result of the GetTweet invocation."
type GetTweetResult struct {
	TweetString string
	Timestamp   *big.Int
}

// GetTweet is a free data retrieval call binding the contract method 0xc3ad5ecb.
//
// Solidity: function getTweet(tweetId uint256) constant returns(tweetString string, timestamp uint256)
func (_TweetAccount *TweetAccountCaller) GetTweet(opts *bind.CallOpts, tweetId *big.Int) (GetTweetResult, error) {
	var (
		ret = new(GetTweetResult)
	)
	out := ret
	err := _TweetAccount.contract.Call(opts, out, "getTweet", tweetId)
	return *ret, err
}

// GetTweet is a free data retrieval call binding the contract method 0xc3ad5ecb.
//
// Solidity: function getTweet(tweetId uint256) constant returns(tweetString string, timestamp uint256)
func (_TweetAccount *TweetAccountSession) GetTweet(tweetId *big.Int) (GetTweetResult, error) {
	return _TweetAccount.Contract.GetTweet(&_TweetAccount.CallOpts, tweetId)
}

// GetTweet is a free data retrieval call binding the contract method 0xc3ad5ecb.
//
// Solidity: function getTweet(tweetId uint256) constant returns(tweetString string, timestamp uint256)
func (_TweetAccount *TweetAccountCallerSession) GetTweet(tweetId *big.Int) (GetTweetResult, error) {
	return _TweetAccount.Contract.GetTweet(&_TweetAccount.CallOpts, tweetId)
}

// IsAdmin is a free data retrieval call binding the contract method 0xb6db75a0.
//
// Solidity: function isAdmin() constant returns(isAdmin bool)
func (_TweetAccount *TweetAccountCaller) IsAdmin(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TweetAccount.contract.Call(opts, out, "isAdmin")
	return *ret0, err
}

// IsAdmin is a free data retrieval call binding the contract method 0xb6db75a0.
//
// Solidity: function isAdmin() constant returns(isAdmin bool)
func (_TweetAccount *TweetAccountSession) IsAdmin() (bool, error) {
	return _TweetAccount.Contract.IsAdmin(&_TweetAccount.CallOpts)
}

// IsAdmin is a free data retrieval call binding the contract method 0xb6db75a0.
//
// Solidity: function isAdmin() constant returns(isAdmin bool)
func (_TweetAccount *TweetAccountCallerSession) IsAdmin() (bool, error) {
	return _TweetAccount.Contract.IsAdmin(&_TweetAccount.CallOpts)
}

// AdminDeleteAccount is a paid mutator transaction binding the contract method 0x3e450fff.
//
// Solidity: function adminDeleteAccount() returns()
func (_TweetAccount *TweetAccountTransactor) AdminDeleteAccount(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TweetAccount.contract.Transact(opts, "adminDeleteAccount")
}

// AdminDeleteAccount is a paid mutator transaction binding the contract method 0x3e450fff.
//
// Solidity: function adminDeleteAccount() returns()
func (_TweetAccount *TweetAccountSession) AdminDeleteAccount() (*types.Transaction, error) {
	return _TweetAccount.Contract.AdminDeleteAccount(&_TweetAccount.TransactOpts)
}

// AdminDeleteAccount is a paid mutator transaction binding the contract method 0x3e450fff.
//
// Solidity: function adminDeleteAccount() returns()
func (_TweetAccount *TweetAccountTransactorSession) AdminDeleteAccount() (*types.Transaction, error) {
	return _TweetAccount.Contract.AdminDeleteAccount(&_TweetAccount.TransactOpts)
}

// AdminRetrieveDonations is a paid mutator transaction binding the contract method 0x5c3e426c.
//
// Solidity: function adminRetrieveDonations(receiver address) returns()
func (_TweetAccount *TweetAccountTransactor) AdminRetrieveDonations(opts *bind.TransactOpts, receiver common.Address) (*types.Transaction, error) {
	return _TweetAccount.contract.Transact(opts, "adminRetrieveDonations", receiver)
}

// AdminRetrieveDonations is a paid mutator transaction binding the contract method 0x5c3e426c.
//
// Solidity: function adminRetrieveDonations(receiver address) returns()
func (_TweetAccount *TweetAccountSession) AdminRetrieveDonations(receiver common.Address) (*types.Transaction, error) {
	return _TweetAccount.Contract.AdminRetrieveDonations(&_TweetAccount.TransactOpts, receiver)
}

// AdminRetrieveDonations is a paid mutator transaction binding the contract method 0x5c3e426c.
//
// Solidity: function adminRetrieveDonations(receiver address) returns()
func (_TweetAccount *TweetAccountTransactorSession) AdminRetrieveDonations(receiver common.Address) (*types.Transaction, error) {
	return _TweetAccount.Contract.AdminRetrieveDonations(&_TweetAccount.TransactOpts, receiver)
}

// Tweet is a paid mutator transaction binding the contract method 0xfb46d4c5.
//
// Solidity: function tweet(tweetString string) returns(result int256)
func (_TweetAccount *TweetAccountTransactor) Tweet(opts *bind.TransactOpts, tweetString string) (*types.Transaction, error) {
	return _TweetAccount.contract.Transact(opts, "tweet", tweetString)
}

// Tweet is a paid mutator transaction binding the contract method 0xfb46d4c5.
//
// Solidity: function tweet(tweetString string) returns(result int256)
func (_TweetAccount *TweetAccountSession) Tweet(tweetString string) (*types.Transaction, error) {
	return _TweetAccount.Contract.Tweet(&_TweetAccount.TransactOpts, tweetString)
}

// Tweet is a paid mutator transaction binding the contract method 0xfb46d4c5.
//
// Solidity: function tweet(tweetString string) returns(result int256)
func (_TweetAccount *TweetAccountTransactorSession) Tweet(tweetString string) (*types.Transaction, error) {
	return _TweetAccount.Contract.Tweet(&_TweetAccount.TransactOpts, tweetString)
}
