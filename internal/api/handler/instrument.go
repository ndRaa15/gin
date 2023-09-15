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

func (h *Handler) GetAllInstrument(ctx *gin.Context) {
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

	instruments, err := h.Instrument.GetAllInstrument(c)
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
		data = instruments
	}
}

func (h *Handler) GetInstrumentByID(ctx *gin.Context) {
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

	res, err := h.Instrument.GetByID(c, uint(id))
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

func (h *Handler) RentInstrument(ctx *gin.Context) {
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

	req := entity.RentInstrument{}
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

	res, err := h.Instrument.RentInstrument(c, uint(idInstrument), &req, idUser)
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

func (h *Handler) GetProvince(ctx *gin.Context) {
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

	idProvince := ctx.Query("id")
	res, err := h.Instrument.GetProvince(c, idProvince)
	if err != nil {
		code = http.StatusBadRequest
		message = errors.ErrBadRequest.Error()
		return
	}

	select {
	case <-c.Done():
		code = http.StatusRequestTimeout
		message = errors.ErrRequestTimeout.Error()
	default:
		message = "Success to get province"
		data = res
	}
}

func (h *Handler) GetCity(ctx *gin.Context) {
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

	idProvince := ctx.Query("province")
	idCity := ctx.Query("id")
	res, err := h.Instrument.GetCity(c, idCity, idProvince)
	if err != nil {
		code = http.StatusBadRequest
		message = errors.ErrBadRequest.Error()
		return
	}

	select {
	case <-c.Done():
		code = http.StatusRequestTimeout
		message = errors.ErrRequestTimeout.Error()
	default:
		message = "Success to get city"
		data = res
	}
}

func (h *Handler) GetCost(ctx *gin.Context) {
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

	req := entity.ShippingCost{}
	if err = ctx.ShouldBindJSON(&req); err != nil {
		message = errors.ErrInvalidRequest.Error()
		code = http.StatusBadRequest
		return
	}

	res, err := h.Instrument.GetCost(c, uint(id), &req)
	if err != nil {
		code = http.StatusBadRequest
		message = errors.ErrBadRequest.Error()
		return
	}

	select {
	case <-c.Done():
		code = http.StatusRequestTimeout
		message = errors.ErrRequestTimeout.Error()
	default:
		message = "Success to get cost"
		data = res
	}
}
