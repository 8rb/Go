mtype = {red, yellow, green}
mytype seleccionado = green

active[5] proctype P(){
	byte c=0
	do
	:: c < 20 -> 
		c++
		if
		:: seleccionado == red -> seleccionado = green
		:: seleccionado == yellow -> seleccionado = red
		:: seleccionado == green -> seleccionado = yellow
		fi
		printf("El color seleccionado es %e %d \n", seleccionado, _pid)
	:: else -> break
	od
}
