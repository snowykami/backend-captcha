package service

import (
	"image/color"
	"liteyuki-captcha/dao"
	"liteyuki-captcha/service/base64Captcha"
)

type redisStore struct{}

var store base64Captcha.Store = &redisStore{}

func (s redisStore) Set(id string, value string) error {
	err := dao.SetStringString(id, value)
	if err != nil {
		return err
	}
	return nil
}

func (s redisStore) Get(id string, clear bool) string {
	value, err := dao.GetStringString(id)
	if err != nil {
		return ""
	}
	return value
}

func (s redisStore) Verify(id, answer string, clear bool) bool {
	value, err := dao.GetStringString(id)
	if err != nil {
		return false
	}
	return value == answer
}

func GenerateCaptcha() (id, b64s, ans string, err error) {
	var driver base64Captcha.Driver
	driver = &base64Captcha.DriverString{
		Height:          80,
		Width:           240,
		NoiseCount:      0,
		ShowLineOptions: 0,
		Length:          4,
		Source:          "1234567890abcdefghijklmnopqrstuvwxyz",
		BgColor: &color.RGBA{
			R: 0xa2, G: 0xd8, B: 0xf4, A: 255,
		},
	}
	c := base64Captcha.NewCaptcha(driver, store)
	id, b64s, ans, err = c.Generate()
	return id, b64s, ans, err

}

func VerifyCaptcha(id, verifyValue string) bool {
	return store.Verify(id, verifyValue, true)
}
