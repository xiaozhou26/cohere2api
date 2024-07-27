# cohere2api
Docker部署

```bash
docker run -d \
  --name aurora \
  -p 8080:8080 \
  ghcr.io/xiaozhou26/cohere2api:latest
```
## Usage
```bash

curl --request POST \
  --url http://127.0.0.1:8080/v1/chat/completions \
  --header 'Authorization: Bearer cohere key' \
  --data '{
  "messages": [
    {
      "role": "user",
      "content": "test"
    }
  ],
  "model": "command-r-plus",
}'

```
## 支持的模型
command-r-plus	
command-r	
command	
command-nightly
command-light	
command-light-nightly
## cohere key获取
https://dashboard.cohere.com/api-keys
