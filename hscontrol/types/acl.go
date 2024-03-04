package types

import (
	"encoding/json"
	"errors"

	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/datatypes"
	"gorm.io/gorm"

	v1 "github.com/juanfont/headscale/gen/go/headscale/v1"
)

var (
	ErrInvalidPolicyFormat = errors.New("invalid policy format")
)

// ACL describes the data model for ACLs used to restrict access to resources.
type ACL struct {
	gorm.Model
	Policy datatypes.JSON
}

func (a *ACL) Proto() *v1.ACL {
	var p map[string]any

	if err := json.Unmarshal(a.Policy, &p); err != nil {
		return nil
	}

	polPb, err := structpb.NewStruct(p)
	if err != nil {
		return nil
	}

	return &v1.ACL{
		Policy:    polPb,
		UpdatedAt: timestamppb.New(a.UpdatedAt),
	}
}
