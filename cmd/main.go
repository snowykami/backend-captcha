package main

import (
	"liteyuki-captcha/api"
)

func main() {
	api.Run()

	select {} // block forever
}
