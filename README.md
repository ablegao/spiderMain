RUN

	git clone https://github.com/ablegao/spiderMain.git
	cd spiderMain
	go get . 
	go build . 
	spiderMain --config=./config.yaml



Config:

	## http 请求的头补充信息
    header:
      User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.119 Safari/537.36
     
    ## 顺序任务
    workflows:
      -
        run: get  # 执行http get 模式抓取网页
        value: https://www.hao123.com
        out-type: html
      -
        run: stdout  # 测试输出
      -
        run: html  # 将上一步得到的数据按HTML 解析
        find: # jquery 语法查找对象
          - a
        attr: href ## 可选参数 ， 获取a .href 如果指定attr ，将返回一个 href 的列表
		#out-type: html ## 当没有指定attr 时， 该值默认为 text , 将返回 <a> 标签的innerHTML 集合 , 如果指定a 标签,则返回所有的a 标签字符串
      -
        run: stdout
	  - 
	    run: each-http # 批量获取上一步指定链接的内容
		
	  #### 更多参数
	  - 
	    run: each-download: # 批量下载上一步得到的所有链接 上一步获取到的数据， 必须是以\n分割的多个路径
		out-path: ./download-dir
	  - 
	    run: write-to-file
		out-path: ./downloads/a.log
        
