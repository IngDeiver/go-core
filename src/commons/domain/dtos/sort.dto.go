package dtos

type SortParamsDto struct {
	Field     string `json:"sort_field,omitempty" form:"sort_field,omitempty"`
	Direction string `json:"sort_direction,omitempty" form:"sort_direction,omitempty" binding:"omitempty,oneof=asc desc"`
}