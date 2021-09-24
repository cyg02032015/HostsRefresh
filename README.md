# FreeGithub
自动获取github相关网站的ip地址，解决github链接不畅通的问题，该项目的README会自动更新，大家可以直接复制使用

## hosts列表
```base
# Host Start
52.177.192.221 github.com
52.177.192.221 api.github.com
52.177.192.221 github.map.fastly.net
52.177.192.221 github.githubassets.com
52.177.192.221 github.global.ssl.fastly.net
52.177.192.221 raw.githubusercontent.com
52.177.192.221 camo.githubusercontent.com
52.177.192.221 avatars.githubusercontent.com
52.177.192.221 avatars0.githubusercontent.com
52.177.192.221 avatars1.githubusercontent.com
52.177.192.221 avatars2.githubusercontent.com
52.177.192.221 avatars3.githubusercontent.com
52.177.192.221 avatars4.githubusercontent.com
52.177.192.221 avatars5.githubusercontent.com
52.177.192.221 avatars6.githubusercontent.com
52.177.192.221 avatars7.githubusercontent.com
52.177.192.221 avatars8.githubusercontent.com
52.177.192.221 favicons.githubusercontent.com
52.177.192.221 developer.apple.com
52.177.192.221 devstreaming-cdn.apple.com
# Host End
```

更新时间：2021-09-24 08:41:24

## 修改本机的hosts文件
### 存放位置
* Windows 系统：C:\Windows\System32\drivers\etc\hosts
* Linux 系统：/etc/hosts
* Mac（苹果电脑）系统：/etc/hosts

### 生效
* Windows：在 CMD 窗口输入：ipconfig /flushdns
* Linux 命令：sudo rcnscd restart
* Mac 命令：sudo killall -HUP mDNSResponder