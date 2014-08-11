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
its best to just require other modules, instead of relying on order
```

Now luaenv will loop waiting for file changes and eval updated files into the same environment.

LICENSE
======

MIT
