package model

import (
	"time"
)

type InterviewCoachingPromotion struct {
	Email                       string    `json:"email"`
	TutorNickname               string    `json:"tutor_nickname"`                // 导师花名
	StudentNickname             string    `json:"student_nickname"`              // 学员花名
	StartTime                   time.Time `json:"start_time"`                    // 辅导开始时间
	EndTime                     time.Time `json:"end_time"`                      // 辅导结束时间
	TargetPosition              string    `json:"target_position"`               // 目标岗位
	PlatformProductType         int       `json:"platform_product_type"`         // 本订单平台产品类型
	ResumeEvaluation            string    `json:"resume_evaluation"`             // 简历评价
	InterviewEvaluation         string    `json:"interview_evaluation"`          // 模拟面试评价
	StudentAdvantages           string    `json:"student_advantages"`            // 学生的优势
	TutorSuggestions            string    `json:"tutor_suggestions"`             // 导师建议
	StudentHighlightTags        []string  `json:"student_highlight_tags"`        // 学员亮点标签
	InterviewScore              int       `json:"interview_score"`               // 模拟面试总评分（60-100）
	PlatformCertificationStatus string    `json:"platform_certification_status"` // 导师背景平台认证情况
}
