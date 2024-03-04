package types

import (
	"errors"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

var (
	ErrInvalidPolicyFormat = errors.New("invalid policy format")
	ErrACLDisabled         = errors.New("change ACL by api is disabled")
)

// ACL describes the data model for ACLs used to restrict access to resources.
type ACL struct {
	gorm.Model
	Policy datatypes.JSON
}
