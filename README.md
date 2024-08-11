# Go Core

### A robust Go core to kickstart your application development

## Features
- **Clean Code**: Adheres to DDD, SOLID principles, domain-driven design, and best practices for maintainable and scalable code.
- **MongoDB Base Service and Repository**: Includes a MongoDB repository with essential CRUD operations, pagination, and filtering. You can easily implement your own database repository without breaking the existing implementation.
- **Easy Queries**: Provides convenient filtering of resources using DTOs, supporting unique and multiple value filters, regex expressions, and more.
- **Logs**: Integrated with zero log for consistent and structured logging.
- **Middlewares**: A comprehensive set of essential middlewares, including JWT authentication, standard error handling, rate limiting, and more.
- **Mail Service**: Ready-to-use mail service for sending basic emails, emails with templates, and attachments.
- **Authentication**: JWT-based authentication with built-in templates for welcome emails and password resets.
- **File Service**: Ready-to-use file service with a default S3 integration. Can be extended to support other file storage services.
- **WebSockets**: Basic WebSocket implementation included (disabled by default).
- **Configurable**: A Gin HTTP server wrapper, allowing you to enable or disable any middleware, modify HTTP server configurations, manage different environments, and more.
- **Documentation**: Well-commented code and a Postman collection to help you quickly build and extend your app.

## Getting Started

### Prerequisites
- Go 1.19 or higher installed on your machine.
- MongoDB instance running (locally or remotely).
- [Mailtrap](https://mailtrap.io/) for testing your emails (or another SMPT server).
- [Postman](https://www.postman.com/downloads/) for API testing (optional but recommended).

### Installation

1. **Clone the repository:**
   ```sh
   git clone https://github.com/IngDeiver/go-core.git
   cd go-core
   
2. **Install dependencies:**
   ```
   go mod tidy
   
3. **Set up environment variables:**
   Run this command and custom with your environment variables
   ```
   cp .env.example .env.local

5. **Run the project:**
   ```
   go run main.go


6. **Testing your setup:**
   
   You can use the provided Postman collection to test the API endpoints. Import the collection from the docs/go-core.postman_collection.json file.
   
## Credits
This project is maintained by IngDeiver.

## License
This project is licensed under the MIT License 
