# diff

## tools

- P4 merge (P4 mergeのみ)
https://www.perforce.com/downloads/visual-merge-tool
path追加

```
git config --global merge.tool p4merge
git config --global mergetool.p4merge.path "C:/Program Files/Perforce/p4merge.exe"
git config --global mergetool.prompt false
git config --global diff.tool p4merge
git config --global difftool.p4merge.path "C:/Program Files/Perforce/p4merge.exe"
git config --global difftool.prompt false
git config --global --list
```

## diff：差分比較

Working directory =(add)=> Staging area =(commit)=> Local Git repository =(push)=> Remote

```
// staging ⇔ working
git diff

// staging ⇔ working＜P4mergeを使う＞
git difftool


// Local Git Repository(last commit) ⇔ working
git diff HEAD

// Local Git Repository(last commit) ⇔ working＜P4mergeを使う＞
git difftool HEAD


// Local Git Repository(last commit) ⇔ staging
git diff --staged HEAD

// Local Git Repository(last commit) ⇔ staging＜P4mergeを使う＞
git difftool --staged HEAD


// Local Git Repository(last commit ) ⇔ Remote Git Repository
// masterがlocal、origin/masterがremote

git diff master origin/master
git difftool master origin/master

```

### commit指定diff

```
// 指定commit ⇔ last commit
git diff <commit-id> HEAD

// last commit ⇔ 1つ前のcommit
git diff HEAD HEAD^
git difftool HEAD HEAD^


// 指定commit(aaaa) ⇔ 指定commit(bbbb)
git diff aaaaa  bbbbbb
git difftool aaaaa bbbbb
ctrl + qで次のファイルdiffがみれる


// ファイル名指定でdiff
git diff -- <file name>
qで抜ける
```
