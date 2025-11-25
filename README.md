# The Orc Shack REST API

Welcome to **The Orc Shack**, a cozy Middle-Earth-themed restaurant serving traditional dishes from across the lands. This project implements a **REST API** for managing the restaurant's dishes, customers, and user interactions.

---

## Table of Contents

- [Objective](#objective)  
- [Features](#features)  
- [Tasks / Specifications](#tasks--specifications)  
- [Tech Stack](#tech-stack)  
- [Getting Started](#getting-started)  
- [API Endpoints](#api-endpoints)  
- [Authentication](#authentication)  
- [Advanced Features](#advanced-features)  
- [Constraints](#constraints)  
- [Evaluation Criteria](#evaluation-criteria)

---

## Objective

Build a REST API for **The Orc Shack** to allow creation, management, and interaction with dishes served at the restaurant. The API will be used to power a front-end website for customers and restaurant staff.

---

## Features

- **Dish Management:** Create, view, update, delete, and list dishes.  
- **Customer Interaction:** Search and rate dishes.  
- **User Management:** Register, login, and authenticate users (Intermediate & Senior levels).  
- **Security Features:** Rate-limiting, brute-force protection, and input validation.  
- **Performance Optimization:** Multi-tenant support and scalable solutions (Senior level).  
- **Optional Enhancements:** AI/ML sentiment analysis for reviews, OAuth2 SSO, advanced rate-limiting.

---

## Tasks / Specifications

### Task 1 – Junior Level
- Implement CRUD operations for dishes:
  - Each dish must include: `name`, `description`, `price`, `image`.  
- Customers can:
  - Search dishes by name or description.  
  - View and rate dishes.  
- **No authentication required** for this level.

### Task 2 – Intermediate Level
- Add **user registration and login**.  
- Protect all endpoints (except registration) with authentication.  
- Users must have:
  - `name`, `email`, and `password`.  
- Add **data validation** for all entities.  
- Implement **brute-force protection** to prevent malicious login attempts.

### Task 3 – Senior Level
- Optimize API performance for low-resource servers.  
- Add **multi-tenant support** for multiple restaurants.  

### Task 4 – Above & Beyond
- Implement **sentiment analysis** for reviews.  
- Add **rate-limiting** per customer.  
- Support **OAuth2 SSO** (e.g., Google).

---

## Tech Stack

- **Primary Language:** Golang   
- **Database:** SQL-based (PostgreSQL, MySQL, etc.)  
- **Frameworks / Libraries:**  
  - Gin
  - ORM for database interactions  
  - JWT / OAuth2 for authentication  

---

## Getting Started

1. **Clone the repository:**

```bash
git clone https://github.com/your-username/orc-shack.git
cd orc-shack

2. **Install dependencies:**

go mod tidy

3. **Set environment variables:**

export JWT_SECRET="your-secret-key"
export DATABASE_URL="your-database-connection-string"

4. **Run the API:**
run .\cmd\api\main.go


