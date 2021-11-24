package kennitala

import (
	"errors"
	"testing"
)

func TestKennitalaIndividualSuccess(t *testing.T) {
	var kennitala Kennitala = "0101303019"
	err := kennitala.IsValidKennitala(KennitalaIndividual)
	if err != nil {
		t.Errorf("Test Fail")
	}
}

func TestKennitalaIsPersonSuccess(t *testing.T) {
	var kennitala Kennitala = "0101303019"
	err := kennitala.IsPerson()
	if err != nil {
		t.Errorf("Test Fail")
	}
}

func TestKennitalaIsNotCompanySuccess(t *testing.T) {
	var kennitala Kennitala = "0101303019"
	err := kennitala.IsValidKennitala(0)
	if err == nil || !errors.Is(err, ErrInvalidKennitalaType) {
		t.Errorf("Test Fail")
	}
}

func TestCompanySuccess(t *testing.T) {
	var kennitala Kennitala = "6204830369" // Marel hf.
	err := kennitala.IsValidKennitala(KennitalaCompany)
	if err != nil {
		t.Errorf("Test Fail")
	}
}
