package __内置函数

/*
	Go 语言拥有一些不需要进行导入操作就可以使用的内置函数。
它们有时可以针对不同的类型进行操作，例如：len、cap 和 append，
或必须用于系统级的操作，例如：panic。因此，它们需要直接获得编译器的支持。

close	用于管道通信
len		用于返回某个类型的长度或数量（字符串、数组、切片、map 和管道）
cap		是容量的意思，用于返回某个类型的最大容量（只能用于切片和 map
copy、append	用于复制和连接切片
panic、recover	两者均用于错误处理机制
print、println	底层打印函数(在部署环境中建议使用 fmt 包)
complex、real imag	用于创建和操作复数
*/
