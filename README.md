# FreeGithub
自动获取github相关网站的ip地址，解决github链接不畅通的问题，该项目的README会自动更新，大家可以直接复制使用

## hosts列表
```base
# Host Start
54.185.200.2 github.com
54.149.114.105 api.github.com
34.219.42.116 github.map.fastly.net
54.213.192.94 github.githubassets.com
54.212.126.15 github.global.ssl.fastly.net
54.213.192.94 raw.githubusercontent.com
20.109.167.188 camo.githubusercontent.com
34.219.183.66 avatars.githubusercontent.com
54.191.85.23 avatars0.githubusercontent.com
54.213.90.230 avatars1.githubusercontent.com
54.218.251.18 avatars2.githubusercontent.com
54.213.90.230 avatars3.githubusercontent.com
54.185.200.2 avatars4.githubusercontent.com
34.221.3.228 avatars5.githubusercontent.com
54.213.90.230 avatars6.githubusercontent.com
52.26.86.110 avatars7.githubusercontent.com
54.212.126.15 avatars8.githubusercontent.com
20.109.167.188 favicons.githubusercontent.com
34.219.183.66 developer.apple.com
34.222.167.119 devstreaming-cdn.apple.com
# Host End
```

更新时间：2021-12-17 08:48:00

## 修改本机的hosts文件
### 存放位置
* Windows 系统：C:\Windows\System32\drivers\etc\hosts
* Linux 系统：/etc/hosts
* Mac（苹果电脑）系统：/etc/hosts

### 生效
* Windows：在 CMD 窗口输入：ipconfig /flushdns
* Linux 命令：sudo rcnscd restart
* Mac 命令：sudo killall -HUP mDNSResponder