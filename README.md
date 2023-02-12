# 基于gin框架封装的单体项目目录结构

- 基于zap的日志打印
- controller接口自动打印接口访问日志，包括traceId、userId、输入输出参数、耗时等内容
- 全局的异常处理
- 统一的json数据返回
- 基于Viper的yaml项目配置
- 基于jwt的鉴权认证，通过中间件指定接口是否需要登录
- 基于validator的输入校验及自定义错误信息返回
- 基于sqlx的数据库访问层