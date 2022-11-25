# FreeGithub
自动获取github相关网站的ip地址，解决github链接不畅通的问题，该项目的README会自动更新，大家可以直接复制使用

## hosts列表
```base
# Host Start
140.82.114.4 github.com
140.82.114.4 api.github.com
151.101.1.6 github.map.fastly.net
3.81.9.22 github.githubassets.com
151.101.1.6 github.global.ssl.fastly.net
54.234.188.16 raw.githubusercontent.com
18.207.213.2 camo.githubusercontent.com
3.235.74.207 avatars.githubusercontent.com
54.227.20.153 avatars0.githubusercontent.com
3.95.217.11 avatars1.githubusercontent.com
3.238.149.41 avatars2.githubusercontent.com
184.73.146.99 avatars3.githubusercontent.com
34.231.20.255 avatars4.githubusercontent.com
3.89.253.88 avatars5.githubusercontent.com
34.231.229.220 avatars6.githubusercontent.com
3.81.109.1 avatars7.githubusercontent.com
3.237.170.112 avatars8.githubusercontent.com
44.200.166.0 favicons.githubusercontent.com
17.253.144.10 developer.apple.com
17.253.144.10 devstreaming-cdn.apple.com
# Host End
```

更新时间：2022-11-25 09:08:35

## 修改本机的hosts文件
### 存放位置
* Windows 系统：C:\Windows\System32\drivers\etc\hosts
* Linux 系统：/etc/hosts
* Mac（苹果电脑）系统：/etc/hosts

### 生效
* Windows：在 CMD 窗口输入：ipconfig /flushdns
* Linux 命令：sudo rcnscd restart
* Mac 命令：sudo killall -HUP mDNSResponder