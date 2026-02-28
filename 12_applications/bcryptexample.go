package _2_applications

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func BcryptExample() {

	bs, err := bcrypt.GenerateFromPassword([]byte("Welcome@123"), bcrypt.MinCost)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println(bs)

	err = bcrypt.CompareHashAndPassword(bs, []byte("Welcome@123"))

	if err != nil {
		fmt.Println(err)
	}

}
