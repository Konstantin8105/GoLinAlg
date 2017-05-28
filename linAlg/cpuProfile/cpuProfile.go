package cpuProfile

import (
	"fmt"
	"time"
)

// GetCPUProfile - show a cpu memory profile
func GetCPUProfile() {
	minFactor := 6
	maxFactor := 26
	sizeOfFloat64 := 8 // bait

	for i := minFactor; i <= maxFactor; i += 2 {
		amountIterations := int(1000000.0 / float64(i*i))
		// memory in bait, not a bit
		memory := 1
		for j := 0; j < i; j++ {
			memory = memory << 1
		}
		amountElements := memory / sizeOfFloat64
		memory = amountElements * sizeOfFloat64

		buffer := make([]node, amountElements)
		for j := 0; j < amountElements; j++ {
			buffer[j].value = float64(j)
		}

		for j := 0; j < amountElements/2; j++ {
			buffer[j].next = &buffer[amountElements-1-j]
			buffer[amountElements-1-j].next = &buffer[j+1]
		}
		buffer[amountElements-1-amountElements/2+1].next = nil

		/*
			for j := 0; j < amountElements-1-1; j++ {
				buffer[j].next = &buffer[j+1]
			}
			buffer[amountElements-1].next = nil
		*/

		startNode := buffer[0]
		//for j := 0; j < len(buffer); j++ {
		//	fmt.Printf("\n%v\t%#v,\t%v", j, buffer[j], (buffer[j].next))
		//}

		cpuTime := calculate(amountElements, amountIterations, startNode)
		fmt.Printf("%v.\t|%viter|\t%v bait\t%.4E ns\n", i, amountIterations, memory, cpuTime)
	}
}

type node struct {
	value float64
	next  *node
}

func calculate(amountElements, amountIterations int, start node) float64 {
	startTime := time.Now()
	for i := 0; i < amountIterations; i++ {
		s := &start
		for {
			if s.next == nil {
				break
			}
			s.value = s.next.value * 3.1415926 * s.value / 4.0
			s = s.next
		}
	}
	diff := time.Since(startTime).Nanoseconds()
	return float64(diff) / float64(amountIterations) / float64(amountElements)
}
