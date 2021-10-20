package dto

type UpdateExpenseDTO struct {
	UserId      int64   `json:"idUser"`
	Nama        string  `json:"nama"`
	ReleaseDate string  `json:"releaseDate"`
	Quantity    float32 `json:"quantity"`
	Category    int64   `json:"category"`
	Description string  `json:"description"`
}
