package main

import (
	"fmt"
	"time"
	tinytime "github.com/wagslane/go-tinytime"
)

func main(){
	tt :=tinytime.New(1585750374)
	tt = tt.Add(time.Hour * 24)
	fmt.Println(tt)
}