/**
 * @Author: yinjinlin
 * @File:  goregx
 * @Description: 正则表达式
 * @Date: 2021/10/29 下午5:03
 */

package goregistry

import (
	"fmt"
	"regexp"
)

func regx() {
	email := "YY@163.com"
	// 电子邮箱合法检查
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` // 匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	emailBool := reg.MatchString(email)
	if !emailBool {
		fmt.Println("电子邮箱不合法....")
		return
	}

	mobile := "18741230098"

	// 手机号码合法检查
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
	reg = regexp.MustCompile(regular)
	mobileBool := reg.MatchString(mobile)
	if !mobileBool {
		fmt.Println("手机号码不合法....")
		return
	}

}

// 座机电话 检查
func PhoneRegx(phone string) {
	// ^[0][1-9]{2,3}-[0-9]{5,10}$

	// regxp := "^([0][1-9]{2,3}-)?[0-9]{7,8}$"

	// regxp := "/^((0\\d{2,3}-\\d{7,8})|(1[3584]\\d{9}))$/"
	regxp := "^(0[0-9]{2,3}\\-)?([2-9][0-9]{6,7})+(\\-[0-9]{1,4})?$|(^11[3|4|5|7|8]\\d{9}$)"
	reg := regexp.MustCompile(regxp)
	phnoeBool := reg.MatchString(phone)
	if !phnoeBool {
		fmt.Println("号码不合法....")
		return
	}

}

// 匹配最后字符
func Matching() {
	carType := map[int64][]string{}
	// 287762808832:[situ_t
	//est.car6 situ_test.car14]
	carType[12] = []string{"situ_test.car14", "situ_test.car3"}
	carType[10] = []string{"situ_test.car10", "situ_test.car13"}
	for _, v := range carType {
		a := "^situ_test\\.car[0-4]$"
		reg := regexp.MustCompile(a)

		for _, v1 := range v {
			str := reg.MatchString(v1)
			if str {
				fmt.Println("A")
			}
		}

	}
}

func Match() {
	a := 1
	low := 1
	if low > a/2 {
		fmt.Println(a / 2)
		fmt.Println("拥堵===")
	} else if low > a/3 {
		fmt.Println(a / 3)
		fmt.Println("缓慢===")
	} else {
		fmt.Println("不拥堵===")
	}
}
