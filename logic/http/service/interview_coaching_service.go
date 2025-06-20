package service

import (
	"os"
	"sort"
	"time"

	"gopkg.in/yaml.v3"
	applogger "jihulab.com/guashu/gs-server/lib/logger"
	"jihulab.com/guashu/gs-server/logic/http/model"
	"jihulab.com/guashu/gs-server/util/constants"
	"jihulab.com/guashu/gs-server/util/email"
	"jihulab.com/guashu/gs-server/util/htmlkit"
)

var highlightsMap map[string]string

func init() {
	data, err := os.ReadFile(constants.HighlightsMapPath)
	if err != nil {
		applogger.Sugar.Error("读取文件失败: %v\n", err)
		panic("读取文件失败.")
	}

	// 解析为映射
	var result map[string]string
	err = yaml.Unmarshal(data, &result)
	if err != nil {
		applogger.Sugar.Error("解析 YAML 失败: %v\n", err)
		panic("解析 YAML 失败.")
	}

	highlightsMap = result
}

type InterviewCoachingService interface {
	SendEmail(model.InterviewCoachingPromotion) error
}

type interviewCoachingService struct {
}

func NewInterviewCoachingService() InterviewCoachingService {
	return &interviewCoachingService{}
}

func (i *interviewCoachingService) SendEmail(requestData model.InterviewCoachingPromotion) error {
	images := [][]byte{}
	imageNames := []string{}
	text := ""
	data := toMap(requestData)

	if requestData.PlatformProductType == int(constants.ResumeDiagnosisOnly) || requestData.PlatformProductType == int(constants.ResumeAndInterview) { // 简历
		htmlCode, err := htmlkit.RenderHTML(constants.ResumeAssessmentPath, data)
		if err != nil {
			applogger.Sugar.Error("渲染html失败: %v\n", err)
			return err
		}
		imageBytes, err := htmlkit.HTMLToImageBytes(htmlCode, nil)
		if err != nil {
			applogger.Sugar.Error("html转图片失败: %v\n", err)
			return err
		}
		images = append(images, imageBytes)
		imageNames = append(imageNames, "简历诊断报告.png")
		text = "<p>感谢您选择我们的求职辅导服务，以下是您的详细评估报告。我们针对您的简历提供了专业分析,详情在附件，请查收！</p>"
	}

	if requestData.PlatformProductType == int(constants.InterviewCoachingOnly) || requestData.PlatformProductType == int(constants.ResumeAndInterview) { // 面试
		htmlCode, err := htmlkit.RenderHTML(constants.InterviewAssessmentPath, data)
		if err != nil {
			applogger.Sugar.Error("渲染html失败: %v\n", err)
			return err
		}
		imageBytes, err := htmlkit.HTMLToImageBytes(htmlCode, nil)
		if err != nil {
			applogger.Sugar.Error("html转图片失败: %v\n", err)
			return err
		}
		images = append(images, imageBytes)
		imageNames = append(imageNames, "简历诊断报告.png")
		text = "<p>感谢您选择我们的求职辅导服务，以下是您的详细评估报告。我们针对您的面试表现提供了专业分析和改进建议,详情在附件，请查收！</p>"
	}

	if requestData.PlatformProductType == int(constants.ResumeAndInterview) { // 简历 + 面试
		text = "<p>感谢您选择我们的求职辅导服务，以下是您的详细评估报告。我们针对您的简历和面试表现提供了专业分析和改进建议,详情在附件，请查收！</p>"
	}

	emailMsg := email.Email{
		From:       constants.SmtpUser,
		To:         []string{requestData.Email},
		Subject:    requestData.StudentNickname + "的求职辅导评估报告",
		Body:       text,
		Images:     images,
		ImageNames: imageNames,
	}

	err := email.SendEmail(emailMsg, constants.SmtpUser, constants.SmtpPass, constants.SmtpHost, constants.SmtpPort)
	if err != nil {
		applogger.Sugar.Error("邮件发送失败: %v\n", err)
		return err
	}
	return nil
}

// toMap 将结构体转换为 map[string]interface{}
func toMap(i model.InterviewCoachingPromotion) map[string]interface{} {
	data := make(map[string]interface{})

	data["studentNickname"] = i.StudentNickname
	data["targetPosition"] = i.TargetPosition
	data["rating"] = scoreToRating(i.InterviewScore)
	data["startTime"] = i.StartTime.UTC().Format("2006年1月2日")
	data["firstName"] = firstCharacter(i.TutorNickname)
	data["tutorNickname"] = i.TutorNickname
	data["totalTime"] = int(i.EndTime.Sub(i.StartTime).Round(time.Minute).Minutes())
	data["interviewScore"] = i.InterviewScore
	data["highlights"] = getHighlightsMap(i.StudentHighlightTags)
	data["studentAdvantages"] = i.StudentAdvantages
	data["tutorSuggestions"] = i.TutorSuggestions
	data["interviewEvaluation"] = i.InterviewEvaluation
	data["platformCertificationStatus"] = i.PlatformCertificationStatus
	data["resumeEvaluation"] = i.ResumeEvaluation

	return data
}

func getHighlightsMap(tags []string) map[string][]string {
	result := make(map[string][]string)
	categoryCounts := make(map[string]int)

	for _, tag := range tags {
		if category, exists := highlightsMap[tag]; exists {
			result[category] = append(result[category], tag)
			categoryCounts[category]++
		} else {
			result["其他"] = append(result["其他"], tag)
			categoryCounts["其他"]++
		}
	}

	// 如果分类数量超过3个，进行筛选
	if len(result) > 3 {
		type CategoryCount struct {
			Name  string
			Count int
		}

		var categories []CategoryCount
		for name, count := range categoryCounts {
			categories = append(categories, CategoryCount{name, count})
		}

		sort.Slice(categories, func(i, j int) bool {
			return categories[i].Count > categories[j].Count
		})

		top3Categories := make(map[string]bool)
		for i, cat := range categories {
			if i >= 3 {
				break
			}
			top3Categories[cat.Name] = true
		}

		// 删除不在前三位的分类
		for category := range result {
			if !top3Categories[category] {
				delete(result, category)
			}
		}
	}

	return result
}

// ScoreToRating 将0-100分的分数转换为评级
// 设计原则：80分以上细分4个等级，60分以下采用鼓励性表述
// 在模板中低于60不显示
func scoreToRating(score int) string {
	switch {
	case score >= 95:
		return "卓越"
	case score >= 90:
		return "优秀"
	case score >= 85:
		return "良好+"
	case score >= 80:
		return "良好"
	case score >= 75:
		return "中等偏上+"
	case score >= 70:
		return "中等偏上"
	case score >= 65:
		return "合格+"
	case score >= 60:
		return "合格"
	case score >= 55:
		return "接近合格"
	case score >= 50:
		return "有待提高"
	case score >= 40:
		return "有一定基础"
	case score >= 30:
		return "基础薄弱"
	case score >= 20:
		return "需要努力"
	default:
		return "继续加油"
	}
}

func firstCharacter(s string) string {
	runes := []rune(s)

	if len(runes) == 0 {
		return ""
	}

	return string(runes[0])
}
