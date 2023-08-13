git：

仓库相关：
git remote add git@github.com:WEIxucheng1024/test.git                   将远程仓库添加到本地
git remote -v                                                           显示所有绑定的远程仓库
git remote rm origin                                                    删除别名为origin的远程仓库链接
git init                                                                创建仓库

分支代码相关
git branch                                                              展示所有分支
git branch name                                                         创建某个分支
git checkout -b name                                                    创建并跳转到新分支(如果没有-b，那么就只有跳转)
git branch -d branchname                                                删除分支
git merge branchname                                                    合并分支
git add *                                                               上传所有修改记录到本地
git commit -m ''                                                        提交本地修改到远程仓库
git push origin master                                                  将本地的分支版本上传到远程并合并
git diff                                                                比较文件的不同，即暂存区和工作区的差异
git reset                                                               回退版本
git rm                                                                  将文件从暂存区和工作区中删除。
git mv                                                                  移动或重命名工作区文件。

go mod:
go mod init XXX                             	初始化项目(XXX为模块名称)
go get github.com/XXX/XXX                   	获取某个包
go env -w GO111MODULES=ON                   	开启go mod模块
go mod edit -replace=旧的模块信息=新的模块信息  	更换依赖包的版本
go mod tidy				  	引用项目需要的依赖增加到go.mod文件。
去掉go.mod文件中项目不需要的依赖



Linux:
ssh ones@101.33.207.188					远程连接服务器
nohup ./recv > recv.log &					后台运行recv程序，并将标准输出打印到recv.log下
tail -f recv.log						查看某个标准输出
上传文件到远程服务器：
scp /Users/weixucheng/Desktop/辣鸡？？？/main ones@101.33.207.188:/tmp
/Users/weixucheng/Desktop/辣鸡？？？/main为本地文件路径和名字
chmod +x main						赋予main可执行权限


Redis:
redis-server						启动redis服务
redis-cli -h 101.33.207.188 -p 6379 -a "password"		启动redis客户端(后面为链接远程redis数据库)
select 1							切换到1号数据库
dbsize							查看当前数据库的存储数量
EXPIRE key 10						给某个key设置十秒过期时间

Redis的CRUD：这里只记录了部分指令，详细的可以在菜鸟教程上看

string：
set key value						写入数据(有对应的key则直接修改)
mset key1 value1 key2 value2				一次保存多个key、value
get key value						拿数据
Mget key1 key2						一次获取多个数据
del key							删除某个数据
flushdb							清空当前数据库(flushall位清空全部)

哈希：
hset user1 name 'Tom'
hset user1 age 21					添加hash键值对
hmset user1 name lisi age 21				一次性添加多个内容
hget user1 name
hget user1 age						获取hash内容
hmget user1 name age					一次性获取多个内容
hgetall key						获取key下的所有字段和字段内容
hvals key 						获取key下的所有字段内容
hkeys key						获取key下的所有字段
hlen key							获取key下的字段数量
hexists key field					查看这个key下是否有field字段，有返回1，没有返回0

list：
Lpush key value1 value2...				按顺序添加元素到list内(从左边开始)
Rpush key value1 value2...				从右边按顺序插入
lrange key 0 -1						获取list下的所有数据(从第0个开始，-1代表最后一个
元素，-2代表倒数第二个元素)
lpop key							从list最左边拿出一个数据(拿出的数据在redis移除)
rpop							从list最右边移出一个数据
del heroslist						移出一个list，也就是删除
llen key							返回list的长度(为空时返回0)
lindex key indix						获取list指定下标的内容

set：
sadd key value1 value2					添加set
smember key						提取set中所有的元素
Sismember key value					判断集合中是否有这个元素，有的话返回1，无则0
srem key value						删除集合中的某个元素，删除成功返回1，失败则为0


Docker
docker ps -a						        展示所有容器
docker start AAAAA					        启动ID为AAAA的容器
docker exec -it AAAA /bin/sh				进入AAAA的容器




MySQL

Mysql -uroot -p123456
使用root，密码123456登录mysql

CREATE USER 'pig'@'%' IDENTIFIED BY '123456';		
创建成员命令(@后面为%标识任意远程主机都可登录，只允许本机的话%改为localhost)

GRANT SELECT, INSERT ON test.user TO 'pig'@'%';
GRANT ALL ON *.* TO 'pig'@'%';
授权，SELECT、INSERT等位权限，test.user为test库下的User表，第二句意思是授权所有库和表

GRANT privileges ON databasename.tablename TO 'username'@'host' WITH GRANT OPTION;
允许该客户给别的客户授权

SET PASSWORD FOR 'username'@'host' = PASSWORD('newpassword');
修改用户的密码
SET PASSWORD = PASSWORD("newpassword");
修改当前登录用户的密码

REVOKE privilege ON databasename.tablename FROM 'username'@'host';
撤销用户授权

DROP USER 'username'@'host';
删除用户

SHOW DATABASES;
展示所有数据库

use MySQL;
进入名字为mysql的数据库

SHOW TABLES;
展示数据库下的所有表




go get -u github.com/asaskevich/govalidator
用于验证数据的格式和有效性的包，比如验证邮箱

go get -u github.com/swaggo/swag/cmd/swag
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files
swagger包，用来在页面上展示目前项目下的接口等，service层的接口方法上面会有大量的@ XXX

go get github.com/spf13/viper
这个包是用来便捷引用一些yml配置文件的包

go get -u github.com/gin-gonic/gin
获取gin框架

go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite
go get gorm.io/driver/mysql
这个是导入gorm，包括gorm连接mysql的包等
















