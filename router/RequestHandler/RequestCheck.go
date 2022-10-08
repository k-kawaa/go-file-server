package RequestHandler

import "net/http"

func RequestHeaderCheck(Request http.Request) {
	if Request.RequestURI != "127.0.0.1" {

	}
}

func Permission_Check() {

}
