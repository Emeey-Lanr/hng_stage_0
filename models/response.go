package models

import "time"

type UserData struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	Stack string `json:"stack"`
}
type MeResponse struct {
	Status    string   `json:"status,omitempty"`
	User      UserData `json:"user,omitempty"`
	Timestamp time.Time `json:"timeStamp,omitempty"`
	Fact string `json:"fact,omitempty"`

}
