package models

import (
	"github.com/lansanacamara/davillex/config"
	"github.com/lansanacamara/davillex/util"
)

type Form struct {
	Name       string
	Email      string
	Phone      string
	Message    string
	Thanks     string
	Validation map[string]string
	Error
}

func (this *Form) Validate() bool {
	this.Validation = make(map[string]string)

	matched := util.MatchRegexp(".+@.+\\..+", this.Email)

	if !util.IsEmpty(this.Email) {
		if matched == false {
			this.Validation["Email"] = config.EMAIL_INVALID
		}
	} else {
		this.Validation["Email"] = config.EMAIL_EMPTY
	}

	if util.IsEmpty(this.Name) {
		this.Validation["Name"] = config.NAME_EMPTY
	}

	if util.IsEmpty(this.Phone) {
		this.Validation["Phone"] = config.PHONE_EMPTY
	}

	if util.IsEmpty(this.Message) {
		this.Validation["Message"] = config.MESSAGE_EMPTY
	}

	return len(this.Validation) == 0
}

func (this *Form) NewUIError(message string) {
	this.UIError = message
}