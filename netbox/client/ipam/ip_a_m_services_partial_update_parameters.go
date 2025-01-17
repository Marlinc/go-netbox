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

package ipam

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

	models "github.com/hosting-de-labs/go-netbox/netbox/models"
)

// NewIPAMServicesPartialUpdateParams creates a new IPAMServicesPartialUpdateParams object
// with the default values initialized.
func NewIPAMServicesPartialUpdateParams() *IPAMServicesPartialUpdateParams {
	var ()
	return &IPAMServicesPartialUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIPAMServicesPartialUpdateParamsWithTimeout creates a new IPAMServicesPartialUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIPAMServicesPartialUpdateParamsWithTimeout(timeout time.Duration) *IPAMServicesPartialUpdateParams {
	var ()
	return &IPAMServicesPartialUpdateParams{

		timeout: timeout,
	}
}

// NewIPAMServicesPartialUpdateParamsWithContext creates a new IPAMServicesPartialUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewIPAMServicesPartialUpdateParamsWithContext(ctx context.Context) *IPAMServicesPartialUpdateParams {
	var ()
	return &IPAMServicesPartialUpdateParams{

		Context: ctx,
	}
}

// NewIPAMServicesPartialUpdateParamsWithHTTPClient creates a new IPAMServicesPartialUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIPAMServicesPartialUpdateParamsWithHTTPClient(client *http.Client) *IPAMServicesPartialUpdateParams {
	var ()
	return &IPAMServicesPartialUpdateParams{
		HTTPClient: client,
	}
}

/*IPAMServicesPartialUpdateParams contains all the parameters to send to the API endpoint
for the ipam services partial update operation typically these are written to a http.Request
*/
type IPAMServicesPartialUpdateParams struct {

	/*Data*/
	Data *models.WritableService
	/*ID
	  A unique integer value identifying this service.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ipam services partial update params
func (o *IPAMServicesPartialUpdateParams) WithTimeout(timeout time.Duration) *IPAMServicesPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam services partial update params
func (o *IPAMServicesPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam services partial update params
func (o *IPAMServicesPartialUpdateParams) WithContext(ctx context.Context) *IPAMServicesPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam services partial update params
func (o *IPAMServicesPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam services partial update params
func (o *IPAMServicesPartialUpdateParams) WithHTTPClient(client *http.Client) *IPAMServicesPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam services partial update params
func (o *IPAMServicesPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the ipam services partial update params
func (o *IPAMServicesPartialUpdateParams) WithData(data *models.WritableService) *IPAMServicesPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the ipam services partial update params
func (o *IPAMServicesPartialUpdateParams) SetData(data *models.WritableService) {
	o.Data = data
}

// WithID adds the id to the ipam services partial update params
func (o *IPAMServicesPartialUpdateParams) WithID(id int64) *IPAMServicesPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam services partial update params
func (o *IPAMServicesPartialUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IPAMServicesPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
