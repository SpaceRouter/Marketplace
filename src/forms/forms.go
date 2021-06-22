package forms

import "github.com/spacerouter/marketplace/models"

type BasicResponse struct {
	Message string
	Ok      bool
}

type StackResponse struct {
	Message   string
	Ok        bool
	Stack     *models.Stack
	Developer *models.Developer
}

type StackSearchResponse struct {
	Message string
	Ok      bool
	Stacks  []StackInfo
}

type StackInfo struct {
	ID          uint
	Name        string
	Icon        string
	Description string
	Developer   *models.Developer
}
