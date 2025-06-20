package model

import (
	"time"
)

// StudentCareerProfile 咨询信息库_学员职业档案表
type StudentCareerProfile struct {
	StudentID      string    `json:"student_id" db:"student_id"`           // 学员学号
	Name           string    `json:"name" db:"name"`                       // 姓名
	ContactPhone   string    `json:"contact_phone" db:"contact_phone"`     // 联系电话
	EmploymentType string    `json:"employment_type" db:"employment_type"` // 工作类型
	Company        string    `json:"company" db:"company"`                 // 工作单位
	Department     string    `json:"department" db:"department"`           // 部门
	Position       string    `json:"position" db:"position"`               // 岗位
	BusinessArea   string    `json:"business_area" db:"business_area"`     // 业务方向
	StartTime      time.Time `json:"start_time" db:"start_time"`           // 开始时间
	EndTime        time.Time `json:"end_time" db:"end_time"`               // 结束时间
	Country        string    `json:"country" db:"country"`                 // 工作国家
	Region         string    `json:"region" db:"region"`                   // 工作地域
}
