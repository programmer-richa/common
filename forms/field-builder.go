package forms

// fieldModifier has a type of function that accepts a *Field
// functional builder
type fieldModifier func(*Field)

// FieldBuilder holds a list of functions that accepts a *Field
type FieldBuilder struct {
	actions []fieldModifier
}

func (fb *FieldBuilder) Name(name string) *FieldBuilder {
	// append an anonymous function that specifies name of a field
	fb.actions = append(fb.actions, func(f *Field) {
		f.Name = name
	})
	return fb
}

func (fb *FieldBuilder) Placeholder(placeholder string) *FieldBuilder {
	// append an anonymous function that specifies placeholder of a field
	fb.actions = append(fb.actions, func(f *Field) {
		f.Placeholder = placeholder
	})
	return fb
}

func (fb *FieldBuilder) Formatters(formatterFunc FormatterFunc) *FieldBuilder {
	// append an anonymous function that specifies formatterFunc of a field
	fb.actions = append(fb.actions, func(f *Field) {
		f.Formatters = append(f.Formatters, formatterFunc)
	})
	return fb
}
func (fb *FieldBuilder) ValidatorFunc(validatorFunc ValidatorFunc) *FieldBuilder {
	// append an anonymous function that specifies validatorFunc of a field
	fb.actions = append(fb.actions, func(f *Field) {
		f.Validators = append(f.Validators, validatorFunc)
	})
	return fb
}
func (fb *FieldBuilder) FieldNames(name string, fbuild FieldBuilder) *FieldBuilder {
	// append an anonymous function that specifies name of a fields
	fb.actions = append(fb.actions, func(f *Field) {
		field := fbuild.Build()
		f.Fields[name] = field
		f.FieldNames = append(f.FieldNames, name)
	})
	return fb
}
func (fb *FieldBuilder) JsFunction(name string, event string) *FieldBuilder {
	// append an anonymous function that specifies name of a fields
	fb.actions = append(fb.actions, func(f *Field) {
		f.JsFunctions[name] = event
		f.EventNames = append(f.EventNames, name)
	})
	return fb
}

func (fb *FieldBuilder) Loader(loaderFunc LoaderFunc) *FieldBuilder {
	// append an anonymous function that specifies loaderFunc of a field
	fb.actions = append(fb.actions, func(f *Field) {
		f.Loader = loaderFunc
	})
	return fb
}

func (fb *FieldBuilder) Required() *FieldBuilder {
	// append an anonymous function that specifies required field
	fb.actions = append(fb.actions, func(f *Field) {
		f.Required = true
	})
	return fb
}
func (fb *FieldBuilder) Selected() *FieldBuilder {
	// append an anonymous function that specifies selected  field
	fb.actions = append(fb.actions, func(f *Field) {
		f.Selected = true
	})
	return fb
}
func (fb *FieldBuilder) Disabled() *FieldBuilder {
	// append an anonymous function that specifies disabled  field
	fb.actions = append(fb.actions, func(f *Field) {
		f.Disabled = true
	})
	return fb
}
func (fb *FieldBuilder) AllowMultipleValues() *FieldBuilder {
	// append an anonymous function that specifies multi-value  field
	fb.actions = append(fb.actions, func(f *Field) {
		f.AllowMultipleValues = true
	})
	return fb
}
func (fb *FieldBuilder) HideLabel() *FieldBuilder {
	// append an anonymous function that specifies hidden lable of the field
	fb.actions = append(fb.actions, func(f *Field) {
		f.HideLabel = true
	})
	return fb
}

func (fb *FieldBuilder) Empty(empty interface{}) *FieldBuilder {
	// append an anonymous function that specifies default value of a field
	fb.actions = append(fb.actions, func(f *Field) {
		f.Empty = empty
	})
	return fb
}

func (fb *FieldBuilder) Min(min int) *FieldBuilder {
	// append an anonymous function that specifies min value of a field
	fb.actions = append(fb.actions, func(f *Field) {
		f.Min = min
	})
	return fb
}
func (fb *FieldBuilder) Max(max int) *FieldBuilder {
	// append an anonymous function that specifies min value of a field
	fb.actions = append(fb.actions, func(f *Field) {
		f.Max = max
	})
	return fb
}
func (fb *FieldBuilder) Label(label string) *FieldBuilder {
	// append an anonymous function that specifies label of a field
	fb.actions = append(fb.actions, func(f *Field) {
		f.Label = label
	})
	return fb
}
func (fb *FieldBuilder) FieldType(fieldType string) *FieldBuilder {
	// append an anonymous function that specifies type of a field
	fb.actions = append(fb.actions, func(f *Field) {
		f.FieldType = fieldType
	})
	return fb
}
func (fb *FieldBuilder) Value(value string) *FieldBuilder {
	// append an anonymous function that specifies value of a field
	fb.actions = append(fb.actions, func(f *Field) {
		f.Value = value
	})
	return fb
}
func (fb *FieldBuilder) RequiredErrorMsg(msg string) *FieldBuilder {
	// append an anonymous function that specifies error msg of a field
	fb.actions = append(fb.actions, func(f *Field) {
		f.RequiredErrorMsg = msg
	})
	return fb
}
func (fb *FieldBuilder) InvalidDataErrorMsg(msg string) *FieldBuilder {
	// append an anonymous function that specifies error msg of a field
	fb.actions = append(fb.actions, func(f *Field) {
		f.InvalidDataErrorMsg = msg
	})
	return fb
}
func (fb *FieldBuilder) JSValidator(js string) *FieldBuilder {
	// append an anonymous function that specifies error msg of a field
	fb.actions = append(fb.actions, func(f *Field) {
		f.JSValidator = js
	})
	return fb
}
func (fb *FieldBuilder) Width(width int) *FieldBuilder {
	// append an anonymous function that specifies error msg of a field
	fb.actions = append(fb.actions, func(f *Field) {
		if width == 0 {
			width = 12
		}
		f.Width = width
	})
	return fb
}

func (fb *FieldBuilder) Build() *Field {
	// Binds all values to field
	f := Field{}
	f.Fields, f.JsFunctions = make(map[string]*Field), make(map[string]string)
	f.EventNames, f.FieldNames = []string{}, []string{}

	for _, a := range fb.actions {
		a(&f)
	}
	return &f
}
