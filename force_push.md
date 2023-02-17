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
あるいは、 --force-with-lease[=<refname>]とともに補助的なオプションとして --force-if-includesを指定すると<br>
(つまり、リモート側のrefが正確にどのコミットを指していなければならないか、<br>
あるいはリモート側のどのrefが保護されているかを言わずに)「プッシュ」時に、<br>
バックグラウンドで暗黙的に更新されていたかもしれないリモート追跡refからの更新が<br>
ローカルで統合されているかどうかを検証してから強制更新が可能になります。
```

▼force pushで他の人のコミットを吹き飛ばすケース<br>
pullを忘れている<br>
pullしてからforce pushまでにコミットが入った<br>
=> --force-with-lease で防ぐ<br>

▼--force-with-lease で他の人のコミットを吹き飛ばすケース<br>
pullを忘れている<br>
pullしてからforce pushまでにコミットが入った<br>
かつfetchしてから --force-with-lease付きでforce pushした場合<br>
※ローカルのリポジトリが知っているリモートブランチは git fetch で origin と同期される<br>
=> 他の人のコミットを吹き飛ばす<br>
remotes/origin/xxxx  (ローカルのリポジトリが知っているリモートブランチ)<br>
origin/xxx (リモートブランチ)<br>
  <br>
=> `git push --force-with-lease --force-if-includes branch名` で防げる


### 例 コミット吹き飛ばした

```
main
A -> B
rc/202301ブランチ
A -> B -> C -> D
↓
A -> B -> C -> D -> D2
rc/202302ブランチ
ローカル
A -> B -> C -> D -> E -> F ->
これにD2追加してforce pushした
リモート
A -> B -> C -> D -> E -> F -> F2
※F2が消失
```
