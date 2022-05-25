// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/cartridge-gg/starknet-indexer/ent/contract"
	"github.com/cartridge-gg/starknet-indexer/ent/schema/big"
	"github.com/cartridge-gg/starknet-indexer/ent/token"
)

// Token is the model entity for the Token schema.
type Token struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// TokenId holds the value of the "tokenId" field.
	TokenId big.Int `json:"tokenId,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TokenQuery when eager-loading is set.
	Edges          TokenEdges `json:"edges"`
	token_owner    *string
	token_contract *string
}

// TokenEdges holds the relations/edges for other nodes in the graph.
type TokenEdges struct {
	// Owner holds the value of the owner edge.
	Owner *Contract `json:"owner,omitempty"`
	// Contract holds the value of the contract edge.
	Contract *Contract `json:"contract,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]*int
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TokenEdges) OwnerOrErr() (*Contract, error) {
	if e.loadedTypes[0] {
		if e.Owner == nil {
			// The edge owner was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: contract.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// ContractOrErr returns the Contract value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TokenEdges) ContractOrErr() (*Contract, error) {
	if e.loadedTypes[1] {
		if e.Contract == nil {
			// The edge contract was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: contract.Label}
		}
		return e.Contract, nil
	}
	return nil, &NotLoadedError{edge: "contract"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Token) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case token.FieldTokenId:
			values[i] = new(big.Int)
		case token.FieldID:
			values[i] = new(sql.NullString)
		case token.ForeignKeys[0]: // token_owner
			values[i] = new(sql.NullString)
		case token.ForeignKeys[1]: // token_contract
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Token", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Token fields.
func (t *Token) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case token.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				t.ID = value.String
			}
		case token.FieldTokenId:
			if value, ok := values[i].(*big.Int); !ok {
				return fmt.Errorf("unexpected type %T for field tokenId", values[i])
			} else if value != nil {
				t.TokenId = *value
			}
		case token.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field token_owner", values[i])
			} else if value.Valid {
				t.token_owner = new(string)
				*t.token_owner = value.String
			}
		case token.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field token_contract", values[i])
			} else if value.Valid {
				t.token_contract = new(string)
				*t.token_contract = value.String
			}
		}
	}
	return nil
}

// QueryOwner queries the "owner" edge of the Token entity.
func (t *Token) QueryOwner() *ContractQuery {
	return (&TokenClient{config: t.config}).QueryOwner(t)
}

// QueryContract queries the "contract" edge of the Token entity.
func (t *Token) QueryContract() *ContractQuery {
	return (&TokenClient{config: t.config}).QueryContract(t)
}

// Update returns a builder for updating this Token.
// Note that you need to call Token.Unwrap() before calling this method if this Token
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Token) Update() *TokenUpdateOne {
	return (&TokenClient{config: t.config}).UpdateOne(t)
}

// Unwrap unwraps the Token entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Token) Unwrap() *Token {
	tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Token is not a transactional entity")
	}
	t.config.driver = tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Token) String() string {
	var builder strings.Builder
	builder.WriteString("Token(")
	builder.WriteString(fmt.Sprintf("id=%v", t.ID))
	builder.WriteString(", tokenId=")
	builder.WriteString(fmt.Sprintf("%v", t.TokenId))
	builder.WriteByte(')')
	return builder.String()
}

// Tokens is a parsable slice of Token.
type Tokens []*Token

func (t Tokens) config(cfg config) {
	for _i := range t {
		t[_i].config = cfg
	}
}
