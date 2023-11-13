package request

type UpdateCountryRequest struct {
	Id   string  `validate:"required"`
	Name string `validate:"required,min=1,max=200" json:"name"`
	Code string `validate:"required,min=1,max=3" json:"code"`
}