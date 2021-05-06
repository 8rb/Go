int turn=1

active proctype P()
{
    do
    ::
        printf("SNC line 1 del proceso (%d)\n", _pid)
        printf("SNC line 2 del proceso (%d)\n", _pid)

        (turn == 1) ->
            printf("SC line 1 del proceso (%d)\n", _pid)
            printf("SC line 2 del proceso (%d)\n", _pid)
            turn = 2
    od
}

active proctype Q()
{
    do
    ::
        printf("SNC line 1 del proceso (%d)\n", _pid)
        printf("SNC line 2 del proceso (%d)\n", _pid)

        (turn == 2) ->
            printf("SC line 1 del proceso (%d)\n", _pid)
            printf("SC line 2 del proceso (%d)\n", _pid)
            turn = 1
    od
}