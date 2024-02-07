package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/rest/httpx"
	"go_zero/common/response"
	"net/http"

	"go_zero/study_api/user/api_jwt/internal/config"
	"go_zero/study_api/user/api_jwt/internal/handler"
	"go_zero/study_api/user/api_jwt/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/users.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf, rest.WithUnauthorizedCallback(JwtUnauthorizedResult))
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

func JwtUnauthorizedResult(w http.ResponseWriter, r *http.Request, err error) {
	fmt.Println(err)
	httpx.WriteJson(w, http.StatusOK, response.Body{10087, nil, "認證失敗"})
}
