package handler

import (
	"net/http"
	"strings"

	"github.com/Ki4EH/lib-service/account/entities"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func (h *Handler) register(c *gin.Context) {
	var input entities.User
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	_, err = h.services.SendConfirmMail(input)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type loginDto struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password_hash" binding:"required"`
}

func (h *Handler) login(c *gin.Context) {
	var input loginDto

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Login, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
	c.SetCookie("token", token, 60*60*60*24, "/", strings.Split(viper.GetString("SERVER_URL"), ":")[0], false, true)
}

func (h *Handler) logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "localhost", false, true)
}

func (h *Handler) verify(c *gin.Context) {
	if err := h.services.ConfirmEmail(c.Param("token")); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.Redirect(300, viper.GetString("CLIENT_URL"))
}
