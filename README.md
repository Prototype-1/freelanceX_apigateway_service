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

---

## 🗂 Directory Structure

freelanceX_apigateway_service/
├── client/ # gRPC client connections for all backend services
│ ├── auth_client.go
│ ├── profile_client.go
│ ├── ...
│
├── handler/ # HTTP handlers for each functional domain
│ ├── auth_service
│ ├── proposal_service
│ ├── project.crm_service
│ ├── message.notification_service
│ ├── invoice.payment_service
│ ├── timeTracker_service
│ └── ...
│
├── middleware/ # JWT authentication & role-checking middleware
│ └── auth.go
│
├── router/ # HTTP router and route registrations
│ └── router.go
│
├── proto/ # gRPC Protobuf definitions (used to generate clients)
│ ├── auth/
│ ├── profile/
│ ├── portfolio/
│ ├── proposal/
│ ├── crm/
│ └── invoice/
│
├── config/ # Configuration loader (e.g., .env, constants)
│ └── config.go
│
├── .env # Environment variables (e.g. JWT_SECRET, service ports)
├── go.mod / go.sum # Dependencies
└── main.go # Entry point


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
| Auth Service           | Login, Register, OAuth, Role Selection                 | `:50051`  |
| Profile Service        | Freelancer profiles: bio, skills, experience           | `:50051`  |
| Portfolio Service      | Work samples, images, links                            | `:50051`  |
| Review Service         | Client reviews for freelancers                         | `:50051`  |
| Proposal Service       | Proposal and Template creation (MongoDB-backed)        | `:50052`  |
| CRM Service            | Project management: creation, discovery, assignment    | `:50053`  |
| Notification Service   | (Planned) In-app and email notifications               | TBD       |
| Time Tracker Service   | (Planned) Logs, time entries, reports                  | TBD       |
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

##  API Groups
**Please fine the routes.go inside internal dir for more routes and info on middleware integration**

###  Auth (`/api/auth`)
- `POST /register` – User registration
- `POST /login` – User login
- `POST /oauth` – Google OAuth login
- `POST /select-role` – Select role (client/freelancer)
- `GET /me` – Get current user
- `POST /logout` – Logout user

###  Freelancer Portfolio (`/portfolio`)
- `POST /create` – Create portfolio
- `GET /get/:freelancer_id` – Get portfolio by freelancer ID
- `DELETE /delete/:portfolio_id` – Delete portfolio

###  Freelancer Profile (`/profile`)
- `POST /create` – Create profile
- `PUT /update` – Update profile
- `GET /get/:user_id` – Get profile by user ID

###  Reviews (`/api/review`)
- `POST /submit` – Submit review
- `GET /get/:freelancer_id` – Get reviews for a freelancer

###  Proposals (`/proposal`)
- `POST /create` – Create proposal
- `GET /get/:id` – Get proposal by ID
- `PUT /update/:id` – Update proposal
- `GET /listall` – List all proposals
- `POST /template/save` – Save template
- `GET /templates/:freelancer_id` – Get templates

###  Clients (`/api/clients`)
- `POST /create`
- `GET /get/:id`
- `PUT /update/:id`
- `DELETE /delete/:id`

###  Projects (`/api/projects`)
- `POST /create`
- `GET /get/user/:id`
- `GET /get/project/:id`
- `GET /discover/:userId`
- `POST /assign`
- `PUT /update/:id`
- `DELETE /delete/:id`

###  Time Tracker (`/api/time-tracker`)
- `POST /logs/create`
- `GET /logs/user/:userId`
- `GET /logs/project/:projectId`
- `PUT /logs/update/:logId`
- `DELETE /logs/delete/:logId`

###  Message Service (`/api/message`)
- `GET /get/all` – Authenticated messages

###  Milestones (`/milestone`)
- `POST /create`
- `PUT /update`
- `GET /project/:project_id`
- `GET /project/:project_id/phase/:phase`

###  Invoices (`/invoices`)
- `POST /` – Create invoice
- `GET /:id` – Get invoice
- `GET /user/:userId` – All invoices for user
- `GET /project/:projectId` – Invoices for a project
- `PUT /:id/status` – Update invoice status

###  Payments (`/payment`)
- `POST /order` – Create Razorpay order
- `GET /checkout` – Checkout page
- `POST /verify` – Verify payment

---

##  WebSocket

- `GET /ws/messages` – Authenticated WebSocket for real-time messages

---

##  Auth Middleware

- Most routes (except `/auth` and `/payment/checkout`) are protected via `middleware.AuthMiddleware()`.
- Users must include a valid JWT token in headers.

---

#### Maintainers

aswin100396@gmail.com