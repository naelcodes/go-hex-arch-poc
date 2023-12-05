// Code generated by ent, DO NOT EDIT.

package travelitem

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/naelcodes/ab-backend/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldLTE(FieldID, id))
}

// TotalPrice applies equality check predicate on the "total_price" field. It's identical to TotalPriceEQ.
func TotalPrice(v float64) predicate.TravelItem {
	vc, err := ValueScanner.TotalPrice.Value(v)
	return predicate.TravelItemOrErr(sql.FieldEQ(FieldTotalPrice, vc), err)
}

// Itinerary applies equality check predicate on the "itinerary" field. It's identical to ItineraryEQ.
func Itinerary(v string) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldEQ(FieldItinerary, v))
}

// TravelerName applies equality check predicate on the "traveler_name" field. It's identical to TravelerNameEQ.
func TravelerName(v string) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldEQ(FieldTravelerName, v))
}

// TicketNumber applies equality check predicate on the "ticket_number" field. It's identical to TicketNumberEQ.
func TicketNumber(v string) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldEQ(FieldTicketNumber, v))
}

// ConjunctionNumber applies equality check predicate on the "conjunction_number" field. It's identical to ConjunctionNumberEQ.
func ConjunctionNumber(v int) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldEQ(FieldConjunctionNumber, v))
}

// TransactionType applies equality check predicate on the "transaction_type" field. It's identical to TransactionTypeEQ.
func TransactionType(v string) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldEQ(FieldTransactionType, v))
}

// ProductType applies equality check predicate on the "product_type" field. It's identical to ProductTypeEQ.
func ProductType(v string) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldEQ(FieldProductType, v))
}

// TotalPriceEQ applies the EQ predicate on the "total_price" field.
func TotalPriceEQ(v float64) predicate.TravelItem {
	vc, err := ValueScanner.TotalPrice.Value(v)
	return predicate.TravelItemOrErr(sql.FieldEQ(FieldTotalPrice, vc), err)
}

// TotalPriceNEQ applies the NEQ predicate on the "total_price" field.
func TotalPriceNEQ(v float64) predicate.TravelItem {
	vc, err := ValueScanner.TotalPrice.Value(v)
	return predicate.TravelItemOrErr(sql.FieldNEQ(FieldTotalPrice, vc), err)
}

// TotalPriceIn applies the In predicate on the "total_price" field.
func TotalPriceIn(vs ...float64) predicate.TravelItem {
	var (
		err error
		v   = make([]any, len(vs))
	)
	for i := range v {
		if v[i], err = ValueScanner.TotalPrice.Value(vs[i]); err != nil {
			break
		}
	}
	return predicate.TravelItemOrErr(sql.FieldIn(FieldTotalPrice, v...), err)
}

// TotalPriceNotIn applies the NotIn predicate on the "total_price" field.
func TotalPriceNotIn(vs ...float64) predicate.TravelItem {
	var (
		err error
		v   = make([]any, len(vs))
	)
	for i := range v {
		if v[i], err = ValueScanner.TotalPrice.Value(vs[i]); err != nil {
			break
		}
	}
	return predicate.TravelItemOrErr(sql.FieldNotIn(FieldTotalPrice, v...), err)
}

// TotalPriceGT applies the GT predicate on the "total_price" field.
func TotalPriceGT(v float64) predicate.TravelItem {
	vc, err := ValueScanner.TotalPrice.Value(v)
	return predicate.TravelItemOrErr(sql.FieldGT(FieldTotalPrice, vc), err)
}

// TotalPriceGTE applies the GTE predicate on the "total_price" field.
func TotalPriceGTE(v float64) predicate.TravelItem {
	vc, err := ValueScanner.TotalPrice.Value(v)
	return predicate.TravelItemOrErr(sql.FieldGTE(FieldTotalPrice, vc), err)
}

// TotalPriceLT applies the LT predicate on the "total_price" field.
func TotalPriceLT(v float64) predicate.TravelItem {
	vc, err := ValueScanner.TotalPrice.Value(v)
	return predicate.TravelItemOrErr(sql.FieldLT(FieldTotalPrice, vc), err)
}

// TotalPriceLTE applies the LTE predicate on the "total_price" field.
func TotalPriceLTE(v float64) predicate.TravelItem {
	vc, err := ValueScanner.TotalPrice.Value(v)
	return predicate.TravelItemOrErr(sql.FieldLTE(FieldTotalPrice, vc), err)
}

