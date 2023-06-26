package posts

type PostStatus string

const(
	PostStatusActive PostStatus = "A"
	PostStatusHidden PostStatus = "H"
)

func PostStatusList() map[string]PostStatus {
	return map[string]PostStatus{
		"Active": PostStatusActive,
		"Hidden": PostStatusHidden,
	}
}