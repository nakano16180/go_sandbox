## Go sandbox

### やったこと
- 以下の記事を読んで、Windows環境でdevcontainerをさわってみた
- 記事中にあるリンクより、まずはマイクロソフトの公式を試した
  - [VSCode devcontainerでローカルを汚さずに、快適なGo言語の開発環境を整える](https://zenn.dev/bun913/articles/f0a6c6177a4716)
- Try a Dev Container Sample... がうまくいかなかったが以下の記事によるとWindows側のDocker Desktopを起動しておく必要があったそう
  - [Docker Desktop for Windows を使用しない Dev Container 環境を構築する #VSCode - Qiita](https://qiita.com/ain1084/items/6cb6d82852c91416ec0e)
- もう一つ試したのは```>Dev Containers Open Folder in Container```
- このリポジトリはまだ何も入れていないディレクトリで上記コマンドを実行したもの。
  - 環境はGo1.22
  - 追加のツールインストールはしなかったが、dependabotだけなんか入ったっぽい
  - [Dev Container on WSL2で開発環境構築](https://zenn.dev/ykdev/articles/14a108290e24f9)
