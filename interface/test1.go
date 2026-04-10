package main

type Fly interface {
	Fly()
}
type Plane struct{}

func (p Plane) Fly() {}

var f Fly = Plane{} //这边如果不做初始化的话就是nil了
