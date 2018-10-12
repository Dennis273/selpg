# SELPG

## 设计说明

* 本程序的功能是将文档部分页数输出。输出结果可以输出至打印设备或是标准输入输出流。

* 文件结构：

  ```
  .
  ├── main.go (源代码)
  ├── readme.md 
  ├── selpg （二进制文件）
  ├── test1 (测试用文件1)
  └── test2 (测试用文件2)
  ```

  其中test1中含有0-1，每个数字一行，test2中数据为"a\fb\fc\f"

## 使用

* 命令行参数有：
  * 起始页 : -s 打印起始页码
  * 中止页：-e 打印中止页码
  * 页长：-l 源文档每页行数 默认为 72
  * 页尾标识：-f 使用 '\f' 判断换页
  * 目标地址：-d 供外部设备打印输出

* 使用方法

  ```
  /selpg$ ./selpg -s[startpage] -e[endpage] -l[pageLength] -f(option) [sourcefile]
  ```

  其中-l参数与-f参数互斥

## 测试

* 打印输出至终端

  ```bash
  /selpg$ go build
  /selpg$ ./selpg -s1 -e3 -l2 test1
  ```

  输出结果：

  ```
  2
  3
  4
  5
  6
  7
  ```

* 打印输出至目标文件

  ```bash
  /selpg$ go build
  /selpg$ ./selpg -s1 -e2 -l3 test1 > temp
  ```

  temp中数据

  ```
  3
  4
  5
  6
  7
  8
  ```

  ​

