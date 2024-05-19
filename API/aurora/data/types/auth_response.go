package types

type AuthData struct {
	Id         string     `json:"id"`
	Name       string     `json:"name"`
	Properties []Property `json:"properties"`
}
