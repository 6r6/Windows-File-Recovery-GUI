### Windows File Recovery GUI

A simple GUI for Windows File Recovery CLI Tool.

Golang + [GoVCL](https://github.com/ying32/govcl) + [Lazarus](http://www.lazarus-ide.org/) + [Res2go](https://github.com/ying32/govcl/blob/master/Tools/res2go).

---

#### How to build this repo with debug windows?

```
cd /master
go build
```

#### How to build for release?

```
# reduce the exe size and hide the cmd window:
go build -i -ldflags="-s -w -H windowsgui"
# then use UPXShell further reduce the exe size
# remember to put liblcl.dll to the same forder
```

#### Do you finished it?

Not yet.

#### Give me a demo

![demo1](https://github.com/6r6/Windows-File-Recovery-GUI/raw/master/LazarusProjectUI/DEMO1.png)