// ItineraryEQ applies the EQ predicate on the "itinerary" field.
func ItineraryEQ(v string) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldEQ(FieldItinerary, v))
}

// ItineraryNEQ applies the NEQ predicate on the "itinerary" field.
func ItineraryNEQ(v string) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldNEQ(FieldItinerary, v))
}

// ItineraryIn applies the In predicate on the "itinerary" field.
func ItineraryIn(vs ...string) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldIn(FieldItinerary, vs...))
}

// ItineraryNotIn applies the NotIn predicate on the "itinerary" field.
func ItineraryNotIn(vs ...string) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldNotIn(FieldItinerary, vs...))
}

// ItineraryGT applies the GT predicate on the "itinerary" field.
func ItineraryGT(v string) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldGT(FieldItinerary, v))
}

// ItineraryGTE applies the GTE predicate on the "itinerary" field.
func ItineraryGTE(v string) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldGTE(FieldItinerary, v))
}

// ItineraryLT applies the LT predicate on the "itinerary" field.
func ItineraryLT(v string) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldLT(FieldItinerary, v))
}

// ItineraryLTE applies the LTE predicate on the "itinerary" field.
func ItineraryLTE(v string) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldLTE(FieldItinerary, v))
}

// ItineraryContains applies the Contains predicate on the "itinerary" field.
func ItineraryContains(v string) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldContains(FieldItinerary, v))
}

// ItineraryHasPrefix applies the HasPrefix predicate on the "itinerary" field.
func ItineraryHasPrefix(v string) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldHasPrefix(FieldItinerary, v))
}

// ItineraryHasSuffix applies the HasSuffix predicate on the "itinerary" field.
func ItineraryHasSuffix(v string) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldHasSuffix(FieldItinerary, v))
}

// ItineraryEqualFold applies the EqualFold predicate on the "itinerary" field.
func ItineraryEqualFold(v string) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldEqualFold(FieldItinerary, v))
}

// ItineraryContainsFold applies the ContainsFold predicate on the "itinerary" field.
func ItineraryContainsFold(v string) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldContainsFold(FieldItinerary, v))
}

// TravelerNameEQ applies the EQ predicate on the "traveler_name" field.
func TravelerNameEQ(v string) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldEQ(FieldTravelerName, v))
}

// TravelerNameNEQ applies the NEQ predicate on the "traveler_name" field.
func TravelerNameNEQ(v string) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldNEQ(FieldTravelerName, v))
}

// TravelerNameIn applies the In predicate on the "traveler_name" field.
func TravelerNameIn(vs ...string) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldIn(FieldTravelerName, vs...))
}

// TravelerNameNotIn applies the NotIn predicate on the "traveler_name" field.
func TravelerNameNotIn(vs ...string) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldNotIn(FieldTravelerName, vs...))
}

// TravelerNameGT applies the GT predicate on the "traveler_name" field.
func TravelerNameGT(v string) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldGT(FieldTravelerName, v))
}

// TravelerNameGTE applies the GTE predicate on the "traveler_name" field.
func TravelerNameGTE(v string) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldGTE(FieldTravelerName, v))
}

// TravelerNameLT applies the LT predicate on the "traveler_name" field.
func TravelerNameLT(v string) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldLT(FieldTravelerName, v))
}

// TravelerNameLTE applies the LTE predicate on the "traveler_name" field.
func TravelerNameLTE(v string) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldLTE(FieldTravelerName, v))
}

// TravelerNameContains applies the Contains predicate on the "traveler_name" field.
func TravelerNameContains(v string) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldContains(FieldTravelerName, v))
}

// TravelerNameHasPrefix applies the HasPrefix predicate on the "traveler_name" field.
func TravelerNameHasPrefix(v string) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldHasPrefix(FieldTravelerName, v))
}

// TravelerNameHasSuffix applies the HasSuffix predicate on the "traveler_name" field.
func TravelerNameHasSuffix(v string) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldHasSuffix(FieldTravelerName, v))
}

// TravelerNameEqualFold applies the EqualFold predicate on the "traveler_name" field.
func TravelerNameEqualFold(v string) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldEqualFold(FieldTravelerName, v))
}

// TravelerNameContainsFold applies the ContainsFold predicate on the "traveler_name" field.
func TravelerNameContainsFold(v string) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldContainsFold(FieldTravelerName, v))
}

// TicketNumberEQ applies the EQ predicate on the "ticket_number" field.
func TicketNumberEQ(v string) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldEQ(FieldTicketNumber, v))
}

