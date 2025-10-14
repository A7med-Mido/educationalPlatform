package handlers

import (
	"strconv"

	"educational-platform/database"
	"educational-platform/models"

	"github.com/gofiber/fiber/v3"
)

// Student dashboard endpoint
func StudentDashboardHandler(c fiber.Ctx) error {
	userID := c.Locals("user_id").(int)

	// Get student's subscriptions
	subscriptions, err := database.GetSubscriptionsByStudentID(userID)
	if err != nil {
		return c.Status(500).JSON(models.APIResponse{
			Success: false,
			Message: "Failed to get subscriptions",
		})
	}

	// Get available videos from subscribed teachers
	videos, err := database.GetVideosForStudent(userID)
	if err != nil {
		return c.Status(500).JSON(models.APIResponse{
			Success: false,
			Message: "Failed to get videos",
		})
	}

	return c.JSON(models.APIResponse{
		Success: true,
		Data: map[string]interface{}{
			"subscriptions": subscriptions,
			"videos":        videos,
		},
	})
}

// Get all teachers (for subscription)
func GetTeachersHandler(c fiber.Ctx) error {
	teachers, err := database.GetAllTeachers()
	if err != nil {
		return c.Status(500).JSON(models.APIResponse{
			Success: false,
			Message: "Failed to get teachers",
		})
	}

	return c.JSON(models.APIResponse{
		Success: true,
		Data:    teachers,
	})
}

// Subscribe to a teacher
func SubscribeToTeacherHandler(c fiber.Ctx) error {
	studentID := c.Locals("user_id").(int)
	teacherIDStr := c.Params("teacher_id")
	teacherID, err := strconv.Atoi(teacherIDStr)
	if err != nil {
		return c.Status(400).JSON(models.APIResponse{
			Success: false,
			Message: "Invalid teacher ID",
		})
	}

	// Check if already subscribed
	subscribed, err := database.IsSubscribed(studentID, teacherID)
	if err != nil {
		return c.Status(500).JSON(models.APIResponse{
			Success: false,
			Message: "Failed to check subscription",
		})
	}

	if subscribed {
		return c.Status(400).JSON(models.APIResponse{
			Success: false,
			Message: "Already subscribed to this teacher",
		})
	}

	// Create subscription
	err = database.CreateSubscription(studentID, teacherID)
	if err != nil {
		return c.Status(500).JSON(models.APIResponse{
			Success: false,
			Message: "Failed to subscribe",
		})
	}

	return c.JSON(models.APIResponse{
		Success: true,
		Message: "Successfully subscribed to teacher",
	})
}

// Unsubscribe from a teacher
func UnsubscribeFromTeacherHandler(c fiber.Ctx) error {
	studentID := c.Locals("user_id").(int)
	teacherIDStr := c.Params("teacher_id")
	teacherID, err := strconv.Atoi(teacherIDStr)
	if err != nil {
		return c.Status(400).JSON(models.APIResponse{
			Success: false,
			Message: "Invalid teacher ID",
		})
	}

	// Check if subscribed
	subscribed, err := database.IsSubscribed(studentID, teacherID)
	if err != nil {
		return c.Status(500).JSON(models.APIResponse{
			Success: false,
			Message: "Failed to check subscription",
		})
	}

	if !subscribed {
		return c.Status(400).JSON(models.APIResponse{
			Success: false,
			Message: "Not subscribed to this teacher",
		})
	}

	// Remove subscription
	err = database.Unsubscribe(studentID, teacherID)
	if err != nil {
		return c.Status(500).JSON(models.APIResponse{
			Success: false,
			Message: "Failed to unsubscribe",
		})
	}

	return c.JSON(models.APIResponse{
		Success: true,
		Message: "Successfully unsubscribed from teacher",
	})
}

// Get student's videos (from subscribed teachers)
func GetStudentVideosHandler(c fiber.Ctx) error {
	userID := c.Locals("user_id").(int)

	videos, err := database.GetVideosForStudent(userID)
	if err != nil {
		return c.Status(500).JSON(models.APIResponse{
			Success: false,
			Message: "Failed to get videos",
		})
	}

	return c.JSON(models.APIResponse{
		Success: true,
		Data:    videos,
	})
}

// Watch video (record view)
func WatchVideoHandler(c fiber.Ctx) error {
	studentID := c.Locals("user_id").(int)
	videoIDStr := c.Params("id")
	videoID, err := strconv.Atoi(videoIDStr)
	if err != nil {
		return c.Status(400).JSON(models.APIResponse{
			Success: false,
			Message: "Invalid video ID",
		})
	}

	// Get video info
	video, err := database.GetVideoByID(videoID)
	if err != nil {
		return c.Status(404).JSON(models.APIResponse{
			Success: false,
			Message: "Video not found",
		})
	}

	// Check if student is subscribed to the teacher
	subscribed, err := database.IsSubscribed(studentID, video.TeacherID)
	if err != nil {
		return c.Status(500).JSON(models.APIResponse{
			Success: false,
			Message: "Failed to check subscription",
		})
	}

	if !subscribed {
		return c.Status(403).JSON(models.APIResponse{
			Success: false,
			Message: "You must subscribe to this teacher to watch their videos",
		})
	}

	// Record the view
	err = database.RecordVideoView(studentID, videoID)
	if err != nil {
		return c.Status(500).JSON(models.APIResponse{
			Success: false,
			Message: "Failed to record view",
		})
	}

	return c.JSON(models.APIResponse{
		Success: true,
		Data:    video,
	})
}

// Get student's subscriptions
func GetStudentSubscriptionsHandler(c fiber.Ctx) error {
	userID := c.Locals("user_id").(int)

	subscriptions, err := database.GetSubscriptionsByStudentID(userID)
	if err != nil {
		return c.Status(500).JSON(models.APIResponse{
			Success: false,
			Message: "Failed to get subscriptions",
		})
	}

	return c.JSON(models.APIResponse{
		Success: true,
		Data:    subscriptions,
	})
}