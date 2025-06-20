package htmlkit

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os/exec"
	"strings"

	"jihulab.com/guashu/gs-server/util/constants"
)

type HTMLToImageOptions struct {
	WkhtmltoimagePath string // 指定 wkhtmltoimage 可执行文件路径
	Format            string
	Quality           int
	Width             int
	Height            int
	ZoomFactor        float64
	JavaScriptDelay   int
	Debug             bool
	FullPage          bool
	SmartWidth        bool // 是否启用智能宽度（尝试根据内容调整）
}

// 默认配置
func defaultOptions() HTMLToImageOptions {
	return HTMLToImageOptions{
		WkhtmltoimagePath: constants.WkhtmltoimagePath, // 默认从 PATH 查找
		Format:            "png",
		Quality:           90,
		Width:             1280,
		Height:            0, // 自动高度
		ZoomFactor:        1.0,
		JavaScriptDelay:   200,
		Debug:             false,
		FullPage:          true,  // 默认捕获整个页面
		SmartWidth:        false, // 禁用智能宽度，使用固定宽度
	}
}

// options: 配置选项 (可为 nil，使用默认值)
func HTMLToImageBytes(htmlContent string, options *HTMLToImageOptions) ([]byte, error) {
	if options == nil {
		opts := defaultOptions()
		options = &opts
	}

	supportedFormats := map[string]bool{
		"png": true, "jpg": true, "jpeg": true, "pdf": true, "bmp": true, "svg": true,
	}
	if !supportedFormats[strings.ToLower(options.Format)] {
		return nil, errors.New("不支持的输出格式，支持的格式: png, jpg, pdf, bmp, svg")
	}

	args := []string{
		"--format", options.Format,
		"--quality", fmt.Sprintf("%d", options.Quality),
		"--width", fmt.Sprintf("%d", options.Width),
	}

	args = append(args, "--height", fmt.Sprintf("%d", options.Height))

	if options.ZoomFactor != 1.0 {
		args = append(args, "--zoom", fmt.Sprintf("%.1f", options.ZoomFactor))
	}

	if options.JavaScriptDelay > 0 {
		args = append(args, "--javascript-delay", fmt.Sprintf("%d", options.JavaScriptDelay))
	}

	if options.FullPage {
		args = append(args, "--disable-smart-width")
	}

	if options.SmartWidth {
		args = append(args, "--enable-smart-width")
	} else {
		args = append(args, "--disable-smart-width")
	}

	if !options.Debug {
		args = append(args, "--quiet")
	}

	args = append(args, "-", "-")

	cmd := exec.Command(options.WkhtmltoimagePath, args...)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, fmt.Errorf("创建标准输入管道失败: %v", err)
	}

	go func() {
		defer stdin.Close()
		io.WriteString(stdin, htmlContent)
	}()

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("转换失败: %v, 错误详情: %s", err, stderr.String())
	}

	return stdout.Bytes(), nil
}
