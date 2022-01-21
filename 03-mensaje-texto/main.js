let numPad = {
    0: " ",
    1: "",
    2: "abc",
    3: "def",
    4: "ghi",
    5: "jkl",
    6: "mno",
    7: "pqrs",
    8: "tuv",
    9: "wxyz",
};

const buscarNumero = letra => {
    let numeros = Object.keys(numPad);    
    let resultado = "";

    numeros.forEach(num => {        
        for (let i = 0; i < numPad[num].length; i++) {
            if (letra === numPad[num][i]) {
                for (let j = 0; j <= i; j++) {
                    resultado = `${resultado}${num}`  
                }                        
            }
        }
    })

    return resultado
}

const convertirTexto = texto => {
    let respuesta = "";

    for (let i = 0; i < texto.length; i++) {
        respuesta += buscarNumero(texto[i])        
    }

    return respuesta
}

console.log(convertirTexto("hola mundo"))
