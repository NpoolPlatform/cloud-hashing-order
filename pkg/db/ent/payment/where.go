// Code generated by entc, DO NOT EDIT.

package payment

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/cloud-hashing-order/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// AppID applies equality check predicate on the "app_id" field. It's identical to AppIDEQ.
func AppID(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppID), v))
	})
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// OrderID applies equality check predicate on the "order_id" field. It's identical to OrderIDEQ.
func OrderID(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOrderID), v))
	})
}

// AccountID applies equality check predicate on the "account_id" field. It's identical to AccountIDEQ.
func AccountID(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAccountID), v))
	})
}

// StartAmount applies equality check predicate on the "start_amount" field. It's identical to StartAmountEQ.
func StartAmount(v uint64) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStartAmount), v))
	})
}

// Amount applies equality check predicate on the "amount" field. It's identical to AmountEQ.
func Amount(v uint64) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAmount), v))
	})
}

// CoinUsdCurrency applies equality check predicate on the "coin_usd_currency" field. It's identical to CoinUsdCurrencyEQ.
func CoinUsdCurrency(v uint64) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCoinUsdCurrency), v))
	})
}

// CoinInfoID applies equality check predicate on the "coin_info_id" field. It's identical to CoinInfoIDEQ.
func CoinInfoID(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCoinInfoID), v))
	})
}

// ChainTransactionID applies equality check predicate on the "chain_transaction_id" field. It's identical to ChainTransactionIDEQ.
func ChainTransactionID(v string) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldChainTransactionID), v))
	})
}

// PlatformTransactionID applies equality check predicate on the "platform_transaction_id" field. It's identical to PlatformTransactionIDEQ.
func PlatformTransactionID(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPlatformTransactionID), v))
	})
}

// CreateAt applies equality check predicate on the "create_at" field. It's identical to CreateAtEQ.
func CreateAt(v uint32) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateAt), v))
	})
}

// UpdateAt applies equality check predicate on the "update_at" field. It's identical to UpdateAtEQ.
func UpdateAt(v uint32) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateAt), v))
	})
}

// DeleteAt applies equality check predicate on the "delete_at" field. It's identical to DeleteAtEQ.
func DeleteAt(v uint32) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeleteAt), v))
	})
}

// AppIDEQ applies the EQ predicate on the "app_id" field.
func AppIDEQ(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppID), v))
	})
}

// AppIDNEQ applies the NEQ predicate on the "app_id" field.
func AppIDNEQ(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAppID), v))
	})
}

// AppIDIn applies the In predicate on the "app_id" field.
func AppIDIn(vs ...uuid.UUID) predicate.Payment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAppID), v...))
	})
}

// AppIDNotIn applies the NotIn predicate on the "app_id" field.
func AppIDNotIn(vs ...uuid.UUID) predicate.Payment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAppID), v...))
	})
}

// AppIDGT applies the GT predicate on the "app_id" field.
func AppIDGT(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAppID), v))
	})
}

// AppIDGTE applies the GTE predicate on the "app_id" field.
func AppIDGTE(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAppID), v))
	})
}

// AppIDLT applies the LT predicate on the "app_id" field.
func AppIDLT(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAppID), v))
	})
}

// AppIDLTE applies the LTE predicate on the "app_id" field.
func AppIDLTE(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAppID), v))
	})
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUserID), v))
	})
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...uuid.UUID) predicate.Payment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUserID), v...))
	})
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...uuid.UUID) predicate.Payment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUserID), v...))
	})
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUserID), v))
	})
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUserID), v))
	})
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUserID), v))
	})
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUserID), v))
	})
}

// OrderIDEQ applies the EQ predicate on the "order_id" field.
func OrderIDEQ(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOrderID), v))
	})
}

// OrderIDNEQ applies the NEQ predicate on the "order_id" field.
func OrderIDNEQ(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOrderID), v))
	})
}

// OrderIDIn applies the In predicate on the "order_id" field.
func OrderIDIn(vs ...uuid.UUID) predicate.Payment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOrderID), v...))
	})
}

// OrderIDNotIn applies the NotIn predicate on the "order_id" field.
func OrderIDNotIn(vs ...uuid.UUID) predicate.Payment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOrderID), v...))
	})
}

// OrderIDGT applies the GT predicate on the "order_id" field.
func OrderIDGT(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOrderID), v))
	})
}

// OrderIDGTE applies the GTE predicate on the "order_id" field.
func OrderIDGTE(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOrderID), v))
	})
}

// OrderIDLT applies the LT predicate on the "order_id" field.
func OrderIDLT(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOrderID), v))
	})
}

// OrderIDLTE applies the LTE predicate on the "order_id" field.
func OrderIDLTE(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOrderID), v))
	})
}

// AccountIDEQ applies the EQ predicate on the "account_id" field.
func AccountIDEQ(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAccountID), v))
	})
}

