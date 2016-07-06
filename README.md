# Socks5Server
一个用golang写的socks5代理服务器，主要代理功能代码沿用了[physacco/socks5](https://github.com/physacco/socks5)的，然后再原来的基础上面增加了Ip白名单和json配置功能。

#安装

    git clone https://github.com/farmercode/Socks5Server.git
    cd Socks5Server
    go build socks5server.go
    ./socks5server
    #如果需要后台执行可以使用下面命令
    ./socks5server > /dev/null 2>&1 &
    #如果需要记录日志文件可以使用nohup
    nohup ./socks5server &
    #或者重定向到自己指定的一个文件
    ./socks5server > server.log 2>&1 &

#配置文件
config.json为配置文件  
 **Listen** 为要代理服务器监听ip和端口，格式为 **host:port** , 如  `127.0.0.01:9527`  
 **WhiteList**为Ip白名单，当WhiteList为空时，所有Ip地址都可以连接，WhiteList不为空时，只允许WhiteList中的Ip进行代理,如 `["127.0.0.1","180.168.85.190"]` 

