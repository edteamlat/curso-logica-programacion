disponible = {
    100: 3,
    50: 6,
    20: 10,
    10: 50,
    1: 50,
}

def total_disponible():
    total = 0
    for billete, cantidad in disponible.items():
   		total += billete * cantidad
    return total


def retirar(monto):
    if monto > total_disponible():
        print("Error, no hay suficiente efectivo")
    
    saldo = monto
    while saldo > 0:
        for denominacion, v in disponible.items():
            cantidad = get_denominacion(denominacion, saldo)
            print("Billete o moneda de: ", denominacion, cantidad)
            saldo -= cantidad * denominacion
        

def get_denominacion(denominacion, saldo):
    cantidad = saldo/denominacion
    if (cantidad > disponible[denominacion]):
        cantidad = disponible[denominacion]
    
    disponible[denominacion] -= cantidad
    return cantidad


retirar(1300)
