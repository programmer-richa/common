package forms

import (
	"fmt"
	"html/template"
	"net/http"
	"testing"

	"github.com/programmer-richa/common/constants"
	"github.com/programmer-richa/common/http_helper"
	"github.com/programmer-richa/common/validator"
)

func testForm() *Form {
	// initialising form
	form := NewForm("registration", "/registration", http.MethodPost)
	username := NewFieldBuilder(Text, "Name").
		Required().
		ValidatorFunc(ReValidator(validator.NamePattern, constants.InvalidName))
	email := username.DeepCopy().
		Label("Email address").
		FieldType(Email).
		ValidatorFunc(ReValidator(validator.EmailPattern, constants.InvalidEmail))

	password := username.DeepCopy().
		Label("Password").
		FieldType(Password).
		ValidatorFunc(PasswordValidator(constants.InvalidPassword))
	cpassword := username.DeepCopy().
		Label("Confirm Password").
		FieldType(Password).
		ValidatorFunc(PasswordValidator(constants.InvalidPassword))

	// username
	form.WithField(form.PrefixFieldName+"Name", username)

	// email
	form.WithField(form.PrefixFieldName+"Email", email)

	// password
	form.WithField(form.PrefixFieldName+"Password", password)

	// password
	form.WithField(form.PrefixFieldName+"CPassword", cpassword)

	// Subscribe
	form.WithField(form.PrefixFieldName+"Subscribe",
		NewFieldBuilder(Checkbox, "Subscribe to email alerts."))

	// Submit
	form.WithField(form.PrefixFieldName+"Submit",
		NewFieldBuilder(Button, "Register Now"))

	return form
}

func TestInvalidForm(t *testing.T) {
	form := testForm()
	tpl := template.Must(template.ParseGlob("templates/*.gohtml"))
	str, _ := http_helper.GetTemplateHtml(tpl, "form.gohtml", form)
	// fmt.Println(err)
	fmt.Println(str)
}
