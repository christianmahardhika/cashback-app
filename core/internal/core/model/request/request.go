package request

type Calculate struct {
	PromoCode string `json:"promo_code" validate:"required"`
	Amount    int    `json:"amount" validate:"required"`
}
