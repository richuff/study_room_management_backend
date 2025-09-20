🤝 如何贡献
> 我们欢迎并感谢任何形式的贡献！无论是修复 typo、提交 issue、改进文档，还是提交新功能，你的帮助都会让项目变得更好。

✅ 开始之前
1. 请先阅读 行为准则（如适用）。
2. 确保你熟悉项目的 README 和目录结构。
3. 如果你要报告 bug 或提议新功能，请先搜索 Issues 是否已存在类似内容。

🐛 报告问题（Issue）
> 请使用我们的 Issue 模板 提交：

1.清晰的标题

2.复现步骤

3.期望行为 vs 实际行为

4.系统环境（如操作系统、浏览器、版本号等）

🔧 提交代码（Pull Request）
1. Fork 仓库
> 点击右上角的「Fork」按钮，将项目复制到你的账户下。
2. 克隆到本地
```bash
git clone https://github.com/你的用户名/你的项目.git
cd 你的项目
```
3. 创建新分支

```bash
git checkout -b fix/你的修复描述
git checkout -b feature/你的新功能名称
```
4. 提交更改
```bash
# 写清楚 commit message，建议使用 Conventional Commits 规范：
fix: 修复了导航栏在移动端显示异常的问题
feat: 新增用户头像上传功能
docs: 更新安装指南中的 Node 版本要求
```
5. 推送到你的 Fork
```bash
git push origin fix/你的修复描述
```
6. 提交 Pull Request
>    标题简洁明了，说明「做了什么」,正文可引用关联的 Issue，例如： Closes #123
