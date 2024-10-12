package dash

import (
	"github.com/gin-gonic/gin"
	doc "github.com/tbxark/go-base-api/docs/dashboard"
	"github.com/tbxark/go-base-api/pkg/web"
)

func (w *Web) bindDocRoute(r gin.IRouter) {
	web.SetupDoc(doc.SwaggerInfoDashboard, "Dashboard", r)
}

//func (w *Web) bindDocRoute(r gin.IRouter) {
//}
