package userDtos

type UserFilterDto struct {
	FirstName  string `form:"firstName"`
	Role  string `form:"role"`
}