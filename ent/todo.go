// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"go-start-project/ent/todo"
	"strings"
	"time"

	"github.com/facebook/ent/dialect/sql"
)

// Todo is the model entity for the Todo schema.
type Todo struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Status holds the value of the "status" field.
	Status todo.Status `json:"status,omitempty"`
	// Priority holds the value of the "priority" field.
	Priority int `json:"priority,omitempty"`
	// Text holds the value of the "text" field.
	Text string `json:"text,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TodoQuery when eager-loading is set.
	Edges         TodoEdges `json:"edges"`
	todo_children *int
}

// TodoEdges holds the relations/edges for other nodes in the graph.
type TodoEdges struct {
	// Parent holds the value of the parent edge.
	Parent *Todo
	// Children holds the value of the children edge.
	Children []*Todo
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TodoEdges) ParentOrErr() (*Todo, error) {
	if e.loadedTypes[0] {
		if e.Parent == nil {
			// The edge parent was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: todo.Label}
		}
		return e.Parent, nil
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// ChildrenOrErr returns the Children value or an error if the edge
// was not loaded in eager-loading.
func (e TodoEdges) ChildrenOrErr() ([]*Todo, error) {
	if e.loadedTypes[1] {
		return e.Children, nil
	}
	return nil, &NotLoadedError{edge: "children"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Todo) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullTime{},   // created_at
		&sql.NullString{}, // status
		&sql.NullInt64{},  // priority
		&sql.NullString{}, // text
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Todo) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // todo_children
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Todo fields.
func (t *Todo) assignValues(values ...interface{}) error {
	if m, n := len(values), len(todo.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	t.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field created_at", values[0])
	} else if value.Valid {
		t.CreatedAt = value.Time
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field status", values[1])
	} else if value.Valid {
		t.Status = todo.Status(value.String)
	}
	if value, ok := values[2].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field priority", values[2])
	} else if value.Valid {
		t.Priority = int(value.Int64)
	}
	if value, ok := values[3].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field text", values[3])
	} else if value.Valid {
		t.Text = value.String
	}
	values = values[4:]
	if len(values) == len(todo.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field todo_children", value)
		} else if value.Valid {
			t.todo_children = new(int)
			*t.todo_children = int(value.Int64)
		}
	}
	return nil
}

// QueryParent queries the parent edge of the Todo.
func (t *Todo) QueryParent() *TodoQuery {
	return (&TodoClient{config: t.config}).QueryParent(t)
}

// QueryChildren queries the children edge of the Todo.
func (t *Todo) QueryChildren() *TodoQuery {
	return (&TodoClient{config: t.config}).QueryChildren(t)
}

// Update returns a builder for updating this Todo.
// Note that, you need to call Todo.Unwrap() before calling this method, if this Todo
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Todo) Update() *TodoUpdateOne {
	return (&TodoClient{config: t.config}).UpdateOne(t)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (t *Todo) Unwrap() *Todo {
	tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Todo is not a transactional entity")
	}
	t.config.driver = tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Todo) String() string {
	var builder strings.Builder
	builder.WriteString("Todo(")
	builder.WriteString(fmt.Sprintf("id=%v", t.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(t.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", t.Status))
	builder.WriteString(", priority=")
	builder.WriteString(fmt.Sprintf("%v", t.Priority))
	builder.WriteString(", text=")
	builder.WriteString(t.Text)
	builder.WriteByte(')')
	return builder.String()
}

// Todos is a parsable slice of Todo.
type Todos []*Todo

func (t Todos) config(cfg config) {
	for _i := range t {
		t[_i].config = cfg
	}
}
