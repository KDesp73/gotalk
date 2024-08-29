# gotalk

A general API endpoint to handle comments even on static sites.

## Endpoints

Check [api-spec.yml](https://github.com/KDesp73/gotalk/blob/main/api/api-spec.yml) for detailed info.

> [!NOTE]
> All endpoints are under `/v1`.

### Open

- **GET** `/ping`  
  Ping the server.

- **POST** `/user/new`  
  Register a new user.  
  **Request Parameters**:  
  - `name` (query, required): The name of the user.
  - `email` (query, required): The email of the user.  

### Authorization Needed

> [!NOTE]
> The following endpoints are under `/auth`.

- **POST** `/users/{userid}/comment`  
  Post a comment.  
  **Request Parameters**:  
  - `userid` (path, required): The ID of the user.
  - `threadid` (query, required): The ID of the thread.  

### Admin Privileges Needed

> [!NOTE]
> The following endpoints are under `/admin`.

- **POST** `/admin/users/{userid}/sudo`  
  Grant admin privileges to a user.  
  **Request Parameters**:  
  - `userid` (path, required): The ID of the user.

- **POST** `/admin/users/{userid}/sudo/revoke`  
  Revoke admin privileges from a user.  
  **Request Parameters**:  
  - `userid` (path, required): The ID of the user.

- **POST** `/threads/new`  
  Create a new thread.  
  **Request Parameters**:  
  - `title` (query, required): The title of the thread.

- **DELETE** `/threads/{threadid}`  
  Delete a thread.  
  **Request Parameters**:  
  - `threadid` (path, required): The ID of the thread.

