## gui_demo


## required
- github.com/andlabs/ui
- go1.19
- cmake



### prerequisites

**`macOS`**
- [osxcross](https://github.com/tpoechtrager/osxcross)
- macOS SDK
- Xcode.dmg


First, we need to download the macOS SDK, which is bundled with Xcode.
Download XCode.dmg from the Apple download site at <https://developer.apple.co>
m/download/more/?name=Xcode%207.3 (7.3.1 is recommended for osxcross). Next,
install the osxcross tools from github.com/tpoechtrager/osxcross (full installation
details are available at that URL or in the Appendix). Completing the
installation will have extracted the macOS SDK and created the compilation
toolchain that will build against these installed APIs.

```sh
CC=o64-clang CXX=o64-clang++
```



**`Linux`**
- gtk3 or gtk3-devel or libgtk-3-dev
- gcc or clang


#### rebuilding libui library(optional)

```sh
git clone git@github.com:andlabs/libui.git -q
cd libui
git checkout alpha4.1 -q
mkdir build
cd build
cmake -DBUILD_SHARED_LIBS=OFF ..
make

cp ./out/libui.a /path/to/your/folder
```

**docker container network**
```sh
vim /etc/docker/daemon.json
# add dns list
{"dns": ["8.8.8.8"]}
```

**docker container goproxy**
```sh
export GOPROXY="https://goproxy.cn,direct"
```



**see also**

- cross compiler: [gox](https://github.com/mitchellh/gox)
