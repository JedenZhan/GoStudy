##　函数

语法:

```go
package main

func main() {
    
}

// 定义方式: func funcName(参数 参数类型, 第二个参数 第二个参数的类型) 返回值类型 {
// 	函数体
// }
func myFunc(a float64, b float64) float64 {
    // 函数体
    return a * b
}
```

函数可以返回多个值

```go
// ...
func main() {
    a, b = myFunc(1, 2)
}
func myFunc(a int, b int) (int, int) { // 多个返回值也要挨个声明类型
    return a, b
}

```

### 函数可以和Struct结合(我们称为methods)

e.g.

```go
func main() {
    type Rectangle struct {
        width, height float64
    }
    
}
func myFunc(r Rectangle) float64 {
    return r.width * r.height
}
```
