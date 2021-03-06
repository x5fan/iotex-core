// Code generated by MockGen. DO NOT EDIT.
// Source: ./explorer/idl/explorer/explorer.go

// Package mock_explorer is a generated GoMock package.
package mock_explorer

import (
	gomock "github.com/golang/mock/gomock"
	explorer "github.com/iotexproject/iotex-core/explorer/idl/explorer"
	reflect "reflect"
)

// MockExplorer is a mock of Explorer interface
type MockExplorer struct {
	ctrl     *gomock.Controller
	recorder *MockExplorerMockRecorder
}

// MockExplorerMockRecorder is the mock recorder for MockExplorer
type MockExplorerMockRecorder struct {
	mock *MockExplorer
}

// NewMockExplorer creates a new mock instance
func NewMockExplorer(ctrl *gomock.Controller) *MockExplorer {
	mock := &MockExplorer{ctrl: ctrl}
	mock.recorder = &MockExplorerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockExplorer) EXPECT() *MockExplorerMockRecorder {
	return m.recorder
}

// GetBlockchainHeight mocks base method
func (m *MockExplorer) GetBlockchainHeight() (int64, error) {
	ret := m.ctrl.Call(m, "GetBlockchainHeight")
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlockchainHeight indicates an expected call of GetBlockchainHeight
func (mr *MockExplorerMockRecorder) GetBlockchainHeight() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockchainHeight", reflect.TypeOf((*MockExplorer)(nil).GetBlockchainHeight))
}

