package route

import (
	"wycto/app/controller/public"
	"wycto/weigo"
)

func AddPublicRouter() {
	weigo.Router("/public/user/", &public.UserController{})
}
