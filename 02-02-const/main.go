package main

import "fmt"

const (
	BEIJING  = 0
	SHANGHAI = 1
	SHENZHEN = 2
)

const (
	STATE_INITIAL = iota
	STATE_PRECCED
	STATE_FINISHED
)

const (
	_                                  = iota
	HTTP_PROCESS, HTTP_SWITCH_PROTOCOL = 100 * iota, 100*iota + 1
	HTTP_SUCCESS, HTTP_CREATED
	HTTP_REDIRECT, HTTP_REDIRECT_PER
	HTTP_USER_INPUT_ERROR, HTTP_USER_UNAUTHED
	HTTP_SERVER_ERROR, HTTP_SERVER_FAILED
)

const (
	a, b = iota + 1, iota + 2 // iota
)

func main() {
	const length int = 10

	fmt.Println("length =", length)
	fmt.Println(
		HTTP_PROCESS, HTTP_SWITCH_PROTOCOL,
		HTTP_SUCCESS, HTTP_CREATED,
		HTTP_REDIRECT, HTTP_REDIRECT_PER,
		HTTP_USER_INPUT_ERROR, HTTP_USER_UNAUTHED,
		HTTP_SERVER_ERROR, HTTP_SERVER_FAILED,
	)

}
