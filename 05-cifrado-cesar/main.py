def cifrado(mensaje, desplazamiento, abecedario):
	mensaje_cifrado = ""
	
    for letra in mensaje:
	    pos = abc.find(letra)
	    final = pos + desplazamiento
	    conversion = abc[final]
	    mensaje_cifrado += conversion
    
    return mensaje_cifrado


def descifrar(mensaje, desplazamiento, abecedario):
	mensaje_descifrado = ""
	
    for letra in mensaje:
		pos = abc.find(letra)
		final = pos - desplazamiento
		conversion = abc[final]
		mensaje_descifrado += conversion

    return mensaje_descifrado


abc= "abcdefghijklmnñopqrstuvwxyzabcdefghijklmnñopqrstuvwxyz"
print("Bienvenido al software para encriptar mensajes")

mensaje = input("escribe el mensaje que quieres encriptar")
desplazamiento =int( input("Escribe la cantidad de desplazamiento 1-27"))

c = cifrado(mensaje,desplazamiento,abc)
print("tu mensaje cifrado es: ",c)

d = descifrar(c,desplazamiento,abc)
print("tu mensaje descifrado es: ",d)