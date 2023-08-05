package main

import (
	"bufio"
	"fmt"
)

type car struct{}

func (c car) drive() {
	fmt.Println("car drive")
}
func (c car) stop() {
	fmt.Println("stop")
}

type trafficColor interface {
	color() string
	check(car car)
}

func nextColor(t trafficColor) trafficColor {
	if t.color() == "red" {
		return newRed("green")
	} else if t.color() == "green" {
		return newRed("yellow")
	} else if t.color() == "yellow" {
		return newRed("red")
	}
	return nil
}

type red struct {
	c string
}

func (r red) color() string { return r.c }
func (r red) check(car car) {
	if r.color() == "red" {
		car.stop()
	} else if r.color() == "yellow" {
		car.stop()
	} else if r.color() == "green" {
		car.drive()
		bufio.Reader{}
	}
}
func newRed(color string) red {
	return red{c: color}
}

func main() {
	r := red{c: "green"}
	r2 := nextColor(r)
	r3 := nextColor(r2)
	r4 := nextColor(r3)
	r5 := nextColor(r4)
	r6 := nextColor(r5)
	r7 := nextColor(r6)

	fmt.Println(nextColor(r))
	fmt.Println(nextColor(r2))
	fmt.Println(nextColor(r3))
	fmt.Println(nextColor(r4))
	fmt.Println(nextColor(r5))
	fmt.Println(nextColor(r6))
	fmt.Println(nextColor(r7))
}
