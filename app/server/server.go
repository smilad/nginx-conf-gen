package server

import (
	"context"
	"github.com/labstack/echo/v4/middleware"
	"github.com/opentracing/opentracing-go"
	echoSwagger "github.com/swaggo/echo-swagger"
	controller "nginx/controller/admin/http"
	_ "nginx/docs"

	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	_ "net/http/pprof"
	"nginx/app/server/routes"
	"nginx/config"
)

// Server struct
type Server struct {
	srv       *echo.Echo
	container DeliveryContainer
}
type DeliveryContainer struct {
	Handler controller.AdminHandler
}

// NewServer New Server constructor
func NewServer(p DeliveryContainer) IServer {
	server := Server{
		srv:       echo.New(),
		container: p,
	}
	return &server
}

func (s *Server) Run() error {
	server := &http.Server{
		Addr:         config.C().Service.Server.Port,
		ReadTimeout:  config.C().Service.Server.ReadTimeout,
		WriteTimeout: config.C().Service.Server.WriteTimeout,
	}

	s.SetUpServer(s.container)
	log.Printf("Server is listening on PORT: %s", config.C().Service.Server.Port)
	if err := s.srv.StartServer(server); err != nil {
		log.Fatalf("Error starting Server: %s ", err)
	}

	return nil

}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.srv.Server.Shutdown(ctx)
}

// Map Server Handlers
func (s *Server) SetUpServer(container DeliveryContainer) {

	// MiddleWare
	s.srv.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize:         1 << 10, // 1 KB
		DisablePrintStack: false,
		DisableStackAll:   false,
	}))
	s.srv.Use(middleware.RequestID())
	s.srv.Use(middleware.Secure())
	s.srv.Use(middleware.BodyLimit("2M"))

	if config.C().Service.Debug {
		s.srv.Use(middleware.Logger())
	}

	//base route
	v1 := s.srv.Group("/v1")

	routes.MapAdminHandler(v1, container.Handler)
	// DO_NOT_TOUCH <<health routes>> UNTOUCHABLE
	s.srv.GET("/health", func(c echo.Context) error {
		span, _ := opentracing.StartSpanFromContext(c.Request().Context(), "health")
		defer span.Finish()
		log.Printf("Health check RequestID: %s", c.Response().Header().Get(echo.HeaderXRequestID))
		return c.String(http.StatusOK, "healthy...")
	})
	s.srv.GET("/swagger/*", echoSwagger.WrapHandler)
}
