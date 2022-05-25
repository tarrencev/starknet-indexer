// Code generated by entc, DO NOT EDIT.

package token

const (
	// Label holds the string label denoting the token type in the database.
	Label = "token"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTokenId holds the string denoting the tokenid field in the database.
	FieldTokenId = "token_id"
	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// EdgeContract holds the string denoting the contract edge name in mutations.
	EdgeContract = "contract"
	// Table holds the table name of the token in the database.
	Table = "tokens"
	// OwnerTable is the table that holds the owner relation/edge.
	OwnerTable = "tokens"
	// OwnerInverseTable is the table name for the Contract entity.
	// It exists in this package in order to avoid circular dependency with the "contract" package.
	OwnerInverseTable = "contracts"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "token_owner"
	// ContractTable is the table that holds the contract relation/edge.
	ContractTable = "tokens"
	// ContractInverseTable is the table name for the Contract entity.
	// It exists in this package in order to avoid circular dependency with the "contract" package.
	ContractInverseTable = "contracts"
	// ContractColumn is the table column denoting the contract relation/edge.
	ContractColumn = "token_contract"
)

// Columns holds all SQL columns for token fields.
var Columns = []string{
	FieldID,
	FieldTokenId,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "tokens"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"token_owner",
	"token_contract",
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