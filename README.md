# FreeGithub
自动获取github相关网站的ip地址，解决github链接不畅通的问题，该项目的README会自动更新，大家可以直接复制使用

## hosts列表
```base
# Host Start
52.204.220.231 github.com
52.242.78.150 api.github.com
34.205.191.178 github.map.fastly.net
3.235.249.140 github.githubassets.com
3.230.158.116 github.global.ssl.fastly.net
54.172.69.188 raw.githubusercontent.com
35.175.125.135 camo.githubusercontent.com
34.207.106.115 avatars.githubusercontent.com
54.198.101.23 avatars0.githubusercontent.com
35.175.121.170 avatars1.githubusercontent.com
100.26.22.225 avatars2.githubusercontent.com
52.205.255.249 avatars3.githubusercontent.com
3.95.18.95 avatars4.githubusercontent.com
18.209.162.23 avatars5.githubusercontent.com
18.209.162.23 avatars6.githubusercontent.com
52.205.255.249 avatars7.githubusercontent.com
3.236.177.188 avatars8.githubusercontent.com
3.236.177.188 favicons.githubusercontent.com
44.192.5.151 developer.apple.com
44.192.94.73 devstreaming-cdn.apple.com
# Host End
```

更新时间：2022-01-28 08:45:31

## 修改本机的hosts文件
### 存放位置
* Windows 系统：C:\Windows\System32\drivers\etc\hosts
* Linux 系统：/etc/hosts
* Mac（苹果电脑）系统：/etc/hosts

### 生效
* Windows：在 CMD 窗口输入：ipconfig /flushdns
* Linux 命令：sudo rcnscd restart
* Mac 命令：sudo killall -HUP mDNSResponder