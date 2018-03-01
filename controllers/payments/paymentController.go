package payments

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mpay/database"
	"github.com/mpay/models"
	"github.com/mpay/services"
)

//Router especifica las rutas a seguir
func Router(router *gin.RouterGroup) {
	router.GET("/:paymentId", getPayment)
	router.PUT("/:paymentId", putPayment)
	router.POST("", postPayment)
}

//getPayment Recupera un payment by id
func getPayment(c *gin.Context) {
	var paymentID = c.Param("paymentId")
	var payment models.Payment
	db := database.GetDB()
	db.First(&payment, paymentID)
	c.JSON(http.StatusOK, paymentID)
}

//postPayment Crea un payment
func postPayment(c *gin.Context) {
	var payment models.Payment
	c.BindJSON(&payment)
	var err error
	var ret *models.Payment
	if ret, err = services.PostPayment(&payment); err != nil {
		c.JSON(http.StatusUnprocessableEntity, database.NewError("Error", err))
	}
	c.JSON(http.StatusCreated, &ret)

}

//putPayment Aprueba o rechaza un payment
func putPayment(c *gin.Context) {
	var paymentID = c.Param("paymentId")
	var err error
	var payment models.Payment
	var updatedPayment models.Payment
	c.BindJSON(&payment)
	if updatedPayment, err = services.PutPayment(payment, paymentID); err != nil {
		c.JSON(http.StatusUnprocessableEntity, database.NewError("Error", err))
	}
	c.JSON(http.StatusOK, updatedPayment)

}
