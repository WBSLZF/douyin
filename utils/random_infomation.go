package utils

import (
	"math/rand"
	"time"
)

var Avatar_list = []string{
	"https://tse3-mm.cn.bing.net/th/id/OIP-C.ZhAC_olTvV6po6JH3XylPwAAAA?w=160&h=180&c=7&r=0&o=5&dpr=1.5&pid=1.7",
	"https://tse2-mm.cn.bing.net/th/id/OIP-C.azOrz7gQlOYvBUQBHv9oCQAAAA?w=214&h=214&c=7&r=0&o=5&dpr=1.5&pid=1.7",
	"https://tse3-mm.cn.bing.net/th/id/OIP-C.7SXPl1PEWyhM6Xh7zHU71AAAAA?w=213&h=214&c=7&r=0&o=5&dpr=1.5&pid=1.7",
	"https://tse2-mm.cn.bing.net/th/id/OIP-C.pqV8i7F46Bgn20ciBnXG1wAAAA?w=141&h=150&c=7&r=0&o=5&dpr=1.5&pid=1.7",
	"https://tse3-mm.cn.bing.net/th/id/OIP-C.YPva28zf9U4PUGpuEM-T-AAAAA?w=205&h=205&c=7&r=0&o=5&dpr=1.5&pid=1.7",
	"https://tse2-mm.cn.bing.net/th/id/OIP-C.SGGilCfdCUzy7khzHbivwAAAAA?w=213&h=213&c=7&r=0&o=5&dpr=1.5&pid=1.7",
	"https://tse3-mm.cn.bing.net/th/id/OIP-C.YIQLigOn1X4_sfwQAZPHCgAAAA?w=197&h=197&c=7&r=0&o=5&dpr=1.5&pid=1.7",
	"https://tse4-mm.cn.bing.net/th/id/OIP-C.xrO4ZXlT0fpTlHs6_TFszQAAAA?w=181&h=181&c=7&r=0&o=5&dpr=1.5&pid=1.7",
}
var Backgroud_list = []string{
	"https://ts1.cn.mm.bing.net/th?id=OIP-C.rRdN7bapsW0ZPNfTc9H3NwHaHa&w=250&h=250&c=8&rs=1&qlt=90&o=6&dpr=1.5&pid=3.1&rm=2",
	//"https://cn.bing.com/images/search?view=detailV2&ccid=SUeJmciy&id=3F084D2BF9A2A9E2C47B9F016AA04DE1AE82BFBB&thid=OIP.SUeJmciyvAVUORj_ZNTP-AHaEo&mediaurl=https%3a%2f%2fts1.cn.mm.bing.net%2fth%2fid%2fR-C.49478999c8b2bc05543918ff64d4cff8%3frik%3du7%252bCruFNoGoBnw%26riu%3dhttp%253a%252f%252fwww.91danji.com%252fupload%252f201381%252f20130801101940490.jpg%26ehk%3dghZmF9%252bsmHq%252fKTyxrSzUtzXGSveq4c0rYQMHNOnOlFA%253d%26risl%3d%26pid%3dImgRaw%26r%3d0&exph=900&expw=1440&q=%e6%b5%b7%e8%b4%bc%e7%8e%8b%e8%83%8c%e6%99%af%e5%9b%be&simid=608050022573869859&FORM=IRPRST&ck=68A4EB659AEFFE27B72B9462B0B45BB1&selectedIndex=25",
	//"https://cn.bing.com/images/search?view=detailV2&ccid=2g9l9Utj&id=A9863A6D505C9A888B95D673197BD6DBCDC286CC&thid=OIP.2g9l9UtjmDo-2kZ9FPGQKgHaFM&mediaurl=https%3a%2f%2fpic1.zhimg.com%2fv2-7d67ce1351f17f44021f914e9260634d_r.jpg%3fsource%3d1940ef5c&exph=1348&expw=1920&q=%e6%b5%b7%e8%b4%bc%e7%8e%8b%e8%83%8c%e6%99%af%e5%9b%be&simid=608043867889754307&FORM=IRPRST&ck=DA4A4FC408692FB9D2F97E64E4ED521E&selectedIndex=6",
	"https://tse2-mm.cn.bing.net/th/id/OIP-C.HWgy1UjtrBbHnM2wMyPqsgHaEo?w=275&h=180&c=7&r=0&o=5&dpr=1.5&pid=1.7",
	//"https://tse3-mm.cn.bing.net/th/id/OIP-C.W2gIcKU70cEBp_lpbT-MpwHaEo?w=269&h=180&c=7&r=0&o=5&dpr=1.5&pid=1.7",
	//"https://tse3-mm.cn.bing.net/th/id/OIP-C.9LslcMn7V-1AMgppfCetTwHaEK?w=285&h=180&c=7&r=0&o=5&dpr=1.5&pid=1.7",
	"https://tse3-mm.cn.bing.net/th/id/OIP-C.ElgmDrjjmwX-cZ0zAaDrfgHaE8?w=241&h=181&c=7&r=0&o=5&dpr=1.5&pid=1.7",
	"https://tse2-mm.cn.bing.net/th/id/OIP-C.p9F9efpzpwU7lPkMIVRMugHaEo?w=284&h=180&c=7&r=0&o=5&dpr=1.5&pid=1.7",
}

var Signature_list = []string{
	"兄弟们给个6666666",
	"这是一个测试简介",
	"大家快来给我一个关注",
	"关注我, 就带大家上分",
	"美妆博主欢迎大家关注",
	"冲冲冲冲冲冲",
}

func Random_information() (avatar string, background string, signature string) {
	rand.Seed(time.Now().Unix())
	avatar = Avatar_list[rand.Intn(len(Avatar_list))]
	background = Backgroud_list[rand.Intn(len(Backgroud_list))]
	signature = Signature_list[rand.Intn(len(Signature_list))]
	return avatar, background, signature
}
