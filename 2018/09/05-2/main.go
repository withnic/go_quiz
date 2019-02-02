package main

// https://twitter.com/empijei/status/1037362222245269504
//#golang pop quiz:
//does this func properly close t?
//Yes
//No (why?)
//Depends (on what?)
func Close(t chan int) {
	select {
	case _, ok := <-t:
		if ok {
			close(t)
		}
	default:
			close(t)
			}
}

func main(){
	t :=make(chan int, 1)

	Close(t)
}