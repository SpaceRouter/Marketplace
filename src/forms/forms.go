package forms

import "github.com/spacerouter/marketplace/models"

type StackInfo struct {
	Message string
	Ok      bool
	Stack   *models.Stack
}
