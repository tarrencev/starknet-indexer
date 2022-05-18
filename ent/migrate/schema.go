// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// BlocksColumns holds the columns for the "blocks" table.
	BlocksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "block_hash", Type: field.TypeString, Unique: true},
		{Name: "parent_block_hash", Type: field.TypeString},
		{Name: "block_number", Type: field.TypeUint64, Unique: true},
		{Name: "state_root", Type: field.TypeString},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"ACCEPTED_ON_L1", "ACCEPTED_ON_L2"}},
		{Name: "timestamp", Type: field.TypeTime},
	}
	// BlocksTable holds the schema information for the "blocks" table.
	BlocksTable = &schema.Table{
		Name:       "blocks",
		Columns:    BlocksColumns,
		PrimaryKey: []*schema.Column{BlocksColumns[0]},
	}
	// ContractsColumns holds the columns for the "contracts" table.
	ContractsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"UNKNOWN", "ERC20", "ERC721"}, Default: "UNKNOWN"},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// ContractsTable holds the schema information for the "contracts" table.
	ContractsTable = &schema.Table{
		Name:       "contracts",
		Columns:    ContractsColumns,
		PrimaryKey: []*schema.Column{ContractsColumns[0]},
	}
	// EventsColumns holds the columns for the "events" table.
	EventsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "from", Type: field.TypeString},
		{Name: "keys", Type: field.TypeJSON},
		{Name: "data", Type: field.TypeJSON},
		{Name: "transaction_events", Type: field.TypeString, Nullable: true},
	}
	// EventsTable holds the schema information for the "events" table.
	EventsTable = &schema.Table{
		Name:       "events",
		Columns:    EventsColumns,
		PrimaryKey: []*schema.Column{EventsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "events_transactions_events",
				Columns:    []*schema.Column{EventsColumns[4]},
				RefColumns: []*schema.Column{TransactionsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// TransactionsColumns holds the columns for the "transactions" table.
	TransactionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "contract_address", Type: field.TypeString},
		{Name: "entry_point_selector", Type: field.TypeString, Nullable: true},
		{Name: "transaction_hash", Type: field.TypeString},
		{Name: "calldata", Type: field.TypeJSON},
		{Name: "signature", Type: field.TypeJSON, Nullable: true},
		{Name: "nonce", Type: field.TypeString, Nullable: true},
		{Name: "block_transactions", Type: field.TypeString, Nullable: true},
	}
	// TransactionsTable holds the schema information for the "transactions" table.
	TransactionsTable = &schema.Table{
		Name:       "transactions",
		Columns:    TransactionsColumns,
		PrimaryKey: []*schema.Column{TransactionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "transactions_blocks_transactions",
				Columns:    []*schema.Column{TransactionsColumns[7]},
				RefColumns: []*schema.Column{BlocksColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// TransactionReceiptsColumns holds the columns for the "transaction_receipts" table.
	TransactionReceiptsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "transaction_hash", Type: field.TypeString},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"UNKNOWN", "RECEIVED", "PENDING", "ACCEPTED_ON_L2", "ACCEPTED_ON_L1", "REJECTED"}},
		{Name: "status_data", Type: field.TypeString},
		{Name: "messages_sent", Type: field.TypeJSON},
		{Name: "l1_origin_message", Type: field.TypeJSON},
		{Name: "block_transaction_receipts", Type: field.TypeString, Nullable: true},
		{Name: "transaction_receipt", Type: field.TypeString, Unique: true},
	}
	// TransactionReceiptsTable holds the schema information for the "transaction_receipts" table.
	TransactionReceiptsTable = &schema.Table{
		Name:       "transaction_receipts",
		Columns:    TransactionReceiptsColumns,
		PrimaryKey: []*schema.Column{TransactionReceiptsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "transaction_receipts_blocks_transaction_receipts",
				Columns:    []*schema.Column{TransactionReceiptsColumns[6]},
				RefColumns: []*schema.Column{BlocksColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "transaction_receipts_transactions_receipt",
				Columns:    []*schema.Column{TransactionReceiptsColumns[7]},
				RefColumns: []*schema.Column{TransactionsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// ContractTransactionsColumns holds the columns for the "contract_transactions" table.
	ContractTransactionsColumns = []*schema.Column{
		{Name: "contract_id", Type: field.TypeString},
		{Name: "transaction_id", Type: field.TypeString},
	}
	// ContractTransactionsTable holds the schema information for the "contract_transactions" table.
	ContractTransactionsTable = &schema.Table{
		Name:       "contract_transactions",
		Columns:    ContractTransactionsColumns,
		PrimaryKey: []*schema.Column{ContractTransactionsColumns[0], ContractTransactionsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "contract_transactions_contract_id",
				Columns:    []*schema.Column{ContractTransactionsColumns[0]},
				RefColumns: []*schema.Column{ContractsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "contract_transactions_transaction_id",
				Columns:    []*schema.Column{ContractTransactionsColumns[1]},
				RefColumns: []*schema.Column{TransactionsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		BlocksTable,
		ContractsTable,
		EventsTable,
		TransactionsTable,
		TransactionReceiptsTable,
		ContractTransactionsTable,
	}
)

func init() {
	EventsTable.ForeignKeys[0].RefTable = TransactionsTable
	TransactionsTable.ForeignKeys[0].RefTable = BlocksTable
	TransactionReceiptsTable.ForeignKeys[0].RefTable = BlocksTable
	TransactionReceiptsTable.ForeignKeys[1].RefTable = TransactionsTable
	ContractTransactionsTable.ForeignKeys[0].RefTable = ContractsTable
	ContractTransactionsTable.ForeignKeys[1].RefTable = TransactionsTable
}
