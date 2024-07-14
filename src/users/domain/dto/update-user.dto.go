package userDtos

type UpdateUserDto struct {
    FirstName string `json:"firstName,omitempty" bson:"firstName,omitempty" binding:"omitempty"`
    LastName  string `json:"lastName,omitempty" bson:"lastName,omitempty" binding:"omitempty"`
    Email     string `json:"email,omitempty" bson:"email,omitempty" binding:"omitempty,email"`
    Password  string `json:"password,omitempty" bson:"password,omitempty" binding:"omitempty"`
}