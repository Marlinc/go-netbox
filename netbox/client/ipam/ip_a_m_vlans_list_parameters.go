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
)

// NewIPAMVlansListParams creates a new IPAMVlansListParams object
// with the default values initialized.
func NewIPAMVlansListParams() *IPAMVlansListParams {
	var ()
	return &IPAMVlansListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIPAMVlansListParamsWithTimeout creates a new IPAMVlansListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIPAMVlansListParamsWithTimeout(timeout time.Duration) *IPAMVlansListParams {
	var ()
	return &IPAMVlansListParams{

		timeout: timeout,
	}
}

// NewIPAMVlansListParamsWithContext creates a new IPAMVlansListParams object
// with the default values initialized, and the ability to set a context for a request
func NewIPAMVlansListParamsWithContext(ctx context.Context) *IPAMVlansListParams {
	var ()
	return &IPAMVlansListParams{

		Context: ctx,
	}
}

// NewIPAMVlansListParamsWithHTTPClient creates a new IPAMVlansListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIPAMVlansListParamsWithHTTPClient(client *http.Client) *IPAMVlansListParams {
	var ()
	return &IPAMVlansListParams{
		HTTPClient: client,
	}
}

/*IPAMVlansListParams contains all the parameters to send to the API endpoint
for the ipam vlans list operation typically these are written to a http.Request
*/
type IPAMVlansListParams struct {

	/*Group*/
	Group *string
	/*GroupID*/
	GroupID *int64
	/*IDIn
	  Multiple values may be separated by commas.

	*/
	IDIn *string
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
	/*Role*/
	Role *string
	/*RoleID*/
	RoleID *int64
	/*Site*/
	Site *string
	/*SiteID*/
	SiteID *int64
	/*Status*/
	Status *string
	/*Tag*/
	Tag *string
	/*Tenant*/
	Tenant *string
	/*TenantID*/
	TenantID *int64
	/*Vid*/
	Vid *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ipam vlans list params
func (o *IPAMVlansListParams) WithTimeout(timeout time.Duration) *IPAMVlansListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam vlans list params
func (o *IPAMVlansListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam vlans list params
func (o *IPAMVlansListParams) WithContext(ctx context.Context) *IPAMVlansListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam vlans list params
func (o *IPAMVlansListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam vlans list params
func (o *IPAMVlansListParams) WithHTTPClient(client *http.Client) *IPAMVlansListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam vlans list params
func (o *IPAMVlansListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGroup adds the group to the ipam vlans list params
func (o *IPAMVlansListParams) WithGroup(group *string) *IPAMVlansListParams {
	o.SetGroup(group)
	return o
}

// SetGroup adds the group to the ipam vlans list params
func (o *IPAMVlansListParams) SetGroup(group *string) {
	o.Group = group
}

// WithGroupID adds the groupID to the ipam vlans list params
func (o *IPAMVlansListParams) WithGroupID(groupID *int64) *IPAMVlansListParams {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the ipam vlans list params
func (o *IPAMVlansListParams) SetGroupID(groupID *int64) {
	o.GroupID = groupID
}

// WithIDIn adds the iDIn to the ipam vlans list params
func (o *IPAMVlansListParams) WithIDIn(iDIn *string) *IPAMVlansListParams {
	o.SetIDIn(iDIn)
	return o
}

// SetIDIn adds the idIn to the ipam vlans list params
func (o *IPAMVlansListParams) SetIDIn(iDIn *string) {
	o.IDIn = iDIn
}

// WithLimit adds the limit to the ipam vlans list params
func (o *IPAMVlansListParams) WithLimit(limit *int64) *IPAMVlansListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the ipam vlans list params
func (o *IPAMVlansListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the ipam vlans list params
func (o *IPAMVlansListParams) WithName(name *string) *IPAMVlansListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the ipam vlans list params
func (o *IPAMVlansListParams) SetName(name *string) {
	o.Name = name
}

// WithOffset adds the offset to the ipam vlans list params
func (o *IPAMVlansListParams) WithOffset(offset *int64) *IPAMVlansListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the ipam vlans list params
func (o *IPAMVlansListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQ adds the q to the ipam vlans list params
func (o *IPAMVlansListParams) WithQ(q *string) *IPAMVlansListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the ipam vlans list params
func (o *IPAMVlansListParams) SetQ(q *string) {
	o.Q = q
}

// WithRole adds the role to the ipam vlans list params
func (o *IPAMVlansListParams) WithRole(role *string) *IPAMVlansListParams {
	o.SetRole(role)
	return o
}

// SetRole adds the role to the ipam vlans list params
func (o *IPAMVlansListParams) SetRole(role *string) {
	o.Role = role
}

// WithRoleID adds the roleID to the ipam vlans list params
func (o *IPAMVlansListParams) WithRoleID(roleID *int64) *IPAMVlansListParams {
	o.SetRoleID(roleID)
	return o
}

// SetRoleID adds the roleId to the ipam vlans list params
func (o *IPAMVlansListParams) SetRoleID(roleID *int64) {
	o.RoleID = roleID
}

// WithSite adds the site to the ipam vlans list params
func (o *IPAMVlansListParams) WithSite(site *string) *IPAMVlansListParams {
	o.SetSite(site)
	return o
}

// SetSite adds the site to the ipam vlans list params
func (o *IPAMVlansListParams) SetSite(site *string) {
	o.Site = site
}

// WithSiteID adds the siteID to the ipam vlans list params
func (o *IPAMVlansListParams) WithSiteID(siteID *int64) *IPAMVlansListParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the ipam vlans list params
func (o *IPAMVlansListParams) SetSiteID(siteID *int64) {
	o.SiteID = siteID
}

// WithStatus adds the status to the ipam vlans list params
func (o *IPAMVlansListParams) WithStatus(status *string) *IPAMVlansListParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the ipam vlans list params
func (o *IPAMVlansListParams) SetStatus(status *string) {
	o.Status = status
}

// WithTag adds the tag to the ipam vlans list params
func (o *IPAMVlansListParams) WithTag(tag *string) *IPAMVlansListParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the ipam vlans list params
func (o *IPAMVlansListParams) SetTag(tag *string) {
	o.Tag = tag
}

// WithTenant adds the tenant to the ipam vlans list params
func (o *IPAMVlansListParams) WithTenant(tenant *string) *IPAMVlansListParams {
	o.SetTenant(tenant)
	return o
}

// SetTenant adds the tenant to the ipam vlans list params
func (o *IPAMVlansListParams) SetTenant(tenant *string) {
	o.Tenant = tenant
}

// WithTenantID adds the tenantID to the ipam vlans list params
func (o *IPAMVlansListParams) WithTenantID(tenantID *int64) *IPAMVlansListParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the ipam vlans list params
func (o *IPAMVlansListParams) SetTenantID(tenantID *int64) {
	o.TenantID = tenantID
}

// WithVid adds the vid to the ipam vlans list params
func (o *IPAMVlansListParams) WithVid(vid *int64) *IPAMVlansListParams {
	o.SetVid(vid)
	return o
}

// SetVid adds the vid to the ipam vlans list params
func (o *IPAMVlansListParams) SetVid(vid *int64) {
	o.Vid = vid
}

// WriteToRequest writes these params to a swagger request
func (o *IPAMVlansListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Group != nil {

		// query param group
		var qrGroup string
		if o.Group != nil {
			qrGroup = *o.Group
		}
		qGroup := qrGroup
		if qGroup != "" {
			if err := r.SetQueryParam("group", qGroup); err != nil {
				return err
			}
		}

	}

	if o.GroupID != nil {

		// query param group_id
		var qrGroupID int64
		if o.GroupID != nil {
			qrGroupID = *o.GroupID
		}
		qGroupID := swag.FormatInt64(qrGroupID)
		if qGroupID != "" {
			if err := r.SetQueryParam("group_id", qGroupID); err != nil {
				return err
			}
		}

	}

	if o.IDIn != nil {

		// query param id__in
		var qrIDIn string
		if o.IDIn != nil {
			qrIDIn = *o.IDIn
		}
		qIDIn := qrIDIn
		if qIDIn != "" {
			if err := r.SetQueryParam("id__in", qIDIn); err != nil {
				return err
			}
		}

	}

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

	if o.Role != nil {

		// query param role
		var qrRole string
		if o.Role != nil {
			qrRole = *o.Role
		}
		qRole := qrRole
		if qRole != "" {
			if err := r.SetQueryParam("role", qRole); err != nil {
				return err
			}
		}

	}

	if o.RoleID != nil {

		// query param role_id
		var qrRoleID int64
		if o.RoleID != nil {
			qrRoleID = *o.RoleID
		}
		qRoleID := swag.FormatInt64(qrRoleID)
		if qRoleID != "" {
			if err := r.SetQueryParam("role_id", qRoleID); err != nil {
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

	if o.Status != nil {

		// query param status
		var qrStatus string
		if o.Status != nil {
			qrStatus = *o.Status
		}
		qStatus := qrStatus
		if qStatus != "" {
			if err := r.SetQueryParam("status", qStatus); err != nil {
				return err
			}
		}

	}

	if o.Tag != nil {

		// query param tag
		var qrTag string
		if o.Tag != nil {
			qrTag = *o.Tag
		}
		qTag := qrTag
		if qTag != "" {
			if err := r.SetQueryParam("tag", qTag); err != nil {
				return err
			}
		}

	}

	if o.Tenant != nil {

		// query param tenant
		var qrTenant string
		if o.Tenant != nil {
			qrTenant = *o.Tenant
		}
		qTenant := qrTenant
		if qTenant != "" {
			if err := r.SetQueryParam("tenant", qTenant); err != nil {
				return err
			}
		}

	}

	if o.TenantID != nil {

		// query param tenant_id
		var qrTenantID int64
		if o.TenantID != nil {
			qrTenantID = *o.TenantID
		}
		qTenantID := swag.FormatInt64(qrTenantID)
		if qTenantID != "" {
			if err := r.SetQueryParam("tenant_id", qTenantID); err != nil {
				return err
			}
		}

	}

	if o.Vid != nil {

		// query param vid
		var qrVid int64
		if o.Vid != nil {
			qrVid = *o.Vid
		}
		qVid := swag.FormatInt64(qrVid)
		if qVid != "" {
			if err := r.SetQueryParam("vid", qVid); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
