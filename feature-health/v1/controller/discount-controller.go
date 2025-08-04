package controller

import (
	core_model "confluence-payment/core-internal/model"
	"confluence-payment/feature-health/v1/service"

	"github.com/gin-gonic/gin"
)

type HealthController struct {
	HealthService *service.HealthService
}

func NewHealthController(healthService *service.HealthService) *HealthController {
	return &HealthController{
		HealthService: healthService,
	}
}

func (h *HealthController) Health(c *gin.Context) {
	res := core_model.CoreResponseObject{}

	data, err, statusCode := h.HealthService.Health()
	if err != nil {
		c.AbortWithStatusJSON(statusCode, data)
		return
	}

	res.Data = data
	res.Succeeded = true
	c.JSON(statusCode, res)
}
