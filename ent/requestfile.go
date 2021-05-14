// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/traPtitech/Jomon/ent/file"
	"github.com/traPtitech/Jomon/ent/request"
	"github.com/traPtitech/Jomon/ent/requestfile"
)

// RequestFile is the model entity for the RequestFile schema.
type RequestFile struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// RequestID holds the value of the "request_id" field.
	RequestID int `json:"request_id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RequestFileQuery when eager-loading is set.
	Edges RequestFileEdges `json:"edges"`
}

// RequestFileEdges holds the relations/edges for other nodes in the graph.
type RequestFileEdges struct {
	// Request holds the value of the request edge.
	Request *Request `json:"request,omitempty"`
	// File holds the value of the file edge.
	File *File `json:"file,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// RequestOrErr returns the Request value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RequestFileEdges) RequestOrErr() (*Request, error) {
	if e.loadedTypes[0] {
		if e.Request == nil {
			// The edge request was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: request.Label}
		}
		return e.Request, nil
	}
	return nil, &NotLoadedError{edge: "request"}
}

// FileOrErr returns the File value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RequestFileEdges) FileOrErr() (*File, error) {
	if e.loadedTypes[1] {
		if e.File == nil {
			// The edge file was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: file.Label}
		}
		return e.File, nil
	}
	return nil, &NotLoadedError{edge: "file"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*RequestFile) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case requestfile.FieldID, requestfile.FieldRequestID:
			values[i] = new(sql.NullInt64)
		case requestfile.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type RequestFile", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the RequestFile fields.
func (rf *RequestFile) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case requestfile.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			rf.ID = int(value.Int64)
		case requestfile.FieldRequestID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field request_id", values[i])
			} else if value.Valid {
				rf.RequestID = int(value.Int64)
			}
		case requestfile.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				rf.CreatedAt = value.Time
			}
		}
	}
	return nil
}

// QueryRequest queries the "request" edge of the RequestFile entity.
func (rf *RequestFile) QueryRequest() *RequestQuery {
	return (&RequestFileClient{config: rf.config}).QueryRequest(rf)
}

// QueryFile queries the "file" edge of the RequestFile entity.
func (rf *RequestFile) QueryFile() *FileQuery {
	return (&RequestFileClient{config: rf.config}).QueryFile(rf)
}

// Update returns a builder for updating this RequestFile.
// Note that you need to call RequestFile.Unwrap() before calling this method if this RequestFile
// was returned from a transaction, and the transaction was committed or rolled back.
func (rf *RequestFile) Update() *RequestFileUpdateOne {
	return (&RequestFileClient{config: rf.config}).UpdateOne(rf)
}

// Unwrap unwraps the RequestFile entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (rf *RequestFile) Unwrap() *RequestFile {
	tx, ok := rf.config.driver.(*txDriver)
	if !ok {
		panic("ent: RequestFile is not a transactional entity")
	}
	rf.config.driver = tx.drv
	return rf
}

// String implements the fmt.Stringer.
func (rf *RequestFile) String() string {
	var builder strings.Builder
	builder.WriteString("RequestFile(")
	builder.WriteString(fmt.Sprintf("id=%v", rf.ID))
	builder.WriteString(", request_id=")
	builder.WriteString(fmt.Sprintf("%v", rf.RequestID))
	builder.WriteString(", created_at=")
	builder.WriteString(rf.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// RequestFiles is a parsable slice of RequestFile.
type RequestFiles []*RequestFile

func (rf RequestFiles) config(cfg config) {
	for _i := range rf {
		rf[_i].config = cfg
	}
}
