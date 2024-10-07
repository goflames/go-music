# 使用官方的Node.js基础镜像作为基础镜像
FROM node:20
# 设置工作目录
WORKDIR /appManage
# 将package.json和package-lock.json复制到工作目录
COPY package*.json ./
# 安装依赖
RUN npm install
# 复制整个项目目录到容器中
COPY . .  
# 暴露端口
EXPOSE 8081
# 运行命令
CMD ["npm", "run", "serve"]

