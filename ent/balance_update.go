// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/cartridge-gg/starknet-indexer/ent/balance"
	"github.com/cartridge-gg/starknet-indexer/ent/contract"
	"github.com/cartridge-gg/starknet-indexer/ent/predicate"
	"github.com/cartridge-gg/starknet-indexer/ent/schema/big"
)

// BalanceUpdate is the builder for updating Balance entities.
type BalanceUpdate struct {
	config
	hooks    []Hook
	mutation *BalanceMutation
}

// Where appends a list predicates to the BalanceUpdate builder.
func (bu *BalanceUpdate) Where(ps ...predicate.Balance) *BalanceUpdate {
	bu.mutation.Where(ps...)
	return bu
}

// SetBalance sets the "balance" field.
func (bu *BalanceUpdate) SetBalance(b big.Int) *BalanceUpdate {
	bu.mutation.ResetBalance()
	bu.mutation.SetBalance(b)
	return bu
}

// SetNillableBalance sets the "balance" field if the given value is not nil.
func (bu *BalanceUpdate) SetNillableBalance(b *big.Int) *BalanceUpdate {
	if b != nil {
		bu.SetBalance(*b)
	}
	return bu
}

// AddBalance adds b to the "balance" field.
func (bu *BalanceUpdate) AddBalance(b big.Int) *BalanceUpdate {
	bu.mutation.AddBalance(b)
	return bu
}

// SetAccountID sets the "account" edge to the Contract entity by ID.
func (bu *BalanceUpdate) SetAccountID(id string) *BalanceUpdate {
	bu.mutation.SetAccountID(id)
	return bu
}

// SetNillableAccountID sets the "account" edge to the Contract entity by ID if the given value is not nil.
func (bu *BalanceUpdate) SetNillableAccountID(id *string) *BalanceUpdate {
	if id != nil {
		bu = bu.SetAccountID(*id)
	}
	return bu
}

// SetAccount sets the "account" edge to the Contract entity.
func (bu *BalanceUpdate) SetAccount(c *Contract) *BalanceUpdate {
	return bu.SetAccountID(c.ID)
}

// SetContractID sets the "contract" edge to the Contract entity by ID.
func (bu *BalanceUpdate) SetContractID(id string) *BalanceUpdate {
	bu.mutation.SetContractID(id)
	return bu
}

// SetNillableContractID sets the "contract" edge to the Contract entity by ID if the given value is not nil.
func (bu *BalanceUpdate) SetNillableContractID(id *string) *BalanceUpdate {
	if id != nil {
		bu = bu.SetContractID(*id)
	}
	return bu
}

// SetContract sets the "contract" edge to the Contract entity.
func (bu *BalanceUpdate) SetContract(c *Contract) *BalanceUpdate {
	return bu.SetContractID(c.ID)
}

// Mutation returns the BalanceMutation object of the builder.
func (bu *BalanceUpdate) Mutation() *BalanceMutation {
	return bu.mutation
}

// ClearAccount clears the "account" edge to the Contract entity.
func (bu *BalanceUpdate) ClearAccount() *BalanceUpdate {
	bu.mutation.ClearAccount()
	return bu
}

