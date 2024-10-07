# Project Title

A brief description of your project, its purpose, and what it accomplishes.

## Table of Contents

- [Technologies Used](#technologies-used)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Usage](#usage)
- [Features](#features)
- [Challenges Faced](#challenges-faced)
- [Testing](#testing)
- [Contributing](#contributing)
- [License](#license)
- [Contact](#contact)

## Technologies Used

- **Backend:**
  - [Go](https://golang.org/) with [Echo Framework](https://echo.labstack.com/)
  - [GORM](https://gorm.io/) for ORM
  - [Asynq](https://github.com/hibiken/asynq) for background task processing
  - [Redis](https://redis.io/) as a broker for Asynq
  - [PostgreSQL](https://www.postgresql.org/) for relational data management

- **Frontend:**
  - HTMX for seamless AJAX interactions
  - HTML, CSS, and JavaScript for the user interface

- **DevOps:**
  - [Docker](https://www.docker.com/) for containerization
  - [Docker Compose](https://docs.docker.com/compose/) for orchestrating multi-container applications
  - [Make](https://www.gnu.org/software/make/) for automation

- **Others:**
  - [Logrus](https://github.com/sirupsen/logrus) for structured logging
  - [go-playground/validator](https://github.com/go-playground/validator) for input validation

## Prerequisites

Before you begin, ensure you have met the following requirements:

- **Docker:** Install Docker from [here](https://docs.docker.com/get-docker/).
- **Docker Compose:** Install Docker Compose from [here](https://docs.docker.com/compose/install/).
- **Make:** Ensure `make` is installed on your system. You can install it from [here](https://www.gnu.org/software/make/).

## Installation

1. **Clone the Repository**

   ```bash
   git clone git@github.com:tsw025/web-analytics.git
   cd web-analytics 
   ```

3. **Build and Start Containers**

   Use the `Makefile` to build and start your Docker containers.

   ```bash
   make app
   ```

   This command will:

    - Build Docker images for your backend and frontend services.
    - Start the services using Docker Compose.
    - Ensure that all dependencies like PostgreSQL and Redis are up and running.

## Usage

Once the application is running, you can access the different components as follows:

- **Frontend:** [http://localhost:8001](http://localhost:8001)
- **Backend API:** [http://localhost:8000](http://localhost:8000)

### API Documentation

- **Docs:** [docs/api.md](docs/api.md)
- **API SPEC:** [docs/api_spec.yaml](docs/api_spec.yaml)


## Features

- **User Authentication:**
    - Registration and login functionalities.
    - Secure password handling and JWT-based authentication.

- **Website Analysis:**
    - Validate and analyze provided URLs.
    - Extract detailed information such as HTML version, page title, headings count, links analysis, and login form detection.
    - Asynchronous processing of analysis tasks to ensure a responsive user experience.

- **Background Task Processing:**
    - Utilizes Asynq with Redis for reliable and scalable background job processing.
    - Dedicated worker service to handle intensive analysis tasks without blocking the main application.

- **Error Handling:**
    - Comprehensive error messages for inaccessible or malformed URLs.
    - Graceful handling of HTTP status codes and useful error descriptions.

- **Logging:**
    - Structured logging with Logrus for better monitoring and debugging.

## Challenges Faced

Developing this project involved tackling several complex areas:

1. **Asynchronous Task Processing:**
    - Ensuring reliable background task execution with Asynq and Redis.
    - Managing task retries and failure scenarios.

2. **HTML Parsing and Analysis:**
    - Accurately extracting and interpreting various elements from diverse web pages.
    - Handling different HTML versions and structures.

3. **Concurrency Control:**
    - Implementing concurrent link accessibility checks without overwhelming system resources.
    - Balancing performance with network usage.

4. **Error Handling:**
    - Providing meaningful and user-friendly error messages.
    - Differentiating between various error types to respond appropriately.

## Testing

While the core functionalities are robust and well-integrated, the project currently lacks comprehensive automated tests. 
Future improvements include:

- **Unit Tests:**
    - Testing individual components like services, repositories, and handlers.

- **Integration Tests:**
    - Ensuring seamless interaction between different parts of the application.

- **End-to-End Tests:**
    - Validating the entire workflow from user input to task processing and result presentation.

Implementing these tests will enhance the reliability and maintainability of the project.
