package core

type MetaService struct {
}

type Meta struct {
	Name    string `json:"name"`
	Version int32  `json:"version,string"`
	Length  uint64 `json:"length,string"`
	Hash    string `json:"hash"`
}
