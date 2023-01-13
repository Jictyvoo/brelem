package documents

import (
	"testing"
	"time"

	"github.com/jictyvoo/brelem/validators"
)

func TestCNH_Validate(t *testing.T) {
	documentList := [...]CNH{
		{
			Number: "30231030596", ExpiresAt: time.Date(2001, 3, 20, 0, 0, 0, 0, time.UTC),
			FirstLicenseDate: time.Date(1986, 6, 20, 0, 0, 0, 0, time.UTC), Type: DriverTypeD,
		},
		{
			Number: "66199902513", ExpiresAt: time.Date(1998, 2, 6, 0, 0, 0, 0, time.UTC),
			FirstLicenseDate: time.Date(1987, 2, 00, 0, 0, 0, 0, time.UTC), Type: DriverTypeA,
		},
		{
			Number: "80169632370", ExpiresAt: time.Date(2004, 8, 13, 0, 0, 0, 0, time.UTC),
			FirstLicenseDate: time.Date(1997, 3, 6, 0, 0, 0, 0, time.UTC), Type: DriverTypeE,
		},
		{
			Number: "92505342466", ExpiresAt: time.Date(2013, 1, 21, 0, 0, 0, 0, time.UTC),
			FirstLicenseDate: time.Date(1987, 1, 24, 0, 0, 0, 0, time.UTC), Type: DriverTypeC,
		},
		{
			Number: "65022784469", ExpiresAt: time.Date(1990, 5, 9, 0, 0, 0, 0, time.UTC),
			FirstLicenseDate: time.Date(1990, 1, 22, 0, 0, 0, 0, time.UTC), Type: DriverTypeE,
		},
		{
			Number: "51070280999", ExpiresAt: time.Date(2013, 7, 1, 0, 0, 0, 0, time.UTC),
			FirstLicenseDate: time.Date(1984, 12, 21, 0, 0, 0, 0, time.UTC), Type: DriverTypeC,
		},
		{
			Number: "58049585715", ExpiresAt: time.Date(2013, 6, 2, 0, 0, 0, 0, time.UTC),
			FirstLicenseDate: time.Date(1997, 11, 13, 0, 0, 0, 0, time.UTC), Type: DriverTypeE,
		},
		{
			Number: "00618463148", ExpiresAt: time.Date(1990, 3, 5, 0, 0, 0, 0, time.UTC),
			FirstLicenseDate: time.Date(1991, 7, 3, 0, 0, 0, 0, time.UTC), Type: DriverTypeA,
		},
		{
			Number: "45643514075", ExpiresAt: time.Date(2003, 3, 4, 0, 0, 0, 0, time.UTC),
			FirstLicenseDate: time.Date(1984, 11, 12, 0, 0, 0, 0, time.UTC), Type: DriverTypeC,
		},
		{
			Number: "27850130146", ExpiresAt: time.Date(2012, 5, 25, 0, 0, 0, 0, time.UTC),
			FirstLicenseDate: time.Date(1987, 6, 19, 0, 0, 0, 0, time.UTC), Type: DriverTypeE,
		},
		{
			Number: "50535482964", ExpiresAt: time.Date(2005, 11, 11, 0, 0, 0, 0, time.UTC),
			FirstLicenseDate: time.Date(1996, 2, 18, 0, 0, 0, 0, time.UTC), Type: DriverTypeB,
		},
		{
			Number: "72800394363", ExpiresAt: time.Date(2017, 2, 12, 0, 0, 0, 0, time.UTC),
			FirstLicenseDate: time.Date(1994, 7, 18, 0, 0, 0, 0, time.UTC), Type: DriverTypeC,
		},
		{
			Number: "34237320316", ExpiresAt: time.Date(2004, 5, 18, 0, 0, 0, 0, time.UTC),
			FirstLicenseDate: time.Date(1991, 2, 8, 0, 0, 0, 0, time.UTC), Type: DriverTypeC,
		},
		{
			Number: "57271673315", ExpiresAt: time.Date(1991, 1, 1, 0, 0, 0, 0, time.UTC),
			FirstLicenseDate: time.Date(1988, 1, 6, 0, 0, 0, 0, time.UTC), Type: DriverTypeC,
		},
		{
			Number: "90356036254", ExpiresAt: time.Date(2012, 8, 16, 0, 0, 0, 0, time.UTC),
			FirstLicenseDate: time.Date(1992, 4, 25, 0, 0, 0, 0, time.UTC), Type: DriverTypeE,
		},
		{
			Number: "42645673394", ExpiresAt: time.Date(2000, 7, 24, 0, 0, 0, 0, time.UTC),
			FirstLicenseDate: time.Date(1997, 12, 17, 0, 0, 0, 0, time.UTC), Type: DriverTypeC,
		},
		{
			Number: "81866483793", ExpiresAt: time.Date(2011, 6, 1, 0, 0, 0, 0, time.UTC),
			FirstLicenseDate: time.Date(1993, 8, 21, 0, 0, 0, 0, time.UTC), Type: DriverTypeE,
		},
		{
			Number: "58241900906", ExpiresAt: time.Date(2001, 7, 5, 0, 0, 0, 0, time.UTC),
			FirstLicenseDate: time.Date(1989, 3, 13, 0, 0, 0, 0, time.UTC), Type: DriverTypeD,
		},
		{
			Number: "19268477514", ExpiresAt: time.Date(2004, 2, 24, 0, 0, 0, 0, time.UTC),
			FirstLicenseDate: time.Date(1991, 2, 17, 0, 0, 0, 0, time.UTC), Type: DriverTypeA,
		},
		{
			Number: "06990392146", ExpiresAt: time.Date(2016, 3, 13, 0, 0, 0, 0, time.UTC),
			FirstLicenseDate: time.Date(1987, 1, 1, 0, 0, 0, 0, time.UTC), Type: DriverTypeB,
		},
		{
			Number: "46629197135", ExpiresAt: time.Date(2021, 10, 16, 0, 0, 0, 0, time.UTC),
			FirstLicenseDate: time.Date(1994, 10, 17, 0, 0, 0, 0, time.UTC), Type: DriverTypeC,
		},
		{
			Number: "11891511083", ExpiresAt: time.Date(2022, 2, 23, 0, 0, 0, 0, time.UTC),
			FirstLicenseDate: time.Date(1998, 3, 11, 0, 0, 0, 0, time.UTC), Type: DriverTypeA,
		},
		{
			Number: "05305602596", ExpiresAt: time.Date(2018, 10, 8, 0, 0, 0, 0, time.UTC),
			FirstLicenseDate: time.Date(1990, 2, 6, 0, 0, 0, 0, time.UTC), Type: DriverTypeC,
		},
		{
			Number: "55512917711", ExpiresAt: time.Date(2010, 6, 10, 0, 0, 0, 0, time.UTC),
			FirstLicenseDate: time.Date(1985, 5, 24, 0, 0, 0, 0, time.UTC), Type: DriverTypeD,
		},
		{
			Number: "57056475202", ExpiresAt: time.Date(2016, 12, 21, 0, 0, 0, 0, time.UTC),
			FirstLicenseDate: time.Date(1992, 10, 19, 0, 0, 0, 0, time.UTC), Type: DriverTypeA,
		},
		{
			Number: "47840603918", ExpiresAt: time.Date(2001, 3, 24, 0, 0, 0, 0, time.UTC),
			FirstLicenseDate: time.Date(1998, 12, 24, 0, 0, 0, 0, time.UTC), Type: DriverTypeD,
		},
		{
			Number: "64959312184", ExpiresAt: time.Date(2006, 10, 13, 0, 0, 0, 0, time.UTC),
			FirstLicenseDate: time.Date(1993, 4, 17, 0, 0, 0, 0, time.UTC), Type: DriverTypeC,
		},
		{
			Number: "90712479234", ExpiresAt: time.Date(1996, 4, 13, 0, 0, 0, 0, time.UTC),
			FirstLicenseDate: time.Date(1985, 7, 6, 0, 0, 0, 0, time.UTC), Type: DriverTypeE,
		},
		{
			Number: "52372377844", ExpiresAt: time.Date(1998, 10, 19, 0, 0, 0, 0, time.UTC),
			FirstLicenseDate: time.Date(1989, 6, 4, 0, 0, 0, 0, time.UTC), Type: DriverTypeB,
		},
		{
			Number: "48146789947", ExpiresAt: time.Date(2019, 8, 15, 0, 0, 0, 0, time.UTC),
			FirstLicenseDate: time.Date(1990, 2, 19, 0, 0, 0, 0, time.UTC), Type: DriverTypeD,
		},
		{
			Number: "32514787270", ExpiresAt: time.Date(2012, 12, 8, 0, 0, 0, 0, time.UTC),
			FirstLicenseDate: time.Date(1989, 3, 14, 0, 0, 0, 0, time.UTC), Type: DriverTypeA,
		},
		{
			Number: "07006144716", ExpiresAt: time.Date(2021, 8, 17, 0, 0, 0, 0, time.UTC),
			FirstLicenseDate: time.Date(1987, 1, 20, 0, 0, 0, 0, time.UTC), Type: DriverTypeA,
		},
		{
			Number: "00506885075", ExpiresAt: time.Date(2014, 9, 8, 0, 0, 0, 0, time.UTC),
			FirstLicenseDate: time.Date(1989, 6, 0, 0, 0, 0, 0, time.UTC), Type: DriverTypeC,
		},
		{
			Number: "66464618306", ExpiresAt: time.Date(1991, 4, 17, 0, 0, 0, 0, time.UTC),
			FirstLicenseDate: time.Date(1993, 5, 6, 0, 0, 0, 0, time.UTC), Type: DriverTypeD,
		},
		{
			Number: "86129232978", ExpiresAt: time.Date(2006, 2, 8, 0, 0, 0, 0, time.UTC),
			FirstLicenseDate: time.Date(1984, 0, 24, 0, 0, 0, 0, time.UTC), Type: DriverTypeB,
		},
		{
			Number: "49545976082", ExpiresAt: time.Date(2007, 1, 5, 0, 0, 0, 0, time.UTC),
			FirstLicenseDate: time.Date(1996, 10, 1, 0, 0, 0, 0, time.UTC), Type: DriverTypeD,
		},
		{
			Number: "20539958980", ExpiresAt: time.Date(2007, 2, 2, 0, 0, 0, 0, time.UTC),
			FirstLicenseDate: time.Date(1987, 7, 10, 0, 0, 0, 0, time.UTC), Type: DriverTypeC,
		},
		{
			Number: "98316782009", ExpiresAt: time.Date(1991, 6, 9, 0, 0, 0, 0, time.UTC),
			FirstLicenseDate: time.Date(1995, 5, 8, 0, 0, 0, 0, time.UTC), Type: DriverTypeE,
		},
	}

	for _, document := range documentList {
		if err := validators.CNH(document.Number); err == nil {
			t.Errorf("Validate document`%+v` as valid", document)
		}
	}
}
