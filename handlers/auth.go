package handlers

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"

	"educational-platform/database"
	"educational-platform/models"

	"github.com/gofiber/fiber/v3"
)

// Simple in-memory session store for MVP
var sessions = make(map[string]map[string]interface{})

func InitAuth() {
	// Initialize session store
}

// HashPassword creates a SHA256 hash of the password
func HashPassword(password string) string {
	hash := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hash[:])
}

// VerifyPassword checks if the provided password matches the hash
func VerifyPassword(password, hash string) bool {
	return HashPassword(password) == hash
}

// GenerateSessionID creates a random session ID
func GenerateSessionID() string {
	bytes := make([]byte, 32)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

// Middleware to check if user is authenticated
func AuthMiddleware(c fiber.Ctx) error {
	sessionID := c.Cookies("session_id")
	if sessionID == "" {
		return c.Status(401).JSON(models.APIResponse{
			Success: false,
			Message: "Not authenticated",
		})
	}

	session, exists := sessions[sessionID]
	if !exists {
		return c.Status(401).JSON(models.APIResponse{
			Success: false,
			Message: "Not authenticated",
		})
	}

	userID := session["user_id"]
	userType := session["user_type"]

	if userID == nil || userType == nil {
		return c.Status(401).JSON(models.APIResponse{
			Success: false,
			Message: "Not authenticated",
		})
	}

	// Add user info to context
	c.Locals("user_id", userID)
	c.Locals("user_type", userType)

	return c.Next()
}

// Middleware to check if user is a teacher
func TeacherAuthMiddleware(c fiber.Ctx) error {
	if err := AuthMiddleware(c); err != nil {
		return err
	}

	userType := c.Locals("user_type").(string)
	if userType != "teacher" {
		return c.Status(403).JSON(models.APIResponse{
			Success: false,
			Message: "Teacher access required",
		})
	}

	return c.Next()
}

// Middleware to check if user is a student
func StudentAuthMiddleware(c fiber.Ctx) error {
	if err := AuthMiddleware(c); err != nil {
		return err
	}

	userType := c.Locals("user_type").(string)
	if userType != "student" {
		return c.Status(403).JSON(models.APIResponse{
			Success: false,
			Message: "Student access required",
		})
	}

	return c.Next()
}

// Login handler
func LoginHandler(c fiber.Ctx) error {
	var req models.LoginRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(400).JSON(models.APIResponse{
			Success: false,
			Message: "Invalid request body",
		})
	}

	// Try to find user as teacher first
	teacher, err := database.GetTeacherByUsername(req.Username)
	if err == nil {
		if VerifyPassword(req.Password, teacher.PasswordHash) {
			sessionID := GenerateSessionID()
			sessions[sessionID] = map[string]interface{}{
				"user_id":   teacher.ID,
				"user_type": "teacher",
				"username":  teacher.Username,
				"name":      teacher.Name,
			}

			c.Cookie(&fiber.Cookie{
				Name:     "session_id",
				Value:    sessionID,
				HTTPOnly: true,
				Secure:   false, // Set to true in production with HTTPS
			})

			return c.JSON(models.APIResponse{
				Success: true,
				Message: "Login successful",
				Data: map[string]interface{}{
					"user_type": "teacher",
					"user_id":   teacher.ID,
					"username":  teacher.Username,
					"name":      teacher.Name,
				},
			})
		}
	}

	// Try to find user as student
	student, err := database.GetStudentByUsername(req.Username)
	if err == nil {
		if VerifyPassword(req.Password, student.PasswordHash) {
			sessionID := GenerateSessionID()
			sessions[sessionID] = map[string]interface{}{
				"user_id":   student.ID,
				"user_type": "student",
				"username":  student.Username,
				"name":      student.Name,
			}

			c.Cookie(&fiber.Cookie{
				Name:     "session_id",
				Value:    sessionID,
				HTTPOnly: true,
				Secure:   false, // Set to true in production with HTTPS
			})

			return c.JSON(models.APIResponse{
				Success: true,
				Message: "Login successful",
				Data: map[string]interface{}{
					"user_type": "student",
					"user_id":   student.ID,
					"username":  student.Username,
					"name":      student.Name,
				},
			})
		}
	}

	return c.Status(401).JSON(models.APIResponse{
		Success: false,
		Message: "Invalid credentials",
	})
}

// Register handler
func RegisterHandler(c fiber.Ctx) error {
	var req models.RegisterRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(400).JSON(models.APIResponse{
			Success: false,
			Message: "Invalid request body",
		})
	}

	passwordHash := HashPassword(req.Password)

	if req.UserType == "teacher" {
		err := database.CreateTeacher(req.Username, req.Email, passwordHash, req.Name)
		if err != nil {
			return c.Status(400).JSON(models.APIResponse{
				Success: false,
				Message: "Failed to create teacher account",
			})
		}
	} else if req.UserType == "student" {
		err := database.CreateStudent(req.Username, req.Email, passwordHash, req.Name)
		if err != nil {
			return c.Status(400).JSON(models.APIResponse{
				Success: false,
				Message: "Failed to create student account",
			})
		}
	} else {
		return c.Status(400).JSON(models.APIResponse{
			Success: false,
			Message: "Invalid user type",
		})
	}

	return c.JSON(models.APIResponse{
		Success: true,
		Message: "Account created successfully",
	})
}

// Logout handler
func LogoutHandler(c fiber.Ctx) error {
	sessionID := c.Cookies("session_id")
	if sessionID != "" {
		delete(sessions, sessionID)
	}

	c.Cookie(&fiber.Cookie{
		Name:     "session_id",
		Value:    "",
		HTTPOnly: true,
		Secure:   false,
		MaxAge:   -1, // Delete cookie
	})

	return c.JSON(models.APIResponse{
		Success: true,
		Message: "Logged out successfully",
	})
}

// Get current user info
func GetCurrentUserHandler(c fiber.Ctx) error {
	sessionID := c.Cookies("session_id")
	if sessionID == "" {
		return c.Status(401).JSON(models.APIResponse{
			Success: false,
			Message: "Not authenticated",
		})
	}

	session, exists := sessions[sessionID]
	if !exists {
		return c.Status(401).JSON(models.APIResponse{
			Success: false,
			Message: "Not authenticated",
		})
	}

	userID := session["user_id"]
	userType := session["user_type"]
	username := session["username"]
	name := session["name"]

	if userID == nil {
		return c.Status(401).JSON(models.APIResponse{
			Success: false,
			Message: "Not authenticated",
		})
	}

	return c.JSON(models.APIResponse{
		Success: true,
		Data: map[string]interface{}{
			"user_id":   userID,
			"user_type": userType,
			"username":  username,
			"name":      name,
		},
	})
}