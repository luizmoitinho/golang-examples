package entities

type Post struct {
	ID       int32
	Category *Category
	Title    string
	Text     string
	Slug     string
}
