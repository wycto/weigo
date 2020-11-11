package model

type User struct {
	Name     string
	Age      int
	Like     []string
	UserInfo map[string]interface{}
}
