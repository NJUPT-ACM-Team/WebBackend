package models

import (
	"time"
)

type OJInfo struct {
	OJId       int `db:"oj_id"`
	Name       string
	Version    string
	Int64IO    string
	JavaClass  string
	Status     string
	StatusInfo string `db:"status_info"`
	LastCheck  time.Time
}

type OJInfoModel struct {
	Model
}

func (this *OJInfoModel) Insert(oj *OJInfo) (int, error) {
	if err := this.OpenDB(); err != nil {
		return 0, err
	}
	defer this.CloseDB()
	last_insert_id, err := this.InlineInsert(oj, "ojinfo", []string{"oj_id"})
	if err != nil {
		return 0, err
	}
	return last_insert_id, nil
}
