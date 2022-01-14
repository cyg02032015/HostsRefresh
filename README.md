# FreeGithub
自动获取github相关网站的ip地址，解决github链接不畅通的问题，该项目的README会自动更新，大家可以直接复制使用

## hosts列表
```base
# Host Start
3.91.246.208 github.com
54.210.149.59 api.github.com
40.122.134.122 github.map.fastly.net
54.221.152.96 github.githubassets.com
174.129.191.255 github.global.ssl.fastly.net
44.192.84.249 raw.githubusercontent.com
3.93.3.61 camo.githubusercontent.com
54.242.155.162 avatars.githubusercontent.com
34.204.193.33 avatars0.githubusercontent.com
3.239.92.147 avatars1.githubusercontent.com
3.84.227.117 avatars2.githubusercontent.com
52.3.232.28 avatars3.githubusercontent.com
34.224.32.88 avatars4.githubusercontent.com
54.159.177.221 avatars5.githubusercontent.com
52.205.14.112 avatars6.githubusercontent.com
54.198.66.190 avatars7.githubusercontent.com
54.196.239.20 avatars8.githubusercontent.com
54.208.153.224 favicons.githubusercontent.com
35.173.219.23 developer.apple.com
3.239.92.147 devstreaming-cdn.apple.com
# Host End
```

更新时间：2022-01-14 08:49:59

## 修改本机的hosts文件
### 存放位置
* Windows 系统：C:\Windows\System32\drivers\etc\hosts
* Linux 系统：/etc/hosts
* Mac（苹果电脑）系统：/etc/hosts

### 生效
* Windows：在 CMD 窗口输入：ipconfig /flushdns
* Linux 命令：sudo rcnscd restart
* Mac 命令：sudo killall -HUP mDNSResponder