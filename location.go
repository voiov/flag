package flag

import "errors"

var m = map[string][2]string{
	"🇦🇨": [2]string{"Ascension Island", "阿森松岛"},
	"🇦🇩": [2]string{"Andorra", "安道尔"},
	"🇦🇪": [2]string{"United Arab Emirates", "阿拉伯联合酋长国"},
	"🇦🇫": [2]string{"Afghanistan", "阿富汗"},
	"🇦🇬": [2]string{"Antigua & Barbuda", "安提瓜和巴布达"},
	"🇦🇮": [2]string{"Anguilla", "安圭拉"},
	"🇦🇱": [2]string{"Albania", "阿尔巴尼亚"},
	"🇦🇲": [2]string{"Armenia", "亚美尼亚"},
	"🇦🇴": [2]string{"Angola", "安哥拉"},
	"🇦🇶": [2]string{"Antarctica", "南极洲"},
	"🇦🇷": [2]string{"Argentina", "阿根廷"},
	"🇦🇸": [2]string{"American Samoa", "美属萨摩亚"},
	"🇦🇹": [2]string{"Austria", "奥地利"},
	"🇦🇺": [2]string{"Australia", "澳大利亚"},
	"🇦🇼": [2]string{"Aruba", "阿鲁巴"},
	"🇦🇽": [2]string{"Åland Islands", "奥兰群岛"},
	"🇦🇿": [2]string{"Azerbaijan", "阿塞拜疆"},
	"🇧🇦": [2]string{"Bosnia & Herzegovina", "波斯尼亚和黑塞哥维那"},
	"🇧🇧": [2]string{"Barbados", "巴巴多斯"},
	"🇧🇩": [2]string{"Bangladesh", "孟加拉国"},
	"🇧🇪": [2]string{"Belgium", "比利时"},
	"🇧🇫": [2]string{"Burkina Faso", "布基纳法索"},
	"🇧🇬": [2]string{"Bulgaria", "保加利亚"},
	"🇧🇭": [2]string{"Bahrain", "巴林"},
	"🇧🇮": [2]string{"Burundi", "布隆迪"},
	"🇧🇯": [2]string{"Benin", "贝宁"},
	"🇧🇱": [2]string{"St. Barthélemy", "圣巴泰勒米"},
	"🇧🇲": [2]string{"Bermuda", "百慕大"},
	"🇧🇳": [2]string{"Brunei", "文莱"},
	"🇧🇴": [2]string{"Bolivia", "玻利维亚"},
	"🇧🇶": [2]string{"Caribbean Netherlands", "加勒比荷兰"},
	"🇧🇷": [2]string{"Brazil", "巴西"},
	"🇧🇸": [2]string{"Bahamas", "巴哈马"},
	"🇧🇹": [2]string{"Bhutan", "不丹"},
	"🇧🇻": [2]string{"Bouvet Island", "布维岛"},
	"🇧🇼": [2]string{"Botswana", "博茨瓦纳"},
	"🇧🇾": [2]string{"Belarus", "白俄罗斯"},
	"🇧🇿": [2]string{"Belize", "伯利兹"},
	"🇨🇦": [2]string{"Canada", "加拿大"},
	"🇨🇨": [2]string{"Cocos (Keeling) Islands", "科科斯（基林）群岛"},
	"🇨🇩": [2]string{"Congo - Kinshasa", "刚果 - 金沙萨"},
	"🇨🇫": [2]string{"Central African Republic", "中非共和国"},
	"🇨🇬": [2]string{"Congo - Brazzaville", "刚果 - 布拉柴维尔"},
	"🇨🇭": [2]string{"Switzerland", "瑞士"},
	"🇨🇮": [2]string{"Côte D’Ivoire", "科特迪瓦"},
	"🇨🇰": [2]string{"Cook Islands", "库克群岛"},
	"🇨🇱": [2]string{"Chile", "智利"},
	"🇨🇲": [2]string{"Cameroon", "喀麦隆"},
	"🇨🇳": [2]string{"China", "中国"},
	"🇨🇴": [2]string{"Colombia", "哥伦比亚"},
	"🇨🇵": [2]string{"Clipperton Island", "克利珀顿岛"},
	"🇨🇷": [2]string{"Costa Rica", "哥斯达黎加"},
	"🇨🇺": [2]string{"Cuba", "古巴"},
	"🇨🇻": [2]string{"Cape Verde", "佛得角"},
	"🇨🇼": [2]string{"Curaçao", "库拉索岛"},
	"🇨🇽": [2]string{"Christmas Island", "圣诞岛"},
	"🇨🇾": [2]string{"Cyprus", "塞浦路斯"},
	"🇨🇿": [2]string{"Czechia", "Czechia"},
	"🇩🇪": [2]string{"Germany", "德国"},
	"🇩🇬": [2]string{"Diego Garcia", "DiegoGarcia"},
	"🇩🇯": [2]string{"Djibouti", "吉布提"},
	"🇩🇰": [2]string{"Denmark", "丹麦"},
	"🇩🇲": [2]string{"Dominica", "多米尼加"},
	"🇩🇴": [2]string{"Dominican Republic", "多米尼加共和国"},
	"🇩🇿": [2]string{"Algeria", "阿尔及利亚"},
	"🇪🇦": [2]string{"Ceuta & Melilla", "休达和梅利利亚"},
	"🇪🇨": [2]string{"Ecuador", "厄瓜多尔"},
	"🇪🇪": [2]string{"Estonia", "爱沙尼亚"},
	"🇪🇬": [2]string{"Egypt", "埃及"},
	"🇪🇭": [2]string{"Western Sahara", "西撒哈拉"},
	"🇪🇷": [2]string{"Eritrea", "厄立特里亚"},
	"🇪🇸": [2]string{"Spain", "西班牙"},
	"🇪🇹": [2]string{"Ethiopia", "埃塞俄比亚"},
	"🇪🇺": [2]string{"European Union", "欧盟"},
	"🇫🇮": [2]string{"Finland", "芬兰"},
	"🇫🇯": [2]string{"Fiji", "斐济"},
	"🇫🇰": [2]string{"Falkland Islands", "福克兰群岛"},
	"🇫🇲": [2]string{"Micronesia", "密克罗尼西亚"},
	"🇫🇴": [2]string{"Faroe Islands", "法罗群岛"},
	"🇫🇷": [2]string{"France", "法国"},
	"🇬🇦": [2]string{"Gabon", "加蓬"},
	"🇬🇧": [2]string{"United Kingdom", "英国"},
	"🇬🇩": [2]string{"Grenada", "格林纳达"},
	"🇬🇪": [2]string{"Georgia", "格鲁吉亚"},
	"🇬🇫": [2]string{"French Guiana", "法属圭亚那"},
	"🇬🇬": [2]string{"Guernsey", "根西岛"},
	"🇬🇭": [2]string{"Ghana", "加纳"},
	"🇬🇮": [2]string{"Gibraltar", "直布罗陀"},
	"🇬🇱": [2]string{"Greenland", "格陵兰岛"},
	"🇬🇲": [2]string{"Gambia", "冈比亚"},
	"🇬🇳": [2]string{"Guinea", "几内亚"},
	"🇬🇵": [2]string{"Guadeloupe", "瓜德罗普岛"},
	"🇬🇶": [2]string{"Equatorial Guinea", "赤道几内亚"},
	"🇬🇷": [2]string{"Greece", "希腊"},
	"🇬🇸": [2]string{"South Georgia & South Sandwich Islands", "南乔治亚岛和南桑威奇群岛"},
	"🇬🇹": [2]string{"Guatemala", "危地马拉"},
	"🇬🇺": [2]string{"Guam", "关岛"},
	"🇬🇼": [2]string{"Guinea-Bissau", "几内亚比绍"},
	"🇬🇾": [2]string{"Guyana", "圭亚那"},
	"🇭🇰": [2]string{"Hong Kong SAR China", "中国香港特别行政区"},
	"🇭🇲": [2]string{"Heard & McDonald Islands", "赫德和麦克唐纳群岛"},
	"🇭🇳": [2]string{"Honduras", "洪都拉斯"},
	"🇭🇷": [2]string{"Croatia", "克罗地亚"},
	"🇭🇹": [2]string{"Haiti", "海地"},
	"🇭🇺": [2]string{"Hungary", "匈牙利"},
	"🇮🇨": [2]string{"Canary Islands", "加那利群岛"},
	"🇮🇩": [2]string{"Indonesia", "印度尼西亚"},
	"🇮🇪": [2]string{"Ireland", "爱尔兰"},
	"🇮🇱": [2]string{"Israel", "以色列"},
	"🇮🇲": [2]string{"Isle of Man", "马恩岛"},
	"🇮🇳": [2]string{"India", "印度"},
	"🇮🇴": [2]string{"British Indian Ocean Territory", "英属印度洋领地"},
	"🇮🇶": [2]string{"Iraq", "伊拉克"},
	"🇮🇷": [2]string{"Iran", "伊朗"},
	"🇮🇸": [2]string{"Iceland", "冰岛"},
	"🇮🇹": [2]string{"Italy", "意大利"},
	"🇯🇪": [2]string{"Jersey", "泽西岛"},
	"🇯🇲": [2]string{"Jamaica", "牙买加"},
	"🇯🇴": [2]string{"Jordan", "乔丹"},
	"🇯🇵": [2]string{"Japan", "日本"},
	"🇰🇪": [2]string{"Kenya", "肯尼亚"},
	"🇰🇬": [2]string{"Kyrgyzstan", "吉尔吉斯斯坦"},
	"🇰🇭": [2]string{"Cambodia", "柬埔寨"},
	"🇰🇮": [2]string{"Kiribati", "基里巴斯"},
	"🇰🇲": [2]string{"Comoros", "科摩罗"},
	"🇰🇳": [2]string{"St. Kitts & Nevis", "圣基茨和尼维斯"},
	"🇰🇵": [2]string{"North Korea", "朝鲜"},
	"🇰🇷": [2]string{"South Korea", "韩国"},
	"🇰🇼": [2]string{"Kuwait", "科威特"},
	"🇰🇾": [2]string{"Cayman Islands", "开曼群岛"},
	"🇰🇿": [2]string{"Kazakhstan", "哈萨克斯坦"},
	"🇱🇦": [2]string{"Laos", "老挝"},
	"🇱🇧": [2]string{"Lebanon", "黎巴嫩"},
	"🇱🇨": [2]string{"St. Lucia", "圣卢西亚"},
	"🇱🇮": [2]string{"Liechtenstein", "列支敦士登"},
	"🇱🇰": [2]string{"Sri Lanka", "斯里兰卡"},
	"🇱🇷": [2]string{"Liberia", "利比里亚"},
	"🇱🇸": [2]string{"Lesotho", "莱索托"},
	"🇱🇹": [2]string{"Lithuania", "立陶宛"},
	"🇱🇺": [2]string{"Luxembourg", "卢森堡"},
	"🇱🇻": [2]string{"Latvia", "拉脱维亚"},
	"🇱🇾": [2]string{"Libya", "利比亚"},
	"🇲🇦": [2]string{"Morocco", "摩洛哥"},
	"🇲🇨": [2]string{"Monaco", "摩纳哥"},
	"🇲🇩": [2]string{"Moldova", "摩尔多瓦"},
	"🇲🇪": [2]string{"Montenegro", "黑山"},
	"🇲🇫": [2]string{"St. Martin", "圣马丁"},
	"🇲🇬": [2]string{"Madagascar", "马达加斯加"},
	"🇲🇭": [2]string{"Marshall Islands", "马绍尔群岛"},
	"🇲🇰": [2]string{"Macedonia", "马其顿"},
	"🇲🇱": [2]string{"Mali", "马里"},
	"🇲🇲": [2]string{"Myanmar (Burma)", "缅甸（缅甸）"},
	"🇲🇳": [2]string{"Mongolia", "蒙古"},
	"🇲🇴": [2]string{"Macau SAR China", "澳门特区中国"},
	"🇲🇵": [2]string{"Northern Mariana Islands", "北马里亚纳群岛"},
	"🇲🇶": [2]string{"Martinique", "马提尼克岛"},
	"🇲🇷": [2]string{"Mauritania", "毛里塔尼亚"},
	"🇲🇸": [2]string{"Montserrat", "蒙特塞拉特"},
	"🇲🇹": [2]string{"Malta", "马耳他"},
	"🇲🇺": [2]string{"Mauritius", "毛里求斯"},
	"🇲🇻": [2]string{"Maldives", "马尔代夫"},
	"🇲🇼": [2]string{"Malawi", "马拉维"},
	"🇲🇽": [2]string{"Mexico", "墨西哥"},
	"🇲🇾": [2]string{"Malaysia", "马来西亚"},
	"🇲🇿": [2]string{"Mozambique", "莫桑比克"},
	"🇳🇦": [2]string{"Namibia", "纳米比亚"},
	"🇳🇨": [2]string{"New Caledonia", "新喀里多尼亚"},
	"🇳🇪": [2]string{"Niger", "尼日尔"},
	"🇳🇫": [2]string{"Norfolk Island", "诺福克岛"},
	"🇳🇬": [2]string{"Nigeria", "尼日利亚"},
	"🇳🇮": [2]string{"Nicaragua", "尼加拉瓜"},
	"🇳🇱": [2]string{"Netherlands", "荷兰"},
	"🇳🇴": [2]string{"Norway", "挪威"},
	"🇳🇵": [2]string{"Nepal", "尼泊尔"},
	"🇳🇷": [2]string{"Nauru", "瑙鲁"},
	"🇳🇺": [2]string{"Niue", "纽埃"},
	"🇳🇿": [2]string{"New Zealand", "新西兰"},
	"🇴🇲": [2]string{"Oman", "阿曼"},
	"🇵🇦": [2]string{"Panama", "巴拿马"},
	"🇵🇪": [2]string{"Peru", "秘鲁"},
	"🇵🇫": [2]string{"French Polynesia", "法属波利尼西亚"},
	"🇵🇬": [2]string{"Papua New Guinea", "巴布亚新几内亚"},
	"🇵🇭": [2]string{"Philippines", "菲律宾"},
	"🇵🇰": [2]string{"Pakistan", "巴基斯坦"},
	"🇵🇱": [2]string{"Poland", "波兰"},
	"🇵🇲": [2]string{"St. Pierre & Miquelon", "圣皮埃尔和密克隆"},
	"🇵🇳": [2]string{"Pitcairn Islands", "皮特凯恩群岛"},
	"🇵🇷": [2]string{"Puerto Rico", "波多黎各"},
	"🇵🇸": [2]string{"Palestinian Territories", "巴勒斯坦领土"},
	"🇵🇹": [2]string{"Portugal", "葡萄牙"},
	"🇵🇼": [2]string{"Palau", "帕劳"},
	"🇵🇾": [2]string{"Paraguay", "巴拉圭"},
	"🇶🇦": [2]string{"Qatar", "卡塔尔"},
	"🇷🇪": [2]string{"Réunion", "留尼汪"},
	"🇷🇴": [2]string{"Romania", "罗马尼亚"},
	"🇷🇸": [2]string{"Serbia", "塞尔维亚"},
	"🇷🇺": [2]string{"Russia", "俄罗斯"},
	"🇷🇼": [2]string{"Rwanda", "卢旺达"},
	"🇸🇦": [2]string{"Saudi Arabia", "沙特阿拉伯"},
	"🇸🇧": [2]string{"Solomon Islands", "所罗门群岛"},
	"🇸🇨": [2]string{"Seychelles", "塞舌尔"},
	"🇸🇩": [2]string{"Sudan", "苏丹"},
	"🇸🇪": [2]string{"Sweden", "瑞典"},
	"🇸🇬": [2]string{"Singapore", "新加坡"},
	"🇸🇭": [2]string{"St. Helena", "圣赫勒拿"},
	"🇸🇮": [2]string{"Slovenia", "斯洛文尼亚"},
	"🇸🇯": [2]string{"Svalbard & Jan Mayen", "斯瓦尔巴和扬马延"},
	"🇸🇰": [2]string{"Slovakia", "斯洛伐克"},
	"🇸🇱": [2]string{"Sierra Leone", "塞拉利昂"},
	"🇸🇲": [2]string{"San Marino", "圣马力诺"},
	"🇸🇳": [2]string{"Senegal", "塞内加尔"},
	"🇸🇴": [2]string{"Somalia", "索马里"},
	"🇸🇷": [2]string{"Suriname", "苏里南"},
	"🇸🇸": [2]string{"South Sudan", "南苏丹"},
	"🇸🇹": [2]string{"São Tomé & Príncipe", "圣多美和普林西比"},
	"🇸🇻": [2]string{"El Salvador", "萨尔瓦多"},
	"🇸🇽": [2]string{"Sint Maarten", "圣马丁"},
	"🇸🇾": [2]string{"Syria", "叙利亚"},
	"🇸🇿": [2]string{"Swaziland", "斯威士兰"},
	"🇹🇦": [2]string{"Tristan Da Cunha", "TristanDa Cunha"},
	"🇹🇨": [2]string{"Turks & Caicos Islands", "特克斯和凯科斯群岛"},
	"🇹🇩": [2]string{"Chad", "乍得"},
	"🇹🇫": [2]string{"French Southern Territories", "法属南部领土"},
	"🇹🇬": [2]string{"Togo", "多哥"},
	"🇹🇭": [2]string{"Thailand", "泰国"},
	"🇹🇯": [2]string{"Tajikistan", "塔吉克斯坦"},
	"🇹🇰": [2]string{"Tokelau", "托克劳"},
	"🇹🇱": [2]string{"Timor-Leste", "东帝汶"},
	"🇹🇲": [2]string{"Turkmenistan", "土库曼斯坦"},
	"🇹🇳": [2]string{"Tunisia", "突尼斯"},
	"🇹🇴": [2]string{"Tonga", "汤加"},
	"🇹🇷": [2]string{"Turkey", "土耳其"},
	"🇹🇹": [2]string{"Trinidad & Tobago", "特立尼达和多巴哥"},
	"🇹🇻": [2]string{"Tuvalu", "图瓦卢"},
	"🇹🇼": [2]string{"Taiwan", "台湾"},
	"🇹🇿": [2]string{"Tanzania", "坦桑尼亚"},
	"🇺🇦": [2]string{"Ukraine", "乌克兰"},
	"🇺🇬": [2]string{"Uganda", "乌干达"},
	"🇺🇲": [2]string{"U.S. Outlying Islands", "美国离岛"},
	"🇺🇳": [2]string{"United Nations", "联合国"},
	"🇺🇸": [2]string{"United States", "美国"},
	"🇺🇾": [2]string{"Uruguay", "乌拉圭"},
	"🇺🇿": [2]string{"Uzbekistan", "乌兹别克斯坦"},
	"🇻🇦": [2]string{"Vatican City", "梵蒂冈城"},
	"🇻🇨": [2]string{"St. Vincent & Grenadines", "圣文森特和格林纳丁斯"},
	"🇻🇪": [2]string{"Venezuela", "委内瑞拉"},
	"🇻🇬": [2]string{"British Virgin Islands", "英属维尔京群岛"},
	"🇻🇮": [2]string{"U.S. Virgin Islands", "美属维尔京群岛"},
	"🇻🇳": [2]string{"Vietnam", "越南"},
	"🇻🇺": [2]string{"Vanuatu", "瓦努阿图"},
	"🇼🇫": [2]string{"Wallis & Futuna", "瓦利斯和富图纳群岛"},
	"🇼🇸": [2]string{"Samoa", "萨摩亚"},
	"🇽🇰": [2]string{"Kosovo", "科索沃"},
	"🇾🇪": [2]string{"Yemen", "也门"},
	"🇾🇹": [2]string{"Mayotte", "马约特岛"},
	"🇿🇦": [2]string{"South Africa", "南非"},
	"🇿🇲": [2]string{"Zambia", "赞比亚"},
	"🇿🇼": [2]string{"Zimbabwe", "津巴布韦"},
}

