# goLang
Nuestra solucion a este ejercicio fue la siguiente:
    Una funcion recibe como parametro un string y inicialmente crea una estructura del tipo Result con todos sus atributos vacios o 0,
    luego al string que fue recibido como parametro lo "parseamos" utilizando la funcion rune (la cual fue la unica que encontramos para parsear una cadena de caracteres) y luego, con esos caracteres ya parseados, rellenamos la estrutuca Result con cada caracter o caracteres en su debido campo, solamente si estos cumplen su respectiva condiciones para ser una cadena valida , y por ultimo retornamos la estructura y el error
