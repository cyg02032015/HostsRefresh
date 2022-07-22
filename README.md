# FreeGithub
自动获取github相关网站的ip地址，解决github链接不畅通的问题，该项目的README会自动更新，大家可以直接复制使用

## hosts列表
```base
# Host Start
140.82.112.4 github.com
140.82.112.4 api.github.com
151.101.1.6 github.map.fastly.net
44.201.154.119 github.githubassets.com
151.101.1.6 github.global.ssl.fastly.net
3.235.168.0 raw.githubusercontent.com
34.236.37.148 camo.githubusercontent.com
54.224.146.143 avatars.githubusercontent.com
35.171.133.29 avatars0.githubusercontent.com
54.224.146.143 avatars1.githubusercontent.com
35.172.146.26 avatars2.githubusercontent.com
54.147.35.37 avatars3.githubusercontent.com
44.200.179.25 avatars4.githubusercontent.com
3.92.126.100 avatars5.githubusercontent.com
34.205.225.226 avatars6.githubusercontent.com
44.200.0.237 avatars7.githubusercontent.com
3.235.168.0 avatars8.githubusercontent.com
54.164.203.141 favicons.githubusercontent.com
17.253.144.10 developer.apple.com
17.253.144.10 devstreaming-cdn.apple.com
# Host End
```

更新时间：2022-07-22 09:27:40

## 修改本机的hosts文件
### 存放位置
* Windows 系统：C:\Windows\System32\drivers\etc\hosts
* Linux 系统：/etc/hosts
* Mac（苹果电脑）系统：/etc/hosts

### 生效
* Windows：在 CMD 窗口输入：ipconfig /flushdns
* Linux 命令：sudo rcnscd restart
* Mac 命令：sudo killall -HUP mDNSResponder