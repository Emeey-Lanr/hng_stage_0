package models



type UserData struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	Stack string `json:"stack"`
}
type MeResponse struct {
	Status    string   `json:"status,omitempty"`
	User    interface{} `json:"user,omitempty"`
	Timestamp  string `json:"timestamp,omitempty"`
	Fact string `json:"fact,omitempty"`

}
