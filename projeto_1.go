package main

import "fmt"

// declaração da variável CONST do ponto de ebulição da água em Kelvin
const ebulicaoK = 373.15

func main() {
	tempKK := 32.0
	tempK := ebulicaoK
	tempC := tempK - tempKK // Convertendo K para C

	fmt.Printf("A temperatura de ebulição da água em K° = %v , a temperatura de ebulição em C° a partir de %v K  é = %v graus celsius ", tempK, tempKK, tempC)
}
