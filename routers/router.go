// @APIVersion 1.0.0
// @Title Restaurant API
// @Description Restaurant service API.
// @Contact shantanu.bhadoria@gmail.com
package routers

import (
	"github.com/shantanubhadoria/dbs-yelp-beego/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
		ns := beego.NewNamespace("/v1",
			beego.NSNamespace("/restaurant",
				beego.NSInclude(
					&controllers.RestaurantController{},
				),
			),
		)
		beego.AddNamespace(ns)
}
