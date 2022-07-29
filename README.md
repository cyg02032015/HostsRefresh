# FreeGithub
自动获取github相关网站的ip地址，解决github链接不畅通的问题，该项目的README会自动更新，大家可以直接复制使用

## hosts列表
```base
# Host Start
140.82.113.4 github.com
140.82.113.4 api.github.com
151.101.1.6 github.map.fastly.net
3.227.20.223 github.githubassets.com
151.101.1.6 github.global.ssl.fastly.net
18.212.205.154 raw.githubusercontent.com
34.235.125.49 camo.githubusercontent.com
54.80.65.183 avatars.githubusercontent.com
34.206.64.37 avatars0.githubusercontent.com
3.231.56.28 avatars1.githubusercontent.com
184.73.57.190 avatars2.githubusercontent.com
35.170.198.157 avatars3.githubusercontent.com
3.93.74.194 avatars4.githubusercontent.com
54.159.35.43 avatars5.githubusercontent.com
3.238.5.205 avatars6.githubusercontent.com
34.239.168.133 avatars7.githubusercontent.com
34.239.168.133 avatars8.githubusercontent.com
3.94.7.160 favicons.githubusercontent.com
17.253.144.10 developer.apple.com
17.253.144.10 devstreaming-cdn.apple.com
# Host End
```

更新时间：2022-07-29 09:21:48

## 修改本机的hosts文件
### 存放位置
* Windows 系统：C:\Windows\System32\drivers\etc\hosts
* Linux 系统：/etc/hosts
* Mac（苹果电脑）系统：/etc/hosts

### 生效
* Windows：在 CMD 窗口输入：ipconfig /flushdns
* Linux 命令：sudo rcnscd restart
* Mac 命令：sudo killall -HUP mDNSResponder