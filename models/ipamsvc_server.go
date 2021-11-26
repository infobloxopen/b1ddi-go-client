// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// IpamsvcServer Server
//
// A DHCP Config Profile (_dhcp/server_) is a named configuration for specified list of hosts.
//
// swagger:model ipamsvcServer
type IpamsvcServer struct {

	// The description for the DHCP Config Profile. May contain 0 to 1024 characters. Can include UTF-8.
	Comment string `json:"comment,omitempty"`

	// Time when the object has been created.
	// Read Only: true
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// Controls who does the DDNS updates.
	//
	// Valid values are:
	// * _client_: DHCP server updates DNS if requested by client.
	// * _server_: DHCP server always updates DNS, overriding an update request from the client, unless the client requests no updates.
	// * _ignore_: DHCP server always updates DNS, even if the client says not to.
	// * _over_client_update_: Same as _server_. DHCP server always updates DNS, overriding an update request from the client, unless the client requests no updates.
	// * _over_no_update_: DHCP server updates DNS even if the client requests that no updates be done. If the client requests to do the update, DHCP server allows it.
	//
	// Defaults to _client_.
	DdnsClientUpdate string `json:"ddns_client_update,omitempty"`

	// The domain suffix for DDNS updates. FQDN, may be empty.
	//
	// Required if _ddns_enabled_ is true.
	//
	// Defaults to empty.
	DdnsDomain string `json:"ddns_domain,omitempty"`

	// Indicates if DDNS updates should be performed for leases.
	// All other _ddns_*_ configuration is ignored when this flag is unset.
	//
	// At a minimum, _ddns_domain_ and _ddns_zones_ must be configured to enable DDNS.
	//
	// Defaults to _false_.
	DdnsEnabled bool `json:"ddns_enabled,omitempty"`

	// Indicates if DDNS should generate a hostname when not supplied by the client.
	//
	// Defaults to _false_.
	DdnsGenerateName bool `json:"ddns_generate_name,omitempty"`

	// The prefix used in the generation of an FQDN.
	//
	// When generating a name, DHCP server will construct the name in the format: [ddns-generated-prefix]-[address-text].[ddns-qualifying-suffix].
	// where address-text is simply the lease IP address converted to a hyphenated string.
	//
	// Defaults to "myhost".
	DdnsGeneratedPrefix string `json:"ddns_generated_prefix,omitempty"`

	// Determines if DDNS updates are enabled at the server level.
	// Defaults to _true_.
	DdnsSendUpdates bool `json:"ddns_send_updates,omitempty"`

	// Instructs the DHCP server to always update the DNS information when a lease is renewed even if its DNS information has not changed.
	//
	// Defaults to _false_.
	DdnsUpdateOnRenew bool `json:"ddns_update_on_renew,omitempty"`

	// When true, DHCP server will apply conflict resolution, as described in RFC 4703, when attempting to fulfill the update request.
	//
	// When false, DHCP server will simply attempt to update the DNS entries per the request, regardless of whether or not they conflict with existing entries owned by other DHCP4 clients.
	//
	// Defaults to _true_.
	DdnsUseConflictResolution bool `json:"ddns_use_conflict_resolution,omitempty"`

	// The DNS zones that DDNS updates can be sent to.
	// There is no resolver fallback. The target zone must be explicitly configured for the update to be performed.
	//
	// Updates are sent to the closest enclosing zone.
	//
	// Error if _ddns_enabled_ is _true_ and the _ddns_domain_ does not have a corresponding
	// entry in _ddns_zones_.
	//
	// Error if there are items with duplicate zone in the list.
	//
	// Defaults to empty list.
	DdnsZones []*IpamsvcDDNSZone `json:"ddns_zones"`

	// The DHCP configuration for the profile. This controls how leases are issued.
	DhcpConfig *IpamsvcDHCPConfig `json:"dhcp_config,omitempty"`

	// The list of DHCP options or group of options.
	// An option list is ordered and may include both option groups
	// and specific options.
	// Multiple occurences of the same option or group is not an error.
	// The last occurence of an option in the list will be used.
	//
	// Error if the graph of referenced groups contains cycles.
	//
	// Defaults to empty list.
	DhcpOptions []*IpamsvcOptionItem `json:"dhcp_options"`

	// The configuration for header option filename field.
	HeaderOptionFilename string `json:"header_option_filename,omitempty"`

	// The configuration for header option server address field.
	HeaderOptionServerAddress string `json:"header_option_server_address,omitempty"`

	// The configuration for header option server name field.
	HeaderOptionServerName string `json:"header_option_server_name,omitempty"`

	// The character to replace non-matching characters with, when hostname rewrite is enabled.
	//
	// Any single ASCII character.
	//
	// Defaults to "_".
	HostnameRewriteChar string `json:"hostname_rewrite_char,omitempty"`

	// Indicates if client supplied hostnames will be rewritten prior to DDNS update by replacing every character that does not match _hostname_rewrite_regex_ by _hostname_rewrite_char_.
	//
	// Defaults to _false_.
	HostnameRewriteEnabled bool `json:"hostname_rewrite_enabled,omitempty"`

	// The regex bracket expression to match valid characters.
	//
	// Must begin with "[" and end with "]" and be a compilable POSIX regex.
	//
	// Defaults to "[^a-zA-Z0-9_.]".
	HostnameRewriteRegex string `json:"hostname_rewrite_regex,omitempty"`

	// The resource identifier.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// The inheritance configuration.
	InheritanceSources *IpamsvcServerInheritance `json:"inheritance_sources,omitempty"`

	// The name of the DHCP Config Profile. Must contain 1 to 256 characters. Can include UTF-8.
	// Required: true
	Name *string `json:"name"`

	// The tags for the DHCP Config Profile in JSON format.
	Tags interface{} `json:"tags,omitempty"`

	// Time when the object has been updated. Equals to _created_at_ if not updated after creation.
	// Read Only: true
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`

	// The resource identifier.
	VendorSpecificOptionOptionSpace string `json:"vendor_specific_option_option_space,omitempty"`
}

// Validate validates this ipamsvc server
func (m *IpamsvcServer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDdnsZones(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDhcpConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDhcpOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInheritanceSources(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IpamsvcServer) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *IpamsvcServer) validateDdnsZones(formats strfmt.Registry) error {
	if swag.IsZero(m.DdnsZones) { // not required
		return nil
	}

	for i := 0; i < len(m.DdnsZones); i++ {
		if swag.IsZero(m.DdnsZones[i]) { // not required
			continue
		}

		if m.DdnsZones[i] != nil {
			if err := m.DdnsZones[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ddns_zones" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ddns_zones" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IpamsvcServer) validateDhcpConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.DhcpConfig) { // not required
		return nil
	}

	if m.DhcpConfig != nil {
		if err := m.DhcpConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dhcp_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dhcp_config")
			}
			return err
		}
	}

	return nil
}

func (m *IpamsvcServer) validateDhcpOptions(formats strfmt.Registry) error {
	if swag.IsZero(m.DhcpOptions) { // not required
		return nil
	}

	for i := 0; i < len(m.DhcpOptions); i++ {
		if swag.IsZero(m.DhcpOptions[i]) { // not required
			continue
		}

		if m.DhcpOptions[i] != nil {
			if err := m.DhcpOptions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dhcp_options" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dhcp_options" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IpamsvcServer) validateInheritanceSources(formats strfmt.Registry) error {
	if swag.IsZero(m.InheritanceSources) { // not required
		return nil
	}

	if m.InheritanceSources != nil {
		if err := m.InheritanceSources.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("inheritance_sources")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("inheritance_sources")
			}
			return err
		}
	}

	return nil
}

func (m *IpamsvcServer) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *IpamsvcServer) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this ipamsvc server based on the context it is used
func (m *IpamsvcServer) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreatedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDdnsZones(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDhcpConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDhcpOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInheritanceSources(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpdatedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IpamsvcServer) contextValidateCreatedAt(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created_at", "body", strfmt.DateTime(m.CreatedAt)); err != nil {
		return err
	}

	return nil
}

func (m *IpamsvcServer) contextValidateDdnsZones(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DdnsZones); i++ {

		if m.DdnsZones[i] != nil {
			if err := m.DdnsZones[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ddns_zones" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ddns_zones" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IpamsvcServer) contextValidateDhcpConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.DhcpConfig != nil {
		if err := m.DhcpConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dhcp_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dhcp_config")
			}
			return err
		}
	}

	return nil
}

func (m *IpamsvcServer) contextValidateDhcpOptions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DhcpOptions); i++ {

		if m.DhcpOptions[i] != nil {
			if err := m.DhcpOptions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dhcp_options" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dhcp_options" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IpamsvcServer) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *IpamsvcServer) contextValidateInheritanceSources(ctx context.Context, formats strfmt.Registry) error {

	if m.InheritanceSources != nil {
		if err := m.InheritanceSources.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("inheritance_sources")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("inheritance_sources")
			}
			return err
		}
	}

	return nil
}

func (m *IpamsvcServer) contextValidateUpdatedAt(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "updated_at", "body", strfmt.DateTime(m.UpdatedAt)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IpamsvcServer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IpamsvcServer) UnmarshalBinary(b []byte) error {
	var res IpamsvcServer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
