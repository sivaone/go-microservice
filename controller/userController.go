package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"

	"go-microservice/ent"
	"go-microservice/service"
	"go-microservice/utils"
)

func UserGetAllController(c *gin.Context) {
	w := c.Writer
	users, err := service.NewUserOps(c.Request.Context()).UsersGetAll()
	if err != nil {
		log.Printf("error retrieving users: %v", err)
		utils.Return(w, false, http.StatusInternalServerError, err, nil)
	}
	log.Println("Successfully retrieved users")
	utils.Return(w, true, http.StatusOK, nil, users)
}

func UserCreateController(c *gin.Context) {
	w := c.Writer
	var newUser ent.User
	err := json.NewDecoder(c.Request.Body).Decode(&newUser)
	if err != nil {
		log.Printf("error converting json to object: %v", err)
		utils.Return(w, false, http.StatusBadRequest, err, nil)
	}

	user, err := service.NewUserOps(c.Request.Context()).CreateUser(newUser)
	if err != nil {
		log.Printf("error saving user to database: %v", err)
		utils.Return(w, false, http.StatusInternalServerError, err, nil)
	}

	log.Println("successfully created new user")
	utils.Return(w, true, http.StatusOK, nil, user)
}
