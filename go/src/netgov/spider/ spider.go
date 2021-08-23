package spider

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/gocolly/colly"
)

func SendOrder(order map[string]interface{}) bool {
	c := colly.NewCollector()
	Objorder := new(Order)
	Objorder.Assignment.Operation = "route"
	if order["tag"] == "department" {
		Objorder.Assignment.ResponseAttributes.EntriesAttributes = EntriesAttributes{
			{FieldID: 15323, Value: order["order"]},
			{FieldID: 15313, Value: order["classification"]},
			{FieldID: 15543, Value: order["department"]},
			{FieldID: 15317, Value: "市长信箱", OptionID: 21555},
			{FieldID: 15321, Value: order["name"]},
			{FieldID: 15659, Value: order["phone"]},
			{FieldID: 15320, Value: order["comment"]},
			{FieldID: 15315, Value: "注：若电话联系当事人时，当事人电话处于关机或无人接听、无法接通等状态时，要求间隔4个小时后再次联系，分3次联系，若连续3次都未联系上时，分别注明联系的时间、拨打时的电话、以及2名经办人的姓名进行特别说明。请注意来话人个人信息保密！", OptionID: 21553},
		}
	}

	if order["tag"] == "community" {
		Objorder.Assignment.ResponseAttributes.EntriesAttributes = EntriesAttributes{
			{FieldID: 15323, Value: order["order"]},
			{FieldID: 15313, Value: order["classification"]},
			{FieldID: 15544, Value: order["community"]},
			{FieldID: 15317, Value: "市长信箱", OptionID: 21555},
			{FieldID: 15321, Value: order["name"]},
			{FieldID: 15659, Value: order["phone"]},
			{FieldID: 15320, Value: order["comment"]},
			{FieldID: 15315, Value: "注：若电话联系当事人时，当事人电话处于关机或无人接听、无法接通等状态时，要求间隔4个小时后再次联系，分3次联系，若连续3次都未联系上时，分别注明联系的时间、拨打时的电话、以及2名经办人的姓名进行特别说明。请注意来话人个人信息保密！", OptionID: 21553},
		}
		Objorder.Assignment.ResponseAttributes.EntriesAttributes = EntriesAttributes{
			{FieldID: 15544, Value: order["community"]},
		}
	}

	data1, err := json.Marshal(Objorder)
	if err != nil {
		log.Fatal("json marshal err")
	}
	Objordertime := new(OrderTime)
	Objordertime.Assignment.NextVertexID = 16069
	Objordertime.Assignment.DurationThresholds = DurationThresholds{
		{Gid: "gid://skylark/YetAnotherWorkflow::Graph/1775", Value: "P3D"},
	}
	Objordertime.Assignment.Operation = "propose"

	data2, err := json.Marshal(Objordertime)
	if err != nil {
		log.Fatal("json marshal err")
	}

	c.OnResponse(func(r *colly.Response) {
		fmt.Print(r.StatusCode)
	})

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Add("Accept", "application/json, text/javascript")
		r.Headers.Add("Content-Type", "application/json")
		r.Headers.Add("X-Requested-With", "XMLHttpRequest")
		r.Headers.Add("Cookie", "last_ns_id=MQ%3D%3D--2c3fc9585e8343b352b1674c724da826c16ae033; ns_remember_tokens=IntcIjFcIjpcInRJVk9pcmtHbnNqcFlqcG9xSTNuXCJ9Ig%3D%3D--bec41a281c8fd1cce464bb452dd59f04cfafff25; yfx_c_g_u_id_10002800=_ck21070509470314714238792733215; yfx_f_l_v_t_10002800=f_t_1625449623464__r_t_1625734506229__v_t_1625734506229__r_c_2; current_user_id=MTU2NTE%3D--b8cca0386bad5d172641686265da45e62480192a; _skylark_session=c2VQWjF6OER2Nkp4L25QYUF6UXV6aTRrbTFKZVd6S0ptbG5xOWRCZ1FFTnhhNXFYN0FRTGJONmZxZ3Z0a1FKWjdYdGZIb1RWV0VQNWRiV2l3WTlVMzVucTZlbXVPZWFtelFLYlplOERsZHRFdnNIZTlDNUdzdS95Nm00Z2I0T0xiTWltNGtvSVFGbE1wVCtTYW9RR21IYW9hZ0lVeTVmWHZlcHgvOTVmTUFyT0ovYzNFZ3RLanpTbGczQnVGb1psLS14K2tJSGFGbHBpcGI5UURFNEhnWnR3PT0%3D--af7a4937d3087399dc2bb755d0e014b796de3538")
		r.Headers.Add("X-CSRF-Token", "AGEeN2n/r5vwtvFP4ipY50+Wpcshmgk+nsjsDzaKA9IXJWipJw9JnHzYdNvs9XXCrsa3MLgCdJpbJXHCfgqldA==")
		r.Headers.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.97 Safari/537.36")

	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println(r.StatusCode)
		fmt.Println(string(r.Body))
		log.Println("访问失败!")
	})

	Dealorder = append(Dealorder, order["order"])
	c.PostRaw("https://gxzh.cdht.gov.cn/namespaces/1/yet_another_workflow/flows/1009/journeys", data1)
	c.PostRaw("https://gxzh.cdht.gov.cn/namespaces/1/yet_another_workflow/flows/1009/journeys", data2)

	return true
}

