package services

import (
	"database/sql"
	"errors"

	"github.com/mpay/database"
	"github.com/mpay/models"
	validator "gopkg.in/go-playground/validator.v9"
)

var validate *validator.Validate

//PostPayment provee un servicio para persistir un pago en la base de datos,
//también corre la validación y descarta aquellos valores que no son necesarios
func PostPayment(payment *models.Payment) (*models.Payment, error) {
	validate = validator.New()
	validate.RegisterCustomTypeFunc(models.ValidatePostPayment, sql.NullString{}, sql.NullInt64{}, sql.NullBool{}, sql.NullFloat64{})
	var err error
	if err = validate.Struct(payment); err != nil {
		return payment, err
	}
	setDefaultValues(payment)
	db := database.GetDB()
	err = db.Save(&payment).Error

	return payment, err
}

//PutPayment es el servicio encargado de aprobar o rechazar un pago.
//Los valores posible son: Approved / Cancel
func PutPayment(data models.Payment, paymentID string) (models.Payment, error) {
	var actualPayment models.Payment
	db := database.GetDB()
	var err error
	if err = db.First(&actualPayment, paymentID).Error; err != nil {
		return actualPayment, err
	}
	if data.Status == "Approved" {
		actualPayment.Status = data.Status
		actualPayment.StatusDetail = "Accredited"
		if err = db.Model(&actualPayment).Update(actualPayment).Error; err != nil {
			return actualPayment, err
		}
		return actualPayment, err

	}
	if data.Status == "Cancel" {
		actualPayment.Status = data.Status
		actualPayment.StatusDetail = "By Admin"
		if err = db.Model(&actualPayment).Update(actualPayment).Error; err != nil {
			return actualPayment, err
		}
		return actualPayment, err
	}

	return actualPayment, errors.New("Status must be Accredited/Cancel")
}

//setDefaultValues setea valores por defecto para los status
func setDefaultValues(payment *models.Payment) {
	payment.Status = "Pending"
	payment.StatusDetail = "Waiting for payment"
}
