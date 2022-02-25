# FreeGithub
自动获取github相关网站的ip地址，解决github链接不畅通的问题，该项目的README会自动更新，大家可以直接复制使用

## hosts列表
```base
# Host Start
54.166.174.183 github.com
3.238.100.205 api.github.com
54.198.20.210 github.map.fastly.net
54.235.38.109 github.githubassets.com
3.82.20.235 github.global.ssl.fastly.net
54.198.20.210 raw.githubusercontent.com
34.204.187.222 camo.githubusercontent.com
54.204.140.149 avatars.githubusercontent.com
44.202.194.139 avatars0.githubusercontent.com
18.204.204.231 avatars1.githubusercontent.com
3.82.20.235 avatars2.githubusercontent.com
3.208.15.223 avatars3.githubusercontent.com
54.167.134.190 avatars4.githubusercontent.com
3.235.135.89 avatars5.githubusercontent.com
54.90.193.209 avatars6.githubusercontent.com
54.167.134.190 avatars7.githubusercontent.com
52.54.77.127 avatars8.githubusercontent.com
3.236.17.170 favicons.githubusercontent.com
35.172.100.254 developer.apple.com
3.233.224.158 devstreaming-cdn.apple.com
# Host End
```

更新时间：2022-02-25 08:55:00

## 修改本机的hosts文件
### 存放位置
* Windows 系统：C:\Windows\System32\drivers\etc\hosts
* Linux 系统：/etc/hosts
* Mac（苹果电脑）系统：/etc/hosts

### 生效
* Windows：在 CMD 窗口输入：ipconfig /flushdns
* Linux 命令：sudo rcnscd restart
* Mac 命令：sudo killall -HUP mDNSResponder