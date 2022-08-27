// 和headfirstgo模块配套
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gogogo1024/headfirstgo"
)

// pass_fail reports whether a grade is passing
func main() {
	myFloatPointer := headfirstgo.CreatePoint()
	fmt.Println(myFloatPointer)
	fmt.Println(*myFloatPointer)
	amount := 6
	headfirstgo.DoubleNumber(&amount)
	fmt.Println(amount)

	headfirstgo.Hello()
	headfirstgo.Hi()
	headfirstgo.ReadFile("./data.txt")

	// slice切片就是数组的视图
	var arrayInt [3]int
	var sliceInt []int
	fmt.Printf("arrayInt:%#v,sliceInt:%#v", arrayInt, sliceInt)
	fmt.Println()

	seconds := time.Now().UnixMilli()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println("can you guess it?")

	fmt.Println(target)

	soda := headfirstgo.Liters(2)

	water := headfirstgo.Milliliters(500)
	fmt.Printf("%0.3f liters equals %0.3f liters\n", soda, soda.ToGallons())
	fmt.Printf("%0.3f milliliters equals %0.3f gallons\n", water, water.ToGallons())

	date := headfirstgo.Date{}
	err := date.SetYear(2000)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(date.Year())

	err = date.SetMonth(10)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(date.Month())

	err = date.SetDay(10)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(date.Day())

	fmt.Println(date)

	// 直接访问Date结构体中未导出的字段会导致报错
	// unknown field year in struct literal
	// date2 := Date{year: 10, month: 13, day: 20}
	// fmt.Println(date2)

	event := headfirstgo.Event{}
	err = event.SetYear(2000)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(event.Year())

	err = event.SetTitle("xasxaxxaxaxasxasx")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(event.Title())

	err = event.SetMonth(12)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(event.Month())

	err = event.SetDay(30)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(event.Day())

	fmt.Println(event)

	mixTape := []string{"God is a girl", "Whip It"}
	var player headfirstgo.Player = headfirstgo.TapePlayer{}
	headfirstgo.PlayList(player, mixTape)
	player = headfirstgo.TapeRecord{}
	headfirstgo.PlayList(player, mixTape)

	// 通过类型断言实现了接口方法和具体类型方法上的切换
	var noiseMaker headfirstgo.NoiseMaker = headfirstgo.Robot("A robot")
	noiseMaker.MakeSound()
	// 类型断言得到具体类型
	robot, ok := noiseMaker.(headfirstgo.Robot)
	if !ok {
		log.Fatal("type assert error")
	}
	// 调用具体类型上的walk方法
	robot.Walk()

	err = headfirstgo.ComedyError("this is comedyError")
	fmt.Println(err)
	err = headfirstgo.CheckTemperature(127.1, 130.0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(headfirstgo.Meters(1.1))
	fmt.Println(headfirstgo.Centimeters(1.1))
	fmt.Println(headfirstgo.Millimeters(1.1))

	defer headfirstgo.ReportPanic()
	headfirstgo.ScanDir("./")

	headfirstgo.Count("votes.txt")

	// goroutine && channel

	pages := make(chan headfirstgo.Page)
	urls := []string{"https://baidu.com", "https://www.sina.com.cn/"}
	for _, url := range urls {
		go headfirstgo.ResponseSize(url, pages)
	}
	for i := 0; i < len(urls); i++ {
		page := <-pages
		fmt.Printf("%s:%d\n", page.URL, page.Size)
	}

	// 函数是一等公民,需要符合参数和返回值类型一致
	var myFunction func() = headfirstgo.Hello
	myFunction()
	var myFunction2 func(path string) = headfirstgo.Count
	myFunction2("votes.txt")

	err = http.ListenAndServe(":8080", nil)
	log.Fatal(err)

	// guess game
	success := false
	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("you have", 10-guesses, "guesses left.")
		fmt.Print("Make a guess: ")

		reader := bufio.NewReader(os.Stdin)
		// 两次err没有出现编译错误的原因
		// 1. 当同一个变量名在同一个作用域中被声明两次时，我们会得到一个编译错误.
		// 2. 只要短变量声明中至少有一个变量名是新的，这是允许的。新变量名被视为声明，而现有的名字被视为赋值。
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		if guess > target {
			fmt.Println("Your guess was high")
		} else if guess < target {
			fmt.Println("Your guess was low")

		} else {
			success = true
			fmt.Println("Your guess was right")
			break
		}
	}
	if !success {
		fmt.Println("sorry you did`t guess my number , It was", target)
	}
}
