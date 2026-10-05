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
- [Lesson 11: Error Handling Done Right](#lesson-11-error-handling-done-right)
  - [Step 1: a custom error type](#step-1-a-custom-error-type)
  - [Step 2: one middleware that turns errors into responses](#step-2-one-middleware-that-turns-errors-into-responses)
  - [Step 3: handlers get much shorter](#step-3-handlers-get-much-shorter)
  - [Panics and unknown routes](#panics-and-unknown-routes)
- [Lesson 12: Authentication and Authorization (Production Style)](#lesson-12-authentication-and-authorization-production-style)
  - [The big picture](#the-big-picture)
  - [Threats and defenses](#threats-and-defenses)
  - [Install packages](#install-packages)
  - [Step 1: Passwords](#step-1-passwords)
  - [Step 2: Tokens](#step-2-tokens)
  - [Step 3: The auth service](#step-3-the-auth-service)
  - [Step 4: Handlers and cookies](#step-4-handlers-and-cookies)
  - [Step 5: Authentication and role middleware](#step-5-authentication-and-role-middleware)
  - [Step 6: Ownership checks](#step-6-ownership-checks)
  - [Step 7: Wire up the routes](#step-7-wire-up-the-routes)
  - [Step 8: First admin and housekeeping](#step-8-first-admin-and-housekeeping)
  - [Try it with curl](#try-it-with-curl)
  - [Design decisions and trade-offs](#design-decisions-and-trade-offs)
  - [Where to go from here](#where-to-go-from-here)
- [Lesson 13: Pagination, Search, and Sorting](#lesson-13-pagination-search-and-sorting)
  - [Why validate sort and order so strictly?](#why-validate-sort-and-order-so-strictly)
  - [Offset pagination vs cursor pagination](#offset-pagination-vs-cursor-pagination)
- [Lesson 14: File Upload and Static Files](#lesson-14-file-upload-and-static-files)
  - [The upload handler](#the-upload-handler)
  - [Serving files](#serving-files)
  - [Sending a file as a download](#sending-a-file-as-a-download)
  - [Safety checklist for uploads](#safety-checklist-for-uploads)
- [Lesson 15: CORS, Rate Limiting, Compression, Security Headers, Logging](#lesson-15-cors-rate-limiting-compression-security-headers-logging)
  - [CORS](#cors)
  - [Rate limiting](#rate-limiting)
  - [Trusted proxies](#trusted-proxies)
  - [Compression](#compression)
  - [Security headers](#security-headers)
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
- Production style authentication: access and refresh tokens, roles, and ownership checks
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

Create `internal/models/models.go`. These models are already shaped for a real project: users have roles and lockout fields, books know who owns them, and refresh tokens get their own table (you will use all of that in Lesson 12).

```go
package models

import (
	"strings"
	"time"
)

const (
	RoleUser  = "user"
	RoleAdmin = "admin"
)

type Book struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"not null"`
	Author    string    `json:"author" gorm:"not null"`
	Price     float64   `json:"price" gorm:"not null"`
	OwnerID   uint      `json:"owner_id" gorm:"index"` // who created the book, used for permission checks
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type User struct {
	ID                  uint       `json:"id" gorm:"primaryKey"`
	Name                string     `json:"name" gorm:"not null"`
	Email               string     `json:"email" gorm:"uniqueIndex;not null"`
	PasswordHash        string     `json:"-" gorm:"not null"` // json:"-" means this can never leak in a response
	Role                string     `json:"role" gorm:"not null"`
	IsActive            bool       `json:"is_active" gorm:"not null"`
	FailedLoginAttempts int        `json:"-" gorm:"not null"`
	LockedUntil         *time.Time `json:"-"`
	LastLoginAt         *time.Time `json:"last_login_at,omitempty"`
	CreatedAt           time.Time  `json:"created_at"`
	UpdatedAt           time.Time  `json:"updated_at"`
}

// RefreshToken is one issued refresh token. We store only its hash, never the token itself.
type RefreshToken struct {
	ID        uint       `gorm:"primaryKey"`
	UserID    uint       `gorm:"index;not null"`
	FamilyID  string     `gorm:"index;not null;size:36"` // every token that descends from one login shares a family
	TokenHash string     `gorm:"uniqueIndex;not null;size:64"`
	ExpiresAt time.Time  `gorm:"index;not null"`
	RevokedAt *time.Time `gorm:"index"`
	UserAgent string     `gorm:"size:255"`
	IP        string     `gorm:"size:45"`
	CreatedAt time.Time
}

// NormalizeEmail makes sure "Test@Example.com " and "test@example.com" are the same account.
func NormalizeEmail(email string) string {
	return strings.ToLower(strings.TrimSpace(email))
}
```

A few things worth noticing:

- GORM fills `CreatedAt` and `UpdatedAt` automatically because of their names.
- Pointer fields like `*time.Time` become nullable columns. `nil` means "never happened".
- We avoid `gorm:"default:true"` on booleans. GORM treats `false` as "not set" and would apply the default, which is a classic surprise. We set `IsActive` explicitly in code instead.
- In a Postgres or MySQL project, you would also add real foreign keys (for example `OwnerID` pointing to `users.id`). SQLite does not enforce them by default, so we keep the learning version simple.

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
	// TranslateError turns driver specific errors (like a unique constraint violation)
	// into GORM errors such as gorm.ErrDuplicatedKey
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{TranslateError: true})
	if err != nil {
		return nil, err
	}

	// AutoMigrate creates or updates tables to match your structs
	if err := db.AutoMigrate(&models.Book{}, &models.User{}, &models.RefreshToken{}); err != nil {
		return nil, err
	}
	return db, nil
}
```

For Postgres you would use `gorm.io/driver/postgres` and `postgres.Open(dsn)` instead. Everything else stays the same.

One honest note: `AutoMigrate` is perfect for learning and small projects. Larger teams use versioned migration tools (like golang-migrate or Atlas) so schema changes are reviewed, repeatable, and reversible.

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
    apperr/        custom error type (Lesson 11)
    auth/          password hashing and tokens (Lesson 12)
    config/        loads environment variables
    database/      database connection
    handlers/      HTTP handlers
    middleware/    custom middleware
    models/        database models
    routes/        all route registration
    services/      business logic, starting with auth (Lesson 12)
    validation/    validator setup
```

The `internal` folder is a Go feature: packages inside it cannot be imported by other modules. It tells the world "this is private to my project".

### Config from environment variables

Never hardcode secrets, and never let the app start with a weak one. Install godotenv for local development:

```bash
go get github.com/joho/godotenv
```

Create `.env`:

```text
PORT=8080
DB_PATH=bookstore.db
GIN_MODE=debug

# Authentication
JWT_SECRET=change-me-this-is-only-for-local-development-1234
JWT_ISSUER=gin-bookstore
JWT_AUDIENCE=gin-bookstore-api
ACCESS_TOKEN_TTL=15m
REFRESH_TOKEN_TTL=168h
BCRYPT_COST=12

# Cookies and browsers
COOKIE_SECURE=false
COOKIE_DOMAIN=
ALLOWED_ORIGINS=http://localhost:3000

# Brute force protection on /auth routes (per IP)
AUTH_RATE_PER_MINUTE=10
AUTH_RATE_BURST=5

# Optional: create the first admin on startup
ADMIN_EMAIL=
ADMIN_PASSWORD=
```

Add `.env` and `*.db` to `.gitignore`. Do not commit secrets. In production, these values come from your platform's environment or a secret manager.

Create `internal/config/config.go`:

```go
package config

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Port    string
	DBPath  string
	GinMode string

	JWTSecret   string
	JWTIssuer   string
	JWTAudience string
	AccessTTL   time.Duration
	RefreshTTL  time.Duration
	BcryptCost  int

	CookieSecure   bool
	CookieDomain   string
	AllowedOrigins []string

	AuthRatePerMinute int
	AuthRateBurst     int

	AdminEmail    string
	AdminPassword string
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file found, using system environment")
	}

	ginMode := getEnv("GIN_MODE", "debug")

	cfg := &Config{
		Port:    getEnv("PORT", "8080"),
		DBPath:  getEnv("DB_PATH", "bookstore.db"),
		GinMode: ginMode,

		JWTSecret:   os.Getenv("JWT_SECRET"),
		JWTIssuer:   getEnv("JWT_ISSUER", "gin-bookstore"),
		JWTAudience: getEnv("JWT_AUDIENCE", "gin-bookstore-api"),
		AccessTTL:   getDuration("ACCESS_TOKEN_TTL", 15*time.Minute),
		RefreshTTL:  getDuration("REFRESH_TOKEN_TTL", 7*24*time.Hour),
		BcryptCost:  getInt("BCRYPT_COST", 12),

		// Secure cookies by default in release mode. Browsers only send them over HTTPS.
		CookieSecure:   getBool("COOKIE_SECURE", ginMode == "release"),
		CookieDomain:   os.Getenv("COOKIE_DOMAIN"),
		AllowedOrigins: splitList(getEnv("ALLOWED_ORIGINS", "http://localhost:3000")),

		AuthRatePerMinute: getInt("AUTH_RATE_PER_MINUTE", 10),
		AuthRateBurst:     getInt("AUTH_RATE_BURST", 5),

		AdminEmail:    os.Getenv("ADMIN_EMAIL"),
		AdminPassword: os.Getenv("ADMIN_PASSWORD"),
	}

	// Fail fast. A server that starts with a weak secret is worse than one that refuses to start.
	if len(cfg.JWTSecret) < 32 {
		log.Fatal("JWT_SECRET must be at least 32 characters long")
	}
	if cfg.BcryptCost < 10 || cfg.BcryptCost > 14 {
		log.Fatal("BCRYPT_COST must be between 10 and 14")
	}
	if cfg.AccessTTL <= 0 || cfg.RefreshTTL <= cfg.AccessTTL {
		log.Fatal("REFRESH_TOKEN_TTL must be longer than ACCESS_TOKEN_TTL, and both must be positive")
	}
	return cfg
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func getInt(key string, fallback int) int {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}
	n, err := strconv.Atoi(v)
	if err != nil {
		log.Fatalf("%s must be a whole number, got %q", key, v)
	}
	return n
}

func getBool(key string, fallback bool) bool {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}
	b, err := strconv.ParseBool(v)
	if err != nil {
		log.Fatalf("%s must be true or false, got %q", key, v)
	}
	return b
}

func getDuration(key string, fallback time.Duration) time.Duration {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}
	d, err := time.ParseDuration(v)
	if err != nil {
		log.Fatalf("%s must be a duration like 15m or 168h, got %q", key, v)
	}
	return d
}

func splitList(s string) []string {
	var out []string
	for _, part := range strings.Split(s, ",") {
		if p := strings.TrimSpace(part); p != "" {
			out = append(out, p)
		}
	}
	return out
}
```

Why do we call `gin.SetMode` ourselves later? Gin reads `GIN_MODE` when the package is first loaded, which happens before godotenv has read your `.env` file. So we set the mode explicitly after loading config.

### The validation package

Move `setupValidator` from Lesson 5 into `internal/validation/validation.go`, rename it to `Setup`, and add one important line at the top:

```go
package validation

import (
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func Setup() {
	// Reject JSON that contains fields we did not define. This blocks "mass assignment",
	// for example a client sending {"role": "admin"} to the registration endpoint.
	binding.EnableDecoderDisallowUnknownFields = true

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

If your handlers grow a lot of business logic, add a `services` package between handlers and the database. Handlers deal with HTTP (parse, respond), services deal with rules (what is allowed), and the database layer deals with storage. Do not add that layer until you actually need it. We will need it in Lesson 12, because authentication has real rules and database transactions that do not belong in a handler.

### Recap

- Split code by responsibility, under `internal/`
- Configuration comes from environment variables, never from source code
- `routes.Setup` builds the engine so tests can reuse it

### Homework

Add a `/healthz` check that also pings the database (use `db.DB()` then `PingContext`) and returns 503 if it fails.

***

## Lesson 11: Error Handling Done Right

**You will learn:** how to stop repeating `c.JSON(...)` error code in every handler and return consistent, safe error responses. We do this before authentication, because the auth code in the next lesson is built on top of it.

Look at our handlers. Every one repeats the same error blocks, and the format of those errors can drift. Let us centralize it.

### Step 1: a custom error type

Create `internal/apperr/apperr.go`:

```go
package apperr

import "net/http"

type Error struct {
	Status  int               `json:"-"`
	Code    string            `json:"code"`
	Message string            `json:"message"`
	Fields  map[string]string `json:"fields,omitempty"`
}

func (e *Error) Error() string { return e.Message }

func New(status int, code, message string) *Error {
	return &Error{Status: status, Code: code, Message: message}
}

func BadRequest(msg string) *Error      { return New(http.StatusBadRequest, "bad_request", msg) }
func Unauthorized(msg string) *Error    { return New(http.StatusUnauthorized, "unauthorized", msg) }
func Forbidden(msg string) *Error       { return New(http.StatusForbidden, "forbidden", msg) }
func NotFound(what string) *Error       { return New(http.StatusNotFound, "not_found", what+" not found") }
func Conflict(msg string) *Error        { return New(http.StatusConflict, "conflict", msg) }
func TooManyRequests(msg string) *Error { return New(http.StatusTooManyRequests, "rate_limited", msg) }

// Validation is for business rule failures that name specific fields.
func Validation(fields map[string]string) *Error {
	return &Error{
		Status:  http.StatusUnprocessableEntity,
		Code:    "validation_failed",
		Message: "some fields are invalid",
		Fields:  fields,
	}
}
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
		var tooBig *http.MaxBytesError

		switch {
		case errors.As(last.Err, &appErr):
			c.JSON(appErr.Status, gin.H{"error": appErr})

		case errors.As(last.Err, &verrs):
			fields := make(map[string]string, len(verrs))
			for _, fe := range verrs {
				fields[fe.Field()] = fieldMessage(fe)
			}
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": apperr.Validation(fields)})

		case errors.As(last.Err, &tooBig):
			c.JSON(http.StatusRequestEntityTooLarge, gin.H{"error": gin.H{
				"code":    "payload_too_large",
				"message": "request body is too large",
			}})

		case last.IsType(gin.ErrorTypeBind):
			c.JSON(http.StatusBadRequest, gin.H{"error": gin.H{
				"code":    "bad_request",
				"message": "request body is missing, malformed, or contains unknown fields",
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

Create `internal/handlers/helpers.go` and move `parseID` there (delete the old copy from `book.go`). While we are here, add a small `bindJSON` helper that every handler will reuse:

```go
package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/yourname/gin-bookstore/internal/apperr"
)

func parseID(c *gin.Context) (uint, error) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return 0, apperr.BadRequest("id must be a positive number")
	}
	return uint(id), nil
}

// bindJSON binds the request body and records an error for the middleware if it fails.
// It returns false when the handler should stop.
func bindJSON(c *gin.Context, obj any) bool {
	if err := c.ShouldBindJSON(obj); err != nil {
		c.Error(err).SetType(gin.ErrorTypeBind)
		return false
	}
	return true
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
	if !bindJSON(c, &in) {
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

Frontend developers will love you for this. Also notice what we did **not** do: unknown errors are logged privately and the client only sees "something went wrong". Raw database errors can reveal table names and queries, so they never leave the server.

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
- Known errors (`apperr`) carry their own status, code, and optional field messages
- Unknown errors are logged privately and shown to clients as a generic 500
- A consistent error shape makes your API pleasant to use

### Homework

Refactor `List`, `Update`, and `Delete` in `book.go` to use `apperr`, `c.Error`, and `bindJSON`. In the next lesson, all of the auth code is written in this style.

***

## Lesson 12: Authentication and Authorization (Production Style)

**You will learn:** how to build login the way real systems do it: safe password storage, short-lived access tokens, rotating refresh tokens, brute force protection, roles, and ownership checks.

A quick honesty check first. The classic tutorial version (one JWT that lasts 24 hours, stored wherever) works in a demo, but it falls apart in the real world: you cannot log anyone out, a stolen token works for a day, and nothing slows down password guessing. This lesson fixes all of that.

### The big picture

Think of a hotel. When you check in, the front desk gives you a **key card** that opens your room, but it only works for a short time. That is the access token. Along with it you get a **booking slip** that proves you are a real guest. When your key card expires, you show the booking slip at the desk and receive a fresh key card **and a new booking slip**. The old slip is cancelled immediately.

Now imagine a thief photocopies your booking slip. The moment the thief uses the old slip, the desk notices that a cancelled slip was presented again. That can only mean it was copied, so the hotel cancels **every** slip from that stay and you must check in again. That is exactly what refresh token rotation with reuse detection does.

```text
POST /auth/login     ->  access token (15 min, in JSON)  +  refresh token (7 days, HttpOnly cookie)
GET  /books ...      ->  Authorization: Bearer <access token>
Access token expired ->  POST /auth/refresh (cookie)     ->  NEW access token + NEW refresh token
Old refresh token used again  ->  the whole session family is revoked, user must log in again
```

### Threats and defenses

What are we protecting against, and what stops it?

- **Stolen database:** passwords are hashed with bcrypt, refresh tokens are stored only as SHA-256 hashes
- **Password guessing:** a per IP rate limit on auth routes, a per account lockout, a password policy
- **Stolen access token:** it dies after 15 minutes
- **Stolen refresh token:** rotation plus reuse detection, and it lives in an `HttpOnly` cookie that JavaScript cannot read
- **CSRF (a malicious site making your browser send requests):** `SameSite=Strict` cookie, cookie limited to `/api/v1/auth`, and the access token travels in a header, not a cookie
- **Finding out which emails exist:** same error message and same response time for "no such user" and "wrong password"
- **Reading or changing other users' data:** ownership checks on every write
- **Becoming an admin by sending `"role": "admin"`:** roles come from the database, and unknown JSON fields are rejected
- **Weak server secret:** the app refuses to start with a short `JWT_SECRET`

### Install packages

```bash
go get github.com/golang-jwt/jwt/v5
go get golang.org/x/crypto
go get github.com/google/uuid
go get golang.org/x/time
go mod tidy
```

Make sure you applied the updated models (Lesson 9), config (Lesson 10), and `ErrorHandler` plus `apperr` (Lesson 11) before continuing.

### Step 1: Passwords

Create `internal/auth/password.go`:

```go
package auth

import (
	"errors"
	"strings"
	"unicode/utf8"

	"golang.org/x/crypto/bcrypt"
)

const (
	minPasswordLength = 10
	maxPasswordBytes  = 72 // bcrypt cannot use more than 72 bytes
)

// A tiny sample. In a real project, load a large list of leaked passwords
// or call a breach checking service.
var commonPasswords = map[string]struct{}{
	"password123": {}, "1234567890": {}, "qwertyuiop": {}, "iloveyou123": {},
	"letmein1234": {}, "welcome1234": {}, "admin12345": {}, "password1234": {},
}

// ValidatePassword enforces a modern policy: length matters more than weird symbols.
func ValidatePassword(pw string) error {
	if utf8.RuneCountInString(pw) < minPasswordLength {
		return errors.New("must be at least 10 characters long")
	}
	if len(pw) > maxPasswordBytes {
		return errors.New("must be at most 72 bytes long")
	}
	if _, bad := commonPasswords[strings.ToLower(pw)]; bad {
		return errors.New("is too common, choose something less guessable")
	}
	return nil
}

func HashPassword(password string, cost int) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func CheckPassword(hash, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
```

Why these choices:

- **bcrypt is slow on purpose.** At cost 12, one guess takes roughly a quarter of a second, which makes bulk guessing very expensive. It also adds a random salt for you. (argon2id is the other good choice if you prefer it.)
- **Length over complexity.** Forcing symbols and digits makes people write `Password1!`. A long passphrase is stronger and easier to remember.
- **The 72 byte limit** is a bcrypt limitation. We reject longer input clearly instead of silently cutting it.

### Step 2: Tokens

Create `internal/auth/tokens.go`:

```go
package auth

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var ErrInvalidToken = errors.New("invalid token")

type AccessClaims struct {
	Role string `json:"role"`
	jwt.RegisteredClaims
}

type TokenManager struct {
	secret    []byte
	issuer    string
	audience  string
	accessTTL time.Duration
}

func NewTokenManager(secret, issuer, audience string, accessTTL time.Duration) *TokenManager {
	return &TokenManager{
		secret:    []byte(secret),
		issuer:    issuer,
		audience:  audience,
		accessTTL: accessTTL,
	}
}

// NewAccessToken returns a signed JWT and the time it expires.
func (m *TokenManager) NewAccessToken(userID uint, role string) (string, time.Time, error) {
	now := time.Now()
	expires := now.Add(m.accessTTL)

	claims := AccessClaims{
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    m.issuer,
			Subject:   strconv.FormatUint(uint64(userID), 10),
			Audience:  jwt.ClaimStrings{m.audience},
			ExpiresAt: jwt.NewNumericDate(expires),
			NotBefore: jwt.NewNumericDate(now),
			IssuedAt:  jwt.NewNumericDate(now),
			ID:        uuid.NewString(), // unique id for this token, handy for audit logs or a denylist
		},
	}

	signed, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(m.secret)
	if err != nil {
		return "", time.Time{}, err
	}
	return signed, expires, nil
}

// ParseAccessToken verifies the signature AND the claims that matter.
func (m *TokenManager) ParseAccessToken(tokenStr string) (*AccessClaims, error) {
	claims := &AccessClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token) (any, error) { return m.secret, nil },
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}), // only the algorithm we use
		jwt.WithIssuer(m.issuer),
		jwt.WithAudience(m.audience),
		jwt.WithExpirationRequired(),
		jwt.WithLeeway(5*time.Second), // tolerate small clock differences
	)
	if err != nil || !token.Valid {
		return nil, ErrInvalidToken
	}
	return claims, nil
}

// NewRefreshToken creates an opaque random token. The raw value goes to the client,
// only the hash is stored in the database.
func NewRefreshToken() (raw string, hash string, err error) {
	b := make([]byte, 32) // 256 bits of randomness
	if _, err = rand.Read(b); err != nil {
		return "", "", err
	}
	raw = base64.RawURLEncoding.EncodeToString(b)
	return raw, HashRefreshToken(raw), nil
}

func HashRefreshToken(raw string) string {
	sum := sha256.Sum256([]byte(raw))
	return hex.EncodeToString(sum[:])
}
```

The reasoning:

- **Access token is a JWT.** Servers can verify it without touching the database, which is fast. But a JWT cannot be revoked once issued, so we keep it very short lived (15 minutes).
- **We check more than the signature.** The parser also demands the right algorithm, issuer, audience, and an expiry. This blocks forged tokens, tokens meant for another service, and the old "algorithm switching" attack.
- **Refresh token is not a JWT.** It is 32 random bytes. Nothing inside it needs to be read by the client, and being opaque means we can revoke it from the database at any time.
- **Why SHA-256 and not bcrypt for refresh tokens?** Passwords are guessable, so they need a slow hash. A 256 bit random token cannot be guessed, so a fast hash is enough, and it lets us look the token up directly by its hash.

### Step 3: The auth service

Handlers should deal with HTTP. The rules (lockout, rotation, revocation) and database transactions belong in a service. Create `internal/services/auth.go`:

```go
package services

import (
	"context"
	"errors"
	"log/slog"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/yourname/gin-bookstore/internal/apperr"
	"github.com/yourname/gin-bookstore/internal/auth"
	"github.com/yourname/gin-bookstore/internal/config"
	"github.com/yourname/gin-bookstore/internal/models"
)

const (
	maxFailedLogins = 5
	lockDuration    = 15 * time.Minute
)

var (
	// One message for "no such user" and "wrong password" so attackers learn nothing
	errInvalidCredentials = apperr.Unauthorized("invalid email or password")
	errInvalidRefresh     = apperr.Unauthorized("invalid or expired refresh token")
)

type ClientMeta struct {
	IP        string
	UserAgent string
}

type TokenPair struct {
	AccessToken      string
	AccessExpiresAt  time.Time
	RefreshToken     string
	RefreshExpiresAt time.Time
}

type AuthService struct {
	db        *gorm.DB
	tokens    *auth.TokenManager
	cfg       *config.Config
	log       *slog.Logger
	dummyHash string
}

func NewAuthService(db *gorm.DB, tokens *auth.TokenManager, cfg *config.Config, log *slog.Logger) *AuthService {
	// A throwaway hash. We compare against it when the email does not exist,
	// so that response time does not reveal which emails are registered.
	dummy, err := auth.HashPassword("not-a-real-password", cfg.BcryptCost)
	if err != nil {
		panic(err) // cannot happen with a valid cost, and Load already checked the cost
	}
	return &AuthService{db: db, tokens: tokens, cfg: cfg, log: log, dummyHash: dummy}
}

func (s *AuthService) Register(ctx context.Context, name, email, password string) (*models.User, error) {
	if err := auth.ValidatePassword(password); err != nil {
		return nil, apperr.Validation(map[string]string{"password": err.Error()})
	}
	hash, err := auth.HashPassword(password, s.cfg.BcryptCost)
	if err != nil {
		return nil, err
	}

	user := models.User{
		Name:         strings.TrimSpace(name),
		Email:        models.NormalizeEmail(email),
		PasswordHash: hash,
		Role:         models.RoleUser, // always. Never take the role from the client.
		IsActive:     true,
	}
	if err := s.db.WithContext(ctx).Create(&user).Error; err != nil {
		if isUniqueViolation(err) {
			return nil, apperr.Conflict("an account with this email already exists")
		}
		return nil, err
	}

	s.log.Info("user registered", "user_id", user.ID)
	return &user, nil
}

func (s *AuthService) Login(ctx context.Context, email, password string, meta ClientMeta) (*models.User, *TokenPair, error) {
	db := s.db.WithContext(ctx)

	var user models.User
	err := db.Where("email = ?", models.NormalizeEmail(email)).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// Do the same expensive work as a real check, so timing gives nothing away
		auth.CheckPassword(s.dummyHash, password)
		s.log.Warn("login failed", "reason", "unknown_email", "ip", meta.IP)
		return nil, nil, errInvalidCredentials
	}
	if err != nil {
		return nil, nil, err
	}

	if user.LockedUntil != nil && time.Now().Before(*user.LockedUntil) {
		s.log.Warn("login blocked", "reason", "account_locked", "user_id", user.ID, "ip", meta.IP)
		return nil, nil, apperr.TooManyRequests("too many failed attempts, try again later")
	}

	if !auth.CheckPassword(user.PasswordHash, password) {
		s.recordFailedLogin(db, &user, meta)
		return nil, nil, errInvalidCredentials
	}
	if !user.IsActive {
		return nil, nil, apperr.Forbidden("this account is disabled")
	}

	now := time.Now()
	err = db.Model(&user).Updates(map[string]any{
		"failed_login_attempts": 0,
		"locked_until":          nil,
		"last_login_at":         now,
	}).Error
	if err != nil {
		return nil, nil, err
	}

	// A new login starts a new "family" of refresh tokens
	pair, err := s.issueTokens(db, &user, uuid.NewString(), meta)
	if err != nil {
		return nil, nil, err
	}

	s.log.Info("login succeeded", "user_id", user.ID, "ip", meta.IP)
	return &user, pair, nil
}

func (s *AuthService) recordFailedLogin(db *gorm.DB, user *models.User, meta ClientMeta) {
	// Increment inside the database so two parallel attempts cannot overwrite each other
	updates := map[string]any{"failed_login_attempts": gorm.Expr("failed_login_attempts + 1")}

	locked := user.FailedLoginAttempts+1 >= maxFailedLogins
	if locked {
		updates["failed_login_attempts"] = 0
		updates["locked_until"] = time.Now().Add(lockDuration)
	}

	if err := db.Model(&models.User{}).Where("id = ?", user.ID).Updates(updates).Error; err != nil {
		s.log.Error("could not record failed login", "user_id", user.ID, "error", err)
	}
	s.log.Warn("login failed", "reason", "wrong_password", "user_id", user.ID, "ip", meta.IP, "locked", locked)
}

func (s *AuthService) issueTokens(tx *gorm.DB, user *models.User, familyID string, meta ClientMeta) (*TokenPair, error) {
	access, accessExpires, err := s.tokens.NewAccessToken(user.ID, user.Role)
	if err != nil {
		return nil, err
	}
	rawRefresh, refreshHash, err := auth.NewRefreshToken()
	if err != nil {
		return nil, err
	}

	refreshExpires := time.Now().Add(s.cfg.RefreshTTL)
	record := models.RefreshToken{
		UserID:    user.ID,
		FamilyID:  familyID,
		TokenHash: refreshHash,
		ExpiresAt: refreshExpires,
		UserAgent: truncate(meta.UserAgent, 255),
		IP:        meta.IP,
	}
	if err := tx.Create(&record).Error; err != nil {
		return nil, err
	}

	return &TokenPair{
		AccessToken:      access,
		AccessExpiresAt:  accessExpires,
		RefreshToken:     rawRefresh,
		RefreshExpiresAt: refreshExpires,
	}, nil
}

// Refresh rotates the refresh token: the old one is revoked and a new pair is issued.
// If an already used token shows up again, the whole family is revoked.
func (s *AuthService) Refresh(ctx context.Context, rawToken string, meta ClientMeta) (*TokenPair, error) {
	var (
		pair   *TokenPair
		reused bool
		family string
	)

	err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var rt models.RefreshToken
		err := tx.Where("token_hash = ?", auth.HashRefreshToken(rawToken)).First(&rt).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errInvalidRefresh
		}
		if err != nil {
			return err
		}

		// Already used or revoked, and someone is presenting it again.
		// Either a client bug or a stolen token. Kill the whole family.
		if rt.RevokedAt != nil {
			reused, family = true, rt.FamilyID
			return revokeFamily(tx, rt.FamilyID)
		}
		if time.Now().After(rt.ExpiresAt) {
			return errInvalidRefresh
		}

		// Claim the token atomically. If two requests race with the same token,
		// only one of them updates a row.
		res := tx.Model(&models.RefreshToken{}).
			Where("id = ? AND revoked_at IS NULL", rt.ID).
			Update("revoked_at", time.Now())
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected == 0 {
			reused, family = true, rt.FamilyID
			return revokeFamily(tx, rt.FamilyID)
		}

		var user models.User
		if err := tx.First(&user, rt.UserID).Error; err != nil {
			return errInvalidRefresh
		}
		if !user.IsActive {
			return apperr.Forbidden("this account is disabled")
		}

		p, err := s.issueTokens(tx, &user, rt.FamilyID, meta)
		if err != nil {
			return err
		}
		pair = p
		return nil
	})
	if err != nil {
		return nil, err
	}
	if reused {
		s.log.Warn("refresh token reuse detected, session family revoked", "family_id", family, "ip", meta.IP)
		return nil, errInvalidRefresh
	}
	return pair, nil
}

func revokeFamily(tx *gorm.DB, familyID string) error {
	return tx.Model(&models.RefreshToken{}).
		Where("family_id = ? AND revoked_at IS NULL", familyID).
		Update("revoked_at", time.Now()).Error
}

// Logout ends the current session (one device).
func (s *AuthService) Logout(ctx context.Context, rawToken string) error {
	return s.db.WithContext(ctx).Model(&models.RefreshToken{}).
		Where("token_hash = ? AND revoked_at IS NULL", auth.HashRefreshToken(rawToken)).
		Update("revoked_at", time.Now()).Error
}

// LogoutAll ends every session of the user (all devices).
func (s *AuthService) LogoutAll(ctx context.Context, userID uint) error {
	return s.db.WithContext(ctx).Model(&models.RefreshToken{}).
		Where("user_id = ? AND revoked_at IS NULL", userID).
		Update("revoked_at", time.Now()).Error
}

func (s *AuthService) ChangePassword(ctx context.Context, userID uint, current, next string) error {
	user, err := s.GetUser(ctx, userID)
	if err != nil {
		return err
	}
	if !auth.CheckPassword(user.PasswordHash, current) {
		return apperr.Validation(map[string]string{"current_password": "is incorrect"})
	}
	if err := auth.ValidatePassword(next); err != nil {
		return apperr.Validation(map[string]string{"new_password": err.Error()})
	}
	hash, err := auth.HashPassword(next, s.cfg.BcryptCost)
	if err != nil {
		return err
	}

	// Changing the password also signs out every device, in one transaction
	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(user).Update("password_hash", hash).Error; err != nil {
			return err
		}
		return tx.Model(&models.RefreshToken{}).
			Where("user_id = ? AND revoked_at IS NULL", userID).
			Update("revoked_at", time.Now()).Error
	})
}

func (s *AuthService) GetUser(ctx context.Context, id uint) (*models.User, error) {
	var user models.User
	if err := s.db.WithContext(ctx).First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperr.NotFound("user")
		}
		return nil, err
	}
	return &user, nil
}

func isUniqueViolation(err error) bool {
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return true
	}
	msg := strings.ToLower(err.Error())
	return strings.Contains(msg, "unique constraint") || strings.Contains(msg, "duplicate key")
}

func truncate(s string, max int) string {
	if len(s) <= max {
		return s
	}
	return s[:max]
}
```

Take a minute with `Refresh`. It is the heart of the whole system, and a few details matter:

- It runs inside a **transaction**, so "revoke the old token" and "issue the new one" either both happen or neither does.
- Revoking uses `WHERE revoked_at IS NULL` and checks `RowsAffected`. That makes the claim atomic: if a thief and the real user both send the same token at the same moment, only one wins.
- On reuse, we **commit** the family revocation and then return the error afterwards. (If we returned the error inside the transaction, GORM would roll the revocation back.)
- Logs describe what happened (`reuse detected`, user id, IP) but never contain passwords or tokens.

### Step 4: Handlers and cookies

Create `internal/handlers/auth.go`:

```go
package handlers

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/yourname/gin-bookstore/internal/apperr"
	"github.com/yourname/gin-bookstore/internal/config"
	"github.com/yourname/gin-bookstore/internal/middleware"
	"github.com/yourname/gin-bookstore/internal/services"
)

const (
	refreshCookieName = "refresh_token"
	refreshCookiePath = "/api/v1/auth" // the browser only sends the cookie to auth endpoints
)

type AuthHandler struct {
	Svc *services.AuthService
	Cfg *config.Config
}

type registerRequest struct {
	Name     string `json:"name" binding:"required,min=2,max=50"`
	Email    string `json:"email" binding:"required,email,max=254"`
	Password string `json:"password" binding:"required"` // the full policy is checked in the service
}

type loginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type changePasswordRequest struct {
	CurrentPassword string `json:"current_password" binding:"required"`
	NewPassword     string `json:"new_password" binding:"required"`
}

func (h *AuthHandler) Register(c *gin.Context) {
	var in registerRequest
	if !bindJSON(c, &in) {
		return
	}
	user, err := h.Svc.Register(c.Request.Context(), in.Name, in.Email, in.Password)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusCreated, user)
}

func (h *AuthHandler) Login(c *gin.Context) {
	var in loginRequest
	if !bindJSON(c, &in) {
		return
	}
	user, pair, err := h.Svc.Login(c.Request.Context(), in.Email, in.Password, clientMeta(c))
	if err != nil {
		c.Error(err)
		return
	}

	h.setRefreshCookie(c, pair)
	resp := tokenResponse(pair)
	resp["user"] = user
	c.JSON(http.StatusOK, resp)
}

func (h *AuthHandler) Refresh(c *gin.Context) {
	raw, err := c.Cookie(refreshCookieName)
	if err != nil || raw == "" {
		c.Error(apperr.Unauthorized("missing refresh token"))
		return
	}

	pair, err := h.Svc.Refresh(c.Request.Context(), raw, clientMeta(c))
	if err != nil {
		// A rejected token is useless, so stop the browser from sending it again.
		// A temporary server problem should not log the user out, so only clear on auth errors.
		var appErr *apperr.Error
		if errors.As(err, &appErr) {
			h.clearRefreshCookie(c)
		}
		c.Error(err)
		return
	}

	h.setRefreshCookie(c, pair)
	c.JSON(http.StatusOK, tokenResponse(pair))
}

func (h *AuthHandler) Logout(c *gin.Context) {
	if raw, err := c.Cookie(refreshCookieName); err == nil && raw != "" {
		if err := h.Svc.Logout(c.Request.Context(), raw); err != nil {
			c.Error(err)
			return
		}
	}
	h.clearRefreshCookie(c)
	c.Status(http.StatusNoContent)
}

func (h *AuthHandler) LogoutAll(c *gin.Context) {
	if err := h.Svc.LogoutAll(c.Request.Context(), middleware.UserID(c)); err != nil {
		c.Error(err)
		return
	}
	h.clearRefreshCookie(c)
	c.Status(http.StatusNoContent)
}

func (h *AuthHandler) ChangePassword(c *gin.Context) {
	var in changePasswordRequest
	if !bindJSON(c, &in) {
		return
	}
	if err := h.Svc.ChangePassword(c.Request.Context(), middleware.UserID(c), in.CurrentPassword, in.NewPassword); err != nil {
		c.Error(err)
		return
	}
	h.clearRefreshCookie(c) // every session was revoked, so this device must log in again
	c.Status(http.StatusNoContent)
}

func (h *AuthHandler) Me(c *gin.Context) {
	user, err := h.Svc.GetUser(c.Request.Context(), middleware.UserID(c))
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *AuthHandler) setRefreshCookie(c *gin.Context, pair *services.TokenPair) {
	c.SetSameSite(http.SameSiteStrictMode)
	maxAge := int(time.Until(pair.RefreshExpiresAt).Seconds())
	// name, value, maxAge, path, domain, secure, httpOnly
	c.SetCookie(refreshCookieName, pair.RefreshToken, maxAge, refreshCookiePath, h.Cfg.CookieDomain, h.Cfg.CookieSecure, true)
}

func (h *AuthHandler) clearRefreshCookie(c *gin.Context) {
	c.SetSameSite(http.SameSiteStrictMode)
	c.SetCookie(refreshCookieName, "", -1, refreshCookiePath, h.Cfg.CookieDomain, h.Cfg.CookieSecure, true)
}

func tokenResponse(pair *services.TokenPair) gin.H {
	return gin.H{
		"access_token": pair.AccessToken,
		"token_type":   "Bearer",
		"expires_in":   int(time.Until(pair.AccessExpiresAt).Seconds()),
	}
}

func clientMeta(c *gin.Context) services.ClientMeta {
	return services.ClientMeta{IP: c.ClientIP(), UserAgent: c.Request.UserAgent()}
}
```

About the cookie flags:

- `HttpOnly`: JavaScript on your page cannot read the cookie, so a cross site scripting (XSS) bug cannot steal the refresh token.
- `Secure`: the browser only sends it over HTTPS.
- `SameSite=Strict`: the browser does not attach it to requests that start from another site, which stops CSRF.
- `Path=/api/v1/auth`: the cookie is not even sent to your normal API routes.

Your frontend keeps the **access token in memory** (a JavaScript variable), sends it in the `Authorization` header, and calls `/auth/refresh` when it expires or when the page reloads. Do not put the access token in `localStorage` if you can avoid it.

If your frontend and API live on different sites (not just different subdomains), you would need `SameSite=None` plus extra CSRF protection. Keep them on the same site when you can. A mobile app has no cookies, so it would send the refresh token in a request body and keep it in the secure storage of the phone.

### Step 5: Authentication and role middleware

Create `internal/middleware/auth.go`:

```go
package middleware

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/yourname/gin-bookstore/internal/apperr"
	"github.com/yourname/gin-bookstore/internal/auth"
	"github.com/yourname/gin-bookstore/internal/models"
)

const (
	ctxUserID = "userID"
	ctxRole   = "role"
)

// Auth checks the access token and stores the caller's id and role in the context.
func Auth(tokens *auth.TokenManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		scheme, tokenStr, found := strings.Cut(c.GetHeader("Authorization"), " ")
		if !found || !strings.EqualFold(scheme, "Bearer") || strings.TrimSpace(tokenStr) == "" {
			deny(c, "missing or malformed Authorization header")
			return
		}

		claims, err := tokens.ParseAccessToken(strings.TrimSpace(tokenStr))
		if err != nil {
			deny(c, "invalid or expired access token")
			return
		}
		id, err := strconv.ParseUint(claims.Subject, 10, 64)
		if err != nil {
			deny(c, "invalid access token")
			return
		}

		c.Set(ctxUserID, uint(id))
		c.Set(ctxRole, claims.Role)
		c.Next()
	}
}

func deny(c *gin.Context, msg string) {
	c.Header("WWW-Authenticate", `Bearer realm="api"`)
	c.Error(apperr.Unauthorized(msg))
	c.Abort()
}

// RequireRole allows the request only if the caller has one of the given roles.
// Always place it after Auth.
func RequireRole(roles ...string) gin.HandlerFunc {
	allowed := make(map[string]struct{}, len(roles))
	for _, r := range roles {
		allowed[r] = struct{}{}
	}
	return func(c *gin.Context) {
		if _, ok := allowed[c.GetString(ctxRole)]; !ok {
			c.Error(apperr.Forbidden("you do not have permission to do this"))
			c.Abort()
			return
		}
		c.Next()
	}
}

// UserID returns the id of the authenticated caller. Only call it behind Auth.
func UserID(c *gin.Context) uint {
	return c.MustGet(ctxUserID).(uint)
}

func IsAdmin(c *gin.Context) bool {
	return c.GetString(ctxRole) == models.RoleAdmin
}

// NoStore stops browsers and proxies from caching responses that contain tokens.
func NoStore() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Cache-Control", "no-store")
		c.Header("Pragma", "no-cache")
		c.Next()
	}
}
```

Two things to know:

- 401 means "I do not know who you are". 403 means "I know who you are, and you are not allowed". Use them correctly and your frontend can react correctly (refresh the token on 401, show a message on 403).
- The role lives inside the access token, so if you demote someone, the old role stays valid until that token expires (at most 15 minutes). For most apps that is fine. For very sensitive actions, load the user from the database inside the handler.

### Step 6: Ownership checks

Authentication answers "who are you". Authorization answers "may you do this to **this** thing". The most common real world API bug is skipping the second question: any logged in user can edit any other user's data by changing the id in the URL. Security people call it BOLA or IDOR, and it tops the OWASP API risk list.

Update `internal/handlers/book.go`. Add `"github.com/yourname/gin-bookstore/internal/middleware"` to the imports, and replace `Create`, `Update`, and `Delete`:

```go
// canModify allows the owner of a book, or an admin.
func canModify(c *gin.Context, book *models.Book) bool {
	return middleware.IsAdmin(c) || book.OwnerID == middleware.UserID(c)
}

func (h *BookHandler) Create(c *gin.Context) {
	var in BookInput
	if !bindJSON(c, &in) {
		return
	}

	book := models.Book{
		Title:   in.Title,
		Author:  in.Author,
		Price:   in.Price,
		OwnerID: middleware.UserID(c), // from the verified token, never from the request body
	}
	if err := h.DB.Create(&book).Error; err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusCreated, book)
}

func (h *BookHandler) Update(c *gin.Context) {
	id, err := parseID(c)
	if err != nil {
		c.Error(err)
		return
	}

	var in BookInput
	if !bindJSON(c, &in) {
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
	if !canModify(c, &book) {
		c.Error(apperr.Forbidden("you can only change your own books"))
		return
	}

	book.Title, book.Author, book.Price = in.Title, in.Author, in.Price
	if err := h.DB.Save(&book).Error; err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, book)
}

func (h *BookHandler) Delete(c *gin.Context) {
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
	if !canModify(c, &book) {
		c.Error(apperr.Forbidden("you can only delete your own books"))
		return
	}

	if err := h.DB.Delete(&book).Error; err != nil {
		c.Error(err)
		return
	}
	c.Status(http.StatusNoContent)
}
```

Notice that `OwnerID` always comes from the token. The request body type `BookInput` has no owner field at all, and thanks to `DisallowUnknownFields`, a client that tries to send one gets a 400.

For data that must stay private (like a user's orders), return **404 instead of 403** when the caller is not the owner, so you do not even confirm the record exists. Books are public to read, so 403 is fine here.

### Step 7: Wire up the routes

The auth routes use `middleware.RateLimit`, which is explained in Lesson 15. It is a single file, so create `internal/middleware/ratelimit.go` with the code from that lesson now and come back here.

At the top of `routes.Setup`, build the pieces, then register the routes. Add imports for `log/slog`, `os`, `time`, `golang.org/x/time/rate`, and your `auth`, `models`, `services`, and `middleware` packages.

```go
func Setup(db *gorm.DB, cfg *config.Config) *gin.Engine {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery(), middleware.ErrorHandler())

	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	tokens := auth.NewTokenManager(cfg.JWTSecret, cfg.JWTIssuer, cfg.JWTAudience, cfg.AccessTTL)
	authSvc := services.NewAuthService(db, tokens, cfg, logger)

	authH := &handlers.AuthHandler{Svc: authSvc, Cfg: cfg}
	adminH := &handlers.AdminHandler{DB: db}
	books := &handlers.BookHandler{DB: db}

	requireLogin := middleware.Auth(tokens)

	api := r.Group("/api/v1")
	{
		// Auth routes: never cached, and limited per IP to slow down password guessing
		authGroup := api.Group("/auth",
			middleware.NoStore(),
			middleware.RateLimit(rate.Every(time.Minute/time.Duration(cfg.AuthRatePerMinute)), cfg.AuthRateBurst),
		)
		{
			authGroup.POST("/register", authH.Register)
			authGroup.POST("/login", authH.Login)
			authGroup.POST("/refresh", authH.Refresh)
			authGroup.POST("/logout", authH.Logout)
			authGroup.POST("/logout-all", requireLogin, authH.LogoutAll)
			authGroup.POST("/change-password", requireLogin, authH.ChangePassword)
		}

		// Public reads
		api.GET("/books", books.List)
		api.GET("/books/:id", books.Get)

		// Logged in users
		protected := api.Group("", requireLogin)
		{
			protected.GET("/me", authH.Me)
			protected.POST("/books", books.Create)
			protected.PUT("/books/:id", books.Update)
			protected.DELETE("/books/:id", books.Delete)
		}

		// Admins only
		admin := api.Group("/admin", requireLogin, middleware.RequireRole(models.RoleAdmin))
		{
			admin.GET("/users", adminH.ListUsers)
		}
	}
	return r
}
```

Here is a tiny admin handler so the admin group has something to protect. Create `internal/handlers/admin.go`:

```go
package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/yourname/gin-bookstore/internal/models"
)

type AdminHandler struct {
	DB *gorm.DB
}

func (h *AdminHandler) ListUsers(c *gin.Context) {
	var users []models.User
	if err := h.DB.Order("id").Limit(100).Find(&users).Error; err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, users)
}
```

### Step 8: First admin and housekeeping

Nobody can register as an admin through the API (that is the point). The usual solution is to create the first admin at startup from environment variables. Add these to `internal/database/database.go`, along with the imports `context`, `fmt`, `log`, `time`, and your `auth` package:

```go
func SeedAdmin(db *gorm.DB, email, password string, cost int) error {
	email = models.NormalizeEmail(email)

	var count int64
	if err := db.Model(&models.User{}).Where("email = ?", email).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return nil // already exists, nothing to do
	}

	if err := auth.ValidatePassword(password); err != nil {
		return fmt.Errorf("ADMIN_PASSWORD %w", err)
	}
	hash, err := auth.HashPassword(password, cost)
	if err != nil {
		return err
	}
	return db.Create(&models.User{
		Name:         "Administrator",
		Email:        email,
		PasswordHash: hash,
		Role:         models.RoleAdmin,
		IsActive:     true,
	}).Error
}

// RunTokenCleanup deletes expired refresh tokens so the table does not grow forever.
func RunTokenCleanup(ctx context.Context, db *gorm.DB, every time.Duration) {
	ticker := time.NewTicker(every)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			if err := db.Where("expires_at < ?", time.Now()).Delete(&models.RefreshToken{}).Error; err != nil {
				log.Printf("token cleanup failed: %v", err)
			}
		}
	}
}
```

You will call both from `main.go`. The final `main.go` is shown in Lesson 17. Remove `ADMIN_PASSWORD` from your environment after the first start if you like, because the account now exists.

### Try it with curl

Start the server, then:

```bash
# 1. register (a weak password is rejected with a 422)
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{"name":"Test User","email":"test@example.com","password":"a-long-passphrase-123"}'

# 2. login: -c saves the refresh cookie into cookies.txt, -i shows the headers
curl -i -c cookies.txt -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","password":"a-long-passphrase-123"}'

# 3. use the access token from the response
curl -X POST http://localhost:8080/api/v1/books \
  -H "Authorization: Bearer PASTE_ACCESS_TOKEN_HERE" \
  -H "Content-Type: application/json" \
  -d '{"title":"Go in Action","author":"William Kennedy","price":39.5}'

# 4. refresh: -b sends the cookie, -c stores the NEW cookie
cp cookies.txt old-cookies.txt
curl -b cookies.txt -c cookies.txt -X POST http://localhost:8080/api/v1/auth/refresh

# 5. replay the OLD cookie. This must fail with 401 and revoke the whole session
curl -b old-cookies.txt -X POST http://localhost:8080/api/v1/auth/refresh

# 6. the NEW cookie is dead too, because the family was revoked
curl -b cookies.txt -X POST http://localhost:8080/api/v1/auth/refresh
```

Also try these on purpose:

- Log in with a wrong password six times in a row and watch the account lock (429).
- Register a second user and try to `PUT` the first user's book. You should get a 403.
- Send `{"name":"x","email":"x@example.com","password":"a-long-passphrase-123","role":"admin"}` to register and see it rejected.
- Call `/api/v1/admin/users` as a normal user (403) and as the seeded admin (200).

In debug mode the cookie is not marked `Secure`, so curl over plain http works. In release mode, cookies are `Secure` by default, so set `COOKIE_SECURE=false` only for local http testing.

### Design decisions and trade-offs

Real engineering is choosing trade-offs on purpose. These are the ones in this design:

- **Access tokens cannot be revoked instantly.** After logout, an access token keeps working until it expires (15 minutes). That is the price of stateless verification. If you need instant revocation, keep a denylist of token ids (`jti`) in Redis and check it in the middleware.
- **Two simultaneous refresh calls with the same token look like theft.** The second one will revoke the session. Frontends should make sure only one refresh request is in flight at a time. Some systems add a short grace period instead, which is more forgiving but less strict.
- **Account lockout can be abused.** An attacker can lock out a victim by failing logins on purpose. The per IP rate limit and the short lock time (15 minutes) limit the damage. Notifying the user by email is the usual next step.
- **Registration tells you if an email is taken.** That is friendlier but allows email enumeration. Systems with high privacy needs always answer "check your inbox" and send an email either way.
- **HS256 with one shared secret** is simple and fine for a single service. When several services must verify tokens, switch to asymmetric signing (RS256 or EdDSA) so only the auth service holds the private key.

### Where to go from here

This is a solid foundation, not the whole story. A real product would add:

- Email verification and a password reset flow (single use, expiring tokens, sent by email)
- Multi factor authentication (TOTP apps or passkeys)
- Social login with OAuth 2.0 and OpenID Connect
- A "your sessions" page. You already store the user agent and IP for every refresh token, so listing and revoking sessions is just a query
- Audit logs in a table instead of only in log files
- Or, for many teams, an identity provider (Keycloak, Auth0, Firebase Auth, and others), so they do not maintain this code themselves. Building it once teaches you exactly what those providers do for you

### Recap

- Short lived JWT access token plus a rotating, hashed, opaque refresh token in an `HttpOnly` cookie
- Reuse of an old refresh token revokes the whole session family
- bcrypt, a password policy, same error and timing for unknown users, lockout, and per IP rate limiting
- Roles from the database, never from the client, and an ownership check on every write
- Business rules live in a service, HTTP details live in the handler

### Homework

Add `PATCH /admin/users/:id/active` that lets an admin deactivate or reactivate a user. Then write a test proving that a deactivated user can no longer refresh their token.

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

	if !canModify(c, &book) {
		c.Error(apperr.Forbidden("you can only change your own books"))
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

## Lesson 15: CORS, Rate Limiting, Compression, Security Headers, Logging

**You will learn:** the production middleware that every public API needs, and how to configure it for real.

### CORS

Browsers block JavaScript on `https://myapp.com` from calling an API on `https://api.example.com` unless the API says it is okay. That is CORS. Think of a building with a visitor policy: the front desk only lets in people from companies on the approved list. Without a CORS middleware, your React frontend will hit errors in the browser console even though curl works fine.

```bash
go get github.com/gin-contrib/cors
```

```go
r.Use(cors.New(cors.Config{
	AllowOrigins:     cfg.AllowedOrigins, // from ALLOWED_ORIGINS, for example https://myapp.com
	AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
	AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
	ExposeHeaders:    []string{"X-Request-ID"},
	AllowCredentials: true, // required so the browser sends the refresh token cookie
	MaxAge:           12 * time.Hour,
}))
```

List exact origins. A wildcard `*` together with credentials is not allowed by browsers, and allowing every origin on an authenticated API is a security smell anyway. On the frontend, remember to send requests with credentials (`credentials: "include"` in `fetch`) when calling the auth endpoints.

### Rate limiting

Stop one client from hammering your API. A per client limiter using the standard `x/time/rate` package, this time with cleanup so memory does not grow forever:

```bash
go get golang.org/x/time
```

```go
package middleware

import (
	"math"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

type visitor struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

// RateLimit allows `limit` requests per second per client IP, with bursts up to `burst`.
func RateLimit(limit rate.Limit, burst int) gin.HandlerFunc {
	var (
		mu          sync.Mutex
		visitors    = make(map[string]*visitor)
		lastCleanup = time.Now()
	)
	retryAfter := strconv.Itoa(int(math.Ceil(1 / float64(limit))))

	return func(c *gin.Context) {
		ip := c.ClientIP()
		now := time.Now()

		mu.Lock()
		// Once a minute, forget clients that have been quiet for 10 minutes
		if now.Sub(lastCleanup) > time.Minute {
			for key, v := range visitors {
				if now.Sub(v.lastSeen) > 10*time.Minute {
					delete(visitors, key)
				}
			}
			lastCleanup = now
		}
		v, ok := visitors[ip]
		if !ok {
			v = &visitor{limiter: rate.NewLimiter(limit, burst)}
			visitors[ip] = v
		}
		v.lastSeen = now
		allowed := v.limiter.Allow()
		mu.Unlock()

		if !allowed {
			c.Header("Retry-After", retryAfter)
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

You now use it twice in the project: a generous limit for the whole API, and a strict one on `/auth` routes (Lesson 12). This version keeps its counters in the memory of one process. When you run several instances behind a load balancer, each has its own counters, so use a shared store like Redis for a limit that holds across all of them.

### Trusted proxies

`c.ClientIP()` reads headers like `X-Forwarded-For`, which any client can fake. Tell Gin which proxies to trust:

```go
_ = r.SetTrustedProxies(nil)                       // trust nobody, use the direct connection IP
// or, when behind your own reverse proxy:
// _ = r.SetTrustedProxies([]string{"10.0.0.1"})
```

If you skip this, your rate limiter (and your login lockout logs) can be fooled by a fake header. And if you deploy behind Nginx, Caddy, or a cloud load balancer while trusting nobody, every request will appear to come from the proxy's IP and share one limit. Configure it for your real setup.

### Compression

```bash
go get github.com/gin-contrib/gzip
```

```go
r.Use(gzip.Gzip(gzip.DefaultCompression))
```

### Security headers

A few response headers switch on protections in the browser. They cost nothing:

```go
func SecurityHeaders(hsts bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		h := c.Writer.Header()
		h.Set("X-Content-Type-Options", "nosniff")  // do not guess content types
		h.Set("X-Frame-Options", "DENY")            // cannot be embedded in an iframe
		h.Set("Referrer-Policy", "no-referrer")     // do not leak URLs to other sites
		if hsts {
			// tell browsers to use HTTPS only, for two years
			h.Set("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
		}
		c.Next()
	}
}
```

Turn HSTS on only in production, over HTTPS. A pure JSON API can also send `Content-Security-Policy: default-src 'none'`, but skip it on routes that serve HTML or Swagger UI, since those need their own policy.

### Request ID and structured logging

A request ID lets you find all logs that belong to one request. Combine it with Go's `log/slog` package for structured logs:

```go
func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.GetHeader("X-Request-ID")
		if id == "" || len(id) > 64 {
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
		path := c.Request.URL.Path // not the full URL, so query strings with secrets stay out of logs

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

A rule worth tattooing somewhere: **never log passwords, tokens, or full request bodies.** Log ids and outcomes instead.

### Putting it together

The order of middleware matters. Here is a sensible final start of `routes.Setup`:

```go
func Setup(db *gorm.DB, cfg *config.Config) *gin.Engine {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	r := gin.New()
	_ = r.SetTrustedProxies(nil)
	r.MaxMultipartMemory = 8 << 20

	r.Use(
		middleware.RequestID(),                                  // first, so everything else can use the id
		middleware.RequestLogger(logger),                        // wraps everything, logs after it finishes
		gin.Recovery(),                                          // catches panics from everything below
		middleware.ErrorHandler(),                               // turns c.Error values into JSON
		middleware.SecurityHeaders(cfg.GinMode == gin.ReleaseMode),
		cors.New(cors.Config{
			AllowOrigins:     cfg.AllowedOrigins,
			AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
			AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
			ExposeHeaders:    []string{"X-Request-ID"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		}),
		gzip.Gzip(gzip.DefaultCompression),
		middleware.RateLimit(20, 40),                            // 20 requests per second per IP, bursts up to 40
		middleware.LimitBody(10<<20),                            // 10 MB request cap
	)

	// ... services and routes from Lesson 12
	return r
}
```

### Recap

- CORS lets your browser frontend talk to the API, with credentials only for origins you list
- Rate limit by client IP (with cleanup), and configure trusted proxies correctly
- Security headers are free protection
- Request IDs and structured logs make debugging production possible, as long as secrets stay out of them
- Middleware order is part of your design

### Homework

Write a test that creates a router with `AuthRateBurst: 2` and proves that the third quick login attempt from one IP gets a 429 with a `Retry-After` header.

***

## Lesson 16: Testing

**You will learn:** how to test your API, including the security behavior, without starting a real server.

Because `routes.Setup` returns a `*gin.Engine`, and an engine implements `http.Handler`, you can call it directly with fake requests using the standard `httptest` package. No network, no ports.

Security code deserves the strictest tests in your project, because bugs there do not show up as errors, they show up as breaches. So we will test the rules from Lesson 12: weak passwords, generic login errors, lockout, ownership, and refresh token rotation.

Create `internal/routes/routes_test.go`:

```go
package routes_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"github.com/yourname/gin-bookstore/internal/config"
	"github.com/yourname/gin-bookstore/internal/database"
	"github.com/yourname/gin-bookstore/internal/routes"
	"github.com/yourname/gin-bookstore/internal/validation"
)

const testPassword = "a-long-passphrase-123"

func testConfig() *config.Config {
	return &config.Config{
		GinMode:           gin.TestMode,
		JWTSecret:         "test-secret-that-is-at-least-32-characters-long",
		JWTIssuer:         "test",
		JWTAudience:       "test-api",
		AccessTTL:         15 * time.Minute,
		RefreshTTL:        24 * time.Hour,
		BcryptCost:        bcrypt.MinCost, // fast hashing keeps tests quick
		AllowedOrigins:    []string{"http://localhost:3000"},
		AuthRatePerMinute: 600,
		AuthRateBurst:     100,
	}
}

func newTestRouter(t *testing.T) *gin.Engine {
	t.Helper()
	gin.SetMode(gin.TestMode)
	validation.Setup()

	// A separate in-memory database per test so tests do not affect each other
	name := strings.ReplaceAll(t.Name(), "/", "_")
	db, err := database.Connect("file:" + name + "?mode=memory&cache=shared")
	if err != nil {
		t.Fatalf("connect db: %v", err)
	}
	return routes.Setup(db, testConfig())
}

// call sends a request to the router. Options add headers or cookies.
func call(r http.Handler, method, path string, body any, opts ...func(*http.Request)) *httptest.ResponseRecorder {
	var buf bytes.Buffer
	if body != nil {
		_ = json.NewEncoder(&buf).Encode(body)
	}

	req := httptest.NewRequest(method, path, &buf)
	req.Header.Set("Content-Type", "application/json")
	for _, opt := range opts {
		opt(req)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func withToken(token string) func(*http.Request) {
	return func(req *http.Request) { req.Header.Set("Authorization", "Bearer "+token) }
}

func withCookie(c *http.Cookie) func(*http.Request) {
	return func(req *http.Request) { req.AddCookie(c) }
}

func refreshCookie(t *testing.T, w *httptest.ResponseRecorder) *http.Cookie {
	t.Helper()
	for _, c := range w.Result().Cookies() {
		if c.Name == "refresh_token" && c.Value != "" {
			return c
		}
	}
	t.Fatalf("no refresh_token cookie in response")
	return nil
}

type session struct {
	access  string
	refresh *http.Cookie
}

func registerAndLogin(t *testing.T, r http.Handler, email string) session {
	t.Helper()

	w := call(r, http.MethodPost, "/api/v1/auth/register", gin.H{
		"name": "Test User", "email": email, "password": testPassword,
	})
	if w.Code != http.StatusCreated {
		t.Fatalf("register: expected 201, got %d (%s)", w.Code, w.Body.String())
	}

	w = call(r, http.MethodPost, "/api/v1/auth/login", gin.H{"email": email, "password": testPassword})
	if w.Code != http.StatusOK {
		t.Fatalf("login: expected 200, got %d (%s)", w.Code, w.Body.String())
	}

	var body struct {
		AccessToken string `json:"access_token"`
	}
	if err := json.Unmarshal(w.Body.Bytes(), &body); err != nil || body.AccessToken == "" {
		t.Fatalf("login: no access token in response (%s)", w.Body.String())
	}
	return session{access: body.AccessToken, refresh: refreshCookie(t, w)}
}

func TestHealth(t *testing.T) {
	r := newTestRouter(t)

	w := call(r, http.MethodGet, "/healthz", nil)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
}

func TestProtectedRouteNeedsToken(t *testing.T) {
	r := newTestRouter(t)

	w := call(r, http.MethodPost, "/api/v1/books", gin.H{"title": "Go", "author": "Someone", "price": 10})
	if w.Code != http.StatusUnauthorized {
		t.Fatalf("no token: expected 401, got %d", w.Code)
	}

	w = call(r, http.MethodPost, "/api/v1/books", gin.H{"title": "Go", "author": "Someone", "price": 10}, withToken("not-a-real-token"))
	if w.Code != http.StatusUnauthorized {
		t.Fatalf("garbage token: expected 401, got %d", w.Code)
	}
}

func TestWeakPasswordIsRejected(t *testing.T) {
	r := newTestRouter(t)

	w := call(r, http.MethodPost, "/api/v1/auth/register", gin.H{
		"name": "Test User", "email": "weak@example.com", "password": "short",
	})
	if w.Code != http.StatusUnprocessableEntity {
		t.Fatalf("expected 422, got %d (%s)", w.Code, w.Body.String())
	}
}

func TestRoleCannotBeSetByClient(t *testing.T) {
	r := newTestRouter(t)

	w := call(r, http.MethodPost, "/api/v1/auth/register", gin.H{
		"name": "Sneaky", "email": "sneaky@example.com", "password": testPassword, "role": "admin",
	})
	if w.Code != http.StatusBadRequest {
		t.Fatalf("unknown field: expected 400, got %d (%s)", w.Code, w.Body.String())
	}
}

func TestLoginErrorsDoNotRevealWhichEmailsExist(t *testing.T) {
	r := newTestRouter(t)
	registerAndLogin(t, r, "real@example.com")

	wrongPassword := call(r, http.MethodPost, "/api/v1/auth/login", gin.H{"email": "real@example.com", "password": "definitely-wrong-123"})
	unknownEmail := call(r, http.MethodPost, "/api/v1/auth/login", gin.H{"email": "nobody@example.com", "password": "definitely-wrong-123"})

	if wrongPassword.Code != http.StatusUnauthorized || unknownEmail.Code != http.StatusUnauthorized {
		t.Fatalf("expected two 401s, got %d and %d", wrongPassword.Code, unknownEmail.Code)
	}
	if wrongPassword.Body.String() != unknownEmail.Body.String() {
		t.Fatalf("responses differ, which leaks information:\n%s\n%s", wrongPassword.Body.String(), unknownEmail.Body.String())
	}
}

func TestAccountLocksAfterRepeatedFailures(t *testing.T) {
	r := newTestRouter(t)
	registerAndLogin(t, r, "lock@example.com")

	for i := 0; i < 5; i++ {
		call(r, http.MethodPost, "/api/v1/auth/login", gin.H{"email": "lock@example.com", "password": "wrong-password-123"})
	}

	// Even the correct password is refused while the account is locked
	w := call(r, http.MethodPost, "/api/v1/auth/login", gin.H{"email": "lock@example.com", "password": testPassword})
	if w.Code != http.StatusTooManyRequests {
		t.Fatalf("expected 429 while locked, got %d (%s)", w.Code, w.Body.String())
	}
}

func TestOnlyTheOwnerCanChangeABook(t *testing.T) {
	r := newTestRouter(t)
	alice := registerAndLogin(t, r, "alice@example.com")
	bob := registerAndLogin(t, r, "bob@example.com")

	w := call(r, http.MethodPost, "/api/v1/books", gin.H{"title": "Go in Action", "author": "William Kennedy", "price": 39.5}, withToken(alice.access))
	if w.Code != http.StatusCreated {
		t.Fatalf("create: expected 201, got %d (%s)", w.Code, w.Body.String())
	}
	var book struct {
		ID      uint `json:"id"`
		OwnerID uint `json:"owner_id"`
	}
	_ = json.Unmarshal(w.Body.Bytes(), &book)
	path := "/api/v1/books/" + strconv.FormatUint(uint64(book.ID), 10)
	update := gin.H{"title": "Changed", "author": "William Kennedy", "price": 10}

	if w := call(r, http.MethodPut, path, update, withToken(bob.access)); w.Code != http.StatusForbidden {
		t.Fatalf("other user: expected 403, got %d", w.Code)
	}
	if w := call(r, http.MethodPut, path, update, withToken(alice.access)); w.Code != http.StatusOK {
		t.Fatalf("owner: expected 200, got %d (%s)", w.Code, w.Body.String())
	}
	if w := call(r, http.MethodDelete, path, nil, withToken(bob.access)); w.Code != http.StatusForbidden {
		t.Fatalf("other user delete: expected 403, got %d", w.Code)
	}
}

func TestRefreshRotationAndReuseDetection(t *testing.T) {
	r := newTestRouter(t)
	s := registerAndLogin(t, r, "rotate@example.com")

	// 1. A normal refresh works and hands out a different token
	w := call(r, http.MethodPost, "/api/v1/auth/refresh", nil, withCookie(s.refresh))
	if w.Code != http.StatusOK {
		t.Fatalf("refresh: expected 200, got %d (%s)", w.Code, w.Body.String())
	}
	newCookie := refreshCookie(t, w)
	if newCookie.Value == s.refresh.Value {
		t.Fatal("refresh token was not rotated")
	}

	// 2. Replaying the OLD token looks like theft
	w = call(r, http.MethodPost, "/api/v1/auth/refresh", nil, withCookie(s.refresh))
	if w.Code != http.StatusUnauthorized {
		t.Fatalf("reused token: expected 401, got %d", w.Code)
	}

	// 3. So the whole family is revoked, and even the NEW token stops working
	w = call(r, http.MethodPost, "/api/v1/auth/refresh", nil, withCookie(newCookie))
	if w.Code != http.StatusUnauthorized {
		t.Fatalf("family should be revoked: expected 401, got %d", w.Code)
	}
}

func TestLogoutRevokesTheRefreshToken(t *testing.T) {
	r := newTestRouter(t)
	s := registerAndLogin(t, r, "logout@example.com")

	if w := call(r, http.MethodPost, "/api/v1/auth/logout", nil, withCookie(s.refresh)); w.Code != http.StatusNoContent {
		t.Fatalf("logout: expected 204, got %d", w.Code)
	}
	if w := call(r, http.MethodPost, "/api/v1/auth/refresh", nil, withCookie(s.refresh)); w.Code != http.StatusUnauthorized {
		t.Fatalf("refresh after logout: expected 401, got %d", w.Code)
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
- Use bcrypt's minimum cost in tests (`bcrypt.MinCost`). The production cost of 12 would make a test suite crawl.
- Put your config in a `testConfig()` function. When you add a setting, tests keep compiling and you update one place.
- Use table driven tests for validation cases: one slice of inputs and expected status codes, one loop.
- Use a fresh in-memory database per test so tests stay independent and can run in any order.
- Test the unhappy paths (bad input, missing token, wrong owner, reused token) at least as much as the happy path. That is where bugs and vulnerabilities live.

### Recap

- `httptest.NewRecorder` plus `r.ServeHTTP` tests the full stack without a network
- Response cookies can be read with `w.Result().Cookies()` and replayed with `req.AddCookie`
- Security rules are behavior, so write tests that try to break them

### Homework

Write a table driven test that sends 6 different bad payloads to `POST /books` and checks that each returns 422. Then add a test that changing a password signs the user out of every session.

***

## Lesson 17: Graceful Shutdown and Server Settings

**You will learn:** how to stop your server without cutting off in-flight requests.

`r.Run()` is convenient, but it gives you no control. When you deploy a new version, the old process gets a stop signal. With `r.Run()` it dies immediately and any request in progress is dropped. A graceful shutdown stops accepting new requests, lets current ones finish, and then exits.

It is like a shop at closing time: lock the front door so no new customers enter, but serve everyone already inside before turning off the lights.

We will use the standard `http.Server`, which also lets us set timeouts. Update `main.go`. This is the final shape, and it also seeds the first admin and starts the token cleanup job from Lesson 12:

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

	// Create the first admin if ADMIN_EMAIL and ADMIN_PASSWORD are set
	if cfg.AdminEmail != "" && cfg.AdminPassword != "" {
		if err := database.SeedAdmin(db, cfg.AdminEmail, cfg.AdminPassword, cfg.BcryptCost); err != nil {
			log.Fatalf("seed admin: %v", err)
		}
	}

	// Cancelled on Ctrl+C or a termination signal from Docker or Kubernetes
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// Background job: remove expired refresh tokens every hour
	go database.RunTokenCleanup(ctx, db, time.Hour)

	srv := &http.Server{
		Addr:              ":" + cfg.Port,
		Handler:           routes.Setup(db, cfg),
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	// Start the server in a goroutine so main can wait for the signal
	go func() {
		log.Printf("listening on %s", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("server error: %v", err)
		}
	}()

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
docker run -p 8080:8080 \
  -e JWT_SECRET=paste-a-random-secret-of-at-least-32-characters \
  -e ALLOWED_ORIGINS=https://myapp.com \
  -e DB_PATH=/data/bookstore.db \
  -v bookstore-data:/data \
  gin-bookstore
```

The image runs in release mode, so cookies are marked `Secure` and browsers only send them over HTTPS. For a quick local test over plain http you can add `-e COOKIE_SECURE=false`, but never do that in production.

The `-v` option mounts a named volume so your SQLite file survives container restarts. Without it, your data disappears with the container.

### Production checklist

- `GIN_MODE=release` (hides debug output and route listing, slightly faster)
- Secrets come from environment variables or a secret manager, never from git. `JWT_SECRET` must be at least 32 random characters (for example from `openssl rand -base64 48`), and you should have a plan to rotate it
- Cookies are `Secure`, `HttpOnly`, and `SameSite=Strict` (the defaults in release mode), and `ALLOWED_ORIGINS` lists only your real frontends
- The first admin is created from `ADMIN_EMAIL` and `ADMIN_PASSWORD`, and nobody can self register as an admin
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

Auth and security

- Short lived access token in the `Authorization` header, rotating refresh token in an `HttpOnly` cookie
- `middleware.Auth` identifies the caller, `middleware.RequireRole` checks the role, and an ownership check protects each resource
- `bcrypt` for passwords, SHA-256 for random refresh tokens, identical responses for unknown users and wrong passwords

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
13. **One long lived JWT and no way to revoke it.** Use short access tokens plus rotating refresh tokens.
14. **Storing refresh tokens in plain text, or keeping them in `localStorage`.** Store hashes on the server and use an `HttpOnly` cookie in the browser.
15. **Trusting the client for ids and roles.** Take the user id and role from the verified token, and check ownership on every write.
16. **Different errors for "no such user" and "wrong password".** It tells attackers which emails are registered.
17. **Accepting any JSON fields.** Reject unknown fields so nobody can sneak in `"role": "admin"`.
18. **Weak or default JWT secrets.** Fail at startup if the secret is short.

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
