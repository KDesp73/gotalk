# gotalk

A general api endpoint to handle comments even on static sites

## Endpoints

Check [api-spec.yml](https://github.com/KDesp73/gotalk/blob/main/api/api-spec.yml) for detailed info

> [!NOTE]
> All endpoints are under /v1

### Open

- GET `/ping`
- POST `/register?name={name}&email={email}`

### Authorization needed

> [!NOTE]
> The following endpoints are under /user

- POST `/comment?userid={userid}&threadid={threadid}&content={content}`

### Admin privilleges needed

> [!NOTE]
> The following endpoints are under /admin

- POST `/sudo?id={userid}`
- POST `/sudo/undo?id={userid}`
- POST `/thread/new`
- DELETE `/thread/delete?threadid={threadid}`
