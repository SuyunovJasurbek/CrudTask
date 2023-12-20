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

func (h *Handler) UpdateUser(c *gin.Context) {

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
		c.JSON(http.StatusBadRequest, models.Error{Message: "invalddddid input id"})
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