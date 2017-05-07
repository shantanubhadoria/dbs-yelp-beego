package controllers

import (
	"github.com/astaxie/beego"
	"github.com/shantanubhadoria/dbs-yelp-beego/models"
	"github.com/astaxie/beego/logs"
)

// RestaurantController operations for Restaurant
type RestaurantController struct {
	beego.Controller
}

// URLMapping ...
func (c *RestaurantController) URLMapping() {
	c.Mapping("GetAll", c.GetAll)
}

// GetAll ...
// @Title GetAll
// @Description get Restaurants
// @Param	keywords	path	string	true	"keywords search string"
// @Param	latitude	path	string	true	"latitude of the search location"
// @Param	longitude	path	string	true	"longitude of the search location"
// @Success 200 {object} models.RestaurantResponse
// @Failure 400 Bad request
// @Failure 404 Not found
// @router /:keywords/:latitude/:longitude [get]
func (o *RestaurantController) GetAll() {
	logs.GetLogger("RestaurantController#GetAll")
  keywords := o.Ctx.Input.Param(":keywords")
	latitude := o.Ctx.Input.Param(":latitude")
	longitude := o.Ctx.Input.Param(":longitude")
	restaurants, err := models.GetAllRestaurant(keywords, latitude, longitude)

	// body, _ := ioutil.ReadAll(response.Body)
	// json.Unmarshal(body, &authResp)
	logs.Info(restaurants)
	if err != nil {
		logs.Error("error")
		o.Data["json"] = []int{1, 2, 3}
	} else {
	// 	o.Data["json"] = []int{1, 2, 3}
		o.Data["json"] = restaurants
	}
	o.ServeJSON()
}
