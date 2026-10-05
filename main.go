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
