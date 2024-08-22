// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/naelcodes/ab-backend/ent/imputation"
	"github.com/naelcodes/ab-backend/ent/invoice"
	"github.com/naelcodes/ab-backend/ent/payment"
	"github.com/naelcodes/ab-backend/ent/predicate"
)

// ImputationUpdate is the builder for updating Imputation entities.
type ImputationUpdate struct {
	config
	hooks    []Hook
	mutation *ImputationMutation
}

// Where appends a list predicates to the ImputationUpdate builder.
func (iu *ImputationUpdate) Where(ps ...predicate.Imputation) *ImputationUpdate {
	iu.mutation.Where(ps...)
	return iu
}

// SetAmountApply sets the "amount_apply" field.
func (iu *ImputationUpdate) SetAmountApply(f float64) *ImputationUpdate {
	iu.mutation.ResetAmountApply()
	iu.mutation.SetAmountApply(f)
	return iu
}

// SetNillableAmountApply sets the "amount_apply" field if the given value is not nil.
func (iu *ImputationUpdate) SetNillableAmountApply(f *float64) *ImputationUpdate {
	if f != nil {
		iu.SetAmountApply(*f)
	}
	return iu
}

// AddAmountApply adds f to the "amount_apply" field.
func (iu *ImputationUpdate) AddAmountApply(f float64) *ImputationUpdate {
	iu.mutation.AddAmountApply(f)
	return iu
}

// SetInvoiceAmount sets the "invoice_amount" field.
func (iu *ImputationUpdate) SetInvoiceAmount(f float64) *ImputationUpdate {
	iu.mutation.ResetInvoiceAmount()
	iu.mutation.SetInvoiceAmount(f)
	return iu
}

// SetNillableInvoiceAmount sets the "invoice_amount" field if the given value is not nil.
func (iu *ImputationUpdate) SetNillableInvoiceAmount(f *float64) *ImputationUpdate {
	if f != nil {
		iu.SetInvoiceAmount(*f)
	}
	return iu
}

// AddInvoiceAmount adds f to the "invoice_amount" field.
func (iu *ImputationUpdate) AddInvoiceAmount(f float64) *ImputationUpdate {
	iu.mutation.AddInvoiceAmount(f)
	return iu
}

// SetPaymentAmount sets the "payment_amount" field.
func (iu *ImputationUpdate) SetPaymentAmount(f float64) *ImputationUpdate {
	iu.mutation.ResetPaymentAmount()
	iu.mutation.SetPaymentAmount(f)
	return iu
}

// SetNillablePaymentAmount sets the "payment_amount" field if the given value is not nil.
func (iu *ImputationUpdate) SetNillablePaymentAmount(f *float64) *ImputationUpdate {
	if f != nil {
		iu.SetPaymentAmount(*f)
	}
	return iu
}

// AddPaymentAmount adds f to the "payment_amount" field.
func (iu *ImputationUpdate) AddPaymentAmount(f float64) *ImputationUpdate {
	iu.mutation.AddPaymentAmount(f)
	return iu
}

// SetTag sets the "tag" field.
func (iu *ImputationUpdate) SetTag(i imputation.Tag) *ImputationUpdate {
	iu.mutation.SetTag(i)
	return iu
}

// SetNillableTag sets the "tag" field if the given value is not nil.
func (iu *ImputationUpdate) SetNillableTag(i *imputation.Tag) *ImputationUpdate {
	if i != nil {
		iu.SetTag(*i)
	}
	return iu
}

// SetInvoiceID sets the "invoice" edge to the Invoice entity by ID.
func (iu *ImputationUpdate) SetInvoiceID(id int) *ImputationUpdate {
	iu.mutation.SetInvoiceID(id)
	return iu
}

// SetInvoice sets the "invoice" edge to the Invoice entity.
func (iu *ImputationUpdate) SetInvoice(i *Invoice) *ImputationUpdate {
	return iu.SetInvoiceID(i.ID)
}

