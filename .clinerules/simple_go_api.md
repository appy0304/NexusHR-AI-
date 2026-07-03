Act as a Principal Software Architect, Staff Backend Engineer, Senior Frontend Engineer, Cloud Architect, DevOps Engineer, and AI Systems Architect with 20+ years of experience building enterprise SaaS platforms.

I already have a working project built with:

Backend:
- Golang
- Gin Framework
- MongoDB
- Layered Architecture
  - Routes
  - Controllers
  - Services
  - DAO
- Swagger Documentation
- CRUD APIs

Frontend:
- Vue 3
- Axios
- Dashboard UI

Current Features:
- Create User
- Get Users
- Update User
- Delete User
- MongoDB Integration
- Swagger UI
- ObjectID
- Timestamps
- API Layer Separation

I want to transform this project into a Production-Grade AI-Powered Employee Management Platform capable of serving 10,000+ employees today and scalable to 100,000+ employees in the future.

The architecture must follow enterprise-grade software engineering principles, clean architecture, scalability, maintainability, observability, security, and AI readiness.

The system should be modular and future-proof so new modules can be added without major refactoring.

Generate a complete project blueprint covering all phases below.

-----------------------------------------------------
PHASE 1
ENTERPRISE EMPLOYEE MANAGEMENT FOUNDATION
-----------------------------------------------------

Convert User into Employee.

Design employee schema with:

- Employee ID
- First Name
- Last Name
- Email
- Phone
- Department
- Designation
- Manager
- Salary
- Joining Date
- Employment Status
- Skills
- Address
- Emergency Contact
- Created At
- Updated At

Explain:

- Why each field exists
- Real-world business use case
- Database design decisions
- Validation rules

Implement:

- Create Employee
- Get Employee
- Get All Employees
- Update Employee
- Delete Employee

Add:

- Pagination
- Sorting
- Filtering
- Search

Explain:
- Why pagination is required
- Why filtering is required
- How large enterprises handle employee records

-----------------------------------------------------
PHASE 2
PRODUCTION API ARCHITECTURE
-----------------------------------------------------

Design APIs using:

- REST Best Practices
- Standard Response Structure
- Error Handling Framework
- Request Validation
- DTO Pattern
- Middleware Pattern

Implement:

- Request IDs
- Correlation IDs
- Structured Logging

Explain:

- API Lifecycle
- Request Flow
- Controller Flow
- Service Flow
- DAO Flow
- Database Flow

Create detailed API architecture diagrams.

-----------------------------------------------------
PHASE 3
AUTHENTICATION & AUTHORIZATION
-----------------------------------------------------

Implement:

JWT Authentication

Roles:

- Super Admin
- HR Admin
- Manager
- Employee

Explain:

- Authentication
- Authorization
- Access Control

Implement:

- Login
- Logout
- Refresh Token
- Password Hashing
- Password Reset

Use:

- bcrypt
- JWT

Create middleware.

Explain every security decision.

-----------------------------------------------------
PHASE 4
ADVANCED EMPLOYEE MODULES
-----------------------------------------------------

Build:

Employee Lifecycle Management

Modules:

- Employee Profile
- Promotion History
- Department Transfers
- Attendance
- Leave Management
- Performance Reviews
- Training Records

Explain:

- HR workflow
- Business logic
- Database design

-----------------------------------------------------
PHASE 5
ANALYTICS & REPORTING
-----------------------------------------------------

Create dashboard metrics:

- Total Employees
- Department Distribution
- New Joinees
- Attrition Rate
- Leave Statistics
- Attendance Trends
- Salary Distribution

Use MongoDB Aggregation Pipelines.

Explain:

- Aggregation Framework
- Dashboard Design
- KPI Calculations

Build charts and visualizations.

-----------------------------------------------------
PHASE 6
AUDIT TRAIL & ACTIVITY LOGGING
-----------------------------------------------------

Track every action:

- Create
- Update
- Delete
- Login
- Logout

Store:

- User
- Timestamp
- Action
- IP Address
- Device

Explain:

- Compliance
- Security
- Enterprise Requirements

-----------------------------------------------------
PHASE 7
DOCUMENT MANAGEMENT SYSTEM
-----------------------------------------------------

Allow uploads:

- Resume
- Offer Letter
- Certificates
- Identity Documents

Use:

- MongoDB Atlas
- AWS S3

Implement:

- Upload
- Download
- Versioning

Explain:

- Storage Strategy
- Security
- Scalability

-----------------------------------------------------
PHASE 8
OBSERVABILITY & PRODUCTION MONITORING
-----------------------------------------------------

