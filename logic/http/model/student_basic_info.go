package model

import (
	"time"
)

// StudentBasicInfo 咨询信息库_学员基本信息表
type StudentBasicInfo struct {
	StudentID                 string    `json:"student_id" db:"student_id"`                                   // 学员学号
	ReferencedStudentID       string    `json:"referenced_student_id" db:"referenced_student_id"`             // 学员学号数据引用
	Name                      string    `json:"name" db:"name"`                                               // 姓名
	Nickname                  string    `json:"nickname" db:"nickname"`                                       // 花名
	Gender                    string    `json:"gender" db:"gender"`                                           // 性别
	StudentType               string    `json:"student_type" db:"student_type"`                               // 学员类别
	Major                     string    `json:"major" db:"major"`                                             // 专业
	EducationLevel            string    `json:"education_level" db:"education_level"`                         // 学历
	GraduationTime            time.Time `json:"graduation_time" db:"graduation_time"`                         // 毕业时间
	CurrentPosition           string    `json:"current_position" db:"current_position"`                       // 当前从事岗位（社招）
	LatestTargetPosition      string    `json:"latest_target_position" db:"latest_target_position"`           // 最新目标岗位
	LatestTargetBusiness      string    `json:"latest_target_business" db:"latest_target_business"`           // 最新目标业务
	TargetCountry             string    `json:"target_country" db:"target_country"`                           // 目标国家
	TargetRegion              string    `json:"target_region" db:"target_region"`                             // 目标地域
	ContactPhone              string    `json:"contact_phone" db:"contact_phone"`                             // 联系电话
	ContactEmail              string    `json:"contact_email" db:"contact_email"`                             // 联系邮箱
	PreviousConsultationCount int       `json:"previous_consultation_count" db:"previous_consultation_count"` // 之前在平台累计咨询次数
	OtherRemarks              string    `json:"other_remarks" db:"other_remarks"`                             // 其他备注
	Attachments               string    `json:"attachments" db:"attachments"`                                 // 附件
}
