package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/stheven26/technical_test/db"
	"github.com/stheven26/technical_test/globals"
	"github.com/stheven26/technical_test/models"
)

func Register(c *gin.Context) {
	var data map[string]string

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"time":    time.Now(),
		})
	}

	user := models.User{
		Phone: data["phone"],
	}

	connection := db.GetConnection()
	connection.Create(&user)

	c.JSON(http.StatusOK, gin.H{
		"message": "Succesfully Register",
		"time":    time.Now(),
	})
}

func Login(c *gin.Context) {
	var data map[string]string
	var user models.User

	connection := db.GetConnection()
	connection.Where(`phone = ?`, data["phone"]).First(&user)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "User not found",
			"time":    time.Now(),
		})
	} else {
		claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
			Issuer:    strconv.Itoa(int(user.ID)),
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		})

		token, err := claims.SignedString([]byte(globals.Key))

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
				"time":    time.Now(),
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Successfully Login!",
			"token":   token,
			"time":    time.Now(),
		})
	}
}

func OneOnOneMessage(c *gin.Context) {
	var data map[string]string

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"time":    time.Now(),
		})
	}

	msg := models.Message{
		Message: data["message"],
		Status:  data["status"],
		Sharing: data["sharing"],
	}

	connection := db.GetConnection()
	connection.Create(&msg)

	if msg.Status == "sent" || msg.Status == "Sent" {
		c.JSON(http.StatusOK, gin.H{
			"message": "successfully sent message",
			"time":    time.Now(),
		})
	}
	if msg.Status == "pending" || msg.Status == "Pending" {
		c.JSON(http.StatusOK, gin.H{
			"message": "pending sending message",
			"time":    time.Now(),
		})
	}
	if msg.Status == "read" || msg.Status == "Read" {
		c.JSON(http.StatusOK, gin.H{
			"message": "read message",
			"time":    time.Now(),
		})
	}

}

func CreateGroup(c *gin.Context) {
	var data map[string]string

	connection := db.GetConnection()

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"time":    time.Now(),
		})
	}

	group := models.Group{
		GroupName: data["group_name"],
		Role:      data["role"],
	}

	if group.Role != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Anda Bukan Admin, Anda Tidak Bisa Membuat Group",
			"time":    time.Now(),
		})
	} else {
		connection.Create(&group)

		c.JSON(http.StatusOK, gin.H{
			"message": "Success Create Group",
			"time":    time.Now(),
		})
	}
}

// func AddMember(c *gin.Context) {
// 	var data map[string]string
// 	var user models.User

// 	connection := db.GetConnection()

// 	connection.Where(`phone = ?, role = ? `, data["phone"], data["role"]).First(&user)

// 	if user.ID == 0 {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"message": "User not found",
// 			"time":    time.Now(),
// 		})
// 	} else {

// 	}

// }

// func SentMessageOnGroup(c *gin.Context) {
// 	var data map[string]string
// 	var user models.Group

// 	connection := db.GetConnection()

// 	connection.Where(`group_name = ?`, data["group_name"]).First(&user)

// 	if user.GroupName == "" {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"message": "Group not found",
// 			"time":    time.Now(),
// 		})
// 	} else {
// 		msg := models.MessageOnGroup{
// 			Message: data["message"],
// 			Status:  data["status"],
// 			Sharing: data["sharing"],
// 		}

// 		connection.Create(&msg)

// 		if msg.Status == "sent" || msg.Status == "Sent" {
// 			c.JSON(http.StatusOK, gin.H{
// 				"message": "successfully sent message on group",
// 				"time":    time.Now(),
// 			})
// 		}
// 		if msg.Status == "pending" || msg.Status == "Pending" {
// 			c.JSON(http.StatusOK, gin.H{
// 				"message": "pending sending message",
// 				"time":    time.Now(),
// 			})
// 		}
// 		if msg.Status == "read" || msg.Status == "Read" {
// 			c.JSON(http.StatusOK, gin.H{
// 				"message": "read message",
// 				"time":    time.Now(),
// 			})
// 		}
// 	}
// }
