package models

type ErrorHandler struct {
	Errors map[string]string
}

func (this *ErrorHandler) HandleErr(err string) {
	this.Errors = make(map[string]string)
	this.Errors["Error"] = err
}
