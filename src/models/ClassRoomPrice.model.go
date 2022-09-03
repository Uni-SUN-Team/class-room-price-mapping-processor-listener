package models

type ClassRoomPrice struct {
	Id           int     `json:"id"`
	ClassRoomId  int     `json:"classRoomId"`
	RegularPrice float64 `json:"regularPrice"`
	SpecialPrice float64 `json:"specialPrice"`
	Advisors     string  `json:"advisors"`
	Categories   string  `json:"categories"`
}
