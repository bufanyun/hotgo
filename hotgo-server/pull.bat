::强制和远程同步，不保留本地修改
::git fetch --all
::git reset --hard origin/master
::git pull


::保存本地修改，拉取远程后恢复本地修改
::git stash
::git pull origin master
::git stash pop
::git stash list
::git stash clear

::拉取远程
git pull origin master