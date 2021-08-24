# flybeat 数据收集处理

data := "a|1|2|d|f|123214"
Match:{"Split":{"seq":"|", "fileds":["index","category","date","msg","id"]}}

data := "{\"a\":123}"
Match:{"Json":{}}

Match:{"Regex":{"match":"", fileds:["index","category","date","msg","id"]}}

data := "a=1,c=2"
Match:{"Kv":{"seq": "=", "delimiter": ","}}
Input:{
	"Redis":{
		"address":"127.0.0.1:6379",
		"topic":"queue2",
		"password":"",
		"db":0,
		"codec":"plain"
	}
}
Input:{
	"Elasticsearch":{
		"address":["http://127.0.0.1:9200"],
		"index":"web%20060102",
		"index_type":"logs"
	}
}

Filter:{"Add":[{"filed":"field", "process":"", "value":"Now(2006-01-02)"}]}
Filter:{"Add":[{"filed": "field", "process": "", "value": "Random(20,10)"}]}
Filter:{"Remove":[{"filed": "field", "process": "HasSuffix(234)", "value": ""}]}
Filter:{"Remove":[{"filed": "field", "process": "HasPrefix(d2)", "value": ""}]}

Filter:{"Drop":[{"filed":"model", "process":"hasPrefix(q)", "value":""}]}
Filter:{"Drop":[{"filed":"model", "process":"hasSuffix(q)", "value":""}]}
Filter:{"Replace":[{"filed":"model", "process":"hasSuffix(q)", "value":""}]}
Filter:{"Modify":[{"filed":"model", "process":"hasSuffix(model, "q")", "value":""}]}

