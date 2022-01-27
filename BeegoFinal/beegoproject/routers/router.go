package routers

import (
	"beegoproject/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
func init() {
    beego.Router("/new", &controllers.CatController{})
}
func init() {
    beego.Router("/category", &controllers.CategoryController{})
}
func init() {
    beego.Router("/breed", &controllers.BreedController{})
}