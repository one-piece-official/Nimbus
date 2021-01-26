// nolint
package useragentparser

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestDeviceBrandDetect(t *testing.T) {
	parser := NewUserAgentParser()
	cases := [][]string{
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
	}
	for _, casePair := range cases {
		brand := parser.Parse(casePair[1]).Device.Brand
		assert.Equal(t, brand, casePair[0])
	}
}
