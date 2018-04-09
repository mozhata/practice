package merr

import "strings"

// 错误码分类：1xx(参数错误)；2xx(操作错误); 3xx(外部依赖错误); 4xx(内部错误)

const (
	OK                     = 0
	AccountBindFailed      = 100100
	InvalidParam           = 100101
	AccountInvalidSource   = 100102
	AccountAccountNotExist = 100103
	InternalError          = 100500
)

var ErrorMap = map[int]map[string]string{
	0: map[string]string{
		"EN-US": "OK",
		"ZH-CN": "成功",
	},
	100100: map[string]string{
		"EN-US": "Account Bind Failed",
		"ZH-CN": "参数解析错误",
	},
	100101: map[string]string{
		"EN-US": "Invalid Param",
		"ZH-CN": "参数不合法",
	},
	100102: map[string]string{
		"EN-US": "Account Invalid Source",
		"ZH-CN": "无效来源",
	},
	100103: map[string]string{
		"EN-US": "Account Account Not Exist",
		"ZH-CN": "账户不存在",
	},

	100500: map[string]string{
		"EN-US": "Internal Errort",
		"ZH-CN": "内部错误",
	},
}

func GetMsg(code int, languages []string) string {
	msgMap, ok := ErrorMap[code]
	if !ok {
		return "Unknown error"
	}
	for _, lang := range languages {
		if msg, ok := msgMap[strings.ToUpper(lang)]; ok {
			if msg != "" {
				return msg
			}
		}
	}
	return "Unknown error"
}