Implement:

Logging
Metrics
Tracing

Use:

- OpenTelemetry
- Prometheus
- Grafana

Track:

- API Latency
- Error Rates
- Database Performance
- Request Tracing

Explain:

- Distributed Tracing
- Monitoring
- Incident Investigation

Provide production debugging workflow.

-----------------------------------------------------
PHASE 9
AI ASSISTANT INTEGRATION
-----------------------------------------------------

Build:

AI HR Assistant

Frontend:

Chat Interface

Backend:

Go APIs

LLM Integration:

- OpenAI
- Claude
- Gemini

Allow queries:

"Who joined this month?"

"Show employees from IT department."

"Who has not completed training?"

Explain:

- Prompt Engineering
- Tool Calling
- Function Calling
- AI Architecture

Create full architecture diagram.

-----------------------------------------------------
PHASE 10
RAG KNOWLEDGE SYSTEM
-----------------------------------------------------

Create:

Enterprise Knowledge Base

Store:

- HR Policies
- Company Handbook
- Training Material
- Leave Rules

Implement:

- Embeddings
- Chunking
- Vector Search
- Semantic Search

Use:

- OpenAI Embeddings
- Pinecone
- Weaviate
- Chroma

Explain:

- RAG
- Retrieval Pipeline
- Embedding Flow
- Query Flow

Allow employees to ask:

"What is the maternity leave policy?"

"How many annual leaves are allowed?"

-----------------------------------------------------
PHASE 11
AI EMPLOYEE INTELLIGENCE
-----------------------------------------------------

Build AI features:

- Employee Performance Summary
- Skill Gap Detection
- Learning Recommendations
- Attrition Risk Analysis
- Promotion Readiness

Explain:

- ML vs LLM
- Predictive Analytics
- Business Value

-----------------------------------------------------
PHASE 12
ENTERPRISE FRONTEND
-----------------------------------------------------

Build a modern SaaS dashboard using:

- Vue 3
- Composition API
- Pinia
- Vue Router

Design:

- Premium UI
- Dark Mode
- Animations
- Responsive Layout
- Dashboard Widgets

Create:

- Employee Directory
- Analytics Dashboard
- AI Assistant
- Reports Section

Follow:

- Enterprise UX
- Accessibility
- Scalability

-----------------------------------------------------
PHASE 13
CLOUD & DEVOPS
-----------------------------------------------------

Containerize:

- Backend
- Frontend

Use:

Docker

Implement:

- Docker Compose
- CI/CD
- GitHub Actions

Deploy:

Frontend:
- Vercel

Backend:
- Render
- Railway
- AWS ECS

Database:
- MongoDB Atlas

Explain:

- Build Pipeline
- Deployment Pipeline
- Rollback Strategy

-----------------------------------------------------
PHASE 14
SCALABILITY FOR 100,000+ EMPLOYEES
-----------------------------------------------------

Design:

- Horizontal Scaling
- Load Balancing
- Caching
- Redis
- Queue Systems

Implement:

- Redis
- RabbitMQ / Kafka

Explain:

- Performance Bottlenecks
- Database Optimization
- Query Optimization
- Indexing Strategy

Create capacity planning documentation.

-----------------------------------------------------
PHASE 15
FUTURE EXPANSION ARCHITECTURE
-----------------------------------------------------

Design the system so future modules can be added easily:

Future Modules:

- Payroll
- Recruitment ATS
- CRM
- Asset Management
- Project Management
- Helpdesk
- LMS
- Employee Engagement
- AI Recruiting Assistant
- AI Resume Screening
- AI Career Coach

Explain:

- Plugin Architecture
- Domain Driven Design
- Microservice Migration Strategy

-----------------------------------------------------
DELIVERABLES
-----------------------------------------------------

Generate:

1. Complete Architecture Diagram
2. Database Design
3. Folder Structure
4. Backend Structure
5. Frontend Structure
6. API Documentation
7. Security Design
8. AI Design
9. DevOps Design
10. Deployment Strategy
11. Scaling Strategy
12. Learning Roadmap
13. Production Readiness Checklist
14. Future Growth Roadmap

For every phase explain:

- Why it exists
- Real-world business problem solved
- Architecture
- Workflow
- Database changes
- API changes
- Frontend changes
- Security implications
- Scalability considerations
- Interview explanation

The final result should be a production-grade AI-Powered Employee Management Platform that demonstrates Full Stack Engineering, Backend Architecture, Database Engineering, DevOps, Cloud, Scalability, and AI Engineering skills.