# tictactoe-mcts

This is a private experiment for me familiarize myself with the [Monte Carlo Tree Search (MCTS)](https://en.wikipedia.org/wiki/Monte_Carlo_tree_search) algorithm. It is used to power the opponent in a game of tic-tac-toe.

## Command Line

These are the supported command line options:

```
$ ./tictactoe-mcts-linux-amd64 --help
Usage:
  tictactoe-mcts-linux-amd64 [OPTIONS]

Application Options:
  -s, --seed=            The RNG seed used by MCTS (default: 1690791430501)
  -t, --timeout=         The maximum amount of time the MCTS algorithm may take per action (default: 1s)
  -p, --parallelization= The amount of parallel goroutines to execute for the MCTS algorithm (default: 4)

Help Options:
  -h, --help             Show this help message
```

## Gameplay

In the example below the AI goes first (player "O"), while the human is player "X".

```
=======
|X|O| |
|O|O|X|
|X|O| |
=======
[2023-07-31 10:06:36] <O> $ 2x2 (explored 123934 nodes in 1002ms)
[2023-07-31 10:06:38] <X> $ 1x1
[2023-07-31 10:06:39] <O> $ 2x1 (explored 9633 nodes in 1001ms)
[2023-07-31 10:06:42] <X> $ 3x2
[2023-07-31 10:06:43] <O> $ 1x2 (explored 957 nodes in 1000ms)
[2023-07-31 10:06:48] <X> $ 1x3
[2023-07-31 10:06:49] <O> $ 3x2 (explored 44 nodes in 1002ms)
[2023-07-31 10:06:49]     $ the winner is player O!
```