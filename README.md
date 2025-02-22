# Overview

Drawing a Sacks Spiral of the center-column bits of Rule 30.  A dot is drawn when the bit is a one.

# Building and Running

```
go build && ./rule30-sacks -file=A-Million-Bits-of-the-Center-Column-of-the-Rule-30-Cellular-Automaton.csv | head -100000 > coordinates.csv

gnuplot < chart-1.gnuplot ; mv output.png sacks-100000.png
```

# Output

There appear to be some horizontal curves near the center:
![Horizontal lines near the center](horizontal-lines.png)

But those also occur for all numbers:

> "The reason that the north and south progressions grow at twice the rate of the east and west progressions is that each north/south number is separated by two turns of the spiral." - https://www.naturalnumbers.org/sparticle.html

![All numbers](all-numbers.png)

The first 100,000 bits that are '1':
![Sacks spiral for 100,000 '1' bits](sacks-100000.png)

The first 100,000 bits that are '1' with a 2:1 aspect ratio:
![Sacks spiral for 100,000 '1' bits with 2:1 aspect ratio](sacks-100000-aspect-2-1.png)

The first 500,768 bits that are '1', out of the first million bits:
![Sacks spiral for 500,768 '1' bits](sacks-500768.png)

# References

- Coordinate formulas: https://numberspiral.com/p/p006.html
- Rule 30 center column bits: https://writings.stephenwolfram.com/2019/10/announcing-the-rule-30-prizes/
- Sacks Spirals for prime numbers: https://thatsmaths.com/2019/09/12/spiraling-primes/
