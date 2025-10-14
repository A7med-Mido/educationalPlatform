package handlers

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"educational-platform/database"
	"educational-platform/models"

	"github.com/gofiber/fiber/v3"
)

// Teacher dashboard endpoint
func TeacherDashboardHandler(c fiber.Ctx) error {
	userID := c.Locals("user_id").(int)
	
	stats, err := database.GetDashboardStats(userID)
	if err != nil {
		return c.Status(500).JSON(models.APIResponse{
			Success: false,
			Message: "Failed to get dashboard stats",
		})
	}

	return c.JSON(models.APIResponse{
		Success: true,
		Data:    stats,
	})
}

// Upload video endpoint
func UploadVideoHandler(c fiber.Ctx) error {
	userID := c.Locals("user_id").(int)

	// Parse multipart form
	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(400).JSON(models.APIResponse{
			Success: false,
			Message: "Failed to parse form",
		})
	}

	// Get video file
	files := form.File["video"]
	if len(files) == 0 {
		return c.Status(400).JSON(models.APIResponse{
			Success: false,
			Message: "No video file provided",
		})
	}

	videoFile := files[0]

	// Get title and description
	title := c.FormValue("title")
	description := c.FormValue("description")

	if title == "" {
		return c.Status(400).JSON(models.APIResponse{
			Success: false,
			Message: "Title is required",
		})
	}

	// Validate file type
	ext := strings.ToLower(filepath.Ext(videoFile.Filename))
	allowedExts := []string{".mp4", ".avi", ".mov", ".mkv", ".webm"}
	isValidExt := false
	for _, allowedExt := range allowedExts {
		if ext == allowedExt {
			isValidExt = true
			break
		}
	}

	if !isValidExt {
		return c.Status(400).JSON(models.APIResponse{
			Success: false,
			Message: "Invalid video file type",
		})
	}

	// Generate unique filename
	timestamp := time.Now().Unix()
	filename := fmt.Sprintf("video_%d_%d%s", userID, timestamp, ext)
	filePath := filepath.Join("./uploads/videos", filename)

	// Save video file
	err = c.SaveFile(videoFile, filePath)
	if err != nil {
		return c.Status(500).JSON(models.APIResponse{
			Success: false,
			Message: "Failed to save video file",
		})
	}

	// Generate thumbnail
	thumbnailPath := GenerateThumbnail(filePath, filename)

	// Get file size
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return c.Status(500).JSON(models.APIResponse{
			Success: false,
			Message: "Failed to get file info",
		})
	}

	// For MVP, we'll set duration to 0 (could be enhanced with ffmpeg)
	duration := 0

	// Save video info to database
	err = database.CreateVideo(userID, title, description, videoFile.Filename, filePath, thumbnailPath, duration, fileInfo.Size())
	if err != nil {
		// Clean up uploaded file if database save fails
		os.Remove(filePath)
		if thumbnailPath != "" {
			os.Remove(thumbnailPath)
		}
		return c.Status(500).JSON(models.APIResponse{
			Success: false,
			Message: "Failed to save video info",
		})
	}

	return c.JSON(models.APIResponse{
		Success: true,
		Message: "Video uploaded successfully",
	})
}

// Get teacher's videos
func GetTeacherVideosHandler(c fiber.Ctx) error {
	userID := c.Locals("user_id").(int)

	videos, err := database.GetVideosByTeacherID(userID)
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

// Delete video
func DeleteVideoHandler(c fiber.Ctx) error {
	userID := c.Locals("user_id").(int)
	videoIDStr := c.Params("id")
	videoID, err := strconv.Atoi(videoIDStr)
	if err != nil {
		return c.Status(400).JSON(models.APIResponse{
			Success: false,
			Message: "Invalid video ID",
		})
	}

	// Get video info first to check ownership and get file paths
	video, err := database.GetVideoByID(videoID)
	if err != nil {
		return c.Status(404).JSON(models.APIResponse{
			Success: false,
			Message: "Video not found",
		})
	}

	// Check if teacher owns this video
	if video.TeacherID != userID {
		return c.Status(403).JSON(models.APIResponse{
			Success: false,
			Message: "Not authorized to delete this video",
		})
	}

	// Delete from database
	err = database.DeleteVideo(videoID)
	if err != nil {
		return c.Status(500).JSON(models.APIResponse{
			Success: false,
			Message: "Failed to delete video",
		})
	}

	// Delete files
	os.Remove(video.FilePath)
	if video.ThumbnailPath != "" {
		os.Remove(video.ThumbnailPath)
	}

	return c.JSON(models.APIResponse{
		Success: true,
		Message: "Video deleted successfully",
	})
}

// Get teacher's students (subscribers)
func GetTeacherStudentsHandler(c fiber.Ctx) error {
	userID := c.Locals("user_id").(int)

	subscriptions, err := database.GetSubscriptionsByTeacherID(userID)
	if err != nil {
		return c.Status(500).JSON(models.APIResponse{
			Success: false,
			Message: "Failed to get students",
		})
	}

	return c.JSON(models.APIResponse{
		Success: true,
		Data:    subscriptions,
	})
}

// Get video analytics
func GetVideoAnalyticsHandler(c fiber.Ctx) error {
	userID := c.Locals("user_id").(int)

	views, err := database.GetVideoViewsByTeacherID(userID)
	if err != nil {
		return c.Status(500).JSON(models.APIResponse{
			Success: false,
			Message: "Failed to get video analytics",
		})
	}

	return c.JSON(models.APIResponse{
		Success: true,
		Data:    views,
	})
}

// Serve video file
func ServeVideoHandler(c fiber.Ctx) error {
	videoIDStr := c.Params("id")
	videoID, err := strconv.Atoi(videoIDStr)
	if err != nil {
		return c.Status(400).JSON(models.APIResponse{
			Success: false,
			Message: "Invalid video ID",
		})
	}

	video, err := database.GetVideoByID(videoID)
	if err != nil {
		return c.Status(404).JSON(models.APIResponse{
			Success: false,
			Message: "Video not found",
		})
	}

	// Check if file exists
	if _, err := os.Stat(video.FilePath); os.IsNotExist(err) {
		return c.Status(404).JSON(models.APIResponse{
			Success: false,
			Message: "Video file not found",
		})
	}

	return c.SendFile(video.FilePath)
}

// Serve thumbnail
func ServeThumbnailHandler(c fiber.Ctx) error {
	videoIDStr := c.Params("id")
	videoID, err := strconv.Atoi(videoIDStr)
	if err != nil {
		return c.Status(400).JSON(models.APIResponse{
			Success: false,
			Message: "Invalid video ID",
		})
	}

	video, err := database.GetVideoByID(videoID)
	if err != nil {
		return c.Status(404).JSON(models.APIResponse{
			Success: false,
			Message: "Video not found",
		})
	}

	if video.ThumbnailPath == "" {
		return c.Status(404).JSON(models.APIResponse{
			Success: false,
			Message: "Thumbnail not found",
		})
	}

	// Check if file exists
	if _, err := os.Stat(video.ThumbnailPath); os.IsNotExist(err) {
		return c.Status(404).JSON(models.APIResponse{
			Success: false,
			Message: "Thumbnail file not found",
		})
	}

	return c.SendFile(video.ThumbnailPath)
}