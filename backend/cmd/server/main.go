package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kohwg/gowoopi/backend/internal/database"
	"github.com/kohwg/gowoopi/backend/internal/handler"
	"github.com/kohwg/gowoopi/backend/internal/middleware"
	"github.com/kohwg/gowoopi/backend/internal/repository/impl"
	"github.com/kohwg/gowoopi/backend/internal/service"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/kohwg/gowoopi/backend/docs"
)

// @title			Gowoopi Table Order API
// @version		1.0
// @description	테이블오더 서비스 API
// @host			localhost:8080
// @BasePath		/
// @securityDefinitions.apikey	BearerAuth
// @in							header
// @name						Authorization

func main() {
	cfg := &database.Config{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnv("DB_PORT", "3306"),
		User:     getEnv("DB_USER", "root"),
		Password: getEnv("DB_PASSWORD", ""),
		DBName:   getEnv("DB_NAME", "table_order"),
	}

	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatal("DB connection failed:", err)
	}
	if err := database.AutoMigrate(db); err != nil {
		log.Fatal("AutoMigrate failed:", err)
	}
	if getEnv("DB_SEED", "false") == "true" {
		if err := database.Seed(db); err != nil {
			slog.Warn("Seed failed", "error", err)
		}
	}

	jwtSecret := getEnv("JWT_SECRET", "dev-secret-key")

	// Repositories
	storeRepo := impl.NewStoreRepository(db)
	adminRepo := impl.NewAdminRepository(db)
	tableRepo := impl.NewTableRepository(db)
	sessionRepo := impl.NewSessionRepository(db)
	categoryRepo := impl.NewCategoryRepository(db)
	menuRepo := impl.NewMenuRepository(db)
	orderRepo := impl.NewOrderRepository(db)

	// Services
	sseMgr := service.NewSSEManager()
	authSvc := service.NewAuthService(storeRepo, tableRepo, sessionRepo, adminRepo, jwtSecret)
	menuSvc := service.NewMenuService(menuRepo, categoryRepo)
	orderSvc := service.NewOrderService(orderRepo, menuRepo, sseMgr)
	tableSvc := service.NewTableService(tableRepo, sessionRepo, orderRepo, sseMgr)

	// Handlers
	authH := handler.NewAuthHandler(authSvc)
	menuH := handler.NewMenuHandler(menuSvc)
	orderH := handler.NewOrderHandler(orderSvc)
	tableH := handler.NewTableHandler(tableSvc)
	sseH := handler.NewSSEHandler(sseMgr)

	r := gin.Default()
	r.Use(middleware.LoggerMiddleware())
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:3001"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Public routes
	r.POST("/api/customer/login", authH.CustomerLogin)
	r.POST("/api/admin/login", authH.AdminLogin)

	// Auth required
	auth := r.Group("", middleware.AuthMiddleware(authSvc))
	auth.POST("/api/auth/refresh", authH.RefreshToken)

	// Customer routes
	customer := auth.Group("/api/customer", middleware.RequireRole("customer"))
	customer.GET("/menus", menuH.GetMenus)
	customer.POST("/orders", orderH.CreateOrder)
	customer.GET("/orders", orderH.GetCustomerOrders)

	// Admin routes
	admin := auth.Group("/api/admin", middleware.RequireRole("admin"))
	admin.GET("/orders/stream", sseH.StreamOrders)
	admin.GET("/orders", orderH.GetAdminOrders)
	admin.PATCH("/orders/:id/status", orderH.UpdateOrderStatus)
	admin.DELETE("/orders/:id", orderH.DeleteOrder)
	admin.POST("/tables/setup", tableH.SetupTable)
	admin.POST("/tables/:id/complete", tableH.CompleteTable)
	admin.GET("/tables/:id/history", tableH.GetTableHistory)
	admin.POST("/menus", menuH.CreateMenu)
	admin.PUT("/menus/:id", menuH.UpdateMenu)
	admin.DELETE("/menus/:id", menuH.DeleteMenu)
	admin.PATCH("/menus/order", menuH.UpdateMenuOrder)

	port := getEnv("SERVER_PORT", "8080")
	slog.Info("Server starting", "port", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Server failed:", err)
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
