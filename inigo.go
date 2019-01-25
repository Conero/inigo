package inigo

// @Date：   2018/8/19 0019 10:58
// @Author:  Joshua Conero
// @Name:    库主文件

// get new Parser
// param format(single param)
// 		opts map[string]string{}|string
//			driver SupportNameRong SupportNameIni
// default BaseParser
func NewParser(params ...interface{}) Parser {
	var driver string
	var opts map[string]interface{}
	if params == nil {
		return new(BaseParser)
	}else {
		paramsLen := len(params)
		if driverTmp, isStr := params[0].(string); isStr {
			driver = driverTmp
		}

		if optsTmp, isOpt := params[0].(map[string]interface{}); isOpt && driver == "" {
			opts = optsTmp
			if driverTmp, isset := opts["driver"]; isset {
				driver = driverTmp.(string)
			}
		}

		if paramsLen > 1{
			if driverTmp, isStr := params[1].(string); isStr {
				driver = driverTmp
			}
		}
	}

	switch driver {
	case SupportNameRong:
		return new(RongParser)
	case SupportNameToml:
		return new(TomlParser)
	default:
		return new(BaseParser)
	}
	return nil
}
