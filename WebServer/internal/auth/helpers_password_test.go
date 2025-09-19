package auth

import (
	"fmt"
	"testing"
)

func TestHashPassword(t *testing.T) {
	// Test Casets
	testCases := []struct {
		input string
	}{
		{
			input: "shrikant123%",
		},
		{
			input: "rc.pheuh.,rpgc{[}]}",
		},
		{
			input: "649844444cg.fp[+))]",
		},
		{
			input: "neothu5777539{}=",
		},
	}

	//Check for no erros and valid hash
	for _, cas := range testCases {
		pswd, err := HashPassword(cas.input)

		if err != nil {
			t.Errorf("Getting this error: %v", err)
			continue
		}
		if pswd == "" {
			t.Errorf("Error with converting to hash")
			continue
		}
		fmt.Println(pswd)

	}
}

func TestCheckPasswordHash(t *testing.t) {
	testCases := []struct {
		inputpassword string
		inputHash     string
	}{
		{
			inputpassword: "shrikant123%",
			inputHash:     "$2a$10$vmO0bs18kC5ILrZ5E4x0GeOkOI.F9JBaiJ21gCB9eHyb9wxUj2PS2",
		},
		{
			inputpassword: "rc.pheuh.,rpgc{[}]}",
			inputHash:     "$2a$10$BySdzxW5CZkF8Gprq83Pgej6oXao7nl6veBYsw1HHwud4kYRjhGhu",
		},
		{
			inputpassword: "649844444cg.fp[+))]",
			inputHash:     "$2a$10$DKsRo6V0h5F4XOa/57ss3OtEWJIJjKKfvLqZkyafGdXRLrEh93n6a",
		},
		{
			inputpassword: "neothu5777539{}=",
			inputHash:     "$2a$10$NsfOm6/3.dxBPSaxcAst.O2nxdlxJpDoXabvjNL/nrpxNH3taA/0u",
		},
	}
}
