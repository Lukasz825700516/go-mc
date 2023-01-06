# How to run

First prepare deobfuscated `server.jar`
- Download server.jar
- Download server's deobfuscation maps
- Download [enigma](https://github.com/FabricMC/Enigma) (or other deobfuscation program)
- Deobfuscate server.jar/META-INF/version/<version>/<version>.jar, and apply downloaded mappings

Now run java with server libraries jars (list can be found in META-INF/libraries.list)
```
java -cp <server.jar>:<libraries.jars> Main.java
```
