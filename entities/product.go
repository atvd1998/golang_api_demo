package entities

type Product struct {
	Title       string `json:"-"`
	Description string `json:"description"`
}
