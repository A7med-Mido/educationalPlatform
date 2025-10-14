package database

import (
	"educational-platform/models"
)

// Teacher queries
func CreateTeacher(username, email, passwordHash, name string) error {
	query := `INSERT INTO teachers (username, email, password_hash, name) VALUES (?, ?, ?, ?)`
	_, err := DB.Exec(query, username, email, passwordHash, name)
	return err
}

func GetTeacherByUsername(username string) (*models.Teacher, error) {
	query := `SELECT id, username, email, password_hash, name, created_at FROM teachers WHERE username = ?`
	row := DB.QueryRow(query, username)
	
	teacher := &models.Teacher{}
	err := row.Scan(&teacher.ID, &teacher.Username, &teacher.Email, &teacher.PasswordHash, &teacher.Name, &teacher.CreatedAt)
	if err != nil {
		return nil, err
	}
	return teacher, nil
}

func GetTeacherByID(id int) (*models.Teacher, error) {
	query := `SELECT id, username, email, password_hash, name, created_at FROM teachers WHERE id = ?`
	row := DB.QueryRow(query, id)
	
	teacher := &models.Teacher{}
	err := row.Scan(&teacher.ID, &teacher.Username, &teacher.Email, &teacher.PasswordHash, &teacher.Name, &teacher.CreatedAt)
	if err != nil {
		return nil, err
	}
	return teacher, nil
}

// Student queries
func CreateStudent(username, email, passwordHash, name string) error {
	query := `INSERT INTO students (username, email, password_hash, name) VALUES (?, ?, ?, ?)`
	_, err := DB.Exec(query, username, email, passwordHash, name)
	return err
}

func GetStudentByUsername(username string) (*models.Student, error) {
	query := `SELECT id, username, email, password_hash, name, created_at FROM students WHERE username = ?`
	row := DB.QueryRow(query, username)
	
	student := &models.Student{}
	err := row.Scan(&student.ID, &student.Username, &student.Email, &student.PasswordHash, &student.Name, &student.CreatedAt)
	if err != nil {
		return nil, err
	}
	return student, nil
}

func GetStudentByID(id int) (*models.Student, error) {
	query := `SELECT id, username, email, password_hash, name, created_at FROM students WHERE id = ?`
	row := DB.QueryRow(query, id)
	
	student := &models.Student{}
	err := row.Scan(&student.ID, &student.Username, &student.Email, &student.PasswordHash, &student.Name, &student.CreatedAt)
	if err != nil {
		return nil, err
	}
	return student, nil
}

