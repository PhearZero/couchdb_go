// Code generated by go-swagger; DO NOT EDIT.

package document

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/RossMerr/couchdb_go/models"
)

// DocPutReader is a Reader for the DocPut structure.
type DocPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DocPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDocPutCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewDocPutAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDocPutBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDocPutUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDocPutNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDocPutConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDocPutCreated creates a DocPutCreated with default headers values
func NewDocPutCreated() *DocPutCreated {
	return &DocPutCreated{}
}

/*DocPutCreated handles this case with default header values.

Document created and stored on disk
*/
type DocPutCreated struct {
	/*Double quoted document’s revision token
	 */
	ETag string
	/*Document URI
	 */
	Location strfmt.URI

	Payload *models.DocumentOK
}

func (o *DocPutCreated) Error() string {
	return fmt.Sprintf("[PUT /{db}/{docid}][%d] docPutCreated  %+v", 201, o.Payload)
}

func (o *DocPutCreated) GetPayload() *models.DocumentOK {
	return o.Payload
}

func (o *DocPutCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header ETag
	o.ETag = response.GetHeader("ETag")

	// response header Location

	location, err := formats.Parse("uri", response.GetHeader("Location"))
	if err != nil {
		return errors.InvalidType("Location", "header", "strfmt.URI", response.GetHeader("Location"))
	}
	o.Location = *(location.(*strfmt.URI))

	o.Payload = new(models.DocumentOK)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDocPutAccepted creates a DocPutAccepted with default headers values
func NewDocPutAccepted() *DocPutAccepted {
	return &DocPutAccepted{}
}

/*DocPutAccepted handles this case with default header values.

Document data accepted, but not yet stored on disk
*/
type DocPutAccepted struct {
	/*Double quoted document’s revision token
	 */
	ETag string
	/*chunked. Available if requested with query parameter open_revs
	 */
	TransferEncoding string

	Payload *models.DocumentOK
}

func (o *DocPutAccepted) Error() string {
	return fmt.Sprintf("[PUT /{db}/{docid}][%d] docPutAccepted  %+v", 202, o.Payload)
}

func (o *DocPutAccepted) GetPayload() *models.DocumentOK {
	return o.Payload
}

func (o *DocPutAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header ETag
	o.ETag = response.GetHeader("ETag")

	// response header Transfer-Encoding
	o.TransferEncoding = response.GetHeader("Transfer-Encoding")

	o.Payload = new(models.DocumentOK)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDocPutBadRequest creates a DocPutBadRequest with default headers values
func NewDocPutBadRequest() *DocPutBadRequest {
	return &DocPutBadRequest{}
}

/*DocPutBadRequest handles this case with default header values.

Invalid request body or parameters
*/
type DocPutBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *DocPutBadRequest) Error() string {
	return fmt.Sprintf("[PUT /{db}/{docid}][%d] docPutBadRequest  %+v", 400, o.Payload)
}

func (o *DocPutBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DocPutBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDocPutUnauthorized creates a DocPutUnauthorized with default headers values
func NewDocPutUnauthorized() *DocPutUnauthorized {
	return &DocPutUnauthorized{}
}

/*DocPutUnauthorized handles this case with default header values.

Write privileges required
*/
type DocPutUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *DocPutUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /{db}/{docid}][%d] docPutUnauthorized  %+v", 401, o.Payload)
}

func (o *DocPutUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DocPutUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDocPutNotFound creates a DocPutNotFound with default headers values
func NewDocPutNotFound() *DocPutNotFound {
	return &DocPutNotFound{}
}

/*DocPutNotFound handles this case with default header values.

Specified database or document ID doesn’t exists
*/
type DocPutNotFound struct {
	Payload *models.ErrorResponse
}

func (o *DocPutNotFound) Error() string {
	return fmt.Sprintf("[PUT /{db}/{docid}][%d] docPutNotFound  %+v", 404, o.Payload)
}

func (o *DocPutNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DocPutNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDocPutConflict creates a DocPutConflict with default headers values
func NewDocPutConflict() *DocPutConflict {
	return &DocPutConflict{}
}

/*DocPutConflict handles this case with default header values.

Document with the specified ID already exists or specified revision is not latest for target document
*/
type DocPutConflict struct {
	Payload *models.ErrorResponse
}

func (o *DocPutConflict) Error() string {
	return fmt.Sprintf("[PUT /{db}/{docid}][%d] docPutConflict  %+v", 409, o.Payload)
}

func (o *DocPutConflict) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DocPutConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
