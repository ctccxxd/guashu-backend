package email

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"mime"
	"mime/quotedprintable"
	"net/smtp"
	"strings"
)

// 邮件结构体
type Email struct {
	From       string
	To         []string
	Subject    string
	Body       string
	Images     [][]byte
	ImageNames []string
}

func SendEmail(email Email, smtpUser, smtpPass, smtpHost string, smtpPort int) error {
	msg, err := buildEmailMessage(email)
	if err != nil {
		return err
	}

	addr := fmt.Sprintf("%s:%d", smtpHost, smtpPort)
	auth := smtp.PlainAuth("", smtpUser, smtpPass, smtpHost)

	conn, err := tls.Dial("tcp", addr, &tls.Config{
		ServerName:         smtpHost,
		InsecureSkipVerify: true,
	})
	if err != nil {
		return fmt.Errorf("dial failed: %w", err)
	}
	defer conn.Close()

	c, err := smtp.NewClient(conn, smtpHost)
	if err != nil {
		return fmt.Errorf("new client failed: %w", err)
	}
	defer c.Close()

	if err = c.Auth(auth); err != nil {
		return fmt.Errorf("auth failed: %w", err)
	}

	if err = c.Mail(email.From); err != nil {
		return fmt.Errorf("mail failed: %w", err)
	}

	for _, to := range email.To {
		if err = c.Rcpt(to); err != nil {
			return fmt.Errorf("rcpt failed: %w", err)
		}
	}

	w, err := c.Data()
	if err != nil {
		return fmt.Errorf("data failed: %w", err)
	}
	_, err = w.Write(msg.Bytes())
	if err != nil {
		return fmt.Errorf("write failed: %w", err)
	}
	err = w.Close()
	if err != nil {
		return fmt.Errorf("close failed: %w", err)
	}

	return c.Quit()
}

func buildEmailMessage(email Email) (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)

	fmt.Fprintf(buf, "From: %s\r\n", email.From)
	fmt.Fprintf(buf, "To: %s\r\n", joinEmailList(email.To))
	fmt.Fprintf(buf, "Subject: =?UTF-8?B?%s?=\r\n", base64.StdEncoding.EncodeToString([]byte(email.Subject)))
	fmt.Fprintf(buf, "MIME-Version: 1.0\r\n")
	fmt.Fprintf(buf, "Content-Type: multipart/mixed; boundary=\"%s\"\r\n\r\n", "boundary12345")
	fmt.Fprintf(buf, "--boundary12345\r\n")

	fmt.Fprintf(buf, "Content-Type: text/html; charset=\"UTF-8\"\r\n")
	fmt.Fprintf(buf, "Content-Transfer-Encoding: quoted-printable\r\n\r\n")

	qpw := quotedprintable.NewWriter(buf)
	_, err := qpw.Write([]byte(email.Body))
	if err != nil {
		return nil, err
	}
	qpw.Close()
	fmt.Fprintf(buf, "\r\n--boundary12345\r\n")

	for i, imgData := range email.Images {
		if i >= len(email.ImageNames) {
			continue
		}
		err := addImageAttachment(buf, imgData, email.ImageNames[i])
		if err != nil {
			return nil, err
		}
		fmt.Fprintf(buf, "--boundary12345\r\n")
	}

	fmt.Fprintf(buf, "--boundary12345--\r\n")
	return buf, nil
}

func addImageAttachment(buf *bytes.Buffer, imgData []byte, fileName string) error {
	mimeType := "image/jpeg"
	if strings.HasSuffix(fileName, ".png") {
		mimeType = "image/png"
	} else if strings.HasSuffix(fileName, ".gif") {
		mimeType = "image/gif"
	} else if strings.HasSuffix(fileName, ".webp") {
		mimeType = "image/webp"
	}

	encodedFileName := mime.QEncoding.Encode("UTF-8", fileName)

	fmt.Fprintf(buf, "Content-Type: %s; name=\"%s\"\r\n", mimeType, encodedFileName)
	fmt.Fprintf(buf, "Content-Disposition: attachment; filename=\"%s\"\r\n", encodedFileName)
	fmt.Fprintf(buf, "Content-Transfer-Encoding: base64\r\n\r\n")

	encoder := base64.NewEncoder(base64.StdEncoding, buf)
	_, err := encoder.Write(imgData)
	encoder.Close()
	if err != nil {
		return err
	}
	fmt.Fprintf(buf, "\r\n")
	return nil
}

func joinEmailList(emails []string) string {
	var result string
	for i, email := range emails {
		if i > 0 {
			result += ", "
		}
		result += email
	}
	return result
}
