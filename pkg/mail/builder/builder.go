package builder

import (
	"bytes"
	"fmt"
	"html/template"
)

const templatePath = "pkg/mail/builder/template/html-email-forms-main/index.html"

type EmailBuilder struct {
	templatePath string
	data         EmailData
}

type EmailData struct {
	FIO         string
	PhoneNumber string
	Email       string
	Address     string
	Comment     string
	LeechSize1  int
	LeechSize2  int
	LeechSize3  int
	TotalCount  int
	PackageType int
	TotalPrice  float64
}

// NewEmailBuilder создает новый экземпляр EmailBuilder
func NewEmailBuilder() *EmailBuilder {
	return &EmailBuilder{
		templatePath: templatePath,
		data:         EmailData{},
	}
}

// Build генерирует HTML письмо на основе шаблона и данных
func (b *EmailBuilder) Build() (string, error) {
	tmpl, err := template.ParseFiles(b.templatePath)
	if err != nil {
		return "", fmt.Errorf("failed to parse template: %w", err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, b.data); err != nil {
		return "", fmt.Errorf("failed to execute template: %w", err)
	}

	return buf.String(), nil
}
