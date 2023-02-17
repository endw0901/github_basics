一番わかり易い記事 <br>
https://onk.hatenablog.jp/entry/2022/12/18/000000


参考記事
https://runebook.dev/ja/docs/git/git-push
https://qiita.com/ototonari/items/fd67801a2812f33ace1a

### force pushの必要性
rebase は履歴がかわるので必要になる
rebaseの必要性 -> branchをリリース日別につくると、直線のコミット履歴が必要

cherry-pickは、そもそものpick漏れが起きそうで、指定して戻したいとか個別運用以外は採用しない方がよさそう

### memo
git push --force-with-lease --force-if-includes branch名

```
あるいは、 --force-with-lease[=<refname>]とともに補助的なオプションとして --force-if-includesを指定すると(つまり、リモート側のrefが正確にどのコミットを指していなければならないか、あるいはリモート側のどのrefが保護されているかを言わずに)「プッシュ」時に、バックグラウンドで暗黙的に更新されていたかもしれないリモート追跡refからの更新がローカルで統合されているかどうかを検証してから強制更新が可能になります。
```

▼force pushで他の人のコミットを吹き飛ばすケース
pullを忘れている
pullしてからforce pushまでにコミットが入った
=> --force-with-lease で防ぐ

▼--force-with-lease で他の人のコミットを吹き飛ばすケース
pullを忘れている
pullしてからforce pushまでにコミットが入った
かつfetchしてから --force-with-lease付きでforce pushした場合
※ローカルのリポジトリが知っているリモートブランチは git fetch で origin と同期される
=> 他の人のコミットを吹き飛ばす
remotes/origin/xxxx  (ローカルのリポジトリが知っているリモートブランチ)
origin/xxx (リモートブランチ)
  
=> git push --force-with-lease --force-if-includes branch名 で防げる
