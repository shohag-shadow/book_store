# Gin Crash Course: Lesson by Lesson

## Index

- [What is Gin, and why should you care?](#what-is-gin-and-why-should-you-care)
- [What you need before starting](#what-you-need-before-starting)
- [What we are building](#what-we-are-building)
- [Lesson 1: Setup and Hello World](#lesson-1-setup-and-hello-world)
  - [What just happened?](#what-just-happened)
  - [A real life picture](#a-real-life-picture)
- [Lesson 2: Routing and Responses](#lesson-2-routing-and-responses)
  - [HTTP methods](#http-methods)
  - [Ways to respond](#ways-to-respond)
  - [Prefer structs over gin.H in real projects](#prefer-structs-over-ginh-in-real-projects)
  - [Use status code constants](#use-status-code-constants)
  - [Custom 404 and 405](#custom-404-and-405)
- [Lesson 3: Path and Query Parameters](#lesson-3-path-and-query-parameters)
  - [Path parameters](#path-parameters)
  - [Wildcard parameters](#wildcard-parameters)
  - [Query parameters](#query-parameters)
  - [When to use which?](#when-to-use-which)
- [Lesson 4: Request Bodies and Binding](#lesson-4-request-bodies-and-binding)
  - [The binding family](#the-binding-family)
  - [Form data](#form-data)
  - [ShouldBind vs Bind](#shouldbind-vs-bind)
  - [Two gotchas](#two-gotchas)
- [Lesson 5: Validation](#lesson-5-validation)
  - [The required gotcha](#the-required-gotcha)
  - [Friendly error messages](#friendly-error-messages)
  - [Use JSON names in errors and add a custom rule](#use-json-names-in-errors-and-add-a-custom-rule)
- [Lesson 6: Route Groups](#lesson-6-route-groups)
  - [Why versioning in the path?](#why-versioning-in-the-path)
  - [Groups carry middleware](#groups-carry-middleware)
- [Lesson 7: Middleware](#lesson-7-middleware)
  - [The picture](#the-picture)
  - [Writing your first middleware](#writing-your-first-middleware)
  - [Where to attach it](#where-to-attach-it)
  - [Stopping the chain](#stopping-the-chain)
  - [Passing data down the chain](#passing-data-down-the-chain)
  - [Built in middleware](#built-in-middleware)
  - [One important trap: goroutines](#one-important-trap-goroutines)
- [Lesson 8: Mini Project: In-Memory Book API](#lesson-8-mini-project-in-memory-book-api)
  - [What to notice](#what-to-notice)
- [Lesson 9: Database with GORM](#lesson-9-database-with-gorm)
  - [Models](#models)
  - [Connecting](#connecting)
  - [Handlers using the database](#handlers-using-the-database)
  - [GORM cheat moves](#gorm-cheat-moves)
- [Lesson 10: Project Structure and Config](#lesson-10-project-structure-and-config)
  - [Config from environment variables](#config-from-environment-variables)
  - [The validation package](#the-validation-package)
  - [Routes in one place](#routes-in-one-place)
  - [Why this layout pays off](#why-this-layout-pays-off)
- [Lesson 11: Authentication with JWT](#lesson-11-authentication-with-jwt)
  - [The picture](#the-picture-1)
  - [Token helpers](#token-helpers)
  - [Register and login handlers](#register-and-login-handlers)
  - [The auth middleware](#the-auth-middleware)
  - [Wire it up](#wire-it-up)
  - [Where this simple version falls short](#where-this-simple-version-falls-short)
- [Lesson 12: Error Handling Done Right](#lesson-12-error-handling-done-right)
  - [Step 1: a custom error type](#step-1-a-custom-error-type)
  - [Step 2: one middleware that turns errors into responses](#step-2-one-middleware-that-turns-errors-into-responses)
  - [Step 3: handlers get much shorter](#step-3-handlers-get-much-shorter)
  - [Panics and unknown routes](#panics-and-unknown-routes)
- [Lesson 13: Pagination, Search, and Sorting](#lesson-13-pagination-search-and-sorting)
  - [Why validate sort and order so strictly?](#why-validate-sort-and-order-so-strictly)
  - [Offset pagination vs cursor pagination](#offset-pagination-vs-cursor-pagination)
- [Lesson 14: File Upload and Static Files](#lesson-14-file-upload-and-static-files)
  - [The upload handler](#the-upload-handler)
  - [Serving files](#serving-files)
  - [Sending a file as a download](#sending-a-file-as-a-download)
  - [Safety checklist for uploads](#safety-checklist-for-uploads)
- [Lesson 15: CORS, Rate Limiting, Compression, Logging](#lesson-15-cors-rate-limiting-compression-logging)
  - [CORS](#cors)
  - [Rate limiting](#rate-limiting)
  - [Trusted proxies](#trusted-proxies)
  - [Compression](#compression)
  - [Request ID and structured logging](#request-id-and-structured-logging)
  - [Putting it together](#putting-it-together)
- [Lesson 16: Testing](#lesson-16-testing)
  - [Testing tips](#testing-tips)
- [Lesson 17: Graceful Shutdown and Server Settings](#lesson-17-graceful-shutdown-and-server-settings)
  - [Why the timeouts matter](#why-the-timeouts-matter)
- [Lesson 18: Bonus: HTML Templates, SSE, Swagger](#lesson-18-bonus-html-templates-sse-swagger)
  - [HTML templates](#html-templates)
  - [Server Sent Events (SSE)](#server-sent-events-sse)
  - [Swagger documentation](#swagger-documentation)
- [Lesson 19: Docker and Deployment](#lesson-19-docker-and-deployment)
  - [Dockerfile](#dockerfile)
  - [Production checklist](#production-checklist)
- [Lesson 20: Cheat Sheet, Common Mistakes, Next Steps](#lesson-20-cheat-sheet-common-mistakes-next-steps)
  - [Cheat sheet](#cheat-sheet)
  - [Common mistakes](#common-mistakes)
  - [Practice projects](#practice-projects)
  - [Where to go next](#where-to-go-next)

Welcome to this Gin course. By the end, you will go from `go mod init` to a tested, containerized REST API with authentication, validation, a database, file uploads, and graceful shutdown.

One rule for this course: do not just read. Open your editor and type every example yourself. Break things on purpose, read the error, fix it. That is how it sticks. Each lesson ends with a short recap and a homework task. Do the homework. Seriously.

## What is Gin, and why should you care?

Gin is a web framework for Go. It sits on top of the standard `net/http` package and adds the things you would otherwise write yourself again and again:

- A very fast router with path parameters
- Middleware (code that runs before or after your handlers)
- Request binding and validation (JSON to struct, with rules)
- Easy JSON responses
- Panic recovery and request logging out of the box

Think of `net/http` as a bare kitchen with a stove. Gin is the same kitchen with the counters, knives, and labeled shelves already set up. You can still cook everything the hard way, but you will not want to.

## What you need before starting

- Go installed (a recent version, 1.22 or newer is a safe choice)
- Basic Go: structs, functions, errors, packages, and modules
- Basic HTTP: methods (GET, POST), status codes, headers, JSON
- A tool to send requests: `curl`, Postman, or Thunder Client in VS Code

## What we are building

A **Bookstore API** that grows lesson by lesson:

- CRUD for books
- Validation and clean error responses
- SQLite database using GORM
- User registration, login, and JWT authentication
- Pagination, search, and sorting
- Cover image upload
- CORS, rate limiting, logging
- Tests, graceful shutdown, Docker

***

## Lesson 1: Setup and Hello World

**You will learn:** how to create a Go module, install Gin, and run your first server.

Open a terminal and run these commands:

```bash
go version
mkdir gin-bookstore
cd gin-bookstore
go mod init github.com/yourname/gin-bookstore
go get github.com/gin-gonic/gin
```

`go mod init` creates `go.mod`, the file that tracks your project and its dependencies. `go get` downloads Gin and records it there.

Now create `main.go`:

```go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	r.Run(":8080")
}
```

Run it:

```bash
go run .
```

In another terminal:

```bash
curl http://localhost:8080/ping
```

You should see `{"message":"pong"}`. Congratulations, you have a web server.

### What just happened?

- `gin.Default()` creates a router (called an engine) with two middlewares already attached: **Logger** (prints every request) and **Recovery** (catches panics so one bad request does not crash your server). If you want a bare engine with nothing attached, use `gin.New()`.
- `r.GET("/ping", handler)` says: when a GET request arrives at `/ping`, run this function.
- `*gin.Context` is the star of the show. It holds the incoming request, the outgoing response, URL parameters, and helper methods. Every handler receives one.
- `gin.H` is just a shortcut for `map[string]any`. It is handy for quick JSON.
- `r.Run(":8080")` starts the server on port 8080.

### A real life picture

Imagine a restaurant. The **router** is the host at the door who looks at your reservation and sends you to the right section. The **handler** is the waiter who actually serves you. The **context** is the order slip that travels with you from the door to the kitchen and back. Keep this picture in mind, we will extend it later with middleware.

You will also notice Gin prints a warning about debug mode. That is normal. We will switch to release mode in the deployment lesson.

### Recap

- `go mod init` and `go get` set up the project
- `gin.Default()` gives you a router with Logger and Recovery
- Handlers take a `*gin.Context`

### Homework

Add a `/hello` route that returns `{"message": "hello from gin"}` and a `/time` route that returns the current server time.

***

## Lesson 2: Routing and Responses

**You will learn:** HTTP methods in Gin, the different ways to respond, and custom 404 pages.

### HTTP methods

Gin has one method per HTTP verb:

```go
r.GET("/books", listBooks)
r.POST("/books", createBook)
r.PUT("/books/:id", replaceBook)
r.PATCH("/books/:id", updateBook)
r.DELETE("/books/:id", deleteBook)
```

What each verb normally means in a REST API:

- **GET** reads data, changes nothing
- **POST** creates something new
- **PUT** replaces a whole resource
- **PATCH** changes part of a resource
- **DELETE** removes a resource

There is also `r.Any("/path", handler)` which matches all methods, and `r.Handle("PURGE", "/cache", handler)` for custom ones.

### Ways to respond

```go
// JSON (you will use this 95 percent of the time)
c.JSON(http.StatusOK, gin.H{"ok": true})

// Pretty printed JSON, nice while debugging
c.IndentedJSON(http.StatusOK, gin.H{"ok": true})

// Plain text with formatting
c.String(http.StatusOK, "Hello %s", "Gin")

// Status only, no body (great for DELETE)
c.Status(http.StatusNoContent)

// Redirect
c.Redirect(http.StatusMovedPermanently, "https://gin-gonic.com")

// Raw bytes with a content type you choose
c.Data(http.StatusOK, "text/plain; charset=utf-8", []byte("raw bytes"))

// Send a file from disk
c.File("./report.pdf")
```

Gin can also render XML and YAML with `c.XML` and `c.YAML`.

### Prefer structs over gin.H in real projects

`gin.H` is quick, but structs give you a stable response shape that you can document and test:

```go
type PingResponse struct {
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
}

r.GET("/ping", func(c *gin.Context) {
	c.JSON(http.StatusOK, PingResponse{Message: "pong", Time: time.Now()})
})
```

### Use status code constants

Write `http.StatusCreated` instead of `201`. It reads better and avoids typos. The ones you will use most:

- `200 OK`, `201 Created`, `204 No Content`
- `400 Bad Request`, `401 Unauthorized`, `403 Forbidden`, `404 Not Found`, `409 Conflict`, `422 Unprocessable Entity`, `429 Too Many Requests`
- `500 Internal Server Error`

### Custom 404 and 405

```go
r.NoRoute(func(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"error": "route not found"})
})

r.HandleMethodNotAllowed = true
r.NoMethod(func(c *gin.Context) {
	c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "method not allowed"})
})
```

### Recap

- One router method per HTTP verb
- `c.JSON` is your main response tool
- Use constants for status codes, structs for stable responses

### Homework

Create routes for all five verbs on `/notes` that simply return a JSON message saying which verb was used. Then hit each one with curl using `-X`.

***

## Lesson 3: Path and Query Parameters

**You will learn:** how to read data from the URL.

### Path parameters

A path parameter is part of the URL itself, marked with a colon:

```go
r.GET("/users/:id", func(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"user_id": id})
})
```

Request `GET /users/42` gives you `id = "42"`. Note that it is always a string. Convert it yourself:

```go
id, err := strconv.Atoi(c.Param("id"))
if err != nil {
	c.JSON(http.StatusBadRequest, gin.H{"error": "id must be a number"})
	return
}
```

Look at that `return`. Forgetting it is the single most common Gin bug. `c.JSON` writes the response but does **not** stop your function. Without `return`, the code below keeps running.

### Wildcard parameters

A `*` captures everything after it:

```go
r.GET("/files/*filepath", func(c *gin.Context) {
	c.String(http.StatusOK, "file path is %s", c.Param("filepath"))
})
```

`GET /files/images/cat.png` gives `/images/cat.png` (with the leading slash).

### Query parameters

Query parameters come after the `?` in the URL, like `/search?q=golang&page=2`:

```go
r.GET("/search", func(c *gin.Context) {
	q := c.Query("q")                  // empty string if missing
	page := c.DefaultQuery("page", "1") // fallback value if missing

	if value, ok := c.GetQuery("sort"); ok {
		// the sort parameter was present, even if empty
		_ = value
	}

	c.JSON(http.StatusOK, gin.H{"q": q, "page": page})
})
```

Arrays and maps also work:

```go
tags := c.QueryArray("tag")   // /search?tag=go&tag=gin
filters := c.QueryMap("f")    // /search?f[color]=red&f[size]=xl
```

### When to use which?

Path parameters identify **which resource** (`/books/7`). Query parameters **adjust the view** (`/books?page=2&sort=price`). A good rule: if removing it changes what the resource is, it belongs in the path. If it only filters or sorts, it belongs in the query.

### Recap

- `c.Param` for path values, `c.Query` and `c.DefaultQuery` for query values
- Everything arrives as a string, so convert and validate
- Always `return` after sending an error response

### Homework

Build `GET /greet/:name` that reads an optional `?lang=` query. Return a greeting in English by default and in two other languages of your choice.

***

## Lesson 4: Request Bodies and Binding

**You will learn:** how to turn incoming JSON or form data into Go structs.

When a client sends a POST with a JSON body, you do not want to parse it by hand. Gin does it with **binding**:

```go
type BookInput struct {
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}

r.POST("/books", func(c *gin.Context) {
	var in BookInput
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, in)
})
```

Test it:

```bash
curl -X POST http://localhost:8080/books \
  -H "Content-Type: application/json" \
  -d '{"title":"Clean Code","author":"Robert Martin","price":29.99}'
```

### The binding family

Each method reads from a different place and uses a different struct tag:

- `ShouldBindJSON` reads the JSON body, tag `json`
- `ShouldBindQuery` reads the query string, tag `form`
- `ShouldBindUri` reads path parameters, tag `uri`
- `ShouldBindHeader` reads headers, tag `header`
- `ShouldBind` guesses from the Content-Type (JSON, XML, form)

Example mixing URI and JSON:

```go
type BookURI struct {
	ID uint `uri:"id" binding:"required"`
}

r.PUT("/books/:id", func(c *gin.Context) {
	var uri BookURI
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var in BookInput
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": uri.ID, "book": in})
})
```

### Form data

For HTML forms or `multipart/form-data`:

```go
type LoginForm struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

r.POST("/login", func(c *gin.Context) {
	var f LoginForm
	if err := c.ShouldBind(&f); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"email": f.Email})
})
```

You can also read single fields with `c.PostForm("email")` and `c.DefaultPostForm("lang", "en")`.

### ShouldBind vs Bind

Gin has both `ShouldBindJSON` and `BindJSON`. The `Bind` versions automatically abort with a 400 and write a header when binding fails. That takes control away from you, so you cannot shape the error response. **Use the `ShouldBind` family.**

### Two gotchas

1. **The request body can only be read once.** If you need to bind the same body into two structs, use `c.ShouldBindBodyWith(&obj, binding.JSON)`, which caches the body.
2. **Optional fields for PATCH.** With a plain `float64`, you cannot tell "client sent 0" from "client did not send it". Use pointers:

```go
type UpdateBookInput struct {
	Title *string  `json:"title"`
	Price *float64 `json:"price"`
}
// nil means "not provided", non-nil means "set it to this value"
```

### Recap

- Define a struct, tag it, call `ShouldBindXxx`
- Pick the binder that matches where the data lives
- Prefer `ShouldBind` over `Bind`

### Homework

Create `POST /signup` that accepts JSON with `name`, `email`, and `age`, and returns the same data back with status 201.

***

## Lesson 5: Validation

**You will learn:** how to reject bad input before it reaches your logic.

Gin uses the `go-playground/validator` library under the hood. You add rules with the `binding` struct tag:

```go
type RegisterInput struct {
	Name            string `json:"name" binding:"required,min=2,max=50"`
	Email           string `json:"email" binding:"required,email"`
	Age             int    `json:"age" binding:"required,gte=13,lte=120"`
	Password        string `json:"password" binding:"required,min=8,max=72"`
	ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=Password"`
	Role            string `json:"role" binding:"required,oneof=reader author"`
}
```

Rules are comma separated and all must pass. The ones you will use all the time:

- `required`: must not be the zero value
- `min`, `max`: length for strings and slices, value for numbers
- `len`: exact length
- `gt`, `gte`, `lt`, `lte`: number comparisons
- `email`, `url`, `uuid`: format checks
- `oneof=a b c`: value must be one of these
- `eqfield=Other`: must equal another field
- `omitempty`: skip the remaining rules if the value is empty
- `dive`: apply the following rules to each item of a slice or map

### The required gotcha

`required` fails on zero values. A price of `0` or an age of `0` is rejected as if it was missing. If zero is a valid value, use a pointer:

```go
type ProductInput struct {
	Stock *int `json:"stock" binding:"required,gte=0"` // 0 is allowed, missing is not
}
```

### Friendly error messages

The default error text is ugly (`Key: 'RegisterInput.Email' Error:Field validation for 'Email' failed on the 'email' tag`). Let us turn it into something a frontend can use:

```go
func validationMessages(err error) map[string]string {
	var ve validator.ValidationErrors
	if !errors.As(err, &ve) {
		return nil
	}

	out := make(map[string]string, len(ve))
	for _, fe := range ve {
		switch fe.Tag() {
		case "required":
			out[fe.Field()] = "this field is required"
		case "email":
			out[fe.Field()] = "must be a valid email address"
		case "min":
			out[fe.Field()] = "must be at least " + fe.Param()
		case "max":
			out[fe.Field()] = "must be at most " + fe.Param()
		case "oneof":
			out[fe.Field()] = "must be one of: " + fe.Param()
		case "eqfield":
			out[fe.Field()] = "must match " + fe.Param()
		default:
			out[fe.Field()] = "is invalid (" + fe.Tag() + ")"
		}
	}
	return out
}
```

Use it in a handler:

```go
if err := c.ShouldBindJSON(&in); err != nil {
	if fields := validationMessages(err); fields != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "validation failed", "fields": fields})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON body"})
	return
}
```

Imports you need: `errors` and `github.com/go-playground/validator/v10`. Run `go mod tidy` after adding them.

### Use JSON names in errors and add a custom rule

By default, `fe.Field()` returns the Go field name (`ConfirmPassword`). Your frontend uses JSON names (`confirm_password`). Fix that, and add a custom validation rule at the same time:

```go
func setupValidator() {
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		return
	}

	// Report JSON field names instead of Go field names
	v.RegisterTagNameFunc(func(f reflect.StructField) string {
		name := strings.SplitN(f.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	// Custom rule: block reserved usernames
	_ = v.RegisterValidation("notreserved", func(fl validator.FieldLevel) bool {
		switch strings.ToLower(fl.Field().String()) {
		case "admin", "root", "support":
			return false
		}
		return true
	})
}
```

Call `setupValidator()` once at startup. Now you can write `binding:"required,notreserved"` on any string field. Imports: `reflect`, `strings`, `github.com/gin-gonic/gin/binding`, and the validator package. In Lesson 10 we will move this into its own package.

### Recap

- Rules live in the `binding` tag
- `required` fails on zero values, use pointers when zero is valid
- Convert validator errors into clean field messages

### Homework

Add a `username` field to `RegisterInput` with `notreserved`, 3 to 20 characters. Test that `admin` is rejected and shows a clear message.

***

## Lesson 6: Route Groups

**You will learn:** how to organize routes and share a prefix or middleware.

Right now every route is registered on `r`. Real APIs have dozens of routes, and many share a prefix like `/api/v1`. Groups fix this:

```go
api := r.Group("/api/v1")
{
	api.GET("/books", listBooks)
	api.GET("/books/:id", getBook)

	admin := api.Group("/admin")
	{
		admin.GET("/stats", stats)
		admin.DELETE("/books/:id", forceDelete)
	}
}
```

The curly braces are just a Go block. They do nothing technically, they only make the grouping visible. The final paths are `/api/v1/books`, `/api/v1/admin/stats`, and so on.

### Why versioning in the path?

When you must change a response shape later, you can add `/api/v2` and keep v1 alive for old clients. Think of it like a hospital with an old and a new wing: patients (clients) keep using the wing they know, while you build the new one.

### Groups carry middleware

You can attach middleware to a whole group, which is the foundation of "public routes vs protected routes":

```go
public := r.Group("/api/v1")
protected := r.Group("/api/v1", authMiddleware) // handlers after the path are middleware
```

We will build `authMiddleware` in the next lesson.

### Recap

- `r.Group(prefix)` returns a group with the same methods as the router
- Groups nest
- Groups share middleware

### Homework

Organize your earlier routes into `/api/v1` with nested groups for `/users` and `/books`.

***

## Lesson 7: Middleware

**You will learn:** the most important idea in Gin, and how to write your own.

### The picture

Think of the security desk at a big office building. Every visitor passes the desk before reaching any office. The guard can:

- Note down who came and when (logging)
- Check your ID (authentication)
- Turn you away (abort)
- Add a visitor badge to your clothes, which the offices can later read (storing data in the context)

Middleware is that desk. It is a function that runs **around** your handler.

```text
request  -> Logger -> Recovery -> Auth -> Handler
response <- Logger <- Recovery <- Auth <- Handler
```

The request goes in through each layer, hits the handler, and the response comes back out through the same layers in reverse.

### Writing your first middleware

A middleware is a function that returns a `gin.HandlerFunc`:

```go
func RequestTimer() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next() // run the rest of the chain (other middleware and the handler)

		log.Printf("%s %s took %v", c.Request.Method, c.Request.URL.Path, time.Since(start))
	}
}
```

Everything before `c.Next()` runs on the way in. Everything after runs on the way out.

### Where to attach it

```go
r.Use(RequestTimer())                // every route
api.Use(RequestTimer())              // every route in this group
r.GET("/slow", RequestTimer(), slow) // just this route
```

Order matters. Middleware runs in the order you add it.

### Stopping the chain

Sometimes the guard says no. Use `Abort`:

```go
func APIKey(expected string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("X-API-Key") != expected {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid API key"})
			return // important
		}
		c.Next()
	}
}
```

`Abort` prevents the **remaining handlers** from running, but your current function keeps going until it returns. So add `return` right after it.

### Passing data down the chain

Middleware can store values in the context, and later handlers can read them:

```go
// in middleware
c.Set("userID", uint(42))

// in a handler
userID := c.MustGet("userID").(uint) // panics if missing, fine for required values

// or the safe way
if v, ok := c.Get("userID"); ok {
	userID := v.(uint)
	_ = userID
}
```

### Built in middleware

- `gin.Logger()` and `gin.Recovery()` come with `gin.Default()`
- `gin.BasicAuth(gin.Accounts{"admin": "secret"})` gives quick HTTP basic auth, useful for internal tools

### One important trap: goroutines

A `*gin.Context` is reused by Gin after the request ends. If you start a goroutine that needs it, hand over a **copy**:

```go
r.GET("/report", func(c *gin.Context) {
	cp := c.Copy()
	go func() {
		time.Sleep(3 * time.Second)
		log.Println("finished work for", cp.Request.URL.Path)
	}()
	c.String(http.StatusAccepted, "started")
})
```

### Recap

- Middleware wraps handlers: code before `c.Next()` runs first, code after runs last
- `Abort` stops the chain, then `return`
- `c.Set` and `c.Get` pass data to later handlers
- Use `c.Copy()` before using the context in a goroutine

### Homework

Write a middleware `RequireJSON()` that rejects POST, PUT, and PATCH requests whose Content-Type is not `application/json` with a 415 status.

***

## Lesson 8: Mini Project: In-Memory Book API

**You will learn:** how all the pieces fit together in a real CRUD API, before adding a database.

Create a fresh folder or replace `main.go`. This single file uses a map as a fake database, protected by a mutex because Gin handles requests concurrently.

```go
package main

import (
	"net/http"
	"sort"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}

type BookInput struct {
	Title  string  `json:"title" binding:"required,min=2,max=200"`
	Author string  `json:"author" binding:"required,min=2,max=100"`
	Price  float64 `json:"price" binding:"required,gt=0"`
}

type store struct {
	mu     sync.RWMutex
	books  map[int]Book
	nextID int
}

func main() {
	s := &store{books: map[int]Book{}, nextID: 1}

	r := gin.Default()
	api := r.Group("/api/v1")
	{
		api.GET("/books", s.list)
		api.GET("/books/:id", s.get)
		api.POST("/books", s.create)
		api.PUT("/books/:id", s.update)
		api.DELETE("/books/:id", s.remove)
	}

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}

func (s *store) list(c *gin.Context) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	out := make([]Book, 0, len(s.books))
	for _, b := range s.books {
		out = append(out, b)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].ID < out[j].ID })

	c.JSON(http.StatusOK, out)
}

func (s *store) get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id must be a number"})
		return
	}

	s.mu.RLock()
	book, ok := s.books[id]
	s.mu.RUnlock()

	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
		return
	}
	c.JSON(http.StatusOK, book)
}

func (s *store) create(c *gin.Context) {
	var in BookInput
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	s.mu.Lock()
	book := Book{ID: s.nextID, Title: in.Title, Author: in.Author, Price: in.Price}
	s.books[book.ID] = book
	s.nextID++
	s.mu.Unlock()

	c.JSON(http.StatusCreated, book)
}

func (s *store) update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id must be a number"})
		return
	}

	var in BookInput
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.books[id]; !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
		return
	}
	book := Book{ID: id, Title: in.Title, Author: in.Author, Price: in.Price}
	s.books[id] = book

	c.JSON(http.StatusOK, book)
}

func (s *store) remove(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id must be a number"})
		return
	}

	s.mu.Lock()
	_, ok := s.books[id]
	delete(s.books, id)
	s.mu.Unlock()

	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
		return
	}
	c.Status(http.StatusNoContent)
}
```

Try the whole flow:

```bash
# create
curl -X POST http://localhost:8080/api/v1/books \
  -H "Content-Type: application/json" \
  -d '{"title":"Clean Code","author":"Robert Martin","price":29.99}'

# list
curl http://localhost:8080/api/v1/books

# get one
curl http://localhost:8080/api/v1/books/1

# update
curl -X PUT http://localhost:8080/api/v1/books/1 \
  -H "Content-Type: application/json" \
  -d '{"title":"Clean Code 2nd Ed","author":"Robert Martin","price":34.99}'

# delete
curl -X DELETE http://localhost:8080/api/v1/books/1
```

Also try sending bad data: a missing title, a negative price, an id of `abc`. Watch the error responses.

### What to notice

- Handlers are methods on a struct. That is how a handler gets access to shared things like the data store. We will use the same trick with a database.
- Create returns 201, delete returns 204, missing things return 404.
- Every error path ends in `return`.

The problem: restart the server and all data is gone. Time for a real database.

### Recap

- CRUD is five routes on one resource
- Handler methods on a struct give you access to dependencies
- Return the right status code for each outcome

### Homework

Add a `PATCH /books/:id` that uses pointer fields so clients can update only the price or only the title.

***

## Lesson 9: Database with GORM

**You will learn:** how to connect Gin to a database using GORM, an ORM for Go.

We will use SQLite so there is nothing to install. Later you can swap to Postgres or MySQL by changing one line.

```bash
go get gorm.io/gorm
go get github.com/glebarez/sqlite
```

The `glebarez/sqlite` driver is pure Go, so you do not need a C compiler (a common headache on Windows).

### Models

Create `internal/models/models.go`:

```go
package models

import "time"

type Book struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"not null"`
	Author    string    `json:"author" gorm:"not null"`
	Price     float64   `json:"price" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type User struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	Name         string    `json:"name" gorm:"not null"`
	Email        string    `json:"email" gorm:"uniqueIndex;not null"`
	PasswordHash string    `json:"-" gorm:"not null"` // json:"-" means never send this in a response
	CreatedAt    time.Time `json:"created_at"`
}
```

GORM fills `CreatedAt` and `UpdatedAt` automatically because of their names.

### Connecting

Create `internal/database/database.go`:

```go
package database

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"

	"github.com/yourname/gin-bookstore/internal/models"
)

func Connect(path string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// AutoMigrate creates or updates tables to match your structs
	if err := db.AutoMigrate(&models.Book{}, &models.User{}); err != nil {
		return nil, err
	}
	return db, nil
}
```

For Postgres you would use `gorm.io/driver/postgres` and `postgres.Open(dsn)` instead. Everything else stays the same.

### Handlers using the database

Create `internal/handlers/book.go`:

```go
package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/yourname/gin-bookstore/internal/models"
)

type BookHandler struct {
	DB *gorm.DB
}

type BookInput struct {
	Title  string  `json:"title" binding:"required,min=2,max=200"`
	Author string  `json:"author" binding:"required,min=2,max=100"`
	Price  float64 `json:"price" binding:"required,gt=0"`
}

func parseID(c *gin.Context) (uint, bool) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id must be a positive number"})
		return 0, false
	}
	return uint(id), true
}

func (h *BookHandler) List(c *gin.Context) {
	var books []models.Book
	if err := h.DB.Order("id").Find(&books).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch books"})
		return
	}
	c.JSON(http.StatusOK, books)
}

func (h *BookHandler) Get(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}

	var book models.Book
	if err := h.DB.First(&book, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})
		return
	}
	c.JSON(http.StatusOK, book)
}

func (h *BookHandler) Create(c *gin.Context) {
	var in BookInput
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book := models.Book{Title: in.Title, Author: in.Author, Price: in.Price}
	if err := h.DB.Create(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create book"})
		return
	}
	c.JSON(http.StatusCreated, book)
}

func (h *BookHandler) Update(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}

	var in BookInput
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var book models.Book
	if err := h.DB.First(&book, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})
		return
	}

	book.Title, book.Author, book.Price = in.Title, in.Author, in.Price
	if err := h.DB.Save(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not update book"})
		return
	}
	c.JSON(http.StatusOK, book)
}

func (h *BookHandler) Delete(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}

	result := h.DB.Delete(&models.Book{}, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not delete book"})
		return
	}
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
		return
	}
	c.Status(http.StatusNoContent)
}
```

### GORM cheat moves

```go
db.Create(&obj)                              // INSERT
db.First(&obj, id)                           // SELECT by primary key, error if not found
db.Find(&list)                               // SELECT many
db.Where("title LIKE ?", "%go%").Find(&list) // filter, always use ? placeholders
db.Save(&obj)                                // UPDATE all fields
db.Model(&obj).Update("price", 20)           // UPDATE one field
db.Delete(&models.Book{}, id)                // DELETE
```

The `?` placeholder is not optional. Never build SQL by joining user input into a string. That is how SQL injection happens.

### Recap

- Handlers hold a `*gorm.DB`
- `AutoMigrate` keeps tables in sync with structs
- Distinguish "not found" (404) from real database errors (500)

### Homework

Add a `Stock int` field to `Book`, restart the server, and see that AutoMigrate adds the column. Return it in the responses.

***

## Lesson 10: Project Structure and Config

**You will learn:** how to organize a Gin project so it stays maintainable, and how to load configuration from the environment.

A single `main.go` is fine for learning. For anything bigger, split by responsibility:

```text
gin-bookstore/
  main.go
  .env
  .gitignore
  go.mod
  internal/
    apperr/        custom error type (Lesson 12)
    auth/          JWT helpers (Lesson 11)
    config/        loads environment variables
    database/      database connection
    handlers/      HTTP handlers
    middleware/    custom middleware
    models/        database models
    routes/        all route registration
    validation/    validator setup
```

The `internal` folder is a Go feature: packages inside it cannot be imported by other modules. It tells the world "this is private to my project".

### Config from environment variables

Never hardcode secrets. Install godotenv for local development:

```bash
go get github.com/joho/godotenv
```

Create `.env`:

```text
PORT=8080
DB_PATH=bookstore.db
JWT_SECRET=replace-this-with-a-long-random-string
GIN_MODE=debug
```

Add `.env` and `*.db` to `.gitignore`. Do not commit secrets.

Create `internal/config/config.go`:

```go
package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port      string
	DBPath    string
	JWTSecret string
	GinMode   string
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file found, using system environment")
	}

	cfg := &Config{
		Port:      getEnv("PORT", "8080"),
		DBPath:    getEnv("DB_PATH", "bookstore.db"),
		JWTSecret: os.Getenv("JWT_SECRET"),
		GinMode:   getEnv("GIN_MODE", "debug"),
	}

	if cfg.JWTSecret == "" {
		log.Fatal("JWT_SECRET is required")
	}
	return cfg
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
```

Why do we call `gin.SetMode` ourselves? Gin reads `GIN_MODE` when the package is first loaded, which happens before godotenv has read your `.env` file. So we set the mode explicitly after loading config.

### The validation package

Move `setupValidator` from Lesson 5 into `internal/validation/validation.go`, rename it to `Setup`, and keep the same body:

```go
package validation

// imports: reflect, strings, binding, validator

func Setup() {
	// same code as setupValidator in Lesson 5
}
```

### Routes in one place

Create `internal/routes/routes.go`:

```go
package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/yourname/gin-bookstore/internal/config"
	"github.com/yourname/gin-bookstore/internal/handlers"
)

func Setup(db *gorm.DB, cfg *config.Config) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	books := &handlers.BookHandler{DB: db}

	api := r.Group("/api/v1")
	{
		api.GET("/books", books.List)
		api.GET("/books/:id", books.Get)
		api.POST("/books", books.Create)
		api.PUT("/books/:id", books.Update)
		api.DELETE("/books/:id", books.Delete)
	}
	return r
}
```

And a tiny `main.go`:

```go
package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/yourname/gin-bookstore/internal/config"
	"github.com/yourname/gin-bookstore/internal/database"
	"github.com/yourname/gin-bookstore/internal/routes"
	"github.com/yourname/gin-bookstore/internal/validation"
)

func main() {
	cfg := config.Load()
	gin.SetMode(cfg.GinMode)
	validation.Setup()

	db, err := database.Connect(cfg.DBPath)
	if err != nil {
		log.Fatalf("database: %v", err)
	}

	r := routes.Setup(db, cfg)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}
```

### Why this layout pays off

`routes.Setup` returns a `*gin.Engine` without starting a server. That means your tests can build the whole app and send fake requests to it (Lesson 16). Also, `main.go` stays tiny, so it is easy to read.

If your handlers grow a lot of business logic, add a `services` package between handlers and the database. Handlers deal with HTTP (parse, respond), services deal with rules (what is allowed), and the database layer deals with storage. Do not add that layer until you actually need it.

### Recap

- Split code by responsibility, under `internal/`
- Configuration comes from environment variables, never from source code
- `routes.Setup` builds the engine so tests can reuse it

### Homework

Add a `/healthz` check that also pings the database (use `db.DB()` then `PingContext`) and returns 503 if it fails.

***

## Lesson 11: Authentication with JWT

**You will learn:** how to register users, hash passwords, issue tokens, and protect routes.

### The picture

Think of a hospital. At reception you show your ID once, and they give you a visitor wristband. After that, every ward just glances at the wristband instead of asking for your ID again. A JWT (JSON Web Token) is that wristband: the server hands it out after login, and the client shows it on every request.

A JWT has three parts (header, payload, signature). The payload carries claims like the user id and the expiry time. The signature proves the server created it. Anyone can **read** the payload, so never put secrets inside. They just cannot **change** it without the secret key.

Install the packages:

```bash
go get github.com/golang-jwt/jwt/v5
go get golang.org/x/crypto
```

### Token helpers

Create `internal/auth/jwt.go`:

```go
package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateToken(userID uint, secret string, ttl time.Duration) (string, error) {
	now := time.Now()
	claims := Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(ttl)),
			IssuedAt:  jwt.NewNumericDate(now),
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secret))
}

func ParseToken(tokenStr, secret string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(t *jwt.Token) (any, error) {
		return []byte(secret), nil
	}, jwt.WithValidMethods([]string{"HS256"}))
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}
```

`WithValidMethods` matters. It stops attackers from sending a token that claims to use a different signing algorithm.

### Register and login handlers

Create `internal/handlers/auth.go`:

```go
package handlers

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/yourname/gin-bookstore/internal/auth"
	"github.com/yourname/gin-bookstore/internal/models"
)

type AuthHandler struct {
	DB        *gorm.DB
	JWTSecret string
}

type RegisterInput struct {
	Name     string `json:"name" binding:"required,min=2,max=50"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=72"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (h *AuthHandler) Register(c *gin.Context) {
	var in RegisterInput
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	email := strings.ToLower(strings.TrimSpace(in.Email))

	var count int64
	h.DB.Model(&models.User{}).Where("email = ?", email).Count(&count)
	if count > 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "email already registered"})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not process password"})
		return
	}

	user := models.User{Name: in.Name, Email: email, PasswordHash: string(hash)}
	if err := h.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create user"})
		return
	}
	c.JSON(http.StatusCreated, user)
}

func (h *AuthHandler) Login(c *gin.Context) {
	var in LoginInput
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	email := strings.ToLower(strings.TrimSpace(in.Email))
	err := h.DB.Where("email = ?", email).First(&user).Error

	// Same message for "no such user" and "wrong password" so attackers learn nothing
	if err != nil || bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(in.Password)) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid email or password"})
		return
	}

	token, err := auth.GenerateToken(user.ID, h.JWTSecret, 24*time.Hour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token, "user": user})
}

func (h *AuthHandler) Me(c *gin.Context) {
	userID := c.MustGet("userID").(uint)

	var user models.User
	if err := h.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}
```

Two security points:

- **Never store plain passwords.** bcrypt is slow on purpose, which makes guessing attacks expensive. It also adds a random salt for you.
- The email check followed by insert has a tiny race window. The `uniqueIndex` on the email column is your real protection, the check just gives a friendlier message.

### The auth middleware

Create `internal/middleware/auth.go`:

```go
package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/yourname/gin-bookstore/internal/auth"
)

func Auth(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr, found := strings.CutPrefix(c.GetHeader("Authorization"), "Bearer ")
		if !found || tokenStr == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing or malformed Authorization header"})
			return
		}

		claims, err := auth.ParseToken(tokenStr, secret)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			return
		}

		c.Set("userID", claims.UserID)
		c.Next()
	}
}
```

### Wire it up

Update the route section in `routes.go`:

```go
books := &handlers.BookHandler{DB: db}
authH := &handlers.AuthHandler{DB: db, JWTSecret: cfg.JWTSecret}

api := r.Group("/api/v1")
{
	// public
	api.POST("/auth/register", authH.Register)
	api.POST("/auth/login", authH.Login)
	api.GET("/books", books.List)
	api.GET("/books/:id", books.Get)

	// protected: a group with the auth middleware attached
	protected := api.Group("", middleware.Auth(cfg.JWTSecret))
	{
		protected.GET("/me", authH.Me)
		protected.POST("/books", books.Create)
		protected.PUT("/books/:id", books.Update)
		protected.DELETE("/books/:id", books.Delete)
	}
}
```

Import the `middleware` package too. Now test:

```bash
# register
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{"name":"Test User","email":"test@example.com","password":"password123"}'

# login, copy the token from the response
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","password":"password123"}'

# call a protected route without a token (expect 401)
curl -X POST http://localhost:8080/api/v1/books \
  -H "Content-Type: application/json" \
  -d '{"title":"Go","author":"Someone","price":10}'

# with the token (expect 201)
curl -X POST http://localhost:8080/api/v1/books \
  -H "Authorization: Bearer PASTE_TOKEN_HERE" \
  -H "Content-Type: application/json" \
  -d '{"title":"Go","author":"Someone","price":10}'
```

### Where this simple version falls short

- Tokens cannot be revoked before they expire. Real systems use short-lived access tokens plus refresh tokens.
- No roles yet. A role check is just another middleware that reads the role from the claims (or the database) and aborts with 403.

### Recap

- Hash passwords with bcrypt, never store plain text
- Login returns a signed JWT, clients send it as `Authorization: Bearer <token>`
- A middleware validates the token and stores the user id in the context
- Groups make "public vs protected" trivial

### Homework

Add a `role` column to `User` (default `reader`). Write a `RequireRole("admin")` middleware that returns 403 for everyone else, and protect the delete route with it.

***

## Lesson 12: Error Handling Done Right

**You will learn:** how to stop repeating `c.JSON(...)` error code in every handler and return consistent error responses.

Look at our handlers. Every one repeats the same error blocks, and the format of those errors can drift. Let us centralize it.

### Step 1: a custom error type

Create `internal/apperr/apperr.go`:

```go
package apperr

import "net/http"

type Error struct {
	Status  int    `json:"-"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e *Error) Error() string { return e.Message }

func New(status int, code, message string) *Error {
	return &Error{Status: status, Code: code, Message: message}
}

func BadRequest(msg string) *Error   { return New(http.StatusBadRequest, "bad_request", msg) }
func Unauthorized(msg string) *Error { return New(http.StatusUnauthorized, "unauthorized", msg) }
func Forbidden(msg string) *Error    { return New(http.StatusForbidden, "forbidden", msg) }
func NotFound(what string) *Error    { return New(http.StatusNotFound, "not_found", what+" not found") }
func Conflict(msg string) *Error     { return New(http.StatusConflict, "conflict", msg) }
```

### Step 2: one middleware that turns errors into responses

Gin lets handlers attach errors to the context with `c.Error(err)`. A middleware can then read them after the handler finishes and write the response in one place.

Create `internal/middleware/errors.go`:

```go
package middleware

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/yourname/gin-bookstore/internal/apperr"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// Nothing to do if there were no errors or a response was already written
		if len(c.Errors) == 0 || c.Writer.Written() {
			return
		}
		last := c.Errors.Last()

		var appErr *apperr.Error
		var verrs validator.ValidationErrors

		switch {
		case errors.As(last.Err, &appErr):
			c.JSON(appErr.Status, gin.H{"error": appErr})

		case errors.As(last.Err, &verrs):
			fields := make(map[string]string, len(verrs))
			for _, fe := range verrs {
				fields[fe.Field()] = fieldMessage(fe)
			}
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": gin.H{
				"code":    "validation_failed",
				"message": "some fields are invalid",
				"fields":  fields,
			}})

		case last.IsType(gin.ErrorTypeBind):
			c.JSON(http.StatusBadRequest, gin.H{"error": gin.H{
				"code":    "bad_request",
				"message": "request body is missing or is not valid JSON",
			}})

		default:
			// Unknown error: log the details, show the client nothing sensitive
			log.Printf("unexpected error on %s %s: %v", c.Request.Method, c.Request.URL.Path, last.Err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{
				"code":    "internal",
				"message": "something went wrong",
			}})
		}
	}
}

func fieldMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "this field is required"
	case "email":
		return "must be a valid email address"
	case "min":
		return "must be at least " + fe.Param()
	case "max":
		return "must be at most " + fe.Param()
	case "gt":
		return "must be greater than " + fe.Param()
	case "oneof":
		return "must be one of: " + fe.Param()
	case "eqfield":
		return "must match " + fe.Param()
	default:
		return "is invalid (" + fe.Tag() + ")"
	}
}
```

Register it near the top of your middleware list with `r.Use(middleware.ErrorHandler())`.

### Step 3: handlers get much shorter

Update `parseID` to return an error:

```go
func parseID(c *gin.Context) (uint, error) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return 0, apperr.BadRequest("id must be a positive number")
	}
	return uint(id), nil
}
```

And a handler now looks like this:

```go
func (h *BookHandler) Get(c *gin.Context) {
	id, err := parseID(c)
	if err != nil {
		c.Error(err)
		return
	}

	var book models.Book
	if err := h.DB.First(&book, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.Error(apperr.NotFound("book"))
			return
		}
		c.Error(err) // unknown error, the middleware turns it into a 500
		return
	}
	c.JSON(http.StatusOK, book)
}

func (h *BookHandler) Create(c *gin.Context) {
	var in BookInput
	if err := c.ShouldBindJSON(&in); err != nil {
		c.Error(err).SetType(gin.ErrorTypeBind)
		return
	}
	// ...
}
```

Note that `c.Error` does not stop the handler. You still need `return`.

The result: every error leaves your API in the same shape:

```json
{
  "error": {
    "code": "validation_failed",
    "message": "some fields are invalid",
    "fields": { "price": "must be greater than 0" }
  }
}
```

Frontend developers will love you for this.

### Panics and unknown routes

Replace the default Recovery with a custom one that returns your JSON format:

```go
r.Use(gin.CustomRecovery(func(c *gin.Context, recovered any) {
	log.Printf("panic: %v", recovered)
	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": gin.H{
		"code":    "internal",
		"message": "something went wrong",
	}})
}))

r.NoRoute(func(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"error": gin.H{"code": "not_found", "message": "route not found"}})
})
```

### Recap

- Handlers attach errors with `c.Error`, one middleware writes the response
- Known errors (`apperr`) carry their own status and code
- Unknown errors are logged privately and shown to clients as a generic 500
- Consistent error shape equals happier API users

### Homework

Refactor the remaining book handlers and the auth handlers to use `apperr` and `c.Error`. Add `Conflict("email already registered")` to register.

***

## Lesson 13: Pagination, Search, and Sorting

**You will learn:** how to build a list endpoint that real frontends can use.

Returning every row from `GET /books` works with 10 books and dies with 100,000. Let us add `page`, `limit`, `q` (search), `sort`, and `order`:

```
GET /api/v1/books?page=2&limit=5&q=go&sort=price&order=desc
```

Binding query parameters to a struct gives us defaults and validation for free:

```go
type ListQuery struct {
	Page   int    `form:"page,default=1" binding:"gte=1"`
	Limit  int    `form:"limit,default=10" binding:"gte=1,lte=100"`
	Search string `form:"q"`
	Sort   string `form:"sort,default=id" binding:"oneof=id title price created_at"`
	Order  string `form:"order,default=asc" binding:"oneof=asc desc"`
}
```

Note the tag is `form` even though these come from the query string. That is how Gin names it.

The handler:

```go
func (h *BookHandler) List(c *gin.Context) {
	var q ListQuery
	if err := c.ShouldBindQuery(&q); err != nil {
		c.Error(err).SetType(gin.ErrorTypeBind)
		return
	}

	query := h.DB.Model(&models.Book{})
	if q.Search != "" {
		like := "%" + q.Search + "%"
		query = query.Where("title LIKE ? OR author LIKE ?", like, like)
	}
	// Session makes the query safe to reuse for both Count and Find
	query = query.Session(&gorm.Session{})

	var total int64
	if err := query.Count(&total).Error; err != nil {
		c.Error(err)
		return
	}

	var books []models.Book
	err := query.
		Order(q.Sort + " " + q.Order).
		Limit(q.Limit).
		Offset((q.Page - 1) * q.Limit).
		Find(&books).Error
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": books,
		"meta": gin.H{
			"page":        q.Page,
			"limit":       q.Limit,
			"total":       total,
			"total_pages": (total + int64(q.Limit) - 1) / int64(q.Limit),
		},
	})
}
```

### Why validate sort and order so strictly?

We paste `q.Sort` and `q.Order` directly into the query. Placeholders (`?`) cannot be used for column names. That is only safe because `oneof` allows nothing except known column names. Never skip that check, or you open a SQL injection hole.

### Offset pagination vs cursor pagination

What we built is **offset pagination**: simple, lets users jump to page 7, but slows down on huge tables and can show duplicates if data changes between pages. **Cursor pagination** ("give me 10 items after id 340") scales better for feeds and infinite scroll. Start with offset, switch only when you feel the pain.

### Recap

- Bind the query string into a struct with defaults and validation
- Count first, then fetch the page
- Whitelist anything that goes into `Order`

### Homework

Add a `min_price` and `max_price` filter. Return 422 if `min_price` is greater than `max_price`.

***

## Lesson 14: File Upload and Static Files

**You will learn:** how to accept uploaded files safely and serve files back.

Let us let users upload a cover image for a book. First add a field to the model (AutoMigrate will add the column):

```go
CoverPath string `json:"cover_path,omitempty"`
```

Install a UUID package for unique file names:

```bash
go get github.com/google/uuid
```

### The upload handler

```go
const maxCoverSize = 2 << 20 // 2 MB

func (h *BookHandler) UploadCover(c *gin.Context) {
	id, err := parseID(c)
	if err != nil {
		c.Error(err)
		return
	}

	var book models.Book
	if err := h.DB.First(&book, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.Error(apperr.NotFound("book"))
			return
		}
		c.Error(err)
		return
	}

	file, err := c.FormFile("cover") // "cover" is the form field name
	if err != nil {
		c.Error(apperr.BadRequest("a file in the 'cover' field is required"))
		return
	}
	if file.Size > maxCoverSize {
		c.Error(apperr.BadRequest("file is too large, maximum is 2 MB"))
		return
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	switch ext {
	case ".jpg", ".jpeg", ".png", ".webp":
	default:
		c.Error(apperr.BadRequest("only jpg, png, and webp images are allowed"))
		return
	}

	// Never trust the client's file name. Generate our own.
	name := uuid.NewString() + ext
	dst := filepath.Join("uploads", name)
	if err := c.SaveUploadedFile(file, dst); err != nil {
		c.Error(err)
		return
	}

	book.CoverPath = "/uploads/" + name
	if err := h.DB.Save(&book).Error; err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, book)
}
```

Imports used here: `path/filepath`, `strings`, `github.com/google/uuid`, plus the ones you already have.

### Serving files

In `routes.go`:

```go
_ = os.MkdirAll("uploads", 0o755) // make sure the folder exists

r.Static("/uploads", "./uploads")  // /uploads/abc.png serves ./uploads/abc.png
r.StaticFile("/favicon.ico", "./assets/favicon.ico")

protected.POST("/books/:id/cover", books.UploadCover)
```

Test the upload:

```bash
curl -X POST http://localhost:8080/api/v1/books/1/cover \
  -H "Authorization: Bearer PASTE_TOKEN_HERE" \
  -F "cover=@./photo.png"
```

### Sending a file as a download

```go
c.File("./reports/summary.pdf")                          // opens inline in a browser
c.FileAttachment("./reports/summary.pdf", "summary.pdf") // forces a download
```

### Safety checklist for uploads

- Limit the size. Besides checking `file.Size`, cap the whole request body with a middleware:

```go
func LimitBody(maxBytes int64) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, maxBytes)
		c.Next()
	}
}
```

- Restrict extensions, and for stronger checks read the first 512 bytes and call `http.DetectContentType` to confirm the file really is an image.
- Generate your own file names. A client could send `../../etc/passwd` as a file name.
- Gin keeps small uploads in memory up to `r.MaxMultipartMemory` (32 MB by default) and spills the rest to disk. Lower it if you do not need that much.
- For production, store files in object storage (like S3) instead of the server disk, especially if you run more than one server instance.

### Recap

- `c.FormFile` and `c.SaveUploadedFile` handle uploads
- `r.Static` serves a folder
- Never trust file names, sizes, or extensions from the client

### Homework

Support multiple uploads with `c.MultipartForm()` and a `gallery[]` form field, limiting it to 5 images per request.

***

## Lesson 15: CORS, Rate Limiting, Compression, Logging

**You will learn:** the production middleware that every public API needs.

### CORS

Browsers block JavaScript on `https://myapp.com` from calling an API on `https://api.example.com` unless the API says it is okay. That is CORS. Think of a building with a visitor policy: the front desk will only let in people from companies on the approved list. Without a CORS middleware, your React frontend will hit errors in the browser console even though curl works fine.

```bash
go get github.com/gin-contrib/cors
```

```go
r.Use(cors.New(cors.Config{
	AllowOrigins:     []string{"http://localhost:3000", "https://myapp.com"},
	AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
	AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
	ExposeHeaders:    []string{"Content-Length"},
	AllowCredentials: true,
	MaxAge:           12 * time.Hour,
}))
```

List exact origins. Allowing `*` together with credentials is a security smell, and browsers refuse it anyway.

### Rate limiting

Stop one client from hammering your API. A simple per client limiter using the standard `x/time/rate` package:

```bash
go get golang.org/x/time
```

```go
package middleware

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func RateLimit(perSecond rate.Limit, burst int) gin.HandlerFunc {
	var mu sync.Mutex
	limiters := map[string]*rate.Limiter{}

	return func(c *gin.Context) {
		ip := c.ClientIP()

		mu.Lock()
		l, ok := limiters[ip]
		if !ok {
			l = rate.NewLimiter(perSecond, burst)
			limiters[ip] = l
		}
		mu.Unlock()

		if !l.Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": gin.H{
				"code":    "rate_limited",
				"message": "too many requests, slow down",
			}})
			return
		}
		c.Next()
	}
}
```

This version lives in memory and never cleans old entries, so it is fine for learning and small apps. For production, clean up idle entries or use a Redis based limiter, especially when you run several instances.

### Trusted proxies

`c.ClientIP()` reads headers like `X-Forwarded-For`, which any client can fake. Tell Gin which proxies to trust:

```go
_ = r.SetTrustedProxies(nil)                       // trust nobody, use the direct connection IP
// or, when behind your own reverse proxy:
// _ = r.SetTrustedProxies([]string{"10.0.0.1"})
```

If you skip this, your rate limiter can be bypassed by sending a fake header.

### Compression

```bash
go get github.com/gin-contrib/gzip
```

```go
r.Use(gzip.Gzip(gzip.DefaultCompression))
```

### Request ID and structured logging

A request ID lets you find all logs that belong to one request. Combine it with Go's `log/slog` package for structured logs:

```go
func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.GetHeader("X-Request-ID")
		if id == "" {
			id = uuid.NewString()
		}
		c.Set("request_id", id)
		c.Header("X-Request-ID", id)
		c.Next()
	}
}

func RequestLogger(log *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path

		c.Next()

		log.Info("request",
			"request_id", c.GetString("request_id"),
			"method", c.Request.Method,
			"path", path,
			"status", c.Writer.Status(),
			"latency", time.Since(start).String(),
			"ip", c.ClientIP(),
		)
	}
}
```

### Putting it together

The order of middleware matters. Here is a sensible final `routes.Setup` start:

```go
func Setup(db *gorm.DB, cfg *config.Config) *gin.Engine {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	r := gin.New()
	_ = r.SetTrustedProxies(nil)
	r.MaxMultipartMemory = 8 << 20

	r.Use(
		middleware.RequestID(),                  // first, so everything else can use the id
		middleware.RequestLogger(logger),        // wraps everything, logs after it finishes
		gin.Recovery(),                          // catches panics from everything below
		middleware.ErrorHandler(),               // turns c.Error values into JSON
		cors.New(cors.Config{
			AllowOrigins: []string{"http://localhost:3000"},
			AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
			AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
			MaxAge:       12 * time.Hour,
		}),
		gzip.Gzip(gzip.DefaultCompression),
		middleware.RateLimit(5, 10),             // 5 requests per second, bursts up to 10
		middleware.LimitBody(10<<20),            // 10 MB request cap
	)

	// ... routes from before
	return r
}
```

### Recap

- CORS lets your browser frontend talk to the API
- Rate limit by client IP, and configure trusted proxies correctly
- Request IDs and structured logs make debugging production possible
- Middleware order is part of your design

### Homework

Make the rate limit stricter (1 request per second) just for `/auth/login` by attaching a second limiter to that route only.

***

## Lesson 16: Testing

**You will learn:** how to test your API without starting a real server.

Because `routes.Setup` returns a `*gin.Engine`, and an engine implements `http.Handler`, you can call it directly with fake requests using the standard `httptest` package. No network, no ports.

Create `internal/routes/routes_test.go`:

```go
package routes_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/yourname/gin-bookstore/internal/config"
	"github.com/yourname/gin-bookstore/internal/database"
	"github.com/yourname/gin-bookstore/internal/routes"
)

func newTestRouter(t *testing.T) *gin.Engine {
	t.Helper()
	gin.SetMode(gin.TestMode)

	// A separate in-memory database per test so tests do not affect each other
	name := strings.ReplaceAll(t.Name(), "/", "_")
	db, err := database.Connect("file:" + name + "?mode=memory&cache=shared")
	if err != nil {
		t.Fatalf("connect db: %v", err)
	}

	return routes.Setup(db, &config.Config{JWTSecret: "test-secret"})
}

func doJSON(r http.Handler, method, path, token string, body any) *httptest.ResponseRecorder {
	var buf bytes.Buffer
	if body != nil {
		_ = json.NewEncoder(&buf).Encode(body)
	}

	req := httptest.NewRequest(method, path, &buf)
	req.Header.Set("Content-Type", "application/json")
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestHealth(t *testing.T) {
	r := newTestRouter(t)

	w := doJSON(r, http.MethodGet, "/healthz", "", nil)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
}

func TestCreateBookRequiresAuth(t *testing.T) {
	r := newTestRouter(t)

	w := doJSON(r, http.MethodPost, "/api/v1/books", "", gin.H{
		"title": "Go in Action", "author": "William Kennedy", "price": 39.5,
	})

	if w.Code != http.StatusUnauthorized {
		t.Fatalf("expected 401, got %d", w.Code)
	}
}

func TestBookFlow(t *testing.T) {
	r := newTestRouter(t)

	// register and log in
	w := doJSON(r, http.MethodPost, "/api/v1/auth/register", "", gin.H{
		"name": "Test User", "email": "test@example.com", "password": "password123",
	})
	if w.Code != http.StatusCreated {
		t.Fatalf("register: expected 201, got %d (%s)", w.Code, w.Body.String())
	}

	w = doJSON(r, http.MethodPost, "/api/v1/auth/login", "", gin.H{
		"email": "test@example.com", "password": "password123",
	})
	var login struct {
		Token string `json:"token"`
	}
	if err := json.Unmarshal(w.Body.Bytes(), &login); err != nil || login.Token == "" {
		t.Fatalf("login: no token in response (%s)", w.Body.String())
	}

	// create a book
	w = doJSON(r, http.MethodPost, "/api/v1/books", login.Token, gin.H{
		"title": "Go in Action", "author": "William Kennedy", "price": 39.5,
	})
	if w.Code != http.StatusCreated {
		t.Fatalf("create: expected 201, got %d (%s)", w.Code, w.Body.String())
	}

	// invalid input must be rejected
	w = doJSON(r, http.MethodPost, "/api/v1/books", login.Token, gin.H{
		"title": "x", "author": "y", "price": -5,
	})
	if w.Code != http.StatusUnprocessableEntity {
		t.Fatalf("validation: expected 422, got %d (%s)", w.Code, w.Body.String())
	}
}
```

Run all tests:

```bash
go test ./...
go test -v ./internal/routes
go test -race ./...
```

The `-race` flag checks for data races, which is valuable for any concurrent server.

### Testing tips

- Set `gin.SetMode(gin.TestMode)` to silence debug output.
- Use table driven tests for validation cases: one slice of inputs and expected status codes, one loop.
- For handlers that need a database, use an in-memory SQLite like above. It is fast and needs no cleanup.
- Test the unhappy paths (bad input, missing token, not found) at least as much as the happy path. That is where bugs live.

### Recap

- `httptest.NewRecorder` plus `r.ServeHTTP` tests the full stack without a network
- A fresh in-memory database per test keeps tests independent
- Test failures and permissions, not just success

### Homework

Write a table driven test that sends 6 different bad payloads to `POST /books` and checks that each returns 422.

***

## Lesson 17: Graceful Shutdown and Server Settings

**You will learn:** how to stop your server without cutting off in-flight requests.

`r.Run()` is convenient, but it gives you no control. When you deploy a new version, the old process gets a stop signal. With `r.Run()` it dies immediately and any request in progress is dropped. A graceful shutdown stops accepting new requests, lets current ones finish, and then exits.

It is like a shop at closing time: lock the front door so no new customers enter, but serve everyone already inside before turning off the lights.

We will use the standard `http.Server`, which also lets us set timeouts. Update `main.go`:

```go
package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/yourname/gin-bookstore/internal/config"
	"github.com/yourname/gin-bookstore/internal/database"
	"github.com/yourname/gin-bookstore/internal/routes"
	"github.com/yourname/gin-bookstore/internal/validation"
)

func main() {
	cfg := config.Load()
	gin.SetMode(cfg.GinMode)
	validation.Setup()

	db, err := database.Connect(cfg.DBPath)
	if err != nil {
		log.Fatalf("database: %v", err)
	}

	srv := &http.Server{
		Addr:              ":" + cfg.Port,
		Handler:           routes.Setup(db, cfg),
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	// Start the server in a goroutine so main can wait for a signal
	go func() {
		log.Printf("listening on %s", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("server error: %v", err)
		}
	}()

	// Wait for Ctrl+C or a termination signal from Docker or Kubernetes
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	<-ctx.Done()
	log.Println("shutting down, waiting for requests to finish...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Printf("forced shutdown: %v", err)
	}
	log.Println("bye")
}
```

### Why the timeouts matter

Without timeouts, a slow or malicious client can hold a connection open forever and eventually exhaust your server. A client that sends headers one byte per minute is a classic attack (called Slowloris). `ReadHeaderTimeout` defends against it.

Adjust `WriteTimeout` for your slowest legitimate endpoint. Long running streams (like the SSE example in the next lesson) need special care, because `WriteTimeout` applies to them too.

### Recap

- Use `http.Server` with timeouts for production
- `Shutdown` drains in-flight requests before exiting
- Listen for `SIGTERM`, which is what Docker and Kubernetes send

### Homework

Add a handler that sleeps for 5 seconds. Start a request to it, press Ctrl+C in the server terminal, and confirm the request still completes.

***

## Lesson 18: Bonus: HTML Templates, SSE, Swagger

**You will learn:** three extras you will meet sooner or later.

### HTML templates

Gin can render server side HTML, useful for admin pages or simple sites:

```go
r.LoadHTMLGlob("templates/*")

r.GET("/", func(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Bookstore",
		"books": []string{"Clean Code", "The Go Programming Language"},
	})
})
```

And `templates/index.tmpl`:

```html
<!DOCTYPE html>
<html>
<head><title>{{ .title }}</title></head>
<body>
  <h1>{{ .title }}</h1>
  <ul>
    {{ range .books }}
      <li>{{ . }}</li>
    {{ end }}
  </ul>
</body>
</html>
```

Go's `html/template` escapes values automatically, which protects you from cross site scripting.

### Server Sent Events (SSE)

When the server needs to push updates to the browser (live notifications, progress bars), SSE is simpler than WebSockets because it is plain HTTP in one direction:

```go
r.GET("/events", func(c *gin.Context) {
	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	c.Stream(func(w io.Writer) bool {
		select {
		case <-c.Request.Context().Done(): // client disconnected
			return false
		case t := <-ticker.C:
			c.SSEvent("tick", t.Format(time.RFC3339))
			return true // keep streaming
		}
	})
})
```

Browser side:

```js
const es = new EventSource("http://localhost:8080/events");
es.addEventListener("tick", (e) => console.log(e.data));
```

Remember the `WriteTimeout` warning from Lesson 17. For streaming endpoints, you may need to relax it or use `http.ResponseController` to extend the deadline per request.

### Swagger documentation

Auto generated API docs are very handy for teams. The `swaggo` tools read comments above your handlers:

```bash
go install github.com/swaggo/swag/cmd/swag@latest
go get github.com/swaggo/gin-swagger
go get github.com/swaggo/files
```

Annotate a handler:

```go
// List godoc
// @Summary  List books
// @Tags     books
// @Produce  json
// @Param    page  query  int  false  "page number"
// @Success  200   {object}  map[string]any
// @Router   /books [get]
func (h *BookHandler) List(c *gin.Context) { /* ... */ }
```

And add general info above `main()`:

```go
// @title       Bookstore API
// @version     1.0
// @BasePath    /api/v1
```

Generate the docs and register the UI route:

```bash
swag init
```

```go
import (
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"
	_ "github.com/yourname/gin-bookstore/docs" // generated by swag init
)

r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
```

Open `http://localhost:8080/swagger/index.html` and you have interactive docs. In production, consider hiding this route or protecting it.

### Recap

- `LoadHTMLGlob` and `c.HTML` for server side pages
- `c.Stream` and `c.SSEvent` for one way live updates
- swaggo generates interactive API docs from comments

***

## Lesson 19: Docker and Deployment

**You will learn:** how to package the API in a container and what to check before going live.

### Dockerfile

Because we used a pure Go SQLite driver, no C compiler is needed. Create `Dockerfile` (use the same Go version as in your `go.mod`):

```dockerfile
FROM golang:1.23-alpine

WORKDIR /app

# Copy dependency files first so Docker can cache the download step
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o server .

ENV GIN_MODE=release
EXPOSE 8080

CMD ["./server"]
```

This is the simple single stage version. Once you are comfortable, look up "Docker multi stage builds" to shrink the final image to a few megabytes by copying only the compiled binary into a tiny base image.

Create `.dockerignore`:

```text
.env
*.db
uploads
.git
```

Build and run:

```bash
docker build -t gin-bookstore .
docker run -p 8080:8080 -e JWT_SECRET=use-a-long-random-secret -e DB_PATH=/data/bookstore.db -v bookstore-data:/data gin-bookstore
```

The `-v` option mounts a named volume so your SQLite file survives container restarts. Without it, your data disappears with the container.

### Production checklist

- `GIN_MODE=release` (hides debug output and route listing, slightly faster)
- Secrets come from environment variables or a secret manager, never from git
- Terminate HTTPS at a reverse proxy like Nginx or Caddy, and set `SetTrustedProxies` to that proxy
- CORS allows only your real frontend origins
- Rate limiting and body size limits enabled
- Graceful shutdown and server timeouts configured
- A `/healthz` endpoint for your platform's health checks
- A real database server (Postgres or MySQL) once you have more than one instance or need serious backups. With GORM it is a driver and DSN change
- Structured logs going to stdout, collected by your platform
- Automated tests running in CI before every deploy

### Recap

- A short Dockerfile gets you a portable build
- Release mode, secrets, HTTPS, and limits are the production basics

### Homework

Switch the database to Postgres using `gorm.io/driver/postgres` and run both services with Docker Compose.

***

## Lesson 20: Cheat Sheet, Common Mistakes, Next Steps

### Cheat sheet

Routing

- `r := gin.Default()` or `gin.New()`
- `r.GET/POST/PUT/PATCH/DELETE(path, handlers...)`
- `g := r.Group("/api", middleware...)`
- `r.Static(urlPath, dir)`, `r.NoRoute(h)`

Reading input

- `c.Param("id")`, `c.Query("q")`, `c.DefaultQuery("p", "1")`
- `c.PostForm("name")`, `c.FormFile("file")`, `c.GetHeader("X-Key")`
- `c.ShouldBindJSON(&v)`, `ShouldBindQuery`, `ShouldBindUri`, `ShouldBindHeader`, `ShouldBind`

Writing output

- `c.JSON(status, v)`, `c.String`, `c.Status`, `c.Redirect`, `c.File`, `c.FileAttachment`, `c.HTML`

Flow control

- `c.Next()` runs the rest of the chain
- `c.Abort()` and `c.AbortWithStatusJSON(...)` stop it
- `c.Set(k, v)`, `c.Get(k)`, `c.MustGet(k)` share data
- `c.Error(err)` records an error for the error middleware
- `c.Copy()` before using the context in a goroutine

Validation tags

- `required`, `min`, `max`, `len`, `gt`, `gte`, `lt`, `lte`, `email`, `url`, `uuid`, `oneof`, `eqfield`, `omitempty`, `dive`

### Common mistakes

1. **Forgetting `return` after sending a response.** The handler keeps running and may write twice.
2. **Using `Bind` instead of `ShouldBind`.** You lose control of the error response.
3. **Using `required` on values where zero is valid.** Use a pointer.
4. **Using the context in a goroutine without `Copy()`.** Data races and weird bugs.
5. **Running debug mode in production.** Set `GIN_MODE=release`.
6. **Not setting trusted proxies.** Clients can fake their IP.
7. **Putting secrets in code.** Use environment variables.
8. **Building SQL from strings.** Always use placeholders, whitelist column names.
9. **Returning raw database errors to clients.** Log them, send a generic message.
10. **No timeouts on the server.** One slow client can hold resources forever.
11. **Everything in `main.go`.** Fine for day one, painful by day thirty.
12. **No tests for failure cases.** Auth and validation bugs hide there.

### Practice projects

Build these without copying from this guide, then compare:

1. **Todo API** with users, so each user sees only their own todos
2. **URL shortener** with redirects and click counting
3. **Blog API** with posts, comments, and tags (practice GORM relations: `HasMany`, `ManyToMany`)
4. **Expense tracker** with monthly summaries and CSV export
5. **File sharing service** with expiring download links

Stretch goals once those are done: refresh tokens, role based access control, Redis caching, WebSocket chat, background jobs, OpenTelemetry tracing.

### Where to go next

- The official docs at gin-gonic.com/docs for the complete feature list
- The Gin repository's examples folder on GitHub, to see features like binding variants and custom renderers in action
- pkg.go.dev/github.com/gin-gonic/gin for the full API reference
- Read the source of `gin.Context`. It is surprisingly readable and teaches you a lot
- Compare Gin with the standard library router (Go 1.22 added method and path pattern matching to `net/http`). Knowing when you do not need a framework is a real skill

You now know enough Gin to build and ship real APIs. Pick a practice project, start it today, and get stuck on purpose. Happy coding!
