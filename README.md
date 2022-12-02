# FreeGithub
自动获取github相关网站的ip地址，解决github链接不畅通的问题，该项目的README会自动更新，大家可以直接复制使用

## hosts列表
```base
# Host Start
140.82.114.3 github.com
140.82.114.3 api.github.com
151.101.1.6 github.map.fastly.net
3.235.227.176 github.githubassets.com
151.101.1.6 github.global.ssl.fastly.net
54.210.61.40 raw.githubusercontent.com
3.237.12.224 camo.githubusercontent.com
54.91.104.76 avatars.githubusercontent.com
3.237.28.184 avatars0.githubusercontent.com
54.211.191.243 avatars1.githubusercontent.com
3.234.154.46 avatars2.githubusercontent.com
54.91.104.76 avatars3.githubusercontent.com
34.237.75.28 avatars4.githubusercontent.com
3.85.202.0 avatars5.githubusercontent.com
44.193.212.48 avatars6.githubusercontent.com
44.197.114.254 avatars7.githubusercontent.com
44.192.92.178 avatars8.githubusercontent.com
34.204.173.83 favicons.githubusercontent.com
17.253.144.10 developer.apple.com
17.253.144.10 devstreaming-cdn.apple.com
# Host End
```

更新时间：2022-12-02 09:04:35

## 修改本机的hosts文件
### 存放位置
* Windows 系统：C:\Windows\System32\drivers\etc\hosts
* Linux 系统：/etc/hosts
* Mac（苹果电脑）系统：/etc/hosts

### 生效
* Windows：在 CMD 窗口输入：ipconfig /flushdns
* Linux 命令：sudo rcnscd restart
* Mac 命令：sudo killall -HUP mDNSResponder