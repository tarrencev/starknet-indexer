// Code generated by entc, DO NOT EDIT.

package transactionreceipt

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/tarrencev/starknet-indexer/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.TransactionReceipt {
	return predicate.TransactionReceipt(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.TransactionReceipt {
	return predicate.TransactionReceipt(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.TransactionReceipt {
	return predicate.TransactionReceipt(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.TransactionReceipt {
	return predicate.TransactionReceipt(func(s *sql.Selector) {
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
func IDNotIn(ids ...string) predicate.TransactionReceipt {
	return predicate.TransactionReceipt(func(s *sql.Selector) {
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
func IDGT(id string) predicate.TransactionReceipt {
	return predicate.TransactionReceipt(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.TransactionReceipt {
	return predicate.TransactionReceipt(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.TransactionReceipt {
	return predicate.TransactionReceipt(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.TransactionReceipt {
	return predicate.TransactionReceipt(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// TransactionHash applies equality check predicate on the "transaction_hash" field. It's identical to TransactionHashEQ.
func TransactionHash(v string) predicate.TransactionReceipt {
	return predicate.TransactionReceipt(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTransactionHash), v))
	})
}

// StatusData applies equality check predicate on the "status_data" field. It's identical to StatusDataEQ.
func StatusData(v string) predicate.TransactionReceipt {
	return predicate.TransactionReceipt(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatusData), v))
	})
}

// TransactionHashEQ applies the EQ predicate on the "transaction_hash" field.
func TransactionHashEQ(v string) predicate.TransactionReceipt {
	return predicate.TransactionReceipt(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTransactionHash), v))
	})
}

// TransactionHashNEQ applies the NEQ predicate on the "transaction_hash" field.
func TransactionHashNEQ(v string) predicate.TransactionReceipt {
	return predicate.TransactionReceipt(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTransactionHash), v))
	})
}

// TransactionHashIn applies the In predicate on the "transaction_hash" field.
func TransactionHashIn(vs ...string) predicate.TransactionReceipt {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TransactionReceipt(func(s *sql.Selector) {
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
func TransactionHashNotIn(vs ...string) predicate.TransactionReceipt {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TransactionReceipt(func(s *sql.Selector) {
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
func TransactionHashGT(v string) predicate.TransactionReceipt {
	return predicate.TransactionReceipt(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTransactionHash), v))
	})
}

// TransactionHashGTE applies the GTE predicate on the "transaction_hash" field.
func TransactionHashGTE(v string) predicate.TransactionReceipt {
	return predicate.TransactionReceipt(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTransactionHash), v))
	})
}

// TransactionHashLT applies the LT predicate on the "transaction_hash" field.
func TransactionHashLT(v string) predicate.TransactionReceipt {
	return predicate.TransactionReceipt(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTransactionHash), v))
	})
}

// TransactionHashLTE applies the LTE predicate on the "transaction_hash" field.
func TransactionHashLTE(v string) predicate.TransactionReceipt {
	return predicate.TransactionReceipt(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTransactionHash), v))
	})
}

// TransactionHashContains applies the Contains predicate on the "transaction_hash" field.
func TransactionHashContains(v string) predicate.TransactionReceipt {
	return predicate.TransactionReceipt(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTransactionHash), v))
	})
}

// TransactionHashHasPrefix applies the HasPrefix predicate on the "transaction_hash" field.
func TransactionHashHasPrefix(v string) predicate.TransactionReceipt {
	return predicate.TransactionReceipt(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTransactionHash), v))
	})
}

// TransactionHashHasSuffix applies the HasSuffix predicate on the "transaction_hash" field.
func TransactionHashHasSuffix(v string) predicate.TransactionReceipt {
	return predicate.TransactionReceipt(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTransactionHash), v))
	})
}

// TransactionHashEqualFold applies the EqualFold predicate on the "transaction_hash" field.
func TransactionHashEqualFold(v string) predicate.TransactionReceipt {
	return predicate.TransactionReceipt(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTransactionHash), v))
	})
}

