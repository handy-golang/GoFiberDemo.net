package tmpl

import (
	"embed"
)

//go:embed email.html
var Email string

type EmailParam struct {
	Message string
	SysTime string
}

//go:embed *
var Static embed.FS

//go:embed index/index.html
var IndexHtml []byte