// SetPaymentID sets the "payment" edge to the Payment entity by ID.
func (iu *ImputationUpdate) SetPaymentID(id int) *ImputationUpdate {
	iu.mutation.SetPaymentID(id)
	return iu
}

// SetPayment sets the "payment" edge to the Payment entity.
func (iu *ImputationUpdate) SetPayment(p *Payment) *ImputationUpdate {
	return iu.SetPaymentID(p.ID)
}

// Mutation returns the ImputationMutation object of the builder.
func (iu *ImputationUpdate) Mutation() *ImputationMutation {
	return iu.mutation
}

// ClearInvoice clears the "invoice" edge to the Invoice entity.
func (iu *ImputationUpdate) ClearInvoice() *ImputationUpdate {
	iu.mutation.ClearInvoice()
	return iu
}

// ClearPayment clears the "payment" edge to the Payment entity.
func (iu *ImputationUpdate) ClearPayment() *ImputationUpdate {
	iu.mutation.ClearPayment()
	return iu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (iu *ImputationUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, iu.sqlSave, iu.mutation, iu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (iu *ImputationUpdate) SaveX(ctx context.Context) int {
	affected, err := iu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (iu *ImputationUpdate) Exec(ctx context.Context) error {
	_, err := iu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iu *ImputationUpdate) ExecX(ctx context.Context) {
	if err := iu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iu *ImputationUpdate) check() error {
	if v, ok := iu.mutation.AmountApply(); ok {
		if err := imputation.AmountApplyValidator(v); err != nil {
			return &ValidationError{Name: "amount_apply", err: fmt.Errorf(`ent: validator failed for field "Imputation.amount_apply": %w`, err)}
		}
	}
	if v, ok := iu.mutation.InvoiceAmount(); ok {
		if err := imputation.InvoiceAmountValidator(v); err != nil {
			return &ValidationError{Name: "invoice_amount", err: fmt.Errorf(`ent: validator failed for field "Imputation.invoice_amount": %w`, err)}
		}
	}
	if v, ok := iu.mutation.PaymentAmount(); ok {
		if err := imputation.PaymentAmountValidator(v); err != nil {
			return &ValidationError{Name: "payment_amount", err: fmt.Errorf(`ent: validator failed for field "Imputation.payment_amount": %w`, err)}
		}
	}
	if v, ok := iu.mutation.Tag(); ok {
		if err := imputation.TagValidator(v); err != nil {
			return &ValidationError{Name: "tag", err: fmt.Errorf(`ent: validator failed for field "Imputation.tag": %w`, err)}
		}
	}
	if _, ok := iu.mutation.InvoiceID(); iu.mutation.InvoiceCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Imputation.invoice"`)
	}
	if _, ok := iu.mutation.PaymentID(); iu.mutation.PaymentCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Imputation.payment"`)
	}
	return nil
}

func (iu *ImputationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := iu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(imputation.Table, imputation.Columns, sqlgraph.NewFieldSpec(imputation.FieldID, field.TypeInt))
	if ps := iu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iu.mutation.AmountApply(); ok {
		vv, err := imputation.ValueScanner.AmountApply.Value(value)
		if err != nil {
			return 0, err
		}
		_spec.SetField(imputation.FieldAmountApply, field.TypeFloat64, vv)
	}
	if value, ok := iu.mutation.AddedAmountApply(); ok {
		vv, err := imputation.ValueScanner.AmountApply.Value(value)
		if err != nil {
			return 0, err
		}
		_spec.AddField(imputation.FieldAmountApply, field.TypeFloat64, vv)
	}
	if value, ok := iu.mutation.InvoiceAmount(); ok {
		vv, err := imputation.ValueScanner.InvoiceAmount.Value(value)
		if err != nil {
			return 0, err
		}
		_spec.SetField(imputation.FieldInvoiceAmount, field.TypeFloat64, vv)
	}
	if value, ok := iu.mutation.AddedInvoiceAmount(); ok {
		vv, err := imputation.ValueScanner.InvoiceAmount.Value(value)
		if err != nil {
			return 0, err
		}
		_spec.AddField(imputation.FieldInvoiceAmount, field.TypeFloat64, vv)
	}
	if value, ok := iu.mutation.PaymentAmount(); ok {
		vv, err := imputation.ValueScanner.PaymentAmount.Value(value)
		if err != nil {
			return 0, err
		}
		_spec.SetField(imputation.FieldPaymentAmount, field.TypeFloat64, vv)
	}
	if value, ok := iu.mutation.AddedPaymentAmount(); ok {
		vv, err := imputation.ValueScanner.PaymentAmount.Value(value)
		if err != nil {
			return 0, err
		}
		_spec.AddField(imputation.FieldPaymentAmount, field.TypeFloat64, vv)
	}
	if value, ok := iu.mutation.Tag(); ok {
		_spec.SetField(imputation.FieldTag, field.TypeEnum, value)
	}
	if iu.mutation.InvoiceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   imputation.InvoiceTable,
			Columns: []string{imputation.InvoiceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(invoice.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.InvoiceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   imputation.InvoiceTable,
			Columns: []string{imputation.InvoiceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(invoice.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if iu.mutation.PaymentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   imputation.PaymentTable,
			Columns: []string{imputation.PaymentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(payment.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.PaymentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   imputation.PaymentTable,
			Columns: []string{imputation.PaymentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(payment.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, iu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{imputation.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	iu.mutation.done = true
	return n, nil
}

// ImputationUpdateOne is the builder for updating a single Imputation entity.
type ImputationUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ImputationMutation
}

// SetAmountApply sets the "amount_apply" field.
func (iuo *ImputationUpdateOne) SetAmountApply(f float64) *ImputationUpdateOne {
	iuo.mutation.ResetAmountApply()
	iuo.mutation.SetAmountApply(f)
	return iuo
}

// SetNillableAmountApply sets the "amount_apply" field if the given value is not nil.
func (iuo *ImputationUpdateOne) SetNillableAmountApply(f *float64) *ImputationUpdateOne {
	if f != nil {
		iuo.SetAmountApply(*f)
	}
	return iuo
}

// AddAmountApply adds f to the "amount_apply" field.
func (iuo *ImputationUpdateOne) AddAmountApply(f float64) *ImputationUpdateOne {
	iuo.mutation.AddAmountApply(f)
	return iuo
}

// SetInvoiceAmount sets the "invoice_amount" field.
func (iuo *ImputationUpdateOne) SetInvoiceAmount(f float64) *ImputationUpdateOne {
	iuo.mutation.ResetInvoiceAmount()
	iuo.mutation.SetInvoiceAmount(f)
	return iuo
}

// SetNillableInvoiceAmount sets the "invoice_amount" field if the given value is not nil.
func (iuo *ImputationUpdateOne) SetNillableInvoiceAmount(f *float64) *ImputationUpdateOne {
	if f != nil {
		iuo.SetInvoiceAmount(*f)
	}
	return iuo
}

// AddInvoiceAmount adds f to the "invoice_amount" field.
func (iuo *ImputationUpdateOne) AddInvoiceAmount(f float64) *ImputationUpdateOne {
	iuo.mutation.AddInvoiceAmount(f)
	return iuo
}

// SetPaymentAmount sets the "payment_amount" field.
func (iuo *ImputationUpdateOne) SetPaymentAmount(f float64) *ImputationUpdateOne {
	iuo.mutation.ResetPaymentAmount()
	iuo.mutation.SetPaymentAmount(f)
	return iuo
}

// SetNillablePaymentAmount sets the "payment_amount" field if the given value is not nil.
func (iuo *ImputationUpdateOne) SetNillablePaymentAmount(f *float64) *ImputationUpdateOne {
	if f != nil {
		iuo.SetPaymentAmount(*f)
	}
	return iuo
}

// AddPaymentAmount adds f to the "payment_amount" field.
func (iuo *ImputationUpdateOne) AddPaymentAmount(f float64) *ImputationUpdateOne {
	iuo.mutation.AddPaymentAmount(f)
	return iuo
}

// SetTag sets the "tag" field.
func (iuo *ImputationUpdateOne) SetTag(i imputation.Tag) *ImputationUpdateOne {
	iuo.mutation.SetTag(i)
	return iuo
}

// SetNillableTag sets the "tag" field if the given value is not nil.
func (iuo *ImputationUpdateOne) SetNillableTag(i *imputation.Tag) *ImputationUpdateOne {
	if i != nil {
		iuo.SetTag(*i)
	}
	return iuo
}

// SetInvoiceID sets the "invoice" edge to the Invoice entity by ID.
func (iuo *ImputationUpdateOne) SetInvoiceID(id int) *ImputationUpdateOne {
	iuo.mutation.SetInvoiceID(id)
	return iuo
}

// SetInvoice sets the "invoice" edge to the Invoice entity.
func (iuo *ImputationUpdateOne) SetInvoice(i *Invoice) *ImputationUpdateOne {
	return iuo.SetInvoiceID(i.ID)
}

// SetPaymentID sets the "payment" edge to the Payment entity by ID.
func (iuo *ImputationUpdateOne) SetPaymentID(id int) *ImputationUpdateOne {
	iuo.mutation.SetPaymentID(id)
	return iuo
}

// SetPayment sets the "payment" edge to the Payment entity.
func (iuo *ImputationUpdateOne) SetPayment(p *Payment) *ImputationUpdateOne {
	return iuo.SetPaymentID(p.ID)
}

// Mutation returns the ImputationMutation object of the builder.
func (iuo *ImputationUpdateOne) Mutation() *ImputationMutation {
	return iuo.mutation
}

// ClearInvoice clears the "invoice" edge to the Invoice entity.
func (iuo *ImputationUpdateOne) ClearInvoice() *ImputationUpdateOne {
	iuo.mutation.ClearInvoice()
	return iuo
}

// ClearPayment clears the "payment" edge to the Payment entity.
func (iuo *ImputationUpdateOne) ClearPayment() *ImputationUpdateOne {
	iuo.mutation.ClearPayment()
	return iuo
}

// Where appends a list predicates to the ImputationUpdate builder.
func (iuo *ImputationUpdateOne) Where(ps ...predicate.Imputation) *ImputationUpdateOne {
	iuo.mutation.Where(ps...)
	return iuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (iuo *ImputationUpdateOne) Select(field string, fields ...string) *ImputationUpdateOne {
	iuo.fields = append([]string{field}, fields...)
	return iuo
}

// Save executes the query and returns the updated Imputation entity.
func (iuo *ImputationUpdateOne) Save(ctx context.Context) (*Imputation, error) {
	return withHooks(ctx, iuo.sqlSave, iuo.mutation, iuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (iuo *ImputationUpdateOne) SaveX(ctx context.Context) *Imputation {
	node, err := iuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (iuo *ImputationUpdateOne) Exec(ctx context.Context) error {
	_, err := iuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iuo *ImputationUpdateOne) ExecX(ctx context.Context) {
	if err := iuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iuo *ImputationUpdateOne) check() error {
	if v, ok := iuo.mutation.AmountApply(); ok {
		if err := imputation.AmountApplyValidator(v); err != nil {
			return &ValidationError{Name: "amount_apply", err: fmt.Errorf(`ent: validator failed for field "Imputation.amount_apply": %w`, err)}
		}
	}
	if v, ok := iuo.mutation.InvoiceAmount(); ok {
		if err := imputation.InvoiceAmountValidator(v); err != nil {
			return &ValidationError{Name: "invoice_amount", err: fmt.Errorf(`ent: validator failed for field "Imputation.invoice_amount": %w`, err)}
		}
	}
	if v, ok := iuo.mutation.PaymentAmount(); ok {
		if err := imputation.PaymentAmountValidator(v); err != nil {
			return &ValidationError{Name: "payment_amount", err: fmt.Errorf(`ent: validator failed for field "Imputation.payment_amount": %w`, err)}
		}
	}
	if v, ok := iuo.mutation.Tag(); ok {
		if err := imputation.TagValidator(v); err != nil {
			return &ValidationError{Name: "tag", err: fmt.Errorf(`ent: validator failed for field "Imputation.tag": %w`, err)}
		}
	}
	if _, ok := iuo.mutation.InvoiceID(); iuo.mutation.InvoiceCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Imputation.invoice"`)
	}
	if _, ok := iuo.mutation.PaymentID(); iuo.mutation.PaymentCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Imputation.payment"`)
	}
	return nil
}

func (iuo *ImputationUpdateOne) sqlSave(ctx context.Context) (_node *Imputation, err error) {
	if err := iuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(imputation.Table, imputation.Columns, sqlgraph.NewFieldSpec(imputation.FieldID, field.TypeInt))
	id, ok := iuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Imputation.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := iuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, imputation.FieldID)
		for _, f := range fields {
			if !imputation.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != imputation.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := iuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iuo.mutation.AmountApply(); ok {
		vv, err := imputation.ValueScanner.AmountApply.Value(value)
		if err != nil {
			return nil, err
		}
		_spec.SetField(imputation.FieldAmountApply, field.TypeFloat64, vv)
	}
	if value, ok := iuo.mutation.AddedAmountApply(); ok {
		vv, err := imputation.ValueScanner.AmountApply.Value(value)
		if err != nil {
			return nil, err
		}
		_spec.AddField(imputation.FieldAmountApply, field.TypeFloat64, vv)
	}
	if value, ok := iuo.mutation.InvoiceAmount(); ok {
		vv, err := imputation.ValueScanner.InvoiceAmount.Value(value)
		if err != nil {
			return nil, err
		}
		_spec.SetField(imputation.FieldInvoiceAmount, field.TypeFloat64, vv)
	}
	if value, ok := iuo.mutation.AddedInvoiceAmount(); ok {
		vv, err := imputation.ValueScanner.InvoiceAmount.Value(value)
		if err != nil {
			return nil, err
		}
		_spec.AddField(imputation.FieldInvoiceAmount, field.TypeFloat64, vv)
	}
	if value, ok := iuo.mutation.PaymentAmount(); ok {
		vv, err := imputation.ValueScanner.PaymentAmount.Value(value)
		if err != nil {
			return nil, err
		}
		_spec.SetField(imputation.FieldPaymentAmount, field.TypeFloat64, vv)
	}
	if value, ok := iuo.mutation.AddedPaymentAmount(); ok {
		vv, err := imputation.ValueScanner.PaymentAmount.Value(value)
		if err != nil {
			return nil, err
		}
		_spec.AddField(imputation.FieldPaymentAmount, field.TypeFloat64, vv)
	}
	if value, ok := iuo.mutation.Tag(); ok {
		_spec.SetField(imputation.FieldTag, field.TypeEnum, value)
	}
	if iuo.mutation.InvoiceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   imputation.InvoiceTable,
			Columns: []string{imputation.InvoiceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(invoice.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.InvoiceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   imputation.InvoiceTable,
			Columns: []string{imputation.InvoiceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(invoice.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if iuo.mutation.PaymentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   imputation.PaymentTable,
			Columns: []string{imputation.PaymentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(payment.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.PaymentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   imputation.PaymentTable,
			Columns: []string{imputation.PaymentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(payment.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Imputation{config: iuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, iuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{imputation.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	iuo.mutation.done = true
	return _node, nil
}
