// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/cloud-hashing-order/pkg/db/ent/gaspaying"
	"github.com/google/uuid"
)

// GasPaying is the model entity for the GasPaying schema.
type GasPaying struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// OrderID holds the value of the "order_id" field.
	OrderID uuid.UUID `json:"order_id,omitempty"`
	// FeeTypeID holds the value of the "fee_type_id" field.
	FeeTypeID uuid.UUID `json:"fee_type_id,omitempty"`
	// PaymentID holds the value of the "payment_id" field.
	PaymentID uuid.UUID `json:"payment_id,omitempty"`
	// DurationMinutes holds the value of the "duration_minutes" field.
	DurationMinutes uint32 `json:"duration_minutes,omitempty"`
	// CreateAt holds the value of the "create_at" field.
	CreateAt uint32 `json:"create_at,omitempty"`
	// UpdateAt holds the value of the "update_at" field.
	UpdateAt uint32 `json:"update_at,omitempty"`
	// DeleteAt holds the value of the "delete_at" field.
	DeleteAt uint32 `json:"delete_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*GasPaying) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case gaspaying.FieldDurationMinutes, gaspaying.FieldCreateAt, gaspaying.FieldUpdateAt, gaspaying.FieldDeleteAt:
			values[i] = new(sql.NullInt64)
		case gaspaying.FieldID, gaspaying.FieldOrderID, gaspaying.FieldFeeTypeID, gaspaying.FieldPaymentID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type GasPaying", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the GasPaying fields.
func (gp *GasPaying) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case gaspaying.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				gp.ID = *value
			}
		case gaspaying.FieldOrderID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field order_id", values[i])
			} else if value != nil {
				gp.OrderID = *value
			}
		case gaspaying.FieldFeeTypeID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field fee_type_id", values[i])
			} else if value != nil {
				gp.FeeTypeID = *value
			}
		case gaspaying.FieldPaymentID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field payment_id", values[i])
			} else if value != nil {
				gp.PaymentID = *value
			}
		case gaspaying.FieldDurationMinutes:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field duration_minutes", values[i])
			} else if value.Valid {
				gp.DurationMinutes = uint32(value.Int64)
			}
		case gaspaying.FieldCreateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field create_at", values[i])
			} else if value.Valid {
				gp.CreateAt = uint32(value.Int64)
			}
		case gaspaying.FieldUpdateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field update_at", values[i])
			} else if value.Valid {
				gp.UpdateAt = uint32(value.Int64)
			}
		case gaspaying.FieldDeleteAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field delete_at", values[i])
			} else if value.Valid {
				gp.DeleteAt = uint32(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this GasPaying.
// Note that you need to call GasPaying.Unwrap() before calling this method if this GasPaying
// was returned from a transaction, and the transaction was committed or rolled back.
func (gp *GasPaying) Update() *GasPayingUpdateOne {
	return (&GasPayingClient{config: gp.config}).UpdateOne(gp)
}

// Unwrap unwraps the GasPaying entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (gp *GasPaying) Unwrap() *GasPaying {
	tx, ok := gp.config.driver.(*txDriver)
	if !ok {
		panic("ent: GasPaying is not a transactional entity")
	}
	gp.config.driver = tx.drv
	return gp
}

// String implements the fmt.Stringer.
func (gp *GasPaying) String() string {
	var builder strings.Builder
	builder.WriteString("GasPaying(")
	builder.WriteString(fmt.Sprintf("id=%v", gp.ID))
	builder.WriteString(", order_id=")
	builder.WriteString(fmt.Sprintf("%v", gp.OrderID))
	builder.WriteString(", fee_type_id=")
	builder.WriteString(fmt.Sprintf("%v", gp.FeeTypeID))
	builder.WriteString(", payment_id=")
	builder.WriteString(fmt.Sprintf("%v", gp.PaymentID))
	builder.WriteString(", duration_minutes=")
	builder.WriteString(fmt.Sprintf("%v", gp.DurationMinutes))
	builder.WriteString(", create_at=")
	builder.WriteString(fmt.Sprintf("%v", gp.CreateAt))
	builder.WriteString(", update_at=")
	builder.WriteString(fmt.Sprintf("%v", gp.UpdateAt))
	builder.WriteString(", delete_at=")
	builder.WriteString(fmt.Sprintf("%v", gp.DeleteAt))
	builder.WriteByte(')')
	return builder.String()
}

// GasPayings is a parsable slice of GasPaying.
type GasPayings []*GasPaying

func (gp GasPayings) config(cfg config) {
	for _i := range gp {
		gp[_i].config = cfg
	}
}
