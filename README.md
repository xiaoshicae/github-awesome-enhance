## github awesome enhance


* 简介: awesome仓库无star数,蛋疼,故在原有基础上增加star数

* 项目结构:
    * github文件内容接口: https://api.github.com/repos/:username/:repo/contents/[file]
    * 保证项目足够轻量,不采用数据库及缓存,直接采用进程缓存
    * 后端(Go)代码部署在阿里云上,前端(Vue)页面放在github-page上
