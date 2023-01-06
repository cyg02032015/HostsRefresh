# FreeGithub
自动获取github相关网站的ip地址，解决github链接不畅通的问题，该项目的README会自动更新，大家可以直接复制使用

## hosts列表
```base
# Host Start
140.82.114.3 github.com
140.82.114.3 api.github.com
151.101.1.6 github.map.fastly.net
3.238.86.195 github.githubassets.com
151.101.1.6 github.global.ssl.fastly.net
3.237.40.61 raw.githubusercontent.com
18.233.99.35 camo.githubusercontent.com
3.238.86.195 avatars.githubusercontent.com
34.204.108.211 avatars0.githubusercontent.com
3.83.93.214 avatars1.githubusercontent.com
3.238.86.195 avatars2.githubusercontent.com
3.238.241.244 avatars3.githubusercontent.com
3.83.93.214 avatars4.githubusercontent.com
3.239.122.221 avatars5.githubusercontent.com
18.205.192.105 avatars6.githubusercontent.com
18.233.99.35 avatars7.githubusercontent.com
34.231.20.62 avatars8.githubusercontent.com
34.231.122.90 favicons.githubusercontent.com
17.253.144.10 developer.apple.com
17.253.144.10 devstreaming-cdn.apple.com
# Host End
```

更新时间：2023-01-06 09:05:44

## 修改本机的hosts文件
### 存放位置
* Windows 系统：C:\Windows\System32\drivers\etc\hosts
* Linux 系统：/etc/hosts
* Mac（苹果电脑）系统：/etc/hosts

### 生效
* Windows：在 CMD 窗口输入：ipconfig /flushdns
* Linux 命令：sudo rcnscd restart
* Mac 命令：sudo killall -HUP mDNSResponder