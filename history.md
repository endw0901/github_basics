# history

```
// help
git help log

// log 最新～最初のcommitまで
git log
qで抜ける

// commit-id短縮表示
git log --abbrev-commit


git log --oneline --graph --decorate
qで抜ける

// commit-idの範囲指定 xxxx～bbbbのcommitまで
git log xxxx...bbbbb

// 特定のファイルのgit log
git log -- <file name>


// 詳細情報
git show <commit-id>

```

## おすすめ(alias登録)
- https://github.com/endw0901/github_basics/blob/master/alias.md
- git log --oneline --graph --decorateを上記でgit histにalias登録しておく
