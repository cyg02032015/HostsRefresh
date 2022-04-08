# FreeGithub
自动获取github相关网站的ip地址，解决github链接不畅通的问题，该项目的README会自动更新，大家可以直接复制使用

## hosts列表
```base
# Host Start
100.26.47.4 github.com
44.202.219.102 api.github.com
100.26.47.4 github.map.fastly.net
3.81.107.37 github.githubassets.com
54.162.10.190 github.global.ssl.fastly.net
100.26.47.4 raw.githubusercontent.com
3.227.253.24 camo.githubusercontent.com
44.201.96.222 avatars.githubusercontent.com
54.144.230.153 avatars0.githubusercontent.com
54.89.124.183 avatars1.githubusercontent.com
3.239.234.34 avatars2.githubusercontent.com
44.202.236.3 avatars3.githubusercontent.com
18.204.2.46 avatars4.githubusercontent.com
3.91.178.105 avatars5.githubusercontent.com
3.238.73.122 avatars6.githubusercontent.com
107.20.99.207 avatars7.githubusercontent.com
3.237.92.127 avatars8.githubusercontent.com
3.238.73.122 favicons.githubusercontent.com
44.200.166.148 developer.apple.com
3.231.216.227 devstreaming-cdn.apple.com
# Host End
```

更新时间：2022-04-08 09:01:55

## 修改本机的hosts文件
### 存放位置
* Windows 系统：C:\Windows\System32\drivers\etc\hosts
* Linux 系统：/etc/hosts
* Mac（苹果电脑）系统：/etc/hosts

### 生效
* Windows：在 CMD 窗口输入：ipconfig /flushdns
* Linux 命令：sudo rcnscd restart
* Mac 命令：sudo killall -HUP mDNSResponder