# Library Management System
A simple Go application that manages a library system with authors, books, and their relationships. This application demonstrates two ways to run the system: with and without Docker.

---

## **1. Run without Docker**
### **Steps**

### **1. Set Up the Database**
1. Open MySQL Workbench or a terminal.
2. Run the following commands to create the database and set up the tables:
   ```sql
   CREATE DATABASE library;
   USE library;

   -- Import the database schema and sample data
   SOURCE db/setup.sql;
   ```
### **2. Start the Application**
1. Open a terminal or command prompt in the folder where the application is located.
2. Run the following command to start the application:
   ```sql
    go run main.go
   ```
## **2. Run with Docker**
### **Steps**

### **1. Run the Application**
1. Open a terminal or command prompt in the folder where this application is located.
2. Run this command to start the application and the database:
   ```sql
    docker-compose up --build
   ```
### **2.Access the Application**:
   ```sql
    Open your browser and go to: http://localhost:8080
   ```
### **3.Stop the Application**:
   ```sql
    docker-compose down
   ```
## **APIs**
### **Create a new author**
   ```sql
   POST /auth/admin/create-author	
   ```
- Request body
   ```sql
   {
   "name": "Nguyen Van A"
   }
   ```
- Response body
   ```sql
   {
  "message": "Author created successfully",
  "id": 1
   }
   ```

### **Get a list of authors**
   ```sql
   GET /auth/authors	
   ```
- Response body
   ```sql 
  {
    "id": 1,
    "name": "Nguyen Van A"
  },
  {
    "id": 2,
    "name": "Nguyen Van B"
  }
  ```

### **Get author details by ID**
   ```sql
    GET /auth/admin/author?id=1
   ```
- Response body
   ```sql 
  {
  "id": 1,
  "name": "Nguyen Van A"
   }
  ```
### **Create a new book**
   ```sql
   POST /auth/admin/create-book	
   ```
- Request body
   ```sql
   {
  "name": "Harry Potter"
   }
   ```
- Response body
   ```sql
   {
  "message": "Book created successfully",
  "id": 1
   }
   ```

### **Get a list of books**
   ```sql
   GET /auth/books	
   ```
- Response body
   ```sql
  {
    "id": 1,
    "name": "Harry Potter"
  },
  {
    "id": 2,
    "name": "Runeterra"
  }
  ```

### **Get book details by ID**
   ```sql
    GET /auth/admin/book?id=1
   ```
- Response body
   ```sql
  {
  "id": 1,
  "name": "Harry Potter"
   }
  ```

   
   
   