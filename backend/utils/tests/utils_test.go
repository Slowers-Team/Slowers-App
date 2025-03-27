package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Slowers-team/Slowers-App/utils"
)

func TestIsEmailValidWithCorrectInput(t *testing.T) {
	correctEmail := "testuser@test.fi"
	result := utils.IsEmailValid(correctEmail)
	assert.True(t, result)
}

func TestIsEmailValidWithCorrectInputWithSpecialCharacters(t *testing.T) {
	correctEmail := "t%est_user.1+2-3@test-mail.456.fi"
	result := utils.IsEmailValid(correctEmail)
	assert.True(t, result)
}

func TestIsEmailValidWithMissingAtSign(t *testing.T) {
	incorrectEmail := "testusertest.fi"
	result := utils.IsEmailValid(incorrectEmail)
	assert.False(t, result)
}

func TestIsEmailValidWithMissingLastDot(t *testing.T) {
	incorrectEmail := "testuser@testfi"
	result := utils.IsEmailValid(incorrectEmail)
	assert.False(t, result)
}

func TestIsEmailValidWithMissingBeginning(t *testing.T) {
	incorrectEmail := "@test.fi"
	result := utils.IsEmailValid(incorrectEmail)
	assert.False(t, result)
}

func TestIsEmailValidWithMissingEnding(t *testing.T) {
	incorrectEmail := "testuser@"
	result := utils.IsEmailValid(incorrectEmail)
	assert.False(t, result)
}

func TestIsEmailValidWithLargeLetters(t *testing.T) {
	incorrectEmail := "testUSER@test.fi"
	result := utils.IsEmailValid(incorrectEmail)
	assert.False(t, result)
}

func TestIsEmailValidWithTooLongTopLevelDomain(t *testing.T) {
	incorrectEmail := "testuser@test.ficom"
	result := utils.IsEmailValid(incorrectEmail)
	assert.False(t, result)
}

func TestIsEmailValidWithNumbersInTopLevelDoman(t *testing.T) {
	incorrectEmail := "testuser@test.fi5"
	result := utils.IsEmailValid(incorrectEmail)
	assert.False(t, result)
}

func TestIsBusinessIdCodeValidWithCorrectInput(t *testing.T) {
	correctIdCode := "1234567-8"
	result := utils.IsBusinessIdCodeValid(correctIdCode)
	assert.True(t, result)
}

func TestIsBusinessIdCodeValidWithDashInWrongPlace(t *testing.T) {
	incorrectIdCode := "1-2345678"
	result := utils.IsBusinessIdCodeValid(incorrectIdCode)
	assert.False(t, result)
}

func TestIsBusinessIdCodeValidWithMissingDash(t *testing.T) {
	incorrectIdCode := "12345678"
	result := utils.IsBusinessIdCodeValid(incorrectIdCode)
	assert.False(t, result)
}
func TestIsBusinessIdCodeValidWithMissingDashButCorrectLength(t *testing.T) {
	incorrectIdCode := "123456789"
	result := utils.IsBusinessIdCodeValid(incorrectIdCode)
	assert.False(t, result)
}

func TestIsBusinessIdCodeValidWithTooShortCode(t *testing.T) {
	incorrectIdCode := "123456-8"
	result := utils.IsBusinessIdCodeValid(incorrectIdCode)
	assert.False(t, result)
}

func TestIsBusinessIdCodeValidWithTooLongCode(t *testing.T) {
	incorrectIdCode := "12345678-9"
	result := utils.IsBusinessIdCodeValid(incorrectIdCode)
	assert.False(t, result)
}

func TestIsBusinessIdCodeValidWithTooLongEnd(t *testing.T) {
	incorrectIdCode := "123456-78"
	result := utils.IsBusinessIdCodeValid(incorrectIdCode)
	assert.False(t, result)
}
func TestIsBusinessIdCodeValidWithLettersAtBeginning(t *testing.T) {
	incorrectIdCode := "123456A-8"
	result := utils.IsBusinessIdCodeValid(incorrectIdCode)
	assert.False(t, result)
}

func TestIsBusinessIdCodeValidWithLettersAtEnding(t *testing.T) {
	incorrectIdCode := "1234567-B"
	result := utils.IsBusinessIdCodeValid(incorrectIdCode)
	assert.False(t, result)
}

func TestIsPostalCodeValidWithCorrectInput(t *testing.T) {
	correctPostalCode := "12345"
	result := utils.IsPostalCodeValid(correctPostalCode)
	assert.True(t, result)
}

func TestIsPostalCodeValidWithTooShortInput(t *testing.T) {
	incorrectPostalCode := "1234"
	result := utils.IsPostalCodeValid(incorrectPostalCode)
	assert.False(t, result)
}

func TestIsPostalCodeValidWithTooLongInput(t *testing.T) {
	incorrectPostalCode := "123456"
	result := utils.IsPostalCodeValid(incorrectPostalCode)
	assert.False(t, result)
}

func TestIsPostalCodeValidWithCodeWithLetters(t *testing.T) {
	incorrectPostalCode := "1a3B5"
	result := utils.IsPostalCodeValid(incorrectPostalCode)
	assert.False(t, result)
}

func TestIsPostalCodeValidWithSpecialCharacters(t *testing.T) {
	incorrectPostalCode := "12@34"
	result := utils.IsPostalCodeValid(incorrectPostalCode)
	assert.False(t, result)
}
