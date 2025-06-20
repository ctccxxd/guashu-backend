package model

import (
	"time"
)

// TutorCareerProfile 咨询信息库_导师职业档案表
type TutorCareerProfile struct {
	TutorID         string    `json:"tutor_id" db:"tutor_id"`                 // 导师编号
	Name            string    `json:"name" db:"name"`                         // 姓名
	Nickname        string    `json:"nickname" db:"nickname"`                 // 花名
	ContactPhone    string    `json:"contact_phone" db:"contact_phone"`       // 联系电话
	Company         string    `json:"company" db:"company"`                   // 工作单位
	Department      string    `json:"department" db:"department"`             // 部门
	Position        string    `json:"position" db:"position"`                 // 岗位
	BusinessArea    string    `json:"business_area" db:"business_area"`       // 业务方向
	EmploymentType  string    `json:"employment_type" db:"employment_type"`   // 工作类型
	StartTime       time.Time `json:"start_time" db:"start_time"`             // 开始时间
	IsCurrentJob    bool      `json:"is_current_job" db:"is_current_job"`     // 是否工作至今
	StartYearMonth  string    `json:"start_year_month" db:"start_year_month"` // 开始年月
	EndTime         time.Time `json:"end_time" db:"end_time"`                 // 结束时间
	EndYearMonth    string    `json:"end_year_month" db:"end_year_month"`     // 结束年月
	TimeInterval    string    `json:"time_interval" db:"time_interval"`       // 时间间隔
	WorkingDuration string    `json:"working_duration" db:"working_duration"` // 工作时长
	Country         string    `json:"country" db:"country"`                   // 工作国家
	Region          string    `json:"region" db:"region"`                     // 工作地域
}
