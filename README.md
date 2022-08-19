# FreeGithub
自动获取github相关网站的ip地址，解决github链接不畅通的问题，该项目的README会自动更新，大家可以直接复制使用

## hosts列表
```base
# Host Start
140.82.112.3 github.com
140.82.112.3 api.github.com
151.101.1.6 github.map.fastly.net
54.224.100.55 github.githubassets.com
151.101.1.6 github.global.ssl.fastly.net
34.239.159.142 raw.githubusercontent.com
54.161.20.161 camo.githubusercontent.com
34.234.93.111 avatars.githubusercontent.com
44.200.75.237 avatars0.githubusercontent.com
23.23.0.80 avatars1.githubusercontent.com
3.81.28.228 avatars2.githubusercontent.com
54.161.20.161 avatars3.githubusercontent.com
3.81.28.228 avatars4.githubusercontent.com
52.206.43.194 avatars5.githubusercontent.com
44.199.205.77 avatars6.githubusercontent.com
54.172.106.171 avatars7.githubusercontent.com
54.226.18.125 avatars8.githubusercontent.com
44.203.34.171 favicons.githubusercontent.com
17.253.144.10 developer.apple.com
17.253.144.10 devstreaming-cdn.apple.com
# Host End
```

更新时间：2022-08-19 09:24:18

## 修改本机的hosts文件
### 存放位置
* Windows 系统：C:\Windows\System32\drivers\etc\hosts
* Linux 系统：/etc/hosts
* Mac（苹果电脑）系统：/etc/hosts

### 生效
* Windows：在 CMD 窗口输入：ipconfig /flushdns
* Linux 命令：sudo rcnscd restart
* Mac 命令：sudo killall -HUP mDNSResponder