// Code generated by entc, DO NOT EDIT.

package balance

import (
	"github.com/cartridge-gg/starknet-indexer/ent/schema/big"
)

const (
	// Label holds the string label denoting the balance type in the database.
	Label = "balance"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldBalance holds the string denoting the balance field in the database.
	FieldBalance = "balance"
	// EdgeAccount holds the string denoting the account edge name in mutations.
	EdgeAccount = "account"
	// EdgeContract holds the string denoting the contract edge name in mutations.
	EdgeContract = "contract"
	// Table holds the table name of the balance in the database.
	Table = "balances"
	// AccountTable is the table that holds the account relation/edge.
	AccountTable = "balances"
	// AccountInverseTable is the table name for the Contract entity.
	// It exists in this package in order to avoid circular dependency with the "contract" package.
	AccountInverseTable = "contracts"
	// AccountColumn is the table column denoting the account relation/edge.
	AccountColumn = "balance_account"
	// ContractTable is the table that holds the contract relation/edge.
	ContractTable = "balances"
	// ContractInverseTable is the table name for the Contract entity.
	// It exists in this package in order to avoid circular dependency with the "contract" package.
	ContractInverseTable = "contracts"
	// ContractColumn is the table column denoting the contract relation/edge.
	ContractColumn = "balance_contract"
)

// Columns holds all SQL columns for balance fields.
var Columns = []string{
	FieldID,
	FieldBalance,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "balances"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"balance_account",
	"balance_contract",
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

var (
	// DefaultBalance holds the default value on creation for the "balance" field.
	DefaultBalance func() big.Int
)
