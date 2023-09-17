package handler

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"gin/global/errors"
	"gin/global/response"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

func (h *Handler) GetAllVenue(ctx *gin.Context) {
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

	venues, err := h.Venue.GetAllVenue(c)
	if err != nil {
		code = http.StatusInternalServerError
		message = err.Error()
		return
	}

	select {
	case <-c.Done():
		code = http.StatusRequestTimeout
		message = errors.ErrRequestTimeout.Error()
	default:
		message = "Success to get all venue"
		data = venues
	}
}

func (h *Handler) GetVenueByID(ctx *gin.Context) {
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
		code = http.StatusBadRequest
		message = errors.ErrInvalidRequest.Error()
		return
	}

	venues, err := h.Venue.GetVenueByID(c, uint(id))
	if err != nil {
		code = http.StatusInternalServerError
		message = err.Error()
		return
	}

	select {
	case <-c.Done():
		code = http.StatusRequestTimeout
		message = errors.ErrRequestTimeout.Error()
	default:
		message = "Success to get all venue"
		data = venues
	}
}

func (h *Handler) RentVenue(ctx *gin.Context) {
	c, cancel := context.WithTimeout(ctx.Request.Context(), 6000*time.Millisecond)
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

	userID, err := uuid.FromString(ctx.MustGet("user").(string))
	if err != nil {
		message = errors.ErrInvalidRequest.Error()
		code = http.StatusBadRequest
		return
	}

	venueDayID, err := strconv.Atoi(ctx.Param("venueday_id"))
	if err != nil {
		code = http.StatusBadRequest
		message = errors.ErrInvalidRequest.Error()
		return
	}

	res, err := h.Venue.RentVenue(c, userID, uint(venueDayID))
	if err != nil {
		code = http.StatusInternalServerError
		message = errors.ErrInternalServer.Error()
		return
	}

	select {
	case <-c.Done():
		code = http.StatusRequestTimeout
		message = errors.ErrRequestTimeout.Error()
	default:
		message = "Success to rent venue"
		data = res
	}
}

func (h *Handler) GetListApplyVenue(ctx *gin.Context) {
	c, cancel := context.WithTimeout(ctx.Request.Context(), 6000*time.Millisecond)
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

	userID := ctx.MustGet("userID").(uuid.UUID)

	res, err := h.Venue.GetListApplyVenue(c, userID)
	if err != nil {
		code = http.StatusInternalServerError
		message = errors.ErrInternalServer.Error()
		return
	}

	select {
	case <-c.Done():
		code = http.StatusRequestTimeout
		message = errors.ErrRequestTimeout.Error()
		return
	default:
		message = "Success to rent venue"
		data = res
	}
}
