# FreeGithub
自动获取github相关网站的ip地址，解决github链接不畅通的问题，该项目的README会自动更新，大家可以直接复制使用

## hosts列表
```base
# Host Start
3.236.73.33 github.com
3.235.68.41 api.github.com
3.238.51.165 github.map.fastly.net
54.175.52.10 github.githubassets.com
52.91.217.162 github.global.ssl.fastly.net
54.147.136.237 raw.githubusercontent.com
54.175.52.10 camo.githubusercontent.com
54.175.160.87 avatars.githubusercontent.com
34.238.167.103 avatars0.githubusercontent.com
54.174.86.239 avatars1.githubusercontent.com
54.174.86.239 avatars2.githubusercontent.com
54.242.0.218 avatars3.githubusercontent.com
3.237.14.57 avatars4.githubusercontent.com
3.235.24.187 avatars5.githubusercontent.com
3.95.240.201 avatars6.githubusercontent.com
52.90.59.155 avatars7.githubusercontent.com
54.83.88.86 avatars8.githubusercontent.com
3.95.240.201 favicons.githubusercontent.com
44.193.72.43 developer.apple.com
34.228.240.255 devstreaming-cdn.apple.com
# Host End
```

更新时间：2021-11-26 08:42:27

## 修改本机的hosts文件
### 存放位置
* Windows 系统：C:\Windows\System32\drivers\etc\hosts
* Linux 系统：/etc/hosts
* Mac（苹果电脑）系统：/etc/hosts

### 生效
* Windows：在 CMD 窗口输入：ipconfig /flushdns
* Linux 命令：sudo rcnscd restart
* Mac 命令：sudo killall -HUP mDNSResponder