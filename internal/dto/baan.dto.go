package dto

type BaanSize int32

const (
	UNKNOWN = 0
	S       = 1
	M       = 2
	L       = 3
	XL      = 4
	XXL     = 5
)

func (bs BaanSize) String() string {
	return [...]string{"UNKNOWN", "S", "M", "L", "XL", "XXL"}[bs]
}

type Baan struct {
	Id            string   `json:"id"`
	NameTH        string   `json:"name_th"`
	DescriptionTH string   `json:"description_th"`
	NameEN        string   `json:"name_en"`
	DescriptionEN string   `json:"description_en"`
	Size          BaanSize `json:"size"`
	Facebook      string   `json:"facebook"`
	FacebookUrl   string   `json:"facebook_url"`
	Instagram     string   `json:"instagram"`
	InstagramUrl  string   `json:"instagram_url"`
	Line          string   `json:"line"`
	LineUrl       string   `json:"line_url"`
	ImageUrl      string   `json:"image_url"`
}

type BaanInfo struct {
	Id       string   `json:"id"`
	NameTH   string   `json:"name_th"`
	NameEN   string   `json:"name_en"`
	ImageUrl string   `json:"image_url"`
	Size     BaanSize `json:"size"`
}

type FindAllBaanRequest struct {
}

type FindAllBaanResponse struct {
	Baans []*Baan `json:"baans"`
}

type FindOneBaanRequest struct {
	Id string `json:"id" validate:"required"`
}

type FindOneBaanResponse struct {
	Baan *Baan `json:"baan"`
}
