// Code generated by entc, DO NOT EDIT.

package transaction

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/tarrencev/starknet-indexer/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// ContractAddress applies equality check predicate on the "contract_address" field. It's identical to ContractAddressEQ.
func ContractAddress(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContractAddress), v))
	})
}

// EntryPointSelector applies equality check predicate on the "entry_point_selector" field. It's identical to EntryPointSelectorEQ.
func EntryPointSelector(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEntryPointSelector), v))
	})
}

// EntryPointType applies equality check predicate on the "entry_point_type" field. It's identical to EntryPointTypeEQ.
func EntryPointType(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEntryPointType), v))
	})
}

// TransactionHash applies equality check predicate on the "transaction_hash" field. It's identical to TransactionHashEQ.
func TransactionHash(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTransactionHash), v))
	})
}

// Nonce applies equality check predicate on the "nonce" field. It's identical to NonceEQ.
func Nonce(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNonce), v))
	})
}

// ContractAddressEQ applies the EQ predicate on the "contract_address" field.
func ContractAddressEQ(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContractAddress), v))
	})
}

// ContractAddressNEQ applies the NEQ predicate on the "contract_address" field.
func ContractAddressNEQ(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldContractAddress), v))
	})
}

// ContractAddressIn applies the In predicate on the "contract_address" field.
func ContractAddressIn(vs ...string) predicate.Transaction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transaction(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldContractAddress), v...))
	})
}

// ContractAddressNotIn applies the NotIn predicate on the "contract_address" field.
func ContractAddressNotIn(vs ...string) predicate.Transaction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transaction(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldContractAddress), v...))
	})
}

// ContractAddressGT applies the GT predicate on the "contract_address" field.
func ContractAddressGT(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldContractAddress), v))
	})
}

// ContractAddressGTE applies the GTE predicate on the "contract_address" field.
func ContractAddressGTE(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldContractAddress), v))
	})
}

// ContractAddressLT applies the LT predicate on the "contract_address" field.
func ContractAddressLT(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldContractAddress), v))
	})
}

// ContractAddressLTE applies the LTE predicate on the "contract_address" field.
func ContractAddressLTE(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldContractAddress), v))
	})
}

// ContractAddressContains applies the Contains predicate on the "contract_address" field.
func ContractAddressContains(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldContractAddress), v))
	})
}

// ContractAddressHasPrefix applies the HasPrefix predicate on the "contract_address" field.
func ContractAddressHasPrefix(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldContractAddress), v))
	})
}

// ContractAddressHasSuffix applies the HasSuffix predicate on the "contract_address" field.
func ContractAddressHasSuffix(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldContractAddress), v))
	})
}

// ContractAddressEqualFold applies the EqualFold predicate on the "contract_address" field.
func ContractAddressEqualFold(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldContractAddress), v))
	})
}

// ContractAddressContainsFold applies the ContainsFold predicate on the "contract_address" field.
func ContractAddressContainsFold(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldContractAddress), v))
	})
}

// EntryPointSelectorEQ applies the EQ predicate on the "entry_point_selector" field.
func EntryPointSelectorEQ(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEntryPointSelector), v))
	})
}

// EntryPointSelectorNEQ applies the NEQ predicate on the "entry_point_selector" field.
func EntryPointSelectorNEQ(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEntryPointSelector), v))
	})
}

// EntryPointSelectorIn applies the In predicate on the "entry_point_selector" field.
func EntryPointSelectorIn(vs ...string) predicate.Transaction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transaction(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEntryPointSelector), v...))
	})
}

// EntryPointSelectorNotIn applies the NotIn predicate on the "entry_point_selector" field.
func EntryPointSelectorNotIn(vs ...string) predicate.Transaction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transaction(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEntryPointSelector), v...))
	})
}

// EntryPointSelectorGT applies the GT predicate on the "entry_point_selector" field.
func EntryPointSelectorGT(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEntryPointSelector), v))
	})
}

// EntryPointSelectorGTE applies the GTE predicate on the "entry_point_selector" field.
func EntryPointSelectorGTE(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEntryPointSelector), v))
	})
}

