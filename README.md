# FreeGithub
自动获取github相关网站的ip地址，解决github链接不畅通的问题，该项目的README会自动更新，大家可以直接复制使用

## hosts列表
```base
# Host Start
0.0.1.1 github.com
0.0.1.1 api.github.com
0.0.1.1 github.map.fastly.net
0.0.1.1 github.githubassets.com
0.0.1.1 github.global.ssl.fastly.net
0.0.1.1 raw.githubusercontent.com
0.0.1.1 camo.githubusercontent.com
0.0.1.1 avatars.githubusercontent.com
0.0.1.1 avatars0.githubusercontent.com
0.0.1.1 avatars1.githubusercontent.com
0.0.1.1 avatars2.githubusercontent.com
0.0.1.1 avatars3.githubusercontent.com
0.0.1.1 avatars4.githubusercontent.com
0.0.1.1 avatars5.githubusercontent.com
0.0.1.1 avatars6.githubusercontent.com
0.0.1.1 avatars7.githubusercontent.com
0.0.1.1 avatars8.githubusercontent.com
0.0.1.1 favicons.githubusercontent.com
0.0.1.1 developer.apple.com
0.0.1.1 devstreaming-cdn.apple.com
# Host End
```

更新时间：2024-06-14 08:58:45

## 修改本机的hosts文件
### 存放位置
* Windows 系统：C:\Windows\System32\drivers\etc\hosts
* Linux 系统：/etc/hosts
* Mac（苹果电脑）系统：/etc/hosts

### 生效
* Windows：在 CMD 窗口输入：ipconfig /flushdns
* Linux 命令：sudo rcnscd restart
* Mac 命令：sudo killall -HUP mDNSResponder