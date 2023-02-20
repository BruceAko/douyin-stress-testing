# 基于locust + boomer的抖音自动化压力测试

## 使用方法

1.启动locust master

```shell
locust --master -f dummy.py
```

2.另起一终端，运行压力测试程序

```shell
go run stress_testing.go $(taskName)
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

3.打开浏览器 <http://127.0.0.1:8089>，即可在locust中监控各项指标
重点关注Requests per second(吞吐量RPS)、Failed requests、90%，95%和99%的响应时间

## 压测结果

### feed

### register

### login

### user

### publish_action

### publish_list

### like_action

### like_list

### comment_action

### comment_list

### relation_action

### follow_list

### follower_list

### friend_list

### message_action

### chat
