package queue

import (
	"fmt"
	"testing"
	"time"
)

func TestRPushQueue(t *testing.T) {

	ll := NewQueue()

	ll.RPush("1")
	ll.RPush("2")
	ll.RPush("3")

	go func() {
		ll.RPush("4")
	}()
	go func() {
		ll.RPush("5")
	}()

	go func() {
		ll.RPush("6")
	}()

	time.Sleep(1 * time.Second)

	if ll.Len() != 6 {
		t.Error("list Len() do error #1")
	}

	listVal := fmt.Sprintf("num=>%v,%v,%v", ll.LPop(), ll.LPop(), ll.LPop())
	if listVal != "num=>1,2,3" {
		t.Error("list do error #2")
	}

	if ll.Len() != 3 {
		t.Error("list Len() do error #3")
	}

	ll.LPop()
	ll.LPop()
	ll.LPop()
	c := ll.LPop()

	if c != nil {
		t.Error("list LPop() do error #4")
	}

	time.Sleep(1 * time.Second)
}

func TestLPushQueue(t *testing.T) {

	ll := NewQueue()

	ll.LPush("1")
	ll.LPush("2")
	ll.LPush("3")

	go func() {
		ll.LPush("4")
	}()
	go func() {
		ll.LPush("5")
	}()

	go func() {
		ll.LPush("6")
	}()

	time.Sleep(1 * time.Second)

	if ll.Len() != 6 {
		t.Error("list Len() do error #1")
	}

	listVal := fmt.Sprintf("num=>%v,%v,%v", ll.RPop(), ll.RPop(), ll.RPop())
	if listVal != "num=>1,2,3" {
		t.Error("list do error #2")
	}

	if ll.Len() != 3 {
		t.Error("list Len() do error #3")
	}

	ll.RPop()
	ll.RPop()
	ll.RPop()
	c := ll.RPop()

	if c != nil {
		t.Error("list RPop() do error #4")
	}

	time.Sleep(1 * time.Second)
}

func TestRegisterRocketMqProducer(t *testing.T) {
	ins, err := RegisterRocketMqProducer([]string{}, "tests", 2)
	if err == nil {
		t.Error("RegisterRocketMqProducer err #1")
	}

	ins, err = RegisterRocketMqProducer([]string{"192.168.1.1:9876"}, "tests", 2)
	if err != nil {
		t.Error("RegisterRocketMqProducer err #2")
	}

	if ins.endPoints[0] != "192.168.1.1:9876" {
		t.Error("RegisterRocketMqProducer err #3")
	}

}

func TestRegisterRocketMqConsumer(t *testing.T) {
	ins, err := RegisterRocketMqConsumer([]string{}, "tests")
	if err == nil {
		t.Error("RegisterRocketMqConsumer err #1")
	}

	ins, err = RegisterRocketMqProducer([]string{"192.168.1.1:9876"}, "tests", 2)
	if err != nil {
		t.Error("RegisterRocketMqConsumer err #2")
	}

	if ins.endPoints[0] != "192.168.1.1:9876" {
		t.Error("RegisterRocketMqConsumer err #3")
	}

}
