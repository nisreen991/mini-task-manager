# Mini Task Manager

## Table of Contents

- [Setup Instructions](#setup-instructions)
  - [Backend Setup](#backend-setup)
  - [Frontend Setup](#frontend-setup)
- [Tech Stack](#tech-stack)
- [Time Spent and Trade-offs](#time-spent-and-trade-offs)

## Setup Instructions

### Backend Setup

1.  **Navigate to the backend directory:**
    ```bash
    cd backend
    ```
2.  **Ensure Go is installed:** This project requires Go to run the backend. You can download it from [golang.org](https://golang.org/dl/).
3.  **Ensure MongoDB is running:** The backend connects to a MongoDB instance. Make sure you have MongoDB installed and running on `mongodb://localhost:27017`.
4.  **Run the backend:**
    ```bash
    go run main.go
    ```

### Frontend Setup

1.  **Navigate to the frontend directory:**
    ```bash
    cd frontend
    ```
2.  **Install Node.js and npm:** Ensure you have Node.js and npm installed. You can download them from [nodejs.org](https://nodejs.org/).
3.  **Install Angular CLI:** If you don't have it, install the Angular CLI globally:
    ```bash
    npm install -g @angular/cli
    ```
4.  **Install frontend dependencies:**
    ```bash
    npm install
    ```
5.  **Run the frontend application:**
    ```bash
    ng serve --open
    ```
    This will open the application in your browser, usually at `http://localhost:4200/`.

## Tech Stack

This project utilizes the following technologies:

*   **Backend:**
    *   **Go:** Chosen for its simplicity, performance, and strong concurrency features, making it suitable for building efficient APIs.
    *   **MongoDB:** A NoSQL database used for flexible data storage, ideal for handling semi-structured task data.
    *   **Gorilla Mux:** A powerful URL router and dispatcher for Go, used for handling API routes.
    *   **BSON/Primitive:** Go drivers for MongoDB, enabling seamless interaction with the database.
    *   **JWT (JSON Web Tokens):** Used for secure authentication and authorization between the frontend and backend.

*   **Frontend:**
    *   **Angular:** A popular framework for building robust and scalable single-page applications. Chosen for its comprehensive features, component-based architecture, and strong tooling.
    *   **TypeScript:** Provides static typing to JavaScript, enhancing code quality and maintainability in the Angular application.
    *   **SCSS (Sass):** A powerful CSS preprocessor that adds features like variables, nesting, and mixins, making styling more efficient and organized.

## Time Spent and Trade-offs

**Time Spent:** Approximately 5-6 hours were spent on developing this application, including both backend and frontend implementation, as well as debugging and setting up the development environment.

**Trade-offs Made:**

*   **Error Handling Granularity:** While basic error handling is in place, more fine-grained error responses could be implemented for a production-ready application (e.g., specific error codes for different types of validation failures).
*   **Input Validation:** Frontend input validation is present, but more robust backend validation could be added to ensure data integrity.
*   **UI/UX Polish:** The primary focus was on core functionality. Further time could be spent on enhancing the user interface and user experience with more advanced styling, animations, and responsiveness.
*   **Testing:** Comprehensive unit and integration tests were not prioritized due to time constraints. For a production application, a full testing suite would be essential.
*   **Scalability:** The current setup is suitable for a mini-task manager. For a larger-scale application, considerations like load balancing, database sharding, and more advanced authentication mechanisms would be necessary.
*   **Security:** Basic JWT authentication is implemented. However, more advanced security measures like refresh tokens, secure cookie handling, and protection against common web vulnerabilities (e.g., XSS, CSRF) could be added.