// TicketNumberNEQ applies the NEQ predicate on the "ticket_number" field.
func TicketNumberNEQ(v string) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldNEQ(FieldTicketNumber, v))
}

// TicketNumberIn applies the In predicate on the "ticket_number" field.
func TicketNumberIn(vs ...string) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldIn(FieldTicketNumber, vs...))
}

// TicketNumberNotIn applies the NotIn predicate on the "ticket_number" field.
func TicketNumberNotIn(vs ...string) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldNotIn(FieldTicketNumber, vs...))
}

// TicketNumberGT applies the GT predicate on the "ticket_number" field.
func TicketNumberGT(v string) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldGT(FieldTicketNumber, v))
}

// TicketNumberGTE applies the GTE predicate on the "ticket_number" field.
func TicketNumberGTE(v string) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldGTE(FieldTicketNumber, v))
}

// TicketNumberLT applies the LT predicate on the "ticket_number" field.
func TicketNumberLT(v string) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldLT(FieldTicketNumber, v))
}

// TicketNumberLTE applies the LTE predicate on the "ticket_number" field.
func TicketNumberLTE(v string) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldLTE(FieldTicketNumber, v))
}

// TicketNumberContains applies the Contains predicate on the "ticket_number" field.
func TicketNumberContains(v string) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldContains(FieldTicketNumber, v))
}

// TicketNumberHasPrefix applies the HasPrefix predicate on the "ticket_number" field.
func TicketNumberHasPrefix(v string) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldHasPrefix(FieldTicketNumber, v))
}

// TicketNumberHasSuffix applies the HasSuffix predicate on the "ticket_number" field.
func TicketNumberHasSuffix(v string) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldHasSuffix(FieldTicketNumber, v))
}

// TicketNumberEqualFold applies the EqualFold predicate on the "ticket_number" field.
func TicketNumberEqualFold(v string) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldEqualFold(FieldTicketNumber, v))
}

// TicketNumberContainsFold applies the ContainsFold predicate on the "ticket_number" field.
func TicketNumberContainsFold(v string) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldContainsFold(FieldTicketNumber, v))
}

// ConjunctionNumberEQ applies the EQ predicate on the "conjunction_number" field.
func ConjunctionNumberEQ(v int) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldEQ(FieldConjunctionNumber, v))
}

// ConjunctionNumberNEQ applies the NEQ predicate on the "conjunction_number" field.
func ConjunctionNumberNEQ(v int) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldNEQ(FieldConjunctionNumber, v))
}

// ConjunctionNumberIn applies the In predicate on the "conjunction_number" field.
func ConjunctionNumberIn(vs ...int) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldIn(FieldConjunctionNumber, vs...))
}

// ConjunctionNumberNotIn applies the NotIn predicate on the "conjunction_number" field.
func ConjunctionNumberNotIn(vs ...int) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldNotIn(FieldConjunctionNumber, vs...))
}

// ConjunctionNumberGT applies the GT predicate on the "conjunction_number" field.
func ConjunctionNumberGT(v int) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldGT(FieldConjunctionNumber, v))
}

// ConjunctionNumberGTE applies the GTE predicate on the "conjunction_number" field.
func ConjunctionNumberGTE(v int) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldGTE(FieldConjunctionNumber, v))
}

// ConjunctionNumberLT applies the LT predicate on the "conjunction_number" field.
func ConjunctionNumberLT(v int) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldLT(FieldConjunctionNumber, v))
}

// ConjunctionNumberLTE applies the LTE predicate on the "conjunction_number" field.
func ConjunctionNumberLTE(v int) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldLTE(FieldConjunctionNumber, v))
}

// TransactionTypeEQ applies the EQ predicate on the "transaction_type" field.
func TransactionTypeEQ(v string) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldEQ(FieldTransactionType, v))
}

// TransactionTypeNEQ applies the NEQ predicate on the "transaction_type" field.
func TransactionTypeNEQ(v string) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldNEQ(FieldTransactionType, v))
}

// TransactionTypeIn applies the In predicate on the "transaction_type" field.
func TransactionTypeIn(vs ...string) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldIn(FieldTransactionType, vs...))
}

// TransactionTypeNotIn applies the NotIn predicate on the "transaction_type" field.
func TransactionTypeNotIn(vs ...string) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldNotIn(FieldTransactionType, vs...))
}

// TransactionTypeGT applies the GT predicate on the "transaction_type" field.
func TransactionTypeGT(v string) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldGT(FieldTransactionType, v))
}

