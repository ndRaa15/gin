package handler

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"gin/global/errors"
	"gin/global/response"
	"gin/internal/domain/entity"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

func (h *Handler) GetAllStudio(ctx *gin.Context) {
	c, cancel := context.WithTimeout(ctx.Request.Context(), 15*time.Second)
	defer cancel()

	var (
		err     error
		message string
		code    = http.StatusOK
		data    interface{}
	)

	defer func() {
		if err != nil {
			response.Error(ctx, code, err, message, data)
			return
		}
		response.Success(ctx, code, message, data)
	}()

	studios, err := h.Studio.GetAll(c)
	if err != nil {
		message = errors.ErrInternalServer.Error()
		code = http.StatusInternalServerError
		return
	}

	select {
	case <-c.Done():
		code = http.StatusRequestTimeout
		message = errors.ErrRequestTimeout.Error()
		return
	default:
		message = "Success to get all instrument"
		data = studios
	}
}

func (h *Handler) GetStudioByID(ctx *gin.Context) {
	c, cancel := context.WithTimeout(ctx.Request.Context(), 15*time.Second)
	defer cancel()

	var (
		err     error
		message string
		code    = http.StatusOK
		data    interface{}
	)

	defer func() {
		if err != nil {
			response.Error(ctx, code, err, message, data)
			return
		}
		response.Success(ctx, code, message, data)
	}()

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		message = errors.ErrInvalidRequest.Error()
		code = http.StatusBadRequest
		return
	}

	res, err := h.Studio.GetByID(c, uint(id))
	if err != nil {
		message = errors.ErrInternalServer.Error()
		code = http.StatusInternalServerError
		return
	}

	select {
	case <-c.Done():
		code = http.StatusRequestTimeout
		message = errors.ErrRequestTimeout.Error()
	default:
		message = "Success to get instrument by id"
		data = res
	}
}

func (h *Handler) RentStudio(ctx *gin.Context) {
	c, cancel := context.WithTimeout(ctx.Request.Context(), 15*time.Second)
	defer cancel()

	var (
		err     error
		message string
		code    = http.StatusOK
		data    interface{}
	)

	defer func() {
		if err != nil {
			response.Error(ctx, code, err, message, data)
			return
		}
		response.Success(ctx, code, message, data)
	}()

	idUser, err := uuid.FromString(ctx.MustGet("user").(string))
	if err != nil {
		message = errors.ErrInvalidRequest.Error()
		code = http.StatusBadRequest
		return
	}

	req := entity.RentStudio{}
	if err = ctx.ShouldBindJSON(&req); err != nil {
		message = errors.ErrInvalidRequest.Error()
		code = http.StatusBadRequest
		return
	}

	idInstrument, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		message = errors.ErrInvalidRequest.Error()
		code = http.StatusBadRequest
		return
	}

	res, err := h.Studio.RentStudio(c, uint(idInstrument), idUser, &req)
	if err != nil {
		message = errors.ErrInternalServer.Error()
		code = http.StatusInternalServerError
		return
	}

	select {
	case <-c.Done():
		code = http.StatusRequestTimeout
		message = errors.ErrRequestTimeout.Error()
	default:
		message = "Success to get instrument"
		data = res
	}
}
