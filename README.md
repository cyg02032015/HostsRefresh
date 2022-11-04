# FreeGithub
自动获取github相关网站的ip地址，解决github链接不畅通的问题，该项目的README会自动更新，大家可以直接复制使用

## hosts列表
```base
# Host Start
140.82.114.3 github.com
140.82.114.3 api.github.com
151.101.1.6 github.map.fastly.net
3.235.244.127 github.githubassets.com
151.101.1.6 github.global.ssl.fastly.net
3.236.107.192 raw.githubusercontent.com
34.235.134.121 camo.githubusercontent.com
3.228.20.75 avatars.githubusercontent.com
3.236.71.6 avatars0.githubusercontent.com
3.236.13.35 avatars1.githubusercontent.com
18.212.147.175 avatars2.githubusercontent.com
50.17.53.3 avatars3.githubusercontent.com
3.239.198.52 avatars4.githubusercontent.com
3.228.20.75 avatars5.githubusercontent.com
44.210.86.58 avatars6.githubusercontent.com
44.198.52.245 avatars7.githubusercontent.com
18.212.147.175 avatars8.githubusercontent.com
44.198.52.245 favicons.githubusercontent.com
17.253.144.10 developer.apple.com
17.253.144.10 devstreaming-cdn.apple.com
# Host End
```

更新时间：2022-11-04 09:24:47

## 修改本机的hosts文件
### 存放位置
* Windows 系统：C:\Windows\System32\drivers\etc\hosts
* Linux 系统：/etc/hosts
* Mac（苹果电脑）系统：/etc/hosts

### 生效
* Windows：在 CMD 窗口输入：ipconfig /flushdns
* Linux 命令：sudo rcnscd restart
* Mac 命令：sudo killall -HUP mDNSResponder