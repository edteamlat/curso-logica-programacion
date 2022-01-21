const generar = () => {
    let tam = prompt("¿Cuántos caracteres quieres que tenga tu contraseña?");
    let may = prompt("Escribe SI o NO, si quieres contemplar mayusculas");
    let num = prompt("Escribe SI o NO, si quieres contemplar números");
    let car = prompt("Escribe SI o NO, si quieres contemplar caracteres especiales");
    let password = "";

    let caracteres = 'abcdefghijklmnopqrstuvwxyz';
    if (may === "SI") {
        caracteres = caracteres + "ABCDEFGHIJKLMNOPQRSTUVWXYZ";
    }
    if (num === "SI") {
        caracteres = caracteres + "0987654321";
    }
    if (car === "SI") {
        caracteres = caracteres + '!"#$%&/()@';
    }

    for (let i = 0; i < tam; i++) {
        password = password + caracteres[Math.trunc(Math.random() * caracteres.length)]
    }

    return password;
}

generar();