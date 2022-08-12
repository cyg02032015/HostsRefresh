# FreeGithub
自动获取github相关网站的ip地址，解决github链接不畅通的问题，该项目的README会自动更新，大家可以直接复制使用

## hosts列表
```base
# Host Start
140.82.114.3 github.com
140.82.114.3 api.github.com
151.101.1.6 github.map.fastly.net
18.237.35.207 github.githubassets.com
151.101.1.6 github.global.ssl.fastly.net
54.245.195.176 raw.githubusercontent.com
52.12.179.144 camo.githubusercontent.com
18.237.35.207 avatars.githubusercontent.com
35.90.220.57 avatars0.githubusercontent.com
54.185.222.150 avatars1.githubusercontent.com
54.185.222.150 avatars2.githubusercontent.com
54.213.200.112 avatars3.githubusercontent.com
18.237.35.207 avatars4.githubusercontent.com
35.90.196.231 avatars5.githubusercontent.com
35.91.188.132 avatars6.githubusercontent.com
54.188.54.10 avatars7.githubusercontent.com
52.40.209.81 avatars8.githubusercontent.com
54.214.229.64 favicons.githubusercontent.com
17.253.144.10 developer.apple.com
17.253.144.10 devstreaming-cdn.apple.com
# Host End
```

更新时间：2022-08-12 09:18:15

## 修改本机的hosts文件
### 存放位置
* Windows 系统：C:\Windows\System32\drivers\etc\hosts
* Linux 系统：/etc/hosts
* Mac（苹果电脑）系统：/etc/hosts

### 生效
* Windows：在 CMD 窗口输入：ipconfig /flushdns
* Linux 命令：sudo rcnscd restart
* Mac 命令：sudo killall -HUP mDNSResponder