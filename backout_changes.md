# backout_changes

- addを取り消したい
 
```
git add xxxx.txt

// staging area => working directoryに戻る
git reset HEAD xxxx.txt
git status

// working directoryで、変更内容を元に戻したい
git checkout -- <file名>
```
