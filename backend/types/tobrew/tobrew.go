package tobrew

import "database/sql"

type ToBrew struct {
	Id         string         `json:"id"`
	Name       string         `json:"name"`
	Bean       string         `json:"bean"`
	Link       sql.NullString `json:"link"`
	Roaster    sql.NullString `json:"roaster"`
	TimeToBrew string         `db:"time_of_brew" json:"timeToBrew"`
	Created    string         `json:"created"`
}
