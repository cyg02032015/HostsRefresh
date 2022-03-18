# FreeGithub
自动获取github相关网站的ip地址，解决github链接不畅通的问题，该项目的README会自动更新，大家可以直接复制使用

## hosts列表
```base
# Host Start
3.238.127.3 github.com
52.203.229.255 api.github.com
54.83.66.11 github.map.fastly.net
18.208.211.138 github.githubassets.com
52.177.179.137 github.global.ssl.fastly.net
18.212.61.238 raw.githubusercontent.com
3.238.217.91 camo.githubusercontent.com
18.212.240.124 avatars.githubusercontent.com
3.83.40.172 avatars0.githubusercontent.com
54.89.43.164 avatars1.githubusercontent.com
18.209.70.208 avatars2.githubusercontent.com
3.234.214.31 avatars3.githubusercontent.com
52.177.179.137 avatars4.githubusercontent.com
3.82.110.153 avatars5.githubusercontent.com
3.238.223.33 avatars6.githubusercontent.com
34.234.96.73 avatars7.githubusercontent.com
3.233.240.12 avatars8.githubusercontent.com
54.211.128.117 favicons.githubusercontent.com
54.172.6.159 developer.apple.com
18.204.211.163 devstreaming-cdn.apple.com
# Host End
```

更新时间：2022-03-18 08:55:55

## 修改本机的hosts文件
### 存放位置
* Windows 系统：C:\Windows\System32\drivers\etc\hosts
* Linux 系统：/etc/hosts
* Mac（苹果电脑）系统：/etc/hosts

### 生效
* Windows：在 CMD 窗口输入：ipconfig /flushdns
* Linux 命令：sudo rcnscd restart
* Mac 命令：sudo killall -HUP mDNSResponder