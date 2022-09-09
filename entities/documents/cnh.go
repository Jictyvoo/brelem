package documents

import (
	"github.com/jictyvoo/brelem/utils"
	"github.com/jictyvoo/brelem/validators"
	"time"
)

type (
	DriverLicenseType rune
	CNH               struct {
		Number           string
		ExpiresAt        time.Time
		FirstLicenseDate time.Time
		Type             DriverLicenseType
	}
)

const (
	DriverTypeA DriverLicenseType = 'A'
	DriverTypeB DriverLicenseType = 'B'
	DriverTypeC DriverLicenseType = 'C'
	DriverTypeD DriverLicenseType = 'D'
	DriverTypeE DriverLicenseType = 'E'
)

func (c CNH) Validate() error {
	if c.ExpiresAt.Before(time.Now()) {
		return utils.ErrExpiredDocument
	}
	return validators.CNH(c.Number)
}
