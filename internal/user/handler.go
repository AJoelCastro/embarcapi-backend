package user

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var (
	users  = make(map[int]User)
	nextID = 1
	mu     sync.Mutex
)

func RegisterRoutes(r *gin.Engine) {
	group := r.Group("/users")
	{
		group.GET("/", listUsers)
		group.POST("/", createUser)
		group.GET(":id", getUser)
		group.PUT(":id", updateUser)
		group.DELETE(":id", deleteUser)
	}
}

func listUsers(c *gin.Context) {
	mu.Lock()
	defer mu.Unlock()
	result := make([]User, 0, len(users))
	for _, u := range users {
		result = append(result, u)
	}
	c.JSON(http.StatusOK, result)
}

func createUser(c *gin.Context) {
	var u User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	mu.Lock()
	u.ID = nextID
	nextID++
	users[u.ID] = u
	mu.Unlock()
	c.JSON(http.StatusCreated, u)
}

func getUser(c *gin.Context) {
	id, ok := getIDParam(c)
	if !ok {
		return
	}
	mu.Lock()
	u, exists := users[id]
	mu.Unlock()
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	c.JSON(http.StatusOK, u)
}

func updateUser(c *gin.Context) {
	id, ok := getIDParam(c)
	if !ok {
		return
	}
	var u User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	mu.Lock()
	_, exists := users[id]
	if !exists {
		mu.Unlock()
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	u.ID = id
	users[id] = u
	mu.Unlock()
	c.JSON(http.StatusOK, u)
}

func deleteUser(c *gin.Context) {
	id, ok := getIDParam(c)
	if !ok {
		return
	}
	mu.Lock()
	_, exists := users[id]
	if !exists {
		mu.Unlock()
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	delete(users, id)
	mu.Unlock()
	c.Status(http.StatusNoContent)
}

func getIDParam(c *gin.Context) (int, bool) {
	idStr := c.Param("id")
	var id int
	_, err := fmt.Sscanf(idStr, "%d", &id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return 0, false
	}
	return id, true
}
