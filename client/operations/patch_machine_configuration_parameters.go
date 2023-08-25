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

// NewPatchMachineConfigurationParams creates a new PatchMachineConfigurationParams object
// with the default values initialized.
func NewPatchMachineConfigurationParams() *PatchMachineConfigurationParams {
	var ()
	return &PatchMachineConfigurationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchMachineConfigurationParamsWithTimeout creates a new PatchMachineConfigurationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchMachineConfigurationParamsWithTimeout(timeout time.Duration) *PatchMachineConfigurationParams {
	var ()
	return &PatchMachineConfigurationParams{

		timeout: timeout,
	}
}

// NewPatchMachineConfigurationParamsWithContext creates a new PatchMachineConfigurationParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchMachineConfigurationParamsWithContext(ctx context.Context) *PatchMachineConfigurationParams {
	var ()
	return &PatchMachineConfigurationParams{

		Context: ctx,
	}
}

// NewPatchMachineConfigurationParamsWithHTTPClient creates a new PatchMachineConfigurationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchMachineConfigurationParamsWithHTTPClient(client *http.Client) *PatchMachineConfigurationParams {
	var ()
	return &PatchMachineConfigurationParams{
		HTTPClient: client,
	}
}

/*PatchMachineConfigurationParams contains all the parameters to send to the API endpoint
for the patch machine configuration operation typically these are written to a http.Request
*/
type PatchMachineConfigurationParams struct {

	/*Body
	  A subset of Machine Configuration Parameters

	*/
	Body *models.MachineConfiguration

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch machine configuration params
func (o *PatchMachineConfigurationParams) WithTimeout(timeout time.Duration) *PatchMachineConfigurationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch machine configuration params
func (o *PatchMachineConfigurationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch machine configuration params
func (o *PatchMachineConfigurationParams) WithContext(ctx context.Context) *PatchMachineConfigurationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch machine configuration params
func (o *PatchMachineConfigurationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch machine configuration params
func (o *PatchMachineConfigurationParams) WithHTTPClient(client *http.Client) *PatchMachineConfigurationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch machine configuration params
func (o *PatchMachineConfigurationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch machine configuration params
func (o *PatchMachineConfigurationParams) WithBody(body *models.MachineConfiguration) *PatchMachineConfigurationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch machine configuration params
func (o *PatchMachineConfigurationParams) SetBody(body *models.MachineConfiguration) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PatchMachineConfigurationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
