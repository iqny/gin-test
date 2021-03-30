package http

import (
	"byobject/app/admin/main/user/conf"
	"byobject/app/admin/main/user/service"
	"byobject/library/net/http/middleware"
	"io/ioutil"
	"net/http"
	"time"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

var (
	uSvc *service.Service
)

func Init(c *conf.Config) *http.Server {
	initService(c)
	return innerRouter(c)
}
func initService(c *conf.Config) {
	uSvc = service.New(c)
}

func innerRouter(c *conf.Config) (server *http.Server) {
	//正式模式
	if c.App.Debug == false {
		gin.SetMode(gin.ReleaseMode)
		gin.DefaultWriter = ioutil.Discard //把日志丢弃掉
	}
	r := gin.Default()

	r.Use(gzip.Gzip(gzip.BestCompression))
	r.Use(middleware.Cors())
	limit:=ratelimit.NewBucketWithQuantumAndClock(time.Second*1, 300,300,nil)
	v1:=r.Group("/v1")
	v1.GET("/", middleware.Limit(limit), func(ctx *gin.Context) {
		//ctx.String(http.StatusOK, "abs")
		var du = "1.00"
		ctx.JSON(http.StatusOK, gin.H{"aa": du})
	})
	r.NoMethod(notFound)
	r.NoRoute(notFound)
	server = &http.Server{
		Addr:           c.App.Host,
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return
}
func notFound(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":   404,
		"msg":    "无效路由",
		"errors": []string{"404 Not Found"},
	})
	return
}
