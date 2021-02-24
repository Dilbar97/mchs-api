package responses

type NewsResponse struct {
	Success		bool	`json:"success"`
	News		[]News	`json:"news"`
}

type News struct {
	Id 				int 		`json:"id"`
	Title			string 		`json:"title"`
	Description 	string 		`json:"description"`
	Image  			string 		`json:"image"`
	Date 			string 		`json:"date"`
}
