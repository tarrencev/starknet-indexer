// Code generated by entc, DO NOT EDIT.

package transactionreceipt

import (
	"fmt"
	"io"
	"strconv"
)

const (
	// Label holds the string label denoting the transactionreceipt type in the database.
	Label = "transaction_receipt"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTransactionHash holds the string denoting the transaction_hash field in the database.
	FieldTransactionHash = "transaction_hash"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldStatusData holds the string denoting the status_data field in the database.
	FieldStatusData = "status_data"
	// FieldMessagesSent holds the string denoting the messages_sent field in the database.
	FieldMessagesSent = "messages_sent"
	// FieldL1OriginMessage holds the string denoting the l1_origin_message field in the database.
	FieldL1OriginMessage = "l1_origin_message"
	// FieldEvents holds the string denoting the events field in the database.
	FieldEvents = "events"
	// EdgeBlock holds the string denoting the block edge name in mutations.
	EdgeBlock = "block"
	// EdgeTransaction holds the string denoting the transaction edge name in mutations.
	EdgeTransaction = "transaction"
	// Table holds the table name of the transactionreceipt in the database.
	Table = "transaction_receipts"
	// BlockTable is the table that holds the block relation/edge.
	BlockTable = "transaction_receipts"
	// BlockInverseTable is the table name for the Block entity.
	// It exists in this package in order to avoid circular dependency with the "block" package.
	BlockInverseTable = "blocks"
	// BlockColumn is the table column denoting the block relation/edge.
	BlockColumn = "block_transaction_receipts"
	// TransactionTable is the table that holds the transaction relation/edge.
	TransactionTable = "transaction_receipts"
	// TransactionInverseTable is the table name for the Transaction entity.
	// It exists in this package in order to avoid circular dependency with the "transaction" package.
	TransactionInverseTable = "transactions"
	// TransactionColumn is the table column denoting the transaction relation/edge.
	TransactionColumn = "transaction_receipts"
)

// Columns holds all SQL columns for transactionreceipt fields.
var Columns = []string{
	FieldID,
	FieldTransactionHash,
	FieldStatus,
	FieldStatusData,
	FieldMessagesSent,
	FieldL1OriginMessage,
	FieldEvents,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "transaction_receipts"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"block_transaction_receipts",
	"transaction_receipts",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

// Status defines the type for the "status" enum field.
type Status string

// Status values.
const (
	StatusUNKNOWN        Status = "UNKNOWN"
	StatusRECEIVED       Status = "RECEIVED"
	StatusPENDING        Status = "PENDING"
	StatusACCEPTED_ON_L2 Status = "ACCEPTED_ON_L2"
	StatusACCEPTED_ON_L1 Status = "ACCEPTED_ON_L1"
	StatusREJECTED       Status = "REJECTED"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusUNKNOWN, StatusRECEIVED, StatusPENDING, StatusACCEPTED_ON_L2, StatusACCEPTED_ON_L1, StatusREJECTED:
		return nil
	default:
		return fmt.Errorf("transactionreceipt: invalid enum value for status field: %q", s)
	}
}

// MarshalGQL implements graphql.Marshaler interface.
func (e Status) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(e.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (e *Status) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*e = Status(str)
	if err := StatusValidator(*e); err != nil {
		return fmt.Errorf("%s is not a valid Status", str)
	}
	return nil
}