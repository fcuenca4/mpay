package models

import (
	"database/sql/driver"
	"reflect"
	"time"
)

type Payment struct {
	ID           int64      `json:"id"`                            // id
	Collector    int64      `json:"collector" validate:"required"` // collector
	Payer        int64      `json:"payer" validate:"required"`     // payer
	CreationDate *time.Time `json:"creation_date"`                 // creation_date
	Amount       int64      `json:"amount" validate:"required"`    // amount
	Status       string     `json:"status"`                        // status
	StatusDetail string     `json:"status_detail"`                 // status_detail
	Metadata     string     `json:"metadata"`                      // metadata
}

func ValidatePostPayment(field reflect.Value) interface{} {

	if valuer, ok := field.Interface().(driver.Valuer); ok {
		val, err := valuer.Value()
		if err == nil {
			return val
		}
		// handle the error how you want
	}

	return nil
}
