package models

import (
	"net/http"
	"fmt"
	"github.com/lansanacamara/davillex/util"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

type Error struct {
	Error      error  // The actual error that occurred -- for logging in AppEngine
					  // and development purposes.
	Message    string // A custom error message we provide for UX in production.
	Code       int    // An error code.

	UIDisabled bool   // Set to 'true' to use w.Write instead of custom UIError.
	UIError    string // Custom message passed to whatever view is going to be shown.
	View       string // The requested view.

	Ignore     bool   // Set to 'true' to ignore http.Error and UIError -- still
					  // provide an error log to AppEngine.
}

type Handle func(w http.ResponseWriter, r *http.Request) *Error

func (this Handle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := this(w, r); err != nil {
		// Log on AppEngine
		c := appengine.NewContext(r)
		log.Errorf(c, "%v", err.Error)

		// Check if UIError was disabled.
		if err.UIDisabled {
			// Log the error for development purposes.
			fmt.Printf("%v", err)
			http.Error(w, err.Message, err.Code)
			return
		}

		// Create UIError and pass to requested view.
		form := &Form{}
		form.NewUIError(err.Message)
		util.RenderView(w, err.View, form)
	}
}


