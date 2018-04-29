package cancu

import (
	"sync"
	"reflect"
	"time"
)

const (
	//TODO Expose to Observer interface should be changeable.
	//TODO Move this to Observer.go
	DEFAULT_MAX_METHOD_SPEND_TIME = 1 //s
)
// Listener registry assign to one specific struct type.
// In other word, every struct has its own Listener Registry
type listenerRegistry struct {
	mu sync.Mutex
	funcValues []reflect.Value
}

// Create a new Listener Registry
func newListenerRegistry() eventInterface {
	return &listenerRegistry{funcValues:make([]reflect.Value,0)}
}

// Attach new function for event
func (lr *listenerRegistry) attachListener(funcValue reflect.Value) {
	lr.mu.Lock()
	defer lr.mu.Unlock()
	lr.funcValues = append(lr.funcValues, funcValue)
}

// deattach function from events
func (lr *listenerRegistry) detachListener(funcValue reflect.Value)  {
	lr.mu.Lock()
	defer lr.mu.Unlock()
	for i,v := range lr.funcValues{
		if v.Pointer() == funcValue.Pointer() {
			lr.funcValues = append(lr.funcValues[:i], lr.funcValues[i+1:]...)
			break
		}
	}
}

func (lr * listenerRegistry) penaltyForFunctions(pFuncValues [] reflect.Value){
	if len(pFuncValues) == 0 {
		return
	}
	for _,v := range pFuncValues{
		lr.detachListener(v)
	}
}
// Sends input value to registered functions.
func (lr * listenerRegistry) notifyEvent(value reflect.Value)  {
	methodIn := []reflect.Value{value}
	toBeRemove := make([]reflect.Value,0)
	lr.mu.Lock()
	var currentTime time.Time
	var stopTime time.Time
	for _,v := range lr.funcValues{
		currentTime = time.Now()
		v.Call(methodIn)
		stopTime = time.Now()
		if stopTime.Second() - currentTime.Second() > DEFAULT_MAX_METHOD_SPEND_TIME {
			toBeRemove = append(toBeRemove,v)
		}
	}
	lr.mu.Unlock()
	lr.penaltyForFunctions(toBeRemove)
}

type eventInterface interface {
	attachListener(funcValue reflect.Value)
	detachListener(funcValue reflect.Value)
	notifyEvent(value reflect.Value)
}


