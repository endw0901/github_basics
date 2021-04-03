# rebase
- 過去のbranch派生地点に巻き戻り、masterで発生した変更をプレイバックし、そしてfeatcher branchでの変更履歴を作る

```
// newBranchAでファイル修正
git checkout -b feature/newBranchA

// masterでファイル修正, commitが発生
git checkout master

// branchの派生が確認できる
git hist

// masterのcommitをfeature/newBranchAにrebaseする
git checkout feature/newBranchA
git rebase master
git hist

```

## conflict

```
// rebase取り消しする場合
git rebase --abort

// rebase conflict解決 => save
git mergetool
git add xxxxx
git rebase --continue

```
