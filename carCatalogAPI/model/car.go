package model

type Car struct {
	RegNum string `json:"regNum"`

	Mark string `json:"mark"`

	Model string `json:"model"`

	Year int32 `json:"year,omitempty"`

	Owner *People `json:"owner"`
}
