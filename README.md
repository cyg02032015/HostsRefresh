# FreeGithub
自动获取github相关网站的ip地址，解决github链接不畅通的问题，该项目的README会自动更新，大家可以直接复制使用

## hosts列表
```base
# Host Start
3.91.106.156 github.com
54.166.177.214 api.github.com
3.235.11.7 github.map.fastly.net
54.234.143.115 github.githubassets.com
3.221.159.117 github.global.ssl.fastly.net
54.91.93.203 raw.githubusercontent.com
3.80.246.56 camo.githubusercontent.com
3.93.143.193 avatars.githubusercontent.com
3.80.246.56 avatars0.githubusercontent.com
3.80.202.4 avatars1.githubusercontent.com
3.238.200.140 avatars2.githubusercontent.com
13.84.31.79 avatars3.githubusercontent.com
54.225.62.223 avatars4.githubusercontent.com
54.152.78.27 avatars5.githubusercontent.com
54.234.143.115 avatars6.githubusercontent.com
13.84.31.79 avatars7.githubusercontent.com
3.216.79.242 avatars8.githubusercontent.com
3.87.161.114 favicons.githubusercontent.com
54.152.78.27 developer.apple.com
54.227.219.183 devstreaming-cdn.apple.com
# Host End
```

更新时间：2021-10-15 08:46:55

## 修改本机的hosts文件
### 存放位置
* Windows 系统：C:\Windows\System32\drivers\etc\hosts
* Linux 系统：/etc/hosts
* Mac（苹果电脑）系统：/etc/hosts

### 生效
* Windows：在 CMD 窗口输入：ipconfig /flushdns
* Linux 命令：sudo rcnscd restart
* Mac 命令：sudo killall -HUP mDNSResponder