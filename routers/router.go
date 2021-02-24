// @APIVersion 1.0.0
// @Title Antares и Winisland projects api
// @Description Общее апи для двух проектов, содержание внутри отличается контентом
// @Contact sanzhar.sarsenbi@gmail.cim
// @TermsOfServiceUrl https://antaresgambling.com/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego"
	"mchs-api/modules/categories"
	"mchs-api/modules/news"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/main",
			beego.NSInclude(
				&news.MainController{},
			),
			beego.NSInclude(
				&categories.MainController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
