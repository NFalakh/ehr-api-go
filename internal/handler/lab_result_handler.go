package handler

import (
    "ehr-api/internal/dto"
    "ehr-api/internal/usecase"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

type LabResultHandler struct {
    usecase usecase.LabResultUsecase
}

func NewLabResultHandler(uc usecase.LabResultUsecase) *LabResultHandler {
    return &LabResultHandler{usecase: uc}
}

func (h *LabResultHandler) Create(c *gin.Context) {
    var req dto.CreateLabResultRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, dto.Response{Success: false, Message: err.Error()})
        return
    }
    res, err := h.usecase.Create(c.Request.Context(), req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, dto.Response{Success: false, Message: err.Error()})
        return
    }
    c.JSON(http.StatusCreated, dto.Response{Success: true, Message: "Lab result created successfully", Data: res})
}

func (h *LabResultHandler) FindAll(c *gin.Context) {
    res, err := h.usecase.FindAll(c.Request.Context())
    if err != nil {
        c.JSON(http.StatusInternalServerError, dto.Response{Success: false, Message: err.Error()})
        return
    }
    c.JSON(http.StatusOK, dto.Response{Success: true, Message: "Success", Data: res})
}

func (h *LabResultHandler) FindByID(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.ParseUint(idStr, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, dto.Response{Success: false, Message: "Invalid ID"})
        return
    }
    res, err := h.usecase.FindByID(c.Request.Context(), uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, dto.Response{Success: false, Message: "Lab result not found"})
        return
    }
    c.JSON(http.StatusOK, dto.Response{Success: true, Message: "Success", Data: res})
}

func (h *LabResultHandler) Update(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.ParseUint(idStr, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, dto.Response{Success: false, Message: "Invalid ID"})
        return
    }
    var req dto.UpdateLabResultRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, dto.Response{Success: false, Message: err.Error()})
        return
    }
    res, err := h.usecase.Update(c.Request.Context(), uint(id), req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, dto.Response{Success: false, Message: err.Error()})
        return
    }
    c.JSON(http.StatusOK, dto.Response{Success: true, Message: "Lab result updated successfully", Data: res})
}

func (h *LabResultHandler) Delete(c *gin.Context) {
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
    c.JSON(http.StatusOK, dto.Response{Success: true, Message: "Lab result deleted successfully"})
}
