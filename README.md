# FreeGithub
自动获取github相关网站的ip地址，解决github链接不畅通的问题，该项目的README会自动更新，大家可以直接复制使用

## hosts列表
```base
# Host Start
3.93.74.152 github.com
34.229.119.8 api.github.com
54.167.69.38 github.map.fastly.net
18.205.7.97 github.githubassets.com
3.82.97.107 github.global.ssl.fastly.net
54.157.197.228 raw.githubusercontent.com
3.93.74.152 camo.githubusercontent.com
3.86.91.107 avatars.githubusercontent.com
35.171.165.129 avatars0.githubusercontent.com
3.237.82.130 avatars1.githubusercontent.com
3.237.35.67 avatars2.githubusercontent.com
18.206.13.104 avatars3.githubusercontent.com
54.80.169.202 avatars4.githubusercontent.com
44.193.11.71 avatars5.githubusercontent.com
44.193.11.71 avatars6.githubusercontent.com
44.193.11.71 avatars7.githubusercontent.com
3.86.91.107 avatars8.githubusercontent.com
54.174.28.81 favicons.githubusercontent.com
34.232.63.122 developer.apple.com
54.80.1.12 devstreaming-cdn.apple.com
# Host End
```

更新时间：2022-03-04 08:59:47

## 修改本机的hosts文件
### 存放位置
* Windows 系统：C:\Windows\System32\drivers\etc\hosts
* Linux 系统：/etc/hosts
* Mac（苹果电脑）系统：/etc/hosts

### 生效
* Windows：在 CMD 窗口输入：ipconfig /flushdns
* Linux 命令：sudo rcnscd restart
* Mac 命令：sudo killall -HUP mDNSResponder