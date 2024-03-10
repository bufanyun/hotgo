#!/usr/bin/env bash
# leehom Chen <clh021@gmail.com>

# 确保脚本所在目录作为当前工作目录
script_dir=$(dirname "${BASH_SOURCE[0]}")
cd "$script_dir" || exit

# 检查输入文件是否存在
# if [[ ! -f "$1" ]]; then
#     echo "Please provide the MySQL dump file as an argument."
#     exit 1
# fi
pwd

# 定义输入输出文件路径
input_file="sqlite.sql"
output_file="converted_sqlite.sql"

# 解析并转换SQL脚本
while IFS= read -r line; do

  # 删除所有以 'SET '（注意空格）开头的行
  if [[ "$line" =~ ^[[:space:]]*SET ]]; then
    continue
  fi

  # 删除以 '/*!' 开头的行
  if [[ "$line" =~ ^/*/ ]]; then
    continue
  fi

  # 删除以 '--' 开头的行（注释行）
  if [[ "$line" =~ ^-- ]]; then
    continue
  fi

  # 删除空行
  if [[ -z "$line" ]]; then
    continue
  fi
  # 替换 AUTO_INCREMENT 为 INTEGER PRIMARY KEY AUTOINCREMENT
  # line=$(echo "$line" | sed -E 's/(CREATE TABLE.*\()(.*,\) /\\1\\2 INTEGER PRIMARY KEY AUTOINCREMENT,/g' -e 's/AUTO_INCREMENT//g')

  # 删除引擎设置（SQLite不需要）
  # 移除 'ENGINE=InnoDB'
  line="${line/ ENGINE=InnoDB/}"

  # 处理 'COMMENT ' 字符串
  if [[ "$line" =~ \s*COMMENT\s*=\s*'[^']* ]]; then
    # 获取COMMENT内容及其后的字符
    comment_and_suffix="${BASH_REMATCH[1]}"
    suffix="${comment_and_suffix##*}"
    comment="${comment_and_suffix%}*}"
    line="${line/$comment_and_suffix/,-- $comment$suffix}"
  fi

  # 处理 TIMESTAMP DEFAULT CURRENT_TIMESTAMP
  # line=$(echo "$line" | sed 's/TIMESTAMP DEFAULT CURRENT_TIMESTAMP/TIMESTAMP NULL/g')

  # 输出转换后的行
  echo "$line"
done <"$input_file" >"$output_file"

# 回到原始目录
cd - || exit

# 显示转换后的文件路径
printf "Converted file saved at: %s\n" "$output_file"
