package tools

import (
	"bytes"
	"github.com/gogf/gf/os/glog"
	"reflect"
	"text/template"
)

// IsContains 查找值val是否在数组array中存在
func IsContains(val interface{}, array interface{}) bool {
	if array == nil {
		return false
	}
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(array)
		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				return true
			}
		}
	}
	return false
}

//通过文本模板进行变量替换
func StringLiteralTemplate(str string, param interface{}) string {
	t, err := template.New("test").Parse(str)
	if err != nil {
		glog.Fatal("Parse string literal template error:", err)
	}
	buf := new(bytes.Buffer) //读写方法的可变大小的字节缓冲
	err = t.Execute(buf, param)
	if err != nil {
		glog.Fatal("Execute string literal template error:", err)
		return ""
	}
	return buf.String()
}
