# 自习室管理系统 API 接口文档

## 基础信息

- **Base URL**: `http://localhost:8081`
- **认证方式**: JWT Bearer Token
- **Content-Type**: `application/json`

## 认证说明

部分接口需要在请求头中携带 JWT Token：

```
Authorization: Bearer <token>
```

---

## 1. 验证码模块

### 1.1 获取验证码

**接口描述**: 发送邮箱验证码

**请求方式**: `GET`

**请求路径**: `/api/captcha`

**请求参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| email | string | 是 | 用户邮箱 |

**请求示例**:
```
GET /api/captcha?email=user@example.com
```

**响应示例**:
```json
{
  "code": 1,
  "data": "已发送验证码",
  "msg": "success"
}
```

---

### 1.2 验证验证码

**接口描述**: 验证邮箱验证码是否正确

**请求方式**: `GET`

**请求路径**: `/api/captcha/verify`

**请求参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| email | string | 是 | 用户邮箱 |
| ucode | string | 是 | 用户输入的验证码 |

**请求示例**:
```
GET /api/captcha/verify?email=user@example.com&ucode=123456
```

**响应示例**:
```json
{
  "code": 1,
  "data": "验证成功",
  "msg": "success"
}
```

---

## 2. 用户模块

### 2.1 用户注册

**接口描述**: 用户注册接口

**请求方式**: `POST`

**请求路径**: `/api/user/register`

**请求参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| name | string | 是 | 用户名 |
| email | string | 是 | 邮箱 |
| password | string | 是 | 密码 |
| code | string | 是 | 邮箱验证码 |

**请求示例**:
```json
{
  "name": "张三",
  "email": "zhangsan@example.com",
  "password": "123456",
  "code": "123456"
}
```

**响应示例**:
```json
{
  "code": 1,
  "data": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "msg": "success"
}
```

**说明**: 注册成功后返回 JWT Token

---

### 2.2 用户登录

**接口描述**: 用户登录接口

**请求方式**: `POST`

**请求路径**: `/api/user/login`

**请求参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| email | string | 是 | 邮箱 |
| password | string | 是 | 密码 |

**请求示例**:
```json
{
  "email": "zhangsan@example.com",
  "password": "123456"
}
```

**响应示例**:
```json
{
  "code": 1,
  "data": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "msg": "success"
}
```

**说明**: 登录成功后返回 JWT Token

---

### 2.3 用户注销

**接口描述**: 注销用户账号（软删除）

**请求方式**: `GET`

**请求路径**: `/api/user/logoff`

**请求参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| user_id | int | 是 | 用户ID |

**请求示例**:
```
GET /api/user/logoff?user_id=1
```

**响应示例**:
```json
{
  "code": 1,
  "data": null,
  "msg": "success"
}
```

---

### 2.4 查看用户信息

**接口描述**: 查看用户详细信息

**请求方式**: `GET`

**请求路径**: `/api/user/checkInfo`

**请求参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| email | string | 是 | 用户邮箱 |

**请求示例**:
```
GET /api/user/checkInfo?email=zhangsan@example.com
```

**响应示例**:
```json
{
  "code": 1,
  "data": {
    "name": "张三",
    "email": "zhangsan@example.com",
    "avatar": "/static/avatar/1_1234567890.jpg",
    "gender": 1,
    "major": "计算机科学与技术",
    "class": "2021级1班"
  },
  "msg": "success"
}
```

---

### 2.5 设置用户信息

**接口描述**: 更新用户个人信息

**请求方式**: `POST`

**请求路径**: `/api/user/setInfo`

**请求参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| user_id | uint64 | 是 | 用户ID |
| stu_id | string | 否 | 学号 |
| gender | uint | 否 | 性别 (0:未知, 1:男, 2:女) |
| major | string | 否 | 专业 |
| class | string | 否 | 班级 |

**请求示例**:
```json
{
  "user_id": 1,
  "stu_id": "2021001",
  "gender": 1,
  "major": "计算机科学与技术",
  "class": "2021级1班"
}
```

**响应示例**:
```json
{
  "code": 1,
  "data": "修改成功",
  "msg": "success"
}
```

---

### 2.6 设置用户头像

**接口描述**: 上传并设置用户头像

**请求方式**: `POST`

**请求路径**: `/api/user/setAvatar`

**认证**: 需要 Bearer Token

