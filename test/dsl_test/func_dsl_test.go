package dsltest

import (
	"testing"

	. "github.com/agiledragon/gomonkey/v2/test/fake"

	. "github.com/agiledragon/gomonkey/v2"
	. "github.com/agiledragon/gomonkey/v2/dsl"
	. "github.com/smartystreets/goconvey/convey"
)

/*
	func Convey(items ...interface{})

Convey 是声明规范范围时使用的方法。每个作用域都有一个描述和一个 func()，它可能包含对 Convey()、Reset() 或 Should 式断言的其他调用。只要
您认为合适，Convey 调用就可以嵌套。

重要说明：测试方法中的顶级 Convey() 必须符合以下签名：

	Convey(description string, t *testing.T, action func())

所有其他调用应如下所示（无需传入 *testing.T）：

	Convey(description string, action func())

别担心，如果你弄错了，goconvey 会 panic，所以你可以修复它。
此外，您可以通过执行以下操作显式获得对 Convey 上下文的访问权限：
*/
func TestPbBuilderFunc(t *testing.T) {
	Convey("TestPbBuilderFunc", t, func() {

		Convey("first dsl", func() {
			patches := NewPatches()
			defer patches.Reset()
			patchBuilder := NewPatchBuilder(patches)

			patchBuilder.
				Func(Belong).
				Stubs().
				With(Eq("zxl"), Any()).
				Will(Return(true)).
				Then(Repeat(Return(false), 2)).
				End()

			flag := Belong("zxl", []string{})
			So(flag, ShouldBeTrue)

			defer func() {
				if p := recover(); p != nil {
					str, ok := p.(string)
					So(ok, ShouldBeTrue)
					So(str, ShouldEqual, "input paras ddd is not matched")
				}
			}()
			Belong("ddd", []string{"abc"})
		})

	})
}
