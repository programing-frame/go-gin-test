# gin-test
# 接口中间测试
## 请求URL
http://localhost:20000/feature
## 请求说明

1.HTTP 方法：POST

2.请求头

| 字段         |  类型   |  是否必填项 | 说明            |
| -------      | ------ | --------- | --------------  |
| Content-Type | string | 是         | application/json|
| app_key      | string | 是         | 应用app_key      |
| access_token | string | 是         | access_token     |
| expire_time  | string | 是         | 超时时间           |

3.请求体

| 字段           | 类型   | 是否必填项 | 说明        |
| -------        | ------ | ---------- | ----------- |
| head.serviceId | string | 是         | 请求的服务ID  |
| body           | json   | 是         | 请求推理的数据 |

请求样例
```json
{
    "head": {
        "serviceId": "9999.serving-proxy.chenjinxin.cn#01234567-1234-5678-9abc-0123456789ab"
    },
    "body": {
        "featureData": {
            "x0": 1.12345,
            "x1": -2.123456,
            "x2": 3.567890,
            "x3": 4.12345,
            "x4": 5.567890
        },
        "sendToRemoteFeatureData": {
            "age": 18
        }
    }
}
```
## 返回说明

| 字段         | 类型   | 是否必填项 | 说明                   |
| ------------ | ------ | ---------- | ---------------------- |
| retcode      | int    | 是         | 接口code，用于问题定位 |
| errmsg       | string | 是         | 接口错误信息返回       |
| data         | json   | 是         | 响应结果              |
| flag         | int    | 是         | 标记              |

响应样例
```json
{
    "retcode": 0,
    "retmsg": "OK",
    "data": {
        "score": 0.618,
    },
    "flag": 0
}
```

# 代码运行测试
```
go run main.go
```
```
curl -v localhost:8080/feature -H 'Content-Type: application/json' -d '{}'
```
# 镜像测试
```
docker build -t klb/adapter:v1 .
docker run -d --name adapter -p 20000:8080 klb/adapter:v1
docker logs -f adapter
```
```
curl -v localhost:20000/feature -H 'Content-Type: application/json' -d '{}'
```