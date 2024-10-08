package response

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/goal/app/enums"
	"github.com/goal-web/goal/app/results"
)

func ParseReqErr(err error) any {
	return results.ResponseResult{
		Message:    err.Error(),
		Code:       int32(enums.CodeParseReqErr),
		ErrMessage: enums.CodeParseReqErr.Message(),
		Data:       nil,
	}
}

func InvalidReq(err error) any {
	return results.ResponseResult{
		Message:    err.Error(),
		Code:       int32(enums.CodeParseReqErr),
		ErrMessage: enums.CodeParseReqErr.Message(),
		Data:       nil,
	}
}

func BizErr(err error) any {
	return results.ResponseResult{
		Message:    err.Error(),
		Code:       int32(enums.CodeBizErr),
		ErrMessage: enums.CodeBizErr.Message(),
		Data:       nil,
	}
}

func Success(data any) any {
	result := results.ResponseResult{
		Message:    "",
		Code:       int32(enums.CodeSuccess),
		ErrMessage: enums.CodeSuccess.Message(),
	}

	if fieldsProvider, ok := data.(contracts.FieldsProvider); ok {
		result.Data = fieldsProvider.ToFields()
	} else {
		result.Data = data
	}

	return result
}
