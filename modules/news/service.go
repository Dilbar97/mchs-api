package news

import (
	"log"
	"mchs-api/components/db"
	"mchs-api/components/helpers"
	"mchs-api/structures/responses"
)

func getNews() responses.NewsResponse {
	var response responses.NewsResponse
	var news []responses.News

	rows, err := db.DB.Query("SELECT id, title, description, image, created_at FROM news")
	if err != nil {
		log.Println(err)
		response.Success = false
		return response
	}

	for rows.Next() {
		var newsItem responses.News
		err = rows.Scan(&newsItem.Id, &newsItem.Title, &newsItem.Description, &newsItem.Image, &newsItem.Date)
		if err != nil {
			log.Println(err)
			response.Success = false
			return response
		}

		newsItem.Image = helpers.GetImageUrl(newsItem.Image)
		news = append(news, newsItem)
	}

	response.Success = true
	response.News = news

	return response
}
