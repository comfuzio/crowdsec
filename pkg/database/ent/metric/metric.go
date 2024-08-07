// Code generated by ent, DO NOT EDIT.

package metric

import (
	"fmt"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the metric type in the database.
	Label = "metric"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldGeneratedType holds the string denoting the generated_type field in the database.
	FieldGeneratedType = "generated_type"
	// FieldGeneratedBy holds the string denoting the generated_by field in the database.
	FieldGeneratedBy = "generated_by"
	// FieldReceivedAt holds the string denoting the received_at field in the database.
	FieldReceivedAt = "received_at"
	// FieldPushedAt holds the string denoting the pushed_at field in the database.
	FieldPushedAt = "pushed_at"
	// FieldPayload holds the string denoting the payload field in the database.
	FieldPayload = "payload"
	// Table holds the table name of the metric in the database.
	Table = "metrics"
)

// Columns holds all SQL columns for metric fields.
var Columns = []string{
	FieldID,
	FieldGeneratedType,
	FieldGeneratedBy,
	FieldReceivedAt,
	FieldPushedAt,
	FieldPayload,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// GeneratedType defines the type for the "generated_type" enum field.
type GeneratedType string

// GeneratedType values.
const (
	GeneratedTypeLP GeneratedType = "LP"
	GeneratedTypeRC GeneratedType = "RC"
)

func (gt GeneratedType) String() string {
	return string(gt)
}

// GeneratedTypeValidator is a validator for the "generated_type" field enum values. It is called by the builders before save.
func GeneratedTypeValidator(gt GeneratedType) error {
	switch gt {
	case GeneratedTypeLP, GeneratedTypeRC:
		return nil
	default:
		return fmt.Errorf("metric: invalid enum value for generated_type field: %q", gt)
	}
}

// OrderOption defines the ordering options for the Metric queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByGeneratedType orders the results by the generated_type field.
func ByGeneratedType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGeneratedType, opts...).ToFunc()
}

// ByGeneratedBy orders the results by the generated_by field.
func ByGeneratedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGeneratedBy, opts...).ToFunc()
}

// ByReceivedAt orders the results by the received_at field.
func ByReceivedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldReceivedAt, opts...).ToFunc()
}

// ByPushedAt orders the results by the pushed_at field.
func ByPushedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPushedAt, opts...).ToFunc()
}

// ByPayload orders the results by the payload field.
func ByPayload(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPayload, opts...).ToFunc()
}
