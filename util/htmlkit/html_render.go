package htmlkit

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
)

func RenderHTML(filePath string, data map[string]interface{}) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("无法打开文件: %w", err)
	}
	defer file.Close()

	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("读取文件失败: %w", err)
	}

	tmpl, err := template.New("htmlTemplate").Parse(string(content))
	if err != nil {
		return "", fmt.Errorf("解析模板失败: %w", err)
	}

	var result bytes.Buffer
	err = tmpl.Execute(&result, data)
	if err != nil {
		return "", fmt.Errorf("渲染模板失败: %w", err)
	}

	return result.String(), nil
}
