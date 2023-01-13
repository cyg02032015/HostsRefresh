# FreeGithub
自动获取github相关网站的ip地址，解决github链接不畅通的问题，该项目的README会自动更新，大家可以直接复制使用

## hosts列表
```base
# Host Start
140.82.114.3 github.com
140.82.114.3 api.github.com
151.101.1.6 github.map.fastly.net
34.237.75.88 github.githubassets.com
151.101.1.6 github.global.ssl.fastly.net
52.91.68.214 raw.githubusercontent.com
3.95.219.176 camo.githubusercontent.com
100.24.118.154 avatars.githubusercontent.com
52.90.117.131 avatars0.githubusercontent.com
100.24.118.154 avatars1.githubusercontent.com
54.242.232.154 avatars2.githubusercontent.com
3.81.165.75 avatars3.githubusercontent.com
3.215.16.180 avatars4.githubusercontent.com
54.173.168.40 avatars5.githubusercontent.com
3.215.79.191 avatars6.githubusercontent.com
54.89.147.202 avatars7.githubusercontent.com
44.200.40.7 avatars8.githubusercontent.com
54.210.39.84 favicons.githubusercontent.com
17.253.144.10 developer.apple.com
17.253.144.10 devstreaming-cdn.apple.com
# Host End
```

更新时间：2023-01-13 09:05:21

## 修改本机的hosts文件
### 存放位置
* Windows 系统：C:\Windows\System32\drivers\etc\hosts
* Linux 系统：/etc/hosts
* Mac（苹果电脑）系统：/etc/hosts

### 生效
* Windows：在 CMD 窗口输入：ipconfig /flushdns
* Linux 命令：sudo rcnscd restart
* Mac 命令：sudo killall -HUP mDNSResponder