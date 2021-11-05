package dto

type ClientOrder struct {
	ClientID int `json:"client_id"`
	Orders []Order `json:"orders"`
}
