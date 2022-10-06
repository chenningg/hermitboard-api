package auth

import "github.com/chenningg/hermitboard-api/ent"

type LoginToAccountPayload struct {
	Account *ent.Account
	Session *Session
}

type LoginToStaffAccountPayload struct {
	StaffAccount *ent.StaffAccount
	Session      *Session
}

type CreateAccountPayload struct {
	Account *ent.Account
	Session *Session
}

type CreateStaffAccountPayload struct {
	StaffAccount *ent.StaffAccount
	Session      *Session
}

type LoginToAccountInput struct {
	Username string
	Password string
}

type LoginToStaffAccountInput struct {
	Username string
	Password string
}
