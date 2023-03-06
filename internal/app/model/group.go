package model

type Group struct {
	Id         int    `json:"id"`
	Department string `json:"department"`
	CuratorId  int    `json:"curator_id"`
}
