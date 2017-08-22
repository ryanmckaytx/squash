package debugconfig

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/solo-io/squash/pkg/models"
)

// NewAddDebugConfigParams creates a new AddDebugConfigParams object
// with the default values initialized.
func NewAddDebugConfigParams() AddDebugConfigParams {
	var ()
	return AddDebugConfigParams{}
}

// AddDebugConfigParams contains all the bound params for the add debug config operation
// typically these are obtained from a http.Request
//
// swagger:parameters addDebugConfig
type AddDebugConfigParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*DebugConfig object that needs to be added
	  Required: true
	  In: body
	*/
	Body *models.DebugConfig
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *AddDebugConfigParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.DebugConfig
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("body", "body"))
			} else {
				res = append(res, errors.NewParseError("body", "body", "", err))
			}

		} else {
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Body = &body
			}
		}

	} else {
		res = append(res, errors.Required("body", "body"))
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
