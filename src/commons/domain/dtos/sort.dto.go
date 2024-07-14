package dtos

type SortParamsDto struct {
	SortField     string `form:"sortField,omitempty" binding:"omitempty"`
	SortDirection string `form:"sortDirection,omitempty" binding:"omitempty,oneof=asc desc"`
}