package blog

import "github.com/luizmoitinho/generics/blog/entities"

func GetCategories() []entities.Category {
	return []entities.Category{
		{ID: 1, Name: "Marketing", Slug: "marketing"},
		{ID: 2, Name: "E-books", Slug: "ebooks"},
		{ID: 3, Name: "Videos", Slug: "videos"},
	}
}
