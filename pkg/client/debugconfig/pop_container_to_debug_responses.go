package debugconfig

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/solo-io/squash/pkg/models"
)

// PopContainerToDebugReader is a Reader for the PopContainerToDebug structure.
type PopContainerToDebugReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PopContainerToDebugReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPopContainerToDebugOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPopContainerToDebugBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPopContainerToDebugNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 408:
		result := NewPopContainerToDebugRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPopContainerToDebugOK creates a PopContainerToDebugOK with default headers values
func NewPopContainerToDebugOK() *PopContainerToDebugOK {
	return &PopContainerToDebugOK{}
}

/*PopContainerToDebugOK handles this case with default header values.

OK
*/
type PopContainerToDebugOK struct {
	Payload *models.DebugConfig
}

func (o *PopContainerToDebugOK) Error() string {
	return fmt.Sprintf("[DELETE /debugconfig/platform/containers/{node}/latest][%d] popContainerToDebugOK  %+v", 200, o.Payload)
}

func (o *PopContainerToDebugOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DebugConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPopContainerToDebugBadRequest creates a PopContainerToDebugBadRequest with default headers values
func NewPopContainerToDebugBadRequest() *PopContainerToDebugBadRequest {
	return &PopContainerToDebugBadRequest{}
}

/*PopContainerToDebugBadRequest handles this case with default header values.

Invalid ID supplied
*/
type PopContainerToDebugBadRequest struct {
}

func (o *PopContainerToDebugBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /debugconfig/platform/containers/{node}/latest][%d] popContainerToDebugBadRequest ", 400)
}

func (o *PopContainerToDebugBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPopContainerToDebugNotFound creates a PopContainerToDebugNotFound with default headers values
func NewPopContainerToDebugNotFound() *PopContainerToDebugNotFound {
	return &PopContainerToDebugNotFound{}
}

/*PopContainerToDebugNotFound handles this case with default header values.

Node note not found
*/
type PopContainerToDebugNotFound struct {
}

func (o *PopContainerToDebugNotFound) Error() string {
	return fmt.Sprintf("[DELETE /debugconfig/platform/containers/{node}/latest][%d] popContainerToDebugNotFound ", 404)
}

func (o *PopContainerToDebugNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPopContainerToDebugRequestTimeout creates a PopContainerToDebugRequestTimeout with default headers values
func NewPopContainerToDebugRequestTimeout() *PopContainerToDebugRequestTimeout {
	return &PopContainerToDebugRequestTimeout{}
}

/*PopContainerToDebugRequestTimeout handles this case with default header values.

Timeout reached
*/
type PopContainerToDebugRequestTimeout struct {
}

func (o *PopContainerToDebugRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /debugconfig/platform/containers/{node}/latest][%d] popContainerToDebugRequestTimeout ", 408)
}

func (o *PopContainerToDebugRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
