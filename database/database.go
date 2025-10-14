package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDatabase() error {
	var err error
	DB, err = sql.Open("sqlite3", "./educational_platform.db")
	if err != nil {
		return err
	}

	// Create tables
	err = createTables()
	if err != nil {
		return err
	}

	return nil
}

func createTables() error {
	// Teachers table
	teachersTable := `
	CREATE TABLE IF NOT EXISTS teachers (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username VARCHAR(50) UNIQUE NOT NULL,
		email VARCHAR(100) UNIQUE NOT NULL,
		password_hash VARCHAR(255) NOT NULL,
		name VARCHAR(100) NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	// Students table
	studentsTable := `
	CREATE TABLE IF NOT EXISTS students (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username VARCHAR(50) UNIQUE NOT NULL,
		email VARCHAR(100) UNIQUE NOT NULL,
		password_hash VARCHAR(255) NOT NULL,
		name VARCHAR(100) NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	// Videos table
	videosTable := `
	CREATE TABLE IF NOT EXISTS videos (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		teacher_id INTEGER NOT NULL,
		title VARCHAR(200) NOT NULL,
		description TEXT,
		filename VARCHAR(255) NOT NULL,
		file_path VARCHAR(500) NOT NULL,
		thumbnail_path VARCHAR(500),
		duration INTEGER, -- in seconds
		file_size INTEGER, -- in bytes
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (teacher_id) REFERENCES teachers(id) ON DELETE CASCADE
	);`

	// Subscriptions table
	subscriptionsTable := `
	CREATE TABLE IF NOT EXISTS subscriptions (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		student_id INTEGER NOT NULL,
		teacher_id INTEGER NOT NULL,
		subscribed_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (student_id) REFERENCES students(id) ON DELETE CASCADE,
		FOREIGN KEY (teacher_id) REFERENCES teachers(id) ON DELETE CASCADE,
		UNIQUE(student_id, teacher_id)
	);`

	// Video views table (to track which students watched which videos)
	videoViewsTable := `
	CREATE TABLE IF NOT EXISTS video_views (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		student_id INTEGER NOT NULL,
		video_id INTEGER NOT NULL,
		watched_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (student_id) REFERENCES students(id) ON DELETE CASCADE,
		FOREIGN KEY (video_id) REFERENCES videos(id) ON DELETE CASCADE,
		UNIQUE(student_id, video_id)
	);`

	tables := []string{teachersTable, studentsTable, videosTable, subscriptionsTable, videoViewsTable}

	for _, table := range tables {
		_, err := DB.Exec(table)
		if err != nil {
			return fmt.Errorf("error creating table: %v", err)
		}
	}

	// Create uploads directory if it doesn't exist
	err := os.MkdirAll("./uploads/videos", 0755)
	if err != nil {
		return err
	}

	err = os.MkdirAll("./uploads/thumbnails", 0755)
	if err != nil {
		return err
	}

	log.Println("Database tables created successfully")
	return nil
}

func CloseDatabase() {
	if DB != nil {
		DB.Close()
	}
}
