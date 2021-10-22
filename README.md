# FreeGithub
自动获取github相关网站的ip地址，解决github链接不畅通的问题，该项目的README会自动更新，大家可以直接复制使用

## hosts列表
```base
# Host Start
54.165.34.62 github.com
3.81.228.209 api.github.com
34.239.158.114 github.map.fastly.net
44.192.110.213 github.githubassets.com
52.170.18.109 github.global.ssl.fastly.net
3.237.40.230 raw.githubusercontent.com
3.85.119.6 camo.githubusercontent.com
44.192.7.96 avatars.githubusercontent.com
3.229.118.161 avatars0.githubusercontent.com
35.175.217.107 avatars1.githubusercontent.com
18.205.41.34 avatars2.githubusercontent.com
3.81.85.180 avatars3.githubusercontent.com
44.197.181.213 avatars4.githubusercontent.com
54.226.77.223 avatars5.githubusercontent.com
18.205.41.34 avatars6.githubusercontent.com
3.92.6.158 avatars7.githubusercontent.com
34.200.225.91 avatars8.githubusercontent.com
18.232.90.43 favicons.githubusercontent.com
54.236.37.103 developer.apple.com
18.234.101.236 devstreaming-cdn.apple.com
# Host End
```

更新时间：2021-10-22 08:50:36

## 修改本机的hosts文件
### 存放位置
* Windows 系统：C:\Windows\System32\drivers\etc\hosts
* Linux 系统：/etc/hosts
* Mac（苹果电脑）系统：/etc/hosts

### 生效
* Windows：在 CMD 窗口输入：ipconfig /flushdns
* Linux 命令：sudo rcnscd restart
* Mac 命令：sudo killall -HUP mDNSResponder