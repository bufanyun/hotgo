echo -e "\n  \n" >> ./README.md

::git rm -r --cached  -f .
git init && git add -A

for /f "tokens=2 delims==" %%a in ('wmic path win32_operatingsystem get LocalDateTime /value') do (set t=%%a)
set Today=%t:~0,4%-%t:~4,2%-%t:~6,2% %t:~8,2%:%t:~10,2%

git commit -m  "%Today%"

git push git@gitee.com:bufanyun/hotgo-ui.git master