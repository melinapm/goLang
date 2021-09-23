# goLang
GRUPO: Melina Perez Mounet, Grela Juan
Nuestra solución a este ejercicio fue la siguiente:
    La funcion principal GenerateStruct que recibe como parámetro un string y crea una estructura del tipo Result con todos sus atributos vacíos o 0.
    Al string recibido como parámetro se lo "parsea" utilizando un array de tipo rune (de esta forma se puede utilizar el string como un array de chars). 
	Se le asigna a la estructura Result el carácter o los caracteres correspondientes según la posición del array y se valida que cumplan  su respectiva condición (que sea de tipo texto, numero, longitud, etc). En caso que no se cumpla alguna de estas condiciones se torna un error, caso contrario se retorna la estructura result y un nil.