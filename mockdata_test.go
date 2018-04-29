package cancu

import "time"
type MockData struct {
	isEventReceived bool
}

type MockEvent struct {
	count int
}

func (md *MockData) mockMethod(mock MockEvent) {
	md.isEventReceived = true
}

func (md *MockData) mockMethodSleep(mock MockEvent) {
	md.isEventReceived = true
	time.Sleep(time.Second * 3)
}

func (md *MockData) mockMethodPointer(mock * MockEvent)  {
	md.isEventReceived = true
}

func (md *MockData) brokenMockMethod()  {

}
