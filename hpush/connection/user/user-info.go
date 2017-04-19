package user

type UserInfo struct {
	appid    string
	username string
	group    string
}

func (u *UserInfo) UserInfoKeyString() (key string) {
	key = u.appid + ";" + u.username
	return
}
