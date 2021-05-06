int turn=1
int cc=0

active proctype P()
{
    do
    ::
        (turn == 1) ->
            c++
            assert(cc < 2)
            c--
            turn = 2
    od
}

active proctype Q()
{
    do
    ::
        (turn == 2) ->
            c++
            assert(cc < 2)
            c--
            turn = 1
    od
}