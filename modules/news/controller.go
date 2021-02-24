package news

import (
	"mchs-api/components"
)

type MainController struct {
	components.Controller
}

// @Title News
// @Description Метод получения всех новостей
// @Success 200 {object} responses.NewsResponse
// @router /news [get]
func (this *MainController) Get() {
	this.Data["json"] = getNews()
	this.ServeJSON()
}
