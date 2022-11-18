# FreeGithub
自动获取github相关网站的ip地址，解决github链接不畅通的问题，该项目的README会自动更新，大家可以直接复制使用

## hosts列表
```base
# Host Start
140.82.112.4 github.com
140.82.112.4 api.github.com
151.101.1.6 github.map.fastly.net
52.203.208.220 github.githubassets.com
151.101.1.6 github.global.ssl.fastly.net
44.211.26.51 raw.githubusercontent.com
3.91.5.109 camo.githubusercontent.com
3.86.221.5 avatars.githubusercontent.com
3.237.44.11 avatars0.githubusercontent.com
3.236.228.140 avatars1.githubusercontent.com
18.204.219.151 avatars2.githubusercontent.com
54.159.138.95 avatars3.githubusercontent.com
3.236.21.205 avatars4.githubusercontent.com
54.173.170.203 avatars5.githubusercontent.com
54.167.197.43 avatars6.githubusercontent.com
44.210.133.156 avatars7.githubusercontent.com
3.236.122.26 avatars8.githubusercontent.com
44.200.159.175 favicons.githubusercontent.com
17.253.144.10 developer.apple.com
17.253.144.10 devstreaming-cdn.apple.com
# Host End
```

更新时间：2022-11-18 09:19:46

## 修改本机的hosts文件
### 存放位置
* Windows 系统：C:\Windows\System32\drivers\etc\hosts
* Linux 系统：/etc/hosts
* Mac（苹果电脑）系统：/etc/hosts

### 生效
* Windows：在 CMD 窗口输入：ipconfig /flushdns
* Linux 命令：sudo rcnscd restart
* Mac 命令：sudo killall -HUP mDNSResponder