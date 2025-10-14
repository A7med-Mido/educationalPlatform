package handlers

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// GenerateThumbnail creates a thumbnail for the video using ffmpeg
func GenerateThumbnail(videoPath, filename string) string {
	// Check if ffmpeg is available
	if !isFFmpegAvailable() {
		return generateDefaultThumbnail(videoPath, filename)
	}

	// Generate thumbnail filename
	ext := filepath.Ext(filename)
	thumbnailFilename := strings.TrimSuffix(filename, ext) + "_thumb.jpg"
	thumbnailPath := filepath.Join("./uploads/thumbnails", thumbnailFilename)

	// Create thumbnail using ffmpeg (take frame at 5 seconds)
	cmd := exec.Command("ffmpeg", 
		"-i", videoPath,
		"-ss", "00:00:05",
		"-vframes", "1",
		"-q:v", "2",
		thumbnailPath,
		"-y") // -y to overwrite existing file

	err := cmd.Run()
	if err != nil {
		return generateDefaultThumbnail(videoPath, filename)
	}

	// Check if thumbnail was created successfully
	if _, err := os.Stat(thumbnailPath); os.IsNotExist(err) {
		return generateDefaultThumbnail(videoPath, filename)
	}

	return thumbnailPath
}

// isFFmpegAvailable checks if ffmpeg is installed and available
func isFFmpegAvailable() bool {
	cmd := exec.Command("ffmpeg", "-version")
	err := cmd.Run()
	return err == nil
}

// generateDefaultThumbnail creates a simple placeholder thumbnail
func generateDefaultThumbnail(videoPath, filename string) string {
	// For MVP, we'll create a simple text-based placeholder
	// In a real application, you might want to use a library like image/draw
	
	ext := filepath.Ext(filename)
	thumbnailFilename := strings.TrimSuffix(filename, ext) + "_thumb.jpg"
	thumbnailPath := filepath.Join("./uploads/thumbnails", thumbnailFilename)

	// Create a simple placeholder file (in real app, generate actual image)
	file, err := os.Create(thumbnailPath)
	if err != nil {
		return ""
	}
	defer file.Close()

	// Write a simple placeholder content
	file.WriteString("Video Thumbnail Placeholder")
	
	return thumbnailPath
}