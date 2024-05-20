package routers

import (
	"github.com/gin-gonic/gin"
	v1 "xlsqac/iachina-crm/routers/api/v1"

	"xlsqac/iachina-crm/pkg/setting"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)

	apiv1 := r.Group("/api/v1")
	//{
	apiv1.GET("/companies", v1.GetCompanies)
	apiv1.GET("/companies/:id", v1.GetCompanyById)
	apiv1.POST("/companies", v1.AddCompanies)
	apiv1.PUT("/companies/:id", v1.UpdateAddCompanies)
	apiv1.PATCH("/companies/:id", v1.UpdateCompanies)
	apiv1.DELETE("/companies/:id", v1.DeleteCompanies)
	//}

	return r
}
