# FreeGithub
自动获取github相关网站的ip地址，解决github链接不畅通的问题，该项目的README会自动更新，大家可以直接复制使用

## hosts列表
```base
# Host Start
140.82.112.3 github.com
140.82.112.3 api.github.com
151.101.1.6 github.map.fastly.net
54.158.81.91 github.githubassets.com
151.101.1.6 github.global.ssl.fastly.net
3.236.241.165 raw.githubusercontent.com
18.206.245.196 camo.githubusercontent.com
44.204.255.133 avatars.githubusercontent.com
3.230.143.22 avatars0.githubusercontent.com
34.230.57.184 avatars1.githubusercontent.com
3.218.207.220 avatars2.githubusercontent.com
100.26.60.127 avatars3.githubusercontent.com
3.216.91.125 avatars4.githubusercontent.com
18.206.245.196 avatars5.githubusercontent.com
3.94.133.41 avatars6.githubusercontent.com
3.237.48.41 avatars7.githubusercontent.com
44.201.18.98 avatars8.githubusercontent.com
3.237.8.5 favicons.githubusercontent.com
17.253.144.10 developer.apple.com
17.253.144.10 devstreaming-cdn.apple.com
# Host End
```

更新时间：2022-09-09 09:49:12

## 修改本机的hosts文件
### 存放位置
* Windows 系统：C:\Windows\System32\drivers\etc\hosts
* Linux 系统：/etc/hosts
* Mac（苹果电脑）系统：/etc/hosts

### 生效
* Windows：在 CMD 窗口输入：ipconfig /flushdns
* Linux 命令：sudo rcnscd restart
* Mac 命令：sudo killall -HUP mDNSResponder