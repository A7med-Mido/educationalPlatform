# Educational Platform API Documentation

A backend-only REST API for an educational platform that allows teachers to upload videos and students to subscribe and watch content.

## 🚀 Quick Start

```bash
# Build and run the application
go build -o educational-platform .
./educational-platform

# Or use the start script
./start.sh
```

**Base URL**: `http://localhost:3000/api`

## 📋 API Endpoints

### Authentication

#### Register User
```http
POST /api/auth/register
Content-Type: application/json

{
  "username": "string",
  "email": "string", 
  "password": "string",
  "name": "string",
  "user_type": "teacher" | "student"
}
```

**Response:**
```json
{
  "success": true,
  "message": "Account created successfully"
}
```

#### Login
```http
POST /api/auth/login
Content-Type: application/json

{
  "username": "string",
  "password": "string"
}
```

**Response:**
```json
{
  "success": true,
  "message": "Login successful",
  "data": {
    "user_id": 1,
    "user_type": "teacher",
    "username": "teacher1",
    "name": "Teacher Name"
  }
}
```

#### Get Current User
```http
GET /api/auth/me
Cookie: session_id=<session_id>
```

**Response:**
```json
{
  "success": true,
  "data": {
    "user_id": 1,
    "user_type": "teacher",
    "username": "teacher1",
    "name": "Teacher Name"
  }
}
```

#### Logout
```http
POST /api/auth/logout
Cookie: session_id=<session_id>
```

**Response:**
```json
{
  "success": true,
  "message": "Logged out successfully"
}
```

### Teacher Endpoints (Requires Teacher Authentication)

#### Get Dashboard Stats
```http
GET /api/teacher/dashboard
Cookie: session_id=<session_id>
```

**Response:**
```json
{
  "success": true,
  "data": {
    "total_videos": 5,
    "total_students": 12,
    "total_views": 45,
    "recent_videos": [...],
    "recent_students": [...]
  }
}
```

#### Upload Video
```http
POST /api/teacher/upload
Content-Type: multipart/form-data
Cookie: session_id=<session_id>

Form Data:
- title: "Video Title"
- description: "Video Description"
- video: <video_file>
```

**Response:**
```json
{
  "success": true,
  "message": "Video uploaded successfully"
}
```

#### Get Teacher's Videos
```http
GET /api/teacher/videos
Cookie: session_id=<session_id>
```

**Response:**
```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "title": "Introduction to Programming",
      "description": "Basic programming concepts",
      "filename": "intro.mp4",
      "file_path": "./uploads/videos/video_1_1234567890.mp4",
      "thumbnail_path": "./uploads/thumbnails/video_1_1234567890_thumb.jpg",
      "duration": 1200,
      "file_size": 52428800,
      "created_at": "2025-10-15T02:30:00Z",
      "teacher_name": "John Doe"
    }
  ]
}
```

#### Delete Video
```http
DELETE /api/teacher/videos/{id}
Cookie: session_id=<session_id>
```

**Response:**
```json
{
  "success": true,
  "message": "Video deleted successfully"
}
```

#### Get Subscribed Students
```http
GET /api/teacher/students
Cookie: session_id=<session_id>
```

**Response:**
```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "student_id": 5,
      "teacher_id": 1,
      "subscribed_at": "2025-10-15T02:30:00Z",
      "student_name": "Jane Smith"
    }
  ]
}
```

#### Get Video Analytics
```http
GET /api/teacher/analytics
Cookie: session_id=<session_id>
```

**Response:**
```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "student_id": 5,
      "video_id": 1,
      "watched_at": "2025-10-15T02:30:00Z",
      "video_title": "Introduction to Programming",
      "student_name": "Jane Smith"
    }
  ]
}
```

### Student Endpoints (Requires Student Authentication)

#### Get Student Dashboard
```http
GET /api/student/dashboard
Cookie: session_id=<session_id>
```

**Response:**
```json
{
  "success": true,
  "data": {
    "subscriptions": [...],
    "videos": [...]
  }
}
```

#### Get Available Videos
```http
GET /api/student/videos
Cookie: session_id=<session_id>
```

**Response:**
```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "title": "Introduction to Programming",
      "description": "Basic programming concepts",
      "teacher_name": "John Doe",
      "created_at": "2025-10-15T02:30:00Z"
    }
  ]
}
```

#### Watch Video (Record View)
```http
POST /api/student/watch/{id}
Cookie: session_id=<session_id>
```

