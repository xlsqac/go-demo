package models

import "gorm.io/gorm"

// Company ia.company 表对应的 model
type Company struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"` // 单位名称
	IsValid int    `json:"-"`    // 0: 正常 1:删除(申请人撤销申请)
	//JoinUserID string `json:"join_user_id"` // 申请人 ID
	//IACTitleId            string `json:"iac_title_id"`             // 协会职务 ID
	//TypeID1               string `json:"type_id_1"`                // 一级类型
	//TypeID2               string `json:"type_id_2"`                // 二级类型
	//Address               string `json:"address"`                  // 办公地址
	//PostCode              string `json:"post_code"`                // 邮编
	//Jian                  string `json:"jian"`                     // 单位简称
	//en                    string `json:"en"`                       // 英文名称
	//FoundDate             string `json:"found_date"`               // 成立日期
	//CreditCode            string `json:"credit_code"`              // 统一社会信用代码
	//RegisteredCapital     string `json:"registered_capital"`       // 注册资金
	//RegAddress            string `json:"reg_address"`              // 注册地址
	//CompanyType           string `json:"company_type"`             // 单位类型
	//TotalAssets           string `json:"total_assets"`             // 总资产规模
	//Tel                   string `json:"tel"`                      // 总机电话
	//Web                   string `json:"web"`                      // 网站地址
	//Fax                   string `json:"fax"`                      // 传真
	//Administration        string `json:"administration"`           // 主管单位
	//OtherTitle            string `json:"other_title"`              // 在社会其他任职
	//LPName                string `json:"lp_name"`                  // 法人姓名
	//LPSex                 string `json:"lp_Sex"`                   // 法人性别
	//LPNotion              string `json:"lp_notion"`                // 法人民族
	//LPEdu                 string `json:"lp_edu"`                   // 法人学历
	//LPPo                  string `json:"lp_po"`                    // 法人政治面貌
	//LPBirth               string `json:"lp_birth"`                 // 法人出生日期
	//Executives            string `json:"executives"`               // 高管人数
	//Headcount             string `json:"headcount"`                // 职工总数
	//Departments           string `json:"departments"`              // 内设部门
	//Branchs               string `json:"branchs"`                  // 分公司数量
	//Summary               string `json:"summary"`                  // 单位简介
	//FileFace              string `json:"file_face"`                // 附件-入会申请书正面
	//FileBack              string `json:"file_back"`                // 附件-入会申请书背面
	//IsJoin                string `json:"is_join"`                  // 是否已入会 0: 已入会, 1: 未入会
	//LastOp                string `json:"last_op"`                  // 最后操作
	//LastActionId          string `json:"last_action_id"`           // 法人出生日期
	//LastPassedActionId    string `json:"last_passed_action_id"`    // 法人出生日期
	//LastReviewStatus      string `json:"last_review_status"`       // 最后一个审核状态, 1:审核通过 2:审核未通过/退回修改 3:等待审核
	//IsOverallReviewPassed string `json:"is_overall_review_passed"` // 所属模块全部审核通过
	//CreateTime string `json:"create_time"` // 创建时间
	//UpdateTime            string `json:"update_time"`              // 更新时间
	//LastPassedReviewId    string `json:"last_passed_review_id"`    // 最后一个审核状态, 1:审核通过 2:审核未通过/退回修改 3:等待审核
	//CheckRule             string `json:"check_rule"`               // 遵守入会章程 0: 是, 1: 否
	//Index                 string `json:"index"`                    // 公司排名
}

// APICompany 用作 select 字段
type APICompany struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// TableName 指定 Company 对应的表名，gorm 默认用蛇形风格的复数形式
func (c *Company) TableName() string {
	return "company" // 返回你想要的自定义表名
}

// GetCompaniesOrderById 获取指定页的公司信息
// 根据 id 升序
func GetCompaniesOrderById(pageNum int, pageSize int, conditions interface{}) (apiCompanies []APICompany, err error) {
	result := gormDB.Model(&Company{}).
		Where(conditions).Offset(pageNum).Limit(pageSize).Order("id").Find(&apiCompanies)
	err = result.Error
	return
}

// GetCompaniesTotal 获取符合条件的公司数量
func GetCompaniesTotal(conditions interface{}) (count int64, err error) {
	result := gormDB.Model(&Company{}).Where(conditions).Count(&count)
	err = result.Error
	return
}

// GetCompany 获取公司信息
func GetCompany(conditions interface{}) (apiCompany APICompany, err error) {
	result := gormDB.Model(&Company{}).Where(conditions).First(&apiCompany)
	err = result.Error
	return

}

// UpdateCompany 更新公司信息
func (c *Company) UpdateCompany(maps interface{}) {
	gormDB.Save(c)
}

// CreateCompany 创建一个公司
func CreateCompany(company *Company) *gorm.DB {
	result := gormDB.Create(company)
	return result
}

// DeleteCompany 删除会员单位
// 软删除，把 is_valid 设为 1
func DeleteCompany(conditions interface{}) error {
	result := gormDB.Model(&Company{}).Where(conditions).Update("is_valid", 1)
	return result.Error
}
