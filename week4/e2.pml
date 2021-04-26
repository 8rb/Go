proctype hijo(){
	printf("Soy el proceso hijo (%d)\n", _pid)
}
active proctype padre(){
	do
	::(_nr_pr == 1) ->
		run hijo()
	od
}
