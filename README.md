# 基于 Go 语言的 Web 开发脚手架

## 总体架构
此脚手架总体使用 `CLD` 分层架构进行设计。
- `C` 为 `Controller` 层，服务的入口，负责处理路由、参数校验、请求转发。
- `L` 为 `Logic` 层，逻辑（服务）层，负责处理业务逻辑。
- `D` 为 `Dao` 层，负责数据与存储相关功能。

## 相关技术：
- MySQL
- Redis
- Gin
- Zap
- Viper

## 使用方法
```bash
git clone git@github.com:xyb7910/web_go.git

cd web_go
```
## 贡献
欢迎所有感兴趣的开发者为我们的项目做出贡献！你可以在 `issue` 区提出问题或建议，也可以直接提交 `pull request` 添加新功能或者修复 `bug` 。在提交 `PR` 之前，请确保你的代码符合了我们的编码规范。

## 许可证
本项目遵循 **MIT License** 授权协议，查阅 [LICENSE文件](https://docs.github.com/en/repositories/managing-your-repositorys-settings-and-features/customizing-your-repository/licensing-a-repository) 了解详细信息。
## 致谢
对使用或者参与我们项目的每一位开发者表示感谢！