// TransactionHashContainsFold applies the ContainsFold predicate on the "transaction_hash" field.
func TransactionHashContainsFold(v string) predicate.TransactionReceipt {
	return predicate.TransactionReceipt(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTransactionHash), v))
	})
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.TransactionReceipt {
	return predicate.TransactionReceipt(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.TransactionReceipt {
	return predicate.TransactionReceipt(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStatus), v))
	})
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.TransactionReceipt {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TransactionReceipt(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStatus), v...))
	})
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.TransactionReceipt {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TransactionReceipt(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStatus), v...))
	})
}

// StatusDataEQ applies the EQ predicate on the "status_data" field.
func StatusDataEQ(v string) predicate.TransactionReceipt {
	return predicate.TransactionReceipt(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatusData), v))
	})
}

// StatusDataNEQ applies the NEQ predicate on the "status_data" field.
func StatusDataNEQ(v string) predicate.TransactionReceipt {
	return predicate.TransactionReceipt(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStatusData), v))
	})
}

// StatusDataIn applies the In predicate on the "status_data" field.
func StatusDataIn(vs ...string) predicate.TransactionReceipt {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TransactionReceipt(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStatusData), v...))
	})
}

// StatusDataNotIn applies the NotIn predicate on the "status_data" field.
func StatusDataNotIn(vs ...string) predicate.TransactionReceipt {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TransactionReceipt(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStatusData), v...))
	})
}

// StatusDataGT applies the GT predicate on the "status_data" field.
func StatusDataGT(v string) predicate.TransactionReceipt {
	return predicate.TransactionReceipt(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStatusData), v))
	})
}

// StatusDataGTE applies the GTE predicate on the "status_data" field.
func StatusDataGTE(v string) predicate.TransactionReceipt {
	return predicate.TransactionReceipt(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStatusData), v))
	})
}

// StatusDataLT applies the LT predicate on the "status_data" field.
func StatusDataLT(v string) predicate.TransactionReceipt {
	return predicate.TransactionReceipt(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStatusData), v))
	})
}

// StatusDataLTE applies the LTE predicate on the "status_data" field.
func StatusDataLTE(v string) predicate.TransactionReceipt {
	return predicate.TransactionReceipt(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStatusData), v))
	})
}

// StatusDataContains applies the Contains predicate on the "status_data" field.
func StatusDataContains(v string) predicate.TransactionReceipt {
	return predicate.TransactionReceipt(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldStatusData), v))
	})
}

// StatusDataHasPrefix applies the HasPrefix predicate on the "status_data" field.
func StatusDataHasPrefix(v string) predicate.TransactionReceipt {
	return predicate.TransactionReceipt(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldStatusData), v))
	})
}

// StatusDataHasSuffix applies the HasSuffix predicate on the "status_data" field.
func StatusDataHasSuffix(v string) predicate.TransactionReceipt {
	return predicate.TransactionReceipt(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldStatusData), v))
	})
}

// StatusDataEqualFold applies the EqualFold predicate on the "status_data" field.
func StatusDataEqualFold(v string) predicate.TransactionReceipt {
	return predicate.TransactionReceipt(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldStatusData), v))
	})
}

// StatusDataContainsFold applies the ContainsFold predicate on the "status_data" field.
func StatusDataContainsFold(v string) predicate.TransactionReceipt {
	return predicate.TransactionReceipt(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldStatusData), v))
	})
}

// HasBlock applies the HasEdge predicate on the "block" edge.
func HasBlock() predicate.TransactionReceipt {
	return predicate.TransactionReceipt(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BlockTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, BlockTable, BlockColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBlockWith applies the HasEdge predicate on the "block" edge with a given conditions (other predicates).
func HasBlockWith(preds ...predicate.Block) predicate.TransactionReceipt {
	return predicate.TransactionReceipt(func(s *sql.Selector) {
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
func And(predicates ...predicate.TransactionReceipt) predicate.TransactionReceipt {
	return predicate.TransactionReceipt(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.TransactionReceipt) predicate.TransactionReceipt {
	return predicate.TransactionReceipt(func(s *sql.Selector) {
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
func Not(p predicate.TransactionReceipt) predicate.TransactionReceipt {
	return predicate.TransactionReceipt(func(s *sql.Selector) {
		p(s.Not())
	})
}
