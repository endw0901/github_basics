- branchA rebase branchBをする前に、branchBのcommitsがsquashされていればrebaseが楽

## 方法
- 事前にbranchAのbkupをcheckout -bで作成し、rebase後にブランチ比較branchA(rebase後)とbkupを比較することで、差がrebase分しかないことがわかる。
```sh
// swich to the branchB
git checkout branchB

// start an interactive rebase (最後の5コミットをsquashしたいとき)
git rebase -i HEAD~5

// 編集
pick abc123 initial commit
s abc124 second
s abc125 third
s abc126 fourth
s abc127 fifth

// save and exit :  Esc -> :wq
```


