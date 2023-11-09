package app

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"todo/config"
	"todo/internal/api/http/v1/handlers"
	"todo/internal/api/http/v1/routes"
	repo "todo/internal/repository/todo"
	service "todo/internal/service/todo"
	"todo/internal/usecase"
)

func Run(cfg *config.Config) {
	// init repository
	todoRepo, err := repo.NewTodoRepository(cfg.StoragePath)
	if err != nil {
		panic(err)
	}

	// init service
	todoSrv := service.NewTodoService(todoRepo)

	// init usecase
	useCase := usecase.NewUseCase(todoSrv)

	// setup handlers
	handler := handlers.NewHandler(useCase)

	// HTTP server
	router := gin.Default()

	// Настройка CORS middleware
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}                                                                                                               // Разрешенный домен
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}                                                                         // Разрешенные методы
	config.AllowHeaders = []string{"Content-Type", "Access-Control-Allow-Headers", "Authorization", "X-Requested-With", "ngrok-skip-browser-warning"} // Разрешенные заголовки

	// Добавление CORS middleware в роутер
	router.Use(cors.New(config))

	// setup routes
	routes.Setup(router, handler)

	// start server
	router.Run(cfg.Address)
}
