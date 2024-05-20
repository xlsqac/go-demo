package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"gorm.io/gorm"
	"log"
	"net/http"
	"xlsqac/iachina-crm/models"
	"xlsqac/iachina-crm/pkg/app"
	"xlsqac/iachina-crm/pkg/setting"
	"xlsqac/iachina-crm/pkg/util"
)

// GetCompanies 获取会员列表
func GetCompanies(c *gin.Context) {
	// 获取符合条件的数据数量
	conditions := make(map[string]interface{})
	conditions["is_valid"] = 0
	total, err := models.GetCompaniesTotal(conditions)
	if err != nil {
		log.Println(err)
	}

	// 获取一页数据
	lists, err := models.GetCompaniesOrderById(util.GetPage(c), setting.PageSize, conditions)
	if err != nil {
		log.Println(err)
	}
	resp := &app.GetCompaniesResp{
		List:  lists,
		Total: total,
	}

	c.JSON(http.StatusOK, resp)
}

// GetCompanyById 根据id 获取指定会员单位信息
func GetCompanyById(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt64()

	conditions := make(map[string]interface{})
	conditions["id"] = id
	conditions["is_valid"] = 0

	company, err := models.GetCompany(conditions)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.Status(http.StatusNotFound)
			return
		}
		log.Println(err)
	}

	resp := &app.CompanyResp{
		APICompany: models.APICompany{
			ID:   company.ID,
			Name: company.Name,
		},
	}
	c.JSON(http.StatusOK, resp)
}

// AddCompanies 新增会员
func AddCompanies(c *gin.Context) {
	//var req app.AddCompaniesReq
	//if err := c.ShouldBindJSON(&req); err != nil {
	//	// 返回错误信息
	//	// gin.H 封装了生成json数据的工具
	//	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//}
	//
	//valid := validation.Validation{}
	//valid.Required(req.Name, "name").Message("单位名称不能为空")
	//valid.MaxSize(req.Name, 100, "name").Message("单位名称最长为 100 字符")
	//
	//company := models.tag{
	//	Name: req.Name,
	//}
	//
	//if !valid.HasErrors() {
	//	models.AddCompany(&company)
	//} else {
	//	for _, err := range valid.Errors {
	//		c.JSON(http.StatusBadRequest, gin.H{
	//			"error": err.Message,
	//		})
	//	}
	//}

	//c.JSON(http.StatusOK, &app.AddCompaniesResp{Company: company})
}

// UpdateAddCompanies 全量更新会员信息
func UpdateAddCompanies(c *gin.Context) {
	//id := com.StrTo(c.Param("id")).MustInt64()
	var req app.UpdateAddCompaniesReq

	if err := c.ShouldBindJSON(&req); err != nil {
		// 返回错误信息
		// gin.H 封装了生成json数据的工具
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	//maps := make(map[string]interface{})
	//
	//maps["id"] = id
	//maps["is_valid"] = 0
	//company := models.GetCompanyById(maps)
	//// ID 等于 0 说明没有这条数据
	//if company.ID == 0 {
	//	models.AddCompany(&company)
	//} else {
	//	maps = make(map[string]interface{})
	//	maps["name"] = req.Name
	//	company = models.UpdateCompany(maps)
	//}
	//
	//c.JSON(http.StatusOK, &UpdateAddCompaniesResp{Id: company.ID, Name: company.Name})
}

// UpdateCompanies 局部更新会员信息
func UpdateCompanies(c *gin.Context) {}

// DeleteCompanies 删除会员
func DeleteCompanies(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt64()
	conditions := make(map[string]interface{})

	conditions["id"] = id
	conditions["is_valid"] = 0
	err := models.DeleteCompany(conditions)
	if err != nil {
		log.Println(err)
	}

	c.Status(http.StatusNoContent)
	return
}
