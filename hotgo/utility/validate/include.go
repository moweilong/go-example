package validate

import "github.com/gogf/gf/v2/util/gconv"

// InSliceExistStr 判断字符或切片字符是否存在指定字符
func InSliceExistStr(elems any, search string) bool {
	switch elems.(type) {
	case []string:
		elem := gconv.Strings(elems)
		for i := 0; i < len(elem); i++ {
			if gconv.String(elem[i]) == search {
				return true
			}
		}
	default:
		return gconv.String(elems) == search
	}
	return false
}
