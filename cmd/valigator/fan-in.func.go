package main

func fanIn(cSet []<-chan string) <-chan string {
	c := make(chan string)
	for i := range cSet {
		go func(in <-chan string) {
			for {
				x := <-in
				c <- x
			}
		}(cSet[i])
	}
	return c
}
