// Code generated by go-swagger; DO NOT EDIT.

package partition

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/rossmerr/couchdb_go/models"
)

// PartitionDesignDocViewReader is a Reader for the PartitionDesignDocView structure.
type PartitionDesignDocViewReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PartitionDesignDocViewReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPartitionDesignDocViewOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewPartitionDesignDocViewNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewPartitionDesignDocViewBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPartitionDesignDocViewUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPartitionDesignDocViewNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPartitionDesignDocViewOK creates a PartitionDesignDocViewOK with default headers values
func NewPartitionDesignDocViewOK() *PartitionDesignDocViewOK {
	return &PartitionDesignDocViewOK{}
}

/*PartitionDesignDocViewOK handles this case with default header values.

Request completed successfully
*/
type PartitionDesignDocViewOK struct {
	/*Response signature
	 */
	ETag string
	/*chunked
	 */
	TransferEncoding string

	Payload *models.Pagination
}

func (o *PartitionDesignDocViewOK) Error() string {
	return fmt.Sprintf("[GET /{db}/_partition/{partition}/_design/{ddoc}/_view/{view}][%d] partitionDesignDocViewOK  %+v", 200, o.Payload)
}

func (o *PartitionDesignDocViewOK) GetPayload() *models.Pagination {
	return o.Payload
}

func (o *PartitionDesignDocViewOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header ETag
	o.ETag = response.GetHeader("ETag")

	// response header Transfer-Encoding
	o.TransferEncoding = response.GetHeader("Transfer-Encoding")

	o.Payload = new(models.Pagination)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPartitionDesignDocViewNotModified creates a PartitionDesignDocViewNotModified with default headers values
func NewPartitionDesignDocViewNotModified() *PartitionDesignDocViewNotModified {
	return &PartitionDesignDocViewNotModified{}
}

/*PartitionDesignDocViewNotModified handles this case with default header values.

Document wasn’t modified since specified revision
*/
type PartitionDesignDocViewNotModified struct {
}

func (o *PartitionDesignDocViewNotModified) Error() string {
	return fmt.Sprintf("[GET /{db}/_partition/{partition}/_design/{ddoc}/_view/{view}][%d] partitionDesignDocViewNotModified ", 304)
}

func (o *PartitionDesignDocViewNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPartitionDesignDocViewBadRequest creates a PartitionDesignDocViewBadRequest with default headers values
func NewPartitionDesignDocViewBadRequest() *PartitionDesignDocViewBadRequest {
	return &PartitionDesignDocViewBadRequest{}
}

/*PartitionDesignDocViewBadRequest handles this case with default header values.

Invalid request
*/
type PartitionDesignDocViewBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *PartitionDesignDocViewBadRequest) Error() string {
	return fmt.Sprintf("[GET /{db}/_partition/{partition}/_design/{ddoc}/_view/{view}][%d] partitionDesignDocViewBadRequest  %+v", 400, o.Payload)
}

func (o *PartitionDesignDocViewBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PartitionDesignDocViewBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPartitionDesignDocViewUnauthorized creates a PartitionDesignDocViewUnauthorized with default headers values
func NewPartitionDesignDocViewUnauthorized() *PartitionDesignDocViewUnauthorized {
	return &PartitionDesignDocViewUnauthorized{}
}

/*PartitionDesignDocViewUnauthorized handles this case with default header values.

Read permission required
*/
type PartitionDesignDocViewUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *PartitionDesignDocViewUnauthorized) Error() string {
	return fmt.Sprintf("[GET /{db}/_partition/{partition}/_design/{ddoc}/_view/{view}][%d] partitionDesignDocViewUnauthorized  %+v", 401, o.Payload)
}

func (o *PartitionDesignDocViewUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PartitionDesignDocViewUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPartitionDesignDocViewNotFound creates a PartitionDesignDocViewNotFound with default headers values
func NewPartitionDesignDocViewNotFound() *PartitionDesignDocViewNotFound {
	return &PartitionDesignDocViewNotFound{}
}

/*PartitionDesignDocViewNotFound handles this case with default header values.

Specified database, design document or view is missed
*/
type PartitionDesignDocViewNotFound struct {
	Payload *models.ErrorResponse
}

func (o *PartitionDesignDocViewNotFound) Error() string {
	return fmt.Sprintf("[GET /{db}/_partition/{partition}/_design/{ddoc}/_view/{view}][%d] partitionDesignDocViewNotFound  %+v", 404, o.Payload)
}

func (o *PartitionDesignDocViewNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PartitionDesignDocViewNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
