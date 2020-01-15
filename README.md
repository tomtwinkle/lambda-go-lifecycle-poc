# AWS lambda Go ライフサイクル 検証

## build (windows)
* install [mage](https://github.com/magefile/mage)

```
$ mage build
```

## lambda
* runtime: Go 1.x
* handler: main

## lifecycle
![sequence dialog](http://www.plantuml.com/plantuml/proxy?src=https://raw.githubusercontent.com/tomtwinkle/lambda-go-lifecycle-poc/master/lifecycle.puml)
