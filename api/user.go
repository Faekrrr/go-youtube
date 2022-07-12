// api/artist.go
package api
import (
 "net/http"
"github.com/gin-gonic/gin"
 "github.com/jinzhu/gorm"
)


// User model
type User struct {
 ID    uint   `json:"id" gorm:"primary_key"`
 Name  string `json:"name"`
 Email string `json:"email" gorm:"unique;not null"`
}


// CreateUserInput : struct for create user post request
type CreateUserInput struct {
 Name  string `json:"name" binding:"required"`
 Email string `json:"email" binding:"required"`
}


// Findcusers : Controller for getting all users
func FindUsers(c *gin.Context) {
 db := c.MustGet("db").(*gorm.DB)
var users []User
db.Find(&users)
c.JSON(http.StatusOK, gin.H{"data": users})
}


// CreateUser : controller for creating new users
func CreateUser(context *gin.Context) {
 db := context.MustGet("db").(*gorm.DB)

// Validate input
 var input CreateUserInput
 if err := context.ShouldBindJSON(&input); err != nil {
	context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
  return
 }

// Create User
 user := User{Name: input.Name, Email: input.Email}
 db.Create(&user)
 context.JSON(http.StatusOK, gin.H{"data": user})
}