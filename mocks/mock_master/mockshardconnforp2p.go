// Code generated by MockGen. DO NOT EDIT.
// Source: cluster/rpc/interface.go

// Package mock_master is a generated GoMock package.
package mock_master

import (
	account "github.com/QuarkChain/goquarkchain/account"
	rpc "github.com/QuarkChain/goquarkchain/cluster/rpc"
	consensus "github.com/QuarkChain/goquarkchain/consensus"
	types "github.com/QuarkChain/goquarkchain/core/types"
	p2p "github.com/QuarkChain/goquarkchain/p2p"
	common "github.com/ethereum/go-ethereum/common"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockShardConnForP2P is a mock of ShardConnForP2P interface
type MockShardConnForP2P struct {
	ctrl     *gomock.Controller
	recorder *MockShardConnForP2PMockRecorder
}

// MockShardConnForP2PMockRecorder is the mock recorder for MockShardConnForP2P
type MockShardConnForP2PMockRecorder struct {
	mock *MockShardConnForP2P
}

// NewMockShardConnForP2P creates a new mock instance
func NewMockShardConnForP2P(ctrl *gomock.Controller) *MockShardConnForP2P {
	mock := &MockShardConnForP2P{ctrl: ctrl}
	mock.recorder = &MockShardConnForP2PMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockShardConnForP2P) EXPECT() *MockShardConnForP2PMockRecorder {
	return m.recorder
}

// AddTransactions mocks base method
func (m *MockShardConnForP2P) AddTransactions(request *p2p.NewTransactionList) (*rpc.HashList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTransactions", request)
	ret0, _ := ret[0].(*rpc.HashList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddTransactions indicates an expected call of AddTransactions
func (mr *MockShardConnForP2PMockRecorder) AddTransactions(request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTransactions", reflect.TypeOf((*MockShardConnForP2P)(nil).AddTransactions), request)
}

// GetMinorBlocks mocks base method
func (m *MockShardConnForP2P) GetMinorBlocks(request *p2p.GetMinorBlockListRequest) (*p2p.GetMinorBlockListResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMinorBlocks", request)
	ret0, _ := ret[0].(*p2p.GetMinorBlockListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMinorBlocks indicates an expected call of GetMinorBlocks
func (mr *MockShardConnForP2PMockRecorder) GetMinorBlocks(request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMinorBlocks", reflect.TypeOf((*MockShardConnForP2P)(nil).GetMinorBlocks), request)
}

// GetMinorBlockHeaders mocks base method
func (m *MockShardConnForP2P) GetMinorBlockHeaders(request *p2p.GetMinorBlockHeaderListRequest) (*p2p.GetMinorBlockHeaderListResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMinorBlockHeaders", request)
	ret0, _ := ret[0].(*p2p.GetMinorBlockHeaderListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMinorBlockHeaders indicates an expected call of GetMinorBlockHeaders
func (mr *MockShardConnForP2PMockRecorder) GetMinorBlockHeaders(request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMinorBlockHeaders", reflect.TypeOf((*MockShardConnForP2P)(nil).GetMinorBlockHeaders), request)
}

// HandleNewTip mocks base method
func (m *MockShardConnForP2P) HandleNewTip(request *rpc.HandleNewTipRequest) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleNewTip", request)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HandleNewTip indicates an expected call of HandleNewTip
func (mr *MockShardConnForP2PMockRecorder) HandleNewTip(request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleNewTip", reflect.TypeOf((*MockShardConnForP2P)(nil).HandleNewTip), request)
}

// HandleNewMinorBlock mocks base method
func (m *MockShardConnForP2P) HandleNewMinorBlock(request *p2p.NewBlockMinor) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleNewMinorBlock", request)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HandleNewMinorBlock indicates an expected call of HandleNewMinorBlock
func (mr *MockShardConnForP2PMockRecorder) HandleNewMinorBlock(request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleNewMinorBlock", reflect.TypeOf((*MockShardConnForP2P)(nil).HandleNewMinorBlock), request)
}

// AddBlockListForSync mocks base method
func (m *MockShardConnForP2P) AddBlockListForSync(request *rpc.AddBlockListForSyncRequest) (*rpc.ShardStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddBlockListForSync", request)
	ret0, _ := ret[0].(*rpc.ShardStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddBlockListForSync indicates an expected call of AddBlockListForSync
func (mr *MockShardConnForP2PMockRecorder) AddBlockListForSync(request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddBlockListForSync", reflect.TypeOf((*MockShardConnForP2P)(nil).AddBlockListForSync), request)
}

// MockISlaveConn is a mock of ISlaveConn interface
type MockISlaveConn struct {
	ctrl     *gomock.Controller
	recorder *MockISlaveConnMockRecorder
}

// MockISlaveConnMockRecorder is the mock recorder for MockISlaveConn
type MockISlaveConnMockRecorder struct {
	mock *MockISlaveConn
}

// NewMockISlaveConn creates a new mock instance
func NewMockISlaveConn(ctrl *gomock.Controller) *MockISlaveConn {
	mock := &MockISlaveConn{ctrl: ctrl}
	mock.recorder = &MockISlaveConnMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockISlaveConn) EXPECT() *MockISlaveConnMockRecorder {
	return m.recorder
}

// AddTransactions mocks base method
func (m *MockISlaveConn) AddTransactions(request *p2p.NewTransactionList) (*rpc.HashList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTransactions", request)
	ret0, _ := ret[0].(*rpc.HashList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddTransactions indicates an expected call of AddTransactions
func (mr *MockISlaveConnMockRecorder) AddTransactions(request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTransactions", reflect.TypeOf((*MockISlaveConn)(nil).AddTransactions), request)
}

// GetMinorBlocks mocks base method
func (m *MockISlaveConn) GetMinorBlocks(request *p2p.GetMinorBlockListRequest) (*p2p.GetMinorBlockListResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMinorBlocks", request)
	ret0, _ := ret[0].(*p2p.GetMinorBlockListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMinorBlocks indicates an expected call of GetMinorBlocks
func (mr *MockISlaveConnMockRecorder) GetMinorBlocks(request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMinorBlocks", reflect.TypeOf((*MockISlaveConn)(nil).GetMinorBlocks), request)
}

// GetMinorBlockHeaders mocks base method
func (m *MockISlaveConn) GetMinorBlockHeaders(request *p2p.GetMinorBlockHeaderListRequest) (*p2p.GetMinorBlockHeaderListResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMinorBlockHeaders", request)
	ret0, _ := ret[0].(*p2p.GetMinorBlockHeaderListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMinorBlockHeaders indicates an expected call of GetMinorBlockHeaders
func (mr *MockISlaveConnMockRecorder) GetMinorBlockHeaders(request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMinorBlockHeaders", reflect.TypeOf((*MockISlaveConn)(nil).GetMinorBlockHeaders), request)
}

// HandleNewTip mocks base method
func (m *MockISlaveConn) HandleNewTip(request *rpc.HandleNewTipRequest) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleNewTip", request)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HandleNewTip indicates an expected call of HandleNewTip
func (mr *MockISlaveConnMockRecorder) HandleNewTip(request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleNewTip", reflect.TypeOf((*MockISlaveConn)(nil).HandleNewTip), request)
}

// HandleNewMinorBlock mocks base method
func (m *MockISlaveConn) HandleNewMinorBlock(request *p2p.NewBlockMinor) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleNewMinorBlock", request)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HandleNewMinorBlock indicates an expected call of HandleNewMinorBlock
func (mr *MockISlaveConnMockRecorder) HandleNewMinorBlock(request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleNewMinorBlock", reflect.TypeOf((*MockISlaveConn)(nil).HandleNewMinorBlock), request)
}

// AddBlockListForSync mocks base method
func (m *MockISlaveConn) AddBlockListForSync(request *rpc.AddBlockListForSyncRequest) (*rpc.ShardStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddBlockListForSync", request)
	ret0, _ := ret[0].(*rpc.ShardStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddBlockListForSync indicates an expected call of AddBlockListForSync
func (mr *MockISlaveConnMockRecorder) AddBlockListForSync(request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddBlockListForSync", reflect.TypeOf((*MockISlaveConn)(nil).AddBlockListForSync), request)
}

// GetSlaveID mocks base method
func (m *MockISlaveConn) GetSlaveID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSlaveID")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetSlaveID indicates an expected call of GetSlaveID
func (mr *MockISlaveConnMockRecorder) GetSlaveID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSlaveID", reflect.TypeOf((*MockISlaveConn)(nil).GetSlaveID))
}

// GetShardMaskList mocks base method
func (m *MockISlaveConn) GetShardMaskList() []*types.ChainMask {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetShardMaskList")
	ret0, _ := ret[0].([]*types.ChainMask)
	return ret0
}

// GetShardMaskList indicates an expected call of GetShardMaskList
func (mr *MockISlaveConnMockRecorder) GetShardMaskList() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetShardMaskList", reflect.TypeOf((*MockISlaveConn)(nil).GetShardMaskList))
}

// MasterInfo mocks base method
func (m *MockISlaveConn) MasterInfo(ip string, port uint16, rootTip *types.RootBlock) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MasterInfo", ip, port, rootTip)
	ret0, _ := ret[0].(error)
	return ret0
}

// MasterInfo indicates an expected call of MasterInfo
func (mr *MockISlaveConnMockRecorder) MasterInfo(ip, port, rootTip interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MasterInfo", reflect.TypeOf((*MockISlaveConn)(nil).MasterInfo), ip, port, rootTip)
}

// HasShard mocks base method
func (m *MockISlaveConn) HasShard(fullShardID uint32) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasShard", fullShardID)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasShard indicates an expected call of HasShard
func (mr *MockISlaveConnMockRecorder) HasShard(fullShardID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasShard", reflect.TypeOf((*MockISlaveConn)(nil).HasShard), fullShardID)
}

// SendPing mocks base method
func (m *MockISlaveConn) SendPing() ([]byte, []*types.ChainMask, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendPing")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].([]*types.ChainMask)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// SendPing indicates an expected call of SendPing
func (mr *MockISlaveConnMockRecorder) SendPing() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendPing", reflect.TypeOf((*MockISlaveConn)(nil).SendPing))
}

// HeartBeat mocks base method
func (m *MockISlaveConn) HeartBeat() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HeartBeat")
	ret0, _ := ret[0].(bool)
	return ret0
}

// HeartBeat indicates an expected call of HeartBeat
func (mr *MockISlaveConnMockRecorder) HeartBeat() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HeartBeat", reflect.TypeOf((*MockISlaveConn)(nil).HeartBeat))
}

// GetUnconfirmedHeaders mocks base method
func (m *MockISlaveConn) GetUnconfirmedHeaders() (*rpc.GetUnconfirmedHeadersResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUnconfirmedHeaders")
	ret0, _ := ret[0].(*rpc.GetUnconfirmedHeadersResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUnconfirmedHeaders indicates an expected call of GetUnconfirmedHeaders
func (mr *MockISlaveConnMockRecorder) GetUnconfirmedHeaders() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUnconfirmedHeaders", reflect.TypeOf((*MockISlaveConn)(nil).GetUnconfirmedHeaders))
}

// GetAccountData mocks base method
func (m *MockISlaveConn) GetAccountData(address *account.Address, height *uint64) (*rpc.GetAccountDataResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountData", address, height)
	ret0, _ := ret[0].(*rpc.GetAccountDataResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountData indicates an expected call of GetAccountData
func (mr *MockISlaveConnMockRecorder) GetAccountData(address, height interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountData", reflect.TypeOf((*MockISlaveConn)(nil).GetAccountData), address, height)
}

// AddRootBlock mocks base method
func (m *MockISlaveConn) AddRootBlock(rootBlock *types.RootBlock, expectSwitch bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddRootBlock", rootBlock, expectSwitch)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddRootBlock indicates an expected call of AddRootBlock
func (mr *MockISlaveConnMockRecorder) AddRootBlock(rootBlock, expectSwitch interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRootBlock", reflect.TypeOf((*MockISlaveConn)(nil).AddRootBlock), rootBlock, expectSwitch)
}

// GenTx mocks base method
func (m *MockISlaveConn) GenTx(numTxPerShard, xShardPercent uint32, tx *types.Transaction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenTx", numTxPerShard, xShardPercent, tx)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenTx indicates an expected call of GenTx
func (mr *MockISlaveConnMockRecorder) GenTx(numTxPerShard, xShardPercent, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenTx", reflect.TypeOf((*MockISlaveConn)(nil).GenTx), numTxPerShard, xShardPercent, tx)
}

// SendMiningConfigToSlaves mocks base method
func (m *MockISlaveConn) SendMiningConfigToSlaves(artificialTxConfig *rpc.ArtificialTxConfig, mining bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMiningConfigToSlaves", artificialTxConfig, mining)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMiningConfigToSlaves indicates an expected call of SendMiningConfigToSlaves
func (mr *MockISlaveConnMockRecorder) SendMiningConfigToSlaves(artificialTxConfig, mining interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMiningConfigToSlaves", reflect.TypeOf((*MockISlaveConn)(nil).SendMiningConfigToSlaves), artificialTxConfig, mining)
}

// AddTransaction mocks base method
func (m *MockISlaveConn) AddTransaction(tx *types.Transaction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTransaction", tx)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddTransaction indicates an expected call of AddTransaction
func (mr *MockISlaveConnMockRecorder) AddTransaction(tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTransaction", reflect.TypeOf((*MockISlaveConn)(nil).AddTransaction), tx)
}

// ExecuteTransaction mocks base method
func (m *MockISlaveConn) ExecuteTransaction(tx *types.Transaction, fromAddress *account.Address, height *uint64) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExecuteTransaction", tx, fromAddress, height)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExecuteTransaction indicates an expected call of ExecuteTransaction
func (mr *MockISlaveConnMockRecorder) ExecuteTransaction(tx, fromAddress, height interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecuteTransaction", reflect.TypeOf((*MockISlaveConn)(nil).ExecuteTransaction), tx, fromAddress, height)
}

// GetMinorBlockByHash mocks base method
func (m *MockISlaveConn) GetMinorBlockByHash(blockHash common.Hash, branch account.Branch) (*types.MinorBlock, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMinorBlockByHash", blockHash, branch)
	ret0, _ := ret[0].(*types.MinorBlock)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMinorBlockByHash indicates an expected call of GetMinorBlockByHash
func (mr *MockISlaveConnMockRecorder) GetMinorBlockByHash(blockHash, branch interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMinorBlockByHash", reflect.TypeOf((*MockISlaveConn)(nil).GetMinorBlockByHash), blockHash, branch)
}

// GetMinorBlockByHeight mocks base method
func (m *MockISlaveConn) GetMinorBlockByHeight(height uint64, branch account.Branch) (*types.MinorBlock, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMinorBlockByHeight", height, branch)
	ret0, _ := ret[0].(*types.MinorBlock)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMinorBlockByHeight indicates an expected call of GetMinorBlockByHeight
func (mr *MockISlaveConnMockRecorder) GetMinorBlockByHeight(height, branch interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMinorBlockByHeight", reflect.TypeOf((*MockISlaveConn)(nil).GetMinorBlockByHeight), height, branch)
}

// GetTransactionByHash mocks base method
func (m *MockISlaveConn) GetTransactionByHash(txHash common.Hash, branch account.Branch) (*types.MinorBlock, uint32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactionByHash", txHash, branch)
	ret0, _ := ret[0].(*types.MinorBlock)
	ret1, _ := ret[1].(uint32)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetTransactionByHash indicates an expected call of GetTransactionByHash
func (mr *MockISlaveConnMockRecorder) GetTransactionByHash(txHash, branch interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionByHash", reflect.TypeOf((*MockISlaveConn)(nil).GetTransactionByHash), txHash, branch)
}

// GetTransactionReceipt mocks base method
func (m *MockISlaveConn) GetTransactionReceipt(txHash common.Hash, branch account.Branch) (*types.MinorBlock, uint32, *types.Receipt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactionReceipt", txHash, branch)
	ret0, _ := ret[0].(*types.MinorBlock)
	ret1, _ := ret[1].(uint32)
	ret2, _ := ret[2].(*types.Receipt)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// GetTransactionReceipt indicates an expected call of GetTransactionReceipt
func (mr *MockISlaveConnMockRecorder) GetTransactionReceipt(txHash, branch interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionReceipt", reflect.TypeOf((*MockISlaveConn)(nil).GetTransactionReceipt), txHash, branch)
}

// GetTransactionsByAddress mocks base method
func (m *MockISlaveConn) GetTransactionsByAddress(address *account.Address, start []byte, limit uint32) ([]*rpc.TransactionDetail, []byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactionsByAddress", address, start, limit)
	ret0, _ := ret[0].([]*rpc.TransactionDetail)
	ret1, _ := ret[1].([]byte)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetTransactionsByAddress indicates an expected call of GetTransactionsByAddress
func (mr *MockISlaveConnMockRecorder) GetTransactionsByAddress(address, start, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionsByAddress", reflect.TypeOf((*MockISlaveConn)(nil).GetTransactionsByAddress), address, start, limit)
}

// GetLogs mocks base method
func (m *MockISlaveConn) GetLogs(branch account.Branch, address []*account.Address, topics []*rpc.Topic, startBlock, endBlock uint64) ([]*types.Log, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLogs", branch, address, topics, startBlock, endBlock)
	ret0, _ := ret[0].([]*types.Log)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLogs indicates an expected call of GetLogs
func (mr *MockISlaveConnMockRecorder) GetLogs(branch, address, topics, startBlock, endBlock interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLogs", reflect.TypeOf((*MockISlaveConn)(nil).GetLogs), branch, address, topics, startBlock, endBlock)
}

// EstimateGas mocks base method
func (m *MockISlaveConn) EstimateGas(tx *types.Transaction, fromAddress *account.Address) (uint32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EstimateGas", tx, fromAddress)
	ret0, _ := ret[0].(uint32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EstimateGas indicates an expected call of EstimateGas
func (mr *MockISlaveConnMockRecorder) EstimateGas(tx, fromAddress interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EstimateGas", reflect.TypeOf((*MockISlaveConn)(nil).EstimateGas), tx, fromAddress)
}

// GetStorageAt mocks base method
func (m *MockISlaveConn) GetStorageAt(address *account.Address, key common.Hash, height *uint64) (common.Hash, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStorageAt", address, key, height)
	ret0, _ := ret[0].(common.Hash)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStorageAt indicates an expected call of GetStorageAt
func (mr *MockISlaveConnMockRecorder) GetStorageAt(address, key, height interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStorageAt", reflect.TypeOf((*MockISlaveConn)(nil).GetStorageAt), address, key, height)
}

// GetCode mocks base method
func (m *MockISlaveConn) GetCode(address *account.Address, height *uint64) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCode", address, height)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCode indicates an expected call of GetCode
func (mr *MockISlaveConnMockRecorder) GetCode(address, height interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCode", reflect.TypeOf((*MockISlaveConn)(nil).GetCode), address, height)
}

// GasPrice mocks base method
func (m *MockISlaveConn) GasPrice(branch account.Branch) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GasPrice", branch)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GasPrice indicates an expected call of GasPrice
func (mr *MockISlaveConnMockRecorder) GasPrice(branch interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GasPrice", reflect.TypeOf((*MockISlaveConn)(nil).GasPrice), branch)
}

// GetWork mocks base method
func (m *MockISlaveConn) GetWork(branch account.Branch) (*consensus.MiningWork, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWork", branch)
	ret0, _ := ret[0].(*consensus.MiningWork)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWork indicates an expected call of GetWork
func (mr *MockISlaveConnMockRecorder) GetWork(branch interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWork", reflect.TypeOf((*MockISlaveConn)(nil).GetWork), branch)
}

// SubmitWork mocks base method
func (m *MockISlaveConn) SubmitWork(work *rpc.SubmitWorkRequest) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubmitWork", work)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubmitWork indicates an expected call of SubmitWork
func (mr *MockISlaveConnMockRecorder) SubmitWork(work interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitWork", reflect.TypeOf((*MockISlaveConn)(nil).SubmitWork), work)
}
