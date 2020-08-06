package input

import (
	"fmt"
	"io/ioutil"
)

func ReadFakeInput() string {
	data, err := ioutil.ReadFile("../../../../../../Downloads/docker-images.in")
	if err != nil {
		fmt.Println("File reading error", err)
		panic(err)
	}
	return string(data)
}
