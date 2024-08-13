# 项目介绍

### 项目路径规划
* app——微服务代码
* controller
* log——zap
* middleware  中间件
* service  服务接口
* dao 数据库调用接口
* common  公共模块
* configs  配置文件
* docs  文档相关
* cmd  命令行
* Dockerfile 容器构建脚本
* views  html模板
* model  数据结构

### go版本
go version go1.17.5 windows/amd64

### 代码调试,会自动下载依赖包
* go run ./main.go

### 打包线上生产代码(vs 变量可能不生效，可以换cmd执行)
set GOOS=linux 
go build -tags netgo main.go

### 打包生产组件
tar -zcvf go_iris_web.tar.gz main config* views* Dockerfile

### 将文件夹都放在conf文件夹下，Copy命令不会把文件夹本身复制到容器中
mv config* views*  conf/

### 需要安装curl，不然健康检查报错