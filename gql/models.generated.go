// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package gql

import (
	"time"

	"github.com/tarrencev/starknet-indexer/ent/block"
	"github.com/tarrencev/starknet-indexer/ent/transaction"
	"github.com/tarrencev/starknet-indexer/ent/transactionreceipt"
)

// BlockWhereInput is used for filtering Block objects.
// Input was generated by ent.
type BlockWhereInput struct {
	Not *BlockWhereInput   `json:"not"`
	And []*BlockWhereInput `json:"and"`
	Or  []*BlockWhereInput `json:"or"`
	// block_hash field predicates
	BlockHash             *string  `json:"blockHash"`
	BlockHashNeq          *string  `json:"blockHashNEQ"`
	BlockHashIn           []string `json:"blockHashIn"`
	BlockHashNotIn        []string `json:"blockHashNotIn"`
	BlockHashGt           *string  `json:"blockHashGT"`
	BlockHashGte          *string  `json:"blockHashGTE"`
	BlockHashLt           *string  `json:"blockHashLT"`
	BlockHashLte          *string  `json:"blockHashLTE"`
	BlockHashContains     *string  `json:"blockHashContains"`
	BlockHashHasPrefix    *string  `json:"blockHashHasPrefix"`
	BlockHashHasSuffix    *string  `json:"blockHashHasSuffix"`
	BlockHashEqualFold    *string  `json:"blockHashEqualFold"`
	BlockHashContainsFold *string  `json:"blockHashContainsFold"`
	// parent_block_hash field predicates
	ParentBlockHash             *string  `json:"parentBlockHash"`
	ParentBlockHashNeq          *string  `json:"parentBlockHashNEQ"`
	ParentBlockHashIn           []string `json:"parentBlockHashIn"`
	ParentBlockHashNotIn        []string `json:"parentBlockHashNotIn"`
	ParentBlockHashGt           *string  `json:"parentBlockHashGT"`
	ParentBlockHashGte          *string  `json:"parentBlockHashGTE"`
	ParentBlockHashLt           *string  `json:"parentBlockHashLT"`
	ParentBlockHashLte          *string  `json:"parentBlockHashLTE"`
	ParentBlockHashContains     *string  `json:"parentBlockHashContains"`
	ParentBlockHashHasPrefix    *string  `json:"parentBlockHashHasPrefix"`
	ParentBlockHashHasSuffix    *string  `json:"parentBlockHashHasSuffix"`
	ParentBlockHashEqualFold    *string  `json:"parentBlockHashEqualFold"`
	ParentBlockHashContainsFold *string  `json:"parentBlockHashContainsFold"`
	// block_number field predicates
	BlockNumber      *uint64  `json:"blockNumber"`
	BlockNumberNeq   *uint64  `json:"blockNumberNEQ"`
	BlockNumberIn    []uint64 `json:"blockNumberIn"`
	BlockNumberNotIn []uint64 `json:"blockNumberNotIn"`
	BlockNumberGt    *uint64  `json:"blockNumberGT"`
	BlockNumberGte   *uint64  `json:"blockNumberGTE"`
	BlockNumberLt    *uint64  `json:"blockNumberLT"`
	BlockNumberLte   *uint64  `json:"blockNumberLTE"`
	// state_root field predicates
	StateRoot             *string  `json:"stateRoot"`
	StateRootNeq          *string  `json:"stateRootNEQ"`
	StateRootIn           []string `json:"stateRootIn"`
	StateRootNotIn        []string `json:"stateRootNotIn"`
	StateRootGt           *string  `json:"stateRootGT"`
	StateRootGte          *string  `json:"stateRootGTE"`
	StateRootLt           *string  `json:"stateRootLT"`
	StateRootLte          *string  `json:"stateRootLTE"`
	StateRootContains     *string  `json:"stateRootContains"`
	StateRootHasPrefix    *string  `json:"stateRootHasPrefix"`
	StateRootHasSuffix    *string  `json:"stateRootHasSuffix"`
	StateRootEqualFold    *string  `json:"stateRootEqualFold"`
	StateRootContainsFold *string  `json:"stateRootContainsFold"`
	// status field predicates
	Status      *block.Status  `json:"status"`
	StatusNeq   *block.Status  `json:"statusNEQ"`
	StatusIn    []block.Status `json:"statusIn"`
	StatusNotIn []block.Status `json:"statusNotIn"`
	// timestamp field predicates
	Timestamp      *time.Time   `json:"timestamp"`
	TimestampNeq   *time.Time   `json:"timestampNEQ"`
	TimestampIn    []*time.Time `json:"timestampIn"`
	TimestampNotIn []*time.Time `json:"timestampNotIn"`
	TimestampGt    *time.Time   `json:"timestampGT"`
	TimestampGte   *time.Time   `json:"timestampGTE"`
	TimestampLt    *time.Time   `json:"timestampLT"`
	TimestampLte   *time.Time   `json:"timestampLTE"`
	// id field predicates
	ID      *string  `json:"id"`
	IDNeq   *string  `json:"idNEQ"`
	IDIn    []string `json:"idIn"`
	IDNotIn []string `json:"idNotIn"`
	IDGt    *string  `json:"idGT"`
	IDGte   *string  `json:"idGTE"`
	IDLt    *string  `json:"idLT"`
	IDLte   *string  `json:"idLTE"`
	// transactions edge predicates
	HasTransactions     *bool                    `json:"hasTransactions"`
	HasTransactionsWith []*TransactionWhereInput `json:"hasTransactionsWith"`
	// transaction_receipts edge predicates
	HasTransactionReceipts     *bool                           `json:"hasTransactionReceipts"`
	HasTransactionReceiptsWith []*TransactionReceiptWhereInput `json:"hasTransactionReceiptsWith"`
}

