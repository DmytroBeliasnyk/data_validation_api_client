package datavalidation

type EmailResponse struct {
	Email          string     `json:"email"`
	Autocorrect    string     `json:"autocorrect"`
	Deliverability string     `json:"deliverability"`
	Valid          valid      `json:"is_valid_format"`
	Disposable     disposable `json:"is_disposable_email"`
	Role           role       `json:"is_role_email"`
	Smtp           smtp       `json:"is_smtp_valid"`
}

type valid struct {
	Value bool   `json:"value"`
	Text  string `json:"text"`
}

type disposable struct {
	Value bool   `json:"value"`
	Text  string `json:"text"`
}

type role struct {
	Value bool   `json:"value"`
	Text  string `json:"text"`
}

type smtp struct {
	Value bool   `json:"value"`
	Text  string `json:"text"`
}
