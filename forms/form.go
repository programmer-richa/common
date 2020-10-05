package forms

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
