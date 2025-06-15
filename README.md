# Blockchain Public Sector Application

This project consists of a frontend built with Nuxt.js and a backend built with Go, using Firebase for authentication and data storage.

## Project Structure

### Frontend
- Built with Nuxt.js and Quasar Framework
- Uses Firebase Authentication for user management
- Uses Firestore for data storage
- Implements two separate use cases:
  - Diploma verification system
  - State documents system

### Backend
- Built with Go
- Implements blockchain verification for both diploma and state documents
- Uses Firebase Admin SDK for authentication

## Setup Instructions

### Prerequisites
- Node.js and npm installed
- Go installed
- Firebase project set up with Authentication and Firestore enabled

### Frontend Setup

1. Install dependencies:
   ```
   cd frontend
   npm install
   ```

2. Create a `.env` file with your Firebase configuration:
   ```
   FIREBASE_API_KEY=your_api_key
   FIREBASE_AUTH_DOMAIN=your_auth_domain
   FIREBASE_PROJECT_ID=your_project_id
   FIREBASE_STORAGE_BUCKET=your_storage_bucket
   FIREBASE_MESSAGING_SENDER_ID=your_messaging_sender_id
   FIREBASE_APP_ID=your_app_id
   ```

3. Run the development server:
   ```
   npm run dev
   ```

### Backend Setup

1. Install dependencies:
   ```
   cd backend
   go mod download
   ```

2. Create a `serviceAccountKey.json` file with your Firebase Admin SDK credentials

3. Run the server:
   ```
   go run main.go
   ```

## Features

### Diploma Verification System
- User registration and login for universities
- Add and manage diplomas
- Verify diploma authenticity
- View diploma details

### State Documents System
- User registration and login for state users
- Add and manage state documents
- Verify document authenticity
- View document details

## Environment Variables

### Frontend
- FIREBASE_API_KEY
- FIREBASE_AUTH_DOMAIN
- FIREBASE_PROJECT_ID
- FIREBASE_STORAGE_BUCKET
- FIREBASE_MESSAGING_SENDER_ID
- FIREBASE_APP_ID

### Backend
- FIREBASE_PROJECT_ID
- FIREBASE_PRIVATE_KEY
- FIREBASE_CLIENT_EMAIL

## Postman Collection

A Postman collection (`Blockchain Public Sector API.postman_collection.json`) is included in the `backend` directory. This collection contains API requests for testing the backend endpoints, including:
- User registration and login
- Adding and managing diplomas
- Verifying diploma authenticity
- Viewing diploma details

You can import this collection into Postman to quickly test the backend API.
