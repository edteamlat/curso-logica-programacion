const encriptar = () => {
    let abc = "abcdefghijklmnñopqrstuvwxyz";
    let mensaje = prompt("Ingresa el mensaje que quieres encriptar");
    let desplazamiento = Number(prompt("Ingresa el desplazamiento requerido"));
    let resultado = "";

    for (let i = 0; i < mensaje.length; i++) {
        let posActual = abc.indexOf(mensaje[i]);
        let posDeseada = posActual + desplazamiento
        let posNueva = posDeseada < abc.length ? posDeseada : posDeseada % abc.length;
        let letra = abc[posNueva];
        
        resultado += letra;
    }

    return resultado;
}

console.log(encriptar());