package xutil

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"runtime"
	"strings"
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

// StructDeduplication 结构体类型数组去重
func StructDeduplication[T comparable](l []T, uqField string) []T {
	uniqueMap := make(map[string]T)
	var result []T
	for _, item := range l {
		uqIndexValue := reflect.ValueOf(item).FieldByName(uqField)
		if uqIndexValue.IsNil() {
			return l
		}
		uqValue := fmt.Sprintf("%v", uqIndexValue)
		if _, ok := uniqueMap[uqValue]; !ok {
			uniqueMap[uqValue] = item
		}
	}
	for _, item := range uniqueMap {
		result = append(result, item)
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
	copyFunc := func(i int, ste reflect.Type, sfv, tfv reflect.Value) {
		tag := ste.Field(i).Tag.Get("copy")
		if tag == "ignore" {
			return
		}
		if !IsZeroOfRefVal(sfv) || sfv.Kind() == reflect.Bool || tag == "force" {
			tfv.Set(sfv)
			return
		}
	}
	sv := reflect.ValueOf(s)
	tv := reflect.ValueOf(t)
	if sv.Kind() != reflect.Ptr || tv.Kind() != reflect.Ptr {
		return errors.New("s and t must be pointer")
	}
	if sv.IsNil() || tv.IsNil() {
		return errors.New("s and t must not be nil")
	}
	if sv.Type().String() != tv.Type().String() {
		return errors.New("s and t must be same type")
	}
	sve := sv.Elem()
	tve := tv.Elem()
	if sve.Kind() != reflect.Struct {
		tve.Set(sve)
		return nil
	}
	ste := reflect.TypeOf(t).Elem()
	for i := 0; i < ste.NumField(); i++ {
		sfv := sve.Field(i)
		tfv := tve.Field(i)
		switch sfv.Kind() {
		case reflect.Struct:
			if err := FillObj(sfv.Addr().Interface(), tfv.Addr().Interface()); err != nil {
				return err
			}
		case reflect.Ptr:
			if !sfv.IsNil() && !tfv.IsNil() {
				if sfv.Elem().Kind() == reflect.Struct {
					if err := FillObj(sfv.Interface(), tfv.Interface()); err != nil {
						return err
					}
					continue
				}
			}
			copyFunc(i, ste, sfv, tfv)
		default:
			copyFunc(i, ste, sfv, tfv)
		}
	}
	return nil
}

// Remove 删除文件或文件夹，支持通配符
func Remove(f string) error {
	f = filepath.Clean(f)
	if strings.HasSuffix(f, "*") {
		fPrefix := filepath.Join(strings.TrimSuffix(f, "*"))
		basePath := filepath.Dir(fPrefix)
		if strings.HasSuffix(f, "/*") {
			basePath = fPrefix
		}
		if runtime.GOOS == "windows" {
			if strings.HasSuffix(f, "\\*") {
				basePath = fPrefix
			}
		}
		entries, err := os.ReadDir(basePath)
		if err != nil {
			return err
		}
		for _, entry := range entries {
			full := filepath.Join(basePath, entry.Name())
			if strings.HasPrefix(full, fPrefix) {
				err := os.RemoveAll(full)
				if err != nil {
					return err
				}
			}
		}
		return nil
	}
	return os.RemoveAll(f)
}