// AccountIDNEQ applies the NEQ predicate on the "account_id" field.
func AccountIDNEQ(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAccountID), v))
	})
}

// AccountIDIn applies the In predicate on the "account_id" field.
func AccountIDIn(vs ...uuid.UUID) predicate.Payment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAccountID), v...))
	})
}

// AccountIDNotIn applies the NotIn predicate on the "account_id" field.
func AccountIDNotIn(vs ...uuid.UUID) predicate.Payment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAccountID), v...))
	})
}

// AccountIDGT applies the GT predicate on the "account_id" field.
func AccountIDGT(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAccountID), v))
	})
}

// AccountIDGTE applies the GTE predicate on the "account_id" field.
func AccountIDGTE(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAccountID), v))
	})
}

// AccountIDLT applies the LT predicate on the "account_id" field.
func AccountIDLT(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAccountID), v))
	})
}

// AccountIDLTE applies the LTE predicate on the "account_id" field.
func AccountIDLTE(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAccountID), v))
	})
}

// StartAmountEQ applies the EQ predicate on the "start_amount" field.
func StartAmountEQ(v uint64) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStartAmount), v))
	})
}

// StartAmountNEQ applies the NEQ predicate on the "start_amount" field.
func StartAmountNEQ(v uint64) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStartAmount), v))
	})
}

// StartAmountIn applies the In predicate on the "start_amount" field.
func StartAmountIn(vs ...uint64) predicate.Payment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStartAmount), v...))
	})
}

// StartAmountNotIn applies the NotIn predicate on the "start_amount" field.
func StartAmountNotIn(vs ...uint64) predicate.Payment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStartAmount), v...))
	})
}

// StartAmountGT applies the GT predicate on the "start_amount" field.
func StartAmountGT(v uint64) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStartAmount), v))
	})
}

// StartAmountGTE applies the GTE predicate on the "start_amount" field.
func StartAmountGTE(v uint64) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStartAmount), v))
	})
}

// StartAmountLT applies the LT predicate on the "start_amount" field.
func StartAmountLT(v uint64) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStartAmount), v))
	})
}

// StartAmountLTE applies the LTE predicate on the "start_amount" field.
func StartAmountLTE(v uint64) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStartAmount), v))
	})
}

// AmountEQ applies the EQ predicate on the "amount" field.
func AmountEQ(v uint64) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAmount), v))
	})
}

// AmountNEQ applies the NEQ predicate on the "amount" field.
func AmountNEQ(v uint64) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAmount), v))
	})
}

// AmountIn applies the In predicate on the "amount" field.
func AmountIn(vs ...uint64) predicate.Payment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAmount), v...))
	})
}

// AmountNotIn applies the NotIn predicate on the "amount" field.
func AmountNotIn(vs ...uint64) predicate.Payment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAmount), v...))
	})
}

// AmountGT applies the GT predicate on the "amount" field.
func AmountGT(v uint64) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAmount), v))
	})
}

// AmountGTE applies the GTE predicate on the "amount" field.
func AmountGTE(v uint64) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAmount), v))
	})
}

// AmountLT applies the LT predicate on the "amount" field.
func AmountLT(v uint64) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAmount), v))
	})
}

// AmountLTE applies the LTE predicate on the "amount" field.
func AmountLTE(v uint64) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAmount), v))
	})
}

// CoinUsdCurrencyEQ applies the EQ predicate on the "coin_usd_currency" field.
func CoinUsdCurrencyEQ(v uint64) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCoinUsdCurrency), v))
	})
}

// CoinUsdCurrencyNEQ applies the NEQ predicate on the "coin_usd_currency" field.
func CoinUsdCurrencyNEQ(v uint64) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCoinUsdCurrency), v))
	})
}

// CoinUsdCurrencyIn applies the In predicate on the "coin_usd_currency" field.
func CoinUsdCurrencyIn(vs ...uint64) predicate.Payment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCoinUsdCurrency), v...))
	})
}

// CoinUsdCurrencyNotIn applies the NotIn predicate on the "coin_usd_currency" field.
func CoinUsdCurrencyNotIn(vs ...uint64) predicate.Payment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCoinUsdCurrency), v...))
	})
}

// CoinUsdCurrencyGT applies the GT predicate on the "coin_usd_currency" field.
func CoinUsdCurrencyGT(v uint64) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCoinUsdCurrency), v))
	})
}

// CoinUsdCurrencyGTE applies the GTE predicate on the "coin_usd_currency" field.
func CoinUsdCurrencyGTE(v uint64) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCoinUsdCurrency), v))
	})
}

// CoinUsdCurrencyLT applies the LT predicate on the "coin_usd_currency" field.
func CoinUsdCurrencyLT(v uint64) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCoinUsdCurrency), v))
	})
}

