package main

import (
	"syscall/js"
	"web/utils"
)

func main() {
	//alert := js.Global().Get("alert")
	//alert.Invoke("Hello World!")

	js.Global().Set("hello", js.FuncOf(hello))
	js.Global().Set("encryptGo", js.FuncOf(encrypt))
	js.Global().Set("decryptGo", js.FuncOf(decrypt))

	// 阻塞主线程，等待 js 调用
	select {}
}

// 定义一个函数，接受一个 js.Value 类型的参数
func hello(this js.Value, args []js.Value) interface{} {
	// 从参数中获取第一个值，并转换为字符串类型
	name := args[0].String()
	// 在控制台打印 "Hello, name!"
	println("Hello,", name+"!")
	// 返回 nil，表示没有返回值
	return nil
}

// 定义一个函数，接受一个 js.Value 类型的参数
func encrypt(this js.Value, args []js.Value) interface{} {
	data := args[0].String()

	// 密钥，长度必须为16、24或32字节
	key := []byte("0123456789abcdef")

	ciphertext, err := utils.EncryptAES(key, data)
	if err != nil {
		println("加密失败:", err)
		return nil
	}
	return ciphertext
}

// 定义一个函数，接受一个 js.Value 类型的参数
func decrypt(this js.Value, args []js.Value) interface{} {
	data := args[0].String()

	// 密钥，长度必须为16、24或32字节
	key := []byte("0123456789abcdef")

	ciphertext, err := utils.DecryptAES(key, data)
	if err != nil {
		println("加密失败:", err)
		return nil
	}
	return ciphertext
}
