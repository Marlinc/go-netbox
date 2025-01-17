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

// NewDcimRackGroupsListParams creates a new DcimRackGroupsListParams object
// with the default values initialized.
func NewDcimRackGroupsListParams() *DcimRackGroupsListParams {
	var ()
	return &DcimRackGroupsListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimRackGroupsListParamsWithTimeout creates a new DcimRackGroupsListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimRackGroupsListParamsWithTimeout(timeout time.Duration) *DcimRackGroupsListParams {
	var ()
	return &DcimRackGroupsListParams{

		timeout: timeout,
	}
}

// NewDcimRackGroupsListParamsWithContext creates a new DcimRackGroupsListParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimRackGroupsListParamsWithContext(ctx context.Context) *DcimRackGroupsListParams {
	var ()
	return &DcimRackGroupsListParams{

		Context: ctx,
	}
}

// NewDcimRackGroupsListParamsWithHTTPClient creates a new DcimRackGroupsListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimRackGroupsListParamsWithHTTPClient(client *http.Client) *DcimRackGroupsListParams {
	var ()
	return &DcimRackGroupsListParams{
		HTTPClient: client,
	}
}

/*DcimRackGroupsListParams contains all the parameters to send to the API endpoint
for the dcim rack groups list operation typically these are written to a http.Request
*/
type DcimRackGroupsListParams struct {

	/*Limit
	  Number of results to return per page.

	*/
	Limit *int64
	/*Name*/
	Name *string
	/*Offset
	  The initial index from which to return the results.

	*/
	Offset *int64
	/*Q*/
	Q *string
	/*Site*/
	Site *string
	/*SiteID*/
	SiteID *int64
	/*Slug*/
	Slug *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim rack groups list params
func (o *DcimRackGroupsListParams) WithTimeout(timeout time.Duration) *DcimRackGroupsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim rack groups list params
func (o *DcimRackGroupsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim rack groups list params
func (o *DcimRackGroupsListParams) WithContext(ctx context.Context) *DcimRackGroupsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim rack groups list params
func (o *DcimRackGroupsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim rack groups list params
func (o *DcimRackGroupsListParams) WithHTTPClient(client *http.Client) *DcimRackGroupsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim rack groups list params
func (o *DcimRackGroupsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the dcim rack groups list params
func (o *DcimRackGroupsListParams) WithLimit(limit *int64) *DcimRackGroupsListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the dcim rack groups list params
func (o *DcimRackGroupsListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the dcim rack groups list params
func (o *DcimRackGroupsListParams) WithName(name *string) *DcimRackGroupsListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the dcim rack groups list params
func (o *DcimRackGroupsListParams) SetName(name *string) {
	o.Name = name
}

// WithOffset adds the offset to the dcim rack groups list params
func (o *DcimRackGroupsListParams) WithOffset(offset *int64) *DcimRackGroupsListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the dcim rack groups list params
func (o *DcimRackGroupsListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQ adds the q to the dcim rack groups list params
func (o *DcimRackGroupsListParams) WithQ(q *string) *DcimRackGroupsListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the dcim rack groups list params
func (o *DcimRackGroupsListParams) SetQ(q *string) {
	o.Q = q
}

// WithSite adds the site to the dcim rack groups list params
func (o *DcimRackGroupsListParams) WithSite(site *string) *DcimRackGroupsListParams {
	o.SetSite(site)
	return o
}

// SetSite adds the site to the dcim rack groups list params
func (o *DcimRackGroupsListParams) SetSite(site *string) {
	o.Site = site
}

// WithSiteID adds the siteID to the dcim rack groups list params
func (o *DcimRackGroupsListParams) WithSiteID(siteID *int64) *DcimRackGroupsListParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the dcim rack groups list params
func (o *DcimRackGroupsListParams) SetSiteID(siteID *int64) {
	o.SiteID = siteID
}

// WithSlug adds the slug to the dcim rack groups list params
func (o *DcimRackGroupsListParams) WithSlug(slug *string) *DcimRackGroupsListParams {
	o.SetSlug(slug)
	return o
}

// SetSlug adds the slug to the dcim rack groups list params
func (o *DcimRackGroupsListParams) SetSlug(slug *string) {
	o.Slug = slug
}

// WriteToRequest writes these params to a swagger request
func (o *DcimRackGroupsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Name != nil {

		// query param name
		var qrName string
		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {
			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if o.Q != nil {

		// query param q
		var qrQ string
		if o.Q != nil {
			qrQ = *o.Q
		}
		qQ := qrQ
		if qQ != "" {
			if err := r.SetQueryParam("q", qQ); err != nil {
				return err
			}
		}

	}

	if o.Site != nil {

		// query param site
		var qrSite string
		if o.Site != nil {
			qrSite = *o.Site
		}
		qSite := qrSite
		if qSite != "" {
			if err := r.SetQueryParam("site", qSite); err != nil {
				return err
			}
		}

	}

	if o.SiteID != nil {

		// query param site_id
		var qrSiteID int64
		if o.SiteID != nil {
			qrSiteID = *o.SiteID
		}
		qSiteID := swag.FormatInt64(qrSiteID)
		if qSiteID != "" {
			if err := r.SetQueryParam("site_id", qSiteID); err != nil {
				return err
			}
		}

	}

	if o.Slug != nil {

		// query param slug
		var qrSlug string
		if o.Slug != nil {
			qrSlug = *o.Slug
		}
		qSlug := qrSlug
		if qSlug != "" {
			if err := r.SetQueryParam("slug", qSlug); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
