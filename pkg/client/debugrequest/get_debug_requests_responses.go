// Code generated by go-swagger; DO NOT EDIT.

package debugrequest

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/solo-io/squash/pkg/models"
)

// GetDebugRequestsReader is a Reader for the GetDebugRequests structure.
type GetDebugRequestsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDebugRequestsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDebugRequestsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 422:
		result := NewGetDebugRequestsUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDebugRequestsOK creates a GetDebugRequestsOK with default headers values
func NewGetDebugRequestsOK() *GetDebugRequestsOK {
	return &GetDebugRequestsOK{}
}

/*GetDebugRequestsOK handles this case with default header values.

OK
*/
type GetDebugRequestsOK struct {
	Payload models.GetDebugRequestsOKBody
}

func (o *GetDebugRequestsOK) Error() string {
	return fmt.Sprintf("[GET /debugrequest][%d] getDebugRequestsOK  %+v", 200, o.Payload)
}

func (o *GetDebugRequestsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDebugRequestsUnprocessableEntity creates a GetDebugRequestsUnprocessableEntity with default headers values
func NewGetDebugRequestsUnprocessableEntity() *GetDebugRequestsUnprocessableEntity {
	return &GetDebugRequestsUnprocessableEntity{}
}

/*GetDebugRequestsUnprocessableEntity handles this case with default header values.

Validation exception
*/
type GetDebugRequestsUnprocessableEntity struct {
}

func (o *GetDebugRequestsUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /debugrequest][%d] getDebugRequestsUnprocessableEntity ", 422)
}

func (o *GetDebugRequestsUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
