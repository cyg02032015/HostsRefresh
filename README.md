# FreeGithub
自动获取github相关网站的ip地址，解决github链接不畅通的问题，该项目的README会自动更新，大家可以直接复制使用

## hosts列表
```base
# Host Start
34.219.120.234 github.com
34.219.120.234 api.github.com
35.85.59.19 github.map.fastly.net
54.190.44.93 github.githubassets.com
18.237.185.16 github.global.ssl.fastly.net
34.221.120.2 raw.githubusercontent.com
54.184.178.155 camo.githubusercontent.com
34.220.30.171 avatars.githubusercontent.com
35.87.6.30 avatars0.githubusercontent.com
35.87.18.50 avatars1.githubusercontent.com
35.164.15.47 avatars2.githubusercontent.com
35.89.94.222 avatars3.githubusercontent.com
52.32.129.217 avatars4.githubusercontent.com
35.89.94.222 avatars5.githubusercontent.com
35.86.117.188 avatars6.githubusercontent.com
34.221.120.2 avatars7.githubusercontent.com
34.212.122.2 avatars8.githubusercontent.com
52.32.129.217 favicons.githubusercontent.com
34.221.120.2 developer.apple.com
35.87.18.50 devstreaming-cdn.apple.com
# Host End
```

更新时间：2022-04-22 09:12:24

## 修改本机的hosts文件
### 存放位置
* Windows 系统：C:\Windows\System32\drivers\etc\hosts
* Linux 系统：/etc/hosts
* Mac（苹果电脑）系统：/etc/hosts

### 生效
* Windows：在 CMD 窗口输入：ipconfig /flushdns
* Linux 命令：sudo rcnscd restart
* Mac 命令：sudo killall -HUP mDNSResponder