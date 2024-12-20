# AchieveIt Backend

## Overview

AchieveIt Backend is a goal tracker API that allows users to create and track their goals, break them down into milestones, and add comments to milestones.

## Features

- User authentication and management
- Goal creation, retrieval, updating, and deletion
- Milestone creation, retrieval, updating, and deletion
- Comment creation, retrieval, updating, and deletion
- Progress tracking for goals based on milestone completion

## Technologies Used

- Go (Golang)
- GORM (ORM for Go)
- Gin (HTTP web framework for Go)
- PostgreSQL (Database)
- UUID (Universally Unique Identifier)

## Getting Started

### Prerequisites

- Go 1.16 or later
- PostgreSQL

### Installation

1. Clone the repository:

```sh
git clone https://github.com/mcbryan1/achieveit-backend.git
cd achieveit-backend
```


2. Install dependencies:

```sh
go mod tidy
```


3. Set up environment variables:

Create a .env file in the root directory and add the following variables:
PORT=8080
DB_URL="host=localhost user=postgres password=your_password dbname=achieveit port=5432"
JWT_SECRET="your_jwt_secret"


4. Install CompileDaemon:

```sh
go install github.com/githubnemo/CompileDaemon@latest
```


5. Run the server with live reloading:

```sh
CompileDaemon -command="./achieveit-backend"
```
The server will start on http://localhost:8080 with live reloading enabled.




# API Endpoints

## User Endpoints

- ``` POST /v1/auth/register:``` Register a new user
- ``` POST /v1/auth/login:``` Authenticate a user


## Goal Endpoints

- ``` POST /v1/goals/create-goal:``` Create a new goal
- ``` GET /v1/goals/fetch-goals:``` Retrieve all goals for the authenticated user
- ``` GET /v1/goals/fetch-goal/:id:``` Retrieve a specific goal
- ``` PUT /v1/goals/update-goal/:id:``` Update a goal
- ``` DELETE /v1/goals/delete-goal/:id:``` Delete a goal


## Milestone Endpoints

- ``` POST /v1/milestones/create-milestone:``` Create a new milestone
- ``` GET /v1/milestones/fetch-milestones:``` Retrieve all milestones for a specific goal (requires ```goal_id``` as query parameter)
- ``` PUT /v1/milestones/update-milestone/:id:``` Update a milestone
- ``` DELETE /v1/milestones/delete-milestone/:id:``` Delete a milestone


## Comments Endpoints

- ``` POST /v1/comments/create-comment:``` Create a new comment
- ``` PUT /v1/comments/update-comment/:id:``` Update a comment
- ``` DELETE /v1/comments/delete-comment/:id:``` Delete a comment


# Example Requests

## Create a Goal

```sh
curl -X POST http://localhost:8080/v1/goals/create-goal -d '{
    "title": "Get a girlfriend",
    "description": "I need a girlfriend for this christmas family gathering",
}' -H "Authorization: Bearer <your_token>"
```

## Fetch a Goal

```sh
curl -X GET http://localhost:8080/v1/goals/fetch-goal/f857d466-0633-4932-a139-4de7d53898ff -H "Authorization: Bearer <your_token>"
```


## Fetch Goals

```sh
curl -X GET http://localhost:8080/v1/goals/fetch-goals -H "Authorization: Bearer <your_token>"
```

## Update Goal

```sh
curl -X PUT http://localhost:8080/v1/goals/update-goal/1 -d '{
    "title": "Get a girlfriend",
    "description": "I need a girlfriend for this christmas family gathering at Tokyo",
}' -H "Authorization: Bearer <your_token>"
```


## Delete Goal

```sh
curl -X DELETE http://localhost:8080/v1/goals/delete-milestone/1 -H "Authorization: Bearer <your_token>"
```




## Create a Milestone

```sh
curl -X POST http://localhost:8080/v1/milestones/create-milestone -d '{
    "title": "Visit the park",
    "goal_id": "f857d466-0633-4932-a139-4de7d53898ff",
    "completed": false
}' -H "Authorization: Bearer <your_token>"
```


## Fetch Milestones

```sh
curl -X GET http://localhost:8080/v1/milestones/fetch-milestones?goal_id=f857d466-0633-4932-a139-4de7d53898ff -H "Authorization: Bearer <your_token>"
```


## Update a Milestone

```sh
curl -X PUT http://localhost:8080/v1/milestones/update-milestone/1 -d '{
    "completed": true
}' -H "Authorization: Bearer <your_token>"
```

## Delete a Milestone

```sh
curl -X DELETE http://localhost:8080/v1/milestones/delete-milestone/1 -H "Authorization: Bearer <your_token>"
```




## Create a Comment

```sh
curl -X POST http://localhost:8080/v1/comments/create-comment -d '{
    "content": "This is going to be great again",
    "milestone_id": "9cfc4506-bef9-4484-a05b-d3405d53489f"
}' -H "Authorization: Bearer <your_token>"
```


## Update a Comment

```sh
curl -X PUT http://localhost:8080/v1/comments/update-comment/1 -d '{
    "content": "Go again"
}' -H "Authorization: Bearer <your_token>"
```

## Delete a Comment

```sh
curl -X DELETE http://localhost:8080/v1/comments/delete-comment/1 -H "Authorization: Bearer <your_token>"
```




## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any changes.



## Acknowledgements

- [Gin](https://github.com/gin-gonic/gin)
- [GORM](https://github.com/go-gorm/gorm)
- [UUID](https://github.com/google/uuid)
- [CompileDaemon](https://github.com/githubnemo/CompileDaemon)




