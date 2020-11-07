# Daily Coding Problem: Problem #691 [Medium]

This problem was asked by Yahoo.

Write a function that returns the bitwise AND of all integers
between M and N, inclusive.

## Analysis

Wow, this is a goofy problem.
If any bit in any integer between M and N, inclusive,
is set to 0 (zero),
that bit's position remains zero throughout.
Almost all pairs (M, N) will cause this function to return
all-zeros.
