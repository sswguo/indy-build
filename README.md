Indy-build: a simple client tool to run build against indy
======

Why use this cli
---

Sometimes it is not easy to set up a local testing environment for a full pnc process to do some testing. So use this cli we can simulate from some level to address some testing process for some local development results.

What has been done in this cli
---

It simulates a draft cutted process of pnc builds against a indy instance, including

* Create build group and build hosted.
* Create settings.xml with created build group
* Download the git repo and build the specified branch or tag with the settings.xml with a folo-tracking url with buildName (-n flag)
* Create pnc-builds hosted if not exists
* Seal the folo tracking
* Promote the build results to pnc-builds

Prerequisites to compile  
---

* golang (v1.11+)  
* go modules enabled(See [this](https://github.com/golang/go/wiki/Modules) for how to enable, and is enabled by default with golang v1.13+)
* Set up $GOPATH to a location

Compile & make
---

* Run "make clean build". And after build, you can find the compiled binary in "build" folder.  

Not to compile?
---

* You can directly download the binary [here](https://github.com/ligangty/indy-build/releases/download/indy-build-0.1/indy-build-0.1_linux_x86_64.tar.gz) for linux_x86 env  

Usage
---

* Before usage, you must have a running indy  
* A very simple usage example is as following

```bash
indy-build maven -i http://localhost:8080 -g https://www.github.com/Commonjava/weft.git -b master -n weft-master-build-1
```
