ABOUT
======

Application that sets up a Lua environment. It really just combines golua and fsnotify for a workflow that I like when prototyping.

USAGE
=====

```shell
~/misc/lua/go $ ls
foo.lua  l.lua  p.lua
~/misc/lua/go $ luaenv 
hi from lua
there is no load order
its best to just dofile/require other modules, instead of relying on order
```

Now luaenv will loop waiting for file changes and eval updated files into the same environment.

OTHER
=====

Files are not loaded recursively, this is a feature which lets you control how and when something will be updated. It makes more sense to load other files from subdirectories, rather than forcing everything to reload.

LICENSE
======

MIT
