#!/bin/bash

# Educational Platform Quick Start Script

echo "🚀 Starting Educational Platform..."

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "❌ Go is not installed. Please install Go 1.25.0 or later."
    exit 1
fi

# Check if the binary exists, if not build it
if [ ! -f "./educational-platform" ]; then
    echo "📦 Building application..."
    go build -o educational-platform .
    if [ $? -ne 0 ]; then
        echo "❌ Build failed. Please check the error messages above."
        exit 1
    fi
    echo "✅ Build successful!"
fi

# Create uploads directories if they don't exist
mkdir -p uploads/videos uploads/thumbnails

# Check if FFmpeg is available for thumbnail generation
if command -v ffmpeg &> /dev/null; then
    echo "✅ FFmpeg detected - thumbnail generation enabled"
else
    echo "⚠️  FFmpeg not found - placeholder thumbnails will be used"
    echo "   Install FFmpeg for better thumbnail generation:"
    echo "   macOS: brew install ffmpeg"
    echo "   Ubuntu: sudo apt install ffmpeg"
fi

# Start the server
echo "🌐 Starting server on http://localhost:3000"
echo "📚 Teacher Dashboard: http://localhost:3000/teacher_dashboard.html"
echo "🎓 Student Dashboard: http://localhost:3000/student_dashboard.html"
echo "🔐 Login Page: http://localhost:3000/login.html"
echo ""
echo "Press Ctrl+C to stop the server"
echo ""

./educational-platform
