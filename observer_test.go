package cancu

import (
	"testing"
	"reflect"
)

func TestNewObserver(t *testing.T)  {
	
}

func TestAddListener(t *testing.T){
	md := MockData{}
	ob := observerManager{eventRegistryMap:make(map[reflect.Type]eventInterface)}
	err := ob.AddListener(md.brokenMockMethod)
	if err == nil {
		t.Error("Error had to be occurred")
		t.Fail()
	}
	err = ob.AddListener(md.mockMethod)
	if err != nil {
		t.Error("Error shouldnt be occurred")
		t.Fail()
	}
	refType := reflect.ValueOf(MockEvent{}).Type()
	if ob.eventRegistryMap[refType] == nil {
		t.Error("Shouldn't be empty")
		t.Fail()
	}
}

func TestDelListener(t *testing.T){
	md := MockData{}
	ob := observerManager{eventRegistryMap:make(map[reflect.Type]eventInterface)}
	err := ob.DelListener(md.mockMethod)
	if err == nil {
		t.Error("Shouldn't be empty")
		t.Fail()
	}
	ob.AddListener(md.mockMethodSleep)
	err = ob.DelListener(md.mockMethodSleep)
	if err != nil {
		t.Fail()
	}
}

func TestNotify(t *testing.T)  {
 	//TODO this will be done after worker structure is added.
}
