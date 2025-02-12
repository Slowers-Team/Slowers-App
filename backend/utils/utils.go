package utils

import (
	"encoding/json"
	"log"
	"regexp"

	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"

	"golang.org/x/image/draw"

	"golang.org/x/crypto/bcrypt"

	"github.com/Slowers-team/Slowers-App/database"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func IsEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}

func AreIDPtrSlicesEql(a []*database.ObjectID, b []*database.ObjectID) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range len(a) {
		if a[i].Hex() != b[i].Hex() {
			return false
		}
	}

	return true
}

func ToJSON(val any) []byte {
	asJSON, err := json.Marshal(val)
	if err != nil {
		log.Fatal(err)
	}
	return asJSON
}

func ResizeImage(input io.Reader, output io.Writer, format string, newWidth, newHeight int) error {
	var src image.Image
	var err error
	switch format {
	case "png":
		src, err = png.Decode(input)
	case "jpg", "jpeg":
		src, err = jpeg.Decode(input)
	case "gif":
		src, err = gif.Decode(input)
	default:
		return fmt.Errorf("unsupported file type")
	}
	if err != nil {
		return fmt.Errorf("failed to decode image: %w", err)
	}
	dst := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))
	draw.NearestNeighbor.Scale(dst, dst.Rect, src, src.Bounds(), draw.Over, nil)
	switch format {
	case "png":
		err = png.Encode(output, dst)
	case "jpg", "jpeg":
		err = jpeg.Encode(output, dst, nil)
	case "gif":
		err = gif.Encode(output, dst, nil)
	default:
		return fmt.Errorf("unsupported output file type")
	}
	if err != nil {
		return fmt.Errorf("failed to encode image: %w", err)
	}
	return nil
}
