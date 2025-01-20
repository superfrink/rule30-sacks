# Overview

Drawing a Sacks Spiral of the center-column bits of Rule 30.  A dot is drawn when the bit is a one.

# Building and Running

```
go build && ./rule30-sacks -file=A-Million-Bits-of-the-Center-Column-of-the-Rule-30-Cellular-Automaton.csv | head -100000 > coordinates.csv

gnuplot < chart-1.gnuplot ; mv output.png sacks-100000.png
```

# Output

The first 100,000 bits:
![Sacks spiral for 100,000 bits](sacks-100000.png)

# References

- Coordinate formulas: https://numberspiral.com/p/p006.html
- Rule 30 center column bits: https://writings.stephenwolfram.com/2019/10/announcing-the-rule-30-prizes/
- Sacks Spirals for prime numbers: https://thatsmaths.com/2019/09/12/spiraling-primes/
