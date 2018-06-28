## 包含基础功能 
1. 授权应用管理 
2. 权限管理 
3. 数据字典
4. 自动测试 
5. 异常处理
6. 日志记录
7. 配置文件灵活指定
8. 数据库升级change log功能
9. docker服务化


## 所使用技术
1. beego
2. mysql
3. glide
4. goconvey
5. hmac
7. casbin
8. jwt
9. sqlite3
10. hateoas
11. redis
12. gorm
13. liquibase



## 新项目
1. go get github.com/beego/bee
2. go get github.com/Masterminds/glide
3. 添加glide 镜像
4. glide up

## 坑
1. orm不支持自定义主键值  推荐用gorm
2. 数据库升级套件不成熟
3. 类型转换

## 总结
1. 多用struct
2. 善用json


## 测试
go test -v
## 或者
1. go get github.com/smartystreets/goconvey
2. cd tests
3. $GOPATH/bin/goconvey

## 压缩 和交叉编译
go get github.com/mitchellh/gox
brew install upx
