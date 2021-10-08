package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"
)

func main() {
	readFileWithBuffer()
	readFileAll()
	writeFileWithBuffer()
	writeFileTrunc()
	appendLogFiles()
	copyFileContents()
	fileStatics()

}

type TextFileStatics struct {
	AlphaCount int //
	NumCount   int //
	SpaceCount int //
	OtherCount int
}

func fileStatics() {
	file, err := os.Open("abc.txt")
	if err != nil { // 文件存在
		fmt.Println("error open file")
		return
	}
	var statics TextFileStatics
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		for _, v := range str {
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough
			case v >= 'A' && v <= 'A':
				statics.AlphaCount++
			case v == ' ' || v == '\t':
				statics.SpaceCount++
			case v >= '0' && v <= '9':
				statics.NumCount++
				break
			default:
				statics.OtherCount++
				break
			}

		}
	}
	fmt.Println("statics", statics)
}

func fileExiss(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil { // 文件存在
		return true, nil
	}
	if os.IsNotExist(err) { // 文件不存在
		return false, nil
	}
	return false, err // 其它错误
}

func copyFileContents() {

	srcFile, err := os.Open("abc.txt")
	if err != nil {
		return
	}
	defer srcFile.Close()
	dstFile, err := os.OpenFile("def.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return
	}
	defer dstFile.Close()

	writer := bufio.NewWriter(dstFile)
	defer writer.Flush()

	l, err := io.Copy(dstFile, srcFile)
	if err != nil {
		fmt.Println("")
		return
	}
	fmt.Println("l", l)
}

func appendLogFiles() {
	str := time.Now().Format("2006-12-31 23:59:59")
	fmt.Println(str)
	filepath := "logs.txt"
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("error open file")
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	writer.WriteString("[")
	writer.WriteString(str)
	writer.WriteString("] logs writting file\n")
	writer.Flush()

}

// 覆盖原来的文件 ,os.O_TRUNC
func writeFileTrunc() {
	str := time.Now().Format("2006-01-01 15:04:03")
	fmt.Println(str)
	filepath := "timestamp.txt"
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("error open file")
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	writer.WriteString(str)
	writer.Flush()
}

func writeFileWithBuffer() {
	file, err := os.OpenFile("b.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("error open file")
		return
	}
	defer file.Close()
	str := "Hello, Gargon\n"
	// write string with buffer writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	writer.Flush()

}

func readFileAll() {
	str, err := ioutil.ReadFile("a.txt")
	if err != nil {
		fmt.Println("read file a.txt error")
		return
	}
	fmt.Println(string(str))

}

func readFileWithBuffer() {
	var err error
	file, err := os.Open("a.txt")
	if err != nil {
		fmt.Println("open file a.txt error")
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	var line string
	for {
		line, err = reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(line)
	}
	fmt.Println("文件读取结束")
}

func readFile1() {
	file, err := os.Open("a.txt")
	if err != nil {
		fmt.Println("open file a.txt error")
		return
	}
	defer file.Close()
	fmt.Printf("file=%v\n", file)
	bytes := make([]byte, 1024)
	file.Read(bytes)
	fmt.Println(string(bytes))
}
