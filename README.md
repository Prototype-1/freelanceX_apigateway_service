# FreelanceX API Gateway Service

The **FreelanceX API Gateway** acts as the unified entry point for all client and frontend requests in the FreelanceX ecosystem. It routes HTTP traffic to the appropriate backend gRPC services, applies request validation, handles role-based authorization, and abstracts away internal service communication details.

---

##  Overview

FreelanceX is a microservices-based freelance collaboration platform. The API Gateway is responsible for:

- Routing HTTP requests to gRPC microservices (via gRPC clients).
- Parsing and validating request bodies and query params.
- Injecting authorization metadata (JWT-based) into gRPC calls.
- Managing public vs. protected endpoints.
- Providing a clean RESTful interface to the frontend (or external consumers) while hiding service complexity.

---

##  Tech Stack

- **Language:** Go (Golang)
- **Framework:** Gin (HTTP router)
- **gRPC Client:** Protobuf-generated stubs for all backend services
- **JWT Auth:** Middleware-based token validation and role enforcement
- **Reverse Proxy Pattern:** Converts REST calls to internal gRPC
- - **Kafka events**
- **Razorpay integrated**
- **Gorilla websocket**
- **SWAGGER API**
- **Prometheus & Graffana**

---

##  Features

- JWT Authentication Middleware
- Role-based Access Control (Client / Freelancer / Admin)
- Request → gRPC translation
- Environment-based service addresses
- Centralized error handling
- Clean modular routing

---

##  Dependencies

- [Gin](https://github.com/gin-gonic/gin) - HTTP router
- [Go gRPC](https://grpc.io/docs/languages/go/)
- [Protobuf Compiler](https://grpc.io/docs/protoc-installation/)
- [godotenv](https://github.com/joho/godotenv) - for loading `.env`
- [redis](github.com/redis/go-redis/v9)
- [jwt](github.com/golang-jwt/jwt/v4)
- [metadata](cloud.google.com/go/compute/metadata)
- [kafka](github.com/Shopify/sarama)
- [websocket](github.com/gorilla/websocket)

---

##  Supported Services

| Service                 | Description                                            | gRPC Port |
|------------------------|--------------------------------------------------------|-----------|
| User Service           | Login, Register, OAuth, Role Selection                 | `:50051`  |
| Proposal Service       | Proposal and Template creation (MongoDB-backed)        | `:50052`  |
| CRM Service            | Project management: creation, discovery, assignment    | `:50053`  |
| Notification Service   | (Planned) In-app and email notifications               | `:50054`  |
| Time Tracker Service   | (Planned) Logs, time entries, reports                  | `:50055` |
| Invoice & Payment Svc  | Invoice generation, payment tracking                   | `:50056`  |

---

##  Setup Instructions

### 1. Clone & Navigate

```bash
git clone https://github.com/Prototype-1/freelancex_apigateway_service.git
cd freelancex_project/freelanceX_apigateway_service
```

## 2. Create .env File

PORT=8080
JWT_SECRET=your_jwt_secret_key
REDIS_ADDR=localhost:6379
GOOGLE_CLIENT_ID
GOOGLE_CLIENT_SECRET
GOOGLE_REDIRECT_URL
RAZORPAY_KEY_ID
RAZORPAY_KEY_SECRET

##  3. Install Dependencies

go mod tidy

## 4. Generate gRPC Clients from Protos (if needed)

protoc --go_out=. --go-grpc_out=. proto/auth/user.proto
# Repeat for each proto

## 5. Run API Gateway

go run main.go

### Running in Production

    Use go build to build a binary:

go build -o apigateway
./apigateway
---

##  Auth Middleware

- Most routes (except `/auth` and `/payment/checkout`) are protected via `middleware.AuthMiddleware()`.
- Users must include a valid JWT token in headers.

---

### For full API documentation
 see api.md in the root or visit https://freelancex.goxtrace.shop/swagger/index.html (partial routers for now)

#### Maintainers

aswin100396@gmail.com