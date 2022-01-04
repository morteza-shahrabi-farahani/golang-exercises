package main

import (
	"time"
)

func Solution(d time.Duration, message string, ch ...chan string) (numberOfAccesses int) {
	counter := 0
	start := time.Now()

	//wait for d seconds
	for time.Duration(time.Now().Sub(start).Seconds()) < d {

	}
	

	for i := 0; i < len(ch); i++ {
		
		if ch[i] == nil || cap(ch[i]) == 0 {
			continue;
		} else {
			select {
			case value, ok := <-ch[i]:
				if !ok {
					//we can't read from this channel
				}
				if ok {
					//we can read from this channel
					//if capacity is bigger than one we should write twice. First one is the last value and second is new value
					if cap(ch[i]) != 1 {
						ch[i] <- value
					}
					ch[i] <- message
				}

			default:
				//default: for example channel doesn't have any member
				counter += 1
				ch[i] <- message
			}
		}

	
	}

	return counter

}
