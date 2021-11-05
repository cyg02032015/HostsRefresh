# FreeGithub
自动获取github相关网站的ip地址，解决github链接不畅通的问题，该项目的README会自动更新，大家可以直接复制使用

## hosts列表
```base
# Host Start
3.236.207.34 github.com
3.237.51.79 api.github.com
3.239.20.115 github.map.fastly.net
54.161.238.180 github.githubassets.com
3.237.38.135 github.global.ssl.fastly.net
184.73.12.44 raw.githubusercontent.com
44.198.188.214 camo.githubusercontent.com
18.206.164.237 avatars.githubusercontent.com
54.87.245.92 avatars0.githubusercontent.com
3.237.51.79 avatars1.githubusercontent.com
3.81.81.25 avatars2.githubusercontent.com
3.237.38.135 avatars3.githubusercontent.com
18.206.169.94 avatars4.githubusercontent.com
52.87.192.83 avatars5.githubusercontent.com
18.233.161.234 avatars6.githubusercontent.com
3.239.20.115 avatars7.githubusercontent.com
3.86.32.200 avatars8.githubusercontent.com
3.90.187.96 favicons.githubusercontent.com
3.90.68.141 developer.apple.com
3.81.22.180 devstreaming-cdn.apple.com
# Host End
```

更新时间：2021-11-05 08:42:09

## 修改本机的hosts文件
### 存放位置
* Windows 系统：C:\Windows\System32\drivers\etc\hosts
* Linux 系统：/etc/hosts
* Mac（苹果电脑）系统：/etc/hosts

### 生效
* Windows：在 CMD 窗口输入：ipconfig /flushdns
* Linux 命令：sudo rcnscd restart
* Mac 命令：sudo killall -HUP mDNSResponder