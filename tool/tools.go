/******************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2020/11/23 下午4:58

*******************************************/
package tool

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"
)

//map转数组 (提取map中的key值）
func ArrayKeys(maps map[int]interface{}) []int {
	//分析参数
	if len(maps) == 0 {
		return make([]int, 0)
	}
	var arr = make([]int, 0)
	for i, _ := range maps {
		arr = append(arr, i)
	}
	return arr
}

//map数组转数组（根据字段提取信息）
func ArrayValue2Array(field string, maps []map[string]interface{}) []int {
	//分析参数
	if len(maps) == 0 {
		return make([]int, 0)
	}
	var arr = make([]int, 0)
	for _, m := range maps {
		v, ok := m[field]
		if ok {
			if vs, p := v.(string); p {
				n, _ := strconv.Atoi(vs)
				arr = append(arr, n)
			}
			if vs, p := v.(int); p {
				arr = append(arr, vs)
			}
		}
	}
	return arr
}

//map数组转map (根据字段，在map切片中提取map）
func ArrayRebuild(field string, maps []map[string]interface{}) map[string]interface{} {
	//分析参数
	if len(maps) == 0 {
		return make(map[string]interface{}, 0)
	}
	var reMap = make(map[string]interface{})
	for _, m := range maps {
		v, ok := m[field]
		if ok {
			if vs, p := v.(int); p {
				reMap[strconv.Itoa(vs)] = m
			}
			if vs, p := v.(string); p {
				reMap[vs] = m
			}
			if vs, p := v.(float64); p {
				reMap[strconv.FormatFloat(vs, 'f', -1, 64)] = m
			}
			if vs, p := v.(float32); p {
				reMap[strconv.FormatFloat(float64(vs), 'f', -1, 64)] = m
			}
		}
	}
	return reMap
}

// 数组map排序

// 思路：1： 先定义两个容器，mapData/keys 存放数据和存放key，
//      2： 对key进行排序 sort.string
//      3： 遍历key，将key中字段所对应的值，存放起来，return

func SortsMap(field string, maps []map[string]interface{}) []map[string]interface{} {
	var mapData = make(map[string]interface{}) // map make不需要指定大小
	var keys = make([]string, 0)               // 切片make时，需要指定大小
	for _, v := range maps {
		vs := v[field]
		if vp, ok := vs.(float64); ok {
			vs = strconv.FormatFloat(vp, 'f', -1, 64)
		}
		if vp, ok := vs.(int); ok {
			vs = strconv.FormatInt(int64(vp), 10)
		}
		if vp, ok := vs.(string); ok {
			vs = vp
		}
		mapData[vs.(string)] = v
		keys = append(keys, vs.(string))
	}
	sort.Strings(keys)
	remapData := make([]map[string]interface{}, 0)
	for _, v := range keys {
		remapData = append(remapData, mapData[v].(map[string]interface{}))
	}
	return remapData
}

//把数组转换为字符串
//@param  string separator       转换分隔符
//@param  interface{}  interface 待转换数据
//@return string
// []int{1,2,3} ---> str : 1,2,3  ,[]string也适用

func ArrayString(separator string, array interface{}) (str string) {
	str = strings.Replace(strings.Trim(fmt.Sprint(array), "[]"), " ", separator, -1)
	return str
}

// 数组去重 int
func ArrayUniqueInt(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	newArr := make([]int, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if arr[i] == 0 {
			continue
		}
		if repeat == false {
			newArr = append(newArr, arr[i])
		}
	}
	return newArr
}

// 数组去重 string
func ArrayUniqueString(arr []string) []string {
	if len(arr) == 0 {
		return arr
	}
	newArr := make([]string, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if arr[i] == "" {
			continue
		}
		if repeat == false {
			newArr = append(newArr, arr[i])
		}
	}
	return newArr
}

// 字符串切割成int型数组
// str := "112,34,56,78"  ---- >  [112 34 56 78]
func StrExplode2IntArr(s string, step string) []int {
	strs := strings.Split(s, ",")
	var outData []int
	for _, v := range strs {
		if len(v) == 0 {
			continue
		}
		intv, _ := strconv.Atoi(v)
		outData = append(outData, intv)
	}
	return outData
	// 1,2,5 类型卡 --> 适合这些门店 4 5 6 9
}

// TrimRgiht
func StringsTrim(s string, cutset string) string {
	str := strings.TrimRight(s, cutset)
	return str
}

//获取字符串长度
//@param  string str 待获取长度字符串
//@return int
func Mb4Strlen(str string) int {
	str = strings.TrimSpace(str)
	if len(str) == 0 {
		return 0
	}
	strRune := []rune(str)
	lens := len(strRune)
	return lens
}

//截取字符串
//@param string str   待截取的字符串
//@param int    index 截取开始位置
//@param int    lens  截取长度
func StuffStr(str string, index int, lens int) string {
	str = strings.TrimSpace(str)
	if len(str) == 0 {
		return str
	}
	strRune := []rune(str)
	if len(strRune) < lens {
		lens = len(strRune)
	}
	return string(strRune[index:lens])

}

// 公钥转换
func GetPemPublic(public_key string) string {
	res := "-----BEGIN PUBLIC KEY-----\n"
	strlen := len(public_key)
	for i := 0; i < strlen; i += 64 {
		if i+64 >= strlen {
			res += public_key[i:] + "\n"
		} else {
			res += public_key[i:i+64] + "\n"
		}
	}
	res += "-----END PUBLIC KEY-----"
	return res
}

// 私钥转换
func GetPemPrivate(private_key string) string {
	res := "-----BEGIN RSA PRIVATE KEY-----\n"
	strlen := len(private_key)
	for i := 0; i < strlen; i += 64 {
		if i+64 >= strlen {
			res += private_key[i:] + "\n"
		} else {
			res += private_key[i:i+64] + "\n"
		}
	}
	res += "-----END RSA PRIVATE KEY-----"
	return res
}

//把数组转换为字符串
//@param  string separator       转换分隔符
//@param  interface{}  interface 待转换数据
//@return string
func Implode(separator string, array interface{}) string {
	return strings.Replace(strings.Trim(fmt.Sprint(array), "[]"), " ", separator, -1)
}

//反射Model、初始化字段名称
func ReflectModel(structPtr interface{}) {
	rType := reflect.TypeOf(structPtr)
	rVal := reflect.ValueOf(structPtr)
	if rType.Kind() == reflect.Ptr {
		rType = rType.Elem()
		rVal = rVal.Elem()
	} else {
		panic("structPtr must be pointer struct.")
	}
	for i := 0; i < rType.NumField(); i++ {
		t := rType.Field(i)
		f := rVal.Field(i)
		key := t.Tag.Get("default")
		if key == "" {
			f.Set(reflect.ValueOf(""))
		} else {
			f.Set(reflect.ValueOf(key))
		}
	}
}

//生成随机数
func GenRandom(randomLen, randomKey int) (random string) {
	n63, _ := strconv.Atoi("1" + strings.Repeat("0", randomLen))
	s := fmt.Sprintf("%"+strconv.Itoa(randomLen)+"v%d", rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(int64(n63)), randomKey)
	random = strings.Trim(s, " ")
	return random[0:randomLen]
}
