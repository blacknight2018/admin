package main

import (
	"admin/Config"
	"admin/DbModel"
	"fmt"
	"testing"
)

func TestSelectTableRecordSetCount(t *testing.T) {
	Config.GetConf()
	var condition = make(map[string]interface{})
	condition["status"] = 1
	ok, cnt := DbModel.SelectTableRecordSetCount("order", condition, nil, nil, "")
	fmt.Println(ok, cnt)
}
