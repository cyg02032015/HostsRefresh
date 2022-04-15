# FreeGithub
自动获取github相关网站的ip地址，解决github链接不畅通的问题，该项目的README会自动更新，大家可以直接复制使用

## hosts列表
```base
# Host Start
3.91.239.117 github.com
3.87.80.225 api.github.com
34.238.152.83 github.map.fastly.net
75.101.226.54 github.githubassets.com
44.197.174.234 github.global.ssl.fastly.net
3.236.40.58 raw.githubusercontent.com
54.226.226.182 camo.githubusercontent.com
18.233.93.27 avatars.githubusercontent.com
107.20.89.129 avatars0.githubusercontent.com
44.200.101.107 avatars1.githubusercontent.com
54.158.123.244 avatars2.githubusercontent.com
18.205.188.87 avatars3.githubusercontent.com
54.172.207.219 avatars4.githubusercontent.com
44.201.223.65 avatars5.githubusercontent.com
54.226.226.182 avatars6.githubusercontent.com
34.203.223.18 avatars7.githubusercontent.com
54.172.207.219 avatars8.githubusercontent.com
3.236.44.228 favicons.githubusercontent.com
3.236.168.132 developer.apple.com
44.193.12.208 devstreaming-cdn.apple.com
# Host End
```

更新时间：2022-04-15 09:04:50

## 修改本机的hosts文件
### 存放位置
* Windows 系统：C:\Windows\System32\drivers\etc\hosts
* Linux 系统：/etc/hosts
* Mac（苹果电脑）系统：/etc/hosts

### 生效
* Windows：在 CMD 窗口输入：ipconfig /flushdns
* Linux 命令：sudo rcnscd restart
* Mac 命令：sudo killall -HUP mDNSResponder