package main

import (
	"fmt"
)

var participantes = []string{"AAA", "BBB", "CCC", "DDD", "EEE"}

func main() {
	for _, v := range sorteo(participantes){
		fmt.Println(v)
	}
}

func sorteo(equipos []string) [][]string {
	if len(equipos)%2 != 0 {
		equipos = append(equipos, "DESCANSA")
	}

	var fixture [][]string
	partidos := len(equipos) / 2
	rivales := len(equipos) - 1

	for i := 0; i < rivales; i++ {
		fixture = append(fixture, []string{})

		for j := 0; j < partidos; j++ {
			fixture[i] = append(fixture[i], []string{equipos[j], equipos[rivales-j]}...)
		}

		lastIndex := len(equipos) - 1		
		var temp []string
		temp = append(temp, equipos[0], equipos[lastIndex])
		temp = append(temp, equipos[1:lastIndex]...)	

		equipos = temp
	}

	return fixture
}