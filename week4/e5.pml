proctype p(){
	printf("soy el proceso (%d)", _pid)
	printf("Hola ......(%d)", _pid)
}

init {
	run p()
	run p()
}
