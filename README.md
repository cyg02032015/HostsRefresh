# FreeGithub
自动获取github相关网站的ip地址，解决github链接不畅通的问题，该项目的README会自动更新，大家可以直接复制使用

## hosts列表
```base
# Host Start
140.82.112.4 github.com
140.82.112.4 api.github.com
151.101.1.6 github.map.fastly.net
54.198.125.236 github.githubassets.com
151.101.1.6 github.global.ssl.fastly.net
54.158.61.222 raw.githubusercontent.com
3.235.24.65 camo.githubusercontent.com
54.166.250.216 avatars.githubusercontent.com
3.94.247.19 avatars0.githubusercontent.com
54.166.250.216 avatars1.githubusercontent.com
23.20.196.33 avatars2.githubusercontent.com
44.201.252.129 avatars3.githubusercontent.com
54.164.32.164 avatars4.githubusercontent.com
44.203.199.65 avatars5.githubusercontent.com
3.236.23.80 avatars6.githubusercontent.com
54.84.29.196 avatars7.githubusercontent.com
3.80.192.227 avatars8.githubusercontent.com
3.81.75.69 favicons.githubusercontent.com
17.253.144.10 developer.apple.com
17.253.144.10 devstreaming-cdn.apple.com
# Host End
```

更新时间：2022-12-30 09:02:52

## 修改本机的hosts文件
### 存放位置
* Windows 系统：C:\Windows\System32\drivers\etc\hosts
* Linux 系统：/etc/hosts
* Mac（苹果电脑）系统：/etc/hosts

### 生效
* Windows：在 CMD 窗口输入：ipconfig /flushdns
* Linux 命令：sudo rcnscd restart
* Mac 命令：sudo killall -HUP mDNSResponder