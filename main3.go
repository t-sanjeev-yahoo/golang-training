package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var i int
	var pi *int
	var pj = new(int)
	pi = &i
	i = 100
	i += 120
	pj = pi
	*pi++

	k := "ello"
	pk := &k

	l := []int{10, 20, 30, 50}
	m := [5]int{12, 12, 43}
	pl := &l
	fmt.Printf("l: %v, t: %T, s: %v\n", l, l, unsafe.Sizeof(l))

	k1 := l
	k1[2] = 1001
	(*pl)[1] = 110
	l = append(l, 255, 355)
	(*pl)[3] = 220
	(*pl) = append((*pl), 100)
	km := m
	km[1] = 200

	fmt.Printf("i: %v, t: %T, s: %v\n", i, i, unsafe.Sizeof(i))
	fmt.Printf("pi: %v, t: %T, s: %v\n", pi, pi, unsafe.Sizeof(pi))
	fmt.Printf("*pi: %v, t: %T, s: %v\n", *pi, *pi, unsafe.Sizeof(*pi))
	fmt.Printf("pj: %v, t: %T, s: %v\n", pj, pj, unsafe.Sizeof(pj))
	fmt.Printf("*pj: %v, t: %T, s: %v\n", *pj, *pj, unsafe.Sizeof(*pj))

	fmt.Printf("k: %v, t: %T, s: %v\n", k, k, unsafe.Sizeof(k))
	fmt.Printf("*pk: %v, t: %T, s: %v\n", *pk, *pk, unsafe.Sizeof(*pk))

	fmt.Printf("l: %v, t: %T, s: %v\n", l, l, unsafe.Sizeof(l))
	fmt.Printf("*pl: %v, t: %T, s: %v\n", *pl, *pl, unsafe.Sizeof(*pl))
	fmt.Printf("pl: %v, t: %T, s: %v\n", pl, pl, unsafe.Sizeof(pl))

	for i := range *pl {
		fmt.Println(i)
	}
	fmt.Printf("l: %v, %T\n", l, l)
	fmt.Printf("k1: %v, %T\n", k1, k1)
	fmt.Printf("m: %v, %T\n", m, m)
	fmt.Printf("km: %v, %T\n", km, km)
	fmt.Printf("km: %p, %p, %p\n", pl, &l, &k1)
	//fmt.Printf("km: %x, %x, %x\n", id(pl), id(l), id(k1))

	a := []int{1, 2, 3, 4, 5}
	b := a
	fmt.Println(a)
	fmt.Println(b)
	b[1] = 6
	fmt.Println(a)
	fmt.Println(b)

	fmt.Printf("a: %v, b:%v\n", a, b)
	fmt.Printf("a: %p, b:%p\n", a, b)
	a = append(a, 100)
	fmt.Printf("a: %v, b:%v\n", a, b)
	fmt.Printf("a: %p, b:%p\n", a, b)

}
