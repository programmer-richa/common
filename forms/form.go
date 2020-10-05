package forms

import "net/url"

type Form struct {
	Fields          map[string]*Field
	FieldNames      []string
	Errors          map[string]error
	Values          FormValues
	validator       FormValidator
	PrefixFieldName string
	EncryptType     string
	Action          string
	Method          string
	Name            string
	Title           string
	DisplayInModal  bool
}

// Factory function to create form

func NewForm(name string, action string, method string) *Form {
	return &Form{Name: name,
		PrefixFieldName: name,
		Action:          action,
		Method:          method,
		Fields:          make(map[string]*Field),
		FieldNames:      make([]string, 0),
	}
}

// Builder implementation for setting values of form.
// WithTitle adds the title to the form.
func (f *Form) WithTitle(title string) *Form {
	f.Title = title
	return f
}

// ShowInModal specifies that form is in model.
func (f *Form) ShowInModal() *Form {
	f.DisplayInModal = true
	return f
}

// WithField adds the Field produced by the FieldBuilder to the Form under the given name.
func (f *Form) WithField(name string, fb *FieldBuilder) *Form {
	fb.Name(name)
	field := fb.Build()
	f.Fields[name] = field
	f.FieldNames = append(f.FieldNames, name)
	return f
}

// WithValidator adds the FormValidator to the form.
func (f *Form) WithValidator(validator FormValidator) *Form {
	f.validator = validator
	return f
}

// Valid validates every field followed by the form's validator if provided.
func (f *Form) Valid(postForm url.Values) bool {
	valid := true

	f.Errors = nil
	f.Values = nil

	formValues := make(FormValues)
	formErrors := make(map[string]error)

	// validate fields
	for _, fname := range f.FieldNames {
		fieldValue, fieldError := f.Fields[fname].Validate(postForm.Get(fname))
		if fieldError != nil {
			valid = false
			formErrors[fname] = fieldError
		} else {
			formValues[fname] = fieldValue
		}
	}

	// validate form
	if valid && f.validator != nil {
		valid = f.validator(&formValues)
	}

	// if its valid, make the values available
	// otherwise make the errors available
	if valid {
		f.Values = formValues
	} else {
		f.Errors = formErrors
	}

	return valid
}
