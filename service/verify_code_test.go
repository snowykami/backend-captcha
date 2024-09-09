package service

import (
	"encoding/base64"
	"fmt"
	"git.liteyuki.icu/backend/golite/logger"
	"os"
	"testing"
)

func TestGenerateCaptcha(t *testing.T) {
	id, b64, ans, err := GenerateCaptcha()
	logger.Debug("id: ", id, " ans: ", ans, " b64: ", b64, " err: ", err)
	filename := fmt.Sprintf("captcha_%s.png", id)
	file, _ := os.Create(filename)
	imageData, _ := base64.StdEncoding.DecodeString(b64[22:])
	_, err = file.Write(imageData)
	if err != nil {
		return
	}
}

func TestVerifyCaptcha(t *testing.T) {
	id, b64, ans, err := GenerateCaptcha()
	logger.Debug("id: ", id, " ans: ", ans, " b64: ", b64, " err: ", err)
	if VerifyCaptcha(id, ans) {
		logger.Debug("Succeed! correct answer passed")
	} else {
		t.Error("Failed! correct answer not passed")
	}

	if VerifyCaptcha(id, "1234") {
		t.Error("Failed! error answer passed")
	} else {
		logger.Debug("Succeed! error answer not passed")
	}
}
