package main

import (
	"fmt"
	"log"
	"os"

	"educational-platform/database"
	"educational-platform/handlers"
	"educational-platform/models"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

func main() {
	// Initialize database
	err := database.InitDatabase()
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer database.CloseDatabase()

	// Initialize authentication
	handlers.InitAuth()

	// Create Fiber app
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return c.Status(code).JSON(models.APIResponse{
				Success: false,
				Message: err.Error(),
			})
		},
	})

	// Middleware
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization"},
	}))

	// API Routes - These must come BEFORE any catch-all routes
	api := app.Group("/api")

	// Authentication routes
	auth := api.Group("/auth")
	auth.Post("/login", handlers.LoginHandler)
	auth.Post("/register", handlers.RegisterHandler)
	auth.Post("/logout", handlers.LogoutHandler)
	auth.Get("/me", handlers.GetCurrentUserHandler)

	// Teacher routes
	teacher := api.Group("/teacher")
	teacher.Use(handlers.TeacherAuthMiddleware)
	teacher.Get("/dashboard", handlers.TeacherDashboardHandler)
	teacher.Post("/upload", handlers.UploadVideoHandler)
	teacher.Get("/videos", handlers.GetTeacherVideosHandler)
	teacher.Delete("/videos/:id", handlers.DeleteVideoHandler)
	teacher.Get("/students", handlers.GetTeacherStudentsHandler)
	teacher.Get("/analytics", handlers.GetVideoAnalyticsHandler)

	// Student routes
	student := api.Group("/student")
	student.Use(handlers.StudentAuthMiddleware)
	student.Get("/dashboard", handlers.StudentDashboardHandler)
	student.Get("/videos", handlers.GetStudentVideosHandler)
	student.Post("/watch/:id", handlers.WatchVideoHandler)
	student.Get("/subscriptions", handlers.GetStudentSubscriptionsHandler)
	student.Post("/subscribe/:teacher_id", handlers.SubscribeToTeacherHandler)
	student.Delete("/unsubscribe/:teacher_id", handlers.UnsubscribeFromTeacherHandler)

	// Public API routes
	api.Get("/teachers", handlers.GetTeachersHandler)
	api.Get("/video/:id", handlers.ServeVideoHandler)
	api.Get("/video/:id/thumbnail", handlers.ServeThumbnailHandler)

	// Static file routes - These come AFTER API routes
	app.Get("/login.html", func(c fiber.Ctx) error {
		return c.SendFile("./static/login.html")
	})
	app.Get("/teacher_dashboard.html", func(c fiber.Ctx) error {
		return c.SendFile("./static/teacher_dashboard.html")
	})
	app.Get("/student_dashboard.html", func(c fiber.Ctx) error {
		return c.SendFile("./static/student_dashboard.html")
	})

	// Default route - This comes LAST
	app.Get("/", func(c fiber.Ctx) error {
		return c.Redirect("/login.html")
	})

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	fmt.Printf("🚀 Educational Platform server starting on port %s\n", port)
	fmt.Println("📚 Teacher Dashboard: http://localhost:" + port + "/teacher_dashboard.html")
	fmt.Println("🎓 Student Dashboard: http://localhost:" + port + "/student_dashboard.html")
	fmt.Println("🔐 Login Page: http://localhost:" + port + "/login.html")
	fmt.Println("🔗 API Base URL: http://localhost:" + port + "/api")

	log.Fatal(app.Listen(":" + port))
}