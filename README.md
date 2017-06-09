# GoDeeper

### Install

Download Qt5.8.0
```bash
  > wget https://download.qt.io/official_releases/qt/5.8/5.8.0/qt-opensource-linux-x64-android-5.8.0.run
```

Install Qt
```bash
  > chmod u+x qt-opensource-linux-x64-android-5.8.0.run
  > ./qt-opensource-linux-x64-android-5.8.0.run
```

Debian/Ubuntu
```bash
  > sudo apt-get -y install build-essential libgl1-mesa-dev
```

Install Go Bindings
```bash
  > go get -v github.com/therecipe/qt/cmd/...
```

Generate Bindings
```bash
  > $GOPATH/bin/qtsetup
```

### Build/test app

```bash
  > $GOPATH/bin/qtdeploy
  > $GOPATH/bin/qtdeploy test desktop src/GoDeeper
```

Build
```bash
  > $GOPATH/bin/qtdeploy build desktop src/GoDeeper
```
