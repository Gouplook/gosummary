/**
 * @Author: yinjinlin
 * @File:  fileBufio.go
 * @Description:
 * @Date: 2022/1/20 下午4:47
 */

package gofile

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

func FileBufio() {
	// 带有缓冲区的
	file, _ := os.Open("./fileread.txt")
	inputReader := bufio.NewReader(file)

	defer file.Close()

	fmt.Println("Please enter some input:....")
	// 按下回车键后停止
	input, err := inputReader.ReadString('\n')
	if err == nil {
		fmt.Printf("The input was: %s\n", input)
	}
}

func Dir() {
	var dir string = string("./test/case")
	finfo, _ := ioutil.ReadDir(dir)
	for _, v := range finfo {
		filenameWithSuffix := path.Base(v.Name())
		fmt.Println(filenameWithSuffix)
	}

}
