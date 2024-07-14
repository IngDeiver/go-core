package dtos
type PaginationParamsDto struct {
	Page  *int `form:"page"  binding:"omitempty,min=1"`
	Limit *int `form:"limit"  binding:"omitempty,min=1"`
}

func (p *PaginationParamsDto) GetPage() int {
	if p.Page == nil {
		return 1
	}
	return *p.Page
}

func (p *PaginationParamsDto) GetLimit() int {
	if p.Limit == nil {
		return 10
	}
	return *p.Limit
}



type PaginationMetadata struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	TotalPages int `json:"totalPages"`
	TotalCount int `json:"totalCount"`
}

type PagedResponse[T any] struct {
	Data              []T               `json:"data"`
	PaginationMetadata PaginationMetadata `json:"paginationMetadata"`
}

