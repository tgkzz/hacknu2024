package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// HelloWorld godoc
//	@Summary		Show hello world message
//	@Description	get string by ID
//	@Tags			example
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	string	"hello world"
//	@Router			/helloworld [get]
func (h *Handler) HelloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, "hello world")
}
