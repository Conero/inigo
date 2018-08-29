package inigo

// @Date：   2018/8/19 0019 10:58
// @Author:  Joshua Conero
// @Name:    库主文件

// opts map[string]string{}, driver
func NewParser(params ...interface{}) Parser {
	if params == nil{
		return new(BaseParser)
	}
	opts := map[string]string{}
	if params[0] != nil{
		opts = params[0].(map[string]string)
	}
	driver := ""
	if len(params) > 1{
		driver = params[1].(string)
	}else if tDrv, has := opts["driver"]; has{
		driver = tDrv
	}

	switch driver {
	case SupportNameRong:
		return new(RongParser)
	default:
		return new(BaseParser)
	}
	
	return nil
}
