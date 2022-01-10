# **ffmpeg golang binding**
# why a new goav fork?
- the google top goav result doesn't work anymore. 
- a recent working fork I tested is [here](https://github.com/asticode/goav)
- make it usable and compatile for FFmpeg 3.x or 4.x under ubuntu 18.04, 20.04
  
# Dependencies
goav depends upon several libraries from FFmpeg (version 3.3 or higher):

-   libavcodec
-   libavdevice
-   libavfilter
-   libavformat
-   libavutil
-   libswresample
-   libswscale

and a few other tools in general:

- pkg-config



## Mac OS X
On Mac OS X, Homebrew saves the day:
```
brew install ffmpeg pkg-config
```
## Ubuntu >= 18.04 LTS
On Ubuntu 18.04 LTS everything can come from the default sources:

```
### General dependencies
sudo apt-get install -y  pkg-config

### Library components
sudo apt-get install -y libavformat-dev libavcodec-dev libavdevice-dev libavutil-dev libswscale-dev libswresample-dev libavfilter-dev
```
## Ubuntu < 18.04 LTS
On older Ubuntu releases you will be unable to satisfy these requirements with the default package sources. We recommend compiling and installing FFmpeg from source. For FFmpeg:
```
sudo apt install \
    autoconf \
    automake \
    build-essential \
    cmake \
    libass-dev \
    libfreetype6-dev \
    libjpeg-dev \
    libtheora-dev \
    libtool \
    libvorbis-dev \
    libx264-dev \
    pkg-config \
    wget \
    yasm \
    zlib1g-dev

wget http://ffmpeg.org/releases/ffmpeg-3.2.tar.bz2
tar -xjf ffmpeg-3.2.tar.bz2
cd ffmpeg-3.2

./configure --disable-static --enable-shared --disable-doc
make
sudo make install
```

[See this script](https://gist.github.com/mkassner/1caa1b45c19521c884d5) for a very detailed installation of all dependencies.


## Windows
It is possible to build goav on Windows without Conda by installing FFmpeg yourself, e.g. from the shared and dev packages.

Unpack them somewhere (like C:\ffmpeg), and then tell goav where they are located.

# Original fork info
from here:
Golang binding for __FFmpeg 4.1.1__.   Forked from [here](https://github.com/amarburg/goav)