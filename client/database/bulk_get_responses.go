// Code generated by go-swagger; DO NOT EDIT.

package database

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/RossMerr/couchdb_go/models"
)

// BulkGetReader is a Reader for the BulkGet structure.
type BulkGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BulkGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBulkGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewBulkGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewBulkGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewBulkGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewBulkGetUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewBulkGetOK creates a BulkGetOK with default headers values
func NewBulkGetOK() *BulkGetOK {
	return &BulkGetOK{}
}

/*BulkGetOK handles this case with default header values.

Request completed successfully
*/
type BulkGetOK struct {
	Payload *models.Results
}

func (o *BulkGetOK) Error() string {
	return fmt.Sprintf("[POST /{db}/_bulk_get][%d] bulkGetOK  %+v", 200, o.Payload)
}

func (o *BulkGetOK) GetPayload() *models.Results {
	return o.Payload
}

func (o *BulkGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Results)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkGetBadRequest creates a BulkGetBadRequest with default headers values
func NewBulkGetBadRequest() *BulkGetBadRequest {
	return &BulkGetBadRequest{}
}

/*BulkGetBadRequest handles this case with default header values.

The request provided invalid JSON data or invalid query parameter
*/
type BulkGetBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *BulkGetBadRequest) Error() string {
	return fmt.Sprintf("[POST /{db}/_bulk_get][%d] bulkGetBadRequest  %+v", 400, o.Payload)
}

func (o *BulkGetBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *BulkGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkGetUnauthorized creates a BulkGetUnauthorized with default headers values
func NewBulkGetUnauthorized() *BulkGetUnauthorized {
	return &BulkGetUnauthorized{}
}

/*BulkGetUnauthorized handles this case with default header values.

Read permission required
*/
type BulkGetUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *BulkGetUnauthorized) Error() string {
	return fmt.Sprintf("[POST /{db}/_bulk_get][%d] bulkGetUnauthorized  %+v", 401, o.Payload)
}

func (o *BulkGetUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *BulkGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkGetNotFound creates a BulkGetNotFound with default headers values
func NewBulkGetNotFound() *BulkGetNotFound {
	return &BulkGetNotFound{}
}

/*BulkGetNotFound handles this case with default header values.

Requested database not found
*/
type BulkGetNotFound struct {
	Payload *models.ErrorResponse
}

func (o *BulkGetNotFound) Error() string {
	return fmt.Sprintf("[POST /{db}/_bulk_get][%d] bulkGetNotFound  %+v", 404, o.Payload)
}

func (o *BulkGetNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *BulkGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkGetUnsupportedMediaType creates a BulkGetUnsupportedMediaType with default headers values
func NewBulkGetUnsupportedMediaType() *BulkGetUnsupportedMediaType {
	return &BulkGetUnsupportedMediaType{}
}

/*BulkGetUnsupportedMediaType handles this case with default header values.

Bad Content-Type value
*/
type BulkGetUnsupportedMediaType struct {
	Payload *models.ErrorResponse
}

func (o *BulkGetUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /{db}/_bulk_get][%d] bulkGetUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *BulkGetUnsupportedMediaType) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *BulkGetUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*BulkGetBody bulk get body
swagger:model BulkGetBody
*/
type BulkGetBody struct {

	// docs
	Docs []*models.BasicDoc `json:"docs"`
}

// Validate validates this bulk get body
func (o *BulkGetBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDocs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *BulkGetBody) validateDocs(formats strfmt.Registry) error {

	if swag.IsZero(o.Docs) { // not required
		return nil
	}

	for i := 0; i < len(o.Docs); i++ {
		if swag.IsZero(o.Docs[i]) { // not required
			continue
		}

		if o.Docs[i] != nil {
			if err := o.Docs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("body" + "." + "docs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *BulkGetBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *BulkGetBody) UnmarshalBinary(b []byte) error {
	var res BulkGetBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
