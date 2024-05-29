// Code generated by ent, DO NOT EDIT.

package ent

import (
	"Bookapp/ent/admin"
	"Bookapp/ent/user"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Admin is the model entity for the Admin schema.
type Admin struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID int `json:"user_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AdminQuery when eager-loading is set.
	Edges        AdminEdges `json:"edges"`
	selectValues sql.SelectValues
}

// AdminEdges holds the relations/edges for other nodes in the graph.
type AdminEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AdminEdges) UserOrErr() (*User, error) {
	if e.User != nil {
		return e.User, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Admin) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case admin.FieldID, admin.FieldUserID:
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Admin fields.
func (a *Admin) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case admin.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			a.ID = int(value.Int64)
		case admin.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				a.UserID = int(value.Int64)
			}
		default:
			a.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Admin.
// This includes values selected through modifiers, order, etc.
func (a *Admin) Value(name string) (ent.Value, error) {
	return a.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the Admin entity.
func (a *Admin) QueryUser() *UserQuery {
	return NewAdminClient(a.config).QueryUser(a)
}

// Update returns a builder for updating this Admin.
// Note that you need to call Admin.Unwrap() before calling this method if this Admin
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Admin) Update() *AdminUpdateOne {
	return NewAdminClient(a.config).UpdateOne(a)
}

// Unwrap unwraps the Admin entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Admin) Unwrap() *Admin {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Admin is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Admin) String() string {
	var builder strings.Builder
	builder.WriteString("Admin(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", a.UserID))
	builder.WriteByte(')')
	return builder.String()
}

// Admins is a parsable slice of Admin.
type Admins []*Admin
