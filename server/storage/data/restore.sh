#!/usr/bin/env bash
# leehom Chen clh021@gmail.com
cd "$( dirname "${BASH_SOURCE[0]}" )" || exit
file="test.db"
if [ -f "$file" ]; then
    echo "Deleting $file..."
    rm "$file"
fi
echo "Restore $file..."
sqlite3 test.db ".read sqlite.sql"
