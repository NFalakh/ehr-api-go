package handler

import (
    "ehr-api/internal/dto"
    "ehr-api/internal/usecase"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

type VitalSignHandler struct {
    usecase usecase.VitalSignUsecase
}

func NewVitalSignHandler(uc usecase.VitalSignUsecase) *VitalSignHandler {
    return &VitalSignHandler{usecase: uc}
}

func (h *VitalSignHandler) Create(c *gin.Context) {
    var req dto.CreateVitalSignRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, dto.Response{Success: false, Message: err.Error()})
        return
    }
    res, err := h.usecase.Create(c.Request.Context(), req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, dto.Response{Success: false, Message: err.Error()})
        return
    }
    c.JSON(http.StatusCreated, dto.Response{Success: true, Message: "Vital sign created successfully", Data: res})
}

func (h *VitalSignHandler) FindAll(c *gin.Context) {
    res, err := h.usecase.FindAll(c.Request.Context())
    if err != nil {
        c.JSON(http.StatusInternalServerError, dto.Response{Success: false, Message: err.Error()})
        return
    }
    c.JSON(http.StatusOK, dto.Response{Success: true, Message: "Success", Data: res})
}

func (h *VitalSignHandler) FindByID(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.ParseUint(idStr, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, dto.Response{Success: false, Message: "Invalid ID"})
        return
    }
    res, err := h.usecase.FindByID(c.Request.Context(), uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, dto.Response{Success: false, Message: "Vital sign not found"})
        return
    }
    c.JSON(http.StatusOK, dto.Response{Success: true, Message: "Success", Data: res})
}

func (h *VitalSignHandler) Update(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.ParseUint(idStr, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, dto.Response{Success: false, Message: "Invalid ID"})
        return
    }
    var req dto.UpdateVitalSignRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, dto.Response{Success: false, Message: err.Error()})
        return
    }
    res, err := h.usecase.Update(c.Request.Context(), uint(id), req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, dto.Response{Success: false, Message: err.Error()})
        return
    }
    c.JSON(http.StatusOK, dto.Response{Success: true, Message: "Vital sign updated successfully", Data: res})
}

func (h *VitalSignHandler) Delete(c *gin.Context) {
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
    c.JSON(http.StatusOK, dto.Response{Success: true, Message: "Vital sign deleted successfully"})
}
