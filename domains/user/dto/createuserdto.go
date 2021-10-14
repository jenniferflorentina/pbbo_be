package dto

type CreateUserDTO struct {
	CountryId int64  `json:"countryId" validate:"gt=0"`
	Email     string `json:"email" validate:"format=email & empty=false"`
	Password  string `json:"password" validate:"empty=false"`
	FirstName string `json:"firstName" validate:"empty=false"`
	LastName  string `json:"lastName" validate:"empty=false"`
	Role      string `json:"role" validate:"empty=false"`
}
