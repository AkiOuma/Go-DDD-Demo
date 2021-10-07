package valobj

type Password struct {
	value string
}

func NewPassword(password string) (*Password, error) {
	pass := &Password{}
	if err := pass.setPassword(pass.decrypt(password)); err != nil {
		return nil, err
	}
	return pass, nil
}

func (pass *Password) validatePassword(password string) error {
	return nil
}

func (pass *Password) ChangePassword(newPassword string) error {
	return pass.setPassword(newPassword)
}

func (pass *Password) GetPassword() string {
	return pass.decrypt(pass.value)
}

func (pass *Password) setPassword(password string) error {
	if err := pass.validatePassword(password); err != nil {
		return err
	}
	pass.value = pass.encrypt(password)
	return nil
}

func (pass *Password) encrypt(password string) string {
	return password
}

func (pass *Password) decrypt(password string) string {
	return password
}

func (pass *Password) Equal(password string) bool {
	return pass.GetPassword() == password
}

func (pass *Password) ResetPassword() string {
	// random password rule here
	newPassword := "someword"
	for {
		if err := pass.setPassword(newPassword); err == nil {
			break
		}
	}
	return newPassword
}
