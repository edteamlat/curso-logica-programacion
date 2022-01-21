import random

def crearPassword(may,num,car,tam):    
	caracteres = list('abcdefghijklmnopqrstuvwxyz')
	if  may=='SI':
		caracteres.extend(list('ABCDEFGHIJKLMNOPQRSTUVWXYZ'))

	if num=='SI':
		caracteres.extend(list('0987654321'))

	if car=='SI':
		caracteres.extend(list('!"#$%&/()@'))

    password = ''

    for x in range(tam):
	    password += random.choice(caracteres)
        
    return password


print("Vamos a generar una contraseña para ti")
may = input("Escribe SI o NO quieres contemplar mayusculas")
num = input("Escribe SI o NO quieres contemplar números")
car = input("Escribe SI o NO quieres contemplar caracteres especiales")
tam = int(input("Escribe el tamaño que quieres tenga la contraseña"))

psw= crearPassword(may,num,car,tam)
print("la contraseña generada es = ", psw)