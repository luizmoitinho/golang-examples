package main

import (
	"fmt"

	"github.com/luizmoitinho/generics/blog"
	"github.com/luizmoitinho/generics/blog/entities"
	"github.com/luizmoitinho/generics/cache"
)

func main() {
	posts := blog.GetPosts()
	cachePosts := cache.New[entities.Post]()
	for _, value := range posts {
		cachePosts.Set(value.Slug, value)
	}

	categories := blog.GetCategories()
	cacheCategory := cache.New[entities.Category]()
	for _, value := range categories {
		cacheCategory.Set(value.Slug, value)
	}

	fmt.Println(cacheCategory.Get("marketing"))
	fmt.Println(cachePosts.Get("como-criar-o-seu-primeiro-ebook-de-forma-gratuita"))
}
