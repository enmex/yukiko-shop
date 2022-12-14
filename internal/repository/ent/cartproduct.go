// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"yukiko-shop/internal/repository/ent/cartproduct"
	"yukiko-shop/internal/repository/ent/product"
	"yukiko-shop/internal/repository/ent/user"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// CartProduct is the model entity for the CartProduct schema.
type CartProduct struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// ProductID holds the value of the "product_id" field.
	ProductID uuid.UUID `json:"product_id,omitempty"`
	// CustomerID holds the value of the "customer_id" field.
	CustomerID uuid.UUID `json:"customer_id,omitempty"`
	// PhotoURL holds the value of the "photo_url" field.
	PhotoURL string `json:"photo_url,omitempty"`
	// Price holds the value of the "price" field.
	Price float64 `json:"price,omitempty"`
	// Quantity holds the value of the "quantity" field.
	Quantity int `json:"quantity,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CartProductQuery when eager-loading is set.
	Edges CartProductEdges `json:"edges"`
}

// CartProductEdges holds the relations/edges for other nodes in the graph.
type CartProductEdges struct {
	// Customer holds the value of the customer edge.
	Customer *User `json:"customer,omitempty"`
	// Product holds the value of the product edge.
	Product *Product `json:"product,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// CustomerOrErr returns the Customer value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CartProductEdges) CustomerOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.Customer == nil {
			// The edge customer was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Customer, nil
	}
	return nil, &NotLoadedError{edge: "customer"}
}

// ProductOrErr returns the Product value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CartProductEdges) ProductOrErr() (*Product, error) {
	if e.loadedTypes[1] {
		if e.Product == nil {
			// The edge product was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: product.Label}
		}
		return e.Product, nil
	}
	return nil, &NotLoadedError{edge: "product"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CartProduct) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case cartproduct.FieldPrice:
			values[i] = new(sql.NullFloat64)
		case cartproduct.FieldQuantity:
			values[i] = new(sql.NullInt64)
		case cartproduct.FieldName, cartproduct.FieldPhotoURL:
			values[i] = new(sql.NullString)
		case cartproduct.FieldID, cartproduct.FieldProductID, cartproduct.FieldCustomerID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type CartProduct", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CartProduct fields.
func (cp *CartProduct) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case cartproduct.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				cp.ID = *value
			}
		case cartproduct.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				cp.Name = value.String
			}
		case cartproduct.FieldProductID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field product_id", values[i])
			} else if value != nil {
				cp.ProductID = *value
			}
		case cartproduct.FieldCustomerID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field customer_id", values[i])
			} else if value != nil {
				cp.CustomerID = *value
			}
		case cartproduct.FieldPhotoURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field photo_url", values[i])
			} else if value.Valid {
				cp.PhotoURL = value.String
			}
		case cartproduct.FieldPrice:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field price", values[i])
			} else if value.Valid {
				cp.Price = value.Float64
			}
		case cartproduct.FieldQuantity:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field quantity", values[i])
			} else if value.Valid {
				cp.Quantity = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryCustomer queries the "customer" edge of the CartProduct entity.
func (cp *CartProduct) QueryCustomer() *UserQuery {
	return (&CartProductClient{config: cp.config}).QueryCustomer(cp)
}

// QueryProduct queries the "product" edge of the CartProduct entity.
func (cp *CartProduct) QueryProduct() *ProductQuery {
	return (&CartProductClient{config: cp.config}).QueryProduct(cp)
}

// Update returns a builder for updating this CartProduct.
// Note that you need to call CartProduct.Unwrap() before calling this method if this CartProduct
// was returned from a transaction, and the transaction was committed or rolled back.
func (cp *CartProduct) Update() *CartProductUpdateOne {
	return (&CartProductClient{config: cp.config}).UpdateOne(cp)
}

// Unwrap unwraps the CartProduct entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (cp *CartProduct) Unwrap() *CartProduct {
	tx, ok := cp.config.driver.(*txDriver)
	if !ok {
		panic("ent: CartProduct is not a transactional entity")
	}
	cp.config.driver = tx.drv
	return cp
}

// String implements the fmt.Stringer.
func (cp *CartProduct) String() string {
	var builder strings.Builder
	builder.WriteString("CartProduct(")
	builder.WriteString(fmt.Sprintf("id=%v", cp.ID))
	builder.WriteString(", name=")
	builder.WriteString(cp.Name)
	builder.WriteString(", product_id=")
	builder.WriteString(fmt.Sprintf("%v", cp.ProductID))
	builder.WriteString(", customer_id=")
	builder.WriteString(fmt.Sprintf("%v", cp.CustomerID))
	builder.WriteString(", photo_url=")
	builder.WriteString(cp.PhotoURL)
	builder.WriteString(", price=")
	builder.WriteString(fmt.Sprintf("%v", cp.Price))
	builder.WriteString(", quantity=")
	builder.WriteString(fmt.Sprintf("%v", cp.Quantity))
	builder.WriteByte(')')
	return builder.String()
}

// CartProducts is a parsable slice of CartProduct.
type CartProducts []*CartProduct

func (cp CartProducts) config(cfg config) {
	for _i := range cp {
		cp[_i].config = cfg
	}
}
