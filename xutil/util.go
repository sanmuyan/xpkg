package xutil

import "regexp"

// xutil 不允许使用第三方库

// PtrTo 把值转为指针
func PtrTo[T any](v T) *T {
	return &v
}

// ComparableDeduplication 数组去重
func ComparableDeduplication[T comparable](l []T) []T {
	uniqueMap := make(map[T]struct{})
	var result []T
	for _, item := range l {
		if _, ok := uniqueMap[item]; !ok {
			uniqueMap[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

// IsSubPath 判断是否为子路径
func IsSubPath(parentPath, subPath string) bool {
	pl := len(parentPath)
	sl := len(subPath)
	if pl == 0 || sl == 0 || pl >= sl {
		return false
	}
	if parentPath[0] != '/' || subPath[0] != '/' {
		return false
	}
	if subPath[0:pl] == parentPath {
		// /api /api/user
		if subPath[pl] == '/' || parentPath[pl-1] == '/' {
			return true
		}
		// /api/user /api/users
		if subPath[pl] == 's' && sl == pl+1 {
			return true
		}
	}
	return false
}

// IsUsername 判断是否为用户名
func IsUsername(s string) bool {
	if ok := regexp.MustCompile(`^[a-z][a-z\d]{3,29}$`).MatchString(s); !ok {
		return false
	}
	return true
}
