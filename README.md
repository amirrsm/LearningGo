# 🧭 30-Day Golang Roadmap  
**Goal:** Become an intermediate Go backend developer ready to build production-level services.

---

## 🧩 Week 1 — Go Fundamentals
**Goal:** Build a solid foundation in syntax, types, and core Go concepts.

### Day 1
- [ ] Install Go & set up workspace (`go env`, `GOPATH`)
- [ ] Learn Go compilation & running (`go run`, `go build`)
- [ ] Explore [Tour of Go — Basics](https://go.dev/tour/basics/1)
- [ ] 🧠 **Project:** Print “Hello, Amirreza” and current time.

### Day 2
- [ ] Learn variables, constants, and zero values
- [ ] Explore arrays, slices, and maps
- [ ] 🧠 **Mini Project:** CLI that counts word frequency in a string

### Day 3
- [ ] Study control structures: `if`, `for`, `switch`
- [ ] Learn about functions (multiple return values)
- [ ] 🧠 **Mini Project:** Simple calculator CLI (+, -, *, /)

### Day 4
- [ ] Understand structs, methods, and pointers
- [ ] 🧠 **Mini Project:** Define `User` struct and display method

### Day 5
- [ ] Learn interfaces & type assertions
- [ ] Understand package organization & imports
- [ ] 🧠 **Mini Project:** `Shape` interface with `Area()` method

### Day 6
- [ ] Learn idiomatic error handling
- [ ] Create custom error types
- [ ] 🧠 **Mini Project:** File reader CLI with graceful error reporting

### Day 7
- [ ] Revise all week 1 topics
- [ ] 🧠 **Challenge:** Concurrent line counter for text files

---

## ⚙️ Week 2 — Concurrency & Web Basics
**Goal:** Understand goroutines, channels, and basic HTTP APIs.

### Day 8
- [ ] Learn goroutines (`go func()`)
- [ ] 🧠 **Task:** Launch 10 goroutines printing their ID

### Day 9
- [ ] Learn channels and `WaitGroup`
- [ ] 🧠 **Project:** Fetch several URLs concurrently

### Day 10
- [ ] Study `select` statement, buffered vs unbuffered channels
- [ ] 🧠 **Project:** Concurrent “ping checker” CLI

### Day 11
- [ ] Learn Go modules (`go mod init`, `go get`)
- [ ] Understand folder structure (`cmd/`, `pkg/`, `internal/`)
- [ ] 🧠 **Task:** Convert one tool into a Go module

### Day 12
- [ ] Learn `net/http`: handlers, responses, JSON
- [ ] 🧠 **Mini Project:** REST API `/ping` → `{ "status": "ok" }`

### Day 13
- [ ] Parse requests & bind structs
- [ ] Create simple middleware
- [ ] 🧠 **Project:** API accepting JSON user data → returns greeting

### Day 14
- [ ] Learn Gin or Chi web frameworks
- [ ] 🧠 **Project:** `/users` route storing users in memory

---

## 🧱 Week 3 — Real Backend Development
**Goal:** Connect your Go API to a database and implement authentication.

### Day 15
- [ ] Learn `database/sql` and connect to PostgreSQL
- [ ] 🧠 **Task:** List users from DB

### Day 16
- [ ] Learn GORM ORM (models, CRUD, migrations)
- [ ] 🧠 **Project:** CRUD for users using Gin + GORM

### Day 17
- [ ] Implement JWT authentication (`golang-jwt/jwt/v5`)
- [ ] 🧠 **Project:** Login → JWT → Middleware auth

### Day 18
- [ ] Learn config management (`godotenv` or `viper`)
- [ ] 🧠 **Task:** Move DB credentials and JWT secret to `.env`

### Day 19
- [ ] Learn request validation (`go-playground/validator`)
- [ ] 🧠 **Project:** Validate signup input (email, password length)

### Day 20
- [ ] Learn clean architecture (handler/service/repo)
- [ ] 🧠 **Task:** Refactor CRUD to clean structure

### Day 21
- [ ] Test API manually (Postman or curl)
- [ ] 🧠 **Task:** Deploy locally via Docker Compose (Go + Postgres)

---

## 🚀 Week 4 — Production-Level Backend
**Goal:** Testing, logging, optimization, and deployment.

### Day 22
- [ ] Learn Go testing (`testing`, `httptest`, assertions)
- [ ] 🧠 **Task:** Unit tests for handlers & utils

### Day 23
- [ ] Implement graceful shutdown (`context.WithTimeout`)
- [ ] 🧠 **Project:** Add proper shutdown signal handling

### Day 24
- [ ] Structured logging with `zap` or `zerolog`
- [ ] Error wrapping (`fmt.Errorf("%w")`)
- [ ] 🧠 **Task:** Add consistent error handling & logs

### Day 25
- [ ] Learn profiling & benchmarking (`pprof`)
- [ ] 🧠 **Task:** Benchmark concurrent function

### Day 26
- [ ] Learn dependency injection (manual)
- [ ] 🧠 **Refactor:** Inject repositories & configs cleanly

### Day 27
- [ ] Dockerize your entire backend
- [ ] 🧠 **Project:** Dockerfile + docker-compose.yml (Go + Postgres)

### Day 28
- [ ] Write integration tests (real DB)
- [ ] 🧠 **Task:** Test signup → login → get user flow

### Day 29
- [ ] Refactor, clean up, and document project
- [ ] 🧠 **Task:** Add README + Makefile + folder cleanup

### Day 30
- [ ] Deploy app (Render, Fly.io, or Docker on your server)
- [ ] Review architecture, concurrency, API design
- [ ] 🧠 **Outcome:** 🎉 Fully working, production-ready Go backend

---

## 🏁 After 30 Days You Will:
- ✅ Understand Go syntax, interfaces, and concurrency  
- ✅ Build and deploy production APIs using Gin & GORM  
- ✅ Use JWT, validation, configuration, and Docker  
- ✅ Be an **intermediate Go developer**, ready for real backend work  

---

> 👨‍🏫 Mentorship Tip:  
> Push code daily to GitHub — even small commits.  
> Consistency > intensity.  
> I’ll help you refine your project structure when you start the backend build in Week 3.