// TransactionReceiptWhereInput is used for filtering TransactionReceipt objects.
// Input was generated by ent.
type TransactionReceiptWhereInput struct {
	Not *TransactionReceiptWhereInput   `json:"not"`
	And []*TransactionReceiptWhereInput `json:"and"`
	Or  []*TransactionReceiptWhereInput `json:"or"`
	// transaction_hash field predicates
	TransactionHash             *string  `json:"transactionHash"`
	TransactionHashNeq          *string  `json:"transactionHashNEQ"`
	TransactionHashIn           []string `json:"transactionHashIn"`
	TransactionHashNotIn        []string `json:"transactionHashNotIn"`
	TransactionHashGt           *string  `json:"transactionHashGT"`
	TransactionHashGte          *string  `json:"transactionHashGTE"`
	TransactionHashLt           *string  `json:"transactionHashLT"`
	TransactionHashLte          *string  `json:"transactionHashLTE"`
	TransactionHashContains     *string  `json:"transactionHashContains"`
	TransactionHashHasPrefix    *string  `json:"transactionHashHasPrefix"`
	TransactionHashHasSuffix    *string  `json:"transactionHashHasSuffix"`
	TransactionHashEqualFold    *string  `json:"transactionHashEqualFold"`
	TransactionHashContainsFold *string  `json:"transactionHashContainsFold"`
	// status field predicates
	Status      *transactionreceipt.Status  `json:"status"`
	StatusNeq   *transactionreceipt.Status  `json:"statusNEQ"`
	StatusIn    []transactionreceipt.Status `json:"statusIn"`
	StatusNotIn []transactionreceipt.Status `json:"statusNotIn"`
	// status_data field predicates
	StatusData             *string  `json:"statusData"`
	StatusDataNeq          *string  `json:"statusDataNEQ"`
	StatusDataIn           []string `json:"statusDataIn"`
	StatusDataNotIn        []string `json:"statusDataNotIn"`
	StatusDataGt           *string  `json:"statusDataGT"`
	StatusDataGte          *string  `json:"statusDataGTE"`
	StatusDataLt           *string  `json:"statusDataLT"`
	StatusDataLte          *string  `json:"statusDataLTE"`
	StatusDataContains     *string  `json:"statusDataContains"`
	StatusDataHasPrefix    *string  `json:"statusDataHasPrefix"`
	StatusDataHasSuffix    *string  `json:"statusDataHasSuffix"`
	StatusDataEqualFold    *string  `json:"statusDataEqualFold"`
	StatusDataContainsFold *string  `json:"statusDataContainsFold"`
	// id field predicates
	ID      *string  `json:"id"`
	IDNeq   *string  `json:"idNEQ"`
	IDIn    []string `json:"idIn"`
	IDNotIn []string `json:"idNotIn"`
	IDGt    *string  `json:"idGT"`
	IDGte   *string  `json:"idGTE"`
	IDLt    *string  `json:"idLT"`
	IDLte   *string  `json:"idLTE"`
	// block edge predicates
	HasBlock     *bool              `json:"hasBlock"`
	HasBlockWith []*BlockWhereInput `json:"hasBlockWith"`
}

