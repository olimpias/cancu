package main

import (
	"fmt"
	"cancu"
	"time"
	"log"
)

type CalculationManager struct {

}

func (cm * CalculationManager) Result(event CalculationCompletedEvent)  {
	fmt.Printf("%d %s %d = %d", event.value1,event.operation,event.value2, event.result)
}

type CalculationCompletedEvent struct {
	result int
	value1 int
	value2 int
	operation string
}


func main()  {
	cm := CalculationManager{}
	err := cancu.Cancu().AddListener(cm.Result)
	if err != nil {
		log.Fatal(err)
	}
	err = cancu.Cancu().Notify(CalculationCompletedEvent{10,2,8,"+"})
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(10)
}
