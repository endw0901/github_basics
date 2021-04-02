# alias

- 短縮コマンド
```
git log --all --graph --decorate --oneline

=> git histでできるようにする
```

```
git hist
git config --global alias.hist "log --all --graph --decorate --oneline"

// aliasで実行
git hist

// alias登録確認
code ~/.gitconfig
```
