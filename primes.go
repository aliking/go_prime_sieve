package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"sync"
)

func findMultiples(factor int, max int, c chan int, wg *sync.WaitGroup, stoppable bool, primes map[int]bool) {
	var product int
	for mult := 2; ; mult++ {
		if stoppable && !primes[factor] {
			//fmt.Println(factor, " processor cut short after", mult-1, "cycles")
			break
		}
		product = factor * mult
		if product > max {
			//fmt.Println(factor, " processor completed cycle")
			break
		}
		c <- product
	}
	wg.Done()
}

func markPrimes(primes map[int]bool, c chan int) {
	for notPrime := range c {
		primes[notPrime] = false
	}
}

func listPrimes(primes map[int]bool) []int {
	var ret []int
	for elem, val := range primes {
		if val {
			ret = append(ret, elem)
		}
	}
	sort.Ints(ret)
	return ret
}

func main() {
	var num int
	var err error

	s := os.Args[1]
	if num, err = strconv.Atoi(s); err != nil {
		panic(err)
	}
	primes := make(map[int]bool)

	for i := 2; i <= num; i++ {
		primes[i] = true
	}

	c := make(chan int)

	var wg sync.WaitGroup

	go markPrimes(primes, c)

	routineMax := int(math.Sqrt(float64(num)))

	for i := 2; i <= routineMax; i++ {
		wg.Add(1)
		go findMultiples(i, num, c, &wg, true, primes)
	}
	wg.Wait()
	close(c)

	final := listPrimes(primes)
	//fmt.Println(final)
	fmt.Println(len(final), "primes under", num)

}
