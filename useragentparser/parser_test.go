// nolint
package useragentparser_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/one-piece-official/Nimbus/useragentparser"

	"github.com/stretchr/testify/assert"
)

//// 额外的大规模 ua 检测，不提交仅参考
//func TestDeviceBrandByTxt(t *testing.T) {
//	file, err := os.Open(fmt.Sprintf("user-agent3.txt"))
//	defer file.Close()
//
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	brandMapping := map[string]string{
//		"realme":     "oppo",
//		"meizu":      "other",
//		"honor":      "huawei",
//		"oneplus":    "oppo",
//		"nubia":      "other",
//		"blackshark": "xiaomi",
//	}
//	scanner := bufio.NewScanner(file)
//
//	scanner.Split(bufio.ScanLines)
//	parser := useragentparser.NewUserAgentParser()
//
//	for scanner.Scan() {
//		var content map[string]interface{}
//		err := json.Unmarshal(scanner.Bytes(), &content)
//		if err != nil {
//			continue
//		}
//
//		ua := content["userAgent"].(string)
//		if strings.Contains(ua, "AliXAdSDK;") ||
//			strings.Contains(ua, "TBAndroid/Native") ||
//			strings.Contains(ua, "SohuVideoMobile") ||
//			strings.Contains(ua, "ting_") ||
//			strings.Contains(ua, "Dart/2") ||
//			strings.Contains(ua, "okhttp") {
//			continue
//		}
//
//		appBrand := strings.ToLower(content["brand"].(string))
//		if content["manufacturer"] != nil {
//			appBrand = strings.ToLower(content["manufacturer"].(string))
//		}
//
//		if strings.Contains(ua, "SHARK") || strings.Contains(ua, "22041216C") {
//			appBrand = "xiaomi"
//		}
//
//		if brandMapping[appBrand] != "" {
//			appBrand = brandMapping[appBrand]
//		}
//		serverBrand := strings.ToLower(content["brand"].(string))
//		detectedBrand := strings.ToLower(parser.Parse(ua).Device.Brand)
//		//if serverBrand != "vivo" && serverBrand != "oppo" {
//		//	continue
//		//}
//		if detectedBrand != appBrand {
//			fmt.Println(ua, appBrand, detectedBrand, serverBrand)
//			if appBrand == "other" {
//				time.Sleep(10)
//			} else {
//				os.Exit(1)
//			}
//		}
//	}
//}

var cases = [][]string{
	{"Huawei", "Mozilla/5.0 (Linux; U; Android 10; zh-CN; LIO-AN00 Build/HUAWEILIO-AN00) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/78.0.3904.108 UCBrowser/13.2.2.1102 Mobile Safari/537.36"},
	{"vivo", "Mozilla/5.0 (Linux; U; Android 5.1.1; zh-CN; vivo Y51 Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/57.0.2987.108 UCBrowser/12.6.6.1046 Mobile Safari/537.36"},
	{"OPPO", "Mozilla/5.0 (Linux; Android 9; PDBM00 Build/PPR1.180610.011; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/70.0.3538.110 Mobile Safari/537.36 IqiyiApp/iqiyi IqiyiVersion/12.1.0 QYStyleModel/(light)"},
	{"OPPO", "Mozilla/5.0 (Linux; U; Android 7.1.1; zh-cn; OPPO A83 Build/N6F26Q) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/70.0.3538.80 Mobile Safari/537.36 HeyTapBrowser/10.7.16.2"},
	{"OPPO", "Mozilla/5.0 (Linux; U; Android 10; zh-cn; PDPT00 Build/QKQ1.200216.002) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/70.0.3538.80 Mobile Safari/537.36 HeyTapBrowser/40.7.16.2"},
	{"Huawei", "Mozilla/5.0 (Linux; Android 10; LIO-AL00; HMSCore 5.1.0.309) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.106 HuaweiBrowser/11.0.6.302 Mobile Safari/537.36"},
	{"Xiaomi", "Mozilla/5.0 (Linux; U; Android 10; zh-CN; MI CC 9 Meitu Edition Build/QKQ1.190828.002) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/69.0.3497.100 UWS/3.22.1.66 Mobile Safari/537.36 AliApp(Youku/9.9.3) UCBS/2.11.1.1 TTID/700159@youku_android_9.9.3 WindVane/8.5.0 Youku/9.9.3 (Android 10; Bridge_SDK; GUID 18671a55f5017ee0539dfe67df76d5dd; UTDID XbrbaPf2beoDACzWvTtM3Hzu; packageName com.youku.phone; appKey 23570660;)"},
	{"vivo", "Mozilla/5.0 (Linux; U; Android 9; zh-CN; V1913A Build/P00610) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/78.0.3904.108 UCBrowser/13.2.2.1102 Mobile Safari/537.36"},
	{"OPPO", "Mozilla/5.0 (Linux; Android 4.4.4; OPPO R7s Build/KTU84P) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/33.0.0.0 Mobile Safari/537.36"},
	{"Huawei", "Mozilla/5.0 (Linux; Android 10; VOG-AL00 Build/HUAWEIVOG-AL00; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/66.0.3359.126 MQQBrowser/6.2 TBS/045120 Mobile Safari/537.36 V1_AND_SQ_8.2.8_1346_YYB_D QQ/8.2.8.4440 NetType/WIFI WebP/0.3.0 Pixel/1080 StatusBarHeight/76 SimpleUISwitch/0 QQTheme/1000"},
	{"Xiaomi", "Mozilla/5.0 (Linux; U; Android 4.4.4; zh-cn; HM NOTE 1S Build/KTU84P) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/71.0.3578.141 Mobile Safari/537.36 XiaoMi/MiuiBrowser/10.9.2"},
	{"OPPO", "Mozilla/5.0 (Linux; Android 10; PDNM00 Build/QKQ1.200216.002; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/76.0.3809.89 Mobile Safari/537.36 T7/12.8 light/1.0 SP-engine/2.27.0 baiduboxapp/12.8.0.10 (Baidu; P1 10)"},
	{"Huawei", "Mozilla/5.0 (Linux; Android 8.0.0; EVA-AL10 Build/HUAWEIEVA-AL10; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/63.0.3239.83 Mobile Safari/537.36 T7/11.1 lite baiduboxapp/4.0.0.10 (Baidu; P1 8.0.0)"},
	{"Huawei", "Mozilla/5.0 (Linux; U; Android 10; zh-CN; TAS-AN00 Build/HUAWEITAS-AN00) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/78.0.3904.108 UCBrowser/13.2.2.1102 Mobile Safari/537.36"},
	{"Huawei", "Mozilla/5.0 (Linux; U; Android 10; zh-cn; MED-AL00 Build/HUAWEIMED-AL00) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/77.0.3865.120 MQQBrowser/11.1 Mobile Safari/537.36 COVC/045517"},
	{"Huawei", "Mozilla/5.0 (Linux; Android 9; VTR-AL00 Build/HUAWEIVTR-AL00; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/71.0.3578.99 Mobile Safari/537.36 hap/1077/huawei com.huawei.fastapp/3.1.1.300 com.yslqo.bettersaying/1.2.0 ({\x22packageName\x22:\x22com.huawei.systemmanager\x22,\x22type\x22:\x22url\x22,\x22extra\x22:\x22{}\x22})"},
	{"Huawei", "Mozilla/5.0 (Linux; U; Android 9; zh-Hans-CN; VTR-AL00 Build/HUAWEIVTR-AL00) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/57.0.2987.108 Quark/3.5.1.118 Mobile Safari/537.36"},
	{"Huawei", "Mozilla/5.0 (Linux; U; Android 9; zh-CN; COR-AL10 Build/HUAWEICOR-AL10) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/57.0.2987.108 UCBrowser/12.6.5.1045 Mobile Safari/537.36"},
	{"Xiaomi", "Dalvik/2.1.0 (Linux; U; Android 10; M2006C3LC MIUI/V12.0.12.0.QCDCNXM)"},
	{"vivo", "Mozilla/5.0 (Linux; Android 11; V2118A Build/RP1A.200720.012; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/83.0.4103.106 Mobile Safari/537.36"},
	{"vivo", "Mozilla/5.0 (Linux; Android 11; V2046A Build/RP1A.200720.012; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/83.0.4103.106 Mobile Safari/537.36"},
	{"vivo", "Mozilla/5.0 (Linux; Android 8.1.0; VIVO Build/VIVOVIVO; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/78.0.3904.108 Mobile Safari/537.36"},
	{"OPPO", "Mozilla%2F5.0+%28Linux%3B+Android+10%3B+PACT00+Build%2FQP1A.190711.020%3B+wv%29+AppleWebKit%2F537.36+%28KHTML%2C+like+Gecko%29+Version%2F4.0+Chrome%2F77.0.3865.92+Mobile+Safari%2F537.36"},
	//{"OPPO", "Mozilla%252F5.0%2B%2528Linux%253B%2BAndroid%2B8.1.0%253B%2BOPPO%2BR11%2BPluskt%2BBuild%252FOPM1.171019.011%253B%2Bwv%2529%2BAppleWebKit%252F537.36%2B%2528KHTML%252C%2Blike%2BGecko%2529%2BVersion%252F4.0%2BChrome%252F62.0.3202.84%2BMobile%2BSafari%252F537.36%2BYUEDU-NA_1080_1920_8.1.0_7.3.0.1_OPPO%2BR11%2BPluskt"},
	{"Huawei", "Mozilla/5.0 (Linux; Android 9; JSN-AL00a; HMSCore 5.3.0.312) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.93 HuaweiBrowser/11.1.1.310 Mobile Safari/537.36"},
	{"Huawei", "Mozilla/5.0 (Linux; Android 10.1; P40 Build/MRA58K; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/77.0.3865.120 MQQBrowser/6.2 TBS/045618 Mobile Safari/537.36"},
	{"OPPO", "ting_8.3.21(OPPO R11s,Android27)"},
	{"vivo", "ting_8.3.12(; V2002A Build,Android29)"},
	{"OPPO", "ting_8.3.21(; PEGM10 ,Android30)"},
	{"Huawei", "ting_9.0.39(LYA-AL00,Android29)"},
	{"Xiaomi", "Mozilla/5.0 (Linux; Android 10; M2105K81C Build/QKQ1.190828.002; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/87.0.4280.101 Mobile Safari/537.36 hap/1.9/xiaomi com.miui.hybrid/1.9.0.5 com.bqteng.enjoyphrase/2.9.3 ({\"packageName\":\"\",\"type\":\"url\",\"extra\":{}})"},
	{"Xiaomi", "Mozilla/5.0 (Linux; Android 11; M2104K10AC Build/RP1A.200720.011; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/92.0.4515.131 Mobile Safari/537.36 hap/1.9/xiaomi com.miui.hybrid/1.9.0.5 com.ernxzc.tonguetwisterart/2.1.0 ({\"packageName\":\"\",\"type\":\"url\",\"extra\":{}})"},
	{"Xiaomi", "Mozilla/5.0 (Linux; Android 11; 21091116C Build/RP1A.200720.011; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/94.0.4606.85 Mobile Safari/537.36 hap/1.9/xiaomi com.miui.hybrid/1.9.0.5 com.inyneo.magicwhodoneit/2.1.5 ({\"packageName\":\"com.sina.weibo\",\"type\":\"url\",\"extra\":{}})"},
	{"Xiaomi", "Mozilla/5.0 (Linux; Android 11; M2101K9C Build/RKQ1.201112.002; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/87.0.4280.141 Mobile Safari/537.36 hap/1.9/xiaomi com.miui.hybrid/1.9.0.5 com.guangtui.novel.full/2.1.1 ({\"packageName\":\"\",\"type\":\"url\",\"extra\":{}})"},
	{"Xiaomi", "Mozilla/5.0 (Linux; U; Android 12; zh-cn; M2104K10I Build/SP1A.210812.016) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/89.0.4389.116 Mobile Safari/537.36 XiaoMi/MiuiBrowser/16.0.22 swan-mibrowser"},
	{"OPPO", "Dalvik/2.1.0 (Linux; U; Android 10; PELM00 MIUI/V12.5.1.0.QDGCNXM)"},
	{"OPPO", "ting_9.0.39(R6007,Android18)"},
	{"Huawei", "Mozilla/5.0 (Linux; Android 10.0; TRT-LX2 Build/HUAWEITRT-LX2; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/59.0.3071.125 Mobile Safari/537.36"},
	{"vivo", "Mozilla/5.0 (Linux; Android 10; V1829A Build/QP1A.190711.020; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/78.0.3904.96 Mobile Safari/537.36"},
	{"OPPO", "Mozilla/5.0 (Linux; Android 8.1.0; OPPO R11s Build/OPM1.171019.011; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/62.0.3202.84 Mobile Safari/537.36"},
	{"vivo", "Mozilla/5.0 (Linux; Android 9; V1831A Build/P00610; wv) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3202.84 Mobile Safari/537.36 VivoBrowser/9.6.14.0"},
	{"vivo", "Dalvik/2.1.0 (Linux; U; Android 10; V2002A Build/QP1A.190711.020)"},
	{"vivo", "AliXAdSDK;10.2.23;Android;10;V2036A"},
	{"Other", "Mozilla/5.0 (Linux; Android 7.0; 1872-A0 Build/NRD90M; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/75.0.3770.156 Mobile Safari/537.36"},
	{"Xiaomi", "Mozilla/5.0 (Linux; U; Android 12; zh-cn; M2104K10I Build/SP1A.210812.016) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/89.0.4389.116 Mobile Safari/537.36 XiaoMi/MiuiBrowser/16.0.22 swan-mibrowser"},
	{"Xiaomi", "Mozilla/5.0 (Linux; Android 12; M2104K10I Build/SP1A.210812.016; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/96.0.4664.104 Mobile Safari/537.36"},
	{"Xiaomi", "Mozilla/5.0 (Linux; Android 12; 22041211AC Build/SP1A.210812.016; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/96.0.4664.104 Mobile Safari/537.36 hap/1.9/xiaomi com.miui.hybrid/1.9.0.6 com.inyneo.magicwhodoneit/2.7.0 ({\"packageName\":\"com.miui.quickappCenter\",\"type\":\"url\",\"extra\":{\"scene\":\"\"}})"},
	{"Xiaomi", "Mozilla/5.0 (Linux; Android 11; Mi9 Pro 5G Build/RKQ1.200826.002; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/86.0.4240.99 XWEB/3225 MMWEBSDK/20210601 Mobile Safari/537.36 MMWEBID/6518 MicroMessenger/8.0.7.1920(0x28000737) Process/appbrand0 WeChat/arm64 Weixin NetType/4G Language/zh_CN ABI/arm64 MiniProgramEnv/android"},
	{"Huawei", "Dalvik/2.1.0 (Linux; U; Android 10; WLZ-AN00 Build/HUAWEIWLZ-AN00)"},
	{"Other", "Mozilla/5.0 (Android;) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/102.0.5005.50 Mobile Safari/537.36 vision_mode/1 bdapp/1.0 (baiduboxvision; baiduboxvision) baiduboxapp/10.0 (Baidu; P1 9) Quark/5.6.8.212 UCBrowser/114514 BingSapphire/22.1.400120302 WuZhui/8.1.1 360 Aphone Browser (114514) HeyTapBrowser/10.7.12.0.2beta"},
}

func BenchmarkUserAgentParser(b *testing.B) {
	parser := useragentparser.NewUserAgentParser()
	size := int64(len(cases))
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		parser.Parse(cases[time.Now().UnixNano()%size][1])
	}
}

func BenchmarkUserAgentParserParallel(b *testing.B) {
	parser := useragentparser.NewUserAgentParser()
	size := int64(len(cases))

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			parser.Parse(cases[time.Now().UnixNano()%size][1])
		}
	})
}

func TestDeviceBrandDetect(t *testing.T) {
	parser := useragentparser.NewUserAgentParser()

	for _, casePair := range cases {
		ua := parser.Parse(casePair[1])
		brand := ua.Device.Brand
		fmt.Println(ua.Device.Brand, ua.Device.Model, ua.Os.Family, ua.Os.Version)
		// fmt.Println(ua.Device.Brand, ua.Os.Family, ua.Os.Version)
		// t.Error(casePair[1])
		assert.Equal(t, casePair[0], brand)
	}
}
