package main

import (
	"time"
)

func Solution(d time.Duration, message string, ch ...chan string) (numberOfAccesses int) {
	//channels := make([]chan string, len(ch))
	counter := 0
	start := time.Now()
	for time.Duration(time.Since(start).Seconds()) < d {
		//fmt.Println(time.Since(start))
		for i := 0; i < len(ch); i++ {

			if ch[i] == nil {
				continue;
			} else {
				select {
				case _, ok := <- ch[i]:
					if ok {
						counter += 1
						ch[i] <- message
						//fmt.Println("now we are at %d", i)
					} else {
						//fmt.Println(i)
					}
				default:
					counter += 1
					ch[i] <- message
					//fmt.Println("now we are at %d", i)
				}
			}
	
		}

	}

	return counter
}

