package app

import (
	"fmt"
	"time"

	"github.com/pkg/errors"
	"github.com/yann0917/dedao/config"
	"github.com/yann0917/dedao/services"
	"github.com/yann0917/dedao/utils"
)

// LoginByCookie login by cookie
func LoginByCookie(cookie string) (err error) {
	var u config.Dedao
	err = services.ParseCookies(cookie, &u.CookieOptions)
	if err != nil {
		return
	}
	// save config
	u.CookieStr = cookie
	config.Instance.SetUser(&u)
	config.Instance.Save()
	return
}

func LoginByQr(isTerminal bool) (code *services.QrCodeResp, err error) {
	token, err := getService().LoginAccessToken()
	if err != nil {
		return
	}
	// fmt.Printf("token:%#v\n", token)
	code, err = getService().GetQrcode(token)
	if err != nil {
		return
	}

	if isTerminal {
		content := "https://m.igetget.com/oauth/qrcode/v2/authorize?token=" + code.Data.QrCodeString
		obj := utils.NewQrCodeTerminal()
		obj.Get(content).Print()
		ticker := time.NewTicker(time.Duration(1) * time.Second)
		fmt.Println("同时支持「得到App」和「微信」扫码")
		for {
			select {
			case <-ticker.C:
				err = CheckLogin(token, code.Data.QrCodeString)
				if err != nil {
					return
				}
			// 10分钟后二维码失效
			case <-time.After(600 * time.Second):
				err = errors.New("登录失败，二维码已过期")
				return
			}
		}
	}
	return
}

// CheckLogin 需要开启定时器轮询是否扫码登录
func CheckLogin(token, qrCode string) error {
	check, cookie, err := getService().CheckLogin(token, qrCode)
	if err != nil {
		return err
	}
	if check.Data.Status == 1 {
		err = LoginByCookie(cookie)
		if err != nil {
			return err
		}
		fmt.Println("扫码成功")
	} else if check.Data.Status == 2 {
		err = errors.New("登录失败，二维码已过期")
		return err
	}
	return nil
}

func SwitchAccount(uid string) (err error) {
	if config.Instance.LoginUserCount() == 0 {
		err = errors.New("cannot found account's")
		return
	}
	err = config.Instance.SwitchUser(&config.User{UIDHazy: uid})

	if err != nil {
		return err
	}
	return
}
