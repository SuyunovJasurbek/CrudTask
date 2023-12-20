package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/SuyunovJasurbek/CrudTask/helper"
	"github.com/SuyunovJasurbek/CrudTask/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateUser(c *gin.Context) {
	var user models.UserHandler

	// step 1: validate input
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: "invalid input body"})
		return
	}

	// step 2: user fields validation
	if err, ok := helper.UserFeildCheck(user); !ok {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	// step 3: create user
	id, err := h.service.CreateUser(c, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Error{Message: "user not created"})
		return
	}

	// step 4: return user id
	c.JSON(http.StatusCreated, models.Response{
		Message: strconv.Itoa(id),
	})
}

func (h *Handler) UpdateUserById(c *gin.Context) {

	// step 1: validate input
	var user models.UserHandler
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, models.Error{Message: "invalid input id"})
		return
	}

	Id, err := strconv.Atoi(id)
	fmt.Println(Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: "invalid input id"})
		return
	}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: "invalid input body"})
		return
	}

	// step 2: user fields validation
	if err, ok := helper.UserFeildCheck(user); !ok {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	// step 3: update user
	updated_at, err := h.service.UpdateUser(c, Id, user)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: "user not updated"})
		return
	}

	// step 4: return user id
	c.JSON(http.StatusOK, models.Response{
		Message: updated_at,
	})
}

func (h *Handler) DeleteUserByID(c *gin.Context) {

	// step 1: validate input
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, models.Error{Message: "invalid input id"})
		return
	}

	Id, err := strconv.Atoi(id)
	fmt.Println(Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: "invalid input id"})
		return
	}

	// step 2: delete user
	deleted_at, err := h.service.DeleteUser(c, Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: "user not deleted"})
		return
	}

	// step 3: return user id
	c.JSON(http.StatusOK, models.Response{
		Message: deleted_at,
	})
}

func (h *Handler) GetUserById(c *gin.Context) {

	// step 1: validate input
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, models.Error{Message: "invalid input id"})
		return
	}

	Id, err := strconv.Atoi(id)
	fmt.Println(Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: "invalid input id"})
		return
	}

	// step 2: get user
	user, err := h.service.GetUser(c, Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: "user not found"})
		return
	}

	// step 3: return user
	c.JSON(http.StatusOK, user)
}

func (h *Handler) GetUsers(c *gin.Context) {

	created_at := c.Query("created_at")
	updated_at := c.Query("updated_at")

	if created_at != "" {
		fmt.Println("kirdiku")
		users, err := h.service.GetUsersByFieldCreatedAt(c, created_at)

		if err != nil {
			c.JSON(http.StatusBadRequest, models.Error{Message: "users not found"})
			return
		}

		c.JSON(http.StatusOK, users)
		return
	}

	if updated_at != "" {
		users, err := h.service.GetUsersByFieldUpdatedAt(c, updated_at)

		if err != nil {
			c.JSON(http.StatusBadRequest, models.Error{Message: "users not found"})
			return
		}

		c.JSON(http.StatusOK, users)
		return
	}

	if c.Query("full_name")=="asc" && c.Query("id") == "desc" {
		users, err := h.service.GetUserFullnameSort(c)

		if err != nil {
			c.JSON(http.StatusBadRequest, models.Error{Message: "users not found"})
			return
		}

		c.JSON(http.StatusOK, users)
		return
	}

	if created_at == "" && updated_at == "" && c.Query("full_name") != "asc" && c.Query("id") != "desc"{
		fmt.Println("kirdi bunisiga")
		users, err := h.service.GetUsers(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.Error{Message: "users not found"})
			return
		}

		c.JSON(http.StatusOK, users)
		return
	}
}