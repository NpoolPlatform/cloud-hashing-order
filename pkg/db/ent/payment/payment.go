// Code generated by entc, DO NOT EDIT.

package payment

import (
	"fmt"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the payment type in the database.
	Label = "payment"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAppID holds the string denoting the app_id field in the database.
	FieldAppID = "app_id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldGoodID holds the string denoting the good_id field in the database.
	FieldGoodID = "good_id"
	// FieldOrderID holds the string denoting the order_id field in the database.
	FieldOrderID = "order_id"
	// FieldAccountID holds the string denoting the account_id field in the database.
	FieldAccountID = "account_id"
	// FieldStartAmount holds the string denoting the start_amount field in the database.
	FieldStartAmount = "start_amount"
	// FieldAmount holds the string denoting the amount field in the database.
	FieldAmount = "amount"
	// FieldFinishAmount holds the string denoting the finish_amount field in the database.
	FieldFinishAmount = "finish_amount"
	// FieldCoinUsdCurrency holds the string denoting the coin_usd_currency field in the database.
	FieldCoinUsdCurrency = "coin_usd_currency"
	// FieldCoinInfoID holds the string denoting the coin_info_id field in the database.
	FieldCoinInfoID = "coin_info_id"
	// FieldState holds the string denoting the state field in the database.
	FieldState = "state"
	// FieldChainTransactionID holds the string denoting the chain_transaction_id field in the database.
	FieldChainTransactionID = "chain_transaction_id"
	// FieldPlatformTransactionID holds the string denoting the platform_transaction_id field in the database.
	FieldPlatformTransactionID = "platform_transaction_id"
	// FieldUserSetPaid holds the string denoting the user_set_paid field in the database.
	FieldUserSetPaid = "user_set_paid"
	// FieldUserPaymentTxid holds the string denoting the user_payment_txid field in the database.
	FieldUserPaymentTxid = "user_payment_txid"
	// FieldCreateAt holds the string denoting the create_at field in the database.
	FieldCreateAt = "create_at"
	// FieldUpdateAt holds the string denoting the update_at field in the database.
	FieldUpdateAt = "update_at"
	// FieldDeleteAt holds the string denoting the delete_at field in the database.
	FieldDeleteAt = "delete_at"
	// Table holds the table name of the payment in the database.
	Table = "payments"
)

// Columns holds all SQL columns for payment fields.
var Columns = []string{
	FieldID,
	FieldAppID,
	FieldUserID,
	FieldGoodID,
	FieldOrderID,
	FieldAccountID,
	FieldStartAmount,
	FieldAmount,
	FieldFinishAmount,
	FieldCoinUsdCurrency,
	FieldCoinInfoID,
	FieldState,
	FieldChainTransactionID,
	FieldPlatformTransactionID,
	FieldUserSetPaid,
	FieldUserPaymentTxid,
	FieldCreateAt,
	FieldUpdateAt,
	FieldDeleteAt,
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

var (
	// DefaultUserSetPaid holds the default value on creation for the "user_set_paid" field.
	DefaultUserSetPaid bool
	// DefaultUserPaymentTxid holds the default value on creation for the "user_payment_txid" field.
	DefaultUserPaymentTxid string
	// DefaultCreateAt holds the default value on creation for the "create_at" field.
	DefaultCreateAt func() uint32
	// DefaultUpdateAt holds the default value on creation for the "update_at" field.
	DefaultUpdateAt func() uint32
	// UpdateDefaultUpdateAt holds the default value on update for the "update_at" field.
	UpdateDefaultUpdateAt func() uint32
	// DefaultDeleteAt holds the default value on creation for the "delete_at" field.
	DefaultDeleteAt func() uint32
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// State defines the type for the "state" enum field.
type State string

// State values.
const (
	StateWait     State = "wait"
	StateDone     State = "done"
	StateCanceled State = "canceled"
	StateTimeout  State = "timeout"
)

func (s State) String() string {
	return string(s)
}

// StateValidator is a validator for the "state" field enum values. It is called by the builders before save.
func StateValidator(s State) error {
	switch s {
	case StateWait, StateDone, StateCanceled, StateTimeout:
		return nil
	default:
		return fmt.Errorf("payment: invalid enum value for state field: %q", s)
	}
}