**请求参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| user_id | int | 是 | 用户ID (form-data) |
| file | file | 是 | 头像文件 (仅支持 jpg/png) |

**请求示例**:
```
POST /api/user/setAvatar
Content-Type: multipart/form-data

user_id=1
file=<binary>
```

**响应示例**:
```json
{
  "code": 1,
  "data": {
    "url": "/static/avatar/1_1234567890.jpg"
  },
  "msg": "success"
}
```

---

### 2.7 忘记密码

**接口描述**: 通过邮箱验证码重置密码

**请求方式**: `POST`

**请求路径**: `/api/user/forgetPwd`

**请求参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| email | string | 是 | 邮箱 |
| password | string | 是 | 原密码 |
| rePassword | string | 是 | 新密码 |
| code | string | 是 | 邮箱验证码 |

**请求示例**:
```json
{
  "email": "zhangsan@example.com",
  "password": "123456",
  "rePassword": "654321",
  "code": "123456"
}
```

**响应示例**:
```json
{
  "code": 1,
  "data": "修改密码成功",
  "msg": "success"
}
```

---

## 3. 自习室模块

> **注意**: 以下所有接口均需要 JWT 认证

### 3.1 查看自习室状态

**接口描述**: 根据自习室ID查看自习室详细信息

**请求方式**: `GET`

**请求路径**: `/api/room/show`

**认证**: 需要 Bearer Token

**请求参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| room_id | int | 是 | 自习室ID |

**请求示例**:
```
GET /api/room/show?room_id=1
Authorization: Bearer <token>
```

**响应示例**:
```json
{
  "code": 1,
  "data": {
    "room_id": 1,
    "room_name": "A101自习室",
    "capacity": 50,
    "status": 1
  },
  "msg": "success"
}
```

---

### 3.2 查看空闲自习室

**接口描述**: 获取所有空闲状态的自习室列表

**请求方式**: `GET`

**请求路径**: `/api/room/showIdle`

**认证**: 需要 Bearer Token

**请求参数**: 无

**请求示例**:
```
GET /api/room/showIdle
Authorization: Bearer <token>
```

**响应示例**:
```json
{
  "code": 0,
  "data": [
    {
      "room_id": 1,
      "room_name": "A101自习室",
      "capacity": 50
    },
    {
      "room_id": 2,
      "room_name": "A102自习室",
      "capacity": 40
    }
  ],
  "msg": "success"
}
```

---

### 3.3 预约自习室

**接口描述**: 预约指定的自习室

**请求方式**: `POST`

**请求路径**: `/api/room/appointment`

**认证**: 需要 Bearer Token

**请求参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| room_id | uint64 | 是 | 自习室ID |
| user_id | uint64 | 是 | 用户ID |

**请求示例**:
```json
{
  "room_id": 1,
  "user_id": 1
}
```

**响应示例**:
```json
{
  "code": 1,
  "data": "预约成功",
  "msg": "success"
}
```

---

### 3.4 查看预约记录

**接口描述**: 查看用户的自习室预约记录

**请求方式**: `GET`

**请求路径**: `/api/room/showAppointment`

**认证**: 需要 Bearer Token

**请求参数**:

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| user_id | int | 是 | 用户ID |

**请求示例**:
```
GET /api/room/showAppointment?user_id=1
Authorization: Bearer <token>
```

**响应示例**:
```json
{
  "code": 1,
  "data": [
    {
      "appointment_id": 1,
      "room_id": 1,
      "room_name": "A101自习室",
      "user_id": 1,
      "appointment_time": "2025-09-20 14:30:00"
    }
  ],
  "msg": "success"
}
```

---

## 4. 响应状态码说明

| Code | 说明 |
|------|------|
| 0 | 查询成功（无数据或空列表） |
| 1 | 操作成功 |
| 0 | 操作失败 |

## 5. 错误响应示例

```json
{
  "code": 0,
  "data": null,
  "msg": "用户不存在"
}
```

## 6. 注意事项

1. 所有时间格式均为 `YYYY-MM-DD HH:mm:ss`
2. 验证码有效期为 5 分钟
3. JWT Token 有效期为 90 分钟
4. 头像文件大小限制请参考服务器配置
5. 自习室模块的所有接口都需要先通过登录获取 Token
6. 用户注销为软删除，不会真正删除数据

## 7. Swagger 文档

项目启动后可访问 Swagger 在线文档：

```
http://localhost:8081/swagger/index.html
```
