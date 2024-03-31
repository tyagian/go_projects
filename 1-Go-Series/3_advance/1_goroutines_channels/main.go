package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	id := getUserByName("John")
	println(id)
	chats := getUserChats(id)
	friends := getUserFriends(id)

	log.Println(chats)
	log.Println(friends)
}

func getUserFriends(id string) []string {
	time.Sleep(time.Second * 1)

	return []string{
		"john",
		"joe",
		"Jane",
		"Tiago",
	}
}

func getUserChats(id string) []string {
	time.Sleep(time.Second * 2)
	return []string{
		"john",
		"jane",
		"joe",
	}
}

func getUserByName(name string) string {
	time.Sleep(time.Second * 1)

	return fmt.Sprintf("%s-2", name)
}
