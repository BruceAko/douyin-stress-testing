# 基于locust + boomer的抖音自动化压力测试

## 使用方法

1.启动locust master

```shell
locust --master -f dummy.py
```

2.另起一终端，运行压力测试程序stress_testing

```shell
./stress_testing --task $(taskName)
```

taskName包括十六个接口所对应的测试：

feed
register
login
user
publish_action
publish_list
like_action
like_list
comment_action
comment_list
relation_action
follow_list
follower_list
friend_list
message_action
chat

3.打开浏览器 <http://127.0.0.1:8089>，即可在locust中监控各项指标，并导出报告

重点关注Requests per second(吞吐量RPS)、Failed requests、90%，95%和99%的响应时间

## 说明

将编译好的程序放在服务器上运行时可能出现GLIBC版本号对不上的情况，因此需要交叉编译：

```bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build --ldflags="-s -w" -o stress_testing
```

## 压测结果

请见《青训营后端结业项目答辩汇报文档》的性能测试部分

[压测结果](https://gzd0wrb2k4.feishu.cn/docx/TlUxdrUiOoT9E8xiN4Ocwhhpnig#doxcnzYQ7ltoNzjIeWnA9uWBrrc)