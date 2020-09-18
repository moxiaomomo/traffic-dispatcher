package util

import (
	"strconv"

	api "github.com/micro/go-micro/v2/api/proto"
)

// ReqParamProp 请求参数属性
type ReqParamProp struct {
	// 参数名
	Name string
	// 是否必须传递
	Required bool
	// 类型 (int / string)
	Type string
}

// TryGetParam 获取指定参数的值
func TryGetParam(req *api.Request, param string) (string, bool) {
	pair, ok := req.Get[param]
	if !ok || len(pair.Values) == 0 {
		return "", false
	}
	return pair.Values[0], true
}

// TryGetAllParams 获取指定参数值list，有一个必选的参数不满足要求时则返回空列表
func TryGetAllParams(req *api.Request, params []ReqParamProp) ([]interface{}, bool) {
	var values []interface{}

	for _, param := range params {
		pair, ok := req.Get[param.Name]
		// logger.Infof("%+v %+v\n", pair, param)
		if !ok || len(pair.Values) == 0 {
			if param.Required {
				return []interface{}{}, false
			}
			values = append(values, parseValue("", param.Type))
		}
		values = append(values, parseValue(pair.Values[0], param.Type))
	}

	return values, true
}

func parseValue(val string, tp string) interface{} {
	switch tp {
	case "int":
		res, err := strconv.Atoi(val)
		if err != nil {
			return 0
		}
		return res
	case "string":
		return val
	}
	return val
}
