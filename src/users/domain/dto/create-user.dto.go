package userDtos

type CreateUserDto struct {
	FirstName string `json:"firstName" bson:"firstName" binding:"required"`
    LastName  string `json:"lastName" bson:"lastName" binding:"required"`
    Email     string `json:"email" bson:"email" binding:"required,email"`
    Password  string `json:"password" bson:"password" binding:"required"`
}
