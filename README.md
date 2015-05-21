# Sessions

Simple server-side sessions for [goji](goji.io).

```go
goji.Use(Sessions.Middleware())
```

## Dependencies
- [github.com/fzzy/radix](https://github.com/fzzy/radix) _If using RedisStore_.


## Usage
In-memory session store:

```go
var secret               = "thisismysecret"
var inMemorySessionStore = sessions.MemoryStore{}
var Sessions             = sessions.NewSessionOptions(secret, &inMemorySessionStore)
```

Using Redis (using fzzy/radix):

```go
var redisSessionStore = sessions.NewRedisStore("tcp", "localhost:6379")
var Sessions          = sessions.NewSessionOptions(secret, redisSessionStore)
```

Use middleware:

```go
goji.Use(Sessions.Middleware())
```

Accessing session variable:

```go
func handler(c web.C, w http.ResponseWriter, r *http.Request) {
    sessionObj := Sessions.GetSessionObject(&c)

    // Regnerate session..
    Sessions.RegenerateSession(&c)

    // Delete session
    Sessions.DeleteSession(&c)
}
```

See examples folder for full example.
