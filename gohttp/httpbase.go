package gohttp

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/transform"

	"io/ioutil"
	"net/http"
)

func Get() {
	r, err := http.Get("http://httpbin.org/get")
	if err != nil {
		panic(err)
	}
	defer func() { _ = r.Body.Close() }()

	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", content)
}

func Post() {
	// body 是请求参数
	r, err := http.Post("http://httpbin.org/post", "", nil)
	if err != nil {
		panic(err)
	}
	defer func() { _ = r.Body.Close() }()

	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", content)
}

func Put() {
	request, err := http.NewRequest(http.MethodPut, "http://httpbin.org/put", nil)
	if err != nil {
		panic(err)
	}
	r, err := http.DefaultClient.Do(request) // enter 键
	if err != nil {
		panic(err)
	}

	defer func() { _ = r.Body.Close() }()

	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", content)
}

func Del() {
	request, err := http.NewRequest(http.MethodDelete, "http://httpbin.org/delete", nil)
	if err != nil {
		panic(err)
	}
	r, err := http.DefaultClient.Do(request) // enter 键
	if err != nil {
		panic(err)
	}

	defer func() { _ = r.Body.Close() }()

	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", content)
}

func Encoding() {
	r, err := http.Get("https://www.baidu.com")
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = r.Body.Close()
	}()

	//可以通过网页的头部测试编码信息
	bufReader := bufio.NewReader(r.Body)
	bytes, _ := bufReader.Peek(1024) //peek 不会移动只读取位置
	e, _, _ := charset.DetermineEncoding(bytes, r.Header.Get("content-type"))

	fmt.Println(e)
	//
	bodyReader := transform.NewReader(bufReader, e.NewDecoder())
	content, _ := ioutil.ReadAll(bodyReader)
	fmt.Printf("%s", content)
}
