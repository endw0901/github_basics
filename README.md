# readme

## tools

- P4 merge (P4 mergeのみ)
https://www.perforce.com/downloads/visual-merge-tool
path追加

```
git config --global merge.tool p4merge
git config --global mergetool.p4merge.path "C:/Program Files/Perforce/p4merge.exe"
git config --global mergetool.prompt false
git config --global diff.tool p4merge
```
