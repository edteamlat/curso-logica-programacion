const disponible = {
    100: 3,
    50: 6,
    20: 10,
    10: 50,
    1: 50,
};

const billetes = Object.keys(disponible);

const totalDisponible = () => {
    let total = 0;
    for (let i = 0; i < billetes.length; i++) {
        total += billetes[i] * disponible[billetes[i]]
    }

    return total;
}


const calcularBilletes = (saldoARetirar, denominacion) => {
    let contador = saldoARetirar / denominacion

    if (contador > disponible[denominacion]) {
        contador = disponible[denominacion]
    }
    disponible[denominacion] -= contador

    return contador
}


const procesar = valorAretirar => {
    if (valorAretirar > totalDisponible()) {
        console.log("Error, no hay suficiente saldo")

        return
    }

    let saldoARetirar = valorAretirar;
    while (saldoARetirar > 0) {

        let contador = 0;
        for (let i = 0; i < billetes.length; i++) {
            contador = calcularBilletes(saldoARetirar, billetes[i]);
            console.log(`${contador} billetes de denonimacion: ${billetes[i]}`);

            saldoARetirar -= contador * billetes[i]
        }
    }
}



procesar(135);
