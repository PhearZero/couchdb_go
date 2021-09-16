// Code generated by go-swagger; DO NOT EDIT.

package database

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/rossmerr/couchdb_go/models"
)

// NewSbSecurityPutParams creates a new SbSecurityPutParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSbSecurityPutParams() *SbSecurityPutParams {
	return &SbSecurityPutParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSbSecurityPutParamsWithTimeout creates a new SbSecurityPutParams object
// with the ability to set a timeout on a request.
func NewSbSecurityPutParamsWithTimeout(timeout time.Duration) *SbSecurityPutParams {
	return &SbSecurityPutParams{
		timeout: timeout,
	}
}

// NewSbSecurityPutParamsWithContext creates a new SbSecurityPutParams object
// with the ability to set a context for a request.
func NewSbSecurityPutParamsWithContext(ctx context.Context) *SbSecurityPutParams {
	return &SbSecurityPutParams{
		Context: ctx,
	}
}

// NewSbSecurityPutParamsWithHTTPClient creates a new SbSecurityPutParams object
// with the ability to set a custom HTTPClient for a request.
func NewSbSecurityPutParamsWithHTTPClient(client *http.Client) *SbSecurityPutParams {
	return &SbSecurityPutParams{
		HTTPClient: client,
	}
}

/* SbSecurityPutParams contains all the parameters to send to the API endpoint
   for the sb security put operation.

   Typically these are written to a http.Request.
*/
type SbSecurityPutParams struct {

	// Body.
	Body *models.Body5

	/* Db.

	   Database name
	*/
	Db string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the sb security put params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SbSecurityPutParams) WithDefaults() *SbSecurityPutParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the sb security put params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SbSecurityPutParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the sb security put params
func (o *SbSecurityPutParams) WithTimeout(timeout time.Duration) *SbSecurityPutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the sb security put params
func (o *SbSecurityPutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the sb security put params
func (o *SbSecurityPutParams) WithContext(ctx context.Context) *SbSecurityPutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the sb security put params
func (o *SbSecurityPutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the sb security put params
func (o *SbSecurityPutParams) WithHTTPClient(client *http.Client) *SbSecurityPutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the sb security put params
func (o *SbSecurityPutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the sb security put params
func (o *SbSecurityPutParams) WithBody(body *models.Body5) *SbSecurityPutParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the sb security put params
func (o *SbSecurityPutParams) SetBody(body *models.Body5) {
	o.Body = body
}

// WithDb adds the db to the sb security put params
func (o *SbSecurityPutParams) WithDb(db string) *SbSecurityPutParams {
	o.SetDb(db)
	return o
}

// SetDb adds the db to the sb security put params
func (o *SbSecurityPutParams) SetDb(db string) {
	o.Db = db
}

// WriteToRequest writes these params to a swagger request
func (o *SbSecurityPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param db
	if err := r.SetPathParam("db", o.Db); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