// Deal Order
func DealOrder(order map[string]interface{}) map[string]interface{} {
	Classification := []interface{}{"消费纠纷", "工资发放", "房屋中介管理", "其他", "占道经营", "农民工权益", "物业服务", "噪音污染", "城镇职工社会保险", "协议解除劳动关系", "油烟扰民", "基层事务管理",
		"商品质量", "劳动合同纠纷", "营业执照、企业年报等", "疾病预防控制", "城管执法", "酒店管理", "水电气", "广告监督", "园林绿化环卫", "房地产开发管理", "违章建筑", "道路、桥梁、隔离栏的维护金融管理",
		"非法用工", "占道停车", "大气污染", "消防安全", "垃圾收运、垃圾站的设置及管理物价监管", "就业和再就业", "工作时间和休息休假", "食品安全", "食品安全"}
	Community := []string{"三瓦窑社区", "临江社区", "交子公园社区", "南新社区", "吉泰社区", "和平社区", "大源社区", "天华社区", "昆华社区", "月牙湖社区", "永安社区", "益州社区", "科创社区"}
	Department := []string{"市场监管所", "劳动就业和社会保障服务办公室", "城市管理办公室", "综合执法中队", "综合办公室", "民生办公室", "民政服务办公室", "火车南站综合管理办公室", "党群办公室", "政务服务办公室", "社区发展办公室", "退役军人服务站", "营商环境建设办公室", "社区治理办公室"}
	var id int
	if !IsContain(Classification, order["classification"]) {
		log.Println("请选择分类:")
		for i, j := range Classification {
			fmt.Print(i, ":", j)
		}
		fmt.Scanln(&id)
		order["classification"] = Classification[id]
	}

	log.Printf("请选择科室（1）/社区（2）:")
	fmt.Scanln(&id)

	if id == 1 {
		order["tag"] = "department"
		log.Println("请选择科室:")
		for i, j := range Department {
			fmt.Print(i, ":", j)
		}
		fmt.Scanln(&id)
		order["department"] = Department[id]
	} else if id == 2 {
		order["tag"] = "community"

		log.Println("请选择社区:")
		for i, j := range Community {
			fmt.Print(i, ":", j)
		}
		fmt.Scanln(&id)
		order["community"] = Community[id]
	} else {
		log.Fatal("输入错误")
	}

	return order

}

// Get not singned list
func GetList(cookie string, url string) ResponseContext {
	c := colly.NewCollector()
	var result ResponseContext
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Cookie", cookie)
		r.Headers.Add("X-Requested-With", "XMLHttpRequest")
		r.Headers.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.131 Safari/537.36")
	})

	c.OnResponse(func(r *colly.Response) {
		log.Println("Get Unsign List Successful!")
		if json.Unmarshal([]byte(string(r.Body)), &result) != nil {
			log.Fatal("Json Decode Fail!")
		} else {
			log.Printf("Json Decode Successful!")
		}
	})

	data := map[string]string{
		"page": "1",
		"rows": "100",
	}
	c.Post(url, data)

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("访问失败!")
	})

	return result
}

func GetOrderDetails(wfid int, cookie string, url string) map[string]interface{} {
	result := make(map[string]interface{})
	c := colly.NewCollector()
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Cookie", cookie)
		r.Headers.Add("X-Requested-With", "XMLHttpRequest")
		r.Headers.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.131 Safari/537.36")
	})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Printf(err.Error())
		log.Println("访问失败!")
	})

	c.OnHTML("textarea", func(e *colly.HTMLElement) {
		result["comment"] = e.Text
	})

	c.OnHTML(`input[id="twfEty.workFormNo"]`, func(e *colly.HTMLElement) {
		result["order"] = e.Attr("value")
	})

	c.OnHTML(`input[id="twfEty.fromTelA"]`, func(e *colly.HTMLElement) {
		result["phone"] = e.Attr("value")
	})

	c.OnHTML(`input[id="twfEty.fromName"]`, func(e *colly.HTMLElement) {
		result["name"] = e.Attr("value")
	})

	c.OnHTML(`input[id="twfEty.dcntAname"]`, func(e *colly.HTMLElement) {
		result["classification"] = e.Attr("value")
	})

	Url := url + strconv.Itoa(wfid) + "&act=9"
	c.Visit(Url)
	fmt.Print(result)
	return result
}
