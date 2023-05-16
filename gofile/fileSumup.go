/*
 * @Author: your name
 * @Date: 2021-03-08 13:06:11
 * @LastEditTime: 2021-03-23 10:51:27
 * @LastEditors: your name
 * @Description: In User Settings Edit
 * @FilePath: /Goland/file/fileSumup.go
 */
/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2021/2/7 17:25
@Description:

*********************************************/
package gofile

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
)

// ==== bufio 总结 =====

// IoFileRead  从外部文件或终端读取文件到内存中或文件中
// ReadString : 是将外部文件读到（应用程序中）输入流中
func IoFileRead(name string) {
	file, _ := os.Open(name)
	reader := bufio.NewReader(file)
	defer file.Close()
	// 每次读以delim 作为分隔符字符串
	for {
		str, err := reader.ReadString('\n')
		// 读到文件末位的时候，文件为 io.EOF 标识符，则退出
		if err == io.EOF {
			break
		}
		fmt.Println(str)
	}
}

// IoFileWrite 写入到文件
func IoFileWrite(name string) {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println("文件创建失败")
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	n, err := writer.Write([]byte("hello world "))
	//n, err := writer.WriteString("func LimitReader(r Reader, n int64) Reader11209871")
	// 注意：
	// 因为writer是带缓存，因此在调用WriterString方法时，其实
	// 内容是先写入到缓存的,所以需要调用Flush方法，将缓冲的数据
	// 真正写入到文件中， 否则文件中会没有数据!!!
	err = writer.Flush()
	if err != nil {
		panic(err)
	}
	fmt.Println(n)
}

// ==== io/ioutil 总结 ======

// IoFileCopy 利用io.Copy文件拷贝
func IoFileCopy(dstName, srcName string) {
	// open  openFile 区别
	//  os.Open 打开的文件只有读权限，
	//  os.OpenFile 可以添加写权限
	srcFile, _ := os.Open(srcName)
	defer srcFile.Close()
	srcRead := bufio.NewReader(srcFile)

	dstFile, _ := os.OpenFile(dstName, os.O_RDWR|os.O_APPEND, 0777)
	defer dstFile.Close()
	dstWrite := bufio.NewWriter(dstFile)

	io.Copy(dstWrite, srcRead)

}

func IoutilFileRead(name string) {
	file, err := os.OpenFile(name, os.O_RDWR, 0777)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	// 一次性读取所有文件
	by, _ := ioutil.ReadAll(file)
	fmt.Println(string(by))

}

// FileStream 输出流： 将应用程序中数据写如到外部资源中
func FileStream(name string) {
	// 1：读取文件
	// 构建字节切片的时候如果需要读取文件,不要使用os.Open,该方法获取的文件只能读取，无法写入
	// name:文件名，flag:操作方， perm:文件权限
	// fileName 文件
	file, err := os.OpenFile(name, os.O_RDWR, 0777)

	if err != nil {
		fmt.Println("文件不存在，正在创建....")
		file, _ = os.Create("fileread.txt")
	}
	defer file.Close()

	// 向输出流中写入数据
	_, _ = file.Write([]byte("要写入的内容1\\r\\n要写入的内容2\\r\\n\\t要写入的内容3\\r\\n"))
	_, _ = file.Write([]byte("\n"))
	_, _ = file.WriteString("==写入这一行......")
	_, _ = file.WriteAt([]byte("HHHHHHH"), 2) // 此方法与追加不用混用。

}

func Dir() {
	var dir string = string("./test/case")
	finfo, _ := ioutil.ReadDir(dir)
	for _, v := range finfo {
		filenameWithSuffix := path.Base(v.Name())
		fmt.Println(filenameWithSuffix)
	}

}

// ReadFileByIo ioutil 文件读写
func ReadFileByIo() {
	inputFile := "./readwritedata/readWrite.txt"
	outputFile := "./readwritedata/readWrite2.txt"
	// ReadFile 返回字节切片，其被存入 buf 中
	buf, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Println(err)
	}
	// 将字节切片转成字符串类型
	fmt.Printf("%s\n", string(buf))
	// 输出到其他文件中，若无文件，则创建该文件
	err = ioutil.WriteFile(outputFile, buf, 0644) // oct, not hex
	if err != nil {
		panic(err.Error())
	}
}
