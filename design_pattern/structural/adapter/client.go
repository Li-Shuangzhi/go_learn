package main

func main() {
	cardImpl := SDCardImpl{info: "1)init"}
	cardImpl.writeSD("2)hello")
	cardImpl.readSD()

	//类适配器测试
	adapter := &TFSDAdapter{&TFCardImpl{info: "1)init"}}
	adapter.writeSD("2)golang nb")
	adapter.readSD()

	//对象适配器测试
	adapter1 := &TFSDAdapter1{t: TFCardImpl{info: "1)init"}}
	adapter1.writeSD("2)go")
	adapter1.readSD()

}
