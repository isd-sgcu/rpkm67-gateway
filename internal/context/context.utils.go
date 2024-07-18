package context

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"math"

	"image/png"
	"mime/multipart"
	"strings"
)

func ExtractFile(file *multipart.FileHeader, allowedContent map[string]struct{}, maxSize int64) (data []byte, err error) {
	format := file.Header["Content-Type"][0]
	if !isExisted(allowedContent, format) {
		return nil, errors.New("Allowed content type is " + fmt.Sprint(strings.Join(mapToArr(allowedContent), ", ")))
	}

	resizedImage, err := resizeImage(file, 500, 500)
	if err != nil {
		return nil, err
	}

	maxSizeMB := maxSize * 1024 * 1024
	fileBytes, err := compressImage(resizedImage, format, int(maxSizeMB))
	if err != nil {
		return nil, err
	}

	return fileBytes, nil
}

func isExisted(e map[string]struct{}, key string) bool {
	_, ok := e[key]
	return ok
}

func mapToArr(m map[string]struct{}) []string {
	var arr []string
	for k := range m {
		arr = append(arr, k)
	}
	return arr
}

func compressImage(img image.Image, format string, maxSizeMB int) ([]byte, error) {
	var quality int = 80
	var compressed []byte
	var err error

	for {
		buf := new(bytes.Buffer)

		switch format {
		case "image/jpeg", "image/jpg":
			err = jpeg.Encode(buf, img, &jpeg.Options{Quality: quality})
		case "image/png":
			err = png.Encode(buf, img)
		default:
			return nil, fmt.Errorf("unsupported file type: %v", format)
		}
		if err != nil {
			return nil, err
		}

		compressed = buf.Bytes()
		if len(compressed) <= maxSizeMB {
			break
		}

		// Reduce quality if size is still too large
		quality -= 10
		if quality <= 0 {
			return nil, fmt.Errorf(fmt.Sprintf("unable to compress image to %v MB", maxSizeMB/(1024*1024)))
		}
	}

	return compressed, nil
}

func resizeImage(fileHeader *multipart.FileHeader, newWidth, newHeight int) (image.Image, error) {
	format := fileHeader.Header["Content-Type"][0]
	file, err := fileHeader.Open()
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	origWidth := img.Bounds().Dx()
	origHeight := img.Bounds().Dy()

	xOffset := int(math.Max(0, float64(origWidth-newWidth)/2))
	yOffset := int(math.Max(0, float64(origHeight-newHeight)/2))

	cropRect := image.Rect(xOffset, yOffset, xOffset+newWidth, yOffset+newHeight)
	var croppedImg image.Image
	if format == "image/jpeg" || format == "image/jpg" {
		croppedImg = img.(*image.YCbCr).SubImage(cropRect)
	} else if format == "image/png" {
		switch img.(type) {
		case *image.RGBA:
			croppedImg = img.(*image.RGBA).SubImage(cropRect)
		case *image.NRGBA:
			croppedImg = img.(*image.NRGBA).SubImage(cropRect)
		default:
			return nil, fmt.Errorf("unsupported image type: %T", img)
		}
	}

	return croppedImg, nil
}
