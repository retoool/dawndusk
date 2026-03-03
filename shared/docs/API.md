# API 文档

## 基础信息

- **Base URL**: `http://localhost:8080/api/v1`
- **认证方式**: JWT Bearer Token
- **Content-Type**: `application/json`

## 认证端点

### 1. 用户注册

**POST** `/auth/register`

注册新用户账号。

**请求体**:
```json
{
  "username": "string (3-50字符)",
  "email": "string (有效邮箱)",
  "password": "string (最少6字符)"
}
```

**响应** (201 Created):
```json
{
  "access_token": "eyJhbGciOiJIUzI1NiIs...",
  "refresh_token": "eyJhbGciOiJIUzI1NiIs...",
  "user": {
    "id": "uuid",
    "username": "testuser",
    "email": "test@example.com",
    "avatar_url": null,
    "timezone": "UTC",
    "is_verified": false
  }
}
```

**错误响应**:
- `400 Bad Request`: 验证失败或邮箱/用户名已存在
- `500 Internal Server Error`: 服务器错误

### 2. 用户登录

**POST** `/auth/login`

使用邮箱和密码登录。

**请求体**:
```json
{
  "email": "test@example.com",
  "password": "password123"
}
```

**响应** (200 OK):
```json
{
  "access_token": "eyJhbGciOiJIUzI1NiIs...",
  "refresh_token": "eyJhbGciOiJIUzI1NiIs...",
  "user": {
    "id": "uuid",
    "username": "testuser",
    "email": "test@example.com",
    "avatar_url": null,
    "timezone": "UTC",
    "is_verified": false
  }
}
```

**错误响应**:
- `401 Unauthorized`: 邮箱或密码错误
- `403 Forbidden`: 账号已被禁用

### 3. 刷新令牌

**POST** `/auth/refresh`

使用 refresh token 获取新的 access token。

**请求体**:
```json
{
  "refresh_token": "eyJhbGciOiJIUzI1NiIs..."
}
```

**响应** (200 OK):
```json
{
  "access_token": "eyJhbGciOiJIUzI1NiIs...",
  "refresh_token": "eyJhbGciOiJIUzI1NiIs...",
  "user": {
    "id": "uuid",
    "username": "testuser",
    "email": "test@example.com",
    "avatar_url": null,
    "timezone": "UTC",
    "is_verified": false
  }
}
```

**错误响应**:
- `401 Unauthorized`: Token 无效或已过期

### 4. 登出

**POST** `/auth/logout`

登出当前用户（客户端应删除本地 token）。

**Headers**:
```
Authorization: Bearer <access_token>
```

**响应** (200 OK):
```json
{
  "message": "Logged out successfully"
}
```

## 用户端点

### 5. 获取当前用户信息

**GET** `/users/me`

获取当前登录用户的详细信息。

**Headers**:
```
Authorization: Bearer <access_token>
```

**响应** (200 OK):
```json
{
  "message": "Get current user",
  "user_id": "uuid"
}
```

**错误响应**:
- `401 Unauthorized`: Token 无效或未提供

## 打卡端点（待实现）

### 6. 创建打卡记录

**POST** `/check-ins`

创建新的打卡记录。

**Headers**:
```
Authorization: Bearer <access_token>
```

**请求体**:
```json
{
  "type": "wake | sleep",
  "scheduled_time": "2024-03-03T07:00:00Z",
  "actual_time": "2024-03-03T07:05:00Z",
  "mood": "happy",
  "note": "感觉很好",
  "location_lat": 39.9042,
  "location_lng": 116.4074
}
```

**响应** (201 Created):
```json
{
  "id": "uuid",
  "user_id": "uuid",
  "type": "wake",
  "scheduled_time": "2024-03-03T07:00:00Z",
  "actual_time": "2024-03-03T07:05:00Z",
  "time_difference": 5,
  "mood": "happy",
  "note": "感觉很好",
  "created_at": "2024-03-03T07:05:00Z"
}
```

### 7. 获取打卡历史

**GET** `/check-ins?limit=20&offset=0`

获取用户的打卡历史记录。

**Headers**:
```
Authorization: Bearer <access_token>
```

**查询参数**:
- `limit`: 每页数量（默认 20）
- `offset`: 偏移量（默认 0）

**响应** (200 OK):
```json
{
  "check_ins": [
    {
      "id": "uuid",
      "user_id": "uuid",
      "type": "wake",
      "scheduled_time": "2024-03-03T07:00:00Z",
      "actual_time": "2024-03-03T07:05:00Z",
      "time_difference": 5,
      "mood": "happy",
      "created_at": "2024-03-03T07:05:00Z"
    }
  ],
  "total": 100,
  "limit": 20,
  "offset": 0
}
```

### 8. 获取今日打卡

**GET** `/check-ins/today`

获取今天的打卡记录。

**Headers**:
```
Authorization: Bearer <access_token>
```

**响应** (200 OK):
```json
{
  "check_ins": [
    {
      "id": "uuid",
      "type": "wake",
      "actual_time": "2024-03-03T07:05:00Z",
      "time_difference": 5
    }
  ]
}
```

### 9. 获取打卡统计

**GET** `/check-ins/stats`

获取用户的打卡统计数据。

**Headers**:
```
Authorization: Bearer <access_token>
```

**响应** (200 OK):
```json
{
  "total_check_ins": 150,
  "wake_check_ins": 75,
  "sleep_check_ins": 75,
  "current_streak": 7,
  "longest_streak": 30,
  "average_time_diff": 3.5,
  "on_time_percentage": 85.5
}
```

