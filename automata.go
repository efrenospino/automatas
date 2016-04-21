package automata

import (
	"errors"
	"strings"
)

const (
	a = "a"
	b = "b"
	c = "c"
)

func validaAFD(cadena string) (bool, error) {
	cadenaSeparada := strings.Split(cadena, "")
	if len(cadenaSeparada) > 1 {
		if cadenaSeparada[0] == a {
			return validaKleene(cadenaSeparada[1:])
		} else if cadenaSeparada[0] == b {
			if cadenaSeparada[1] == a {
				return validaKleene(cadenaSeparada[2:])
			}
			err := errors.New("Si la cadena empieza en 'b', debe seguir 'a'")
			return false, err
		} else if cadenaSeparada[0] == c {
			if cadenaSeparada[1] == b && cadenaSeparada[2] == a {
				return validaKleene(cadenaSeparada[3:])
			} else if cadenaSeparada[1] == a {
				return validaKleene(cadenaSeparada[2:])
			}
			err := errors.New("Si la cadena empieza en 'c', debe seguir 'ba' o 'a'")
			return false, err
		}
	}
	if cadenaSeparada[0] != a {
		err := errors.New("Si existe un solo caracter debe ser 'a'")
		return false, err
	}
	return true, nil
}

func validaKleene(cadena []string) (bool, error) {
	for i := 0; i < len(cadena)-1; i++ {
		if cadena[i] == a {
		} else if cadena[i] == b && (cadena[i+1] == a || cadena[i+1] == b) {
		} else if cadena[i] == c && (cadena[i+1] == a || cadena[i+1] == c) {
		} else {
			err := errors.New("Las cadenas 'bc' o 'cb' no son validas despues de la primera 'a'")
			return false, err
		}
	}
	if cadena[len(cadena)-1] != a {
		err := errors.New("La cadena debe terminar en 'a'")
		return false, err
	}
	return true, nil
}
