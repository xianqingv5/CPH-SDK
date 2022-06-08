#! /bin/bash
RED="\033[31m"
GREEN="\033[32m"
YELLOW="\033[33m"
END="\033[0m"

# 输入操作名
if [ "$1" = "" ];
then
    echo -e "请输入操作名"
fi

# 判断mod文件是否存在
if [ ! -f "go.mod" ];
then
    echo -e "go.mod文件不存在，请在项目目录中运行..."
    exit 1
fi

# 编译文件名 go_ssp
GoName=`head -n 1 go.mod| awk '{print $2}'`

# 停止运行
function stop() {
    # 打印
    print_color_string 'red' "stopping..."
    # 查找启动的go编译文件名称
    boot_id=`ps -ef | grep $GoName |grep -v grep|awk '{print $2}'`
    # 判断启动数量
    count=`ps -ef |grep $GoName |grep -v grep|wc -l`
    if [ $count != 0 ];
    then
        kill $boot_id
    fi
    print_color_string 'green' "stopped!"
}

# 构建
function build() {
    # 如果文件存在，删除
    if [ -f "./bin/$GoName" ];
    then
        rm "./bin/$GoName"
    fi
    print_color_string 'green' "building...."
    # 构建
    go build -o ./bin/$GoName
    # 获取上一个命令的退出状态
    if [ $? != 0 ];
    then
        print_color_string 'red' "fail to build"
    else
        print_color_string 'green' "success to build"
    fi
}

# 启动
function start() {
    print_color_string 'green' "start...."
    # 如果不存在，构建
    if [ ! -f $GoName ];
    then
        print_color_string 'yellow' "no execute file, building..."
        build
    fi

    #启动
    nohup ./bin/$GoName > worker.log 2>&1 &
    print_color_string 'green' "started!"
}

# 查看启动状态
function status() {
    count=`ps -ef | grep $GoName | grep -v grep | wc -l`
    if [ $count != 0 ];
    then
        print_color_string 'green' "$GoName is running...."
    else
        print_color_string 'red' "$GoName is not running...."
    fi
}

# 重启
function restart() {
    build
    stop
    sleep 1
    start
}

# 将打印加上颜色
print_color_string() {
    color=$1
    string=$2
    case $color in
        red)
                echo -e "$RED $string $END"
                ;;
        green)
                echo -e "$GREEN $string $END"
                ;;
        yellow)
                echo -e "$YELLOW $string $END"
                ;;
        *)
                echo $string
                ;;
    esac
}

case $1 in
    build)
        build;;
    stop)
        stop;;
    start)
        start;;
    restart)
        restart;;
    status)
        status;;
    *)
        echo "你输入了一个错误的命令"   
esac