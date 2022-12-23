# FreeGithub
自动获取github相关网站的ip地址，解决github链接不畅通的问题，该项目的README会自动更新，大家可以直接复制使用

## hosts列表
```base
# Host Start
140.82.112.3 github.com
140.82.112.3 api.github.com
151.101.1.6 github.map.fastly.net
18.232.171.78 github.githubassets.com
151.101.1.6 github.global.ssl.fastly.net
44.211.84.85 raw.githubusercontent.com
3.234.236.162 camo.githubusercontent.com
3.235.229.163 avatars.githubusercontent.com
 avatars0.githubusercontent.com
44.200.193.97 avatars1.githubusercontent.com
44.192.247.67 avatars2.githubusercontent.com
44.200.193.97 avatars3.githubusercontent.com
44.210.233.94 avatars4.githubusercontent.com
44.211.84.85 avatars5.githubusercontent.com
44.192.247.67 avatars6.githubusercontent.com
3.236.45.87 avatars7.githubusercontent.com
3.239.2.66 avatars8.githubusercontent.com
44.211.84.85 favicons.githubusercontent.com
17.253.144.10 developer.apple.com
17.253.144.10 devstreaming-cdn.apple.com
# Host End
```

更新时间：2022-12-23 09:00:47

## 修改本机的hosts文件
### 存放位置
* Windows 系统：C:\Windows\System32\drivers\etc\hosts
* Linux 系统：/etc/hosts
* Mac（苹果电脑）系统：/etc/hosts

### 生效
* Windows：在 CMD 窗口输入：ipconfig /flushdns
* Linux 命令：sudo rcnscd restart
* Mac 命令：sudo killall -HUP mDNSResponder