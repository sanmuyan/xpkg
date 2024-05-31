package xutil

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
	p := []byte(parentPath)
	s := []byte(subPath)
	if p[0] != 47 || s[0] != 47 || len(p) >= len(s) {
		return false
	}
	j := 0
	for index := range p {
		if len(s) == index {
			break
		}
		if p[index] == s[index] {
			j = index
		}
	}
	// 获取两个路径相等的部分
	// p 相等的部分以 / 结尾  例子 /api/ /api/user
	// s 相等的部分下一个元素等于 / 例子 /api /api/user
	// s 相等的部分下一个元素等于 s 例子 /api/user /api/users
	if p[j] == 47 || s[j+1] == 47 || s[j+1] == 115 {
		return true
	}
	return false
}
