# FreeGithub
自动获取github相关网站的ip地址，解决github链接不畅通的问题，该项目的README会自动更新，大家可以直接复制使用

## hosts列表
```base
# Host Start
34.221.117.78 github.com
54.69.142.59 api.github.com
34.217.191.6 github.map.fastly.net
34.217.209.194 github.githubassets.com
34.221.214.185 github.global.ssl.fastly.net
52.10.45.44 raw.githubusercontent.com
35.87.130.190 camo.githubusercontent.com
34.217.209.194 avatars.githubusercontent.com
34.217.209.194 avatars0.githubusercontent.com
34.215.67.13 avatars1.githubusercontent.com
35.87.130.190 avatars2.githubusercontent.com
20.69.183.61 avatars3.githubusercontent.com
34.215.67.13 avatars4.githubusercontent.com
34.221.214.185 avatars5.githubusercontent.com
34.217.209.194 avatars6.githubusercontent.com
34.217.209.194 avatars7.githubusercontent.com
34.219.51.170 avatars8.githubusercontent.com
34.221.214.185 favicons.githubusercontent.com
34.213.224.131 developer.apple.com
35.87.130.190 devstreaming-cdn.apple.com
# Host End
```

更新时间：2022-02-11 08:52:44

## 修改本机的hosts文件
### 存放位置
* Windows 系统：C:\Windows\System32\drivers\etc\hosts
* Linux 系统：/etc/hosts
* Mac（苹果电脑）系统：/etc/hosts

### 生效
* Windows：在 CMD 窗口输入：ipconfig /flushdns
* Linux 命令：sudo rcnscd restart
* Mac 命令：sudo killall -HUP mDNSResponder