package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"mime/multipart"
	"net/textproto"
	"path/filepath"
	"regexp"

	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"

	"golang.org/x/image/draw"

	"golang.org/x/crypto/bcrypt"
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

func BufferToMultipartFileHeader(buf *bytes.Buffer, filename string) (*multipart.FileHeader, error) {
	// Determine the content type based on the file extension
	var contentType string
	ext := filepath.Ext(filename)
	switch ext {
	case ".png":
		contentType = "image/png"
	case ".jpeg", ".jpg":
		contentType = "image/jpeg"
	default:
		return nil, errors.New("unsupported file type")
	}

	// Create a new multipart writer
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Create a form file field
	part, err := writer.CreateFormFile("file", filename)
	if err != nil {
		return nil, err
	}

	// Write the buffer content to the form file field
	_, err = buf.WriteTo(part)
	if err != nil {
		return nil, err
	}

	// Close the writer to finalize the multipart message
	writer.Close()

	// Manually create the FileHeader
	fileHeader := &multipart.FileHeader{
		Filename: filename,
		Header: textproto.MIMEHeader{
			"Content-Disposition": []string{fmt.Sprintf(`form-data; name="file"; filename="%s"`, filename)},
			"Content-Type":        []string{contentType},
		},
		Size: int64(buf.Len()),
	}

	return fileHeader, nil
}
