# Educational Platform API

A backend-only REST API built with Go and Fiber that allows teachers to upload videos and students to subscribe and watch content.

## Features

### For Teachers:
- **Dashboard**: View statistics (total videos, students, views)
- **Video Upload**: Upload video files with title and description
- **Video Management**: View, delete uploaded videos
- **Student Management**: See subscribed students
- **Analytics**: Track video views and engagement
- **Thumbnail Generation**: Automatic thumbnail creation for videos

### For Students:
- **Dashboard**: View available videos from subscribed teachers
- **Teacher Discovery**: Browse and subscribe to teachers
- **Video Watching**: Watch videos from subscribed teachers
- **Subscription Management**: Subscribe/unsubscribe from teachers

## Tech Stack

- **Backend**: Go with Fiber web framework
- **Database**: SQLite with raw SQL queries (no ORM)
- **Frontend**: HTML, CSS, JavaScript (vanilla)
- **Authentication**: Session-based authentication
- **File Storage**: Local file system for videos and thumbnails

## Installation & Setup

1. **Install Dependencies**:
   ```bash
   go mod tidy
   ```

2. **Install FFmpeg** (for thumbnail generation):
   ```bash
   # macOS
   brew install ffmpeg
   
   # Ubuntu/Debian
   sudo apt update
   sudo apt install ffmpeg
   
   # Windows
   # Download from https://ffmpeg.org/download.html
   ```

3. **Run the Application**:
   ```bash
   go run .
   ```

4. **Access the API**:
   - API Base URL: http://localhost:3000/api
   - Health Check: http://localhost:3000/health
   - API Documentation: http://localhost:3000/

## Usage

### Getting Started

1. **Register as a Teacher**:
   - Go to the login page
   - Click "Don't have an account? Register here"
   - Select "Teacher" and fill in your details
   - Login with your credentials

2. **Register as a Student**:
   - Go to the login page
   - Click "Don't have an account? Register here"
   - Select "Student" and fill in your details
   - Login with your credentials

### Teacher Workflow

1. **Upload Videos**:
   - Go to Teacher Dashboard
   - Fill in video title and description
   - Select video file (supports .mp4, .avi, .mov, .mkv, .webm)
   - Click "Upload Video"

2. **Manage Content**:
   - View all uploaded videos
   - Delete videos if needed
   - Check analytics and student engagement

3. **Monitor Students**:
   - See list of subscribed students
   - Track video views and engagement

### Student Workflow

1. **Discover Teachers**:
   - Go to "Browse Teachers" tab
   - See available teachers
   - Subscribe to teachers you're interested in

2. **Watch Videos**:
   - Go to "My Videos" tab
   - See videos from subscribed teachers
   - Click "Watch Video" to view content

3. **Manage Subscriptions**:
   - Go to "My Subscriptions" tab
   - See your current subscriptions
   - Unsubscribe if needed

## API Endpoints

### Authentication
- `POST /api/auth/login` - Login
- `POST /api/auth/register` - Register
- `POST /api/auth/logout` - Logout
- `GET /api/auth/me` - Get current user

### Teacher Endpoints
- `GET /api/teacher/dashboard` - Teacher dashboard stats
- `POST /api/teacher/upload` - Upload video
- `GET /api/teacher/videos` - Get teacher's videos
- `DELETE /api/teacher/videos/:id` - Delete video
- `GET /api/teacher/students` - Get subscribed students
- `GET /api/teacher/analytics` - Get video analytics

### Student Endpoints
- `GET /api/student/dashboard` - Student dashboard
- `GET /api/student/videos` - Get available videos
- `POST /api/student/watch/:id` - Record video view
- `GET /api/student/subscriptions` - Get subscriptions
- `POST /api/student/subscribe/:teacher_id` - Subscribe to teacher
- `DELETE /api/student/unsubscribe/:teacher_id` - Unsubscribe from teacher

### Public Endpoints
- `GET /api/teachers` - Get all teachers
- `GET /api/video/:id` - Serve video file
- `GET /api/video/:id/thumbnail` - Serve thumbnail

## Database Schema

The application uses SQLite with the following tables:

- **teachers**: Teacher accounts
- **students**: Student accounts
- **videos**: Video metadata
- **subscriptions**: Student-teacher relationships
- **video_views**: Video viewing records

## File Structure

```
educational-platform/
├── main.go                 # Main server file
├── database.go            # Database initialization
├── models.go              # Data models
├── queries.go             # Database queries
├── auth.go                # Authentication handlers
├── teacher_handlers.go    # Teacher API endpoints
├── student_handlers.go    # Student API endpoints
├── thumbnail.go           # Thumbnail generation
├── static/                # Frontend files
│   ├── login.html
│   ├── teacher_dashboard.html
│   └── student_dashboard.html
├── uploads/               # File storage
│   ├── videos/           # Video files
│   └── thumbnails/       # Thumbnail files
└── educational_platform.db # SQLite database
```

## Configuration

- **Port**: Default is 3000, can be changed with `PORT` environment variable
- **Database**: SQLite file `educational_platform.db` in project root
- **File Storage**: `./uploads/` directory for videos and thumbnails

## Security Notes

This is an MVP version with basic security. For production use, consider:

- Password hashing improvements (bcrypt instead of SHA256)
- File upload validation and virus scanning
- Rate limiting
- HTTPS enforcement
- Input sanitization
- SQL injection prevention (though raw queries are used carefully)

## Troubleshooting

1. **Thumbnail generation fails**: Ensure FFmpeg is installed and accessible
2. **Database errors**: Check file permissions for SQLite database
3. **File upload fails**: Ensure `uploads/` directory exists and is writable
4. **Port already in use**: Change the PORT environment variable

## License

This project is for educational purposes. Feel free to modify and use as needed.

#### This F**king Shit is Made By Cursor "just promting no manual coding".
