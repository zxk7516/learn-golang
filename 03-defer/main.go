package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)


func sum(n1 , n2 int) int {
  // defer 在return 之后立刻执行
  // defer 被压入 stack 中, 先入后出
  // defer 将引用的变量进行值拷贝


  n1++
  n2++


  defer fmt.Println("ok1 n1=",n1);
  defer fmt.Println("ok2 n2=",n2);

  n1++
  n2++


  res := n1 + n2
  fmt.Println("ok 3 res=",res)
  return res
}

func main() {
  res := sum(1,2)
  fmt.Println("final res=",res)

  test()
}


func test() {
  file,err := os.Open("a.txt")
  if err != nil {
    fmt.Println("file open fialed")
    return
  }
  defer file.Close()

  f2,err := ioutil.ReadFile("main.go")
  if err != nil {
    fmt.Println("file2 open fialed")
    return
  }
  str := string(f2)
  fmt.Println(str)

   var lines []string
   scanner := bufio.NewScanner(file)
   for scanner.Scan() {
     lines = append(lines, scanner.Text())
   }
   fmt.Println("lines: ", lines)

   outfile, err := os.Create("file.txt")
   if err != nil {
     fmt.Println("create file.txt fialed")
     return
   }
   defer outfile.Close()

   var a = []string {"Rio", "Luben", "Milan"}
   for i:=0;i<len(a);i++ {
     outfile.WriteString(a[i])
     outfile.WriteString("\n")
   }
}
