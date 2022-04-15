package useragentparser

import (
	"regexp"
)

func defaultOsParsers() []*osParser {
	return []*osParser{
		{
			Reg:     regexp.MustCompile(`(?:(?:Orca-)?Android|Adr)[ /](?:[a-z]+ )?(\d+[.\d]*)`),
			Expr:    `(?:(?:Orca-)?Android|Adr)[ /](?:[a-z]+ )?(\d+[\.\d]*)`,
			Family:  "Android",
			Version: "$1",
		},
		{
			Reg:     regexp.MustCompile(`(?:Apple-)?(?:iPhone|iPad|iPod)(?:.*Mac OS X.*Version/(\d+\.\d+)|; Opera)?`),
			Expr:    `(?:Apple-)?(?:iPhone|iPad|iPod)(?:.*Mac OS X.*Version/(\d+\.\d+)|; Opera)?`,
			Family:  "Android",
			Version: "$1",
		},
	}
}

func defaultBotParsers() []*botParser {
	return []*botParser{
		{
			Reg:    regexp.MustCompile(`360Spider(-Image|-Video)?`),
			Expr:   `360Spider(-Image|-Video)?`,
			Family: "360Spider",
		},
		{
			Reg:    regexp.MustCompile(`baiduspider(-image)?|baidu Transcoder|baidu.*spider`),
			Expr:   `baiduspider(-image)?|baidu Transcoder|baidu.*spider`,
			Family: "Baidu Spider",
		},
		{
			Reg:    regexp.MustCompile(`Sosospider|Sosoimagespider`),
			Expr:   `Sosospider|Sosoimagespider`,
			Family: "Soso Spider",
		},
		{
			Reg:    regexp.MustCompile(`Bytespider`),
			Expr:   `Bytespider`,
			Family: "Bytespider",
		},
	}
}

