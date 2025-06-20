package model

import (
	"time"
)

// CounselingRecord 咨询信息库_30天陪跑辅导记录
type CounselingRecord30 struct {
	// 订单相关信息
	OrderNumber     string `json:"order_number" db:"order_number"`         // 订单的编号/订单编号
	TutorID         string `json:"tutor_id" db:"tutor_id"`                 // 导师编号
	TutorNickname   string `json:"tutor_nickname" db:"tutor_nickname"`     // 导师花名
	StudentID       string `json:"student_id" db:"student_id"`             // 学员学号
	StudentNickname string `json:"student_nickname" db:"student_nickname"` // 学员花名
	SessionCount    int    `json:"session_count" db:"session_count"`       // 本次是该订单第几次陪跑学生

	// 辅导时间信息
	StartTime time.Time `json:"start_time" db:"start_time"` // 辅导开始时间
	EndTime   time.Time `json:"end_time" db:"end_time"`     // 辅导结束时间

	// 辅导内容信息
	TargetPosition      string   `json:"target_position" db:"target_position"`           // 目标岗位
	ThemeType           string   `json:"theme_type" db:"theme_type"`                     // 本次陪跑的主题类型
	ResumeEvaluation    string   `json:"resume_evaluation" db:"resume_evaluation"`       // 简历评价
	InterviewEvaluation string   `json:"interview_evaluation" db:"interview_evaluation"` // 模拟面试评价
	StudentAdvantage    string   `json:"student_advantage" db:"student_advantage"`       // 学生在该项的优势
	TutorSuggestion     string   `json:"tutor_suggestion" db:"tutor_suggestion"`         // 导师该项建议
	TutorSummary        string   `json:"tutor_summary" db:"tutor_summary"`               // 导师陪跑总结
	StudentHighlight    []string `json:"student_highlight" db:"student_highlight"`       // 总评学员亮点标签
	InterviewScore      int      `json:"interview_score" db:"interview_score"`           // 模拟面试总评分（60-100）
}
