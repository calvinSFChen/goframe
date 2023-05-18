package utils

import "regexp"

func IsPhone(phone string) bool {

	regRuler := `^(13[0-9]|14[01456879]|15[0-35-9]|16[2567]|17[0-8]|18[0-9]|19[0-35-9])\d{8}$`
	// 正则调用规则
	reg := regexp.MustCompile(regRuler)
	// 返回 MatchString 是否匹配
	return reg.MatchString(phone)
}

func IsEmail(email string) bool {
	// 正则调用规则
	regRuler := `^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`
	reg := regexp.MustCompile(regRuler)
	return reg.MatchString(email)
}
