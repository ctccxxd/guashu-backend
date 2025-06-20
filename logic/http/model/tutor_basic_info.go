package model

import (
	"time"
)

// TutorBasicInfo 咨询信息库_导师基本信息表
type TutorBasicInfo struct {
	TutorID                       string    `json:"tutor_id" db:"tutor_id"`                                                   // 导师编号
	ReferencedTutorID             string    `json:"referenced_tutor_id" db:"referenced_tutor_id"`                             // 导师编号数据引用列
	Name                          string    `json:"name" db:"name"`                                                           // 姓名
	PlatformCertificationStatus   string    `json:"platform_certification_status" db:"platform_certification_status"`         // 导师背景平台认证情况
	Nickname                      string    `json:"nickname" db:"nickname"`                                                   // 花名
	ContactPhone                  string    `json:"contact_phone" db:"contact_phone"`                                         // 联系电话
	Gender                        string    `json:"gender" db:"gender"`                                                       // 性别
	HighestEducationLevel         string    `json:"highest_education_level" db:"highest_education_level"`                     // 最高学历
	GraduationSchool              string    `json:"graduation_school" db:"graduation_school"`                                 // 毕业院校
	Major                         string    `json:"major" db:"major"`                                                         // 在校专业
	GraduationTime                time.Time `json:"graduation_time" db:"graduation_time"`                                     // 毕业时间
	PersonalIntroduction          string    `json:"personal_introduction" db:"personal_introduction"`                         // 个人介绍
	SpecializedPositions          string    `json:"specialized_positions" db:"specialized_positions"`                         // 从事并擅长的岗位类型
	HasProgrammingSkills          bool      `json:"has_programming_skills" db:"has_programming_skills"`                       // 是否有擅长的编程语言
	ProgrammingLanguages          string    `json:"programming_languages" db:"programming_languages"`                         // 擅长的编程语言
	WillingToTakePaidOrders       bool      `json:"willing_to_take_paid_orders" db:"willing_to_take_paid_orders"`             // 后续是否想在平台接付费单
	AvailablePaidServices         string    `json:"available_paid_services" db:"available_paid_services"`                     // 可以承接的付费业务（可多选）
	SupportedServices             string    `json:"supported_services" db:"supported_services"`                               // 导师目前支持的业务
	ResumeCoachingPrice           float64   `json:"resume_coaching_price" db:"resume_coaching_price"`                         // 简历辅导价格
	ProductResumePrice            float64   `json:"product_resume_price" db:"product_resume_price"`                           // 产品简历辅导的价格
	WorkplaceInquiryPrice         float64   `json:"workplace_inquiry_price" db:"workplace_inquiry_price"`                     // 团队内部工作环境打听咨询价格
	ProductWorkplacePrice         float64   `json:"product_workplace_price" db:"product_workplace_price"`                     // 产品团队内部工作环境打听咨询价格
	MockInterviewPrice            float64   `json:"mock_interview_price" db:"mock_interview_price"`                           // 模拟面试价格
	ProductMockInterviewPrice     float64   `json:"product_mock_interview_price" db:"product_mock_interview_price"`           // 产品模拟面试的价格
	CareerCoachingPrice           float64   `json:"career_coaching_price" db:"career_coaching_price"`                         // 职业规划价格
	ProductCareerCoachingPrice    float64   `json:"product_career_coaching_price" db:"product_career_coaching_price"`         // 产品职业规划的价格
	ThirtyDayCoachPrice           float64   `json:"thirty_day_coach_price" db:"thirty_day_coach_price"`                       // 30天求职教练价格
	ProductThirtyDayCoachPrice    float64   `json:"product_thirty_day_coach_price" db:"product_thirty_day_coach_price"`       // 产品30天求职教练的价格
	AgreeToTrainingAndSupervision bool      `json:"agree_to_training_and_supervision" db:"agree_to_training_and_supervision"` // 是否同意遵守后续接单平台同意培训和监督
	BusyStatus                    string    `json:"busy_status" db:"busy_status"`                                             // 导师 闲/忙 状态标志
	WasFormerStudent              bool      `json:"was_former_student" db:"was_former_student"`                               // 是否曾是平台学员
	TotalOrders                   int       `json:"total_orders" db:"total_orders"`                                           // 平台累计接单次数
	OtherRemarks                  string    `json:"other_remarks" db:"other_remarks"`                                         // 其他备注
	Attachments                   string    `json:"attachments" db:"attachments"`                                             // 附件
}
