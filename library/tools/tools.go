package tools

import (
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/gconv"
	"reflect"
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

//communicateId

func CreateCommunicateId(fromUser, toUser int) string {

	var commId string = gconv.String(fromUser) + "_" + gconv.String(toUser)
	if fromUser > toUser {
		commId = gconv.String(toUser) + "_" + gconv.String(fromUser)
	}
	glog.Info(commId)
	commId, _ = gmd5.Encrypt(commId)
	return commId
}