//AddNew is for that some new location added and the package has now update,you can add a new
func AddNew(code string, names [2]string) {
	e := Emoji(code)
	m[e] = names
}

//Delete delete the Map m which key equel code,and return the key,if no key,return empty string
func Delete(code string) string {
	e := Emoji(code)
	if e != "" && len(m[e]) != 0 {
		delete(m, e)
		return code
	}
	return ""
}

//Flag has some method to show a location's names
type Flag struct {
	code string
}

//Emoji return a emoji flag
func (f Flag) Emoji() string {
	return Emoji(f.code)
}

//Code return the flag's letter that in ISO 3166-1 alpha-2
func (f Flag) Code() string {
	return f.code
}

//ChineseName return flag's chinese name
func (f Flag) ChineseName() string {
	return ChineseName(f.code)
}

// EnglishName return flag's English name
func (f Flag) EnglishName() string {
	return EnglishName(f.code)
}

//New recive 2 Letter which in ISO 3166-1 alpha-2 and return a flag struct
func New(code string) (Flag, error) {
	for _, v := range code {
		if v < 'A' || v > 'z' {
			return Flag{}, errors.New("invalid countries code")
		}
		if v > 'Z' && v < 'a' {
			return Flag{}, errors.New("invalid countries code")
		}
	}
	return Flag{code}, nil
}
