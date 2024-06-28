package xutil

import (
	"errors"
	"reflect"
	"regexp"
)

// xutil 不允许使用第三方库

// PtrTo 把值转为指针
func PtrTo[T any](v T) *T {
	return &v
}

// Deduplication 数组去重
func Deduplication[T comparable](l []T) []T {
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

// IsZero 判断可比较类型否为非零
func IsZero[T comparable](vals ...T) bool {
	if len(vals) == 0 {
		return true
	}
	var z T
	for _, v := range vals {
		if v == z {
			return true
		}
	}
	return false
}

// IsContains 判断是否为数组成员
func IsContains[T comparable](v T, vals []T) bool {
	if len(vals) == 0 {
		return false
	}
	m := make(map[T]struct{}, len(vals))
	for _, _v := range vals {
		m[_v] = struct{}{}
	}
	_, ok := m[v]
	if ok {
		return true
	}
	return false
}

// RemoveError 去除错误返回
func RemoveError[T any](data T, err error) T {
	return data
}

// IsZeroOfRefVal 判断反射类型是否为零值
func IsZeroOfRefVal(vals ...reflect.Value) bool {
	if len(vals) == 0 {
		return true
	}
	for _, v := range vals {
		if reflect.DeepEqual(v.Interface(), reflect.Zero(v.Type()).Interface()) {
			return true
		}
	}
	return false
}

// FillObj 使用 s 填充 t 对象但排除零值
func FillObj(s, t any) error {
	sv := reflect.ValueOf(s)
	tv := reflect.ValueOf(t)
	if sv.Kind() != reflect.Ptr || tv.Kind() != reflect.Ptr {
		return errors.New("s and t must be pointer")
	}
	sve := sv.Elem()
	tve := tv.Elem()
	if sve.Kind() != tve.Kind() || sv.Type() != tv.Type() {
		return errors.New("s and t must be same type")
	}
	for i := 0; i < tve.NumField(); i++ {
		v := sve.Field(i)
		if !IsZeroOfRefVal(v) || v.Type().Kind() == reflect.Bool {
			tve.Field(i).Set(sve.Field(i))
		}
	}
	return nil
}
