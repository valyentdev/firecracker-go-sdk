// Code generated by go-swagger; DO NOT EDIT.

// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
// 	http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/valyentdev/firecracker-go-sdk/client/models"
)

// NewPutEntropyDeviceParams creates a new PutEntropyDeviceParams object
// with the default values initialized.
func NewPutEntropyDeviceParams() *PutEntropyDeviceParams {
	var ()
	return &PutEntropyDeviceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutEntropyDeviceParamsWithTimeout creates a new PutEntropyDeviceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutEntropyDeviceParamsWithTimeout(timeout time.Duration) *PutEntropyDeviceParams {
	var ()
	return &PutEntropyDeviceParams{

		timeout: timeout,
	}
}

// NewPutEntropyDeviceParamsWithContext creates a new PutEntropyDeviceParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutEntropyDeviceParamsWithContext(ctx context.Context) *PutEntropyDeviceParams {
	var ()
	return &PutEntropyDeviceParams{

		Context: ctx,
	}
}

// NewPutEntropyDeviceParamsWithHTTPClient creates a new PutEntropyDeviceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutEntropyDeviceParamsWithHTTPClient(client *http.Client) *PutEntropyDeviceParams {
	var ()
	return &PutEntropyDeviceParams{
		HTTPClient: client,
	}
}

/*PutEntropyDeviceParams contains all the parameters to send to the API endpoint
for the put entropy device operation typically these are written to a http.Request
*/
type PutEntropyDeviceParams struct {

	/*Body
	  Guest entropy device properties

	*/
	Body *models.EntropyDevice

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put entropy device params
func (o *PutEntropyDeviceParams) WithTimeout(timeout time.Duration) *PutEntropyDeviceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put entropy device params
func (o *PutEntropyDeviceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put entropy device params
func (o *PutEntropyDeviceParams) WithContext(ctx context.Context) *PutEntropyDeviceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put entropy device params
func (o *PutEntropyDeviceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put entropy device params
func (o *PutEntropyDeviceParams) WithHTTPClient(client *http.Client) *PutEntropyDeviceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put entropy device params
func (o *PutEntropyDeviceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put entropy device params
func (o *PutEntropyDeviceParams) WithBody(body *models.EntropyDevice) *PutEntropyDeviceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put entropy device params
func (o *PutEntropyDeviceParams) SetBody(body *models.EntropyDevice) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PutEntropyDeviceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
