package model

import (
	"time"
)

// CounselingOrderRecord 咨询信息库_辅导订单记录表
type CounselingOrderRecord struct {
	OrderNumber              string    `json:"order_number" db:"order_number"`                             // 订单编号
	ReferencedOrderNumber    string    `json:"referenced_order_number" db:"referenced_order_number"`       // 订单编号引用数据列
	TutorID                  string    `json:"tutor_id" db:"tutor_id"`                                     // 导师编号
	TutorNickname            string    `json:"tutor_nickname" db:"tutor_nickname"`                         // 导师花名
	StudentID                string    `json:"student_id" db:"student_id"`                                 // 学员学号
	StudentNickname          string    `json:"student_nickname" db:"student_nickname"`                     // 学员花名
	StudentConsultationCount int       `json:"student_consultation_count" db:"student_consultation_count"` // 本订单学生是第几次来平台咨询
	TutorStudentOrderCount   int       `json:"tutor_student_order_count" db:"tutor_student_order_count"`   // 本订单是该导师第几次辅导该学生
	StartTime                time.Time `json:"start_time" db:"start_time"`                                 // 辅导开始时间
	EndTime                  time.Time `json:"end_time" db:"end_time"`                                     // 辅导结束时间
	TargetPosition           string    `json:"target_position" db:"target_position"`                       // 目标岗位
	PlatformProductType      int       `json:"platform_product_type" db:"platform_product_type"`           // 本订单平台产品类型
	ResumeEvaluation         string    `json:"resume_evaluation" db:"resume_evaluation"`                   // 简历评价
	InterviewEvaluation      string    `json:"interview_evaluation" db:"interview_evaluation"`             // 模拟面试评价
	StudentAdvantages        string    `json:"student_advantages" db:"student_advantages"`                 // 学生的优势
	TutorSuggestions         string    `json:"tutor_suggestions" db:"tutor_suggestions"`                   // 导师建议
	StudentHighlightTags     []string  `json:"student_highlight_tags" db:"student_highlight_tags"`         // 学员亮点标签
	InterviewScore           int       `json:"interview_score" db:"interview_score"`                       // 模拟面试总评分（60-100）
	StudentFeedback          string    `json:"student_feedback" db:"student_feedback"`                     // 咨询感想反馈（学员填写）
	EffectivenessScore       int       `json:"effectiveness_score" db:"effectiveness_score"`               // 咨询效果反馈分（学员填写）
}
