//go:build !tmplgen

package tmpl

import "html/template"

type List struct {
	Counter *template.Template
	Hello   *template.Template
	Test    *template.Template
}
