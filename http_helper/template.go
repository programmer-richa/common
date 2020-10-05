package http_helper

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"bitbucket.org/schawla34/commons/custom_error"
)

func WriteJSON(w http.ResponseWriter, js []byte) {
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write(js)
	if err != nil {
		custom_error.Logger(custom_error.Error, err)
	}
}

func GetTemplateHtml(tpl *template.Template, filename string, data interface{}) (string, error) {
	var bb bytes.Buffer
	err := tpl.ExecuteTemplate(&bb, filename, data)
	if err != nil {
		return "", err
	}
	return bb.String(), nil
}

func MapToQueryString(m map[string]interface{}) string {
	b := new(bytes.Buffer)
	for key, value := range m {
		fmt.Fprintf(b, "%s=%v&", key, value)
	}
	s := b.String()
	s = strings.Trim(s, "&")
	return s
}
