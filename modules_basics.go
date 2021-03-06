package main

import "github.com/endw0901/go_template" // http, httpsは不要
/*
▼moduleについて
・GOPATHに縛られず、どこからでもpackageをimportできる
・基本的には、github等のレポジトリを使う

▼go build
go get => go path directoryにimportされる

▼moduleの作り方
(1)githubアカウントにrepository作成：go_math

(2)ローカルにgo_math directory作成(mkdir go_math) => 配下にcalcとgeometryフォルダを作成

(3)同フォルダにgo.mod作成
$ go mod init github.com/endw0901/go_math
go: creating new go.mod: module github.com/endw0901/go_math

(4)git init
git initで.gitの隠しファイルを作成
Initialized empty Git repository in C:/Users/masaya/go/src/go_math/.git/
※この時点でリモート接続はない
 => ls -a で隠しファイルも含めて表示可能

(5)リモートレポジトリ作成し、originとネーミング
git remote add origin https://github.com/endw0901/go_math.git
※https必要、.git必要

(6)作成したリモートリポジトリの確認
git remote -v
origin  https://github.com/endw0901/go_math.git (fetch)
origin  https://github.com/endw0901/go_math.git (push)

(7)現在のディレクトリのファイルをすべてステージに上げる
git add -A

(8)コンフィグ変数を設定し、コミットするユーザーを設定
git config user.name "endw0901"
git config user.email "endw0901@gmail.com"

(9)コミット
git commit -m "first commit"

(10)push
git push -u -f origin master

(11)バージョニング
Semantic Versioning = ユニバーサルなバージョニング仕様
x.y.z
x:major
y:minor
z:patch version number
v1.2.3
major1,minor2,patch3

git tag -a v1.0.0 -m "initial release"
git push origin master --tags

=> github上にreleaseが追加される

▼Patch
git add -A

git commit -m "xxxxxxxxxxxxxxx"

git tag -a v1.0.1 -m "xxxxxxxxxxxx"

git push origin master --tags

myapp1のgo.modをvimでv1.0.0 -> v1.0.1に修正
=> go buildでダウンロード

=> lsでmyapps1.exeが生成されているのが分かる

▼Patch後の実行ファイルを実行
./myapp1.exe

▼Patch前のverに戻す
vim go.modでv1.0.0に修正

go get -u github.com/endw0901/go_math@v1.0.1
*/

func main() {

}
