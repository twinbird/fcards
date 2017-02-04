# fcards
コマンドライン単語帳

##### インストール

```go
go install https://github.com/twinbird/fcards
```

もしくはReleaseからダウンロードして, 環境変数にPATHを設定してください.


###### 使い方

単語帳のTSVファイルを用意します.
左に英語, 右に日本語の形式で書くのが推奨です.

```sh
$ cat fcards.tsv
Add A to B      AをBに加える
Remove A from B BからAを取り除く
Move A from B to C      AをBからCへ動かす
Replace A with B        AをBに差し替える
Make A B        AをBにする
Change A to B   AをBに変更する
Update A to B   AをBに更新する
Ensure A        Aであることを確実にする
Use A   Aを使う
Fix A   Aを直す
```

コマンドを実行すると, カレントにある「fcards.tsv」をロードして起動します.

```sh
$ ./fcards.exe
Type ':q' for save and quit.

1 / 10
Add A to B
>

```

終了は「:q」とタイプしてEnterを押してください.  
現在のページを保存して終了します.

###### ヘルプ

```sh
$ ./fcards.exe -h
コマンドラインの単語帳
TSVファイルは「表面」,「裏面」の2列で作成してください
  -f string
        単語帳のファイル名を指定する (default "./fcards.tsv")
  -r    表面と裏面を反転して始める
  -reset
        単語帳を最初から始める
```
