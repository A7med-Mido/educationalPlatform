# Educational Platform MVP - Complete Usage Guide

## 🎉 Application Successfully Restructured and Fixed!

Your educational platform is now fully functional with a clean, organized codebase. Here's everything you need to know to use it.

## 📁 Project Structure

```
educational-platform/
├── main.go                    # Main server file
├── go.mod                     # Go dependencies
├── go.sum                     # Dependency checksums
├── README.md                  # This file
├── models/                    # Data models
│   └── models.go
├── database/                  # Database layer
│   ├── database.go           # Database initialization
│   └── queries.go            # All SQL queries
├── handlers/                  # HTTP handlers
│   ├── auth.go               # Authentication handlers
│   ├── teacher_handlers.go   # Teacher-specific endpoints
│   ├── student_handlers.go   # Student-specific endpoints
│   └── thumbnail.go          # Thumbnail generation
├── static/                    # Frontend files
│   ├── login.html
│   ├── teacher_dashboard.html
│   └── student_dashboard.html
└── uploads/                   # File storage
    ├── videos/               # Video files
    └── thumbnails/           # Thumbnail files
```

## 🚀 How to Run the Application

### 1. Prerequisites
- Go 1.25.0 or later
- FFmpeg (optional, for thumbnail generation)

### 2. Installation
```bash
# Navigate to the project directory
cd /Users/ahmed/development/educationalPlatform

# Install dependencies
go mod tidy

# Build the application
go build -o educational-platform .

# Run the application
./educational-platform
```

### 3. Access the Application
- **Login Page**: http://localhost:3000/login.html
- **Teacher Dashboard**: http://localhost:3000/teacher_dashboard.html
- **Student Dashboard**: http://localhost:3000/student_dashboard.html

## 👥 User Roles and Workflows

### 🔐 Authentication
1. **Register**: Create a new account (Teacher or Student)
2. **Login**: Use your username and password
3. **Logout**: End your session

### 👨‍🏫 Teacher Workflow

#### Getting Started
1. Go to http://localhost:3000/login.html
2. Click "Don't have an account? Register here"
3. Select "Teacher" and fill in your details
4. Login with your credentials

#### Managing Content
1. **Upload Videos**:
   - Go to Teacher Dashboard
   - Fill in video title and description
   - Select video file (supports .mp4, .avi, .mov, .mkv, .webm)
   - Click "Upload Video"
   - Thumbnail will be generated automatically (if FFmpeg is installed)

2. **View Dashboard**:
   - See total videos, students, and views
   - View recent videos and students
   - Monitor engagement statistics

3. **Manage Videos**:
   - View all uploaded videos
   - Delete videos if needed
   - Check video analytics

4. **Monitor Students**:
   - See list of subscribed students
   - Track video views and engagement

### 🎓 Student Workflow

#### Getting Started
1. Go to http://localhost:3000/login.html
2. Click "Don't have an account? Register here"
3. Select "Student" and fill in your details
4. Login with your credentials

#### Discovering Content
1. **Browse Teachers**:
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

## 🔧 API Endpoints

### Authentication
- `POST /api/auth/login` - Login
- `POST /api/auth/register` - Register
- `POST /api/auth/logout` - Logout
- `GET /api/auth/me` - Get current user info

### Teacher Endpoints (Requires Teacher Authentication)
- `GET /api/teacher/dashboard` - Teacher dashboard stats
- `POST /api/teacher/upload` - Upload video
- `GET /api/teacher/videos` - Get teacher's videos
- `DELETE /api/teacher/videos/:id` - Delete video
- `GET /api/teacher/students` - Get subscribed students
- `GET /api/teacher/analytics` - Get video analytics

### Student Endpoints (Requires Student Authentication)
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

## 🗄️ Database Schema

The application uses SQLite with these tables:

- **teachers**: Teacher accounts
- **students**: Student accounts  
- **videos**: Video metadata
- **subscriptions**: Student-teacher relationships
- **video_views**: Video viewing records

## 📁 File Storage

- **Videos**: Stored in `./uploads/videos/`
- **Thumbnails**: Stored in `./uploads/thumbnails/`
- **Database**: `./educational_platform.db`

## 🛠️ Configuration

- **Port**: Default is 3000, change with `PORT` environment variable
- **Database**: SQLite file `educational_platform.db` in project root
- **File Storage**: `./uploads/` directory for videos and thumbnails

## 🔒 Security Features

- Session-based authentication
- Password hashing (SHA256)
- File type validation for uploads
- User authorization checks
- SQL injection prevention

## 🐛 Troubleshooting

### Common Issues

1. **Thumbnail generation fails**:
   - Install FFmpeg: `brew install ffmpeg` (macOS) or `sudo apt install ffmpeg` (Ubuntu)
   - The app will create placeholder thumbnails if FFmpeg is not available

2. **Database errors**:
   - Check file permissions for SQLite database
   - Ensure the database file is not locked by another process

3. **File upload fails**:
   - Ensure `uploads/` directory exists and is writable
   - Check file size limits
   - Verify file type is supported (.mp4, .avi, .mov, .mkv, .webm)

4. **Port already in use**:
   - Change the PORT environment variable: `PORT=8080 ./educational-platform`
   - Or kill the existing process: `pkill -f educational-platform`

### Testing the Application

1. **Test API endpoints**:
   ```bash
   # Test authentication
   curl -X POST http://localhost:3000/api/auth/register \
     -H "Content-Type: application/json" \
     -d '{"username":"testteacher","email":"test@example.com","password":"password123","name":"Test Teacher","user_type":"teacher"}'
   
   # Test login
   curl -X POST http://localhost:3000/api/auth/login \
     -H "Content-Type: application/json" \
     -d '{"username":"testteacher","password":"password123"}'
   ```

2. **Test file serving**:
   ```bash
   # Test login page
   curl http://localhost:3000/login.html
   
   # Test API
   curl http://localhost:3000/api/auth/me
   ```

## 🎯 Key Features Implemented

✅ **Complete Authentication System**
- User registration and login
- Session management
- Role-based access control

✅ **Teacher Dashboard**
- Video upload with thumbnail generation
- Student management
- Analytics and statistics
- Video management (view, delete)

✅ **Student Dashboard**
- Teacher discovery and subscription
- Video browsing and watching
- Subscription management

✅ **Database Layer**
- SQLite with raw SQL queries (no ORM)
- Proper table relationships
- Efficient querying

✅ **File Management**
- Video upload and storage
- Thumbnail generation
- File serving

✅ **Modern Frontend**
- Responsive HTML/CSS
- JavaScript for API interaction
- Clean, professional UI

## 🚀 Next Steps for Enhancement

1. **Security Improvements**:
   - Use bcrypt for password hashing
   - Add rate limiting
   - Implement CSRF protection
   - Add input validation

2. **Features**:
   - Video streaming optimization
   - Comments and ratings
   - Course organization
   - Progress tracking

3. **Infrastructure**:
   - Docker containerization
   - Production deployment
   - Database migrations
   - Logging and monitoring

## 📞 Support

The application is now fully functional and ready for use! All major bugs have been fixed, and the codebase is properly organized for future development.

**Happy Teaching and Learning! 🎓📚**
