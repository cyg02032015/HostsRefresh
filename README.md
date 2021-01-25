# FreeGithub
自动获取github相关网站的ip地址，解决github链接不畅通的问题，该项目的README会自动更新，大家可以直接复制使用

## hosts列表
```base
#Github Host Start
github.githubassets.com 
camo.githubusercontent.com 
github.map.fastly.net 
github.global.ssl.fastly.net 
github.com 
api.github.com 
raw.githubusercontent.com 
favicons.githubusercontent.com 
avatars8.githubusercontent.com 
avatars7.githubusercontent.com 
avatars6.githubusercontent.com 
avatars5.githubusercontent.com 
avatars4.githubusercontent.com 
avatars3.githubusercontent.com 
avatars2.githubusercontent.com 
avatars1.githubusercontent.com 
avatars0.githubusercontent.com 
avatars.githubusercontent.com 
# Github Host End
```

更新时间：2021-01-25 20:07:18

## 修改本机的hosts文件
### 存放位置
* Windows 系统：C:\Windows\System32\drivers\etc\hosts
* Linux 系统：/etc/hosts
* Mac（苹果电脑）系统：/etc/hosts

### 生效
* Windows：在 CMD 窗口输入：ipconfig /flushdns
* Linux 命令：sudo rcnscd restart
* Mac 命令：sudo killall -HUP mDNSResponder