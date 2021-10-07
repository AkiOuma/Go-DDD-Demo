package application

type Mail interface {
	SendPasswordResetMail(userID string, newPassword string)
}
