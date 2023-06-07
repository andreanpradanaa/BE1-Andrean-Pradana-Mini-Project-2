package customer

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RouterCustomer struct {
	CustomerRequestHandler RequestHandlerCustomer
}

func NewRouter(
	dbCrud *gorm.DB,
) RouterCustomer {
	return RouterCustomer{CustomerRequestHandler: NewCustomerRequestHandler(
		dbCrud,
	)}
}

func (r RouterCustomer) Handle(routeVersion *gin.RouterGroup) {
	basepath := "/customer"
	user := routeVersion.Group(basepath)

	user.POST("/",
		r.CustomerRequestHandler.CreateCustomer,
	)

	user.GET("/:id",
		r.CustomerRequestHandler.GetCustomerById,
	)

	user.PUT("/:id",
		r.CustomerRequestHandler.UpdateCustomer,
	)
	user.DELETE("/:email",
		r.CustomerRequestHandler.DeleteCustomer,
	)
}