// ClearContract clears the "contract" edge to the Contract entity.
func (bu *BalanceUpdate) ClearContract() *BalanceUpdate {
	bu.mutation.ClearContract()
	return bu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bu *BalanceUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(bu.hooks) == 0 {
		affected, err = bu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BalanceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bu.mutation = mutation
			affected, err = bu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(bu.hooks) - 1; i >= 0; i-- {
			if bu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = bu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (bu *BalanceUpdate) SaveX(ctx context.Context) int {
	affected, err := bu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bu *BalanceUpdate) Exec(ctx context.Context) error {
	_, err := bu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bu *BalanceUpdate) ExecX(ctx context.Context) {
	if err := bu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bu *BalanceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   balance.Table,
			Columns: balance.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: balance.FieldID,
			},
		},
	}
	if ps := bu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bu.mutation.Balance(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: balance.FieldBalance,
		})
	}
	if value, ok := bu.mutation.AddedBalance(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: balance.FieldBalance,
		})
	}
	if bu.mutation.AccountCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   balance.AccountTable,
			Columns: []string{balance.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: contract.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.AccountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   balance.AccountTable,
			Columns: []string{balance.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: contract.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if bu.mutation.ContractCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   balance.ContractTable,
			Columns: []string{balance.ContractColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: contract.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.ContractIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   balance.ContractTable,
			Columns: []string{balance.ContractColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: contract.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{balance.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// BalanceUpdateOne is the builder for updating a single Balance entity.
type BalanceUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BalanceMutation
}

// SetBalance sets the "balance" field.
func (buo *BalanceUpdateOne) SetBalance(b big.Int) *BalanceUpdateOne {
	buo.mutation.ResetBalance()
	buo.mutation.SetBalance(b)
	return buo
}

// SetNillableBalance sets the "balance" field if the given value is not nil.
func (buo *BalanceUpdateOne) SetNillableBalance(b *big.Int) *BalanceUpdateOne {
	if b != nil {
		buo.SetBalance(*b)
	}
	return buo
}

// AddBalance adds b to the "balance" field.
func (buo *BalanceUpdateOne) AddBalance(b big.Int) *BalanceUpdateOne {
	buo.mutation.AddBalance(b)
	return buo
}

// SetAccountID sets the "account" edge to the Contract entity by ID.
func (buo *BalanceUpdateOne) SetAccountID(id string) *BalanceUpdateOne {
	buo.mutation.SetAccountID(id)
	return buo
}

// SetNillableAccountID sets the "account" edge to the Contract entity by ID if the given value is not nil.
func (buo *BalanceUpdateOne) SetNillableAccountID(id *string) *BalanceUpdateOne {
	if id != nil {
		buo = buo.SetAccountID(*id)
	}
	return buo
}

// SetAccount sets the "account" edge to the Contract entity.
func (buo *BalanceUpdateOne) SetAccount(c *Contract) *BalanceUpdateOne {
	return buo.SetAccountID(c.ID)
}

// SetContractID sets the "contract" edge to the Contract entity by ID.
func (buo *BalanceUpdateOne) SetContractID(id string) *BalanceUpdateOne {
	buo.mutation.SetContractID(id)
	return buo
}

// SetNillableContractID sets the "contract" edge to the Contract entity by ID if the given value is not nil.
func (buo *BalanceUpdateOne) SetNillableContractID(id *string) *BalanceUpdateOne {
	if id != nil {
		buo = buo.SetContractID(*id)
	}
	return buo
}

// SetContract sets the "contract" edge to the Contract entity.
func (buo *BalanceUpdateOne) SetContract(c *Contract) *BalanceUpdateOne {
	return buo.SetContractID(c.ID)
}

// Mutation returns the BalanceMutation object of the builder.
func (buo *BalanceUpdateOne) Mutation() *BalanceMutation {
	return buo.mutation
}

// ClearAccount clears the "account" edge to the Contract entity.
func (buo *BalanceUpdateOne) ClearAccount() *BalanceUpdateOne {
	buo.mutation.ClearAccount()
	return buo
}

// ClearContract clears the "contract" edge to the Contract entity.
func (buo *BalanceUpdateOne) ClearContract() *BalanceUpdateOne {
	buo.mutation.ClearContract()
	return buo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (buo *BalanceUpdateOne) Select(field string, fields ...string) *BalanceUpdateOne {
	buo.fields = append([]string{field}, fields...)
	return buo
}

// Save executes the query and returns the updated Balance entity.
func (buo *BalanceUpdateOne) Save(ctx context.Context) (*Balance, error) {
	var (
		err  error
		node *Balance
	)
	if len(buo.hooks) == 0 {
		node, err = buo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BalanceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			buo.mutation = mutation
			node, err = buo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(buo.hooks) - 1; i >= 0; i-- {
			if buo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = buo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, buo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (buo *BalanceUpdateOne) SaveX(ctx context.Context) *Balance {
	node, err := buo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (buo *BalanceUpdateOne) Exec(ctx context.Context) error {
	_, err := buo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (buo *BalanceUpdateOne) ExecX(ctx context.Context) {
	if err := buo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (buo *BalanceUpdateOne) sqlSave(ctx context.Context) (_node *Balance, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   balance.Table,
			Columns: balance.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: balance.FieldID,
			},
		},
	}
	id, ok := buo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Balance.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := buo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, balance.FieldID)
		for _, f := range fields {
			if !balance.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != balance.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := buo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := buo.mutation.Balance(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: balance.FieldBalance,
		})
	}
	if value, ok := buo.mutation.AddedBalance(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: balance.FieldBalance,
		})
	}
	if buo.mutation.AccountCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   balance.AccountTable,
			Columns: []string{balance.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: contract.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.AccountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   balance.AccountTable,
			Columns: []string{balance.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: contract.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if buo.mutation.ContractCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   balance.ContractTable,
			Columns: []string{balance.ContractColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: contract.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.ContractIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   balance.ContractTable,
			Columns: []string{balance.ContractColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: contract.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Balance{config: buo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, buo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{balance.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
