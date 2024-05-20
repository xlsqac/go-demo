package app

import "xlsqac/iachina-crm/models"

// CompanyResp 返回一个包含公司信息的响应
// get, post, put 都会返回
type CompanyResp struct {
	models.APICompany
}

type GetCompaniesResp struct {
	List  []models.APICompany `json:"lists"`
	Total int64               `json:"total"`
}