// EntryPointSelectorLT applies the LT predicate on the "entry_point_selector" field.
func EntryPointSelectorLT(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEntryPointSelector), v))
	})
}

// EntryPointSelectorLTE applies the LTE predicate on the "entry_point_selector" field.
func EntryPointSelectorLTE(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEntryPointSelector), v))
	})
}

// EntryPointSelectorContains applies the Contains predicate on the "entry_point_selector" field.
func EntryPointSelectorContains(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldEntryPointSelector), v))
	})
}

// EntryPointSelectorHasPrefix applies the HasPrefix predicate on the "entry_point_selector" field.
func EntryPointSelectorHasPrefix(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldEntryPointSelector), v))
	})
}

// EntryPointSelectorHasSuffix applies the HasSuffix predicate on the "entry_point_selector" field.
func EntryPointSelectorHasSuffix(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldEntryPointSelector), v))
	})
}

// EntryPointSelectorEqualFold applies the EqualFold predicate on the "entry_point_selector" field.
func EntryPointSelectorEqualFold(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldEntryPointSelector), v))
	})
}

// EntryPointSelectorContainsFold applies the ContainsFold predicate on the "entry_point_selector" field.
func EntryPointSelectorContainsFold(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldEntryPointSelector), v))
	})
}

// EntryPointTypeEQ applies the EQ predicate on the "entry_point_type" field.
func EntryPointTypeEQ(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEntryPointType), v))
	})
}

// EntryPointTypeNEQ applies the NEQ predicate on the "entry_point_type" field.
func EntryPointTypeNEQ(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEntryPointType), v))
	})
}

// EntryPointTypeIn applies the In predicate on the "entry_point_type" field.
func EntryPointTypeIn(vs ...string) predicate.Transaction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transaction(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEntryPointType), v...))
	})
}

// EntryPointTypeNotIn applies the NotIn predicate on the "entry_point_type" field.
func EntryPointTypeNotIn(vs ...string) predicate.Transaction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transaction(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEntryPointType), v...))
	})
}

// EntryPointTypeGT applies the GT predicate on the "entry_point_type" field.
func EntryPointTypeGT(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEntryPointType), v))
	})
}

// EntryPointTypeGTE applies the GTE predicate on the "entry_point_type" field.
func EntryPointTypeGTE(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEntryPointType), v))
	})
}

// EntryPointTypeLT applies the LT predicate on the "entry_point_type" field.
func EntryPointTypeLT(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEntryPointType), v))
	})
}

// EntryPointTypeLTE applies the LTE predicate on the "entry_point_type" field.
func EntryPointTypeLTE(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEntryPointType), v))
	})
}

// EntryPointTypeContains applies the Contains predicate on the "entry_point_type" field.
func EntryPointTypeContains(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldEntryPointType), v))
	})
}

// EntryPointTypeHasPrefix applies the HasPrefix predicate on the "entry_point_type" field.
func EntryPointTypeHasPrefix(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldEntryPointType), v))
	})
}

// EntryPointTypeHasSuffix applies the HasSuffix predicate on the "entry_point_type" field.
func EntryPointTypeHasSuffix(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldEntryPointType), v))
	})
}

// EntryPointTypeEqualFold applies the EqualFold predicate on the "entry_point_type" field.
func EntryPointTypeEqualFold(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldEntryPointType), v))
	})
}

// EntryPointTypeContainsFold applies the ContainsFold predicate on the "entry_point_type" field.
func EntryPointTypeContainsFold(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldEntryPointType), v))
	})
}

// TransactionHashEQ applies the EQ predicate on the "transaction_hash" field.
func TransactionHashEQ(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTransactionHash), v))
	})
}

// TransactionHashNEQ applies the NEQ predicate on the "transaction_hash" field.
func TransactionHashNEQ(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTransactionHash), v))
	})
}

// TransactionHashIn applies the In predicate on the "transaction_hash" field.
func TransactionHashIn(vs ...string) predicate.Transaction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transaction(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTransactionHash), v...))
	})
}

