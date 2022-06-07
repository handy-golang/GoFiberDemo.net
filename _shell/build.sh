#!/bin/bash
# 加载变量
source "./_shell/init.sh"
#############

echo " =========== 正在进行 server 端编译 =========== "

go mod tidy &&
  go build -o ${buildName}
echo " server 端编译 完成"

echo " =========== 开始进行 文件整理 =========== "

echo "清理并创建 dist 目录"
rm -rf ${outPutPath}
mkdir ${outPutPath} &&
  echo "移动 goRun 文件"
mv ${buildName} ${outPutPath} &&
  exit