// TransactionWhereInput is used for filtering Transaction objects.
// Input was generated by ent.
type TransactionWhereInput struct {
	Not *TransactionWhereInput   `json:"not"`
	And []*TransactionWhereInput `json:"and"`
	Or  []*TransactionWhereInput `json:"or"`
	// contract_address field predicates
	ContractAddress             *string  `json:"contractAddress"`
	ContractAddressNeq          *string  `json:"contractAddressNEQ"`
	ContractAddressIn           []string `json:"contractAddressIn"`
	ContractAddressNotIn        []string `json:"contractAddressNotIn"`
	ContractAddressGt           *string  `json:"contractAddressGT"`
	ContractAddressGte          *string  `json:"contractAddressGTE"`
	ContractAddressLt           *string  `json:"contractAddressLT"`
	ContractAddressLte          *string  `json:"contractAddressLTE"`
	ContractAddressContains     *string  `json:"contractAddressContains"`
	ContractAddressHasPrefix    *string  `json:"contractAddressHasPrefix"`
	ContractAddressHasSuffix    *string  `json:"contractAddressHasSuffix"`
	ContractAddressEqualFold    *string  `json:"contractAddressEqualFold"`
	ContractAddressContainsFold *string  `json:"contractAddressContainsFold"`
	// entry_point_selector field predicates
	EntryPointSelector             *string  `json:"entryPointSelector"`
	EntryPointSelectorNeq          *string  `json:"entryPointSelectorNEQ"`
	EntryPointSelectorIn           []string `json:"entryPointSelectorIn"`
	EntryPointSelectorNotIn        []string `json:"entryPointSelectorNotIn"`
	EntryPointSelectorGt           *string  `json:"entryPointSelectorGT"`
	EntryPointSelectorGte          *string  `json:"entryPointSelectorGTE"`
	EntryPointSelectorLt           *string  `json:"entryPointSelectorLT"`
	EntryPointSelectorLte          *string  `json:"entryPointSelectorLTE"`
	EntryPointSelectorContains     *string  `json:"entryPointSelectorContains"`
	EntryPointSelectorHasPrefix    *string  `json:"entryPointSelectorHasPrefix"`
	EntryPointSelectorHasSuffix    *string  `json:"entryPointSelectorHasSuffix"`
	EntryPointSelectorEqualFold    *string  `json:"entryPointSelectorEqualFold"`
	EntryPointSelectorContainsFold *string  `json:"entryPointSelectorContainsFold"`
	// transaction_hash field predicates
	TransactionHash             *string  `json:"transactionHash"`
	TransactionHashNeq          *string  `json:"transactionHashNEQ"`
	TransactionHashIn           []string `json:"transactionHashIn"`
	TransactionHashNotIn        []string `json:"transactionHashNotIn"`
	TransactionHashGt           *string  `json:"transactionHashGT"`
	TransactionHashGte          *string  `json:"transactionHashGTE"`
	TransactionHashLt           *string  `json:"transactionHashLT"`
	TransactionHashLte          *string  `json:"transactionHashLTE"`
	TransactionHashContains     *string  `json:"transactionHashContains"`
	TransactionHashHasPrefix    *string  `json:"transactionHashHasPrefix"`
	TransactionHashHasSuffix    *string  `json:"transactionHashHasSuffix"`
	TransactionHashEqualFold    *string  `json:"transactionHashEqualFold"`
	TransactionHashContainsFold *string  `json:"transactionHashContainsFold"`
	// type field predicates
	Type      *transaction.Type  `json:"type"`
	TypeNeq   *transaction.Type  `json:"typeNEQ"`
	TypeIn    []transaction.Type `json:"typeIn"`
	TypeNotIn []transaction.Type `json:"typeNotIn"`
	// nonce field predicates
	Nonce             *string  `json:"nonce"`
	NonceNeq          *string  `json:"nonceNEQ"`
	NonceIn           []string `json:"nonceIn"`
	NonceNotIn        []string `json:"nonceNotIn"`
	NonceGt           *string  `json:"nonceGT"`
	NonceGte          *string  `json:"nonceGTE"`
	NonceLt           *string  `json:"nonceLT"`
	NonceLte          *string  `json:"nonceLTE"`
	NonceContains     *string  `json:"nonceContains"`
	NonceHasPrefix    *string  `json:"nonceHasPrefix"`
	NonceHasSuffix    *string  `json:"nonceHasSuffix"`
	NonceEqualFold    *string  `json:"nonceEqualFold"`
	NonceContainsFold *string  `json:"nonceContainsFold"`
	// id field predicates
	ID      *string  `json:"id"`
	IDNeq   *string  `json:"idNEQ"`
	IDIn    []string `json:"idIn"`
	IDNotIn []string `json:"idNotIn"`
	IDGt    *string  `json:"idGT"`
	IDGte   *string  `json:"idGTE"`
	IDLt    *string  `json:"idLT"`
	IDLte   *string  `json:"idLTE"`
	// block edge predicates
	HasBlock     *bool              `json:"hasBlock"`
	HasBlockWith []*BlockWhereInput `json:"hasBlockWith"`
}
