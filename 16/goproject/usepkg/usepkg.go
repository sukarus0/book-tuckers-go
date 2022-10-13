package main

import (
	"fmt" // (1) 기본 제공 패키지
	"goproject/usepkg/custompkg" // (2) 현재 모듈 패키지

	"github.com/guptarohit/asciigraph"  // (3) 외부 저장소 패키지
	"github.com/tuckersGo/musthaveGo/ch16/expkg" // (3) 외부 저장소 패키지
)

func main() {
	custompkg.PrintCustom()
	expkg.PrintSample()

	data := []float64{3,4,5,6,9,7,5,8,5,10,2,7,2,5,6}
	graph := asciigraph.Plot(data)
	fmt.Println(graph)
}
