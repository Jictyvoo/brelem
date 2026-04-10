package brdoc

import (
	"time"

	"github.com/jictyvoo/brelem/internal/cnhcalc"
	"github.com/jictyvoo/brelem/internal/docerrors"
)

// LengthCNH is the number of digits in a valid CNH.
const LengthCNH = cnhcalc.Length

// DriverLicenseType represents the category of a Brazilian driver's license.
type DriverLicenseType rune

const (
	DriverTypeA DriverLicenseType = 'A'
	DriverTypeB DriverLicenseType = 'B'
	DriverTypeC DriverLicenseType = 'C'
	DriverTypeD DriverLicenseType = 'D'
	DriverTypeE DriverLicenseType = 'E'
)

// CNH represents a validated Brazilian driver's license number with optional business fields.
type CNH struct {
	digits    [LengthCNH]byte
	expiresAt time.Time
	firstLic  time.Time
	licType   DriverLicenseType
}

// CNHOption configures optional fields on a CNH.
type CNHOption func(*CNH)

// WithExpiresAt sets the CNH expiration date.
func WithExpiresAt(t time.Time) CNHOption {
	return func(c *CNH) { c.expiresAt = t }
}

// WithFirstLicenseDate sets the date of the first license.
func WithFirstLicenseDate(t time.Time) CNHOption {
	return func(c *CNH) { c.firstLic = t }
}

// WithType sets the driver license type.
func WithType(lt DriverLicenseType) CNHOption {
	return func(c *CNH) { c.licType = lt }
}

// NewCNH parses and validates a CNH number string.
// Optional business fields can be set via CNHOption functions.
func NewCNH(raw string, opts ...CNHOption) (CNH, error) {
	c, err := parseCNH(raw)
	if err != nil {
		return CNH{}, err
	}

	for _, opt := range opts {
		opt(&c)
	}
	return c, nil
}

// parseCNH validates the CNH number without options, enabling zero-alloc
// when no options are passed (avoids escape from opt(&c)).
func parseCNH(raw string) (CNH, error) {
	var (
		c     CNH
		v     = cnhcalc.New()
		count uint
	)

	for _, r := range raw {
		if r >= '0' && r <= '9' {
			if count < LengthCNH {
				c.digits[count] = byte(r - '0')
			}
			v.Feed(r)
			count++
		}
	}

	if count == 0 {
		return CNH{}, docerrors.New(docerrors.ReasonEmpty, "cnh")
	}

	if _, reason := v.Finish(); reason != docerrors.ReasonNoError {
		return CNH{}, docerrors.New(reason, "cnh")
	}
	return c, nil
}

// VerifierDigits returns the two check digits as numeric values (0-9).
func (c CNH) VerifierDigits() [2]uint8 {
	return [2]uint8{c.digits[9], c.digits[10]}
}

// String returns the CNH as an 11-digit string (CNH has no standard formatting).
func (c CNH) String() string {
	return c.Digits()
}

// Digits returns the raw 11-digit string.
func (c CNH) Digits() string {
	var buf [LengthCNH]byte
	for i, d := range c.digits {
		buf[i] = d + '0'
	}
	return string(buf[:])
}

// IsZero reports whether the CNH is the zero value (not constructed via NewCNH).
func (c CNH) IsZero() bool {
	return c == CNH{}
}

// ExpiresAt returns the expiration date if set.
func (c CNH) ExpiresAt() time.Time {
	return c.expiresAt
}

// FirstLicenseDate returns the date of the first license if set.
func (c CNH) FirstLicenseDate() time.Time {
	return c.firstLic
}

// Type returns the driver license type if set.
func (c CNH) Type() DriverLicenseType {
	return c.licType
}

// IsExpired reports whether the CNH has expired based on the current time.
// Returns false if ExpiresAt was not set (zero time).
func (c CNH) IsExpired() bool {
	if c.expiresAt.IsZero() {
		return false
	}
	return c.expiresAt.Before(time.Now())
}
