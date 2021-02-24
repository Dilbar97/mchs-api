package categories

import (
	"log"
	"mchs-api/components/db"
	"mchs-api/modules/adds"
	"mchs-api/structures/responses"
)

func getCategories() responses.CategoriesResponse {
	var response responses.CategoriesResponse
	var categories []responses.Categories

	query := "SELECT id, title FROM categories WHERE is_active=true"
	rows, err := db.DB.Query(query)
	if err != nil {
		log.Println(err)
		response.Success = false
		return response
	}

	for rows.Next() {
		var category responses.Categories
		err = rows.Scan(&category.Id, &category.Title)
		if err != nil {
			log.Println(err)
			response.Success = false
			return response
		}
		category.Adds = adds.GetAddsByCategoryId(category.Id)

		categories = append(categories, category)
	}

	response.Success = true
	response.Categories = categories

	return response
}
