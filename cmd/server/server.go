package server

import (
	"gin/internal/api/handler"

	"gin/internal/application/repository"
	"gin/internal/application/service"
	"gin/internal/infrastructure/mysql"
	"gin/internal/middleware"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	CodeSuccess = iota
	ErrBadConfig
	ErrInternalServer
)

type server struct {
	router  *gin.Engine
	server  *http.Server
	handler *handler.Handler
}

func New() (*server, error) {
	s := &server{
		router: gin.Default(),
		server: &http.Server{
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
		},
	}
	db, err := mysql.NewMySqlClient()
	if err != nil {
		log.Printf("[musiku-server] failed to initialize musiku database : %v\n", err)
		return nil, err
	}
	log.Printf("[musiku-server] succes to initialize musiku database. Database connected\n")

	if err := mysql.Migration(db); err != nil {
		log.Printf("[musiku-server] failed to migrate musiku database : %v\n", err)
		return nil, err
	}

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)

	s.handler = handler.NewHandler(userService)

	s.router = gin.Default()

	return s, nil
}

func Run() int {
	s, err := New()

	if err != nil {
		return ErrBadConfig
	}

	s.Start()

	if err := s.router.Run(); err != nil {
		return ErrInternalServer
	}

	return CodeSuccess
}

func (s *server) Start() {
	log.Printf("[musiku-server] Server is running at %s:%s", os.Getenv("CONFIG_SERVER_HOST"), os.Getenv("CONFIG_SERVER_PORT"))
	log.Println("[musiku-server] starting server...")

	s.router.Use(middleware.CORS())

	s.router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "hi, i'm musiku server"})
	})

	route := s.router.Group("/api/v1")

	user := route.Group("/user")
	user.POST("/register", s.handler.Register)
	user.POST("/login", s.handler.Login)
	user.PATCH("/verify/:id", s.handler.VerifyAccount)

	user.Use(middleware.ValidateJWTToken())
	user.PATCH("/update", s.handler.UpdateUser)
	user.PATCH("/photo-profile", s.handler.UploadPhotoProfile)
}
