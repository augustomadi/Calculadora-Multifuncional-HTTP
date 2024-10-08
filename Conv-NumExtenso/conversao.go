package main 


import "strings"


func StringToInt(numeroPorExtenso string) int {
    palavras := strings.Fields(numeroPorExtenso)
    mapeamento := map[string]int{
        "zero": 0, "um": 1, "dois": 2, "três": 3, "quatro": 4,
        "cinco": 5, "seis": 6, "sete": 7, "oito": 8, "nove": 9,
        "dez": 10, "onze": 11, "doze": 12, "treze": 13, "catorze": 14,
        "quinze": 15, "dezesseis": 16, "dezessete": 17, "dezoito": 18, "dezenove": 19,
        "vinte": 20, "trinta": 30, "quarenta": 40, "cinquenta": 50,
        "sessenta": 60, "setenta": 70, "oitenta": 80, "noventa": 90,
        "cem": 100, "cento": 100, "duzentos": 200, "trezentos": 300,
        "quatrocentos": 400, "quinhentos": 500, "seiscentos": 600,
        "setecentos": 700, "oitocentos": 800, "novecentos": 900,
        "mil": 1000,
    }
    
    
    var total int
    var valorAtual int
    
    for _, palavra := range palavras {
        if valor, encontrado := mapeamento[palavra]; encontrado {
            valorAtual += valor
        } else if palavra == "e" {
            continue
        } else if valor, encontrado := mapeamento[strings.TrimSuffix(palavra, "e")]; encontrado {
            valorAtual += valor
        } else {
            total += valorAtual
            valorAtual = 0
        }
    }
    
    total += valorAtual
    return total
}




