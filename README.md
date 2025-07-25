# Hackathon PVC - Registration Platform - (In Development)

## Navigation
- [Run The Project](#run)
- [Go To Context](#context)
- [Go to Planning](#planning)

---


## Run

### Prerequisites
- **Node.js** (for frontend)
- **Go 1.21+** (for backend)

### Install Dependencies

**Frontend:**
```bash
cd front
npm install
```

**Backend:**
```bash
cd backend
go mod tidy
```

### Start Development Servers

**Frontend (React):**
```bash
cd front
npm run dev
```
*Runs on: http://localhost:5173*

**Backend (Go API):**
```bash
cd backend
go run cmd/main.go
```
*Runs on: http://localhost:8080*

### Access the Application
- **Frontend**: http://localhost:5173
- **API Health Check**: http://localhost:8080/api/health
- **Registration API**: http://localhost:8080/api/registrations


## Context

### Landing Page Purpose
A registration platform for hackathon participants where developers, designers, and entrepreneurs can:
- Sign up for the event
- Share their project ideas
- Find potential teammates based on skills and interests
- View available prizes and event information

### Current Architecture
**Monolith Structure:**
- **Frontend**: React app with registration forms and event information
- **Backend**: Go API handling registrations and data storage
- **Database**: SQLite for participant data

The backend provides a REST API that the frontend consumes, with clean separation between presentation and business logic.

---

## Planning

### Short-term Goals
1. **Complete Registration System**: Finish all CRUD operations for participant management
2. **Team Matching**: Algorithm to suggest potential teammates based on skills and project ideas
3. **Admin Dashboard**: Interface for event organizers to manage registrations

### Deployment Strategy
**Single Binary Approach:**
- Embed the React frontend into the Go backend
- Generate one executable file containing both frontend and backend
- Deploy to AWS EC2 as a single, self-contained application
- Benefits: Simple deployment, no separate hosting needed, easier scaling

### Future Enhancements
- Real-time notifications for team matching
- Project submission and voting system
- Integration with external APIs (GitHub, LinkedIn)
- Multi-event support for different hackathons

---

**Tech Stack**: Go, React, SQLite → Single Binary for EC2
