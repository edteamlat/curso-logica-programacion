package main

import (
	"fmt"
)

const (
	Espacio = rune(' ')
)

var abecedario = []rune("abcdefghijklmnñopqrstuvwxyz")

func main() {

	mensaje := "este es un mensaje muy secreto y mañana lo veremos"
	
	desplazamiento := 11
	fmt.Println("mensaje a cifrar", mensaje)

	mensajeCifrado := cifrar(mensaje, desplazamiento)
	fmt.Println("Mensaje cifrado", mensajeCifrado)
	
	mensajeOriginal := descifrar(mensajeCifrado, desplazamiento)
	fmt.Println("Mensaje descifrado", mensajeOriginal)
}

func cifrar(mensaje string, desplazamiento int) string {
	resp := ""
	
	for _, letra := range mensaje {
		if letra == Espacio {
			resp += string(Espacio)
			continue
		}
	
		index := indexFromRunes(abecedario, letra)
		final := index + desplazamiento
		if final > len(abecedario) {
			final = final - len(abecedario)
		}
	
		resp += string(abecedario[final])
	}

	return resp
}

func descifrar(mensaje string, desplazamiento int) string {
	resp := ""
	
	for _, letra := range mensaje {
	
		if letra == Espacio {
			resp += string(Espacio)
			continue
		}
	
		index := indexFromRunes(abecedario, letra)
		final := index - desplazamiento
	
		if final < 0 {
			final = final + len(abecedario)
		}
	
		resp += string(abecedario[final])
	}

	return resp
}

func indexFromRunes(text []rune, letter rune) int {
	for i, v := range text {
		if v == letter {
			return i
		}
	}

	return -1
}