// nolint
func defaultDeviceParsers() []*deviceParser {
	return []*deviceParser{
		{
			Reg:   regexp.MustCompile(`(?:OB-)?OPPO[ _]?([a-zA-Z0-9]+)|N1T|R8001|OPG01|A00[12]OP|(?:X90[07][0679]|U70[57]T?|X909T?|R(?:10[01]1|2001|201[07]|6007|7005|7007|80[13579]|81[13579]|82[01379]|83[013]|800[067]|8015|810[679]|811[13]|820[057])[KLSTW]?|N520[79]|N5117|A33f|A33fw|A37fw?|(; (P[A-G{1}][A-Z0-9]+)))(?:[);/ ]|$)|R7kf|R7plusf|R7Plusm|A1601|CPH[0-9]{4}|CPH19(69|79|23|1[179])|PB(A[TM]00|CT10|BT30|CM[13]0|[FD]M00)|P(DAM10|ADM00|AF[TM]00|ADT00|AHM00|BBM[03]0|BBT00|BDT00|BFT00|[CB]E[MT]00|CA[MT]00|C[CDG]M00|CA[MT]10|[CD]PM00|CRM00|CDT00|CD[TM]10|CHM[013]0|CKM[08]0|CLM[15]0|DEM[13]0|DHM00|DK[TM]00|DPT00|DB[TM]00|DCM00|[CD]NM00|DVM00|DY[TM]20|DNT00|EA[TM]00)|Realme[ _]|(?:RMX[0-9]+|(?:OPPO[ _]?)?CPH1861)(?:[);/ ]|$)|(?:du_)?ONEPLUS|(?:A0001|A200[135]|AC200[13]|A300[03]|A3010|A5000|A600[03]|A601[03]|BE201[1235]|BE202[5689]|E100[135]|GM191[03517]|GM190[0135]|GM192[05]|HD191[013]|HD190[01357]|HD1925|IN201[013579]|IN202[0135]|KB200[01357]|LE211[07]|LE212[03])(?: Build|\))`),
			Expr:  `(?:OB-)?OPPO[ _]?([a-zA-Z0-9]+)|N1T|R8001|OPG01|A00[12]OP|(?:X90[07][0679]|U70[57]T?|X909T?|R(?:10[01]1|2001|201[07]|6007|7005|7007|80[13579]|81[13579]|82[01379]|83[013]|800[067]|8015|810[679]|811[13]|820[057])[KLSTW]?|N520[79]|N5117|A33f|A33fw|A37fw?|(; (P[A-G{1}][A-Z0-9]+)))(?:[);/ ]|$)|R7kf|R7plusf|R7Plusm|A1601|CPH[0-9]{4}|CPH19(69|79|23|1[179])|PB(A[TM]00|CT10|BT30|CM[13]0|[FD]M00)|P(DAM10|ADM00|AF[TM]00|ADT00|AHM00|BBM[03]0|BBT00|BDT00|BFT00|[CB]E[MT]00|CA[MT]00|C[CDG]M00|CA[MT]10|[CD]PM00|CRM00|CDT00|CD[TM]10|CHM[013]0|CKM[08]0|CLM[15]0|DEM[13]0|DHM00|DK[TM]00|DPT00|DB[TM]00|DCM00|[CD]NM00|DVM00|DY[TM]20|DNT00|EA[TM]00)|Realme[ _]|(?:RMX[0-9]+|(?:OPPO[ _]?)?CPH1861)(?:[);/ ]|$)|ONEPLUS`,
			Brand: "OPPO",
			ModelParsers: []*deviceModelParser{
				{
					Reg:   regexp.MustCompile(`R([0-9]{3,4}[KSTW]?)(?:[);/ ]|$)|(CPH[0-9]{4})|(?:OB-)?OPPO[ _]?([a-z0-9]+)`),
					Expr:  `R([0-9]{3,4}[KSTW]?)(?:[);/ ]|$)|(CPH[0-9]{4})|(?:OB-)?OPPO[ _]?([a-z0-9]+)`,
					Model: "$1",
				},
			},
		},
		{
			Reg:   regexp.MustCompile(`(?:vivo|VIVO|iqoo|IQOO)|(?:V\d{4}(A|T|BA|CA|BT|CT|ET|EA|GA|DT|DA|A0)|X50 Pro\+|I1927)(?:[);/ ]|$)`),
			Expr:  `(?:vivo|iqoo)|(?:V1730(D[AT]|GA)|V18(18CA|01A0|13B[AT]|18T|09[AT]|1[346][AT]|[13]8[AT]|14A|24[B]?A|2[19][AT]|3[12][AT]|36[AT])|V1731CA|V1732[AT]|V1818CT|V19[01]1[AT]|V1932[AT]|V191[3469][AT]|V192[1348]A|V193[04]A|V194[15]A|V1938CT|V1955A|V1938T|V1730EA|V19[26]2A|V196[35]A|V198[16]A|V1936A[L]?|V19[59]0A|V200[125]A|V2006|1819|V201[12]A|V202[0345]C?A|V202[235-9]|V205[47]A|V203[0268]|V2031[E]?A|V203[46]A|V204[013]|V204[6789]A|V20(6[158]|99|5[56]|66[B]?|7[23]|80)A|V2046|X50 Pro\+|I1927)(?:[);/ ]|$)`,
			Brand: "vivo",
			ModelParsers: []*deviceModelParser{
				{
					Reg:   regexp.MustCompile(`(?:VIV-|BBG-)?vivo[ _]([^/;]+) Build|(?:VIV-|BBG-)?vivo[ _]([^);/]+)(?:[);/]|$)`),
					Expr:  `(?:VIV-|BBG-)?vivo[ _]([^/;]+) Build|(?:VIV-|BBG-)?vivo[ _]([^);/]+)(?:[);/]|$)`,
					Model: "$1",
				},
			},
		},
		{
			Reg:   regexp.MustCompile(`(HW-)?(?:HUAWEI|Huawei|MediaPad T1|Ideos|HONOR[ _]?|(?:(?:AGS|AGS2|ALE|ALP|AMN|ANE|ARE|ARS|ASK|ATH|ATU|AUM|BAC|BAH[23]?|BG2|BGO|B[ZK]K|BKL|BL[ALN]|BND|BTV|CA[GMNZ]|CH[CM]|CHE[12]?|CLT|CMR|COL|COR|CPN|CRO|CRR|CUN|DIG|DLI|DRA|DUA|DUB|DUK|EDI|ELE|EML|EVA|EVR|FDR|FIG|FLA|FRD|GEM|GRA|HDN|HLK|HMA|Hol|HRY|HWI|H[36]0|INE|JAT|JDN|JDN2|JKM|JMM|JSN|KII|KIW|KNT|KOB|KSA|LDN|LEO|LIO|LLD|LND|LON|LRA|LUA|LY[AO]|MAR|MHA|MRD|MYA|NCE|NEM|NEO|NXT|PAR|PCT|PIC|PLE|PLK|POT|PRA|RIO|RNE|RVL|SCC|SCL|SCM|SEA|SHT|SLA|SNE|SPN|STF|STK|TAG|TIT|TNY|TRT|VCE|VEN|VIE|VKY|VNS|VOG|VRD|VTR|WAS|YAL|G(?:527|620S|621|630|735)|Y(?:221|330|550|6[23]5))-(?:[A-Z]{0,2}[0-9]{1,4}[a-zA-Z]{0,3}?)|H1711|U(?:8230|8500|8661|8665|8667|8800|8818|8860|9200|9508))(?:[);/ ]|$))|hi6210sft|PE-(UL00|TL[12]0|TL00M)|T1-(A21[Lw]|A23L|701u|823L)|G7-(?:L01|TL00)|HW-01K|JNY-(LX[12]|AL10)|OXF-AN[01]0|TAS-(A[LN]00|L29|TL00)|WLZ-(AL10|AN00)|NIC-LX1A|MRX-(AL09|W09)|CDY-(?:[AT]N00|AN90|NX9A)|GLK-(?:[AT]L00|LX1U)|JER-[AT]N10|ELS-(?:[AT]N[10]0|NX9|N39|N04)|AKA-(AL10|L29)|MON-(W|AL)19|BMH-AN[12]0|AQM-([AT]L[01]0|LX1)|MOA-(AL[02]0|LX9N)|NTS-AL00|ART-(?:[AT]L00[xm]|L29N?|L28)|JEF-(?:[AT]N00|AN20)|MED-(?:[AT]L00|LX9N?)|EBG-AN[01]0|ANA-(?:[AT]N00|NX9)|BZ[AK]-W00|BZT-(W09|AL[01]0)|HDL-(AL09|W09)|HWV3[123]|HW-02L|TEL-[AT]N(?:00a?|10)|KKG-AN00|MXW-AN00|JKM-AL00[ab]|TAH-(?:N29|AN00)m|C8817D|T1-821W|d-01[JK]|d-02[HK]|KRJ-W09|HWT31|Y320-U10|Y541-U02|VAT-L19|70[14]HW|60[58]HW|NOH-(?:NX9|AN00)|TNNH-AN00|LIO-(?:[TA]L00|[LN]29|AN00)|KOB2-[LW]09|PPA-LX2|AGS3-L09|DNN-LX9|NEY-NX9|LON-AL00|HLK-L41|503HW| P40 | P50 | P20 | P30 | P10 `),
			Expr:  `(HW-)?(?:HUAWEI|MediaPad T1|Ideos|HONOR[ _]?|(?:(?:AGS|AGS2|ALE|ALP|AMN|ANE|ARE|ARS|ASK|ATH|ATU|AUM|BAC|BAH[23]?|BG2|BGO|B[ZK]K|BKL|BL[ALN]|BND|BTV|CA[GMNZ]|CH[CM]|CHE[12]?|CLT|CMR|COL|COR|CPN|CRO|CRR|CUN|DIG|DLI|DRA|DUA|DUB|DUK|EDI|ELE|EML|EVA|EVR|FDR|FIG|FLA|FRD|GEM|GRA|HDN|HLK|HMA|Hol|HRY|HWI|H[36]0|INE|JAT|JDN|JDN2|JKM|JMM|JSN|KII|KIW|KNT|KOB|KSA|LDN|LEO|LIO|LLD|LND|LON|LRA|LUA|LY[AO]|MAR|MHA|MRD|MYA|NCE|NEM|NEO|NXT|PAR|PCT|PIC|PLE|PLK|POT|PRA|RIO|RNE|RVL|SCC|SCL|SCM|SEA|SHT|SLA|SNE|SPN|STF|STK|TAG|TIT|TNY|TRT|VCE|VEN|VIE|VKY|VNS|VOG|VRD|VTR|WAS|YAL|G(?:527|620S|621|630|735)|Y(?:221|330|550|6[23]5))-(?:[A-Z]{0,2}[0-9]{1,4}[A-Z]{0,3}?)|H1711|U(?:8230|8500|8661|8665|8667|8800|8818|8860|9200|9508))(?:[);/ ]|$))|hi6210sft|PE-(UL00|TL[12]0|TL00M)|T1-(A21[Lw]|A23L|701u|823L)|G7-(?:L01|TL00)|HW-01K|JNY-(LX[12]|AL10)|OXF-AN[01]0|TAS-(A[LN]00|L29|TL00)|WLZ-(AL10|AN00)|NIC-LX1A|MRX-(AL09|W09)|CDY-(?:[AT]N00|AN90|NX9A)|GLK-(?:[AT]L00|LX1U)|JER-[AT]N10|ELS-(?:[AT]N[10]0|NX9|N39|N04)|AKA-(AL10|L29)|MON-(W|AL)19|BMH-AN[12]0|AQM-([AT]L[01]0|LX1)|MOA-(AL[02]0|LX9N)|NTS-AL00|ART-(?:[AT]L00[xm]|L29N?|L28)|JEF-(?:[AT]N00|AN20)|MED-(?:[AT]L00|LX9N?)|EBG-AN[01]0|ANA-(?:[AT]N00|NX9)|BZ[AK]-W00|BZT-(W09|AL[01]0)|HDL-(AL09|W09)|HWV3[123]|HW-02L|TEL-[AT]N(?:00a?|10)|KKG-AN00|MXW-AN00|JKM-AL00[ab]|TAH-(?:N29|AN00)m|C8817D|T1-821W|d-01[JK]|d-02[HK]|KRJ-W09|HWT31|Y320-U10|Y541-U02|VAT-L19|70[14]HW|60[58]HW|NOH-(?:NX9|AN00)|TNNH-AN00|LIO-(?:[TA]L00|[LN]29|AN00)|KOB2-[LW]09|PPA-LX2|AGS3-L09|DNN-LX9|NEY-NX9|LON-AL00|HLK-L41|503HW'| P40 | P50 | P20 | P30 | P10 `,
			Brand: "Huawei",
			ModelParsers: []*deviceModelParser{
				{
					Reg:   regexp.MustCompile(`Huawei[ _\-]?([^/;]*) Build|(?:HW-)?Huawei(?:/1\.0/0?(?:Huawei))?[_\- /]?([a-z0-9\-_]+)|Huawei; ([a-z0-9 \-]+)`),
					Expr:  `Huawei[ _\-]?([^/;]*) Build|(?:HW-)?Huawei(?!Browser)(?:/1\.0/0?(?:Huawei))?[_\- /]?([a-z0-9\-_]+)|Huawei; ([a-z0-9 \-]+)`,
					Model: "$1",
				},
			},
		},
		{
			Reg:   regexp.MustCompile(`(?:iTunes-)?Apple[ _]?TV|(?:Apple-|iTunes-)?(?:iPad|iPhone)|iPh[0-9],[0-9]|CFNetwork`),
			Expr:  `(?:iTunes-)?Apple[ _]?TV|(?:Apple-|iTunes-)?(?:iPad|iPhone)|iPh[0-9],[0-9]|CFNetwork`,
			Brand: "Apple",
			ModelParsers: []*deviceModelParser{
				{
					Reg:   regexp.MustCompile(`iTunes-iPhone/[0-9]+(?:\.[0-9]+)* \(([^;]+);|(?:Apple-)?iPhone ?(3GS?|4S?|5[CS]?|6(:? Plus)?)?`),
					Expr:  `iTunes-iPhone/[0-9]+(?:\.[0-9]+)* \(([^;]+);|(?:Apple-)?iPhone ?(3GS?|4S?|5[CS]?|6(:? Plus)?)?`,
					Model: "iPhone $1",
				},
			},
		},
		{
			Reg:   regexp.MustCompile(`Xiaomi[ _]([^/;]+)(?: Build|$)|Mi9 Pro 5G|(?:(MI|Mi) [a-zA-Z0-9]+|Mi-4c|MI-One[ _]?[a-z0-9]+|MIX(?: 2S?)?)(?:[);/ ]|$)|HM (?:[^/;]+) (?:Build|MIUI)|(?:2014501|2014011|201481[12378]|201302[23]|2013061) Build|Redmi|POCOPHONE|(?:SHARK )?(KLE|KSR|MBU)-[AH]0|SK[RW]-[AH]0|PRS-[AH]0|POCO F1|DLT-[AH]0|MIBOX[234]([_ ]PRO)?|MiTV4[ACSX]?|AWM-A0|MI CC 9 Meitu Edition|MiBOX1S|M2006J10C|M2006C3(?:L[IGC]|LVG|MN?G|MT)|M2007J1(?:7[CGI]|SC)|M2002J9[SEG]|HM2014819|WT88047|210611(?:8C|19BI)|M2004J(?:7[AB]|19)C|M2012K11(?:[CGI]|A[CI])|M2011K2[CG]|M2011J18C|M2006C3[ML]II|M2003J15SC|M2007J3S[ICYGP]|M2007J22[CG]|M2103K19[CGYI]|M2101K(?:[79]AG|7AI|7B[GI]|6[GIRP]|7BNY|9[GCR])|M2010J19S[CGYIL]|M1908C3JGG|M2102(?:K1AC|K1[CG]|J2SC)|HM NOTE 1(?:LTE|W)|MI[_ ]PLAY|XIG01|Qin 1s\+|Qin 2(?: Pro)?|MI_(NOTE_Pro|5X|4i|(?:A2|8)_Lite)|A001XM|lancelot|XIG02|21061119(?:[AD]G|AL)|2107119DC|M2101K(?:7BL|9AI)|M2012K10C|M2104K10AC|M1901F71|21051182[CG]|21081111RG|2109119D[GI]|21091116(?:AI|[AU]?C|i)|220111(?:7T[GIY]|19TI)|21121(?:119S[CG]|23AC)|220111[67]SG|2107113S[GR]|M2105K81A?C|2109119BC|220112[23]C|2201117SY|joyeuse|galahad|begonia|beryllium| MDG1`),
			Expr:  `Xiaomi[ _]([^/;]+)(?: Build|$)|Mi9 Pro 5G|(?:MI [a-z0-9]+|Mi-4c|MI-One[ _]?[a-z0-9]+|MIX(?: 2S?)?)(?:[);/ ]|$)|HM (?:[^/;]+) (?:Build|MIUI)|(?:2014501|2014011|201481[12378]|201302[23]|2013061) Build|Redmi|MI_NOTE_Pro|POCOPHONE|(?:SHARK )?(KLE|MBU)-[AH]0|SKR-[AH]0|SKW-[AH]0|POCO F1|DLT-[AH]0|MIBOX[234]([_ ]PRO)?|MiTV4[CSX]?|MiTV-(MSSP[01]|AXSO0)|AWM-A0|MI CC 9 Meitu Edition|MiBOX1S|MiTV4A|M2006J10C|M2006C3(?:L[IGC]|LVG|MN?G)|M2007J1SC|M2002J9[EG]|HM2014819|WT88047|M2004J(?:7[AB]|19)C|M2006C3MII|M2003J15SC|M2007J3S[CYG]|HM NOTE 1(?:LTE|W)|MI[_ ]PLAY|XIG01`,
			Brand: "Xiaomi",
			ModelParsers: []*deviceModelParser{
				{
					Reg:   regexp.MustCompile(`(MI(?:-One)?[ _](?:[^;/]*))Build|(MI [a-z0-9]+|MI-One[ _]?[a-z0-9]+)(?:[);/ ]|$)|HM ([^/;]+) (?:Build|MIUI)|Xiaomi[ _]([^/;]+)(?: Build|$)`),
					Expr:  `(MI(?:-One)?[ _](?:[^;/]*))Build|(MI [a-z0-9]+|MI-One[ _]?[a-z0-9]+)(?:[);/ ]|$)|HM ([^/;]+) (?:Build|MIUI)|Xiaomi[ _]([^/;]+)(?: Build|$)`,
					Model: "$1",
				},
			},
		},
	}
}
