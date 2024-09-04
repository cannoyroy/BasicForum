# 项目介绍

在这个项目中要求开发一个简易的论坛系统，分为学生端和管理端两个部分，基本实现发帖和审核的功能。

技术栈：Vue3+Axios；Golang(gin+gorm+viper)；MySQL

项目目标：基本理解前后端的交互逻辑以及网站的开发逻辑

前提申明：未使用相关组件，所以页面简陋。

### 功能

- [x] 前端辅助功能

点击图标跳转首页

返回按钮

- [x] 用户登录

用户名+密码

- [x] 用户注册

用户名+真实姓名+密码+身份

若失败，弹窗告知理由

若用户名/姓名超过长度，则阻止输入

- [x] 学生发帖

帖子内容（后台同时发送发帖人信息），字数统计/警告，内容不作截断处理

- [x] 学生获取所有发布的帖子

换行展示

- [x] 修改帖子
- [x] 删除帖子
- [x] 举报帖子
- [x] 查看举报审批
- [x] 管理员获取未审批举报忒子
- [x] 管理员审核被举报帖子

### 项目结构

```
server // 后端
├── internal
│   ├── global
│   │   └── config.go
│   ├── midwares
│   │   └── midwares.go
│   ├── models
│   │   └── models.go
│   ├── pkg\database
│   │   └── mysql.go
│   ├── router
│   │   └── router.go
│   └── services
│       │── common.go
│       └── services.go
├── conf
│   └── config
├── pkg
│   └── utils
├── main.go 		//项目入口文件
│
│
forumjh // 前端
├── node_modules // 第三方库和模块
│   └── ……
│
├── src
│   ├── assets  // 静态资源
│   │   └── ……
│   │
│   ├── styles
│   │   └── header.css
│   │
│   ├── views  # 页面组件
│   │   ├── LoginView.vue
│   │   ├── RegView.vue
│   │   ├── ROOT
│   │   │   ├──StudentBoard.vue
│   │   │   └──TeacherBoard.vue
│   │   ├── Student
│   │   │   ├──StudentDelePost.vue
│   │   │   └──StudentGetPost.vue
│   │   │   ├──StudentGetReport.vue
│   │   │   └──StudentPost.vue
│   │   │   ├──StudentPutPost.vue
│   │   │   └──StudentReportPost.vue
│   │   └── Teacher
│   │       ├──AdminGetReport.vue
│   │       └──AdminProReport.vue
│   │
│   ├── router
│   │   └── index.js
│   │
│   ├── App.vue
│   └── main.js
│
├── .gitignore
├── index.html
├── jsconfig.json
├── package-lock.json
├── package.json
└── vite.config.js
```

### ~~项目缺陷~~-TO DO LIST

- 前端-VUE报错`Unable to resolve `@import "./src/style/header.css"` from D:/BaiduSyncdisk/项目汇总/soft/2024年暑假招新大作业/forumjh/src/views`
- ~~前端-发帖长度限制~~
- 前端-界面UI 极限情况
- ~~前端-安全性问题~~
- ~~前端-注销按钮~~
- ~~前端-router/index.js 二级路由~~
- ~~后端-目录~~框架【努力修复中，但感觉无力回天，应该前期就预定好固定的项目规范文档】
- ~~后端-main.go viper/数据库 单独调用，不要放到这里~~
- ~~后端-user-id 主键自增？~~
- ~~后端-举报记录保留~~
- ~~后端-安全，是否是管理员~~
- ~~后端-举报已被举报通过的帖子仍会成功，并且使该帖子状态变回0~~

### 其他备注

在“审核被举报的帖子”板块，认为`user_id`参数用处不大，所以后端就没对其进行使用。因为审核帖子只针对帖子本身内容，而与发帖的人和举报的人无关。且“(使用者本人)”的表述也带有奇异，后期若修改此处时，务必确定好其定义。

运行时需要添加后端的数据库配置文件，并且进行相关配置安装：

node安装；`npm install`；`npm install axios`

**SQL表创建脚本**

