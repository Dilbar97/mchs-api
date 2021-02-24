package adds

import (
	"log"
	"mchs-api/components/db"
	"mchs-api/structures/responses"
)

func GetAddsByCategoryId(categoryId int) []responses.Adds {
	var adds []responses.Adds

	query := "SELECT id, title, description FROM additional_info WHERE category_id = $1"
	rows, err := db.DB.Query(query, categoryId)
	if err != nil {
		log.Println(err)
		return adds
	}

	for rows.Next() {
		var add responses.Adds
		err = rows.Scan(&add.Id, &add.Title, &add.Description)
		if err != nil {
			log.Println(err)
			return adds
		}

		adds = append(adds, add)
	}

	return adds
}
