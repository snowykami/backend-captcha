package main

import (
	"git.liteyuki.icu/backend/captcha/api"
)

func main() {
	api.Run()

	select {} // block forever
}
