package users

type UserStatus string

const(
	UserStatusActive UserStatus = "A"
	UserStatusDisabled UserStatus = "D"
	UserStatusHidden UserStatus = "H"
)

func UserStatusList() map[string]UserStatus {
	return map[string]UserStatus{
		"Active": UserStatusActive,
		"Disabled": UserStatusDisabled,
		// "Hidden": UserStatusHidden,
	}
}