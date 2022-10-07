package RequestHandler

import (
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, Request *http.Request) {
	fmt.Println("ahaa")
}
