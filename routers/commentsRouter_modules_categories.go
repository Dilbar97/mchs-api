package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["mchs-api/modules/categories:MainController"] = append(beego.GlobalControllerRouter["mchs-api/modules/categories:MainController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/categories",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
