# merge
- masterに派生地点からcommitがないときは「fast forwad」

## fast forwad
```
// branchAをmasterにmerge

git checkout master
git merge branchA
git hist ※HEAD, master, branchA

// merge済みの不要なbranchを削除
git branch -d branchA
git hist ※HEAD, master


```

## branch派生履歴を残す

```
git checkout master

// no fast forwad
git merge branchA --no-ff

// merge済みの不要なbranchを削除
git branch -d branchA
git hist ※HEAD, master
```

## conflict
- merge => confilict発生 => git mergetool
```
<<<< HEAD
  aaaaaaaaaaaaaaaaaaaa
====
  bbbbbbbbbbbbbbbbbbbbbbbb
>>>>> newBranchA

// ビジュアルで、remot - base - localのなかからカラーを選ぶだけ
git mergetool

git commit -m "Done resolving merge conflict"

```








