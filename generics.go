package main

import "golang.org/x/exp/constraints"

func SomaInteiros(m map[string]int64) int64 {
	var soma int64

	for _, v := range m {
		soma += v
	}

	return soma
}
func SomaFloat(m map[string]float64) float64 {
	var soma float64

	for _, v := range m {
		soma += v
	}

	return soma
}

type Number interface{
	~int64 | ~float64 // ~ funciona comm um operador de aproximação
}

type MyNumber int64

func SomaGenerica[T Number](m map[string]T) T {
	var soma T

	for _, v := range m {
		soma += v
	}

	return soma
}

func Equals[T comparable](number1 T, number2 T) T {
	if number1 == number2 {
		return number1
	}
	return number2
}

func Max[T constraints.Ordered](number1 T, number2 T) T{
	if number1 >  number2 {
		return number1
	}

	return number2
}

type stringfy interface{
	String() string
}

func (m MyNumber) String() string {
	return string(m)
}

func Concat[T stringfy](vals []T) string {
	result := ""

	for _, val := range vals {
		result += val.String()
	}

	return result
}

func main(){
	println("Comum")
	println(SomaInteiros(map[string]int64{"a": 1, "b": 2, "c": 3}))
	println(SomaFloat(map[string]float64{"a": 1.5, "b": 2.5, "c": 3.5}))
	
	println("Generics")
	println(SomaGenerica(map[string]int64{"a": 1, "b": 2, "c": 3}))
	println(SomaGenerica(map[string]float64{"a": 1.5, "b": 2.5, "c": 3.5}))
	
	println("Generic Types")
	var a,b,c MyNumber
	a = 1
	b = 2
	c = 3

	println(SomaGenerica(map[string]MyNumber{"a": a, "b": b, "c": c}))

	println(Equals(1,1))
	println(Equals(1,2))

	println(Max(13, 5))
	println(Max(10, 24))

	println(Concat([]MyNumber{1,2,3,4,5}))
}