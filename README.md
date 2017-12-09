# gin-gae
GAE上で動かすginフレームワークの設定諸々


# GAEのSDKをDL

https://cloud.google.com/appengine/downloads?hl=es#Google_App_Engine_SDK_for_Go

DLして解凍されたgoogle_appengineにpathを通す

# install module
```
go  get github.com/gin-gonic/gin 
go  get github.com/mjibson/goon 
go  get github.com/icza/session 
go  get google.golang.org/appengine
```


# JS CSS周り

```
brew install yarn
yarn install
```

下記でsrc --> assetsにdistされます

```
# React
yarn run watch:babel

# sass
yarn run watch:sass
```

# サーバー起動
```
goapp serve
```
