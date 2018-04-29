package cancu

import (
	"sync"
)



var observe Observer;
var once sync.Once

func Cancu() Observer  {
	once.Do(func() {
		observe = newObserver()
	})
	return observe
}


