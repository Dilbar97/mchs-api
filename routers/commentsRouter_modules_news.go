package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["mchs-api/modules/news:MainController"] = append(beego.GlobalControllerRouter["mchs-api/modules/news:MainController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/news",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
