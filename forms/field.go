package forms

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
)

// Field represents an element of a form.
// For instance, text, password field.
type Field struct {
	Name                string
	Placeholder         string
	Formatters          []FormatterFunc
	Loader              LoaderFunc
	Validators          []ValidatorFunc
	Required            bool
	Selected            bool
	Disabled            bool
	AllowMultipleValues bool
	HideLabel           bool
	Empty               interface{}
	Min                 int
	Max                 int
	Label               string
	FieldType           string
	Value               string
	Fields              map[string]*Field
	FieldNames          []string
	Width               int
	JsFunctions         map[string]string
	EventNames          []string
	RequiredErrorMsg    string
	InvalidDataErrorMsg string
	JSValidator         string
}

// Validate tests if the data passed in the field follows the specified validation rules.
func (f *Field) Validate(rawValue string) (interface{}, error) {
	if rawValue == "" {
		if f.Required {
			return nil, errors.New(fmt.Sprintf("%s is required", f.Name))
		} else {
			return f.Empty, nil
		}
	}

	// format raw input
	rawValue = f.format(rawValue)

	// deserialize
	value, err := f.Loader(rawValue)
	if err != nil {
		return nil, err
	}

	// validate
	err = f.validate(value)
	if err != nil {
		return nil, err
	}

	return value, nil
}

// format formats the data
func (f *Field) format(rawValue string) string {
	for _, f := range f.Formatters {
		rawValue = f(rawValue)
	}
	return rawValue
}

// validate tests if the data passed in the field follows the specified validation rules.
func (f *Field) validate(value interface{}) error {
	for _, v := range f.Validators {
		if err := v(value); err != nil {
			return err
		}
	}
	return nil
}

// Creating Copy of Field
// Prototype Design pattern to create copy of form field without initialising
// whole object from scratch
func (f *Field) DeepCopy() *Field {
	// note: no error handling below
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	_ = e.Encode(f)
	// peek into structure
	// fmt.Println(string(b.Bytes()))
	d := gob.NewDecoder(&b)
	result := Field{}
	_ = d.Decode(&result)
	return &result
}
