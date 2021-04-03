# reset

```
// 1つ前のコミットに戻す
git reset HEAD^1

// 特定のhead番号に戻る
git hist
git reset HEAD@{2}

```

## reflog
- git resetなど、過去の作業履歴を一覧化できる

```
git reflog


// 特定の履歴箇所に戻る
git hist
git reset <commit番号>


```
