//go:build !embed_dash

package dash

import (
	"github.com/TBXark/sphere/server/ginx"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func (w *Web) RegisterDashStatic(route gin.IRouter) {
	if dashFs, err := ginx.Fs(w.config.HTTP.Static, nil, ""); err == nil && dashFs != nil {
		d := route.Group("/", gzip.Gzip(gzip.DefaultCompression))
		d.StaticFS("/", dashFs)
	}
}
