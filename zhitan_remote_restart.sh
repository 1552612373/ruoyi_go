#!/bin/bash

# 设置变量
APP_NAME="zhitan_exe"
REMOTE_PATH="/var/golang/zhitan"
LOCAL_BINARY="./$APP_NAME"

# Step 1: 构建 Linux 可执行文件
echo "[$(date)] 开始构建 $APP_NAME..."
GOOS=linux GOARCH=amd64 go build -o $APP_NAME
if [ $? -ne 0 ]; then
    echo "❌ 构建失败，请检查代码错误"
    exit 1
fi

# Step 2: SSH 到服务器，先停止服务（杀进程 + 清理）
echo "[$(date)] 正在连接服务器，停止旧服务并清理旧文件..."
ssh root@ali <<EOF
    # 停止服务：查找并杀死进程
    PID=\$(ps -ef | grep './$APP_NAME' | grep -v 'grep' | awk '{print \$2}')
    if [ -n "\$PID" ]; then
        echo "发现旧进程 PID: \$PID，正在终止..."
        kill \$PID 2>/dev/null || echo "kill 失败，尝试强制终止..."
        kill -9 \$PID 2>/dev/null
    fi

    # 删除旧文件（如果存在）
    rm -f $REMOTE_PATH/$APP_NAME
EOF

if [ $? -ne 0 ]; then
    echo "❌ 服务器端清理失败"
    exit 1
fi

# Step 3: 上传新文件到服务器
echo "[$(date)] 正在上传新文件到服务器..."
scp $LOCAL_BINARY root@ali:$REMOTE_PATH/
if [ $? -ne 0 ]; then
    echo "❌ 上传失败，请检查路径或权限"
    exit 1
fi


# Step 4: 启动新服务
echo "[$(date)] 正在启动服务..."
ssh root@ali "cd $REMOTE_PATH && chmod +x $APP_NAME && ./restart.sh"

# 检查服务是否运行
SERVICE_RUNNING=$(ssh root@ali "pgrep -x zhitan_exe")

if [ -n "$SERVICE_RUNNING" ]; then
    echo "✅ 部署成功！🎉"
else
    echo "❌ 部署失败：服务未启动，请检查日志"
    ssh root@ali "cat $REMOTE_PATH/nohup.out | tail -n 20"
    exit 1
fi
