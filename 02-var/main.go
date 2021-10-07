package main

import "fmt"

var gA = 100

func main() {
  var a int
  fmt.Println("a =", a);
  fmt.Printf("type of a = %T\n",  a);

  var b int = 100
  fmt.Println("b =", b);
  fmt.Printf("type of b = %T\n",  b);

  var bb string = "abcd"
  fmt.Printf("bb = %s, type of bb = %T\n", bb, bb);


  var c = 100
  fmt.Printf("c = %d, type of c is %T\n", c, c)

  var cc = "abcd"
  fmt.Printf("cc = %s, type of cc is %T\n", cc, cc)

  e := 100
  fmt.Printf("e = %d, type of e is %T\n", e, e)


  f := "abcd"
  fmt.Printf("f = %s, type of f is %T\n", f, f)

  g := 3.14
  fmt.Printf("g = %f, type of g is %T\n", g, g)

  var xx,yy int = 100,200
  fmt.Println("xx =", xx, "yy =", yy);
  var kk,ll = 100, "Aceld"
  fmt.Println("kk =", kk, "ll =", ll);

  var(
    vv int = 100
    jj bool = true
  )
  fmt.Println("vv =", vv, "jj =", jj)



}
