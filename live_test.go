package live

import (
	"fmt"
	"testing"
	"time"
)

func TestMongo(t *testing.T) {
	fmt.Println("Test Check for MongoDb Service")

	svc := NewService()
	svc.Ping = func() *Ping {
		p := new(Ping)
		p.Type = "HttpResponse"
		p.Command = "localhost:27123"
		return p
	}()

	svc.KeepAlive()

	done := false
	time.Sleep(100 * time.Minute)
}
