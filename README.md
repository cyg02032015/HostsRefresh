# FreeGithub
自动获取github相关网站的ip地址，解决github链接不畅通的问题，该项目的README会自动更新，大家可以直接复制使用

## hosts列表
```base
# Host Start
140.82.113.3 github.com
140.82.113.3 api.github.com
151.101.1.6 github.map.fastly.net
44.211.245.73 github.githubassets.com
151.101.1.6 github.global.ssl.fastly.net
54.152.215.146 raw.githubusercontent.com
34.207.151.78 camo.githubusercontent.com
54.163.24.91 avatars.githubusercontent.com
54.152.215.146 avatars0.githubusercontent.com
54.160.255.189 avatars1.githubusercontent.com
3.95.250.11 avatars2.githubusercontent.com
35.172.33.103 avatars3.githubusercontent.com
54.91.136.250 avatars4.githubusercontent.com
54.90.91.74 avatars5.githubusercontent.com
3.86.154.137 avatars6.githubusercontent.com
54.160.255.189 avatars7.githubusercontent.com
18.232.174.36 avatars8.githubusercontent.com
52.91.252.35 favicons.githubusercontent.com
17.253.144.10 developer.apple.com
17.253.144.10 devstreaming-cdn.apple.com
# Host End
```

更新时间：2022-12-16 09:01:16

## 修改本机的hosts文件
### 存放位置
* Windows 系统：C:\Windows\System32\drivers\etc\hosts
* Linux 系统：/etc/hosts
* Mac（苹果电脑）系统：/etc/hosts

### 生效
* Windows：在 CMD 窗口输入：ipconfig /flushdns
* Linux 命令：sudo rcnscd restart
* Mac 命令：sudo killall -HUP mDNSResponder