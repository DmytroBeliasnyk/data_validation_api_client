package datavalidation

type PhoneResponse struct {
	Phone    string  `json:"phone"`
	Valid    bool    `json:"valid"`
	Format   format  `json:"format"`
	Country  country `json:"country"`
	Location string  `json:"location"`
	Type     string  `json:"type"`
	Carrier  string  `json:"carrier"`
}

type format struct {
	International string `json:"international"`
	Local         string `json:"local"`
}
type country struct {
	Code   string `json:"code"`
	Name   string `json:"name"`
	Prefix string `json:"prefix"`
}
