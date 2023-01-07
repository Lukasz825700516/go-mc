# How to run

First prepare deobfuscated `server.jar`
- Download server.jar
- Download server's deobfuscation maps
- Download [enigma](https://github.com/FabricMC/Enigma) (or other deobfuscation program)
- Deobfuscate server.jar/META-INF/version/<version>/<version>.jar, and apply downloaded mappings

Now run java with server libraries jars (list can be found in META-INF/libraries.list),
and save its output to some file (`packets`)
```
java -cp <server.jar>:<libraries.jars> Main.java > packets
```

Now prepare decompiled server sources and run
```
go run gen.go < packets
```
