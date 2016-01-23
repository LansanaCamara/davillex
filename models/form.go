package models

import (
	"github.com/lansanacamara/davillex/util"
	"github.com/lansanacamara/davillex/config"
)

type Form struct {
	Name    string
	Email   string
	Phone   string
	Message string
	Thanks  string
	ErrorHandler
}

func (this *Form) Validate() bool {
	this.Errors = make(map[string]string)

	matched := util.MatchRegexp(".+@.+\\..+", this.Email)

	if !util.IsEmpty(this.Email) {
		if matched == false {
			this.Errors["Email"] = config.EMAIL_INVALID
		}
	} else {
		this.Errors["Email"] = config.EMAIL_EMPTY
	}

	if util.IsEmpty(this.Name) {
		this.Errors["Name"] = config.NAME_EMPTY
	}

	if util.IsEmpty(this.Phone) {
		this.Errors["Phone"] = config.PHONE_EMPTY
	}

	if util.IsEmpty(this.Message) {
		this.Errors["Message"] = config.MESSAGE_EMPTY
	}

	return len(this.Errors) == 0
}
