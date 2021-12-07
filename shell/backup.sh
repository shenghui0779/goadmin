mysqldump -u root -pwang970425 -h121.199.68.249 goadmin > goadmin.sql

git add -A
git commit -m "备份sql $(date '+%Y-%m-%d %H:%M:%S')" 
git push origin dev