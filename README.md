# FreeGithub
自动获取github相关网站的ip地址，解决github链接不畅通的问题，该项目的README会自动更新，大家可以直接复制使用

## hosts列表
```base
# Host Start
140.82.112.4 github.com
140.82.112.4 api.github.com
151.101.1.6 github.map.fastly.net
35.90.149.9 github.githubassets.com
151.101.1.6 github.global.ssl.fastly.net
35.161.229.135 raw.githubusercontent.com
34.211.52.251 camo.githubusercontent.com
34.211.52.251 avatars.githubusercontent.com
54.188.65.2 avatars0.githubusercontent.com
52.35.24.184 avatars1.githubusercontent.com
54.184.156.147 avatars2.githubusercontent.com
35.161.229.135 avatars3.githubusercontent.com
54.186.223.156 avatars4.githubusercontent.com
34.223.4.240 avatars5.githubusercontent.com
34.211.52.251 avatars6.githubusercontent.com
54.218.164.178 avatars7.githubusercontent.com
35.89.133.33 avatars8.githubusercontent.com
35.90.106.175 favicons.githubusercontent.com
17.253.144.10 developer.apple.com
17.253.144.10 devstreaming-cdn.apple.com
# Host End
```

更新时间：2023-02-03 09:07:48

## 修改本机的hosts文件
### 存放位置
* Windows 系统：C:\Windows\System32\drivers\etc\hosts
* Linux 系统：/etc/hosts
* Mac（苹果电脑）系统：/etc/hosts

### 生效
* Windows：在 CMD 窗口输入：ipconfig /flushdns
* Linux 命令：sudo rcnscd restart
* Mac 命令：sudo killall -HUP mDNSResponder