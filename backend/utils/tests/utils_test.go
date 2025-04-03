package tests

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Slowers-team/Slowers-App/database"
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

func TestIsPhoneNumberValidWithCorrectShortInput(t *testing.T) {
	correctPhoneNumber := "0101234567"
	result := utils.IsPhoneNumberValid(correctPhoneNumber)
	assert.True(t, result)
}

func TestIsPhoneNumberValidWithCorrectLongInput(t *testing.T) {
	correctPhoneNumber := "1234567890123"
	result := utils.IsPhoneNumberValid(correctPhoneNumber)
	assert.True(t, result)
}

func TestIsPhoneNumberValidWithTooLongPhoneNumber(t *testing.T) {
	incorrectPhoneNumber := "12345678901234"
	result := utils.IsPhoneNumberValid(incorrectPhoneNumber)
	assert.False(t, result)
}

func TestIsPhoneNumberValidWithTooShortPhoneNumber(t *testing.T) {
	incorrectPhoneNumber := "010123456"
	result := utils.IsPhoneNumberValid(incorrectPhoneNumber)
	assert.False(t, result)
}

func TestIsPhoneNumberValidWithPhoneNumberWithLetters(t *testing.T) {
	incorrectPhoneNumber := "0101234ABC"
	result := utils.IsPhoneNumberValid(incorrectPhoneNumber)
	assert.False(t, result)
}

func TestImageIsValidWhenNoteIsNotEmpty(t *testing.T) {
	var image database.Image
	image.Note = "Important note"
	result := utils.ImageNoteIsNotEmpty(image)
	assert.True(t, result)
}

func TestImageIsNotValidWhenNoteIsEmpty(t *testing.T) {
	var image database.Image
	image.Note = ""
	result := utils.ImageNoteIsNotEmpty(image)
	assert.False(t, result)
}

func TestPNGImageReturnsPNGandNIL(t *testing.T) {
	filetype := "image/png"
	result1, result2 := utils.SetImageFormat(filetype)
	assert.Equal(t, result1, "png")
	assert.Equal(t, result2, nil)
}

func TestJPEGImageReturnJPEGandNIL(t *testing.T) {
	filetype := "image/jpeg"
	result1, result2 := utils.SetImageFormat(filetype)
	assert.Equal(t, result1, "jpeg")
	assert.Equal(t, result2, nil)
}

func TestWrongImageTypeReturnsError(t *testing.T) {
	filetype := "image/svg"
	result1, result2 := utils.SetImageFormat(filetype)
	assert.Equal(t, result1, "")
	assert.Equal(t, result2, errors.New("image should be in JPEG or PNG format"))
}

func TestCorrectSizeImageIsValid(t *testing.T) {
	var size int64 =1280
	result := utils.ImageIsNotTooLarge(size)
	assert.True(t, result)

}

func TestTooLargeImageIsNotValid(t *testing.T) {
	var size int64 =999999999999
	result := utils.ImageIsNotTooLarge(size)
	assert.False(t, result)
}

func TestCorrectSizeImageIsLargerThanZero(t *testing.T) {
	var size int64 = 1280
	result := utils.ImageIsLargerThanZero(size)
	assert.True(t,result)
}

func TestZeroSizeImageIsNotValid(t *testing.T) {
	var size int64 = 0
	result := utils.ImageIsLargerThanZero(size)
	assert.False(t,result)
}

func TestNegativeSizeImageIsNotValid(t *testing.T) {
	var size int64 = -100
	result := utils.ImageIsLargerThanZero(size)
	assert.False(t,result)
}