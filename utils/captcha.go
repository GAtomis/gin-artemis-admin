/*
 * @Description: 验证码
 * @Author: Gavin
 * @Date: 2022-07-20 18:54:01
 * @LastEditTime: 2022-07-21 14:06:22
 * @LastEditors: Gavin
 */
package utils

import (
	"bytes"
	"net/http"
	"time"

	"github.com/dchest/captcha"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

const (
	SESSION_SECRET  = "golang-tech-stack"
	SESSION_MAX_AGE = 3600
)

func SessionConfig() sessions.Store {

	store := cookie.NewStore([]byte(SESSION_SECRET))

	store.Options(sessions.Options{
		MaxAge: SESSION_MAX_AGE,
		Path:   "/",
	})
	return store
}

/**
 * @description: 生成图片
 * @param {http.ResponseWriter} w
 * @param {*http.Request} r
 * @param {*} id
 * @param {*} ext
 * @param {string} lang
 * @param {bool} download
 * @param {*} width
 * @param {int} height
 * @return {*}
 * @Date: 2022-07-20 21:58:09
 */
func Serve(w http.ResponseWriter, r *http.Request, id, ext, lang string, download bool, width, height int) error {
	w.Header().Set("Cache-Control", "no-cache,no-store,must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	var content bytes.Buffer
	switch ext {
	case ".png":
		w.Header().Set("Content-Type", "image/png")
		_ = captcha.WriteImage(&content, id, width, height)
	case ".wav":
		w.Header().Set("Content-Type", "audio/x-wav")
		captcha.WriteAudio(&content, id, lang)
	default:
		return captcha.ErrNotFound
	}
	if download {
		w.Header().Set("Content-type", "application/octet-stream")
	}
	http.ServeContent(w, r, id+ext, time.Time{}, bytes.NewReader(content.Bytes()))
	return nil

}

func Captcha(c *gin.Context, length ...int) {
	l := captcha.DefaultLen
	w, h := 107, 36
	switch len(length) {
	case 1:
		l = length[0]
	case 2:
		w = length[1]
	case 3:
		h = length[2]
	default:
	}
	captchaId := captcha.NewLen(l)
	s := sessions.Default(c)
	s.Set("captcha", captchaId)
	_ = s.Save()
	_ = Serve(c.Writer, c.Request, captchaId, ".png", "zh", false, w, h)

}

func CaptchaVerify(c *gin.Context, code string) bool {
	s := sessions.Default(c)
	if captchaId := s.Get("captcha"); captchaId != nil {
		s.Delete("captcha")
		_ = s.Save()
		b := captcha.VerifyString(captchaId.(string), code)
		return b
	} else {
		return false
	}
}
