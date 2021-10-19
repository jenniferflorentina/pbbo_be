package dto

type UpdateExpenseDTO struct {
	Nama          string   `json:"Nama"`
	ReleaseDate   string   `json:"ReleaseDate"`
	Quantity      float32  `json:"Quantity"`
	Category	  int64    `json:"Category"`
	Description   string   `json:"Description"`
}
