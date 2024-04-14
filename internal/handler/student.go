package handler

import (
	"backend/internal/handler/dto"
	"backend/internal/models"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// GetUserByName godoc
//
//	@Summary		Get User ID by Name
//	@Description	Retrieves the user ID by the user's name from the database.
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			name	query		string				true	"Name of the User"
//	@Success		200		{object}	string				"ID of the User"
//	@Failure		400		{object}	models.ErrResponse	"Bad Request: Insufficient query arguments or no user found"
//	@Failure		500		{object}	models.ErrResponse	"Internal Server Error"
//	@Router			/v1/student/get-student-id-by-name [get]
func (h *Handler) GetUserByName(c *gin.Context) {
	name, ok := c.GetQuery("name")
	if !ok {
		log.Print(models.ErrInsufficientQueryArg)
		c.AbortWithStatusJSON(http.StatusBadRequest, models.ErrResponse{Msg: models.ErrInsufficientQueryArg.Error()})
		return
	}

	res, err := h.service.Student.GetUserIdByName(c.Request.Context(), name)
	if err != nil {
		log.Print(err)
		if errors.Is(err, models.ErrNoDocument) {
			c.AbortWithStatusJSON(http.StatusBadRequest, models.ErrResponse{Msg: err.Error()})
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, models.ErrResponse{Msg: err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *Handler) RequestStudentQuestion(c *gin.Context) {
	var req dto.GetStudentQuestion
	if err := c.BindJSON(&req); err != nil {
		log.Print(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, models.ErrResponse{Msg: err.Error()})
		return
	}

	resp, err := h.service.Student.AnswerStudentReq(c.Request.Context(), req)
	if err != nil {
		log.Print(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, models.ErrResponse{Msg: err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
