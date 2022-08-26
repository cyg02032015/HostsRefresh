# FreeGithub
自动获取github相关网站的ip地址，解决github链接不畅通的问题，该项目的README会自动更新，大家可以直接复制使用

## hosts列表
```base
# Host Start
140.82.113.3 github.com
140.82.113.3 api.github.com
151.101.1.6 github.map.fastly.net
3.238.144.136 github.githubassets.com
151.101.1.6 github.global.ssl.fastly.net
3.239.5.157 raw.githubusercontent.com
44.197.101.72 camo.githubusercontent.com
3.81.141.61 avatars.githubusercontent.com
44.203.237.124 avatars0.githubusercontent.com
44.197.175.237 avatars1.githubusercontent.com
3.234.227.138 avatars2.githubusercontent.com
3.88.104.60 avatars3.githubusercontent.com
54.221.154.238 avatars4.githubusercontent.com
3.228.220.94 avatars5.githubusercontent.com
23.23.46.76 avatars6.githubusercontent.com
34.207.178.24 avatars7.githubusercontent.com
44.200.45.217 avatars8.githubusercontent.com
54.92.198.0 favicons.githubusercontent.com
17.253.144.10 developer.apple.com
17.253.144.10 devstreaming-cdn.apple.com
# Host End
```

更新时间：2022-08-26 09:41:05

## 修改本机的hosts文件
### 存放位置
* Windows 系统：C:\Windows\System32\drivers\etc\hosts
* Linux 系统：/etc/hosts
* Mac（苹果电脑）系统：/etc/hosts

### 生效
* Windows：在 CMD 窗口输入：ipconfig /flushdns
* Linux 命令：sudo rcnscd restart
* Mac 命令：sudo killall -HUP mDNSResponder