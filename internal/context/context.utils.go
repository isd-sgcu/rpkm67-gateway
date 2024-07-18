package context

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/jpeg"
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

	println("File size: ", len(fileBytes))

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
		if err := jpeg.Encode(buf, img, &jpeg.Options{Quality: quality}); err != nil {
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
