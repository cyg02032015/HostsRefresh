# FreeGithub
自动获取github相关网站的ip地址，解决github链接不畅通的问题，该项目的README会自动更新，大家可以直接复制使用

## hosts列表
```base
# Host Start
34.238.40.212 github.com
3.82.208.30 api.github.com
54.145.61.56 github.map.fastly.net
54.166.169.44 github.githubassets.com
54.161.234.19 github.global.ssl.fastly.net
3.94.184.52 raw.githubusercontent.com
54.145.61.56 camo.githubusercontent.com
3.80.144.219 avatars.githubusercontent.com
3.219.35.66 avatars0.githubusercontent.com
54.204.55.250 avatars1.githubusercontent.com
54.204.55.250 avatars2.githubusercontent.com
54.89.52.237 avatars3.githubusercontent.com
3.82.208.30 avatars4.githubusercontent.com
54.161.234.19 avatars5.githubusercontent.com
3.87.133.163 avatars6.githubusercontent.com
18.207.224.3 avatars7.githubusercontent.com
3.95.151.219 avatars8.githubusercontent.com
3.235.105.88 favicons.githubusercontent.com
3.95.151.219 developer.apple.com
44.198.175.100 devstreaming-cdn.apple.com
# Host End
```

更新时间：2021-10-29 08:40:34

## 修改本机的hosts文件
### 存放位置
* Windows 系统：C:\Windows\System32\drivers\etc\hosts
* Linux 系统：/etc/hosts
* Mac（苹果电脑）系统：/etc/hosts

### 生效
* Windows：在 CMD 窗口输入：ipconfig /flushdns
* Linux 命令：sudo rcnscd restart
* Mac 命令：sudo killall -HUP mDNSResponder