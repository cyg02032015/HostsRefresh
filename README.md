# FreeGithub
自动获取github相关网站的ip地址，解决github链接不畅通的问题，该项目的README会自动更新，大家可以直接复制使用

## hosts列表
```base
# Host Start
54.172.18.251 github.com
67.202.26.16 api.github.com
3.83.40.162 github.map.fastly.net
44.192.75.252 github.githubassets.com
3.236.97.131 github.global.ssl.fastly.net
35.153.144.102 raw.githubusercontent.com
54.82.114.5 camo.githubusercontent.com
3.238.161.229 avatars.githubusercontent.com
18.209.16.237 avatars0.githubusercontent.com
3.82.243.73 avatars1.githubusercontent.com
3.85.97.183 avatars2.githubusercontent.com
44.197.103.68 avatars3.githubusercontent.com
54.86.99.56 avatars4.githubusercontent.com
54.166.213.83 avatars5.githubusercontent.com
54.167.205.107 avatars6.githubusercontent.com
44.197.251.145 avatars7.githubusercontent.com
34.229.132.208 avatars8.githubusercontent.com
3.238.201.245 favicons.githubusercontent.com
3.89.112.110 developer.apple.com
18.209.16.237 devstreaming-cdn.apple.com
# Host End
```

更新时间：2021-10-08 08:46:46

## 修改本机的hosts文件
### 存放位置
* Windows 系统：C:\Windows\System32\drivers\etc\hosts
* Linux 系统：/etc/hosts
* Mac（苹果电脑）系统：/etc/hosts

### 生效
* Windows：在 CMD 窗口输入：ipconfig /flushdns
* Linux 命令：sudo rcnscd restart
* Mac 命令：sudo killall -HUP mDNSResponder