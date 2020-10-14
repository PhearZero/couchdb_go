// Code generated by go-swagger; DO NOT EDIT.

package design_documents

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

// DesignDocPutReader is a Reader for the DesignDocPut structure.
type DesignDocPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DesignDocPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDesignDocPutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDesignDocPutNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDesignDocPutOK creates a DesignDocPutOK with default headers values
func NewDesignDocPutOK() *DesignDocPutOK {
	return &DesignDocPutOK{}
}

/*DesignDocPutOK handles this case with default header values.

Request completed successfully
*/
type DesignDocPutOK struct {
	/*Double quoted document’s revision token
	 */
	ETag string
	/*Document URI
	 */
	Location strfmt.URI

	Payload *models.Pagination
}

func (o *DesignDocPutOK) Error() string {
	return fmt.Sprintf("[PUT /{db}/_design/{ddoc}][%d] designDocPutOK  %+v", 200, o.Payload)
}

func (o *DesignDocPutOK) GetPayload() *models.Pagination {
	return o.Payload
}

func (o *DesignDocPutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header ETag
	o.ETag = response.GetHeader("ETag")

	// response header Location

	location, err := formats.Parse("uri", response.GetHeader("Location"))
	if err != nil {
		return errors.InvalidType("Location", "header", "strfmt.URI", response.GetHeader("Location"))
	}
	o.Location = *(location.(*strfmt.URI))

	o.Payload = new(models.Pagination)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDesignDocPutNotFound creates a DesignDocPutNotFound with default headers values
func NewDesignDocPutNotFound() *DesignDocPutNotFound {
	return &DesignDocPutNotFound{}
}

/*DesignDocPutNotFound handles this case with default header values.

Requested database not found
*/
type DesignDocPutNotFound struct {
	Payload *models.ErrorResponse
}

func (o *DesignDocPutNotFound) Error() string {
	return fmt.Sprintf("[PUT /{db}/_design/{ddoc}][%d] designDocPutNotFound  %+v", 404, o.Payload)
}

func (o *DesignDocPutNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DesignDocPutNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
