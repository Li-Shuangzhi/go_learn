package main

type TFSDAdapter struct {
	*TFCardImpl
}

func (a *TFSDAdapter) readSD() {
	a.readTF()
}

func (a *TFSDAdapter) writeSD(str string) {
	a.writeTF(str)
}

type TFSDAdapter1 struct {
	t TFCardImpl
}

func (a *TFSDAdapter1) readSD() {
	a.t.readTF()
}

func (a *TFSDAdapter1) writeSD(str string) {
	a.t.writeTF(str)
}
