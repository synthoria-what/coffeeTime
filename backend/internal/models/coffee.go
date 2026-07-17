package models

type Coffee struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Brand       string   `json:"brand"`
	Country     string   `json:"country"`
	Tags        []string `json:"tags"`
	ImgUrl      string   `json:"img_url"`
}

type CreateCoffeeRequest struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Brand       string   `json:"brand"`
	Country     string   `json:"country"`
	Tags        []string `json:"tags"`
	ImgUrl      string   `json:"img_url"`
}

type UpdateCoffeeRequest struct {
	Name        *string   `json:"name"`
	Description *string   `json:"description"`
	Brand       *string   `json:"brand"`
	Country     *string   `json:"country"`
	Tags        *[]string `json:"tags"`
	ImgUrl      *string   `json:"img_url"`
}
