package db

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/juanfont/headscale/hscontrol/types"
)

// SetACL inserts or updates the ACL in the database.
func (hsdb *HSDatabase) SetACL(acl *types.ACL) (*types.ACL, error) {
	err := hsdb.DB.
		Model(&types.ACL{}).
		Where("id = ?", 1).
		First(&types.ACL{}).
		Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrap(err, "updating ACL in db")
	}

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		if err := hsdb.DB.Create(&acl).Error; err != nil {
			return nil, errors.Wrap(err, "creating ACL in db")
		}

		return acl, nil
	}

	if err := hsdb.DB.
		Model(acl).
		Clauses(clause.Returning{}).
		Where("id = ?", 1).
		Update("policy", acl.Policy).
		Error; err != nil {
		return nil, errors.Wrap(err, "updating ACL in db")
	}

	return acl, nil
}

// GetACL returns the ACL from the database.
func (hsdb *HSDatabase) GetACL() (*types.ACL, error) {
	acl := types.ACL{}
	if err := hsdb.DB.
		Where("id = ?", 1).
		First(&acl).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return hsdb.SetACL(&types.ACL{Policy: []byte("{}")})
		} else {
			return nil, errors.Wrap(err, "fetching ACL from db")
		}
	}

	return &acl, nil
}