// CoinUsdCurrencyLTE applies the LTE predicate on the "coin_usd_currency" field.
func CoinUsdCurrencyLTE(v uint64) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCoinUsdCurrency), v))
	})
}

// CoinInfoIDEQ applies the EQ predicate on the "coin_info_id" field.
func CoinInfoIDEQ(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCoinInfoID), v))
	})
}

// CoinInfoIDNEQ applies the NEQ predicate on the "coin_info_id" field.
func CoinInfoIDNEQ(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCoinInfoID), v))
	})
}

// CoinInfoIDIn applies the In predicate on the "coin_info_id" field.
func CoinInfoIDIn(vs ...uuid.UUID) predicate.Payment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCoinInfoID), v...))
	})
}

// CoinInfoIDNotIn applies the NotIn predicate on the "coin_info_id" field.
func CoinInfoIDNotIn(vs ...uuid.UUID) predicate.Payment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCoinInfoID), v...))
	})
}

// CoinInfoIDGT applies the GT predicate on the "coin_info_id" field.
func CoinInfoIDGT(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCoinInfoID), v))
	})
}

// CoinInfoIDGTE applies the GTE predicate on the "coin_info_id" field.
func CoinInfoIDGTE(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCoinInfoID), v))
	})
}

// CoinInfoIDLT applies the LT predicate on the "coin_info_id" field.
func CoinInfoIDLT(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCoinInfoID), v))
	})
}

// CoinInfoIDLTE applies the LTE predicate on the "coin_info_id" field.
func CoinInfoIDLTE(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCoinInfoID), v))
	})
}

// StateEQ applies the EQ predicate on the "state" field.
func StateEQ(v State) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldState), v))
	})
}

// StateNEQ applies the NEQ predicate on the "state" field.
func StateNEQ(v State) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldState), v))
	})
}

// StateIn applies the In predicate on the "state" field.
func StateIn(vs ...State) predicate.Payment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldState), v...))
	})
}

// StateNotIn applies the NotIn predicate on the "state" field.
func StateNotIn(vs ...State) predicate.Payment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldState), v...))
	})
}

// ChainTransactionIDEQ applies the EQ predicate on the "chain_transaction_id" field.
func ChainTransactionIDEQ(v string) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldChainTransactionID), v))
	})
}

// ChainTransactionIDNEQ applies the NEQ predicate on the "chain_transaction_id" field.
func ChainTransactionIDNEQ(v string) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldChainTransactionID), v))
	})
}

// ChainTransactionIDIn applies the In predicate on the "chain_transaction_id" field.
func ChainTransactionIDIn(vs ...string) predicate.Payment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldChainTransactionID), v...))
	})
}

// ChainTransactionIDNotIn applies the NotIn predicate on the "chain_transaction_id" field.
func ChainTransactionIDNotIn(vs ...string) predicate.Payment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldChainTransactionID), v...))
	})
}

// ChainTransactionIDGT applies the GT predicate on the "chain_transaction_id" field.
func ChainTransactionIDGT(v string) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldChainTransactionID), v))
	})
}

// ChainTransactionIDGTE applies the GTE predicate on the "chain_transaction_id" field.
func ChainTransactionIDGTE(v string) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldChainTransactionID), v))
	})
}

// ChainTransactionIDLT applies the LT predicate on the "chain_transaction_id" field.
func ChainTransactionIDLT(v string) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldChainTransactionID), v))
	})
}

// ChainTransactionIDLTE applies the LTE predicate on the "chain_transaction_id" field.
func ChainTransactionIDLTE(v string) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldChainTransactionID), v))
	})
}

// ChainTransactionIDContains applies the Contains predicate on the "chain_transaction_id" field.
func ChainTransactionIDContains(v string) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldChainTransactionID), v))
	})
}

// ChainTransactionIDHasPrefix applies the HasPrefix predicate on the "chain_transaction_id" field.
func ChainTransactionIDHasPrefix(v string) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldChainTransactionID), v))
	})
}

// ChainTransactionIDHasSuffix applies the HasSuffix predicate on the "chain_transaction_id" field.
func ChainTransactionIDHasSuffix(v string) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldChainTransactionID), v))
	})
}

// ChainTransactionIDEqualFold applies the EqualFold predicate on the "chain_transaction_id" field.
func ChainTransactionIDEqualFold(v string) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldChainTransactionID), v))
	})
}

// ChainTransactionIDContainsFold applies the ContainsFold predicate on the "chain_transaction_id" field.
func ChainTransactionIDContainsFold(v string) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldChainTransactionID), v))
	})
}

// PlatformTransactionIDEQ applies the EQ predicate on the "platform_transaction_id" field.
func PlatformTransactionIDEQ(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPlatformTransactionID), v))
	})
}

