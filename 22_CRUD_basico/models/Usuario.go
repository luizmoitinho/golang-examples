package models

//Usuario .. dados persistentes de usuario
type Usuario struct {
	ID    int64  `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}
