# stash

```
// local, stagingをcleanにする
git stash
git status

// stashを戻す
git stash apply


// 最新のstashを削除
git stash list
git stash drop

// 最新のstashを適用し、stashを削除する
git stash pop
```


## 複数stash

```
git stash save "aaaa"
git stash save "bbbb"

git stash list

// 2つめのstashを表示(0から始まるindex)
git stash show stash@{1}

// 複数のstashから選んで適用
// 2つめのstashを適用する
git stash apply stash@{1}

// stash@{0}, 1, 2が 0,1にindex振られなおす
git stash drop stash@{1}

// 複数のstashの一括削除
git stash clear
```
