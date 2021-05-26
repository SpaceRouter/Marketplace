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
		c.AbortWithStatusJSON(http.StatusBadRequest, forms.StackResponse{
			Message:   err.Error(),
			Ok:        false,
			Stack:     nil,
			Developer: nil,
		})
		return
	}
	var stack models.Stack
	result := s.DB.Preload(clause.Associations).First(&stack, parseId)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, forms.StackResponse{
				Message:   result.Error.Error(),
				Ok:        false,
				Stack:     nil,
				Developer: nil,
			})
			return
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, forms.StackResponse{
				Message:   result.Error.Error(),
				Ok:        false,
				Stack:     nil,
				Developer: nil,
			})
			return
		}
	}
	developer := models.Developer{}
	s.DB.First(&developer, map[string]interface{}{"ID": stack.ID})

	c.JSON(http.StatusOK, forms.StackResponse{
		Message:   "",
		Ok:        true,
		Stack:     &stack,
		Developer: &developer,
	})
}

func (s *StackController) GetStackSearch(c *gin.Context) {
	search := c.Param("search")
	var stacks []models.Stack
	result := s.DB.Where("name LIKE ?", "%"+search+"%").Find(&stacks)

	if result.Error != nil {
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusInternalServerError, forms.StackSearchResponse{
				Message: result.Error.Error(),
				Ok:      false,
				Stacks:  nil,
			})
			return
		}
	}

	var stacksInfo []forms.StackInfo

	for _, stack := range stacks {
		stacksInfo = append(stacksInfo, StackToInfo(stack, s.DB))
	}

	c.JSON(http.StatusOK, forms.StackSearchResponse{
		Message: "",
		Ok:      true,
		Stacks:  stacksInfo,
	})
}

func (s *StackController) GetStackByUserId(c *gin.Context) {
	id := c.Param("id")
	var stacks []models.Stack
	result := s.DB.Where("name=?", id).Find(&stacks)

	if result.Error != nil {
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusInternalServerError, forms.StackSearchResponse{
				Message: result.Error.Error(),
				Ok:      false,
				Stacks:  nil,
			})
			return
		}
	}

	var stacksInfo []forms.StackInfo

	for _, stack := range stacks {
		stacksInfo = append(stacksInfo, StackToInfo(stack, s.DB))
	}

	c.JSON(http.StatusOK, forms.StackSearchResponse{
		Message: "",
		Ok:      true,
		Stacks:  stacksInfo,
	})
}

func StackToInfo(stack models.Stack, db *gorm.DB) forms.StackInfo {
	developer := models.Developer{}
	db.First(&developer, map[string]interface{}{"ID": stack.ID})
	return forms.StackInfo{
		ID:          stack.ID,
		Name:        stack.Name,
		Icon:        stack.Icon,
		Description: stack.Description,
		Developer:   &developer,
	}
}
