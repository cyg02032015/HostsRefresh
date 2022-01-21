# FreeGithub
自动获取github相关网站的ip地址，解决github链接不畅通的问题，该项目的README会自动更新，大家可以直接复制使用

## hosts列表
```base
# Host Start
3.90.239.124 github.com
3.92.209.225 api.github.com
18.207.251.32 github.map.fastly.net
44.200.171.136 github.githubassets.com
54.237.223.189 github.global.ssl.fastly.net
3.85.241.213 raw.githubusercontent.com
3.89.153.51 camo.githubusercontent.com
54.85.159.76 avatars.githubusercontent.com
34.204.193.226 avatars0.githubusercontent.com
18.233.223.195 avatars1.githubusercontent.com
54.87.41.194 avatars2.githubusercontent.com
44.200.109.46 avatars3.githubusercontent.com
100.27.25.242 avatars4.githubusercontent.com
174.129.145.106 avatars5.githubusercontent.com
54.234.1.143 avatars6.githubusercontent.com
18.233.223.195 avatars7.githubusercontent.com
54.234.74.138 avatars8.githubusercontent.com
3.83.129.157 favicons.githubusercontent.com
54.234.1.143 developer.apple.com
44.202.127.77 devstreaming-cdn.apple.com
# Host End
```

更新时间：2022-01-21 08:44:21

## 修改本机的hosts文件
### 存放位置
* Windows 系统：C:\Windows\System32\drivers\etc\hosts
* Linux 系统：/etc/hosts
* Mac（苹果电脑）系统：/etc/hosts

### 生效
* Windows：在 CMD 窗口输入：ipconfig /flushdns
* Linux 命令：sudo rcnscd restart
* Mac 命令：sudo killall -HUP mDNSResponder