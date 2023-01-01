// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"yukiko-shop/internal/repository/ent/cartproduct"
	"yukiko-shop/internal/repository/ent/product"
	"yukiko-shop/internal/repository/ent/user"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// CartProductCreate is the builder for creating a CartProduct entity.
type CartProductCreate struct {
	config
	mutation *CartProductMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (cpc *CartProductCreate) SetName(s string) *CartProductCreate {
	cpc.mutation.SetName(s)
	return cpc
}

// SetProductID sets the "product_id" field.
func (cpc *CartProductCreate) SetProductID(u uuid.UUID) *CartProductCreate {
	cpc.mutation.SetProductID(u)
	return cpc
}

// SetCustomerID sets the "customer_id" field.
func (cpc *CartProductCreate) SetCustomerID(u uuid.UUID) *CartProductCreate {
	cpc.mutation.SetCustomerID(u)
	return cpc
}

// SetPhotoURL sets the "photo_url" field.
func (cpc *CartProductCreate) SetPhotoURL(s string) *CartProductCreate {
	cpc.mutation.SetPhotoURL(s)
	return cpc
}

// SetNillablePhotoURL sets the "photo_url" field if the given value is not nil.
func (cpc *CartProductCreate) SetNillablePhotoURL(s *string) *CartProductCreate {
	if s != nil {
		cpc.SetPhotoURL(*s)
	}
	return cpc
}

// SetPrice sets the "price" field.
func (cpc *CartProductCreate) SetPrice(f float64) *CartProductCreate {
	cpc.mutation.SetPrice(f)
	return cpc
}

// SetQuantity sets the "quantity" field.
func (cpc *CartProductCreate) SetQuantity(i int) *CartProductCreate {
	cpc.mutation.SetQuantity(i)
	return cpc
}

// SetID sets the "id" field.
func (cpc *CartProductCreate) SetID(u uuid.UUID) *CartProductCreate {
	cpc.mutation.SetID(u)
	return cpc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (cpc *CartProductCreate) SetNillableID(u *uuid.UUID) *CartProductCreate {
	if u != nil {
		cpc.SetID(*u)
	}
	return cpc
}

// SetCustomer sets the "customer" edge to the User entity.
func (cpc *CartProductCreate) SetCustomer(u *User) *CartProductCreate {
	return cpc.SetCustomerID(u.ID)
}

// SetProduct sets the "product" edge to the Product entity.
func (cpc *CartProductCreate) SetProduct(p *Product) *CartProductCreate {
	return cpc.SetProductID(p.ID)
}

// Mutation returns the CartProductMutation object of the builder.
func (cpc *CartProductCreate) Mutation() *CartProductMutation {
	return cpc.mutation
}

// Save creates the CartProduct in the database.
func (cpc *CartProductCreate) Save(ctx context.Context) (*CartProduct, error) {
	var (
		err  error
		node *CartProduct
	)
	cpc.defaults()
	if len(cpc.hooks) == 0 {
		if err = cpc.check(); err != nil {
			return nil, err
		}
		node, err = cpc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CartProductMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cpc.check(); err != nil {
				return nil, err
			}
			cpc.mutation = mutation
			if node, err = cpc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(cpc.hooks) - 1; i >= 0; i-- {
			if cpc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cpc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cpc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cpc *CartProductCreate) SaveX(ctx context.Context) *CartProduct {
	v, err := cpc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cpc *CartProductCreate) Exec(ctx context.Context) error {
	_, err := cpc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cpc *CartProductCreate) ExecX(ctx context.Context) {
	if err := cpc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cpc *CartProductCreate) defaults() {
	if _, ok := cpc.mutation.ID(); !ok {
		v := cartproduct.DefaultID()
		cpc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cpc *CartProductCreate) check() error {
	if _, ok := cpc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "CartProduct.name"`)}
	}
	if _, ok := cpc.mutation.ProductID(); !ok {
		return &ValidationError{Name: "product_id", err: errors.New(`ent: missing required field "CartProduct.product_id"`)}
	}
	if _, ok := cpc.mutation.CustomerID(); !ok {
		return &ValidationError{Name: "customer_id", err: errors.New(`ent: missing required field "CartProduct.customer_id"`)}
	}
	if _, ok := cpc.mutation.Price(); !ok {
		return &ValidationError{Name: "price", err: errors.New(`ent: missing required field "CartProduct.price"`)}
	}
	if _, ok := cpc.mutation.Quantity(); !ok {
		return &ValidationError{Name: "quantity", err: errors.New(`ent: missing required field "CartProduct.quantity"`)}
	}
	if _, ok := cpc.mutation.CustomerID(); !ok {
		return &ValidationError{Name: "customer", err: errors.New(`ent: missing required edge "CartProduct.customer"`)}
	}
	if _, ok := cpc.mutation.ProductID(); !ok {
		return &ValidationError{Name: "product", err: errors.New(`ent: missing required edge "CartProduct.product"`)}
	}
	return nil
}

func (cpc *CartProductCreate) sqlSave(ctx context.Context) (*CartProduct, error) {
	_node, _spec := cpc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cpc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (cpc *CartProductCreate) createSpec() (*CartProduct, *sqlgraph.CreateSpec) {
	var (
		_node = &CartProduct{config: cpc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: cartproduct.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: cartproduct.FieldID,
			},
		}
	)
	if id, ok := cpc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := cpc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: cartproduct.FieldName,
		})
		_node.Name = value
	}
	if value, ok := cpc.mutation.PhotoURL(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: cartproduct.FieldPhotoURL,
		})
		_node.PhotoURL = value
	}
	if value, ok := cpc.mutation.Price(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: cartproduct.FieldPrice,
		})
		_node.Price = value
	}
	if value, ok := cpc.mutation.Quantity(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: cartproduct.FieldQuantity,
		})
		_node.Quantity = value
	}
	if nodes := cpc.mutation.CustomerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cartproduct.CustomerTable,
			Columns: []string{cartproduct.CustomerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.CustomerID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cpc.mutation.ProductIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cartproduct.ProductTable,
			Columns: []string{cartproduct.ProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: product.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ProductID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CartProductCreateBulk is the builder for creating many CartProduct entities in bulk.
type CartProductCreateBulk struct {
	config
	builders []*CartProductCreate
}

// Save creates the CartProduct entities in the database.
func (cpcb *CartProductCreateBulk) Save(ctx context.Context) ([]*CartProduct, error) {
	specs := make([]*sqlgraph.CreateSpec, len(cpcb.builders))
	nodes := make([]*CartProduct, len(cpcb.builders))
	mutators := make([]Mutator, len(cpcb.builders))
	for i := range cpcb.builders {
		func(i int, root context.Context) {
			builder := cpcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CartProductMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, cpcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, cpcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, cpcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (cpcb *CartProductCreateBulk) SaveX(ctx context.Context) []*CartProduct {
	v, err := cpcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cpcb *CartProductCreateBulk) Exec(ctx context.Context) error {
	_, err := cpcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cpcb *CartProductCreateBulk) ExecX(ctx context.Context) {
	if err := cpcb.Exec(ctx); err != nil {
		panic(err)
	}
}