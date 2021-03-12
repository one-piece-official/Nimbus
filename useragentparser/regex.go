package useragentparser

import "regexp"

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
			Reg:   regexp.MustCompile(`(?is)(?:iTunes-)?Apple[ _]?TV|(?:Apple-|iTunes-)?(?:iPad|iPhone)|iPh[0-9],[0-9]|CFNetwork`),
			Expr:  `(?is)(?:iTunes-)?Apple[ _]?TV|(?:Apple-|iTunes-)?(?:iPad|iPhone)|iPh[0-9],[0-9]|CFNetwork`,
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
			Reg:   regexp.MustCompile(`(?is)ONEPLUS`),
			Expr:  `(?is)ONEPLUS`,
			Brand: "OnePlus",
			ModelParsers: []*deviceModelParser{
				{
					Reg:   regexp.MustCompile(`(?:du_)?ONEPLUS ?([^;/]+) Build`),
					Expr:  `(?:du_)?ONEPLUS ?([^;/]+) Build`,
					Model: "$1",
				},
			},
		},
		{
			Reg:   regexp.MustCompile(`(?is)(?:vivo|iqoo|v\d{4}(a|t|ba|ca|bt|ct|et|ea|ga|dt|da|a0))`),
			Expr:  `(?is)(?:vivo|iqoo|v\d{4}(a|t|ba|ca|bt|ct|et|ea|ga|dt|da|a0))`,
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
			Reg:   regexp.MustCompile(`(?is)(HW-)?(?:Huawei|MediaPad T1|Ideos|Honor[ _]?|(?:(?:AGS|AGS2|ALE|ALP|AMN|ANE|ARE|ARS|ASK|ATH|ATU|AUM|BAC|BAH[23]?|BG2|BGO|B[ZK]K|BKL|BL[ALN]|BND|BTV|CA[GMNZ]|CH[CM]|CHE[12]?|CLT|CMR|COL|COR|CPN|CRO|CRR|CUN|DIG|DLI|DRA|DUA|DUB|DUK|EDI|ELE|EML|EVA|EVR|FDR|FIG|FLA|FRD|GEM|GRA|HDN|HLK|HMA|Hol|HRY|HWI|H[36]0|INE|JAT|JDN|JDN2|JKM|JMM|JSN|KII|KIW|KNT|KOB|KSA|LDN|LEO|LIO|LLD|LND|LON|LRA|LUA|LY[AO]|MAR|MHA|MRD|MYA|NCE|NEM|NEO|NXT|PAR|PCT|PIC|PLE|PLK|POT|PRA|RIO|RNE|RVL|SCC|SCL|SCM|SEA|SHT|SLA|SNE|SPN|STF|STK|TAG|TIT|TNY|TRT|VCE|VEN|VIE|VKY|VNS|VOG|VRD|VTR|WAS|YAL|G(?:527|620S|621|630|735)|Y(?:221|330|550|6[23]5))-(?:[A-Z]{0,2}[0-9]{1,4}[A-Z]{0,3}?)|H1711|U(?:8230|8500|8661|8665|8667|8800|8818|8860|9200|9508))(?:[);/ ]|$))|hi6210sft|PE-(UL00|TL[12]0|TL00M)|T1-(A21[Lw]|A23L|701u|823L)|G7-(?:L01|TL00)|HW-01K|JNY-(LX[12]|AL10)|OXF-AN[01]0|TAS-(A[LN]00|L29|TL00)|WLZ-(AL10|AN00)|NIC-LX1A|MRX-(AL09|W09)|CDY-(?:[AT]N00|AN90|NX9A)|GLK-(?:[AT]L00|LX1U)|JER-[AT]N10|ELS-(?:[AT]N[10]0|NX9|N39|N04)|AKA-(AL10|L29)|MON-(W|AL)19|BMH-AN[12]0|AQM-([AT]L[01]0|LX1)|MOA-(AL[02]0|LX9N)|NTS-AL00|ART-(?:[AT]L00[xm]|L29N?|L28)|JEF-(?:[AT]N00|AN20)|MED-(?:[AT]L00|LX9N?)|EBG-AN[01]0|ANA-(?:[AT]N00|NX9)|BZ[AK]-W00|BZT-(W09|AL[01]0)|HDL-(AL09|W09)|HWV3[123]|HW-02L|TEL-[AT]N(?:00a?|10)|KKG-AN00|MXW-AN00|JKM-AL00[ab]|TAH-(?:N29|AN00)m|C8817D|T1-821W|d-01[JK]|d-02[HK]|KRJ-W09|HWT31|Y320-U10|Y541-U02|VAT-L19|70[14]HW|60[58]HW|NOH-(?:NX9|AN00)|TNNH-AN00|LIO-(?:[TA]L00|[LN]29|AN00)|KOB2-[LW]09|PPA-LX2|AGS3-L09|DNN-LX9|NEY-NX9|LON-AL00|HLK-L41|503HW`),
			Expr:  `(?is)(HW-)?(?:Huawei|MediaPad T1|Ideos|Honor[ _]?|(?:(?:AGS|AGS2|ALE|ALP|AMN|ANE|ARE|ARS|ASK|ATH|ATU|AUM|BAC|BAH[23]?|BG2|BGO|B[ZK]K|BKL|BL[ALN]|BND|BTV|CA[GMNZ]|CH[CM]|CHE[12]?|CLT|CMR|COL|COR|CPN|CRO|CRR|CUN|DIG|DLI|DRA|DUA|DUB|DUK|EDI|ELE|EML|EVA|EVR|FDR|FIG|FLA|FRD|GEM|GRA|HDN|HLK|HMA|Hol|HRY|HWI|H[36]0|INE|JAT|JDN|JDN2|JKM|JMM|JSN|KII|KIW|KNT|KOB|KSA|LDN|LEO|LIO|LLD|LND|LON|LRA|LUA|LY[AO]|MAR|MHA|MRD|MYA|NCE|NEM|NEO|NXT|PAR|PCT|PIC|PLE|PLK|POT|PRA|RIO|RNE|RVL|SCC|SCL|SCM|SEA|SHT|SLA|SNE|SPN|STF|STK|TAG|TIT|TNY|TRT|VCE|VEN|VIE|VKY|VNS|VOG|VRD|VTR|WAS|YAL|G(?:527|620S|621|630|735)|Y(?:221|330|550|6[23]5))-(?:[A-Z]{0,2}[0-9]{1,4}[A-Z]{0,3}?)|H1711|U(?:8230|8500|8661|8665|8667|8800|8818|8860|9200|9508))(?:[);/ ]|$))|hi6210sft|PE-(UL00|TL[12]0|TL00M)|T1-(A21[Lw]|A23L|701u|823L)|G7-(?:L01|TL00)|HW-01K|JNY-(LX[12]|AL10)|OXF-AN[01]0|TAS-(A[LN]00|L29|TL00)|WLZ-(AL10|AN00)|NIC-LX1A|MRX-(AL09|W09)|CDY-(?:[AT]N00|AN90|NX9A)|GLK-(?:[AT]L00|LX1U)|JER-[AT]N10|ELS-(?:[AT]N[10]0|NX9|N39|N04)|AKA-(AL10|L29)|MON-(W|AL)19|BMH-AN[12]0|AQM-([AT]L[01]0|LX1)|MOA-(AL[02]0|LX9N)|NTS-AL00|ART-(?:[AT]L00[xm]|L29N?|L28)|JEF-(?:[AT]N00|AN20)|MED-(?:[AT]L00|LX9N?)|EBG-AN[01]0|ANA-(?:[AT]N00|NX9)|BZ[AK]-W00|BZT-(W09|AL[01]0)|HDL-(AL09|W09)|HWV3[123]|HW-02L|TEL-[AT]N(?:00a?|10)|KKG-AN00|MXW-AN00|JKM-AL00[ab]|TAH-(?:N29|AN00)m|C8817D|T1-821W|d-01[JK]|d-02[HK]|KRJ-W09|HWT31|Y320-U10|Y541-U02|VAT-L19|70[14]HW|60[58]HW|NOH-(?:NX9|AN00)|TNNH-AN00|LIO-(?:[TA]L00|[LN]29|AN00)|KOB2-[LW]09|PPA-LX2|AGS3-L09|DNN-LX9|NEY-NX9|LON-AL00|HLK-L41|503HW'`,
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
			Reg:   regexp.MustCompile(`(?is)Xiaomi[ _]([^/;]+)(?: Build|$)|Mi9 Pro 5G|(?:MI [a-z0-9]+|Mi-4c|MI-One[ _]?[a-z0-9]+|MIX(?: 2S?)?)(?:[);/ ]|$)|HM (?:[^/;]+) (?:Build|MIUI)|(?:2014501|2014011|201481[12378]|201302[23]|2013061) Build|Redmi|MI_NOTE_Pro|POCOPHONE|(?:SHARK )?(KLE|MBU)-[AH]0|SKR-[AH]0|SKW-[AH]0|POCO F1|DLT-[AH]0|MIBOX[234]([_ ]PRO)?|MiTV4[CSX]?|MiTV-(MSSP[01]|AXSO0)|AWM-A0|MI CC 9 Meitu Edition|MiBOX1S|MiTV4A|M2006J10C|M2006C3(?:L[IGC]|LVG|MN?G)|M2007J1SC|M2002J9[EG]|HM2014819|WT88047|M2004J(?:7[AB]|19)C|M2006C3MII|M2003J15SC|M2007J3S[CYG]|HM NOTE 1(?:LTE|W)|MI[_ ]PLAY|XIG01`),
			Expr:  `(?is)Xiaomi[ _]([^/;]+)(?: Build|$)|Mi9 Pro 5G|(?:MI [a-z0-9]+|Mi-4c|MI-One[ _]?[a-z0-9]+|MIX(?: 2S?)?)(?:[);/ ]|$)|HM (?:[^/;]+) (?:Build|MIUI)|(?:2014501|2014011|201481[12378]|201302[23]|2013061) Build|Redmi|MI_NOTE_Pro|POCOPHONE|(?:SHARK )?(KLE|MBU)-[AH]0|SKR-[AH]0|SKW-[AH]0|POCO F1|DLT-[AH]0|MIBOX[234]([_ ]PRO)?|MiTV4[CSX]?|MiTV-(MSSP[01]|AXSO0)|AWM-A0|MI CC 9 Meitu Edition|MiBOX1S|MiTV4A|M2006J10C|M2006C3(?:L[IGC]|LVG|MN?G)|M2007J1SC|M2002J9[EG]|HM2014819|WT88047|M2004J(?:7[AB]|19)C|M2006C3MII|M2003J15SC|M2007J3S[CYG]|HM NOTE 1(?:LTE|W)|MI[_ ]PLAY|XIG01`,
			Brand: "Xiaomi",
			ModelParsers: []*deviceModelParser{
				{
					Reg:   regexp.MustCompile(`(MI(?:-One)?[ _](?:[^;/]*))Build|(MI [a-z0-9]+|MI-One[ _]?[a-z0-9]+)(?:[);/ ]|$)|HM ([^/;]+) (?:Build|MIUI)|Xiaomi[ _]([^/;]+)(?: Build|$)`),
					Expr:  `(MI(?:-One)?[ _](?:[^;/]*))Build|(MI [a-z0-9]+|MI-One[ _]?[a-z0-9]+)(?:[);/ ]|$)|HM ([^/;]+) (?:Build|MIUI)|Xiaomi[ _]([^/;]+)(?: Build|$)`,
					Model: "$1",
				},
			},
		},
		{
			Reg:   regexp.MustCompile(`(?is)(?:OB-)?OPPO[ _]?([a-z0-9]+)|N1T|R8001|OPG01|A00[12]OP|(?:X90[07][0679]|U70[57]T?|X909T?|R(?:10[01]1|2001|201[07]|6007|7005|7007|80[13579]|81[13579]|82[01379]|83[013]|800[067]|8015|810[679]|811[13]|820[057])[KLSTW]?|N520[79]|N5117|A33f|A33fw|A37fw?|(; (P[A-G{1}][A-Z0-9]+)))(?:[);/ ]|$)|R7kf|R7plusf|R7Plusm|A1601|CPH[0-9]{4}|CPH19(69|79|23|1[179])|PB(A[TM]00|CT10|BT30|CM[13]0|[FD]M00)|P(DAM10|ADM00|AF[TM]00|ADT00|AHM00|BBM[03]0|BBT00|BDT00|BFT00|[CB]E[MT]00|CA[MT]00|C[CDG]M00|CA[MT]10|[CD]PM00|CRM00|CDT00|CD[TM]10|CHM[013]0|CKM[08]0|CLM[15]0|DEM[13]0|DHM00|DK[TM]00|DPT00|DB[TM]00|DCM00|[CD]NM00|DVM00|DY[TM]20|DNT00|EA[TM]00)|Realme[ _]|(?:RMX[0-9]+|(?:OPPO[ _]?)?CPH1861)(?:[);/ ]|$)`),
			Expr:  `(?is)(?:OB-)?OPPO[ _]?([a-z0-9]+)|N1T|R8001|OPG01|A00[12]OP|(?:X90[07][0679]|U70[57]T?|X909T?|R(?:10[01]1|2001|201[07]|6007|7005|7007|80[13579]|81[13579]|82[01379]|83[013]|800[067]|8015|810[679]|811[13]|820[057])[KLSTW]?|N520[79]|N5117|A33f|A33fw|A37fw?|(; (P[A-G{1}][A-Z0-9]+)))(?:[);/ ]|$)|R7kf|R7plusf|R7Plusm|A1601|CPH[0-9]{4}|CPH19(69|79|23|1[179])|PB(A[TM]00|CT10|BT30|CM[13]0|[FD]M00)|P(DAM10|ADM00|AF[TM]00|ADT00|AHM00|BBM[03]0|BBT00|BDT00|BFT00|[CB]E[MT]00|CA[MT]00|C[CDG]M00|CA[MT]10|[CD]PM00|CRM00|CDT00|CD[TM]10|CHM[013]0|CKM[08]0|CLM[15]0|DEM[13]0|DHM00|DK[TM]00|DPT00|DB[TM]00|DCM00|[CD]NM00|DVM00|DY[TM]20|DNT00|EA[TM]00)|Realme[ _]|(?:RMX[0-9]+|(?:OPPO[ _]?)?CPH1861)(?:[);/ ]|$)`,
			Brand: "OPPO",
			ModelParsers: []*deviceModelParser{
				{
					Reg:   regexp.MustCompile(`R([0-9]{3,4}[KSTW]?)(?:[);/ ]|$)|(CPH[0-9]{4})|(?:OB-)?OPPO[ _]?([a-z0-9]+)`),
					Expr:  `R([0-9]{3,4}[KSTW]?)(?:[);/ ]|$)|(CPH[0-9]{4})|(?:OB-)?OPPO[ _]?([a-z0-9]+)`,
					Model: "$1",
				},
			},
		},
	}
}
