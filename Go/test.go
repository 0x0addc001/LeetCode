package main

//import (
//	"fmt"
//)

//| 功能    | C语言                  | Go语言                             |
//| ----- | -------------------- | -------------------------------- |
//| 程序入口  | `int main()`         | `func main()`                    |
//| 输出    | `printf("Hello\n");` | `fmt.Println("Hello")`           |
//| 变量声明  | `int a = 10;`        | `a := 10`（函数内）或 `var a int = 10` |
//| 多变量赋值 | 无语法支持                | `a, b := 1, 2`                   |
//| 条件语句  | `if (x > 0)`         | `if x > 0`                       |
//| 代码块   | 必须使用 `{}`，语句后加 `;`   | `{}` 必须，语句末 **不加 `;`**           |
//| 数组    | `int arr[10];`       | `arr := [10]int{}`               |
//| 字符串   | `char* s = "hello";` | `s := "hello"`                   |
//| 注释    | `// 单行` 或 `/* 多行 */` | 相同                               |

/*
如何使用指针
指针使用流程：
1.定义指针变量。
2.为指针变量赋值。
3.访问指针变量中指向地址的值。
在指针类型前面加上 * 号（前缀）来获取指针所指向的内容。
*/

// var 数组名 [长度或...]元素类型
// var 切片名 []元素类型 （长度可变，可append）

//func main() {
//s := "gopher"
//var ptr * string
//ptr = &s
//ptr := &s
//fmt.Println("Hello World", s)
//fmt.Println("Hello World", *ptr)
//for i := 1; i < 10; i++ {
//	fmt.Println("i =", i)
//}

//var name string
//var age int
//fmt.Print("Please enter your name and age")
//fmt.Scan(&name, &age)
//fmt.Printf("Please confirm your info again: %s and %d", name, age)

//a := 1
//b := 2
//swap(&a, &b)
//fmt.Print(a, b)
//}
//func swap(a, b *int) {
//	*a, *b = *b, *a
//}

// 类似于字典
//type Book struct {
//	title   string
//	author  string
//	subject string
//	book_id int
//}

//func main() {

//// 创建一个新的结构体
//fmt.Println(Book{"Go语言", "Kyle", "Go语言教程", 6495407})
//
//// 也可以使用 key => value 格式
//fmt.Println(Book{title: "Go语言", author: "Kyle", subject: "Go语言教程", book_id: 6495407})
//
//// 忽略的字段为 0 或 空
//fmt.Println(Book{title: "Go语言", author: "Kyle"})
//}

// Go 语言中 range 关键字用于 for 循环中迭代数组(array)、切片(slice)、通道(channel)或集合(map)的元素。
// 在数组和切片中它返回元素的索引和索引对应的值，在集合中返回 key-value 对。

// 声明一个包含 2 的幂次方的切片
//var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
//
//func main() {
//	// 遍历 pow 切片，i 是索引，v 是值
//	for i, v := range pow {
//		// 打印 2 的 i 次方等于 v
//		fmt.Printf("2**%d = %d\n", i, v)
//	}
//}

//func main() {
//	for i, c := range "hello" {
//		fmt.Printf("index: %d, char: %c\n", i, c)
//	}
//	//range 也可以用在 map 的键值对上。
//	kvs := map[string]string{"a": "apple", "b": "banana"}
//	for k, v := range kvs {
//		fmt.Printf("%s -> %s\n", k, v)
//	}
//
//	str := "123"
//	num, err := strconv.Atoi(str)
//	if err != nil {
//		fmt.Println("转换错误:", err)
//	} else {
//		fmt.Printf("字符串 '%s' 转换为整数为：%d\n", str, num)
//	}
//}

// Go 的接口设计简单却功能强大，是实现 多态 和 解耦 的重要工具。
/* 定义接口 */
//type interface_name interface {
//	method_name1 [return_type]
//	method_name2 [return_type]
//	method_name3 [return_type]
//...
//method_namen [return_type]
//}
//
///* 定义结构体 */
//type struct_name struct {
//	/* variables */
//}
//
///* 实现接口方法 */
//func (struct_name_variable struct_name) method_name1() [return_type] {
///* 方法实现 */
//}
//...
//func (struct_name_variable struct_name) method_namen() [return_type] {
///* 方法实现*/
//}

