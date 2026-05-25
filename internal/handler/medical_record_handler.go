package handler

import (
    "ehr-api/internal/dto"
    "ehr-api/internal/usecase"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

type MedicalRecordHandler struct {
    usecase usecase.MedicalRecordUsecase
}

func NewMedicalRecordHandler(uc usecase.MedicalRecordUsecase) *MedicalRecordHandler {
    return &MedicalRecordHandler{usecase: uc}
}

func (h *MedicalRecordHandler) Create(c *gin.Context) {
    var req dto.CreateMedicalRecordRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, dto.Response{Success: false, Message: err.Error()})
        return
    }
    res, err := h.usecase.Create(c.Request.Context(), req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, dto.Response{Success: false, Message: err.Error()})
        return
    }
    c.JSON(http.StatusCreated, dto.Response{Success: true, Message: "Medical record created successfully", Data: res})
}

func (h *MedicalRecordHandler) FindAll(c *gin.Context) {
    res, err := h.usecase.FindAll(c.Request.Context())
    if err != nil {
        c.JSON(http.StatusInternalServerError, dto.Response{Success: false, Message: err.Error()})
        return
    }
    c.JSON(http.StatusOK, dto.Response{Success: true, Message: "Success", Data: res})
}

func (h *MedicalRecordHandler) FindByID(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.ParseUint(idStr, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, dto.Response{Success: false, Message: "Invalid ID"})
        return
    }
    res, err := h.usecase.FindByID(c.Request.Context(), uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, dto.Response{Success: false, Message: "Medical record not found"})
        return
    }
    c.JSON(http.StatusOK, dto.Response{Success: true, Message: "Success", Data: res})
}

func (h *MedicalRecordHandler) Update(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.ParseUint(idStr, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, dto.Response{Success: false, Message: "Invalid ID"})
        return
    }
    var req dto.UpdateMedicalRecordRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, dto.Response{Success: false, Message: err.Error()})
        return
    }
    res, err := h.usecase.Update(c.Request.Context(), uint(id), req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, dto.Response{Success: false, Message: err.Error()})
        return
    }
    c.JSON(http.StatusOK, dto.Response{Success: true, Message: "Medical record updated successfully", Data: res})
}

func (h *MedicalRecordHandler) Delete(c *gin.Context) {
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
    c.JSON(http.StatusOK, dto.Response{Success: true, Message: "Medical record deleted successfully"})
}
