package handler

import (
	"net/http"

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
	if err, ok := helper.UserCreateFeildCheck(user); !ok {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	// step 3: create user
	err := h.service.CreateUser(c, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	
	// step 4: return user id
	c.JSON(http.StatusCreated, models.Response{
		Message: "user created",
	})
}
