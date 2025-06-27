#!/bin/bash

echo "🚀 Starting Zico Token Frontend..."

# Navigate to frontend directory
cd frontend

# Check if node_modules exists
if [ ! -d "node_modules" ]; then
    echo "📦 Installing dependencies..."
    npm install
fi

# Start the frontend
echo "🌐 Starting React development server..."
npm start 