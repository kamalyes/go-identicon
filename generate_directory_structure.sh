#!/bin/bash
###
 # @Author: kamalyes 501893067@qq.com
 # @Date: 2024-10-18 17:53:55
 # @LastEditors: kamalyes 501893067@qq.com
 # @LastEditTime: 2024-10-18 17:55:59
 # @FilePath: \go-identicon\generate_directory_structure.sh
 # @Description: 
 # 
 # Copyright (c) 2024 by kamalyes, All Rights Reserved. 
### 

# 定义输出文件名
OUTPUT_FILE="DIRECTORY_STRUCTURE.md"

# 创建或清空输出文件
echo "# Git Project Directory Structure" > "$OUTPUT_FILE"
echo "" >> "$OUTPUT_FILE"

# 函数：递归遍历目录
function list_directory {
    local dir="$1"
    local prefix="$2"

    # 遍历目录中的所有文件和子目录 
    for entry in "$dir"/*; do
        if [ -d "$entry" ]; then
            # 如果是目录，写入 Markdown 格式 
            echo "${prefix}- **$(basename "$entry")**" >> "$OUTPUT_FILE"
            # 递归调用
            list_directory "$entry" "    $prefix"
        elif [ -f "$entry" ]; then
            # 如果是文件，写入 Markdown 格式 
            echo "${prefix}- $(basename "$entry")" >> "$OUTPUT_FILE"
        fi
    done
}

# 从当前目录开始遍历
list_directory "." ""

# 添加文件结束标记
echo "" >> "$OUTPUT_FILE"
echo "---" >> "$OUTPUT_FILE"
echo "Generated on $(date)" >> "$OUTPUT_FILE"

# 完成提示
echo "Directory structure has been written to $OUTPUT_FILE"
