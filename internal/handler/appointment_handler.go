package handler

import (
    "ehr-api/internal/dto"
    "ehr-api/internal/usecase"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

type AppointmentHandler struct {
    usecase usecase.AppointmentUsecase
}

func NewAppointmentHandler(uc usecase.AppointmentUsecase) *AppointmentHandler {
    return &AppointmentHandler{usecase: uc}
}

func (h *AppointmentHandler) Create(c *gin.Context) {
    var req dto.CreateAppointmentRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, dto.Response{Success: false, Message: err.Error()})
        return
    }
    res, err := h.usecase.Create(c.Request.Context(), req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, dto.Response{Success: false, Message: err.Error()})
        return
    }
    c.JSON(http.StatusCreated, dto.Response{Success: true, Message: "Appointment created successfully", Data: res})
}

func (h *AppointmentHandler) FindAll(c *gin.Context) {
    res, err := h.usecase.FindAll(c.Request.Context())
    if err != nil {
        c.JSON(http.StatusInternalServerError, dto.Response{Success: false, Message: err.Error()})
        return
    }
    c.JSON(http.StatusOK, dto.Response{Success: true, Message: "Success", Data: res})
}

func (h *AppointmentHandler) FindByID(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.ParseUint(idStr, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, dto.Response{Success: false, Message: "Invalid ID"})
        return
    }
    res, err := h.usecase.FindByID(c.Request.Context(), uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, dto.Response{Success: false, Message: "Appointment not found"})
        return
    }
    c.JSON(http.StatusOK, dto.Response{Success: true, Message: "Success", Data: res})
}

func (h *AppointmentHandler) Update(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.ParseUint(idStr, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, dto.Response{Success: false, Message: "Invalid ID"})
        return
    }
    var req dto.UpdateAppointmentRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, dto.Response{Success: false, Message: err.Error()})
        return
    }
    res, err := h.usecase.Update(c.Request.Context(), uint(id), req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, dto.Response{Success: false, Message: err.Error()})
        return
    }
    c.JSON(http.StatusOK, dto.Response{Success: true, Message: "Appointment updated successfully", Data: res})
}

func (h *AppointmentHandler) Delete(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.ParseUint(idStr, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, dto.Response{Success: false, Message: "Invalid ID"})
        return
    }
    if err := h.usecase.Delete(c.Request.Context(), uint(id)); err != nil {
        c.JSON(http.StatusInternalServerError, dto.Response{Success: false, Message: err.Error()})
        return
    }
    c.JSON(http.StatusOK, dto.Response{Success: true, Message: "Appointment deleted successfully"})
}
