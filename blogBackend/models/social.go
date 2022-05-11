package models

// Social 社交信息；
type Social struct {
	ID    uint64 `json:"id" db:"id"`
	Title string `json:"title" db:"title"`
	Icon  string `json:"icon" db:"icon"`
	Color string `json:"color" db:"color"`
	Href  string `json:"href" db:"href"`
}
