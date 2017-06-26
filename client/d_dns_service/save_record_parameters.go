package d_dns_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/muka/dyndns/models"
)

// NewSaveRecordParams creates a new SaveRecordParams object
// with the default values initialized.
func NewSaveRecordParams() *SaveRecordParams {
	var ()
	return &SaveRecordParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSaveRecordParamsWithTimeout creates a new SaveRecordParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSaveRecordParamsWithTimeout(timeout time.Duration) *SaveRecordParams {
	var ()
	return &SaveRecordParams{

		timeout: timeout,
	}
}

// NewSaveRecordParamsWithContext creates a new SaveRecordParams object
// with the default values initialized, and the ability to set a context for a request
func NewSaveRecordParamsWithContext(ctx context.Context) *SaveRecordParams {
	var ()
	return &SaveRecordParams{

		Context: ctx,
	}
}

// NewSaveRecordParamsWithHTTPClient creates a new SaveRecordParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSaveRecordParamsWithHTTPClient(client *http.Client) *SaveRecordParams {
	var ()
	return &SaveRecordParams{
		HTTPClient: client,
	}
}

/*SaveRecordParams contains all the parameters to send to the API endpoint
for the save record operation typically these are written to a http.Request
*/
type SaveRecordParams struct {

	/*Body*/
	Body *models.APIRecord

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the save record params
func (o *SaveRecordParams) WithTimeout(timeout time.Duration) *SaveRecordParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the save record params
func (o *SaveRecordParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the save record params
func (o *SaveRecordParams) WithContext(ctx context.Context) *SaveRecordParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the save record params
func (o *SaveRecordParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the save record params
func (o *SaveRecordParams) WithHTTPClient(client *http.Client) *SaveRecordParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the save record params
func (o *SaveRecordParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the save record params
func (o *SaveRecordParams) WithBody(body *models.APIRecord) *SaveRecordParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the save record params
func (o *SaveRecordParams) SetBody(body *models.APIRecord) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *SaveRecordParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body == nil {
		o.Body = new(models.APIRecord)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}