// Video queries
func CreateVideo(teacherID int, title, description, filename, filePath, thumbnailPath string, duration int, fileSize int64) error {
	query := `INSERT INTO videos (teacher_id, title, description, filename, file_path, thumbnail_path, duration, file_size) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := DB.Exec(query, teacherID, title, description, filename, filePath, thumbnailPath, duration, fileSize)
	return err
}

func GetVideosByTeacherID(teacherID int) ([]models.Video, error) {
	query := `
		SELECT v.id, v.teacher_id, v.title, v.description, v.filename, v.file_path, 
		       v.thumbnail_path, v.duration, v.file_size, v.created_at, t.name
		FROM videos v
		JOIN teachers t ON v.teacher_id = t.id
		WHERE v.teacher_id = ?
		ORDER BY v.created_at DESC
	`
	rows, err := DB.Query(query, teacherID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var videos []models.Video
	for rows.Next() {
		var video models.Video
		err := rows.Scan(&video.ID, &video.TeacherID, &video.Title, &video.Description, 
			&video.Filename, &video.FilePath, &video.ThumbnailPath, &video.Duration, 
			&video.FileSize, &video.CreatedAt, &video.TeacherName)
		if err != nil {
			return nil, err
		}
		videos = append(videos, video)
	}
	return videos, nil
}

func GetVideoByID(videoID int) (*models.Video, error) {
	query := `
		SELECT v.id, v.teacher_id, v.title, v.description, v.filename, v.file_path, 
		       v.thumbnail_path, v.duration, v.file_size, v.created_at, t.name
		FROM videos v
		JOIN teachers t ON v.teacher_id = t.id
		WHERE v.id = ?
	`
	row := DB.QueryRow(query, videoID)
	
	video := &models.Video{}
	err := row.Scan(&video.ID, &video.TeacherID, &video.Title, &video.Description, 
		&video.Filename, &video.FilePath, &video.ThumbnailPath, &video.Duration, 
		&video.FileSize, &video.CreatedAt, &video.TeacherName)
	if err != nil {
		return nil, err
	}
	return video, nil
}

func GetVideosForStudent(studentID int) ([]models.Video, error) {
	query := `
		SELECT v.id, v.teacher_id, v.title, v.description, v.filename, v.file_path, 
		       v.thumbnail_path, v.duration, v.file_size, v.created_at, t.name
		FROM videos v
		JOIN teachers t ON v.teacher_id = t.id
		JOIN subscriptions s ON s.teacher_id = t.id
		WHERE s.student_id = ?
		ORDER BY v.created_at DESC
	`
	rows, err := DB.Query(query, studentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var videos []models.Video
	for rows.Next() {
		var video models.Video
		err := rows.Scan(&video.ID, &video.TeacherID, &video.Title, &video.Description, 
			&video.Filename, &video.FilePath, &video.ThumbnailPath, &video.Duration, 
			&video.FileSize, &video.CreatedAt, &video.TeacherName)
		if err != nil {
			return nil, err
		}
		videos = append(videos, video)
	}
	return videos, nil
}

func DeleteVideo(videoID int) error {
	query := `DELETE FROM videos WHERE id = ?`
	_, err := DB.Exec(query, videoID)
	return err
}

// Subscription queries
func CreateSubscription(studentID, teacherID int) error {
	query := `INSERT INTO subscriptions (student_id, teacher_id) VALUES (?, ?)`
	_, err := DB.Exec(query, studentID, teacherID)
	return err
}

func GetSubscriptionsByStudentID(studentID int) ([]models.Subscription, error) {
	query := `
		SELECT s.id, s.student_id, s.teacher_id, s.subscribed_at, t.name
		FROM subscriptions s
		JOIN teachers t ON s.teacher_id = t.id
		WHERE s.student_id = ?
		ORDER BY s.subscribed_at DESC
	`
	rows, err := DB.Query(query, studentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subscriptions []models.Subscription
	for rows.Next() {
		var sub models.Subscription
		err := rows.Scan(&sub.ID, &sub.StudentID, &sub.TeacherID, &sub.SubscribedAt, &sub.TeacherName)
		if err != nil {
			return nil, err
		}
		subscriptions = append(subscriptions, sub)
	}
	return subscriptions, nil
}

func GetSubscriptionsByTeacherID(teacherID int) ([]models.Subscription, error) {
	query := `
		SELECT s.id, s.student_id, s.teacher_id, s.subscribed_at, st.name
		FROM subscriptions s
		JOIN students st ON s.student_id = st.id
		WHERE s.teacher_id = ?
		ORDER BY s.subscribed_at DESC
	`
	rows, err := DB.Query(query, teacherID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subscriptions []models.Subscription
	for rows.Next() {
		var sub models.Subscription
		err := rows.Scan(&sub.ID, &sub.StudentID, &sub.TeacherID, &sub.SubscribedAt, &sub.StudentName)
		if err != nil {
			return nil, err
		}
		subscriptions = append(subscriptions, sub)
	}
	return subscriptions, nil
}

func IsSubscribed(studentID, teacherID int) (bool, error) {
	query := `SELECT COUNT(*) FROM subscriptions WHERE student_id = ? AND teacher_id = ?`
	var count int
	err := DB.QueryRow(query, studentID, teacherID).Scan(&count)
	return count > 0, err
}

func Unsubscribe(studentID, teacherID int) error {
	query := `DELETE FROM subscriptions WHERE student_id = ? AND teacher_id = ?`
	_, err := DB.Exec(query, studentID, teacherID)
	return err
}

// Video view queries
func RecordVideoView(studentID, videoID int) error {
	query := `INSERT OR IGNORE INTO video_views (student_id, video_id) VALUES (?, ?)`
	_, err := DB.Exec(query, studentID, videoID)
	return err
}

func GetVideoViewsByTeacherID(teacherID int) ([]models.VideoView, error) {
	query := `
		SELECT vv.id, vv.student_id, vv.video_id, vv.watched_at, v.title, s.name
		FROM video_views vv
		JOIN videos v ON vv.video_id = v.id
		JOIN students s ON vv.student_id = s.id
		WHERE v.teacher_id = ?
		ORDER BY vv.watched_at DESC
	`
	rows, err := DB.Query(query, teacherID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var views []models.VideoView
	for rows.Next() {
		var view models.VideoView
		err := rows.Scan(&view.ID, &view.StudentID, &view.VideoID, &view.WatchedAt, &view.VideoTitle, &view.StudentName)
		if err != nil {
			return nil, err
		}
		views = append(views, view)
	}
	return views, nil
}

func GetVideoViewCount(videoID int) (int, error) {
	query := `SELECT COUNT(*) FROM video_views WHERE video_id = ?`
	var count int
	err := DB.QueryRow(query, videoID).Scan(&count)
	return count, err
}

// Dashboard statistics queries
func GetDashboardStats(teacherID int) (*models.DashboardStats, error) {
	stats := &models.DashboardStats{}

	// Total videos
	query := `SELECT COUNT(*) FROM videos WHERE teacher_id = ?`
	err := DB.QueryRow(query, teacherID).Scan(&stats.TotalVideos)
	if err != nil {
		return nil, err
	}

	// Total students (subscribers)
	query = `SELECT COUNT(DISTINCT student_id) FROM subscriptions WHERE teacher_id = ?`
	err = DB.QueryRow(query, teacherID).Scan(&stats.TotalStudents)
	if err != nil {
		return nil, err
	}

	// Total views
	query = `
		SELECT COUNT(*) FROM video_views vv
		JOIN videos v ON vv.video_id = v.id
		WHERE v.teacher_id = ?
	`
	err = DB.QueryRow(query, teacherID).Scan(&stats.TotalViews)
	if err != nil {
		return nil, err
	}

	// Recent videos (last 5)
	query = `
		SELECT v.id, v.teacher_id, v.title, v.description, v.filename, v.file_path, 
		       v.thumbnail_path, v.duration, v.file_size, v.created_at, t.name
		FROM videos v
		JOIN teachers t ON v.teacher_id = t.id
		WHERE v.teacher_id = ?
		ORDER BY v.created_at DESC
		LIMIT 5
	`
	rows, err := DB.Query(query, teacherID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var video models.Video
		err := rows.Scan(&video.ID, &video.TeacherID, &video.Title, &video.Description, 
			&video.Filename, &video.FilePath, &video.ThumbnailPath, &video.Duration, 
			&video.FileSize, &video.CreatedAt, &video.TeacherName)
		if err != nil {
			return nil, err
		}
		stats.RecentVideos = append(stats.RecentVideos, video)
	}

	// Recent students (last 5)
	query = `
		SELECT DISTINCT s.id, s.username, s.email, s.name, s.created_at
		FROM students s
		JOIN subscriptions sub ON s.id = sub.student_id
		WHERE sub.teacher_id = ?
		ORDER BY sub.subscribed_at DESC
		LIMIT 5
	`
	rows, err = DB.Query(query, teacherID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var student models.Student
		err := rows.Scan(&student.ID, &student.Username, &student.Email, &student.Name, &student.CreatedAt)
		if err != nil {
			return nil, err
		}
		stats.RecentStudents = append(stats.RecentStudents, student)
	}

	return stats, nil
}

// Get all teachers for student browsing
func GetAllTeachers() ([]models.Teacher, error) {
	query := `SELECT id, username, name, created_at FROM teachers ORDER BY name`
	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var teachers []models.Teacher
	for rows.Next() {
		var teacher models.Teacher
		err := rows.Scan(&teacher.ID, &teacher.Username, &teacher.Name, &teacher.CreatedAt)
		if err != nil {
			return nil, err
		}
		teachers = append(teachers, teacher)
	}
	return teachers, nil
}