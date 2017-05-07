package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/shantanubhadoria/dbs-yelp-beego/controllers:RestaurantController"] = append(beego.GlobalControllerRouter["github.com/shantanubhadoria/dbs-yelp-beego/controllers:RestaurantController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/:keywords/:latitude/:longitude`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

}
