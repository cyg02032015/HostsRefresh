# FreeGithub
自动获取github相关网站的ip地址，解决github链接不畅通的问题，该项目的README会自动更新，大家可以直接复制使用

## hosts列表
```base
# Host Start
140.82.114.3 github.com
140.82.114.3 api.github.com
151.101.1.6 github.map.fastly.net
35.92.182.211 github.githubassets.com
151.101.1.6 github.global.ssl.fastly.net
52.10.6.19 raw.githubusercontent.com
52.10.6.19 camo.githubusercontent.com
35.90.113.215 avatars.githubusercontent.com
54.202.255.32 avatars0.githubusercontent.com
54.202.255.32 avatars1.githubusercontent.com
35.162.250.28 avatars2.githubusercontent.com
34.219.78.45 avatars3.githubusercontent.com
35.87.18.50 avatars4.githubusercontent.com
54.201.24.83 avatars5.githubusercontent.com
35.164.1.202 avatars6.githubusercontent.com
54.191.47.205 avatars7.githubusercontent.com
54.244.3.111 avatars8.githubusercontent.com
35.92.51.191 favicons.githubusercontent.com
17.253.144.10 developer.apple.com
17.253.144.10 devstreaming-cdn.apple.com
# Host End
```

更新时间：2022-10-21 09:25:12

## 修改本机的hosts文件
### 存放位置
* Windows 系统：C:\Windows\System32\drivers\etc\hosts
* Linux 系统：/etc/hosts
* Mac（苹果电脑）系统：/etc/hosts

### 生效
* Windows：在 CMD 窗口输入：ipconfig /flushdns
* Linux 命令：sudo rcnscd restart
* Mac 命令：sudo killall -HUP mDNSResponder