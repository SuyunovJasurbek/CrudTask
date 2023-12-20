package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/SuyunovJasurbek/CrudTask/helper"
	"github.com/SuyunovJasurbek/CrudTask/models"
	"github.com/gin-gonic/gin"
)
// Create User
// @Summary  Create User
// @Description  Create User
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        user body models.UserHandler true "User info"
// @Success      201  {object}  models.Response
// @Response     400 {object}  models.Response
// @Failure  	 500  {object}  models.Response
// @Router        /api/v1/users/	[post]
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
// Update User
// @Summary  Update User
// @Description  Update User
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        id path string true "User ID"
// @Param        user body models.UserHandler true "User info"
// @Success      200  {object}  models.Response
// @Response     400 {object}  models.Response
// @Failure  	 500  {object}  models.Response
// @Router        /api/v1/users/{id}	[put]
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
// Delete User
// @Summary  Delete User
// @Description  Delete User
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        id path string true "User ID"
// @Success      200  {object}  models.Response
// @Response     400 {object}  models.Response
// @Failure  	 500  {object}  models.Response
// @Router        /api/v1/users/{id}	[delete]
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
// Get User
// @Summary  Get User
// @Description  Get User
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        id path string true "User ID"
// @Success      200  {object}  models.GetUser
// @Response     400 {object}  models.Response
// @Failure  	 500  {object}  models.Response
// @Router        /api/v1/users/{id}	[get]
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
// Get Users
// @Summary  Get Users
// @Description  Get Users
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        created_at query string false "User Created At example:2023-12-20 09:18:37"
// @Param        updated_at query string false "User Updated At example:2023-12-20 09:18:37"
// @Param        asc query string false "User Full Name"
// @Param        desc query string false "User ID"
// @Success      200  {object}  []models.GetUser
// @Response     400 {object}  models.Response
// @Failure  	 500  {object}  models.Response
// @Router        /api/v1/users/	[get]
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
		users, err := h.service.GetUsersFullnameSort(c)

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
// multi

// Create Users
// @Summary  Create Users
// @Description  Create Users
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        users body []models.UserHandler true "Users info"
// @Success      201  {object}  models.Response
// @Response     400 {object}  models.Response
// @Failure  	 500  {object}  models.Response
// @Router        /api/v1/users/multi	[post]
func (h *Handler) CreateUsers(c *gin.Context) {
	var users []models.UserHandler
	if err := c.ShouldBindJSON(&users); err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: "invalid input body"})
		return
	}
	for v, user := range users {
		if err, ok := helper.UserFeildCheck(user); !ok {
			c.JSON(http.StatusBadRequest, models.Error{Message:strconv.Itoa(v+1)+" json: "+ err.Message})
			return
		}
	}

	err := h.service.CreateUsers(c, users)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Error{Message: "users not created"})
		return
	}

	c.JSON(http.StatusCreated, models.Response{
		Message: "users created",
	})
}
// Update Users
// @Summary  Update Users
// @Description  Update Users
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        users body []models.UsersUpdateHandler true "Users info"
// @Success      200  {object}  models.Response
// @Response     400 {object}  models.Response
// @Failure  	 500  {object}  models.Response
// @Router        /api/v1/users/multi	[put]
func (h *Handler) UpdateUsers(c *gin.Context) {
	var users []models.UsersUpdateHandler
	if err := c.ShouldBindJSON(&users); err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: "invalid input body"})
		return
	}
	for v, user := range users {
		if err, ok := helper.UserUpdateFeildCheck(user); !ok {
			c.JSON(http.StatusBadRequest, models.Error{Message:strconv.Itoa(v+1)+" json: "+ err.Message})
			return
		}
	}

	err := h.service.UpdateUsers(c, users)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Error{Message: "users not updated"})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Message: "users updated",
	})
}
// Delete Users
// @Summary  Delete Users
// @Description  Delete Users
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        users body []models.UserIds true "Users info"
// @Success      200  {object}  models.Response
// @Response     400 {object}  models.Response
// @Failure  	 500  {object}  models.Response
// @Router        /api/v1/users/multi	[delete]
func (h *Handler) DeleteUsers(c *gin.Context) {
	var users []models.UserIds
	if err := c.ShouldBindJSON(&users); err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: "invalid input body"})
		return
	}
	for _, user := range users {
		if user.ID == 0 {
			c.JSON(http.StatusBadRequest, models.Error{Message: "id must be greater than 0"})
			return
		}
	}

	err := h.service.DeleteUsers(c, users)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Error{Message: "users not deleted"})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Message: "users deleted",
	})
}