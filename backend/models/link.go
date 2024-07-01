package models

type Link struct {
	ID       int    `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Code     string `json:"code"`
	FullLink string `json:"full_link"`
}
