package gomonkey

import "reflect"

gomonkey 是一个使单元测试中的猴子修补变得容易的库，猴子修补的核心思想来自 Bouke，您可以阅读这篇博文以了解其工作原理。

# Features
* 支持功能补丁
* 支持公共成员方法的补丁
* 支持私有成员方法的补丁
* 支持接口补丁
* 支持函数变量的补丁
* 支持全局变量的补丁
* 支持功能指定序列的补丁
* 支持成员方法指定序列的补丁
* 支持接口指定序列的补丁
* 支持函数变量指定序列的补丁

# Notes
* 如果启用内联，gomonkey 无法修补函数或成员方法，请通过添加 -gcflags=-l（低于 go1.10）或 -gcflags=all=-l（ go1.10及以上）。
* 当一个 goroutine 正在 patch 一个函数或者一个成员方法同时被另一个 goroutine 访问时，可能会发生恐慌。也就是说gomonkey不是线程安全的。

# Supported Platform:
* ARCH:
	amd64
	arm64
	386
* OS:
	Linux
	MAC OS X
	Windows

# Installation
* v2.1.0 以下，例如 v2.0.2
$ go get github.com/agiledragon/gomonkey@v2.0.2
* v2.1.0及以上，例如v2.2.0
$ go get github.com/agiledragon/gomonkey/v2@v2.2.0

# Test Method
$ cd test
$ go test -gcflags=all=-l

# Using gomonkey
请参考测试用例作为成语，非常完整和详细。

# Index
func GetResultValues(funcType reflect.Type, results ...interface{}) []reflect.Value
type OutputCell struct {
	Values Params
	Times  int
}
type Params []interface{}
type Patches struct {
	// contains filtered or unexported fields
}
func ApplyFunc(target, double interface{}) *Patches
func ApplyFuncSeq(target interface{}, outputs []OutputCell) *Patches
func ApplyFuncVar(target, double interface{}) *Patches
func ApplyFuncVarSeq(target interface{}, outputs []OutputCell) *Patches
func ApplyGlobalVar(target, double interface{}) *Patches
func ApplyMethod(target reflect.Type, methodName string, double interface{}) *Patches
func ApplyMethodSeq(target reflect.Type, methodName string, outputs []OutputCell) *Patches
func NewPatches() *Patches
func (this *Patches) ApplyCore(target, double reflect.Value) *Patches
func (this *Patches) ApplyFunc(target, double interface{}) *Patches
func (this *Patches) ApplyFuncSeq(target interface{}, outputs []OutputCell) *Patches
func (this *Patches) ApplyFuncVar(target, double interface{}) *Patches
func (this *Patches) ApplyFuncVarSeq(target interface{}, outputs []OutputCell) *Patches
func (this *Patches) ApplyGlobalVar(target, double interface{}) *Patches
func (this *Patches) ApplyMethod(target reflect.Type, methodName string, double interface{}) *Patches
func (this *Patches) ApplyMethodSeq(target reflect.Type, methodName string, outputs []OutputCell) *Patches
func (this *Patches) Reset()