**Response:**
```json
{
  "success": true,
  "data": {
    "id": 1,
    "title": "Introduction to Programming",
    "description": "Basic programming concepts",
    "file_path": "./uploads/videos/video_1_1234567890.mp4"
  }
}
```

#### Get Subscriptions
```http
GET /api/student/subscriptions
Cookie: session_id=<session_id>
```

**Response:**
```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "student_id": 5,
      "teacher_id": 1,
      "subscribed_at": "2025-10-15T02:30:00Z",
      "teacher_name": "John Doe"
    }
  ]
}
```

#### Subscribe to Teacher
```http
POST /api/student/subscribe/{teacher_id}
Cookie: session_id=<session_id>
```

**Response:**
```json
{
  "success": true,
  "message": "Successfully subscribed to teacher"
}
```

#### Unsubscribe from Teacher
```http
DELETE /api/student/unsubscribe/{teacher_id}
Cookie: session_id=<session_id>
```

**Response:**
```json
{
  "success": true,
  "message": "Successfully unsubscribed from teacher"
}
```

### Public Endpoints

#### Get All Teachers
```http
GET /api/teachers
```

**Response:**
```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "username": "teacher1",
      "name": "John Doe",
      "created_at": "2025-10-15T02:30:00Z"
    }
  ]
}
```

#### Serve Video File
```http
GET /api/video/{id}
```

Returns the video file for streaming/downloading.

#### Serve Thumbnail
```http
GET /api/video/{id}/thumbnail
```

Returns the thumbnail image for the video.

### System Endpoints

#### Health Check
```http
GET /health
```

**Response:**
```json
{
  "success": true,
  "message": "Educational Platform API is running",
  "data": {
    "version": "1.0.0",
    "status": "healthy"
  }
}
```

#### API Documentation
```http
GET /
```

**Response:**
```json
{
  "success": true,
  "message": "Educational Platform API",
  "data": {
    "version": "1.0.0",
    "endpoints": {
      "authentication": "/api/auth",
      "teachers": "/api/teacher",
      "students": "/api/student",
      "public": "/api/teachers, /api/video",
      "health": "/health"
    },
    "documentation": "See README.md for API documentation"
  }
}
```

## 🔐 Authentication

The API uses session-based authentication with HTTP cookies:

1. **Login**: Send credentials to `/api/auth/login`
2. **Session Cookie**: Server sets `session_id` cookie
3. **Authenticated Requests**: Include the cookie in subsequent requests
4. **Logout**: Call `/api/auth/logout` to invalidate session

## 📁 File Upload

Video uploads use `multipart/form-data` with the following fields:
- `title`: Video title (required)
- `description`: Video description (optional)
- `video`: Video file (required, supports .mp4, .avi, .mov, .mkv, .webm)

## 🗄️ Database Schema

- **teachers**: Teacher accounts
- **students**: Student accounts
- **videos**: Video metadata
- **subscriptions**: Student-teacher relationships
- **video_views**: Video viewing records

## 🛠️ Error Handling

All endpoints return consistent error responses:

```json
{
  "success": false,
  "message": "Error description"
}
```

Common HTTP status codes:
- `200`: Success
- `400`: Bad Request (validation errors)
- `401`: Unauthorized (not authenticated)
- `403`: Forbidden (insufficient permissions)
- `404`: Not Found
- `500`: Internal Server Error

## 🧪 Testing the API

### Using curl:

```bash
# Register a teacher
curl -X POST http://localhost:3000/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{"username":"teacher1","email":"teacher@example.com","password":"password123","name":"Teacher One","user_type":"teacher"}'

# Login
curl -X POST http://localhost:3000/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username":"teacher1","password":"password123"}'

# Upload video (with session cookie)
curl -X POST http://localhost:3000/api/teacher/upload \
  -F "title=My Video" \
  -F "description=Video description" \
  -F "video=@/path/to/video.mp4" \
  -b "session_id=<session_id>"
```

### Using Postman:

1. Import the API endpoints
2. Set up environment variables for base URL
3. Use the session cookie for authenticated requests

## 🚀 Production Considerations

- Use HTTPS in production
- Implement rate limiting
- Add input validation and sanitization
- Use environment variables for configuration
- Set up proper logging and monitoring
- Consider using a production database (PostgreSQL, MySQL)
- Implement proper password hashing (bcrypt)
- Add CSRF protection
