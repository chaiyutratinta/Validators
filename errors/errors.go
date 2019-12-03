package errors

type Errors []Error

type Error struct {
	FieldName string
	Rule      string
	ErrMsg    string
}

type TmpErrors struct {
	Errs Errors
	Tmp  Error
}

type ErrorsBuilder interface {
	Field(string) ErrorsBuilder
	Rule(string) ErrorsBuilder
	ErrMsg(string) ErrorsBuilder
	Set()
	GetError() Errors
}

func New() ErrorsBuilder {
	return &TmpErrors{}
}

func (e *TmpErrors) Field(f string) ErrorsBuilder {
	e.Tmp.FieldName = f

	return e
}

func (e *TmpErrors) Rule(r string) ErrorsBuilder {
	e.Tmp.Rule = r

	return e
}

func (e *TmpErrors) ErrMsg(err string) ErrorsBuilder {
	e.Tmp.ErrMsg = err

	return e
}

func (e *TmpErrors) Set() {
	e.Errs = append(e.Errs, e.Tmp)
}

func (e *TmpErrors) GetError() Errors {

	return e.Errs
}
