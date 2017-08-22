package debugsessions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/solo-io/squash/pkg/models"
)

// NewPutDebugSessionParams creates a new PutDebugSessionParams object
// with the default values initialized.
func NewPutDebugSessionParams() PutDebugSessionParams {
	var ()
	return PutDebugSessionParams{}
}

// PutDebugSessionParams contains all the bound params for the put debug session operation
// typically these are obtained from a http.Request
//
// swagger:parameters putDebugSession
type PutDebugSessionParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*the debug session.
	  Required: true
	  In: body
	*/
	Body *models.DebugSession
	/*debug config id.
	  Required: true
	  In: path
	*/
	DebugConfigID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *PutDebugSessionParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.DebugSession
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

	rDebugConfigID, rhkDebugConfigID, _ := route.Params.GetOK("debugConfigId")
	if err := o.bindDebugConfigID(rDebugConfigID, rhkDebugConfigID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PutDebugSessionParams) bindDebugConfigID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.DebugConfigID = raw

	return nil
}
