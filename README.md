# Daily Coding Problem: Problem #691 [Medium]

This problem was asked by Yahoo.

Write a function that returns the bitwise AND of all integers
between M and N, inclusive.

## Analysis

[Brute force solution](puzzle1.go)

[Semi-magical algorithm](puzzle2.go)

Wow, this is a goofy problem.
If any bit in any integer between M and N, inclusive,
has a 0 (zero) value,
that bit's position remains zero throughout the bitwise-anding.
Almost all pairs (M, N) will cause this function to return
all-zeros.
If M has the value of zero, the bitwise-anded-together value
will stay zero throughout.

Even if you brute-force bitwise-and subsequent numbers inside
a for-loop, flow-of-control can exit the loop whenever the
anded-so-far value gets to zero.

There's a semi-magical solution floating around the net:

```go
func rangeBitwiseAnd(m, n uint) uint {
    var a uint = 0
    for m != n {
        m >>= 1
        n >>= 1
        a++
    }
    return m << a
}
```

In essense this finds the highest bit(s) of both numbers that
are set, and zeros the lower bit(s).

For example M = 202, N = 205:

```
  76543210
M 11001010
N 11001101
```

M and N have 011001___ in common.
The semi-magical algorithm will shift M and N left 3 bits,
where they are both valued 11001.
The loop exits.
The code shifts M left 3 bits to get 11001000.

There's absolutely no way anyone thinks of this algorithm during an interview.
Either the candidate has it memorized, or they do the brute force algorithm,
and they're lucky to realize the loop can exit when the anded-value gets to zero.

This seems like a good question for an interview at first glance,
but after a few minutes of thought, it turns sour.
I don't think an interviewer will get any idea of a candidate's coding prowess,
the "ah ha!" insight is so odd as to probably never arrive.
This will just puzzle and plague the candidate for no reason.
Don't use it.
