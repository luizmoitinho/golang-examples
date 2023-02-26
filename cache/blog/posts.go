package blog

import "github.com/luizmoitinho/generics/blog/entities"

func GetPosts() []entities.Post {
	return []entities.Post{
		{
			ID:       1,
			Category: &entities.Category{ID: 1, Name: "Marketing"},
			Title:    "MKT Digital",
			Text:     "Aprendendo sobre MKT Digital",
			Slug:     "aprendendo-sobre-mkt-digital",
		},
		{
			ID:       2,
			Category: &entities.Category{ID: 2, Name: "E-books"},
			Title:    "Criando o primeiro E-book",
			Text:     "Como criar o seu primeiro e-book de forma gratuita",
			Slug:     "como-criar-o-seu-primeiro-ebook-de-forma-gratuita",
		},
	}
}
