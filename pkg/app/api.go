package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ContactUsForm struct {
	Email   string `json:"email"`
	Name    string `json:"name"`
	Message string `json:"message"`
	Budget  string `json:"budget"`
}

func ContactUs(c *gin.Context) {
	var form ContactUsForm
	err := c.ShouldBindJSON(&form)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	err = SendMessageToChat(fmt.Sprintf("%s\n%s\n%s\n%s", form.Email, form.Name, form.Budget, form.Message))
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	c.Status(http.StatusOK)
	return
}
