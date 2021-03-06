package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

/*PortRange port range

swagger:model PortRange
*/
type PortRange struct {

	/* end
	 */
	End int64 `json:"end,omitempty"`

	/* start
	 */
	Start int64 `json:"start,omitempty"`
}

// Validate validates this port range
func (m *PortRange) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
