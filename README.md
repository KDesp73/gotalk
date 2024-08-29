# Gotalk API

## Overview

The Gotalk API provides a general endpoint for handling comments, making it suitable for both dynamic and static websites. This API allows users to register, post comments, manage threads, and administer user privileges.

## Version

- **API Version**: 0.0.1

## Base URL

The base URL for the API is `/v1`.

## Subroutes

Endpoints that require registration are under the `/auth` path

Endpoints that require admin privileges are under the `/admin` path

## Endpoints

### 1. Ping the Server

- **GET** `/ping`
  - **Summary**: Ping the server to check its status.
  - **Responses**:
    - `200`: Returns "pong".

### 2. User Management

#### Register a New User

- **POST** `/user/new`
  - **Summary**: Register a new user.
  - **Parameters**:
    - `name` (query, required): The name of the user.
    - `email` (query, required): The email of the user.
  - **Responses**:
    - `201`: Registration complete.
    - `400`: `<field> not set`.
    - `409`: `<field> already exists`.

### 3. Comment Management

#### Post a Comment

- **POST** `/auth/users/{userid}/comment`
  - **Summary**: Post a comment to a specific thread.
  - **Parameters**:
    - `userid` (path, required): The ID of the user.
    - `threadid` (query, required): The ID of the thread.
  - **Responses**:
    - `201`: Comment posted successfully.
    - `401`: Unauthorized.
    - `400`: Bad request.

#### Delete a Comment

- **DELETE** `/auth/comments/{commentid}`
  - **Summary**: Delete a specific comment.
  - **Parameters**:
    - `commentid` (path, required): The ID of the comment.
    - `threadid` (query, required): The ID of the thread.
  - **Responses**:
    - `204`: Comment deleted successfully.
    - `404`: Comment not found.
    - `401`: Unauthorized.
    - `400`: Bad request.

### 4. Admin Management

#### Get List of Users

- **GET** `/admin/users`
  - **Summary**: Retrieve the list of users.
  - **Responses**:
    - `200`: Users retrieved successfully.
    - `401`: Unauthorized.

#### Grant Admin Privileges

- **POST** `/admin/users/{userid}/sudo`
  - **Summary**: Grant admin privileges to a user.
  - **Parameters**:
    - `userid` (path, required): The ID of the user.
  - **Responses**:
    - `200`: Admin privileges granted.
    - `401`: Unauthorized.
    - `404`: User not found.
    - `500`: Failed to grant admin privileges.

#### Revoke Admin Privileges

- **POST** `/admin/users/{userid}/sudo/revoke`
  - **Summary**: Revoke admin privileges from a user.
  - **Parameters**:
    - `userid` (path, required): The ID of the user.
  - **Responses**:
    - `200`: Admin privileges revoked.
    - `401`: Unauthorized.
    - `404`: User not found.
    - `500`: Failed to revoke admin privileges.

### 5. Thread Management

#### Retrieve All Thread IDs

- **GET** `/admin/threads`
  - **Summary**: Retrieve all thread IDs.
  - **Responses**:
    - `200`: Threads retrieved successfully.
    - `401`: Unauthorized.

#### Create a New Thread

- **POST** `/admin/threads/new`
  - **Summary**: Create a new thread.
  - **Parameters**:
    - `title` (query, required): The title of the thread.
  - **Responses**:
    - `201`: Thread created successfully.
    - `400`: Thread title not set.
    - `409`: Thread title already exists.
    - `401`: Unauthorized.

#### Delete a Thread

- **DELETE** `/admin/threads/{threadid}`
  - **Summary**: Delete a specific thread.
  - **Parameters**:
    - `threadid` (path, required): The ID of the thread.
  - **Responses**:
    - `204`: Thread deleted successfully.
    - `400`: Unable to parse request.
    - `401`: Unauthorized.
    - `404`: Thread not found.

