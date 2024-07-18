package context

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/jpeg"

	"image/png"
	"mime/multipart"
	"strings"
)

func ExtractFile(file *multipart.FileHeader, allowedContent map[string]struct{}, maxSize int64) (data []byte, err error) {
	if !isExisted(allowedContent, file.Header["Content-Type"][0]) {
		return nil, errors.New("Allowed content type is " + fmt.Sprint(strings.Join(mapToArr(allowedContent), ", ")))
	}

	maxSizeMB := maxSize * 1024 * 1024
	fileBytes, err := compressImage(file, int(maxSizeMB))
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

func compressImage(fileHeader *multipart.FileHeader, maxSizeMB int) ([]byte, error) {
	fileExt := fileHeader.Header["Content-Type"][0]
	file, err := fileHeader.Open()
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	var quality int = 80
	var compressed []byte

	for {
		buf := new(bytes.Buffer)

		switch fileExt {
		case "image/jpeg", "image/jpg":
			err = jpeg.Encode(buf, img, &jpeg.Options{Quality: quality})
		case "image/png":
			err = png.Encode(buf, img)
		default:
			return nil, fmt.Errorf("unsupported file type: %v", fileExt)
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

// func resizeImage(img image.Image, width, height int) *image.RGBA {
// 	bounds := img.Bounds()

// 	if width == 0 && height == 0 {
// 		return nil
// 	}

// 	if width == 0 {
// 		width = bounds.Dx() * height / bounds.Dy()
// 	}
// 	if height == 0 {
// 		height = bounds.Dy() * width / bounds.Dx()
// 	}

// 	if width > 500 || height > 500 {
// 		scaleFactor := float64(500) / math.Max(float64(width), float64(height))
// 		width = int(float64(width) * scaleFactor)
// 		height = int(float64(height) * scaleFactor)
// 	}

// 	newImg := image.NewRGBA(image.Rect(0, 0, width, height))
// 	draw.CatmullRom.Scale(newImg, newImg.Bounds(), img, bounds, draw.Over, nil)

// 	return newImg
// }
