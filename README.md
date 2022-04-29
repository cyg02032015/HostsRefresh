# FreeGithub
自动获取github相关网站的ip地址，解决github链接不畅通的问题，该项目的README会自动更新，大家可以直接复制使用

## hosts列表
```base
# Host Start
52.203.90.68 github.com
54.90.234.49 api.github.com
54.81.146.159 github.map.fastly.net
34.200.243.248 github.githubassets.com
3.237.36.9 github.global.ssl.fastly.net
3.236.55.178 raw.githubusercontent.com
3.235.145.255 camo.githubusercontent.com
54.91.141.41 avatars.githubusercontent.com
3.90.49.48 avatars0.githubusercontent.com
54.227.89.138 avatars1.githubusercontent.com
34.227.13.130 avatars2.githubusercontent.com
44.201.2.60 avatars3.githubusercontent.com
34.227.242.52 avatars4.githubusercontent.com
3.83.20.131 avatars5.githubusercontent.com
20.97.226.225 avatars6.githubusercontent.com
54.227.55.195 avatars7.githubusercontent.com
34.227.13.130 avatars8.githubusercontent.com
34.229.109.193 favicons.githubusercontent.com
54.157.210.195 developer.apple.com
44.204.124.57 devstreaming-cdn.apple.com
# Host End
```

更新时间：2022-04-29 09:12:12

## 修改本机的hosts文件
### 存放位置
* Windows 系统：C:\Windows\System32\drivers\etc\hosts
* Linux 系统：/etc/hosts
* Mac（苹果电脑）系统：/etc/hosts

### 生效
* Windows：在 CMD 窗口输入：ipconfig /flushdns
* Linux 命令：sudo rcnscd restart
* Mac 命令：sudo killall -HUP mDNSResponder