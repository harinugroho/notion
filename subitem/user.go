package subitem

type User struct {
	Object    string `json:"object"`
	Id        string `json:"id"`
	Name      string `json:"name"`
	AvatarUrl string `json:"avatar_url"`
	Type      string `json:"type"`
}
