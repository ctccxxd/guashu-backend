package constants

type OrderType int

const (
	ResumeDiagnosisOnly   OrderType = iota + 1 // 只做简历诊断
	InterviewCoachingOnly                      // 只做面试辅导
	ResumeAndInterview                         // 包含简历诊断和面试辅导两项服务
)

const (
	SmtpUser = ""
	SmtpPass = "" // 授权码
	SmtpHost = "smtp.163.com"
	SmtpPort = 465 // SSL端口
)

const (
	WkhtmltoimagePath = "thirdParty/wkhtmltoimage.exe"
)

const (
	HighlightsMapPath       = "high_lights_map.yaml"
	InterviewAssessmentPath = "template/interview_assessment.html"
	ResumeAssessmentPath    = "template/resume_assessment.html"
)
