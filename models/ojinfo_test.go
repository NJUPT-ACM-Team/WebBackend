package models

import (
	"testing"
	"time"
)

func TestOJInfo(t *testing.T) {
	ojim := OJInfoModel{}
	ojinfo := OJInfo{
		Name:       "poj",
		Version:    "1",
		Int64IO:    "%I64d",
		JavaClass:  "Main",
		Status:     "ok",
		StatusInfo: "OK",
		LastCheck:  time.Now(),
	}
	id, err := ojim.Insert(&ojinfo)
	t.Log("id: ", id)
	if err != nil {
		t.Errorf("TestOjInfo: %s", err)
	}
}