// TransactionTypeGTE applies the GTE predicate on the "transaction_type" field.
func TransactionTypeGTE(v string) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldGTE(FieldTransactionType, v))
}

// TransactionTypeLT applies the LT predicate on the "transaction_type" field.
func TransactionTypeLT(v string) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldLT(FieldTransactionType, v))
}

// TransactionTypeLTE applies the LTE predicate on the "transaction_type" field.
func TransactionTypeLTE(v string) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldLTE(FieldTransactionType, v))
}

// TransactionTypeContains applies the Contains predicate on the "transaction_type" field.
func TransactionTypeContains(v string) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldContains(FieldTransactionType, v))
}

// TransactionTypeHasPrefix applies the HasPrefix predicate on the "transaction_type" field.
func TransactionTypeHasPrefix(v string) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldHasPrefix(FieldTransactionType, v))
}

// TransactionTypeHasSuffix applies the HasSuffix predicate on the "transaction_type" field.
func TransactionTypeHasSuffix(v string) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldHasSuffix(FieldTransactionType, v))
}

// TransactionTypeEqualFold applies the EqualFold predicate on the "transaction_type" field.
func TransactionTypeEqualFold(v string) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldEqualFold(FieldTransactionType, v))
}

// TransactionTypeContainsFold applies the ContainsFold predicate on the "transaction_type" field.
func TransactionTypeContainsFold(v string) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldContainsFold(FieldTransactionType, v))
}

// ProductTypeEQ applies the EQ predicate on the "product_type" field.
func ProductTypeEQ(v string) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldEQ(FieldProductType, v))
}

// ProductTypeNEQ applies the NEQ predicate on the "product_type" field.
func ProductTypeNEQ(v string) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldNEQ(FieldProductType, v))
}

// ProductTypeIn applies the In predicate on the "product_type" field.
func ProductTypeIn(vs ...string) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldIn(FieldProductType, vs...))
}

// ProductTypeNotIn applies the NotIn predicate on the "product_type" field.
func ProductTypeNotIn(vs ...string) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldNotIn(FieldProductType, vs...))
}

// ProductTypeGT applies the GT predicate on the "product_type" field.
func ProductTypeGT(v string) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldGT(FieldProductType, v))
}

// ProductTypeGTE applies the GTE predicate on the "product_type" field.
func ProductTypeGTE(v string) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldGTE(FieldProductType, v))
}

// ProductTypeLT applies the LT predicate on the "product_type" field.
func ProductTypeLT(v string) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldLT(FieldProductType, v))
}

// ProductTypeLTE applies the LTE predicate on the "product_type" field.
func ProductTypeLTE(v string) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldLTE(FieldProductType, v))
}

// ProductTypeContains applies the Contains predicate on the "product_type" field.
func ProductTypeContains(v string) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldContains(FieldProductType, v))
}

// ProductTypeHasPrefix applies the HasPrefix predicate on the "product_type" field.
func ProductTypeHasPrefix(v string) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldHasPrefix(FieldProductType, v))
}

// ProductTypeHasSuffix applies the HasSuffix predicate on the "product_type" field.
func ProductTypeHasSuffix(v string) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldHasSuffix(FieldProductType, v))
}

// ProductTypeEqualFold applies the EqualFold predicate on the "product_type" field.
func ProductTypeEqualFold(v string) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldEqualFold(FieldProductType, v))
}

// ProductTypeContainsFold applies the ContainsFold predicate on the "product_type" field.
func ProductTypeContainsFold(v string) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldContainsFold(FieldProductType, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.TravelItem {
	return predicate.TravelItem(sql.FieldNotIn(FieldStatus, vs...))
}

// HasInvoice applies the HasEdge predicate on the "invoice" edge.
func HasInvoice() predicate.TravelItem {
	return predicate.TravelItem(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, InvoiceTable, InvoiceColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasInvoiceWith applies the HasEdge predicate on the "invoice" edge with a given conditions (other predicates).
func HasInvoiceWith(preds ...predicate.Invoice) predicate.TravelItem {
	return predicate.TravelItem(func(s *sql.Selector) {
		step := newInvoiceStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.TravelItem) predicate.TravelItem {
	return predicate.TravelItem(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.TravelItem) predicate.TravelItem {
	return predicate.TravelItem(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.TravelItem) predicate.TravelItem {
	return predicate.TravelItem(sql.NotPredicates(p))
}
