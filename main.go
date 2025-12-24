package main

import (
	"beego-postgresql/database"
	"beego-postgresql/telemetry"
	_ "beego-postgresql/routers"
	"github.com/beego/beego/v2/server/web"
	"net/http"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func main() {
	// Init OpenTelemetry traces
	cleanup := telemetry.InitTracer("beego-postgresql-app")
	defer cleanup()

	database.Init()

	web.RunWithMiddleWares(":8080",
		func(next http.Handler) http.Handler {
			return otelhttp.NewHandler(
				next,
				"beego-http-server",
			)
		},
	)
}