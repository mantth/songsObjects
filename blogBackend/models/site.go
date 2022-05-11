package models

// Site 网站信息；
type Site struct {
	Avatar string `json:"avatar" db:"avatar"`
	Slogan string `json:"slogan" db:"slogan"`
	Name   string `json:"name" db:"name"`
	Domain string `json:"domain" db:"domain"`
	Notice string `json:"notice" db:"notice"`
	Desc   string `json:"desc" db:"desc"`
}
