// Code generated by ent, DO NOT EDIT.

package travelitem

import (
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

const (
	// Label holds the string label denoting the travelitem type in the database.
	Label = "travel_item"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTotalPrice holds the string denoting the total_price field in the database.
	FieldTotalPrice = "total_price"
	// FieldItinerary holds the string denoting the itinerary field in the database.
	FieldItinerary = "itinerary"
	// FieldTravelerName holds the string denoting the traveler_name field in the database.
	FieldTravelerName = "traveler_name"
	// FieldTicketNumber holds the string denoting the ticket_number field in the database.
	FieldTicketNumber = "ticket_number"
	// FieldConjunctionNumber holds the string denoting the conjunction_number field in the database.
	FieldConjunctionNumber = "conjunction_number"
	// FieldTransactionType holds the string denoting the transaction_type field in the database.
	FieldTransactionType = "transaction_type"
	// FieldProductType holds the string denoting the product_type field in the database.
	FieldProductType = "product_type"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// EdgeInvoice holds the string denoting the invoice edge name in mutations.
	EdgeInvoice = "invoice"
	// Table holds the table name of the travelitem in the database.
	Table = "air_booking"
	// InvoiceTable is the table that holds the invoice relation/edge.
	InvoiceTable = "air_booking"
	// InvoiceInverseTable is the table name for the Invoice entity.
	// It exists in this package in order to avoid circular dependency with the "invoice" package.
	InvoiceInverseTable = "invoice"
	// InvoiceColumn is the table column denoting the invoice relation/edge.
	InvoiceColumn = "id_invoice"
)

// Columns holds all SQL columns for travelitem fields.
var Columns = []string{
	FieldID,
	FieldTotalPrice,
	FieldItinerary,
	FieldTravelerName,
	FieldTicketNumber,
	FieldConjunctionNumber,
	FieldTransactionType,
	FieldProductType,
	FieldStatus,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "air_booking"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"id_invoice",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// ValueScanner of all TravelItem fields.
	ValueScanner struct {
		TotalPrice field.TypeValueScanner[float64]
	}
)

// Status defines the type for the "status" enum field.
type Status string

// StatusPending is the default value of the Status enum.
const DefaultStatus = StatusPending

// Status values.
const (
	StatusPending   Status = "pending"
	StatusInvoiced  Status = "invoiced"
	StatusVoid      Status = "void"
	StatusReceipted Status = "receipted"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusPending, StatusInvoiced, StatusVoid, StatusReceipted:
		return nil
	default:
		return fmt.Errorf("travelitem: invalid enum value for status field: %q", s)
	}
}

// OrderOption defines the ordering options for the TravelItem queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByTotalPrice orders the results by the total_price field.
func ByTotalPrice(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTotalPrice, opts...).ToFunc()
}

// ByItinerary orders the results by the itinerary field.
func ByItinerary(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldItinerary, opts...).ToFunc()
}

// ByTravelerName orders the results by the traveler_name field.
func ByTravelerName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTravelerName, opts...).ToFunc()
}

// ByTicketNumber orders the results by the ticket_number field.
func ByTicketNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTicketNumber, opts...).ToFunc()
}

// ByConjunctionNumber orders the results by the conjunction_number field.
func ByConjunctionNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldConjunctionNumber, opts...).ToFunc()
}

// ByTransactionType orders the results by the transaction_type field.
func ByTransactionType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTransactionType, opts...).ToFunc()
}

// ByProductType orders the results by the product_type field.
func ByProductType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProductType, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByInvoiceField orders the results by invoice field.
func ByInvoiceField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newInvoiceStep(), sql.OrderByField(field, opts...))
	}
}
func newInvoiceStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(InvoiceInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, InvoiceTable, InvoiceColumn),
	)
}
