# FreeGithub
自动获取github相关网站的ip地址，解决github链接不畅通的问题，该项目的README会自动更新，大家可以直接复制使用

## hosts列表
```base
# Host Start
140.82.114.4 github.com
140.82.114.4 api.github.com
151.101.1.6 github.map.fastly.net
54.165.28.21 github.githubassets.com
151.101.1.6 github.global.ssl.fastly.net
44.204.62.139 raw.githubusercontent.com
18.204.243.220 camo.githubusercontent.com
3.237.204.253 avatars.githubusercontent.com
3.86.238.65 avatars0.githubusercontent.com
44.210.147.199 avatars1.githubusercontent.com
3.89.63.76 avatars2.githubusercontent.com
34.232.48.149 avatars3.githubusercontent.com
35.153.199.152 avatars4.githubusercontent.com
3.84.10.137 avatars5.githubusercontent.com
54.227.31.224 avatars6.githubusercontent.com
54.198.33.141 avatars7.githubusercontent.com
34.224.98.86 avatars8.githubusercontent.com
3.89.69.17 favicons.githubusercontent.com
17.253.144.10 developer.apple.com
17.253.144.10 devstreaming-cdn.apple.com
# Host End
```

更新时间：2022-09-02 09:27:11

## 修改本机的hosts文件
### 存放位置
* Windows 系统：C:\Windows\System32\drivers\etc\hosts
* Linux 系统：/etc/hosts
* Mac（苹果电脑）系统：/etc/hosts

### 生效
* Windows：在 CMD 窗口输入：ipconfig /flushdns
* Linux 命令：sudo rcnscd restart
* Mac 命令：sudo killall -HUP mDNSResponder