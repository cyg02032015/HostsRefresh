# FreeGithub
自动获取github相关网站的ip地址，解决github链接不畅通的问题，该项目的README会自动更新，大家可以直接复制使用

## hosts列表
```base
# Host Start
140.82.112.4 github.com
140.82.113.6 api.github.com
185.199.108.133 github.map.fastly.net
185.199.108.154 github.githubassets.com
199.232.69.194 github.global.ssl.fastly.net
185.199.108.133 raw.githubusercontent.com
185.199.108.133 camo.githubusercontent.com
185.199.108.133 avatars.githubusercontent.com
185.199.108.133 avatars0.githubusercontent.com
185.199.108.133 avatars1.githubusercontent.com
185.199.108.133 avatars2.githubusercontent.com
185.199.108.133 avatars3.githubusercontent.com
185.199.108.133 avatars4.githubusercontent.com
185.199.108.133 avatars5.githubusercontent.com
185.199.108.133 avatars6.githubusercontent.com
185.199.108.133 avatars7.githubusercontent.com
185.199.108.133 avatars8.githubusercontent.com
185.199.108.133 favicons.githubusercontent.com
17.253.25.202 developer.apple.com
17.253.25.201 devstreaming-cdn.apple.com
# Host End
```

更新时间：2021-02-26 08:28:42

## 修改本机的hosts文件
### 存放位置
* Windows 系统：C:\Windows\System32\drivers\etc\hosts
* Linux 系统：/etc/hosts
* Mac（苹果电脑）系统：/etc/hosts

### 生效
* Windows：在 CMD 窗口输入：ipconfig /flushdns
* Linux 命令：sudo rcnscd restart
* Mac 命令：sudo killall -HUP mDNSResponder