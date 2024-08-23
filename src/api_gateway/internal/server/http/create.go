package http

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/swagger"

	// docs are generated by Swag CLI, you have to import them.
	// replace with your own docs folder, usually "github.com/username/reponame/docs"
	_ "github.com/1layar/universe/src/api_gateway/docs"
)

func Create() *fiber.App {
	app := fiber.New(fiber.Config{
		// Enable trusted proxy check
		EnableTrustedProxyCheck: true,
		// Define your trusted proxies (use the appropriate CIDR for your proxy network)
		TrustedProxies: []string{"127.0.0.1/32", "10.0.0.0/8", "172.16.0.0/12", "192.168.0.0/16"},
		ProxyHeader:    "X-Forwarded-For",
	})

	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New())
	app.Use(requestid.New())
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed, // 1
	}))
	app.Use(limiter.New(limiter.Config{
		Max:        20,
		Expiration: time.Second * 60,
	}))

	app.Get("/api/v1/doc/*", swagger.New(swagger.Config{ // custom
		DocExpansion:         "list",
		PersistAuthorization: true,
	}))

	return app
}