```sql
CREATE TABLE IF NOT EXISTS users (
  user_id INT PRIMARY KEY,
  username VARCHAR(50) NOT NULL UNIQUE,
  name VARCHAR(100) NOT NULL,
  password VARCHAR(255) NOT NULL CHECK (LENGTH(password) > 8 AND LENGTH(password) < 16),
  user_type INT NOT NULL CHECK (user_type IN (1, 2)),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

```sql
CREATE TABLE posts (
    post_id INT PRIMARY KEY,
    user_id INT,
    username VARCHAR(50) NOT NULL,
    name VARCHAR(100) NOT NULL,
    user_type INT NOT NULL CHECK (user_type IN (1, 2)),
    content VARCHAR(255),
    reason VARCHAR(255),
    state INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

```sql
CREATE TABLE report (
    report_id INT PRIMARY KEY,
    post_id INT,
    report_user_id INT,
    reason VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

```sql
CREATE TABLE trash (
    report_id INT PRIMARY KEY,
    post_id INT,
    report_user_id INT,
    reason VARCHAR(255),
    state INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```



GitHub忽略了上传的 .gitignore 文件

# 项目笔记

## 后端

### 环境配置

```go
import (
	"fmt" // 1
	"log" // 2
	"net/http" // 3
	"time" // 4

	"github.com/gin-gonic/gin" // 5
	"github.com/spf13/viper" // 6
	"gorm.io/driver/mysql" // 7
	"gorm.io/gorm" // 8
)
```

**`fmt`**

   - **用途**: `fmt` 是 Go 的标准库之一，主要用于格式化输入和输出。
   - **作用**: 在项目中，通常用于打印日志信息、格式化字符串等操作。

 **`log`**

   - **用途**: `log` 也是 Go 的标准库，用于记录日志信息。
   - **作用**: 你可以使用 `log` 包来记录重要的运行时信息，如错误日志、调试信息等，帮助你在开发和维护中追踪程序的行为。

 **`net/http`**

   - **用途**: `net/http` 是 Go 的标准库，提供了 HTTP 客户端和服务器的实现。
   - **作用**: 它为开发 Web 应用程序提供了核心功能，例如处理 HTTP 请求和响应。在使用 `gin` 框架时，`net/http` 包被用来处理底层的 HTTP 请求。

 **`time`**

   - **用途**: `time` 是 Go 的标准库，提供了时间的处理功能。
   - **作用**: `time` 包用于获取当前时间、格式化时间、处理时间差等操作。在 Web 应用中，常用于记录事件发生的时间、计算请求的处理时间等。

 **`github.com/gin-gonic/gin`**

   - **用途**: `gin` 是一个流行的 Go 语言 Web 框架，具有高性能和简洁的 API。
   - **作用**: `gin` 用于快速构建 Web 应用程序或 API。它提供了路由、处理请求、生成响应等常见的 Web 开发功能。

 **`github.com/spf13/viper`**

   - **用途**: `viper` 是一个强大的 Go 配置管理包。
   - **作用**: `viper` 用于读取配置文件、环境变量等，帮助管理应用程序的配置。它支持 JSON、TOML、YAML、HCL 和 envfile 等多种格式的配置文件。

**`gorm.io/driver/mysql`**

   - **用途**: `gorm.io/driver/mysql` 是 `Gorm` ORM（对象关系映射）框架的 MySQL 驱动。
   - **作用**: 通过这个驱动，你可以使用 `Gorm` 与 MySQL 数据库进行交互，如执行 SQL 查询、迁移数据库结构等。

 **`gorm.io/gorm`**

   - **用途**: `gorm` 是 Go 语言的一个流行的 ORM 库，简化了数据库操作。
   - **作用**: 它允许开发者通过结构体来操作数据库表，如增删改查记录、自动迁移数据库表结构等，大大简化了与数据库交互的代码。

### 读取配置文件

```go
viper.SetConfigName("config")
viper.AddConfigPath(".")
viper.SetConfigType("yaml")
if err := viper.ReadInConfig(); err != nil {
	log.Fatalf("Error reading config file, %s", err)
}
```

1. `viper.SetConfigName("config")`：
   - 这一行设置了配置文件的名称（不包括扩展名）。在这个例子中，配置文件的名称是 `config`，所以 Viper 将查找名为 `config` 的文件。
2. `viper.AddConfigPath(".")`：
   - 这一行指定了配置文件所在的路径。`.` 表示当前工作目录。Viper 将在当前工作目录中查找配置文件。
3. `viper.SetConfigType("yaml")`：
   - 这一行设置了配置文件的类型。在这个例子中，配置文件是 YAML 格式的。
4. `if err := viper.ReadInConfig(); err != nil {`：
   - 这一行尝试读取配置文件。`ReadInConfig` 是一个阻塞操作，它会搜索所有添加的路径，并按顺序读取找到的第一个配置文件。如果找到了配置文件并成功读取，`err` 将为 `nil`；如果没有找到配置文件或读取时发生错误，`err` 将包含错误信息。
5. `log.Fatalf("Error reading config file, %s", err)`：
   - 如果 `ReadInConfig` 返回了一个非 `nil` 的错误，这行代码将使用 `log.Fatalf` 打印错误消息并退出程序。`%s` 是一个格式占位符，用于插入 `err` 变量的值。

### 数据库连接

```go
dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	viper.GetString("database.user"),
	viper.GetString("database.password"),
	viper.GetString("database.host"),
	viper.GetString("database.port"),
	viper.GetString("database.name"),
)
db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
if err != nil {
	log.Fatalf("Error connecting to database: %s", err)
}
```

1. `dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",`：
   - `fmt.Sprintf` 是 Go 中的一个函数，用于格式化字符串。这里它用于构造一个数据源名称（DSN），这是连接到数据库所需的信息字符串。
   - `%s` 是一个格式占位符，用于在字符串中插入变量的值。
   - `viper.GetString` 是 Viper 库中用于获取配置文件中指定键的字符串值的函数。在这个例子中，它用于从配置文件中获取数据库连接信息。
   - `database.user`、`database.password`、`database.host`、`database.port` 和 `database.name` 是配置文件中的键，它们分别对应数据库的用户名、密码、主机地址、端口号和数据库名称。
   DSN 的格式通常如下：
   ```
   [username]:[password]@tcp([host]:[port])/[database]?[parameters]
   ```
   在这个例子中，DSN 包含以下参数：
   - `charset=utf8mb4`：指定数据库的字符集为 `utf8mb4`，这是为了支持存储 Unicode 字符。
   - `parseTime=True`：告诉 GORM 在处理数据库中的时间字段时，应该将它们解析为 Go 的 `time.Time` 类型。
   - `loc=Local`：设置本地时区，这样 GORM 就会使用本地时区来解析时间。
2. `db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})`：
   
   - `gorm.Open` 是 GORM 库中用于打开数据库连接的函数。它接受两个参数：数据库驱动程序的连接器和 GORM 的配置。
   - `mysql.Open` 是 GORM 提供的 MySQL 驱动程序中的函数，用于创建一个数据库连接器实例。它接受一个 DSN 字符串作为参数。
   - `&gorm.Config{}` 是一个 GORM 配置实例的指针。在这个例子中，它是一个空配置，意味着将使用 GORM 的默认设置。你可以根据需要修改这个配置。
3. `if err != nil { log.Fatalf("Error connecting to database: %s", err) }`：
   - 这行代码检查 `gorm.Open` 函数是否返回了错误。如果有错误，`err` 将不为 `nil`。
   - `log.Fatalf` 是 Go 标准库 `log` 中的一个函数，用于记录一条消息并随后调用 `os.Exit(1)` 来退出程序。这里，它记录了一个错误消息，并包含了错误详情。

### 迁移框架

```go
db.AutoMigrate(&users{})
```

首先，想象一下你正在玩一个建设城市的游戏。在这个游戏中，你有一个蓝图（Blueprint），它描述了你想要建造的建筑的所有细节，比如建筑的大小、形状、颜色等。你的任务是使用这个蓝图来建造实际的城市建筑。
在这个比喻中，`User` 结构体就像是你的蓝图，而数据库中的表就像是你要建造的建筑。`AutoMigrate` 方法就是你的建筑团队，它负责根据蓝图来建造或更新建筑。
下面是详细的步骤：
**1. 蓝图（User 结构体）**
在 Go 语言中，`User` 结构体定义了你希望在数据库中创建的表的布局。它看起来可能像这样：

```go
type User struct {
    ID        uint   `gorm:"primaryKey"`
    Username  string `gorm:"unique"`
    Password  string
    Email     string `gorm:"unique"`
}
```
这个结构体定义了四个字段：`ID`、`Username`、`Password` 和 `Email`。每个字段旁边的 `gorm:"..."` 是一个标签，它提供了额外的信息，告诉 GORM 如何处理这个字段。例如，`gorm:"primaryKey"` 表示 `ID` 字段是主键，而 `gorm:"unique"` 表示 `Username` 和 `Email` 字段应该是唯一的。

**2. 建筑团队（AutoMigrate 方法）**

现在，我们有了蓝图，我们需要一个团队来建造实际的建筑。在 GORM 中，`AutoMigrate` 方法就是这个团队。
当调用 `db.AutoMigrate(&User{})` 时，以下是发生的事情：
- **检查建筑是否存在**：`AutoMigrate` 首先检查数据库中是否已经有一个名为 `users` 的表（这是根据 `User` 结构体默认推断出的表名）。
- **建造新建筑**：如果表不存在，`AutoMigrate` 将创建一个新的表，并根据 `User` 结构体的定义设置所有的列和属性。例如，它会创建一个 `ID` 列作为主键，并且自增，还会为 `Username` 和 `Email` 设置唯一约束。
- **更新现有建筑**：如果表已经存在，`AutoMigrate` 将比较现有表的结构和 `User` 结构体的定义。如果发现差异（比如添加了新的字段或修改了现有字段的类型），它将更新表结构以匹配蓝图。

**3. 开始建造（执行 AutoMigrate）**

当你执行 `db.AutoMigrate(&User{})`，GORM 就开始了建造过程：
- 它读取 `User` 结构体的定义。
- 它在数据库中查找对应的表。
- 它根据需要创建或更新表。
这个过程是自动的，这也是为什么它被称为 `AutoMigrate`。它简化了数据库迁移的过程，让你不需要手动编写 SQL 语句来创建或修改表结构。

**总结**

所以，`db.AutoMigrate(&User{})` 是一个告诉 GORM 根据你的 Go 结构体定义来创建或更新数据库表的方法。这就像是你给了建筑团队一个蓝图，然后团队就按照蓝图来建造或更新你的城市建筑。简单、高效，而且减少了出错的可能性。

### POST登录请求

```go
r.POST("/api/user/login", func(c *gin.Context) {})
```

1. `r.POST`：
   - `r` 是一个 `gin.Engine` 类型的实例，它是在调用 `gin.Default()` 时创建的。
   - `POST` 是一个方法，用于指定该路由将处理 HTTP POST 请求。
2. `"/api/user/login"`：
   - 这是一个字符串，表示路由的路径。在这个例子中，它是一个相对路径，表示当客户端向服务器发送一个 POST 请求到 `/api/user/login` 路径时，该路由将匹配这个请求。
3. `func(c *gin.Context) {}`：
   - 这是一个匿名函数，它是路由的处理函数。每当匹配到对应的路由时，这个函数就会被调用。
   - `c *gin.Context` 是该函数的参数，`c` 是一个 `gin.Context` 类型的实例，它包含了请求的所有信息，包括请求的头部（Headers）、路径参数（Path Parameters）、查询参数（Query Parameters）、表单值（Form Values）等，并且提供了响应客户端的方法。

`gin.Context` 是一个非常重要的结构体，它被用来传递请求过程中的一些重要信息，并提供响应请求的方法。以下是一些你可以使用 `gin.Context` 来执行的操作：
- 获取请求参数：如查询参数、表单值、路径参数等。
- 设置响应状态码：如 `c.Status(200)`。
- 写入响应头：如 `c.Header("Content-Type", "application/json")`。
- 写入响应体：如 `c.String(200, "Hello, World!")` 或 `c.JSON(200, gin.H{"message": "Hello, World!"})`。
- 中断中间件链：如 `c.Abort()`。
- 获取中间件链中的上下文数据：如 `c.Get("key")`。

```go
c.ShouldBindJSON(&req)
```

- **赋值操作**：`ShouldBindJSON` 是 `gin.Context` 类型的一个方法，用于尝试将 JSON 数据从请求体解码到指定的结构体 `req`。`&req` 是一个指针，它指向一个结构体实例。`ShouldBindJSON` 方法会将 JSON 数据解码到这个结构体实例上。

- **返回操作**：如果 JSON 数据格式不正确或者解码失败，`ShouldBindJSON` 方法将返回一个非 `nil` 的错误值，表示解码过程中出现了问题。

```go
c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "Invalid request"})
```

- `c` 是一个 `gin.Context` 类型的实例，它代表了当前的 HTTP 请求和响应。

- `c.JSON` 是 `gin.Context` 类型的一个方法，用于向客户端发送一个 JSON 响应。
- `http.StatusBadRequest` 是一个常量，代表 HTTP 状态码 400，表示客户端请求无效。
- `gin.H{"code": 400, "msg": "Invalid request"}` 是一个 `gin.H` 类型的实例，它是一个方便的 map 类型，用于构建 JSON 响应。在这个例子中，它包含了一个 `code` 键和一个 `msg` 键，分别代表错误代码和错误消息。

```go
db.Where("username = ?", req.Username).First(&user).Error
```

执行一个查询操作，旨在从数据库中检索与特定条件匹配的记录。

- `db`：这是一个 GORM 数据库连接实例，它代表了与数据库的连接。

- `Where("username = ?", req.Username)`：这是一个查询构造器，用于指定查询条件。`Where` 方法接受一个条件表达式，这里是 `"username = ?"`，表示查询应该匹配 `username` 字段等于某个值。`?` 是一个占位符，用于在执行查询时提供具体的值。
- `First(&user)`：这是 GORM 的查询方法，用于执行查询并返回第一个匹配的记录。`&user` 是一个指针，指向一个 `User` 结构体实例，用于接收查询结果。
- `db.Error` 是 GORM 数据库连接实例的一个属性，它用于返回最后一个操作的错误。如果数据库操作成功，`db.Error` 将返回 `nil`；如果操作失败，它将返回一个错误对象。

### 一些常量

```
http.StatusOK	http.StatusBadRequest	http.StatusInternalServerError
```

- `http.StatusOK` 是 Go 语言标准库 `net/http` 包中的一个常量，它代表 HTTP 协议中的状态码 200。这个状态码是 HTTP 协议中最常见的状态码之一，它表示请求成功并且服务器已经理解了请求，并且返回了客户端请求的资源。
- `http.StatusBadRequest` 是 Go 语言标准库 `net/http` 包中的一个常量，它代表 HTTP 协议中的状态码 400。这个状态码是 HTTP 协议中的一个错误状态码，它表示客户端发送的请求有语法错误，服务器无法理解该请求。
- `http.StatusInternalServerError` 是 Go 语言标准库 `net/http` 包中的一个常量，它代表 HTTP 协议中的状态码 500。这个状态码是 HTTP 协议中的一个错误状态码，它表示服务器遇到了一个无法处理的错误，导致服务器无法完成请求。

```
gorm.ErrRecordNotFound
```

`gorm.ErrRecordNotFound` 是 GORM 框架中的一个错误类型，用于表示查询数据库时没有找到任何记录。当你的查询语句应该返回至少一条记录，但实际上数据库中没有匹配的记录时，GORM 会返回 `gorm.ErrRecordNotFound` 错误。

### GO 和SQL连接时，表和结构体要对应：

在 Go 中，使用 Gorm 连接结构体与数据库表主要通过以下方式完成：

###### 1. **默认映射**

Gorm 会根据结构体的名字自动映射到对应的表名。例如：

```go
type User struct {
    gorm.Model
    Username string `json:"username" gorm:"unique"`
    Name     string `json:"name"`
    Password string `json:"password"`
    UserType int    `json:"user_type"`
}
```

Gorm 默认会将 `User` 结构体映射到名为 `users` 的表。表名是结构体名称的小写复数形式。

###### 2. **自定义表名**

如果你想自定义结构体对应的表名，可以实现 `TableName` 方法：

```go
func (User) TableName() string {
    return "my_users"
}
```

这将使 `User` 结构体映射到名为 `my_users` 的表。

###### 3. **字段映射**

Gorm 会根据结构体字段名映射到表中的列名，默认情况下，字段名对应的列名是字段名的小写形式。例如：

```go
type Post struct {
    ID       uint   `gorm:"primaryKey"` // Gorm 默认会将 ID 映射为主键
    Title    string `json:"title"`
    Content  string `json:"content"`
    AuthorID uint   `json:"author_id"`
}
```

- `Title` 字段会映射到 `title` 列。
- `Content` 字段会映射到 `content` 列。
- `AuthorID` 字段会映射到 `author_id` 列。

###### 4. **自定义列名**

你可以通过 struct tag 自定义列名：

```go
type Post struct {
    ID       uint   `gorm:"primaryKey"`     // 默认映射到 `id` 列
    Title    string `gorm:"column:post_title"` // 映射到 `post_title` 列
    Content  string `gorm:"column:post_content"`
    AuthorID uint   `gorm:"column:author_id"`
}
```

###### 5. **忽略字段**

如果你有某个字段不想映射到数据库表中，可以使用 `gorm:"-"` 忽略：

```go
type User struct {
    gorm.Model
    Username string `json:"username" gorm:"unique"`
    Password string `json:"password"`
    TempData string `gorm:"-"` // TempData 不会映射到数据库表中
}
```

###### 6. **设置主键**

可以使用 `gorm:"primaryKey"` 设置自定义主键：

```go
type User struct {
    MyID    string `gorm:"primaryKey"` // 使用自定义主键
    Username string `json:"username" gorm:"unique"`
}
```

###### 7. **其他常见标签**

- `type`：指定列的数据类型，例如 `gorm:"type:varchar(100)"`。
- `default`：指定列的默认值，例如 `gorm:"default:'unknown'"`。
- `index`：为列创建索引，例如 `gorm:"index"`。

### MYSQL插入表

```go
if err := db.Create(&newUser).Error; err != nil {
	c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "Internal server error"})
	return
}
```

1. `db.Create(&newUser)`：
   - `db` 是一个 `gorm.DB` 类型的实例，代表与数据库的连接。
   - `Create` 是 `gorm.DB` 类型提供的一个方法，用于执行创建记录的操作。
   - `&newUser` 是一个指针，指向一个新的 `users` 结构体实例，其中包含了新用户的属性。
2. `if err != nil {`：
   - 这行代码开始了一个 `if` 语句，用于检查执行 `Create` 操作后返回的错误 `err` 是否为 `nil`。
   - 如果 `err` 是 `nil`，意味着数据库操作成功执行；如果 `err` 不是 `nil`，则表示数据库操作失败。
3. `c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "Internal server error"})`：
   - 这行代码是 `if` 语句的错误处理分支。
   - `c` 是一个 `gin.Context` 类型的实例，代表当前的 HTTP 请求和响应。
   - `c.JSON` 是 `gin.Context` 类型提供的一个方法，用于向客户端发送一个 JSON 响应。
   - `http.StatusInternalServerError` 是 Go 语言标准库中的一个常量，代表 HTTP 状态码 500，表示服务器内部错误。
   - `gin.H{"code": 500, "msg": "Internal server error"}` 是一个 `gin.H` 类型的实例，它是一个方便的 map 类型，用于构建 JSON 响应。在这个例子中，它包含了一个 `code` 键和一个 `msg` 键，分别代表错误代码和错误消息。
4. `return`：
   - 这行代码是 `if` 语句的错误处理分支的结束。
   - 当 `if` 语句的条件为 `true` 时（即数据库操作失败），`return` 语句将终止当前的函数调用，并且不会继续执行函数体中的后续代码。

### MYSQL创表

在MySQL中创建新表，以存储账号基本信息

```sql
CREATE TABLE IF NOT EXISTS users (
  user_id INT PRIMARY KEY AUTO_INCREMENT,
  username VARCHAR(50) NOT NULL UNIQUE,
  name VARCHAR(100) NOT NULL,
  password VARCHAR(255) NOT NULL CHECK (LENGTH(password) > 8 AND LENGTH(password) < 16),
  user_type INT NOT NULL CHECK (user_type IN (1, 2)),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

- `CREATE TABLE IF NOT EXISTS users`：如果名为 `users` 的表不存在，则创建一个名为 `users` 的新表。
- `(`：表定义的开始括号。
- `user_id INT PRIMARY KEY AUTO_INCREMENT`：定义一个名为 `user_id` 的列，数据类型为 `INT`（整数），作为主键，并且每次插入新记录时自动增加。
- `username VARCHAR(50) NOT NULL UNIQUE`：定义一个名为 `username` 的列，数据类型为 `VARCHAR(50)`（最多 50 个字符的字符串），不能为空，并且必须是唯一的，即每个用户名只能有一个。
- `name VARCHAR(100) NOT NULL`：定义一个名为 `name` 的列，数据类型为 `VARCHAR(100)`，不能为空。
- `password VARCHAR(255) NOT NULL CHECK (LENGTH(password) >= 8 AND LENGTH(password) <= 16)`：定义一个名为 `password` 的列，数据类型为 `VARCHAR(255)`，不能为空，并且有一个 `CHECK` 约束，确保密码长度在 8 到 16 个字符之间。
- `user_type INT NOT NULL CHECK (user_type IN (1, 2))`：定义一个名为 `user_type` 的列，数据类型为 `INT`，不能为空，并且有一个 `CHECK` 约束，确保用户类型只能是 1 或 2。
- `created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP`：定义一个名为 `created_at` 的列，数据类型为 `TIMESTAMP`，默认值为当前时间戳。
- `)`：表定义的结束括号。

```sql
CREATE TABLE posts (
    post_id INT PRIMARY KEY,
    user_id INT,
    username VARCHAR(50) NOT NULL,
    name VARCHAR(100) NOT NULL,
    user_type INT NOT NULL CHECK (user_type IN (1, 2)),
    content VARCHAR(255),
    reason VARCHAR(255),
    state INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

```sql
CREATE TABLE report (
    report_id INT PRIMARY KEY,
    post_id INT,
    report_user_id INT,
    reason VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### GO语言判断字符串是否由数字构成

```go
package main

import (
	"strconv"
)

func isNumber(str string) bool {
	_, err := strconv.Atoi(str)
	return err == nil
}

func main() {
	if isNumber("12a3") {
		println("yes")
	} else {
		println("no")
	}
}
```

`Atoi`函数将字符串转换为整数类型，如果转换成功，则说明该字符串为数字；如果转换失败，则说明该字符串不是数字。

### session和jwt

#### Session

- **定义**：Session是一种服务器端机制，用于存储用户会话信息。当用户第一次访问应用时，服务器会创建一个会话，并为这个会话分配一个唯一的标识（Session ID）。
- **工作原理**：
  1. 用户登录后，服务器生成一个Session ID。
  2. 服务器将Session ID发送给客户端，通常是通过设置Cookie来实现。
  3. 客户端在随后的请求中携带这个Session ID。
  4. 服务器根据Session ID识别用户，并从服务器存储中检索会话信息。
- **存储位置**：Session信息通常存储在服务器的内存、数据库或文件系统中。

#### JWT (JSON Web Tokens)

- **定义**：JWT是一种紧凑的、自包含的方式，用于在各方之间以JSON对象的形式安全地传输信息。它可以在信息中添加签名（使用HMAC算法或使用RSA/ECDSA公钥/私钥对），以确保信息不被篡改。
- **工作原理**：
  1. 用户登录后，服务器生成一个JWT。
  2. 服务器将JWT发送给客户端。
  3. 客户端在随后的请求中将JWT附加到HTTP请求的Authorization头部中。
  4. 服务器验证JWT的有效性，并从中提取用户信息。
- **结构**：JWT由三部分组成：头部（Header）、载荷（Payload）、签名（Signature）。
#### Session的使用场景
- **需要频繁访问服务器存储**：如果应用需要频繁读写会话数据，使用Session更合适，因为数据存储在服务器端，操作更为灵活。
- **安全性要求较高**：由于Session ID不包含用户信息，而是存储在服务器端，因此相对更安全。
- **易于管理会话状态**：服务器可以随时创建、销毁或更新会话信息，易于实现用户登录状态的监控和管理。
#### JWT的使用场景
- **跨域认证**：JWT支持跨域认证，特别适用于单页应用（SPA）或多端应用（如移动应用）。
- **无状态认证**：由于JWT本身包含了用户信息，服务器无需存储会话状态，适合于分布式系统和微服务架构。
- **性能要求高**：由于不需要每次请求都查询数据库，使用JWT可以减少服务器端的压力，提高性能。
- **RESTful API**：JWT常用于RESTful API，因为它们是无状态的，且JWT可以很容易地通过HTTP头部传输。

### ref() & reactive()

在 Vue 3 中，`reactive()` 和 `ref()` 是 Composition API 提供的两种不同的方式来创建响应式数据。它们的主要区别在于它们可以创建的响应式数据类型以及它们如何处理数据。

** ref()**

`ref()` 函数用于创建一个响应式的引用，它返回一个 `ref` 对象，这个对象包含一个内部值和一个 `.value` 属性。你可以通过 `.value` 属性来读取或修改内部值。`ref()` 主要用于基本数据类型（如字符串、数字、布尔值等），而不适用于对象和数组。

**reactive()****

`reactive()` 函数用于创建一个响应式的代理对象，它返回一个 `reactive` 对象，这个对象是一个响应式代理，可以代理一个普通对象。当你修改代理对象内部的数据时，Vue 能够知道并作出相应的反应。`reactive()` 主要用于对象和数组。

**区别总结**

1. **适用类型**：
   - `ref()` 适用于基本数据类型。
   - `reactive()` 适用于对象和数组。
2. **内部结构**：
   - `ref()` 返回一个包含内部值的 `ref` 对象。
   - `reactive()` 返回一个代理对象，这个对象可以代理一个普通对象。
3. **修改数据**：
   - 使用 `ref()` 时，你需要直接修改 `.value` 属性来修改内部值。
   - 使用 `reactive()` 时，你可以直接修改代理对象内部的数据，Vue 会自动处理响应式更新。
4. **模板使用**：
   - 在模板中，`ref()` 的值需要通过 `.value` 属性来访问。
   - `reactive()` 的值可以直接在模板中使用，不需要额外的访问属性。
5. **性能**：
   - `ref()` 通常比 `reactive()` 更轻量级，因为它不涉及代理对象。
   - `reactive()` 创建了代理对象，这可能会带来一定的性能开销。
   总的来说，`ref()` 和 `reactive()` 都是 Vue 3 提供的响应式数据创建方式，它们各自适用于不同的数据类型，并且具有不同的使用场景和性能特点。在实际开发中，根据你的具体需求选择合适的方式。

### 中间件

在Go语言中，中间件（middleware）是一个在请求处理过程中执行的函数。它可以在请求到达你的实际业务逻辑之前或之后执行一些操作，比如设置CORS头部、日志记录、身份验证、请求处理等。中间件非常有用，因为它可以让你将通用功能模块化，从而使你的代码更加整洁和易于维护。
通俗易懂地解释，中间件就像是你在厨房做菜时的一系列“调味品”，它们可以给你的菜（请求）增添额外的风味（功能）。
在Go的Web框架中，比如Gin，中间件的使用非常简单。你只需要定义一个函数，然后将它添加到框架的路由器中。当请求到达路由器时，中间件函数会按照添加的顺序执行。
以下是一个简单的Gin中间件示例，它用于记录每个请求的时间：
```go
package main
import (
	"github.com/gin-gonic/gin"
	"time"
)
// logger 中间件，用于记录请求的时间
func logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		// 继续执行下一个中间件或路由
		c.Next()
		// 获取处理请求所需的时间
		end := time.Now()
		latency := end.Sub(start)
		// 打印请求日志
		log.Printf("[GIN-DEBUG] %s - [%s] %s %s %d %s %s\n",
			start.Format("2006-01-02 15:04:05.000000"),
			c.ClientIP(),
			c.Request.Method,
			c.Request.URL.Path,
			c.Writer.Status(),
			latency,
			c.Request.Proto,
		)
	}
}
func main() {
	// 创建一个新的Gin实例
	router := gin.Default()
	// 注册logger中间件
	router.Use(logger())
	// 注册一个简单的GET路由
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// 启动服务器
	router.Run(":8080")
}
```
在这个例子中，我们定义了一个名为`logger`的中间件，它记录了每个请求的开始时间、客户端IP、请求方法、请求路径、响应状态码、请求处理所需的时间以及请求的协议。这个中间件被添加到Gin路由器中，所以每当有请求到达时，它都会被执行。
中间件是非常强大的工具，可以让你轻松地添加各种功能到你的Web应用中。在实际开发中，你可能需要创建多个中间件来处理不同的需求。

### Q&A

###### 为什么要用.json存储用户名和密码等相关用户数据?

优点

1. **简单易用**：
   - **格式简单**：JSON 格式结构简单，易于理解和操作。
   - **内建支持**：Python 的 `json` 模块提供了内建的支持，方便读取和写入 JSON 文件。
2. **轻量级**：
   - **无需复杂配置**：相比于使用数据库，JSON 文件不需要额外的数据库管理系统或复杂的配置。
   - **适合小规模应用**：对于小型应用或开发初期，JSON 文件提供了一个轻量级的数据存储方式。
3. **可读性**：
   - **人类可读**：JSON 文件是文本格式，容易查看和编辑。
   - **调试方便**：在开发和调试过程中，直接查看和修改 JSON 文件中的数据可以更方便。
4. **便于传输**：
   - **标准格式**：JSON 是一个广泛使用的标准格式，便于在不同系统和语言间传输数据。

缺点

1. **性能问题**：
   - **性能较低**：对于大型用户数据或高并发请求，读取和写入 JSON 文件的性能可能不如数据库。
   - **锁定问题**：文件系统可能会遇到锁定问题，影响并发访问的性能。
2. **缺乏安全性**：
   - **没有加密**：JSON 文件不提供加密或安全存储功能，用户数据的安全性可能需要额外措施。
   - **易于篡改**：如果文件权限不设置得当，JSON 文件容易被未授权的用户修改。
3. **数据管理**：
   - **不支持复杂查询**：JSON 文件不支持数据库那样复杂的查询和索引功能。
   - **数据一致性**：在高并发环境下，确保数据一致性和完整性可能变得困难。
4. **扩展性**：
   - **难以扩展**：随着用户数据的增加，JSON 文件可能会变得不够高效或难以管理。

###### BUG1：安装好GO后，发现pip命令行居然无法使用了

然后用Pip3就可以了……虽然仍然不知道为什么为什么本来可以用的pip不能用了

## 前端

### VUE框架

### HTML

- `<ul>`：这是一个无序列表（Unordered List）的标签，用于包裹一系列的列表项。在网页中，`<ul>` 通常会以项目符号（如圆点、方块等）的形式展示列表项。
- `<li>`：这是列表项（List Item）的标签，用于定义列表中的具体项目。每个 `<li>` 标签代表列表中的一个条目。
### OPTIONS请求

`OPTIONS` requests are what we call "preflight" requests in Cross-origin resource sharing (CORS).

They are necessary when you're making requests across different origins in specific situations.

This preflight request is made by some browsers as a safety measure to ensure that the request being done is trusted by the server. Meaning the server understands that the method, origin and headers being sent on the request are safe to act upon.

Your server should not ignore but handle these requests whenever you're attempting to do cross origin requests.

A good resource can be found here http://enable-cors.org/

A way to handle these to get comfortable is to ensure that for any path with `OPTIONS` method the server sends a response with this header

```
Access-Control-Allow-Origin: *
```

This will tell the browser that the server is willing to answer requests from any origin.

For more information on how to add CORS support to your server see the following flowchart

![](http://www.html5rocks.com/static/images/cors_server_flowchart.png)

后端：

在Go语言中，如果想要在服务器端设置CORS头部，以确保对于任何路径的OPTIONS方法，服务器都发送一个带有`Access-Control-Allow-Origin: *`头部的响应，你可以使用Gin框架来轻松实现这一点。

```go
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// 创建一个新的Gin实例
	router := gin.Default()

	// 设置CORS头部
	router.Use(cors())

	// 注册OPTIONS请求的处理函数
	router.OPTIONS("api/user/login", func(c *gin.Context) {
		// 你的业务逻辑
	})


	// 启动服务器
	router.Run(":8080")
}

// cors 中间件设置CORS头部
func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 设置允许的源，这里设置为*允许任何源
		c.Header("Access-Control-Allow-Origin", "*")
		// 设置允许的请求方法
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		// 设置允许的请求头
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// 如果请求类型是OPTIONS，则直接返回
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		// 继续执行下一个中间件或路由
		c.Next()
	}
}

```

在这个例子中，我们创建了一个名为`cors`的中间件，它设置了CORS头部，包括`Access-Control-Allow-Origin`、`Access-Control-Allow-Methods`和`Access-Control-Allow-Headers`。这个中间件被应用到所有路由之前，确保了对于任何路径的OPTIONS方法，服务器都会发送一个带有正确CORS头部的响应。

请注意，设置`Access-Control-Allow-Origin`为`*`是允许任何源的跨源请求，这在开发阶段是常见的做法，但在生产环境中，你应该设置为特定的源，以提高安全性。