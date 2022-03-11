# FreeGithub
自动获取github相关网站的ip地址，解决github链接不畅通的问题，该项目的README会自动更新，大家可以直接复制使用

## hosts列表
```base
# Host Start
54.173.67.40 github.com
18.208.126.153 api.github.com
54.152.216.171 github.map.fastly.net
40.70.206.76 github.githubassets.com
107.22.8.82 github.global.ssl.fastly.net
3.94.29.26 raw.githubusercontent.com
107.21.83.79 camo.githubusercontent.com
18.232.111.166 avatars.githubusercontent.com
54.88.61.17 avatars0.githubusercontent.com
3.238.174.87 avatars1.githubusercontent.com
107.21.83.79 avatars2.githubusercontent.com
54.88.61.17 avatars3.githubusercontent.com
3.216.23.224 avatars4.githubusercontent.com
54.205.161.48 avatars5.githubusercontent.com
54.152.82.231 avatars6.githubusercontent.com
44.198.179.200 avatars7.githubusercontent.com
3.238.174.87 avatars8.githubusercontent.com
54.166.212.121 favicons.githubusercontent.com
54.205.161.48 developer.apple.com
34.204.201.127 devstreaming-cdn.apple.com
# Host End
```

更新时间：2022-03-11 08:59:09

## 修改本机的hosts文件
### 存放位置
* Windows 系统：C:\Windows\System32\drivers\etc\hosts
* Linux 系统：/etc/hosts
* Mac（苹果电脑）系统：/etc/hosts

### 生效
* Windows：在 CMD 窗口输入：ipconfig /flushdns
* Linux 命令：sudo rcnscd restart
* Mac 命令：sudo killall -HUP mDNSResponder