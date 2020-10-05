package forms

// FormValidator represents a form validator function
type FormValidator func(formValues *FormValues) bool
