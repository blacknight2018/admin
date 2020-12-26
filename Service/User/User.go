package User

import (
	"admin/DbModel"
	"admin/Result"
)

func QueryUser(nickName string, limit int, offset int) Result.Result {
	var ret Result.Result
	ret.Code = Result.UnKnow
	type name struct {
		Total int            `json:"total"`
		User  []DbModel.User `json:"user"`
	}
	var retData name
	if ok, data := DbModel.SelectUserSetByNickName(nickName, limit, offset); ok {
		retData.User = data
		retData.Total = DbModel.SelectUserSetCountByNickName(nickName, nil, nil)
		ret.Code = Result.Ok
		ret.Data = retData
	}
	return ret
}
