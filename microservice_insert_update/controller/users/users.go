package user

import (
	"errors"
	"net/http"
	"microservice_insert_update/conn"
	user "microservice_insert_update/models/user"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

// UserCollection statically declared
const UserCollection = "user"

var (
	errNotExist        = errors.New("Users are not exist")
	errInvalidID       = errors.New("Invalid ID")
	errInvalidBody     = errors.New("Invalid request body")
	errInsertionFailed = errors.New("Error in the user insertion")
	errUpdationFailed  = errors.New("Error in the user updation")
	errDeletionFailed  = errors.New("Error in the user deletion")
)

// CreateUser Endpoint
func CreateUser(c *gin.Context) {
	// Get DB from Mongo Config
	db := conn.GetMongoDB()
	user := user.User{}
	err := c.Bind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errInvalidBody.Error()})
		return
	}
	user.ID = bson.NewObjectId()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	err = db.C(UserCollection).Insert(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errInsertionFailed.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "user": &user})
}

// UpdateUser Endpoint
func UpdateUser(c *gin.Context) {
	// Get DB from Mongo Config
	db := conn.GetMongoDB()
	var id bson.ObjectId = bson.ObjectIdHex(c.Param("id")) // Get Param
	existingUser, err := user.UserInfo(id, UserCollection)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errInvalidID.Error()})
		return
	}
	// user := user.User{}
	err = c.Bind(&existingUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errInvalidBody.Error()})
		return
	}
	existingUser.ID = id
	existingUser.UpdatedAt = time.Now()
	err = db.C(UserCollection).Update(bson.M{"_id": &id}, existingUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errUpdationFailed.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "user": &existingUser})
}
