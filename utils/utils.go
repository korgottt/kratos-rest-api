package utils

import (
	"crypto/rand"
	"fmt"
	"log"

	"github.com/google/uuid"
)

//GenerateUUID generate random uuid
func GenerateUUID() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return uuid
}

//IsValidUUID validate uuid
func IsValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}
