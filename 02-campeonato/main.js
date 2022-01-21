const sorteo = (equipos) => {

    if (equipos.length % 2 > 0) {
        equipos.push('Descansa');
    }

    let fixture = []
    let rivales = equipos.length - 1;
    let partidos = equipos.length / 2;


    for (let i = 0; i < rivales; i++) {
        fixture[i] = []
        for (let j = 0; j < partidos; j++) {
            fixture[i].push([equipos[j], equipos[rivales - j]]);
        }
        equipos.splice(1, 0, equipos.pop())
    }

    return fixture;
}

let participantes = [
    "AAA",
    "BBB",
    "CCC",
    "DDD",
    "EEE",
    "FFF"
]

console.log(sorteo(participantes))