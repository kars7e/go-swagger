package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// Tag tag
// swagger:model Tag
type Tag struct {

	// id
	ID int64 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this tag
func (m *Tag) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