// PlatformTransactionIDNEQ applies the NEQ predicate on the "platform_transaction_id" field.
func PlatformTransactionIDNEQ(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPlatformTransactionID), v))
	})
}

// PlatformTransactionIDIn applies the In predicate on the "platform_transaction_id" field.
func PlatformTransactionIDIn(vs ...uuid.UUID) predicate.Payment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPlatformTransactionID), v...))
	})
}

// PlatformTransactionIDNotIn applies the NotIn predicate on the "platform_transaction_id" field.
func PlatformTransactionIDNotIn(vs ...uuid.UUID) predicate.Payment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPlatformTransactionID), v...))
	})
}

// PlatformTransactionIDGT applies the GT predicate on the "platform_transaction_id" field.
func PlatformTransactionIDGT(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPlatformTransactionID), v))
	})
}

// PlatformTransactionIDGTE applies the GTE predicate on the "platform_transaction_id" field.
func PlatformTransactionIDGTE(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPlatformTransactionID), v))
	})
}

// PlatformTransactionIDLT applies the LT predicate on the "platform_transaction_id" field.
func PlatformTransactionIDLT(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPlatformTransactionID), v))
	})
}

// PlatformTransactionIDLTE applies the LTE predicate on the "platform_transaction_id" field.
func PlatformTransactionIDLTE(v uuid.UUID) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPlatformTransactionID), v))
	})
}

// CreateAtEQ applies the EQ predicate on the "create_at" field.
func CreateAtEQ(v uint32) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateAt), v))
	})
}

// CreateAtNEQ applies the NEQ predicate on the "create_at" field.
func CreateAtNEQ(v uint32) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateAt), v))
	})
}

// CreateAtIn applies the In predicate on the "create_at" field.
func CreateAtIn(vs ...uint32) predicate.Payment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreateAt), v...))
	})
}

// CreateAtNotIn applies the NotIn predicate on the "create_at" field.
func CreateAtNotIn(vs ...uint32) predicate.Payment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreateAt), v...))
	})
}

// CreateAtGT applies the GT predicate on the "create_at" field.
func CreateAtGT(v uint32) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateAt), v))
	})
}

// CreateAtGTE applies the GTE predicate on the "create_at" field.
func CreateAtGTE(v uint32) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateAt), v))
	})
}

// CreateAtLT applies the LT predicate on the "create_at" field.
func CreateAtLT(v uint32) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateAt), v))
	})
}

// CreateAtLTE applies the LTE predicate on the "create_at" field.
func CreateAtLTE(v uint32) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateAt), v))
	})
}

// UpdateAtEQ applies the EQ predicate on the "update_at" field.
func UpdateAtEQ(v uint32) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtNEQ applies the NEQ predicate on the "update_at" field.
func UpdateAtNEQ(v uint32) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtIn applies the In predicate on the "update_at" field.
func UpdateAtIn(vs ...uint32) predicate.Payment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdateAt), v...))
	})
}

// UpdateAtNotIn applies the NotIn predicate on the "update_at" field.
func UpdateAtNotIn(vs ...uint32) predicate.Payment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdateAt), v...))
	})
}

// UpdateAtGT applies the GT predicate on the "update_at" field.
func UpdateAtGT(v uint32) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtGTE applies the GTE predicate on the "update_at" field.
func UpdateAtGTE(v uint32) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtLT applies the LT predicate on the "update_at" field.
func UpdateAtLT(v uint32) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtLTE applies the LTE predicate on the "update_at" field.
func UpdateAtLTE(v uint32) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateAt), v))
	})
}

// DeleteAtEQ applies the EQ predicate on the "delete_at" field.
func DeleteAtEQ(v uint32) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeleteAt), v))
	})
}

// DeleteAtNEQ applies the NEQ predicate on the "delete_at" field.
func DeleteAtNEQ(v uint32) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeleteAt), v))
	})
}

// DeleteAtIn applies the In predicate on the "delete_at" field.
func DeleteAtIn(vs ...uint32) predicate.Payment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDeleteAt), v...))
	})
}

// DeleteAtNotIn applies the NotIn predicate on the "delete_at" field.
func DeleteAtNotIn(vs ...uint32) predicate.Payment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDeleteAt), v...))
	})
}

// DeleteAtGT applies the GT predicate on the "delete_at" field.
func DeleteAtGT(v uint32) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeleteAt), v))
	})
}

// DeleteAtGTE applies the GTE predicate on the "delete_at" field.
func DeleteAtGTE(v uint32) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeleteAt), v))
	})
}

// DeleteAtLT applies the LT predicate on the "delete_at" field.
func DeleteAtLT(v uint32) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeleteAt), v))
	})
}

// DeleteAtLTE applies the LTE predicate on the "delete_at" field.
func DeleteAtLTE(v uint32) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeleteAt), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Payment) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Payment) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Payment) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		p(s.Not())
	})
}
