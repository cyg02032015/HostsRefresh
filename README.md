# FreeGithub
自动获取github相关网站的ip地址，解决github链接不畅通的问题，该项目的README会自动更新，大家可以直接复制使用

## hosts列表
```base
# Host Start
140.82.113.4 github.com
140.82.113.4 api.github.com
151.101.1.6 github.map.fastly.net
34.217.112.255 github.githubassets.com
151.101.1.6 github.global.ssl.fastly.net
34.220.5.141 raw.githubusercontent.com
52.24.169.109 camo.githubusercontent.com
35.93.30.85 avatars.githubusercontent.com
35.89.173.18 avatars0.githubusercontent.com
34.220.5.141 avatars1.githubusercontent.com
54.202.52.40 avatars2.githubusercontent.com
35.88.97.126 avatars3.githubusercontent.com
35.91.117.238 avatars4.githubusercontent.com
35.90.214.107 avatars5.githubusercontent.com
34.220.5.141 avatars6.githubusercontent.com
35.92.205.241 avatars7.githubusercontent.com
54.202.52.40 avatars8.githubusercontent.com
54.202.52.40 favicons.githubusercontent.com
17.253.144.10 developer.apple.com
17.253.144.10 devstreaming-cdn.apple.com
# Host End
```

更新时间：2022-10-07 09:45:57

## 修改本机的hosts文件
### 存放位置
* Windows 系统：C:\Windows\System32\drivers\etc\hosts
* Linux 系统：/etc/hosts
* Mac（苹果电脑）系统：/etc/hosts

### 生效
* Windows：在 CMD 窗口输入：ipconfig /flushdns
* Linux 命令：sudo rcnscd restart
* Mac 命令：sudo killall -HUP mDNSResponder