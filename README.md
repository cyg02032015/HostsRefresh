# FreeGithub
自动获取github相关网站的ip地址，解决github链接不畅通的问题，该项目的README会自动更新，大家可以直接复制使用

## hosts列表
```base
# Host Start
20.36.179.54 github.com
20.36.179.54 api.github.com
20.36.179.54 github.map.fastly.net
20.36.179.54 github.githubassets.com
20.36.179.54 github.global.ssl.fastly.net
20.36.179.54 raw.githubusercontent.com
20.36.179.54 camo.githubusercontent.com
20.36.179.54 avatars.githubusercontent.com
20.36.179.54 avatars0.githubusercontent.com
20.36.179.54 avatars1.githubusercontent.com
20.36.179.54 avatars2.githubusercontent.com
20.36.179.54 avatars3.githubusercontent.com
20.36.179.54 avatars4.githubusercontent.com
20.36.179.54 avatars5.githubusercontent.com
20.36.179.54 avatars6.githubusercontent.com
20.36.179.54 avatars7.githubusercontent.com
20.36.179.54 avatars8.githubusercontent.com
20.36.179.54 favicons.githubusercontent.com
20.36.179.54 developer.apple.com
20.36.179.54 devstreaming-cdn.apple.com
# Host End
```

更新时间：2021-10-01 08:45:52

## 修改本机的hosts文件
### 存放位置
* Windows 系统：C:\Windows\System32\drivers\etc\hosts
* Linux 系统：/etc/hosts
* Mac（苹果电脑）系统：/etc/hosts

### 生效
* Windows：在 CMD 窗口输入：ipconfig /flushdns
* Linux 命令：sudo rcnscd restart
* Mac 命令：sudo killall -HUP mDNSResponder