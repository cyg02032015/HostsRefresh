# FreeGithub
自动获取github相关网站的ip地址，解决github链接不畅通的问题，该项目的README会自动更新，大家可以直接复制使用

## hosts列表
```base
# Host Start
140.82.114.4 github.com
140.82.114.4 api.github.com
151.101.1.6 github.map.fastly.net
3.227.211.134 github.githubassets.com
151.101.1.6 github.global.ssl.fastly.net
54.92.150.102 raw.githubusercontent.com
18.208.137.32 camo.githubusercontent.com
3.92.212.46 avatars.githubusercontent.com
3.236.120.183 avatars0.githubusercontent.com
3.84.214.156 avatars1.githubusercontent.com
44.193.198.148 avatars2.githubusercontent.com
34.204.181.194 avatars3.githubusercontent.com
52.91.246.190 avatars4.githubusercontent.com
54.173.41.45 avatars5.githubusercontent.com
34.228.192.233 avatars6.githubusercontent.com
44.200.115.54 avatars7.githubusercontent.com
3.92.212.46 avatars8.githubusercontent.com
35.170.75.30 favicons.githubusercontent.com
17.253.144.10 developer.apple.com
17.253.144.10 devstreaming-cdn.apple.com
# Host End
```

更新时间：2022-11-11 09:22:28

## 修改本机的hosts文件
### 存放位置
* Windows 系统：C:\Windows\System32\drivers\etc\hosts
* Linux 系统：/etc/hosts
* Mac（苹果电脑）系统：/etc/hosts

### 生效
* Windows：在 CMD 窗口输入：ipconfig /flushdns
* Linux 命令：sudo rcnscd restart
* Mac 命令：sudo killall -HUP mDNSResponder