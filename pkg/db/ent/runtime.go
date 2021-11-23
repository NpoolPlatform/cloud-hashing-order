// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/NpoolPlatform/cloud-hashing-order/pkg/db/ent/compensate"
	"github.com/NpoolPlatform/cloud-hashing-order/pkg/db/ent/gaspaying"
	"github.com/NpoolPlatform/cloud-hashing-order/pkg/db/ent/goodpaying"
	"github.com/NpoolPlatform/cloud-hashing-order/pkg/db/ent/order"
	"github.com/NpoolPlatform/cloud-hashing-order/pkg/db/ent/outofgas"
	"github.com/NpoolPlatform/cloud-hashing-order/pkg/db/ent/payment"
	"github.com/NpoolPlatform/cloud-hashing-order/pkg/db/ent/schema"
	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	compensateFields := schema.Compensate{}.Fields()
	_ = compensateFields
	// compensateDescCreateAt is the schema descriptor for create_at field.
	compensateDescCreateAt := compensateFields[5].Descriptor()
	// compensate.DefaultCreateAt holds the default value on creation for the create_at field.
	compensate.DefaultCreateAt = compensateDescCreateAt.Default.(func() uint32)
	// compensateDescUpdateAt is the schema descriptor for update_at field.
	compensateDescUpdateAt := compensateFields[6].Descriptor()
	// compensate.DefaultUpdateAt holds the default value on creation for the update_at field.
	compensate.DefaultUpdateAt = compensateDescUpdateAt.Default.(func() uint32)
	// compensate.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	compensate.UpdateDefaultUpdateAt = compensateDescUpdateAt.UpdateDefault.(func() uint32)
	// compensateDescDeleteAt is the schema descriptor for delete_at field.
	compensateDescDeleteAt := compensateFields[7].Descriptor()
	// compensate.DefaultDeleteAt holds the default value on creation for the delete_at field.
	compensate.DefaultDeleteAt = compensateDescDeleteAt.Default.(func() uint32)
	// compensateDescID is the schema descriptor for id field.
	compensateDescID := compensateFields[0].Descriptor()
	// compensate.DefaultID holds the default value on creation for the id field.
	compensate.DefaultID = compensateDescID.Default.(func() uuid.UUID)
	gaspayingFields := schema.GasPaying{}.Fields()
	_ = gaspayingFields
	// gaspayingDescCreateAt is the schema descriptor for create_at field.
	gaspayingDescCreateAt := gaspayingFields[5].Descriptor()
	// gaspaying.DefaultCreateAt holds the default value on creation for the create_at field.
	gaspaying.DefaultCreateAt = gaspayingDescCreateAt.Default.(func() uint32)
	// gaspayingDescUpdateAt is the schema descriptor for update_at field.
	gaspayingDescUpdateAt := gaspayingFields[6].Descriptor()
	// gaspaying.DefaultUpdateAt holds the default value on creation for the update_at field.
	gaspaying.DefaultUpdateAt = gaspayingDescUpdateAt.Default.(func() uint32)
	// gaspaying.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	gaspaying.UpdateDefaultUpdateAt = gaspayingDescUpdateAt.UpdateDefault.(func() uint32)
	// gaspayingDescDeleteAt is the schema descriptor for delete_at field.
	gaspayingDescDeleteAt := gaspayingFields[7].Descriptor()
	// gaspaying.DefaultDeleteAt holds the default value on creation for the delete_at field.
	gaspaying.DefaultDeleteAt = gaspayingDescDeleteAt.Default.(func() uint32)
	// gaspayingDescID is the schema descriptor for id field.
	gaspayingDescID := gaspayingFields[0].Descriptor()
	// gaspaying.DefaultID holds the default value on creation for the id field.
	gaspaying.DefaultID = gaspayingDescID.Default.(func() uuid.UUID)
	goodpayingFields := schema.GoodPaying{}.Fields()
	_ = goodpayingFields
	// goodpayingDescCreateAt is the schema descriptor for create_at field.
	goodpayingDescCreateAt := goodpayingFields[3].Descriptor()
	// goodpaying.DefaultCreateAt holds the default value on creation for the create_at field.
	goodpaying.DefaultCreateAt = goodpayingDescCreateAt.Default.(func() uint32)
	// goodpayingDescUpdateAt is the schema descriptor for update_at field.
	goodpayingDescUpdateAt := goodpayingFields[4].Descriptor()
	// goodpaying.DefaultUpdateAt holds the default value on creation for the update_at field.
	goodpaying.DefaultUpdateAt = goodpayingDescUpdateAt.Default.(func() uint32)
	// goodpaying.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	goodpaying.UpdateDefaultUpdateAt = goodpayingDescUpdateAt.UpdateDefault.(func() uint32)
	// goodpayingDescDeleteAt is the schema descriptor for delete_at field.
	goodpayingDescDeleteAt := goodpayingFields[5].Descriptor()
	// goodpaying.DefaultDeleteAt holds the default value on creation for the delete_at field.
	goodpaying.DefaultDeleteAt = goodpayingDescDeleteAt.Default.(func() uint32)
	// goodpayingDescID is the schema descriptor for id field.
	goodpayingDescID := goodpayingFields[0].Descriptor()
	// goodpaying.DefaultID holds the default value on creation for the id field.
	goodpaying.DefaultID = goodpayingDescID.Default.(func() uuid.UUID)
	orderFields := schema.Order{}.Fields()
	_ = orderFields
	// orderDescDiscount is the schema descriptor for discount field.
	orderDescDiscount := orderFields[5].Descriptor()
	// order.DefaultDiscount holds the default value on creation for the discount field.
	order.DefaultDiscount = orderDescDiscount.Default.(uint32)
	// orderDescSpecialReductionAmount is the schema descriptor for special_reduction_amount field.
	orderDescSpecialReductionAmount := orderFields[6].Descriptor()
	// order.DefaultSpecialReductionAmount holds the default value on creation for the special_reduction_amount field.
	order.DefaultSpecialReductionAmount = orderDescSpecialReductionAmount.Default.(uint64)
	// orderDescCreateAt is the schema descriptor for create_at field.
	orderDescCreateAt := orderFields[10].Descriptor()
	// order.DefaultCreateAt holds the default value on creation for the create_at field.
	order.DefaultCreateAt = orderDescCreateAt.Default.(func() uint32)
	// orderDescUpdateAt is the schema descriptor for update_at field.
	orderDescUpdateAt := orderFields[11].Descriptor()
	// order.DefaultUpdateAt holds the default value on creation for the update_at field.
	order.DefaultUpdateAt = orderDescUpdateAt.Default.(func() uint32)
	// order.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	order.UpdateDefaultUpdateAt = orderDescUpdateAt.UpdateDefault.(func() uint32)
	// orderDescDeleteAt is the schema descriptor for delete_at field.
	orderDescDeleteAt := orderFields[12].Descriptor()
	// order.DefaultDeleteAt holds the default value on creation for the delete_at field.
	order.DefaultDeleteAt = orderDescDeleteAt.Default.(func() uint32)
	// orderDescID is the schema descriptor for id field.
	orderDescID := orderFields[0].Descriptor()
	// order.DefaultID holds the default value on creation for the id field.
	order.DefaultID = orderDescID.Default.(func() uuid.UUID)
	outofgasFields := schema.OutOfGas{}.Fields()
	_ = outofgasFields
	// outofgasDescCreateAt is the schema descriptor for create_at field.
	outofgasDescCreateAt := outofgasFields[4].Descriptor()
	// outofgas.DefaultCreateAt holds the default value on creation for the create_at field.
	outofgas.DefaultCreateAt = outofgasDescCreateAt.Default.(func() uint32)
	// outofgasDescUpdateAt is the schema descriptor for update_at field.
	outofgasDescUpdateAt := outofgasFields[5].Descriptor()
	// outofgas.DefaultUpdateAt holds the default value on creation for the update_at field.
	outofgas.DefaultUpdateAt = outofgasDescUpdateAt.Default.(func() uint32)
	// outofgas.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	outofgas.UpdateDefaultUpdateAt = outofgasDescUpdateAt.UpdateDefault.(func() uint32)
	// outofgasDescDeleteAt is the schema descriptor for delete_at field.
	outofgasDescDeleteAt := outofgasFields[6].Descriptor()
	// outofgas.DefaultDeleteAt holds the default value on creation for the delete_at field.
	outofgas.DefaultDeleteAt = outofgasDescDeleteAt.Default.(func() uint32)
	// outofgasDescID is the schema descriptor for id field.
	outofgasDescID := outofgasFields[0].Descriptor()
	// outofgas.DefaultID holds the default value on creation for the id field.
	outofgas.DefaultID = outofgasDescID.Default.(func() uuid.UUID)
	paymentFields := schema.Payment{}.Fields()
	_ = paymentFields
	// paymentDescCreateAt is the schema descriptor for create_at field.
	paymentDescCreateAt := paymentFields[8].Descriptor()
	// payment.DefaultCreateAt holds the default value on creation for the create_at field.
	payment.DefaultCreateAt = paymentDescCreateAt.Default.(func() uint32)
	// paymentDescUpdateAt is the schema descriptor for update_at field.
	paymentDescUpdateAt := paymentFields[9].Descriptor()
	// payment.DefaultUpdateAt holds the default value on creation for the update_at field.
	payment.DefaultUpdateAt = paymentDescUpdateAt.Default.(func() uint32)
	// payment.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	payment.UpdateDefaultUpdateAt = paymentDescUpdateAt.UpdateDefault.(func() uint32)
	// paymentDescDeleteAt is the schema descriptor for delete_at field.
	paymentDescDeleteAt := paymentFields[10].Descriptor()
	// payment.DefaultDeleteAt holds the default value on creation for the delete_at field.
	payment.DefaultDeleteAt = paymentDescDeleteAt.Default.(func() uint32)
	// paymentDescID is the schema descriptor for id field.
	paymentDescID := paymentFields[0].Descriptor()
	// payment.DefaultID holds the default value on creation for the id field.
	payment.DefaultID = paymentDescID.Default.(func() uuid.UUID)
}
