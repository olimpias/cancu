package cancu

import (
	"reflect"
	"errors"
	"sync"
)

// TODO write documentation...
type Observer interface {
	Notify(value interface{}) error
	AddListener(method interface{}) error
	DelListener(method interface{}) error
}

var funcMustTakeOneParameterError = errors.New("function must take only one parameter as input")
var inputValueMustbeFuncError = errors.New("input is not a function")
var methodIsNotRegisteredError = errors.New("requested method is not registered")
var noRegistryListenerFoundError = errors.New("there is no registry for event")


type observerManager struct {
	mu sync.RWMutex
	eventRegistryMap map[reflect.Type]eventInterface
}

// Creates a new observer manager. It is expected to be singleton in the running application...
func newObserver() *observerManager {
	return &observerManager{eventRegistryMap:make(map[reflect.Type]eventInterface)}
}

func (o *observerManager) Notify(value interface{}) error {
	typeValue := reflect.ValueOf(value)
	o.mu.RLock()
	defer o.mu.RUnlock()
	eventInterface := o.eventRegistryMap[typeValue.Type()]
	if eventInterface == nil {
		return noRegistryListenerFoundError
	}
	// TODO create worker structure...
	eventInterface.notifyEvent(typeValue)
	return nil
}
func (o *observerManager) AddListener(method interface{}) error {
	err := validateInputParameter(method)
	if err != nil {
		return err
	}
	o.registerTheMethod(method)
	return nil;
}

func (o *observerManager) registerTheMethod(method interface{}) {
	funcRef := reflect.ValueOf(method)
	methodInType := funcRef.Type().In(0)
	o.mu.Lock()
	defer o.mu.Unlock()
	eventInterface := o.eventRegistryMap[methodInType]
	if eventInterface == nil {
		eventInterface = newListenerRegistry()
	}
	eventInterface.attachListener(funcRef)
	o.eventRegistryMap[methodInType] = eventInterface
}

func (o *observerManager) deRegisterTheMethod(method interface{}) error  {
	funcRef := reflect.ValueOf(method)
	methodInType := funcRef.Type().In(0)
	o.mu.RLock()
	defer o.mu.RUnlock()
	eventInterface := o.eventRegistryMap[methodInType]
	if eventInterface == nil {
		 return methodIsNotRegisteredError
	}
	eventInterface.detachListener(funcRef)
	return nil
}

func (o *observerManager) DelListener(method interface{}) error {
	err := validateInputParameter(method)
	if err != nil {
		return err
	}
	return o.deRegisterTheMethod(method);
}

func validateInputParameter(method interface{}) error  {
	funcRef := reflect.ValueOf(method)
	if  funcRef.Kind() != reflect.Func{
		return inputValueMustbeFuncError
	}
	funcRefType := funcRef.Type()
	if funcRefType.NumIn() != 1 {
		return funcMustTakeOneParameterError
	}
	return nil
}
