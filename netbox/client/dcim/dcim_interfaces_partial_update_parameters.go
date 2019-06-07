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

	models "github.com/hosting-de-labs/go-netbox/netbox/models"
)

// NewDcimInterfacesPartialUpdateParams creates a new DcimInterfacesPartialUpdateParams object
// with the default values initialized.
func NewDcimInterfacesPartialUpdateParams() *DcimInterfacesPartialUpdateParams {
	var ()
	return &DcimInterfacesPartialUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimInterfacesPartialUpdateParamsWithTimeout creates a new DcimInterfacesPartialUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimInterfacesPartialUpdateParamsWithTimeout(timeout time.Duration) *DcimInterfacesPartialUpdateParams {
	var ()
	return &DcimInterfacesPartialUpdateParams{

		timeout: timeout,
	}
}

// NewDcimInterfacesPartialUpdateParamsWithContext creates a new DcimInterfacesPartialUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimInterfacesPartialUpdateParamsWithContext(ctx context.Context) *DcimInterfacesPartialUpdateParams {
	var ()
	return &DcimInterfacesPartialUpdateParams{

		Context: ctx,
	}
}

// NewDcimInterfacesPartialUpdateParamsWithHTTPClient creates a new DcimInterfacesPartialUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimInterfacesPartialUpdateParamsWithHTTPClient(client *http.Client) *DcimInterfacesPartialUpdateParams {
	var ()
	return &DcimInterfacesPartialUpdateParams{
		HTTPClient: client,
	}
}

/*DcimInterfacesPartialUpdateParams contains all the parameters to send to the API endpoint
for the dcim interfaces partial update operation typically these are written to a http.Request
*/
type DcimInterfacesPartialUpdateParams struct {

	/*Data*/
	Data *models.WritableDeviceInterface
	/*ID
	  A unique integer value identifying this interface.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim interfaces partial update params
func (o *DcimInterfacesPartialUpdateParams) WithTimeout(timeout time.Duration) *DcimInterfacesPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim interfaces partial update params
func (o *DcimInterfacesPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim interfaces partial update params
func (o *DcimInterfacesPartialUpdateParams) WithContext(ctx context.Context) *DcimInterfacesPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim interfaces partial update params
func (o *DcimInterfacesPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim interfaces partial update params
func (o *DcimInterfacesPartialUpdateParams) WithHTTPClient(client *http.Client) *DcimInterfacesPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim interfaces partial update params
func (o *DcimInterfacesPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim interfaces partial update params
func (o *DcimInterfacesPartialUpdateParams) WithData(data *models.WritableDeviceInterface) *DcimInterfacesPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim interfaces partial update params
func (o *DcimInterfacesPartialUpdateParams) SetData(data *models.WritableDeviceInterface) {
	o.Data = data
}

// WithID adds the id to the dcim interfaces partial update params
func (o *DcimInterfacesPartialUpdateParams) WithID(id int64) *DcimInterfacesPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim interfaces partial update params
func (o *DcimInterfacesPartialUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimInterfacesPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
