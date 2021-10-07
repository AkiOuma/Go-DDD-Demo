package entity

import (
	"gin-ddd-demo/module/command/wxwork/domain/valobj"
	"strconv"
)

type ContactEntity struct {
	accessToken string
}

func NewContactEntity(accessToken string) *ContactEntity {
	return &ContactEntity{
		accessToken: accessToken,
	}
}

func (entity *ContactEntity) GetWxUserList() ([]*valobj.UserDetail, error) {
	// response, err := http.Get(entity.BuildURL())
	// if err != nil {
	// 	return nil, err
	// }
	// defer response.Body.Close()

	// if err := json.Unmarshal(); err != nil {
	// 	return nil, err
	// }

	// simulate get user detail
	users := make([]*valobj.UserDetail, 2)
	for i := 0; i < len(users); i++ {
		index := strconv.Itoa(i)
		user := valobj.NewUserDetail()
		user.SetUserID("user" + index)
		user.SetName("user" + index)
		user.SetEmail("user" + index + "@yuki.com")
		user.SetAvatar("user" + index)
		user.SetPhone("18812345678")
		users[i] = user
	}
	return users, nil
}

// func (entity *ContactEntity) buildURL() string {
// 	return fmt.Sprintf("accesstoken=%s", entity.accessToken)
// }
