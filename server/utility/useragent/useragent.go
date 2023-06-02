// Package useragent
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package useragent

import (
	"fmt"
	"github.com/gogf/gf/v2/text/gstr"
	"hotgo/internal/consts"
	"regexp"
	"strings"
)

// GetOs 获取OS名称
func GetOs(userAgent string) string {
	osName := consts.Unknown
	if userAgent == "" {
		return osName
	}

	var (
		strRe, _   = regexp.Compile(`(?i:\((.*?)\))`)
		levelNames = ":micromessenger:dart:Windows NT:Windows Mobile:Windows Phone:Windows Phone OS:Macintosh|Macintosh:Mac OS:CrOS|CrOS:iPhone OS:iPad|iPad:OS:Android:Linux:blackberry:hpwOS:Series:Symbian:PalmOS:SymbianOS:J2ME:Sailfish:Bada:MeeGo:webOS|hpwOS:Maemo:"
		namesArr   = strings.Split(strings.Trim(levelNames, ":"), ":")
		regStrArr  = make([]string, len(namesArr))
	)

	for k, name := range namesArr {
		regStrArr[k] = fmt.Sprintf("(%s[\\s?\\/XxSs0-9_.]+)", name)
	}

	userAgent = strRe.FindString(userAgent)
	var (
		nameRe, _ = regexp.Compile(fmt.Sprintf("(?i:%s)", strings.Join(regStrArr, "|")))
		names     = nameRe.FindAllString(userAgent, -1)
		name      = ""
	)

	for _, s := range names {
		if len(name) == 0 {
			name = strings.TrimSpace(s)
		} else {
			if strings.Contains(name, "Macintosh") && s != "" {
				name = strings.TrimSpace(s)
			} else if strings.Contains(name, s) {
				name = strings.TrimSpace(s)
			} else if !strings.Contains(s, name) {
				if strings.Contains(name, "iPhone") || strings.Contains(name, "iPad") {
					s = gstr.Trim(s, "Mac OS X")
				}

				if s != "" {
					name += " " + strings.TrimSpace(s)
				}
			}
			break
		}

		if strings.Contains(name, "Windows NT") {
			name = getWinOsNameWithWinNT(name)
			break
		}
	}

	if name != "" {
		osName = name
	}
	return osName
}

// GetBrowser 获取浏览器名称
func GetBrowser(userAgent string) string {
	var (
		deviceName = consts.Unknown
		levelNames = ":VivoBrowser:QQDownload:QQBrowser:QQ:MQQBrowser:MicroMessenger:TencentTraveler:LBBROWSER:TaoBrowser:BrowserNG:UCWEB:TwonkyBeamBrowser:NokiaBrowser:OviBrowser:NF-Browser:OneBrowser:Obigo:DiigoBrowser:baidubrowser:baiduboxapp:xiaomi:Redmi:MI:Lumia:Micromax:MSIEMobile:IEMobile:EdgiOS:Yandex:Mercury:Openwave:TouchPad:UBrowser:Presto:Maxthon:MetaSr:Trident:Opera:IEMobile:Edge:Chrome:Chromium:OPR:CriOS:Firefox:FxiOS:fennec:CrMo:Safari:Nexus One:Nexus S:Nexus:Blazer:teashark:bolt:HTC:Dell:Motorola:Samsung:LG:Sony:SonyST:SonyLT:SonyEricsson:Asus:Palm:Vertu:Pantech:Fly:Wiko:i-mobile:Alcatel:Nintendo:Amoi:INQ:ONEPLUS:Tapatalk:PDA:Novarra-Vision:NetFront:Minimo:FlyFlow:Dolfin:Nokia:Series:AppleWebKit:Mobile:Mozilla:Version:"
		namesArr   = strings.Split(strings.Trim(levelNames, ":"), ":")
		regStrArr  []string
	)

	for _, name := range namesArr {
		regStrArr = append(regStrArr, fmt.Sprintf("(%s[\\s?\\/0-9.]+)", name))
	}

	var (
		regexpStr = fmt.Sprintf("(?i:%s)", strings.Join(regStrArr, "|"))
		nameRe, _ = regexp.Compile(regexpStr)
		names     = nameRe.FindAllString(userAgent, -1)
		level     = 0
	)

	for _, name := range names {
		replaceRe, _ := regexp.Compile(`(?i:[\s?\/0-9.]+)`)
		l := strings.Index(levelNames, fmt.Sprintf(":%s:", replaceRe.ReplaceAllString(name, "")))
		if level == 0 {
			deviceName = strings.TrimSpace(name)
		}

		if l >= 0 && (level == 0 || level > l) {
			level = l
			deviceName = strings.TrimSpace(name)
		}
	}
	return deviceName
}

func getWinOsNameWithWinNT(sName string) string {
	osName := "Windows"
	types := map[string]string{
		"Windows NT 10":  "Windows 10",
		"Windows NT 6.3": "Windows 8",
		"Windows NT 6.2": "Windows 8",
		"Windows NT 6.1": "Windows 7",
		"Windows NT 6.0": "Windows Vista/Server 2008",
		"Windows NT 5.2": "Windows Server 2003",
		"Windows NT 5.1": "Windows XP",
		"Windows NT 5":   "Windows 2000",
		"Windows NT 4":   "Windows NT4",
	}

	for keyWord, name := range types {
		if strings.Contains(sName, keyWord) {
			osName = name
			break
		}
	}
	return osName
}
