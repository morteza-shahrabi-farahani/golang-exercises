package main

import (
	"time"
)

func Solution(d time.Duration, message string, ch ...chan string) (numberOfAccesses int) {
	//channels := make([]chan string, len(ch))
	//time.Sleep(3*time.Second)
	counter := 0
	start := time.Now()

	//_ := time.After(d * time.Second)
	
	//fmt.Println(time.Since(start))
	for time.Duration(time.Now().Sub(start).Seconds()) < d {

	}
	//time.Sleep(d * time.Second)
	for i := 0; i < len(ch); i++ {
		
		if ch[i] == nil || cap(ch[i]) == 0 {
			//fmt.Printf("channel %v is nil or cap = 0\n", i)
			continue;
		} else {
		// 	if len(ch[i]) == 0 {
		// 		counter += 1
		// 		ch[i] <- message
		// 		fmt.Printf("channel %v is empty\n", i)
		// 		//close(ch[i])
		// 	} else {
		// 		value, ok2 := <- ch[i]
		// 		if ok2 {
		// 			counter += 1
		// 			fmt.Printf("(%v): %v\n", ok2, value)
		// 			fmt.Printf("channel %v is open\n", i)
		// 			ch[i] <- value
		// 			ch[i] <- message
		// 			//close(ch[i])
		// 		} else {
		// 			fmt.Printf("channel %v is closed\n", i)
		// 		}
		// 	}

			select {
			case value, ok := <-ch[i]:
				if !ok {
					//fmt.Println("we can't read from channel now %d", i)
				}
				if ok {
					//fmt.Println("we can read from channel %d", i)
					//fmt.Printf("value is %v\n", value)
					//fmt.Printf("message is %v\n", message)
					//fmt.Printf("cap is %v\n", cap(ch[i]))
					if cap(ch[i]) != 1 {
						ch[i] <- value
					}
					//fmt.Printf("hello word\n")
					ch[i] <- message
					//counter++
				}
			//case <-time.After(d * time.Second):

			default:
				
				counter += 1
				//fmt.Println("default state for = %d", i)
				ch[i] <- message
				//fmt.Println("now we are at %d", i)
				//close(ch[i])
			}
		}

	
	}

	return counter

}
