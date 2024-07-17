package userDtos

import userDomain "github.com/ingdeiver/go-core/src/users/domain"

// NOTE: always property should have to json,bson  and binding declarations
type CreateUserDto struct {
	FirstName string `json:"firstName" bson:"firstName" binding:"required"`
    LastName  string `json:"lastName" bson:"lastName" binding:"required"`
    Email     string `json:"email" bson:"email" binding:"required,email"`
    Password  string `json:"password" bson:"password" binding:"required,min=8,password"`
    Role      userDomain.Role  `json:"role,omitempty" bson:"role,omitempty" binding:"role,omitempty"`
}
