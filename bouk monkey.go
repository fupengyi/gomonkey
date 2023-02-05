package gomonkey

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"reflect"
	"strings"
)

# Go monkeypatching 🐵 🐒
Go 的实际任意猴子修补。对真的。
阅读此博文以了解其工作原理：https://bou.ke/blog/monkey-patching-in-go/

## 我认为 Go 中的 monkeypatching 是不可能的？
通过常规语言构造是不可能的，但我们总是可以让计算机屈服于我们的意志！ Monkey 通过在运行时重写正在运行的可执行文件并插入到您想要调用的函数的跳
转来实现 monkeypatching。这听起来很不安全，我不建议任何人在测试环境之外这样做。

如果你打算使用这个库，请确保你阅读了 README 底部的注释。

## Using monkey
Monkey 的 API 非常简单直接。调用 monkey.Patch(<target function>, <replacement function>) 来替换一个函数。例如：
package main

import (
	"fmt"
	"os"
	"strings"
	
	"bou.ke/monkey"
)

func main() {
	monkey.Patch(fmt.Println, func(a ...interface{}) (n int, err error) {
		s := make([]interface{}, len(a))
		for i, v := range a {
			s[i] = strings.Replace(fmt.Sprint(v), "hell", "*bleep*", -1)
		}
		return fmt.Fprintln(os.Stdout, s...)
	})
	fmt.Println("what the hell?") // what the *bleep*?
}

然后您可以调用 monkey.Unpatch(<target function>) 再次取消修补该方法。替换函数可以是任何函数值，无论是匿名函数、绑定函数还是其他函数。

如果要修补实例方法，则需要使用 monkey.PatchInstanceMethod(<type>, <name>, <replacement>)。您可以使用 reflect.TypeOf 获取类型，您
的替换函数只需将实例作为第一个参数。要禁用所有网络连接，您可以执行以下操作，例如：
package main

import (
	"fmt"
	"net"
	"net/http"
	"reflect"
	
	"bou.ke/monkey"
)

func main() {
	var d *net.Dialer // Has to be a pointer to because `Dial` has a pointer receiver	// 必须是指向的指针，因为 `Dial` 有一个指针接收器
	monkey.PatchInstanceMethod(reflect.TypeOf(d), "Dial", func(_ *net.Dialer, _, _ string) (net.Conn, error) {
		return nil, fmt.Errorf("no dialing allowed")
	})
	_, err := http.Get("http://google.com")
	fmt.Println(err) // Get http://google.com: no dialing allowed
}
请注意，目前无法只为一个实例修补该方法，PatchInstanceMethod 将为所有实例修补它。不要费心尝试 monkey.Patch(instance.Method, replacement)
，它不会起作用。 monkey.UnpatchInstanceMethod(<type>, <name>) 将撤消 PatchInstanceMethod。

如果你想删除所有当前应用的 monkeypatches 只需调用 monkey.UnpatchAll。这在测试拆卸功能中可能很有用。

如果你想从替换中调用原始函数，你需要使用 monkey.PatchGuard。 patchguard 允许您轻松删除和恢复补丁，以便您可以调用原始函数。例如：
package main

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"
	
	"bou.ke/monkey"
)

func main() {
	var guard *monkey.PatchGuard
	guard = monkey.PatchInstanceMethod(reflect.TypeOf(http.DefaultClient), "Get", func(c *http.Client, url string) (*http.Response, error) {
		guard.Unpatch()
		defer guard.Restore()
		
		if !strings.HasPrefix(url, "https://") {
			return nil, fmt.Errorf("only https requests allowed")
		}
		
		return c.Get(url)
	})
	
	_, err := http.Get("http://google.com")
	fmt.Println(err) // only https requests allowed
	resp, err := http.Get("https://google.com")
	fmt.Println(resp.Status, err) // 200 OK <nil>
}

## Notes
1.如果启用内联，Monkey 有时无法修补函数。尝试在禁用内联的情况下运行测试，例如：go test -gcflags=-l。相同的命令行参数也可用于构建。
2.Monkey 无法在某些不允许同时写入和执行内存页面的面向安全的操作系统上运行。目前的方法并没有真正可靠的解决方法。
3.Monkey 不是线程安全的。或者任何类型的保险箱。
4.我在 OSX 10.10.2 和 Ubuntu 14.04 上测试了 monkey。它应该适用于任何基于 unix 的 x86 或 x86-64 系统。
