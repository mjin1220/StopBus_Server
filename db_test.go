package main

import (
	"log"
	"strconv"
	"testing"
	"time"
)

func TestAddUserToken(t *testing.T) {
	ret := addUserToken(User{"testToken_" + strconv.Itoa(time.Now().Nanosecond()), "testUUID_" + strconv.Itoa(time.Now().Nanosecond())})
	if ret != nil {
		log.Printf("expected %v; got %v\n", 1, ret)
	}
}

func TestAddGetIn(t *testing.T) {
	ret := addGetIn(Reserv{"testToken_2", "tRouteID1", "StationID"})
	if ret != nil {
		log.Printf("expected %v; got %v\n", 1, ret)
	}
}

// func TestGetInUserTokens(t *testing.T) {
// 	tokens, err := getGetInUserTokens("234000026", "203000066")
// 	if err != nil {
// 		t.Fatalf("%v", err)
// 	}
// 	fmt.Println(tokens)
// }
