# 一个索引工具
### 根据目录层级找到最大的N个目录

## 参数列表
```
sparks-MacBook-Pro:longtech-tools spark$ ./longtech-tools -h
Usage of ./longtech-tools:
  -d int
        搜索深度, 默认为最深层级搜索
  -m int
        输出最大目录条目, 默认为10 (default 10)
  -p string
        目录路径, 默认为当前路径 (default "./")
```

## linux 用法

```
# 下载执行文件
$wget https://github.com/threen134/longtech-tools/releases/download/v/longtech-tools
# 添加权限
$chmod +x ./longtech-tools
# 使用
$./longtech-tools -d 3 -p ~/workspace -m 2
输出第1级最大2的目录列表
/Users/spark/workspace/restClient, 832
/Users/spark/workspace/dns-ui, 864
输出第2级最大2的目录列表
/Users/spark/workspace/java/grpc-java, 1920
/Users/spark/workspace/doc/PrivateDNS-CP.wiki, 5536
输出第3级最大2的目录列表
/Users/spark/workspace/restClient/.git/objects, 3968
/Users/spark/workspace/node/app-id-sample-js/node_modules, 13920
```