package nine

/*
海量数据处理
	面对得两个问题:
		- 一台机器的内存存不下
		- 一台机器处理起来太慢

	海量数据处理的核心(分治):
		- 单机: 利用外存,分批加入内存处理
		- 多机: 对数据分片, 利用多机内存存储
		- 多机: 并行计算,利用多线程,多机并行处理

	一些处理技巧:
		- 外部排序: 多路归并, 桶排序
		- 哈希分片
		- 位图

	常见的问题:
		海量数据的排序
		海量数据的查询
		海量数据的 TOP_K
		海量数据求频率 TOP_K
		海量数据 去重/找重
		两个(多个)海量数据文件找重

解决此类问题的方法,假设在内存足够的情况下进行治理计算,找到合适的数据结果及方法解决问题后,在内存不充足的情况下分治解决问题

·有一个IP地址白名单文件，包含10亿个IP地址，判断某IP是否在白名单中？
·10亿个整数，判断某个整数是否在其中？
·10亿个整数，放在文件中，内存有限，如何求最大的TOP100个整数？
·100GB的搜索关键字文件，统计出现频率TOP100关键词
·一个文件中包含10亿条URL，有可能会重复，将重复的去掉
·给你a、b两个文件，各自有50亿条URL，每条URL占用64字节，内存限制是4GB，找出a、b文件共同的URL
*/
/*
·按照金额大小给10GB的订单文件进行排序

加入当前外存(硬盘)存储了这 10GB 数据, 内存大小为 1.2GB,
*/
