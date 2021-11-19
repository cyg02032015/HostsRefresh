# FreeGithub
自动获取github相关网站的ip地址，解决github链接不畅通的问题，该项目的README会自动更新，大家可以直接复制使用

## hosts列表
```base
# Host Start
40.69.137.44 github.com
35.183.103.137 api.github.com
99.79.50.248 github.map.fastly.net
99.79.45.63 github.githubassets.com
99.79.45.63 github.global.ssl.fastly.net
99.79.45.63 raw.githubusercontent.com
99.79.45.63 camo.githubusercontent.com
35.183.103.137 avatars.githubusercontent.com
35.183.103.137 avatars0.githubusercontent.com
35.182.23.190 avatars1.githubusercontent.com
99.79.50.248 avatars2.githubusercontent.com
99.79.50.248 avatars3.githubusercontent.com
99.79.45.63 avatars4.githubusercontent.com
35.182.23.190 avatars5.githubusercontent.com
99.79.45.63 avatars6.githubusercontent.com
35.183.243.128 avatars7.githubusercontent.com
35.183.243.128 avatars8.githubusercontent.com
35.182.23.190 favicons.githubusercontent.com
99.79.50.248 developer.apple.com
35.183.103.137 devstreaming-cdn.apple.com
# Host End
```

更新时间：2021-11-19 08:43:07

## 修改本机的hosts文件
### 存放位置
* Windows 系统：C:\Windows\System32\drivers\etc\hosts
* Linux 系统：/etc/hosts
* Mac（苹果电脑）系统：/etc/hosts

### 生效
* Windows：在 CMD 窗口输入：ipconfig /flushdns
* Linux 命令：sudo rcnscd restart
* Mac 命令：sudo killall -HUP mDNSResponder