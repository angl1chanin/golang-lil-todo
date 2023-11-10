package app

import (
	"fmt"
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

	gin.SetMode(gin.DebugMode)

	// HTTP server
	router := gin.Default()

	// Настройка CORS middleware
	corsCfg := cors.DefaultConfig()
	corsCfg.AllowOrigins = []string{"*"}                                                                                                               // Разрешенный домен
	corsCfg.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"}                                                                // Разрешенные методы
	corsCfg.AllowHeaders = []string{"Content-Type", "Access-Control-Allow-Headers", "Authorization", "X-Requested-With", "ngrok-skip-browser-warning"} // Разрешенные заголовки

	// Добавление CORS middleware в роутер
	router.Use(cors.New(corsCfg))

	// setup routes
	routes.Setup(router, handler)

	// start server
	err = router.Run(cfg.Address)
	if err != nil {
		fmt.Println(err)
	}
}
