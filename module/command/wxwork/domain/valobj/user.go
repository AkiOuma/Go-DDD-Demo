package valobj

type UserDetail struct {
	userID string
	name   string
	email  string
	phone  string
	avatar string
}

func NewUserDetail() *UserDetail {
	return &UserDetail{}
}

func (detail *UserDetail) UserID() string {
	return detail.userID
}

func (detail *UserDetail) Name() string {
	return detail.name
}

func (detail *UserDetail) Email() string {
	return detail.email
}

func (detail *UserDetail) Phone() string {
	return detail.phone
}

func (detail *UserDetail) Avatar() string {
	return detail.avatar
}

func (detail *UserDetail) SetUserID(userID string) {
	detail.userID = userID
}

func (detail *UserDetail) SetName(name string) {
	detail.name = name
}

func (detail *UserDetail) SetEmail(email string) {
	detail.email = email
}

func (detail *UserDetail) SetPhone(phone string) {
	detail.phone = phone
}

func (detail *UserDetail) SetAvatar(avatar string) {
	detail.avatar = avatar
}
