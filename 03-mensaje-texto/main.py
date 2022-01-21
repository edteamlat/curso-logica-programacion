numPad = {
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
}


def find_number_from_char(v, t):
    count = 0 
    
    for elemento in v:
        count = count + 1
        
        if (elemento == 1):
            break
        
    return count


def print_keys(text):
    anterior = ""
    cadena = ""
    
    for t in text:
        if (t == anterior):
            cadena += ""
        else
            anterior = t
        
        for k, v in numPad.items():
            if (t in v):
                repeat = find_number_from_char(v , t)

                for r in range(repeat)
                    cadena += str(k)
    return cadena                

print_keys("no")