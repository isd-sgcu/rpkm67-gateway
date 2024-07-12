package dto

type DecomposedFile struct {
	Filename string
	Data     []byte
}

type Object struct {
	Url string `json:"url"`
	Key string `json:"key"`
}

type UploadObjectRequest struct {
	Filename string `json:"filename"`
	Data     []byte `json:"data"`
}

type UploadObjectResponse struct {
	Object *Object `json:"object"`
}

type FindByKeyObjectRequest struct {
	Key string `json:"key"`
}

type FindByKeyObjectResponse struct {
	Object *Object `json:"object"`
}

type DeleteObjectRequest struct {
	Key string `json:"key"`
}

type DeleteObjectResponse struct {
	Success bool `json:"success"`
}
