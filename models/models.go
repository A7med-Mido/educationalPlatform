package models

import "time"

// Teacher represents a teacher in the system
type Teacher struct {
	ID           int       `json:"id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"` // Don't include in JSON responses
	Name         string    `json:"name"`
	CreatedAt    time.Time `json:"created_at"`
}

// Student represents a student in the system
type Student struct {
	ID           int       `json:"id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"` // Don't include in JSON responses
	Name         string    `json:"name"`
	CreatedAt    time.Time `json:"created_at"`
}

// Video represents a video uploaded by a teacher
type Video struct {
	ID            int       `json:"id"`
	TeacherID     int       `json:"teacher_id"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	Filename      string    `json:"filename"`
	FilePath      string    `json:"file_path"`
	ThumbnailPath string    `json:"thumbnail_path"`
	Duration      int       `json:"duration"` // in seconds
	FileSize      int64     `json:"file_size"` // in bytes
	CreatedAt     time.Time `json:"created_at"`
	TeacherName   string    `json:"teacher_name,omitempty"` // For display purposes
}

// Subscription represents a student's subscription to a teacher
type Subscription struct {
	ID          int       `json:"id"`
	StudentID   int       `json:"student_id"`
	TeacherID   int       `json:"teacher_id"`
	SubscribedAt time.Time `json:"subscribed_at"`
	TeacherName string    `json:"teacher_name,omitempty"` // For display purposes
	StudentName string    `json:"student_name,omitempty"` // For display purposes
}

// VideoView represents a student watching a video
type VideoView struct {
	ID        int       `json:"id"`
	StudentID int       `json:"student_id"`
	VideoID   int       `json:"video_id"`
	WatchedAt time.Time `json:"watched_at"`
	VideoTitle string   `json:"video_title,omitempty"` // For display purposes
	StudentName string  `json:"student_name,omitempty"` // For display purposes
}

// LoginRequest represents login credentials
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// RegisterRequest represents registration data
type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	UserType string `json:"user_type"` // "teacher" or "student"
}

// VideoUploadRequest represents video upload data
type VideoUploadRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// DashboardStats represents statistics for dashboard
type DashboardStats struct {
	TotalVideos     int `json:"total_videos"`
	TotalStudents   int `json:"total_students"`
	TotalViews      int `json:"total_views"`
	RecentVideos    []Video `json:"recent_videos"`
	RecentStudents  []Student `json:"recent_students"`
}

// APIResponse represents a standard API response
type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
