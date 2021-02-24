package responses

type CategoriesResponse struct {
	Success 	bool			`json:"success"`
	Categories 	[]Categories	`json:"categories"`
}

type Categories struct {
	Id 		int 	`json:"id"`
	Title 	string	`json:"title"`
	Adds 	[]Adds	`json:"adds"`
}
