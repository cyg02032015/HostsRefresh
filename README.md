# FreeGithub
自动获取github相关网站的ip地址，解决github链接不畅通的问题，该项目的README会自动更新，大家可以直接复制使用

## hosts列表
```base
# Host Start
140.82.113.3 github.com
140.82.113.3 api.github.com
151.101.1.6 github.map.fastly.net
54.221.150.124 github.githubassets.com
151.101.1.6 github.global.ssl.fastly.net
100.26.143.157 raw.githubusercontent.com
52.200.68.233 camo.githubusercontent.com
3.80.54.99 avatars.githubusercontent.com
3.237.47.17 avatars0.githubusercontent.com
54.167.78.100 avatars1.githubusercontent.com
44.210.80.72 avatars2.githubusercontent.com
54.167.41.204 avatars3.githubusercontent.com
3.237.195.146 avatars4.githubusercontent.com
54.205.165.130 avatars5.githubusercontent.com
3.239.93.250 avatars6.githubusercontent.com
54.242.50.121 avatars7.githubusercontent.com
3.235.75.224 avatars8.githubusercontent.com
50.16.22.113 favicons.githubusercontent.com
17.253.144.10 developer.apple.com
17.253.144.10 devstreaming-cdn.apple.com
# Host End
```

更新时间：2023-02-10 09:09:51

## 修改本机的hosts文件
### 存放位置
* Windows 系统：C:\Windows\System32\drivers\etc\hosts
* Linux 系统：/etc/hosts
* Mac（苹果电脑）系统：/etc/hosts

### 生效
* Windows：在 CMD 窗口输入：ipconfig /flushdns
* Linux 命令：sudo rcnscd restart
* Mac 命令：sudo killall -HUP mDNSResponder