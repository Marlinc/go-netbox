// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2018 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDcimRearPortsDeleteParams creates a new DcimRearPortsDeleteParams object
// with the default values initialized.
func NewDcimRearPortsDeleteParams() *DcimRearPortsDeleteParams {
	var ()
	return &DcimRearPortsDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimRearPortsDeleteParamsWithTimeout creates a new DcimRearPortsDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimRearPortsDeleteParamsWithTimeout(timeout time.Duration) *DcimRearPortsDeleteParams {
	var ()
	return &DcimRearPortsDeleteParams{

		timeout: timeout,
	}
}

// NewDcimRearPortsDeleteParamsWithContext creates a new DcimRearPortsDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimRearPortsDeleteParamsWithContext(ctx context.Context) *DcimRearPortsDeleteParams {
	var ()
	return &DcimRearPortsDeleteParams{

		Context: ctx,
	}
}

// NewDcimRearPortsDeleteParamsWithHTTPClient creates a new DcimRearPortsDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimRearPortsDeleteParamsWithHTTPClient(client *http.Client) *DcimRearPortsDeleteParams {
	var ()
	return &DcimRearPortsDeleteParams{
		HTTPClient: client,
	}
}

/*DcimRearPortsDeleteParams contains all the parameters to send to the API endpoint
for the dcim rear ports delete operation typically these are written to a http.Request
*/
type DcimRearPortsDeleteParams struct {

	/*ID
	  A unique integer value identifying this rear port.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim rear ports delete params
func (o *DcimRearPortsDeleteParams) WithTimeout(timeout time.Duration) *DcimRearPortsDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim rear ports delete params
func (o *DcimRearPortsDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim rear ports delete params
func (o *DcimRearPortsDeleteParams) WithContext(ctx context.Context) *DcimRearPortsDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim rear ports delete params
func (o *DcimRearPortsDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim rear ports delete params
func (o *DcimRearPortsDeleteParams) WithHTTPClient(client *http.Client) *DcimRearPortsDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim rear ports delete params
func (o *DcimRearPortsDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim rear ports delete params
func (o *DcimRearPortsDeleteParams) WithID(id int64) *DcimRearPortsDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim rear ports delete params
func (o *DcimRearPortsDeleteParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimRearPortsDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
