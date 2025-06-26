##  API Groups
*Please fine the routes.go inside internal dir for more routes and info on middleware integration*

###  Auth (/api/auth)
- POST /register – User registration
- POST /login – User login
- POST /oauth – Google OAuth login
- POST /select-role – Select role (client/freelancer)
- GET /me – Get current user
- POST /logout – Logout user

###  Freelancer Portfolio (/portfolio)
- POST /create – Create portfolio
- GET /get/:freelancer_id – Get portfolio by freelancer ID
- DELETE /delete/:portfolio_id – Delete portfolio

###  Freelancer Profile (/profile)
- POST /create – Create profile
- PUT /update – Update profile
- GET /get/:user_id – Get profile by user ID

###  Reviews (/api/review)
- POST /submit – Submit review
- GET /get/:freelancer_id – Get reviews for a freelancer

###  Proposals (/proposal)
- POST /create – Create proposal
- GET /get/:id – Get proposal by ID
- PUT /update/:id – Update proposal
- GET /listall – List all proposals
- POST /template/save – Save template
- GET /templates/:freelancer_id – Get templates

###  Clients (/api/clients)
- POST /create
- GET /get/:id
- PUT /update/:id
- DELETE /delete/:id

###  Projects (/api/projects)
- POST /create
- GET /get/user/:id
- GET /get/project/:id
- GET /discover/:userId
- POST /assign
- PUT /update/:id
- DELETE /delete/:id

###  Time Tracker (/api/time-tracker)
- POST /logs/create
- GET /logs/user/:userId
- GET /logs/project/:projectId
- PUT /logs/update/:logId
- DELETE /logs/delete/:logId

###  Message Service (/api/message)
- GET /get/all – Authenticated messages

###  Milestones (/milestone)
- POST /create
- PUT /update
- GET /project/:project_id
- GET /project/:project_id/phase/:phase

###  Invoices (/invoices)
- POST / – Create invoice
- GET /:id – Get invoice
- GET /user/:userId – All invoices for user
- GET /project/:projectId – Invoices for a project
- PUT /:id/status – Update invoice status

###  Payments (/payment)
- POST /order – Create Razorpay order
- GET /checkout – Checkout page
- POST /verify – Verify payment

---

##  WebSocket

- GET /ws/messages – Authenticated WebSocket for real-time messages