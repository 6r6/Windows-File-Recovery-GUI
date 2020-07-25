### Windows File Recovery GUI

---

Golang + VCL GUI + Lazarus + Res2go

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
```

#### Do you finished it?

Not yet.

#### Give me a demo

