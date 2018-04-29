package cancu

import (
	"testing"
	"reflect"
)



func TestAttachListener(t *testing.T) {
	lr := listenerRegistry{funcValues: make([]reflect.Value, 0)}
	md := MockData{}
	value := reflect.ValueOf(md.mockMethod)
	lr.attachListener(value)
	if len(lr.funcValues) == 0 {
		t.Fail()
	}
	if lr.funcValues[0].Kind() != value.Kind() {
		t.Fail()
	}
}

func TestDetachListener(t *testing.T) {
	lr := listenerRegistry{funcValues: make([]reflect.Value, 0)}
	md := MockData{}
	value := reflect.ValueOf(md.mockMethod)
	lr.attachListener(value)
	value = reflect.ValueOf(md.mockMethod)
	lr.detachListener(value)
	if len(lr.funcValues) != 0 {
		t.Fail()
	}
}

func TestNotifyEvent(t *testing.T) {
	lr := listenerRegistry{funcValues: make([]reflect.Value, 0)}
	me := MockEvent{}
	md := MockData{}
	value := reflect.ValueOf(md.mockMethod)
	valueEvent := reflect.ValueOf(me)
	lr.attachListener(value)
	lr.notifyEvent(valueEvent)
	if !md.isEventReceived {
		t.Fail()
	}
}

func TestMethodPenalty(t *testing.T)  {
	lr := listenerRegistry{funcValues: make([]reflect.Value, 0)}
	me := MockEvent{}
	md := MockData{}
	value := reflect.ValueOf(md.mockMethodSleep)
	valueEvent := reflect.ValueOf(me)
	lr.attachListener(value)
	lr.notifyEvent(valueEvent)
	if len(lr.funcValues) != 0 {
		t.Fail()
	}
}

func TestNewListenerRegistry(t *testing.T)  {
	eventInter := newListenerRegistry()
	lr := eventInter.(*listenerRegistry)
	if lr.funcValues == nil {
		t.Fail()
	}
}
