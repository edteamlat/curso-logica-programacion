participantes = ["AAA", "BBB", "CCC", "DDD", "EEE", "FFF"]

def sorteo(participantes):
    if( len(participantes) %2 != 0 ):
        participantes.append("descansa")

    fixture=[]
    rivales = len(equipos)-1
    partidos = len(equipos) / 2

    for r in range(rivales):
        cantidad = rivales
        
		for p in range(partidos):
			fixture.extend([participantes[p]])
            fixture.extend([participantes[cantidad]]])
            cantidad -= 1

        participantes.insert(1, participantes[-1])	
        participantes.pop()
        fixture.extend(["##### jornada", r])

    return fixture

print(sorteo(participantes))