## 宠物端点（待实现）

### 10. 获取宠物信息

**GET** `/pet`

获取用户的宠物信息。

**Headers**:
```
Authorization: Bearer <access_token>
```

**响应** (200 OK):
```json
{
  "id": "uuid",
  "user_id": "uuid",
  "name": "小白",
  "type": "cat",
  "level": 5,
  "experience": 1250,
  "health": 100,
  "happiness": 95,
  "exp_for_next_level": 300,
  "created_at": "2024-01-01T00:00:00Z",
  "updated_at": "2024-03-03T07:05:00Z"
}
```

**错误响应**:
- `404 Not Found`: 用户还没有宠物

### 11. 创建宠物

**POST** `/pet`

为用户创建宠物（每个用户只能有一只）。

**Headers**:
```
Authorization: Bearer <access_token>
```

**请求体**:
```json
{
  "name": "小白",
  "type": "cat"
}
```

**响应** (201 Created):
```json
{
  "id": "uuid",
  "user_id": "uuid",
  "name": "小白",
  "type": "cat",
  "level": 1,
  "experience": 0,
  "health": 100,
  "happiness": 100,
  "exp_for_next_level": 100,
  "created_at": "2024-03-03T07:05:00Z",
  "updated_at": "2024-03-03T07:05:00Z"
}
```

**错误响应**:
- `400 Bad Request`: 用户已有宠物

### 12. 更新宠物信息

**PUT** `/pet`

更新宠物名称。

**Headers**:
```
Authorization: Bearer <access_token>
```

**请求体**:
```json
{
  "name": "小黑"
}
```

**响应** (200 OK):
```json
{
  "id": "uuid",
  "name": "小黑",
  "type": "cat",
  "level": 5,
  "experience": 1250
}
```

### 13. 获取装饰品目录

**GET** `/pet/decorations`

获取所有可用的宠物装饰品。

**Headers**:
```
Authorization: Bearer <access_token>
```

**响应** (200 OK):
```json
{
  "decorations": [
    {
      "id": "uuid",
      "name": "红色帽子",
      "description": "可爱的红色帽子",
      "category": "hat",
      "image_url": "https://...",
      "unlock_level": 5,
      "unlock_check_ins": 50,
      "rarity": "common"
    }
  ]
}
```

### 14. 获取已拥有的装饰品

**GET** `/pet/decorations/owned`

获取用户已解锁的装饰品。

**Headers**:
```
Authorization: Bearer <access_token>
```

**响应** (200 OK):
```json
{
  "decorations": [
    {
      "id": "uuid",
      "decoration_id": "uuid",
      "name": "红色帽子",
      "category": "hat",
      "is_equipped": true,
      "unlocked_at": "2024-02-01T00:00:00Z"
    }
  ]
}
```

### 15. 装备装饰品

**POST** `/pet/decorations/:id/equip`

装备指定的装饰品。

**Headers**:
```
Authorization: Bearer <access_token>
```

**响应** (200 OK):
```json
{
  "message": "Decoration equipped successfully"
}
```

## 错误响应格式

所有错误响应遵循统一格式：

```json
{
  "error": "错误描述",
  "code": "ERROR_CODE"
}
```

### 常见错误码

- `UNAUTHORIZED`: 未授权
- `FORBIDDEN`: 禁止访问
- `NOT_FOUND`: 资源不存在
- `BAD_REQUEST`: 请求参数错误
- `INTERNAL_SERVER`: 服务器内部错误
- `EMAIL_EXISTS`: 邮箱已存在
- `USERNAME_EXISTS`: 用户名已存在
- `INVALID_CREDENTIALS`: 凭证无效
- `USER_NOT_FOUND`: 用户不存在

## 认证流程

1. **注册/登录**: 获取 `access_token` 和 `refresh_token`
2. **API 调用**: 在 Header 中携带 `Authorization: Bearer <access_token>`
3. **Token 过期**: 使用 `refresh_token` 调用 `/auth/refresh` 获取新 token
4. **登出**: 调用 `/auth/logout` 并删除本地存储的 token

## Token 有效期

- **Access Token**: 15 分钟
- **Refresh Token**: 7 天

## 速率限制

- 每个 IP 每分钟最多 100 个请求
- 认证端点每个 IP 每分钟最多 10 个请求

## 测试示例

### 使用 curl

```bash
# 注册
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser","email":"test@example.com","password":"password123"}'

# 登录
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","password":"password123"}'

# 获取用户信息
curl -X GET http://localhost:8080/api/v1/users/me \
  -H "Authorization: Bearer YOUR_ACCESS_TOKEN"
```

### 使用 JavaScript (Fetch)

```javascript
// 注册
const register = async () => {
  const response = await fetch('http://localhost:8080/api/v1/auth/register', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({
      username: 'testuser',
      email: 'test@example.com',
      password: 'password123'
    })
  });
  const data = await response.json();
  return data;
};

// 登录
const login = async () => {
  const response = await fetch('http://localhost:8080/api/v1/auth/login', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({
      email: 'test@example.com',
      password: 'password123'
    })
  });
  const data = await response.json();
  localStorage.setItem('access_token', data.access_token);
  localStorage.setItem('refresh_token', data.refresh_token);
  return data;
};

// 获取用户信息
const getMe = async () => {
  const token = localStorage.getItem('access_token');
  const response = await fetch('http://localhost:8080/api/v1/users/me', {
    headers: { 'Authorization': `Bearer ${token}` }
  });
  return await response.json();
};
```
