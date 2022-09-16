# FreeGithub
自动获取github相关网站的ip地址，解决github链接不畅通的问题，该项目的README会自动更新，大家可以直接复制使用

## hosts列表
```base
# Host Start
140.82.112.3 github.com
140.82.112.3 api.github.com
151.101.1.6 github.map.fastly.net
54.82.63.1 github.githubassets.com
151.101.1.6 github.global.ssl.fastly.net
18.234.108.23 raw.githubusercontent.com
44.200.72.139 camo.githubusercontent.com
54.82.63.1 avatars.githubusercontent.com
54.236.216.176 avatars0.githubusercontent.com
3.236.115.7 avatars1.githubusercontent.com
23.23.25.237 avatars2.githubusercontent.com
44.201.178.106 avatars3.githubusercontent.com
35.173.48.203 avatars4.githubusercontent.com
52.90.113.50 avatars5.githubusercontent.com
44.203.183.242 avatars6.githubusercontent.com
35.175.111.53 avatars7.githubusercontent.com
54.174.207.47 avatars8.githubusercontent.com
18.205.96.166 favicons.githubusercontent.com
17.253.144.10 developer.apple.com
17.253.144.10 devstreaming-cdn.apple.com
# Host End
```

更新时间：2022-09-16 09:51:06

## 修改本机的hosts文件
### 存放位置
* Windows 系统：C:\Windows\System32\drivers\etc\hosts
* Linux 系统：/etc/hosts
* Mac（苹果电脑）系统：/etc/hosts

### 生效
* Windows：在 CMD 窗口输入：ipconfig /flushdns
* Linux 命令：sudo rcnscd restart
* Mac 命令：sudo killall -HUP mDNSResponder