// GetAddressBalance mocks base method
func (m *MockExplorer) GetAddressBalance(address string) (string, error) {
	ret := m.ctrl.Call(m, "GetAddressBalance", address)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAddressBalance indicates an expected call of GetAddressBalance
func (mr *MockExplorerMockRecorder) GetAddressBalance(address interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAddressBalance", reflect.TypeOf((*MockExplorer)(nil).GetAddressBalance), address)
}

// GetAddressDetails mocks base method
func (m *MockExplorer) GetAddressDetails(address string) (explorer.AddressDetails, error) {
	ret := m.ctrl.Call(m, "GetAddressDetails", address)
	ret0, _ := ret[0].(explorer.AddressDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAddressDetails indicates an expected call of GetAddressDetails
func (mr *MockExplorerMockRecorder) GetAddressDetails(address interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAddressDetails", reflect.TypeOf((*MockExplorer)(nil).GetAddressDetails), address)
}

// GetCreateDeposit mocks base method
func (m *MockExplorer) GetCreateDeposit(createDepositID string) (explorer.CreateDeposit, error) {
	ret := m.ctrl.Call(m, "GetCreateDeposit", createDepositID)
	ret0, _ := ret[0].(explorer.CreateDeposit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCreateDeposit indicates an expected call of GetCreateDeposit
func (mr *MockExplorerMockRecorder) GetCreateDeposit(createDepositID interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCreateDeposit", reflect.TypeOf((*MockExplorer)(nil).GetCreateDeposit), createDepositID)
}

// GetCreateDepositsByAddress mocks base method
func (m *MockExplorer) GetCreateDepositsByAddress(address string, offset, limit int64) ([]explorer.CreateDeposit, error) {
	ret := m.ctrl.Call(m, "GetCreateDepositsByAddress", address, offset, limit)
	ret0, _ := ret[0].([]explorer.CreateDeposit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCreateDepositsByAddress indicates an expected call of GetCreateDepositsByAddress
func (mr *MockExplorerMockRecorder) GetCreateDepositsByAddress(address, offset, limit interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCreateDepositsByAddress", reflect.TypeOf((*MockExplorer)(nil).GetCreateDepositsByAddress), address, offset, limit)
}

// GetSettleDeposit mocks base method
func (m *MockExplorer) GetSettleDeposit(settleDepositID string) (explorer.SettleDeposit, error) {
	ret := m.ctrl.Call(m, "GetSettleDeposit", settleDepositID)
	ret0, _ := ret[0].(explorer.SettleDeposit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSettleDeposit indicates an expected call of GetSettleDeposit
func (mr *MockExplorerMockRecorder) GetSettleDeposit(settleDepositID interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSettleDeposit", reflect.TypeOf((*MockExplorer)(nil).GetSettleDeposit), settleDepositID)
}

// GetSettleDepositsByAddress mocks base method
func (m *MockExplorer) GetSettleDepositsByAddress(address string, offset, limit int64) ([]explorer.SettleDeposit, error) {
	ret := m.ctrl.Call(m, "GetSettleDepositsByAddress", address, offset, limit)
	ret0, _ := ret[0].([]explorer.SettleDeposit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSettleDepositsByAddress indicates an expected call of GetSettleDepositsByAddress
func (mr *MockExplorerMockRecorder) GetSettleDepositsByAddress(address, offset, limit interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSettleDepositsByAddress", reflect.TypeOf((*MockExplorer)(nil).GetSettleDepositsByAddress), address, offset, limit)
}

// GetLastBlocksByRange mocks base method
func (m *MockExplorer) GetLastBlocksByRange(offset, limit int64) ([]explorer.Block, error) {
	ret := m.ctrl.Call(m, "GetLastBlocksByRange", offset, limit)
	ret0, _ := ret[0].([]explorer.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLastBlocksByRange indicates an expected call of GetLastBlocksByRange
func (mr *MockExplorerMockRecorder) GetLastBlocksByRange(offset, limit interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLastBlocksByRange", reflect.TypeOf((*MockExplorer)(nil).GetLastBlocksByRange), offset, limit)
}

// GetBlockByID mocks base method
func (m *MockExplorer) GetBlockByID(blkID string) (explorer.Block, error) {
	ret := m.ctrl.Call(m, "GetBlockByID", blkID)
	ret0, _ := ret[0].(explorer.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlockByID indicates an expected call of GetBlockByID
func (mr *MockExplorerMockRecorder) GetBlockByID(blkID interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockByID", reflect.TypeOf((*MockExplorer)(nil).GetBlockByID), blkID)
}

// GetCoinStatistic mocks base method
func (m *MockExplorer) GetCoinStatistic() (explorer.CoinStatistic, error) {
	ret := m.ctrl.Call(m, "GetCoinStatistic")
	ret0, _ := ret[0].(explorer.CoinStatistic)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCoinStatistic indicates an expected call of GetCoinStatistic
func (mr *MockExplorerMockRecorder) GetCoinStatistic() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCoinStatistic", reflect.TypeOf((*MockExplorer)(nil).GetCoinStatistic))
}

// GetConsensusMetrics mocks base method
func (m *MockExplorer) GetConsensusMetrics() (explorer.ConsensusMetrics, error) {
	ret := m.ctrl.Call(m, "GetConsensusMetrics")
	ret0, _ := ret[0].(explorer.ConsensusMetrics)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConsensusMetrics indicates an expected call of GetConsensusMetrics
func (mr *MockExplorerMockRecorder) GetConsensusMetrics() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConsensusMetrics", reflect.TypeOf((*MockExplorer)(nil).GetConsensusMetrics))
}

// GetCandidateMetrics mocks base method
func (m *MockExplorer) GetCandidateMetrics() (explorer.CandidateMetrics, error) {
	ret := m.ctrl.Call(m, "GetCandidateMetrics")
	ret0, _ := ret[0].(explorer.CandidateMetrics)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCandidateMetrics indicates an expected call of GetCandidateMetrics
func (mr *MockExplorerMockRecorder) GetCandidateMetrics() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCandidateMetrics", reflect.TypeOf((*MockExplorer)(nil).GetCandidateMetrics))
}

// GetCandidateMetricsByHeight mocks base method
func (m *MockExplorer) GetCandidateMetricsByHeight(h int64) (explorer.CandidateMetrics, error) {
	ret := m.ctrl.Call(m, "GetCandidateMetricsByHeight", h)
	ret0, _ := ret[0].(explorer.CandidateMetrics)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCandidateMetricsByHeight indicates an expected call of GetCandidateMetricsByHeight
func (mr *MockExplorerMockRecorder) GetCandidateMetricsByHeight(h interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCandidateMetricsByHeight", reflect.TypeOf((*MockExplorer)(nil).GetCandidateMetricsByHeight), h)
}

// SendTransfer mocks base method
func (m *MockExplorer) SendTransfer(request explorer.SendTransferRequest) (explorer.SendTransferResponse, error) {
	ret := m.ctrl.Call(m, "SendTransfer", request)
	ret0, _ := ret[0].(explorer.SendTransferResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendTransfer indicates an expected call of SendTransfer
func (mr *MockExplorerMockRecorder) SendTransfer(request interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendTransfer", reflect.TypeOf((*MockExplorer)(nil).SendTransfer), request)
}

// SendVote mocks base method
func (m *MockExplorer) SendVote(request explorer.SendVoteRequest) (explorer.SendVoteResponse, error) {
	ret := m.ctrl.Call(m, "SendVote", request)
	ret0, _ := ret[0].(explorer.SendVoteResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendVote indicates an expected call of SendVote
func (mr *MockExplorerMockRecorder) SendVote(request interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendVote", reflect.TypeOf((*MockExplorer)(nil).SendVote), request)
}

// SendSmartContract mocks base method
func (m *MockExplorer) SendSmartContract(request explorer.Execution) (explorer.SendSmartContractResponse, error) {
	ret := m.ctrl.Call(m, "SendSmartContract", request)
	ret0, _ := ret[0].(explorer.SendSmartContractResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendSmartContract indicates an expected call of SendSmartContract
func (mr *MockExplorerMockRecorder) SendSmartContract(request interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendSmartContract", reflect.TypeOf((*MockExplorer)(nil).SendSmartContract), request)
}

// PutSubChainBlock mocks base method
func (m *MockExplorer) PutSubChainBlock(request explorer.PutSubChainBlockRequest) (explorer.PutSubChainBlockResponse, error) {
	ret := m.ctrl.Call(m, "PutSubChainBlock", request)
	ret0, _ := ret[0].(explorer.PutSubChainBlockResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutSubChainBlock indicates an expected call of PutSubChainBlock
func (mr *MockExplorerMockRecorder) PutSubChainBlock(request interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutSubChainBlock", reflect.TypeOf((*MockExplorer)(nil).PutSubChainBlock), request)
}

// SendAction mocks base method
func (m *MockExplorer) SendAction(request explorer.SendActionRequest) (explorer.SendActionResponse, error) {
	ret := m.ctrl.Call(m, "SendAction", request)
	ret0, _ := ret[0].(explorer.SendActionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendAction indicates an expected call of SendAction
func (mr *MockExplorerMockRecorder) SendAction(request interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendAction", reflect.TypeOf((*MockExplorer)(nil).SendAction), request)
}

// GetPeers mocks base method
func (m *MockExplorer) GetPeers() (explorer.GetPeersResponse, error) {
	ret := m.ctrl.Call(m, "GetPeers")
	ret0, _ := ret[0].(explorer.GetPeersResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPeers indicates an expected call of GetPeers
func (mr *MockExplorerMockRecorder) GetPeers() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPeers", reflect.TypeOf((*MockExplorer)(nil).GetPeers))
}

// GetReceiptByExecutionID mocks base method
func (m *MockExplorer) GetReceiptByExecutionID(id string) (explorer.Receipt, error) {
	ret := m.ctrl.Call(m, "GetReceiptByExecutionID", id)
	ret0, _ := ret[0].(explorer.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReceiptByExecutionID indicates an expected call of GetReceiptByExecutionID
func (mr *MockExplorerMockRecorder) GetReceiptByExecutionID(id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReceiptByExecutionID", reflect.TypeOf((*MockExplorer)(nil).GetReceiptByExecutionID), id)
}

// GetReceiptByActionID mocks base method
func (m *MockExplorer) GetReceiptByActionID(id string) (explorer.Receipt, error) {
	ret := m.ctrl.Call(m, "GetReceiptByActionID", id)
	ret0, _ := ret[0].(explorer.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReceiptByActionID indicates an expected call of GetReceiptByActionID
func (mr *MockExplorerMockRecorder) GetReceiptByActionID(id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReceiptByActionID", reflect.TypeOf((*MockExplorer)(nil).GetReceiptByActionID), id)
}

// ReadExecutionState mocks base method
func (m *MockExplorer) ReadExecutionState(request explorer.Execution) (string, error) {
	ret := m.ctrl.Call(m, "ReadExecutionState", request)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadExecutionState indicates an expected call of ReadExecutionState
func (mr *MockExplorerMockRecorder) ReadExecutionState(request interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadExecutionState", reflect.TypeOf((*MockExplorer)(nil).ReadExecutionState), request)
}

// GetBlockOrActionByHash mocks base method
func (m *MockExplorer) GetBlockOrActionByHash(hashStr string) (explorer.GetBlkOrActResponse, error) {
	ret := m.ctrl.Call(m, "GetBlockOrActionByHash", hashStr)
	ret0, _ := ret[0].(explorer.GetBlkOrActResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlockOrActionByHash indicates an expected call of GetBlockOrActionByHash
func (mr *MockExplorerMockRecorder) GetBlockOrActionByHash(hashStr interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockOrActionByHash", reflect.TypeOf((*MockExplorer)(nil).GetBlockOrActionByHash), hashStr)
}

// CreateDeposit mocks base method
func (m *MockExplorer) CreateDeposit(request explorer.CreateDepositRequest) (explorer.CreateDepositResponse, error) {
	ret := m.ctrl.Call(m, "CreateDeposit", request)
	ret0, _ := ret[0].(explorer.CreateDepositResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateDeposit indicates an expected call of CreateDeposit
func (mr *MockExplorerMockRecorder) CreateDeposit(request interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDeposit", reflect.TypeOf((*MockExplorer)(nil).CreateDeposit), request)
}

// GetDeposits mocks base method
func (m *MockExplorer) GetDeposits(subChainID, offset, limit int64) ([]explorer.Deposit, error) {
	ret := m.ctrl.Call(m, "GetDeposits", subChainID, offset, limit)
	ret0, _ := ret[0].([]explorer.Deposit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeposits indicates an expected call of GetDeposits
func (mr *MockExplorerMockRecorder) GetDeposits(subChainID, offset, limit interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeposits", reflect.TypeOf((*MockExplorer)(nil).GetDeposits), subChainID, offset, limit)
}

// SettleDeposit mocks base method
func (m *MockExplorer) SettleDeposit(request explorer.SettleDepositRequest) (explorer.SettleDepositResponse, error) {
	ret := m.ctrl.Call(m, "SettleDeposit", request)
	ret0, _ := ret[0].(explorer.SettleDepositResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SettleDeposit indicates an expected call of SettleDeposit
func (mr *MockExplorerMockRecorder) SettleDeposit(request interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SettleDeposit", reflect.TypeOf((*MockExplorer)(nil).SettleDeposit), request)
}

// SuggestGasPrice mocks base method
func (m *MockExplorer) SuggestGasPrice() (int64, error) {
	ret := m.ctrl.Call(m, "SuggestGasPrice")
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SuggestGasPrice indicates an expected call of SuggestGasPrice
func (mr *MockExplorerMockRecorder) SuggestGasPrice() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SuggestGasPrice", reflect.TypeOf((*MockExplorer)(nil).SuggestGasPrice))
}

// EstimateGasForTransfer mocks base method
func (m *MockExplorer) EstimateGasForTransfer(request explorer.SendTransferRequest) (int64, error) {
	ret := m.ctrl.Call(m, "EstimateGasForTransfer", request)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EstimateGasForTransfer indicates an expected call of EstimateGasForTransfer
func (mr *MockExplorerMockRecorder) EstimateGasForTransfer(request interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EstimateGasForTransfer", reflect.TypeOf((*MockExplorer)(nil).EstimateGasForTransfer), request)
}

// EstimateGasForVote mocks base method
func (m *MockExplorer) EstimateGasForVote() (int64, error) {
	ret := m.ctrl.Call(m, "EstimateGasForVote")
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EstimateGasForVote indicates an expected call of EstimateGasForVote
func (mr *MockExplorerMockRecorder) EstimateGasForVote() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EstimateGasForVote", reflect.TypeOf((*MockExplorer)(nil).EstimateGasForVote))
}

// EstimateGasForSmartContract mocks base method
func (m *MockExplorer) EstimateGasForSmartContract(request explorer.Execution) (int64, error) {
	ret := m.ctrl.Call(m, "EstimateGasForSmartContract", request)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EstimateGasForSmartContract indicates an expected call of EstimateGasForSmartContract
func (mr *MockExplorerMockRecorder) EstimateGasForSmartContract(request interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EstimateGasForSmartContract", reflect.TypeOf((*MockExplorer)(nil).EstimateGasForSmartContract), request)
}

// GetStateRootHash mocks base method
func (m *MockExplorer) GetStateRootHash(blockHeight int64) (string, error) {
	ret := m.ctrl.Call(m, "GetStateRootHash", blockHeight)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStateRootHash indicates an expected call of GetStateRootHash
func (mr *MockExplorerMockRecorder) GetStateRootHash(blockHeight interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStateRootHash", reflect.TypeOf((*MockExplorer)(nil).GetStateRootHash), blockHeight)
}
