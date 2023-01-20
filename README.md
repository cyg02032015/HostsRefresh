# FreeGithub
自动获取github相关网站的ip地址，解决github链接不畅通的问题，该项目的README会自动更新，大家可以直接复制使用

## hosts列表
```base
# Host Start
185.199.108.153 github.com
185.199.108.153 api.github.com
151.101.1.6 github.map.fastly.net
44.210.23.26 github.githubassets.com
151.101.1.6 github.global.ssl.fastly.net
3.234.243.140 raw.githubusercontent.com
3.236.93.218 camo.githubusercontent.com
35.173.47.89 avatars.githubusercontent.com
3.236.146.119 avatars0.githubusercontent.com
54.204.136.0 avatars1.githubusercontent.com
3.238.225.47 avatars2.githubusercontent.com
54.88.93.143 avatars3.githubusercontent.com
44.197.112.110 avatars4.githubusercontent.com
54.88.93.143 avatars5.githubusercontent.com
35.172.146.142 avatars6.githubusercontent.com
107.20.17.186 avatars7.githubusercontent.com
3.91.205.158 avatars8.githubusercontent.com
3.238.225.47 favicons.githubusercontent.com
17.253.144.10 developer.apple.com
17.253.144.10 devstreaming-cdn.apple.com
# Host End
```

更新时间：2023-01-20 09:04:37

## 修改本机的hosts文件
### 存放位置
* Windows 系统：C:\Windows\System32\drivers\etc\hosts
* Linux 系统：/etc/hosts
* Mac（苹果电脑）系统：/etc/hosts

### 生效
* Windows：在 CMD 窗口输入：ipconfig /flushdns
* Linux 命令：sudo rcnscd restart
* Mac 命令：sudo killall -HUP mDNSResponder