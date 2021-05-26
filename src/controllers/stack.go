package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/spacerouter/marketplace/forms"
	"github.com/spacerouter/marketplace/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"net/http"
	"strconv"
)

type StackController struct {
	DB *gorm.DB
}

func (s *StackController) GetStackById(c *gin.Context) {
	id := c.Param("id")
	parseId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, forms.StackInfo{
			Message: err.Error(),
			Ok:      false,
			Stack:   nil,
		})
		return
	}
	var stack models.Stack
	result := s.DB.Preload(clause.Associations).Find(&stack, parseId)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, forms.StackInfo{
				Message: result.Error.Error(),
				Ok:      false,
				Stack:   nil,
			})
			return
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, forms.StackInfo{
				Message: result.Error.Error(),
				Ok:      false,
				Stack:   nil,
			})
			return
		}
	}

	/*
		if s.DB.RowsAffected != 1 {
			c.AbortWithStatusJSON(http.StatusNotFound, forms.StackInfo{
				Message: "Not found",
				Ok:      false,
				Stack:   nil,
			})
			return
		}

	*/

	c.JSON(http.StatusOK, forms.StackInfo{
		Message: "",
		Ok:      true,
		Stack:   &stack,
	})
}
