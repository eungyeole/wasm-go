package main

import (
	"io/ioutil"
	"net/http"
	"syscall/js"
)

func request(url string) []byte {
	println("1")
	resp, err := http.Get(url)
	println("2")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	println("3")
	data, err := ioutil.ReadAll(resp.Body)
	println("4")
	if err != nil {
		panic(err)
	}
	return data
}
func main() {
	//js.Global().Get("document").Call("getElementById", "test").Set("value", string(res))
	js.Global().Set("MyGoFunc", MyGoFunc())
	select {}
}

func MyGoFunc() js.Func {

	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var res = request("https://api.eungyeol.live/feed/list?page=0")
		return string(res)
	})
}
