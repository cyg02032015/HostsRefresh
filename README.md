# FreeGithub
自动获取github相关网站的ip地址，解决github链接不畅通的问题，该项目的README会自动更新，大家可以直接复制使用

## hosts列表
```base
# Host Start
34.211.39.125 github.com
34.219.21.203 api.github.com
35.86.147.145 github.map.fastly.net
35.163.19.139 github.githubassets.com
54.245.187.195 github.global.ssl.fastly.net
54.202.19.116 raw.githubusercontent.com
54.214.197.180 camo.githubusercontent.com
54.245.187.195 avatars.githubusercontent.com
54.214.197.180 avatars0.githubusercontent.com
20.115.166.105 avatars1.githubusercontent.com
54.70.68.231 avatars2.githubusercontent.com
54.244.69.119 avatars3.githubusercontent.com
35.163.19.139 avatars4.githubusercontent.com
34.211.39.125 avatars5.githubusercontent.com
54.202.131.139 avatars6.githubusercontent.com
34.209.250.210 avatars7.githubusercontent.com
35.163.19.139 avatars8.githubusercontent.com
54.214.197.180 favicons.githubusercontent.com
34.209.250.210 developer.apple.com
34.209.250.210 devstreaming-cdn.apple.com
# Host End
```

更新时间：2021-12-03 08:43:48

## 修改本机的hosts文件
### 存放位置
* Windows 系统：C:\Windows\System32\drivers\etc\hosts
* Linux 系统：/etc/hosts
* Mac（苹果电脑）系统：/etc/hosts

### 生效
* Windows：在 CMD 窗口输入：ipconfig /flushdns
* Linux 命令：sudo rcnscd restart
* Mac 命令：sudo killall -HUP mDNSResponder