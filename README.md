# flybeat 数据收集处理

# 描述
1：现阶段支持api下方对应的策略到该组件

# 后续
1：后续新增流计算
2：新增配置文件新增策略

# input
	# kafka 
	# redis 
	## demo
	{"Redis": {"address": "localhost:6379", "username": "", "password": "", "db": "DB0", "topic": "queue7"}}


# output
    # flydb 自研时间序列数据库(待发布)
	# es
	## demo
	{"Elasticsearch": {"address": "http://127.0.0.1:9200", "username": "", "password": "", "index": "web%20060102", "index_type": "logs"}}

# match
    # {"Split":{"seq":"|", "fileds":["index","category","date","msg","id"]}}
	# {"Regex":{"match":"", fileds:["index","category","date","msg","id"]}}
	# {"Grok":{"match": "%{COMMONAPACHELOG}"}}
	# {"Kv":{"seq": "=", "delimiter": ","}}
	# {"Json":{}}
	 
# filter
	## {"Add":[{"filed":"field", "process":"", "value":"Now(2006-01-02)"}]}
	## {"Add":[{"filed": "field", "process": "", "value": "Random(20,10)"}]}
	## {"Remove":[{"filed": "field", "process": "HasSuffix(234)", "value": ""}]}
	## {"Remove":[{"filed": "field", "process": "HasPrefix(d2)", "value": ""}]}
	## {"Drop":[{"filed":"model", "process":"hasPrefix(q)", "value":""}]}
	## {"Drop":[{"filed":"model", "process":"hasSuffix(q)", "value":""}]}
	## {"Replace":[{"filed":"model", "process":"hasSuffix(q)", "value":""}]}
	##  {"Modify":[{"filed":"model", "process":"hasSuffix(model, "q")", "value":""}]}
 
