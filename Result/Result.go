package Result

import "encoding/json"

type Result struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func (r Result) Get() string {
	var ret string
	r.Msg = code2Msg(r.Code)
	if bytes, err := json.Marshal(r); err == nil {
		ret = string(bytes)
	}
	return ret
}
