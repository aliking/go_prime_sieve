# go_prime_sieve
An implementation of  a [prime sieve](https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes) using go to experiment with concurrency. The executable tool takes a number and outputs the counts of primes less than that number (with limited tweaking it can also output the primes, obviously).

Sequential implementations of this algorithm usually first generate a list of all numbers less than the maximum. Starting with 2, eliminate all multiples of 2 - they are not primes. Then take the next unmarked number (3) and mark multiples of that as not prime. And so on, until all that is left is primes.

This implementation attempts to optimize by creating a goroutine for each number in the list so that each can mark it's multiples without waiting.

This likely runs faster, but can lead to redundancy by checking numbers that, using a sequential method, would never get checked since they would already be known as composite numbers.

In an attempt to work around that, the repo contains two executables: sieve and sieve_fast. sieve_fast allows any number checking goroutine to notice that its initial number is not a prime and stop processing immediately.
### Benchmark on MacbookPro 2.3 GHz Intel Core i7

***
```
time ./sieve 10000000
664579 primes under 10000000

real	0m49.321s
user	2m11.151s
sys	0m39.499s
```
******
```
time ./sieve_fast 10000000
664579 primes under 10000000

real	0m15.279s
user	0m31.331s
sys	0m6.202s
```
***


### Comparison to standard algorithm
***
... version with concurrency hacked out:
```
664579 primes under 10000000

real	0m17.497s
user	0m17.843s
sys	0m0.757s
```
Turns out the redundant number checking is a much bigger deal than I expected, especially on a system where this looks like the parallelism is not high. 



bonus: benchmark of older version of sieve and sieve_fast. This version dispatches goroutines for all of the numbers less than max. Newer versions only test numbers between 2 and the square root of max:
***
```
time ./sieve 10000000
664579 primes under 10000000

real	2m48.997s
user	6m4.152s
sys	3m34.534s
```
******
```
time ./sieve_fast 10000000
664579 primes under 10000000

real	0m28.817s
user	1m16.492s
sys	0m18.017s
```
***
