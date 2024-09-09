# Captcha

仅内部使用的验证码服务，不得直接用于前端请求

## Service

- port: 8080 -> 8085  (容器内映射出的端口)

## API

---

### POST `/verify/create`

生成验证码
- form-data
    - expire `int`: 过期时间，单位秒
- return
    - image `string`: base64 image
    - id `string`: captcha id

---

### POST `/verify/check`
验证答案
- form-data
    - id `string`: captcha id
    - answer `string`: captcha answer
- return
    - result `bool`: true/false