//// 定义接口
//type Shape interface {
//	Area() float64
//	Perimeter() float64
//}
//
//// 定义一个结构体
//type Circle struct {
//	Radius float64
//}
//
//// Circle 实现 Shape 接口
//func (c Circle) Area() float64 {
//	return math.Pi * c.Radius * c.Radius
//}
//
//func (c Circle) Perimeter() float64 {
//	return 2 * math.Pi * c.Radius
//}
//
//func main() {
//	c := Circle{Radius: 5}
//	var s Shape = c // 接口变量可以存储实现了接口的类型
//	fmt.Println("Area:", s.Area())
//	fmt.Println("Perimeter:", s.Perimeter())
//}

//空接口 interface{} 是 Go 的特殊接口，表示所有类型的超集。

// | 特性        | Java Exception    | Go panic/recover          |
// | --------- | ----------------- | ------------------------- |
// | 可被捕获      | ✅ try-catch       | ✅ defer + recover         |
// | 语法机制      | try-catch-finally | defer + panic + recover   |
// | 是否常用于错误处理 | ✅ 是               | ❌ 否（Go 更推荐用 `error`）      |
// | 传递方式      | 抛出异常，栈展开          | panic 自动层层传递，recover 可拦截  |
// | 可以自定义类型   | ✅ 可继承 Exception 类 | ❌ panic 接收任意类型，通常用 string |

//+----------------+
//|   error        | <-- 推荐，标准错误处理机制
//+----------------+
//|
//↓
//+----------------+
//|   panic        | <-- 崩溃机制，非正常错误
//+----------------+
//|
//↓
//+----------------+
//|   recover      | <-- 拦截 panic，让程序恢复
//+----------------+

//func riskyMethod() {
//	panic("something went wrong") // 类似 throw
//}
//
//func main() {
//	defer func() {
//		if r := recover(); r != nil {
//			fmt.Println("Recovered from panic:", r)
//		}
//	}()
//
//	riskyMethod()
//}

//1. error：标准错误机制
//import (
//	"errors"
//	"fmt"
//)
//
//func divide(a, b int) (int, error) {
//	if b == 0 {
//		return 0, errors.New("cannot divide by zero")
//	}
//	return a / b, nil
//}
//
//func main() {
//	result, err := divide(10, 0)
//	if err != nil {
//		fmt.Println("Error:", err)
//		return
//	}
//	fmt.Println("Result:", result)
//}

//2. panic：非预期错误，直接中断程序
//func crash() {
//	panic("something went terribly wrong")
//}
//
//func main() {
//	fmt.Println("Start")
//	crash()
//	fmt.Println("End") // 不会被执行
//}

// 3. recover：拯救 panic，让程序不崩溃
// 必须配合 defer 使用
//func safe() {
//	defer func() {
//		if r := recover(); r != nil {
//			fmt.Println("Recovered from panic:", r)
//		}
//	}()
//
//	panic("boom!")  // 会被 recover 捕获
//	fmt.Println("This won't be printed")
//}
//
//func main() {
//	safe()
//	fmt.Println("Program continues normally")
//}

//import (
//	"fmt"
//	"time"
//)
//
//func sayHello() {
//	for i := 0; i < 5; i++ {
//		fmt.Println("Hello")
//		time.Sleep(100 * time.Millisecond)
//	}
//}
//
//func main() {
//	go sayHello() // 启动 Goroutine
//	for i := 0; i < 5; i++ {
//		fmt.Println("Main")
//		time.Sleep(100 * time.Millisecond)
//	}
//}

//// 通道（Channel）是用于 Goroutine 之间的数据传递。
//func sum(s []int, c chan int) {
//	sum := 0
//	for _, v := range s {
//		sum += v
//	}
//	c <- sum // 把 sum 发送到通道 c
//}
//
//func main() {
//	s := []int{7, 2, 8, -9, 4, 0}
//
//	c := make(chan int)
//	go sum(s[:len(s)/2], c)
//	go sum(s[len(s)/2:], c)
//	x, y := <-c, <-c // 从通道 c 中接收
//
//	fmt.Println(x, y, x+y)
//}

//并发编程小结
//Go 语言通过 Goroutine 和 Channel 提供了强大的并发支持，简化了传统线程模型的复杂性。配合调度器和同步工具，可以轻松实现高性能并发程序。
//
//Goroutines 是轻量级线程，使用 go 关键字启动。
//Channels 用于 goroutines 之间的通信。
//Select 语句 用于等待多个 channel 操作。
//
//
//死锁 (Deadlock)：
//
//示例：所有 Goroutine 都在等待，但没有任何数据可用。
//解决：避免无限等待、正确关闭通道。
//
//
//数据竞争 (Data Race)：
//
//示例：多个 Goroutine 同时访问同一变量。
//解决：使用 Mutex 或 Channel 同步访问。
