# FreeGithub
自动获取github相关网站的ip地址，解决github链接不畅通的问题，该项目的README会自动更新，大家可以直接复制使用

## hosts列表
```base
# Host Start
3.238.201.248 github.com
52.87.249.11 api.github.com
54.174.113.122 github.map.fastly.net
3.238.201.248 github.githubassets.com
18.209.7.231 github.global.ssl.fastly.net
20.225.176.131 raw.githubusercontent.com
54.89.114.171 camo.githubusercontent.com
54.211.84.250 avatars.githubusercontent.com
44.192.18.240 avatars0.githubusercontent.com
54.209.124.21 avatars1.githubusercontent.com
3.227.252.164 avatars2.githubusercontent.com
3.94.95.77 avatars3.githubusercontent.com
52.87.249.11 avatars4.githubusercontent.com
54.166.173.206 avatars5.githubusercontent.com
3.80.231.36 avatars6.githubusercontent.com
107.20.22.157 avatars7.githubusercontent.com
18.209.7.231 avatars8.githubusercontent.com
50.19.174.164 favicons.githubusercontent.com
54.82.34.129 developer.apple.com
44.200.134.165 devstreaming-cdn.apple.com
# Host End
```

更新时间：2022-05-13 09:21:49

## 修改本机的hosts文件
### 存放位置
* Windows 系统：C:\Windows\System32\drivers\etc\hosts
* Linux 系统：/etc/hosts
* Mac（苹果电脑）系统：/etc/hosts

### 生效
* Windows：在 CMD 窗口输入：ipconfig /flushdns
* Linux 命令：sudo rcnscd restart
* Mac 命令：sudo killall -HUP mDNSResponder