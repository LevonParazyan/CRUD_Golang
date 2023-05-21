# # Golang CRUD API with Gin and GORM
This project is a basic CRUD (Create, Read, Update, Delete) API built with Golang using the Gin framework and GORM ORM. It allows you to perform operations on a MySQL database.
## Table of Contents
- [Installation](#installation)
- [Usage](#usage)
- [Endpoints](#endpoints)
  - [Create a User](#create-a-user)
  - [Get All Users](#get-all-users)
  - [Get a User by ID](#get-a-user-by-id)
  - [Update a User](#update-a-user)
  - [Delete a User](#delete-a-user)
- [Contributing](#contributing)
- [License](#license)
## Installation
To run the Golang CRUD API locally, follow these steps:
1. Make sure you have Golang and Mysql installed on your machine.
2. Clone the repository to your local machine.
3. Navigate to the project directory.
4. Download the necessary dependencies using the `go mod download` command.
5. Set up a MySQL database. Make sure to create a .env file and pass your credentials in it according to .env.example.
6. Run the application using the command `go run main.go`.
## Usage
Once the API is running, you can make HTTP requests to interact with the records in the MySQL database. Use tools like Postman to send requests to the available endpoints.
## Endpoints
### Create a User
- **Endpoint:** `POST /users`
- **Description:** Create a new user.
- **Request Body:** Object representing the user.
- **Response:** Object representing the created user.
### Get All Users
- **Endpoint:** `GET /users`
- **Description:** Retrieve all users.
- **Response:** Array containing all users.
### Get a User by ID
- **Endpoint:** `GET /users/:id`
- **Description:** Retrieve a specific user by its ID.
- **Parameters:**
  - `id`: The unique identifier of the user.
- **Response:** Object representing the user.
### Update a User
- **Endpoint:** `PATCH /users/:id`
- **Description:** Update a specific user by its ID.
- **Parameters:**
  - `id`: The unique identifier of the user.
- **Request Body:** Object representing the updated user.
- **Response:** Object representing the updated user.
### Delete a User
- **Endpoint:** `DELETE /users/:id`
- **Description:** Delete a specific user by its ID.
- **Parameters:**
  - `id`: The unique identifier of the user.
- **Response:** No response body.
## Contributing
Contributions to the Golang CRUD API are welcome! If you have any improvements or new features to add, feel free to submit a pull request. Please follow the existing code style and conventions.
## License
This project is licensed under the [MIT License](LICENSE).
