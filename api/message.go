package ants

type Message struct {
	SignInURL              string
	SignOutURL             string
	NeedAdminAuthorization bool
	IsLogin                bool
	IsAdmin                bool
	Error                  bool
	ErrorMessage           string
}
