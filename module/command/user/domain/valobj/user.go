package valobj

type UserProfile struct {
	userID           string
	userName         string
	email            string
	phone            string
	organizationCode string
	avatar           string
}

func NewUserProfile(
	userID string,
	userName string,
	email string,
	phone string,
	organizationCode string,
	avatar string,
) *UserProfile {
	return &UserProfile{
		userID:           userID,
		userName:         userName,
		email:            email,
		phone:            phone,
		organizationCode: organizationCode,
		avatar:           avatar,
	}
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
