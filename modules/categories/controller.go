package categories

import "mchs-api/components"

type MainController struct {
	components.Controller
}

// @Title News
// @Description Метод получения всех категории
// @Success 200 {object} responses.CategoriesResponse
// @router /categories [get]
func (this *MainController) Get() {
	this.Data["json"] = getCategories()
	this.ServeJSON()
}