// TransactionHashNotIn applies the NotIn predicate on the "transaction_hash" field.
func TransactionHashNotIn(vs ...string) predicate.Transaction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transaction(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTransactionHash), v...))
	})
}

// TransactionHashGT applies the GT predicate on the "transaction_hash" field.
func TransactionHashGT(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTransactionHash), v))
	})
}

// TransactionHashGTE applies the GTE predicate on the "transaction_hash" field.
func TransactionHashGTE(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTransactionHash), v))
	})
}

// TransactionHashLT applies the LT predicate on the "transaction_hash" field.
func TransactionHashLT(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTransactionHash), v))
	})
}

// TransactionHashLTE applies the LTE predicate on the "transaction_hash" field.
func TransactionHashLTE(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTransactionHash), v))
	})
}

// TransactionHashContains applies the Contains predicate on the "transaction_hash" field.
func TransactionHashContains(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTransactionHash), v))
	})
}

// TransactionHashHasPrefix applies the HasPrefix predicate on the "transaction_hash" field.
func TransactionHashHasPrefix(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTransactionHash), v))
	})
}

// TransactionHashHasSuffix applies the HasSuffix predicate on the "transaction_hash" field.
func TransactionHashHasSuffix(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTransactionHash), v))
	})
}

// TransactionHashEqualFold applies the EqualFold predicate on the "transaction_hash" field.
func TransactionHashEqualFold(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTransactionHash), v))
	})
}

// TransactionHashContainsFold applies the ContainsFold predicate on the "transaction_hash" field.
func TransactionHashContainsFold(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTransactionHash), v))
	})
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v Type) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldType), v))
	})
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v Type) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldType), v))
	})
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...Type) predicate.Transaction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transaction(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldType), v...))
	})
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...Type) predicate.Transaction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transaction(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldType), v...))
	})
}

// NonceEQ applies the EQ predicate on the "nonce" field.
func NonceEQ(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNonce), v))
	})
}

// NonceNEQ applies the NEQ predicate on the "nonce" field.
func NonceNEQ(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldNonce), v))
	})
}

// NonceIn applies the In predicate on the "nonce" field.
func NonceIn(vs ...string) predicate.Transaction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transaction(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldNonce), v...))
	})
}

// NonceNotIn applies the NotIn predicate on the "nonce" field.
func NonceNotIn(vs ...string) predicate.Transaction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transaction(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldNonce), v...))
	})
}

// NonceGT applies the GT predicate on the "nonce" field.
func NonceGT(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldNonce), v))
	})
}

// NonceGTE applies the GTE predicate on the "nonce" field.
func NonceGTE(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldNonce), v))
	})
}

// NonceLT applies the LT predicate on the "nonce" field.
func NonceLT(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldNonce), v))
	})
}

// NonceLTE applies the LTE predicate on the "nonce" field.
func NonceLTE(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldNonce), v))
	})
}

// NonceContains applies the Contains predicate on the "nonce" field.
func NonceContains(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldNonce), v))
	})
}

// NonceHasPrefix applies the HasPrefix predicate on the "nonce" field.
func NonceHasPrefix(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldNonce), v))
	})
}

// NonceHasSuffix applies the HasSuffix predicate on the "nonce" field.
func NonceHasSuffix(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldNonce), v))
	})
}

// NonceEqualFold applies the EqualFold predicate on the "nonce" field.
func NonceEqualFold(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldNonce), v))
	})
}

// NonceContainsFold applies the ContainsFold predicate on the "nonce" field.
func NonceContainsFold(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldNonce), v))
	})
}

// HasBlock applies the HasEdge predicate on the "block" edge.
func HasBlock() predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BlockTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, BlockTable, BlockColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBlockWith applies the HasEdge predicate on the "block" edge with a given conditions (other predicates).
func HasBlockWith(preds ...predicate.Block) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BlockInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, BlockTable, BlockColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Transaction) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Transaction) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Transaction) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		p(s.Not())
	})
}