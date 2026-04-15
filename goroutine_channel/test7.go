package main

var Done = make(chan struct{})

func cancelled() bool {
	select{
	case <-Done:
		return true
	default:
		return false
	}
}
