package authDto


type RestorePasswordDto struct {
    Password        string `json:"password" binding:"required,min=8,password"`
    ConfirmPassword string `json:"confirmPassword" binding:"required,eqfield=Password"`
}