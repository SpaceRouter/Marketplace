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

// GetStackById godoc
// @Summary Get stack by ID
// @Description Get all stack information from stack ID
// @ID get_stack_by_id
// @Produce  json
// @Success 200 {object} forms.StackResponse
// @Failure 500,400,401 {object} forms.StackResponse
// @Router /v1/stack/{id} [get]
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
	result := s.DB.Preload(clause.Associations).Preload("Services."+clause.Associations).First(&stack, parseId)

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

// GetStackSearch godoc
// @Summary Search for stacks by name
// @Description Search for stacks reduced information by name
// @ID get_stack_search
// @Produce  json
// @Success 200 {object} forms.StackSearchResponse
// @Failure 500,400,401 {object} forms.StackSearchResponse
// @Router /v1/search/name/{search} [get]
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

// GetStackByUserId godoc
// @Summary Get stack by user ID
// @Description Get all stacks information uploaded by a user
// @ID get_stack_by_user_ID
// @Produce  json
// @Success 200 {object} forms.StackSearchResponse
// @Failure 500,400,401 {object} forms.StackSearchResponse
// @Router /v1/search/developer/{search} [get]
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
