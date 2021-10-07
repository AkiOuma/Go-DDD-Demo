package valobj

type UserProfile struct {
	userID           string
	userName         string
	email            string
	phone            string
	organizationCode string
	avatar           string
}

func (profile *UserProfile) UserID() string {
	return profile.userID
}

func (profile *UserProfile) UserName() string {
	return profile.userName
}

func (profile *UserProfile) Email() string {
	return profile.email
}

func (profile *UserProfile) Phone() string {
	return profile.phone
}

func (profile *UserProfile) OrganizationCode() string {
	return profile.organizationCode
}

func (profile *UserProfile) Avatar() string {
	return profile.avatar
}

func (profile *UserProfile) SetUserID(userID string) {
	profile.userID = userID
}

func (profile *UserProfile) SetUserName(userName string) {
	profile.userName = userName
}

func (profile *UserProfile) SetEmail(email string) {
	profile.email = email
}

func (profile *UserProfile) SetPhone(phone string) {
	profile.phone = phone
}

func (profile *UserProfile) SetOrganizationCode(organizationCode string) {
	profile.organizationCode = organizationCode
}

func (profile *UserProfile) SetAvatar(avatar string) {
	profile.avatar = avatar
}
