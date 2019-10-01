
## 并发与并行

Go是一种并发语言, 但不是并行

并发是指一段时间内能同时完成多个任务

举例: 跑步时系鞋带, 是一种并发, 系鞋带并没有打断跑步, 只是暂停而已(需要暂停执行另一个任务的叫做任务切换)

跑步时听音乐, 也是一种并发, 二者同时进行(多个cpu核心同时控制,真正的多核并行)



所以并行与并发的关系: **`并行`只是一种实现并发的方式,另一种实现并发的方式是`任务切换`**

### 那Go是怎么实现并发的呢

答案是—> **协程**

#### 协程(goroutine)

> goroutine 是指并发执行的`函数/方法`(任务), goroutine 也可以理解为轻量级的线程, 创建goroutine的开销远远小于线程, Go里面创建成千上万个goroutine很常见

**goroute VS 线程**

- goroutine 开销更小, goroutine 只占用几kb的内存, 而且可以根据需要动态缩放, 而线程占用的空间比较大, 而且是固定不变的

- goroutine 是对线程的一层封装, 所有创建的goroutine本质上都是复用少量的若干个线程,比如可能只有一个线程, 而很多goroutine跑在这个线程里面, 分时复用, 假如其中一个 goroutine阻塞了该线程的执行, golang 会新建一个线程将其他的 goroutine 切换到新线程上面

- goroutine 通过 channel, 管道来实现 goroutine 之间的通信, channel 能解决共享内存冲突的问题, 也是相对于线程的一个优势

**开启一个 goroutine**

```go
package main

import "fmt"

func hello() {
    fmt.Println("Hello GoRoutine")
}

func main() {
    go hello() // 开启一个新的goroutine, 和main方法并行,main函数是默认goroutine, 是主线程
    fmt.Println("main function")
}
// 因为两个并发执行, 不知道谁先执行完, 所以可能有不同的结果
```
**开启多个 goroutine**

```go
package main

import (
    "fmt"
    "time"
)

func numbers() {
    for i := 1; i <= 5; i ++ {
        time.Sleep(250 * time.Millisecond)
        fmt.Printf("%d ", i)
    }
}

func alphabets() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}

func main() {
    go numbers()
    go alphabets()
    time.Sleep(2000 * time.Millisecond)
    fmt.Println("All end")
}
```

#### 管道(channel)

> channel是 goroutine 通讯的管道, 就像水管一样, 可以双向流动

**声明channel**

每个channel都需要绑定一个数据类型,这个数据类型就是channel可以传输的数据类型`chan`

show time
```go
package main

import ("fmt")

func main() {
    var a chan int
    if (a == nil) {
        fmt.Println("a未赋值")
        a = make(chan int)
        fmt.Printf("type of a 是 %T", a)
    }
}
```

