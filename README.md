# clusterioPluginExample
Example plugin for [clusterio](github.com/Danielv123/factorioClusterio)

This is the most basic example of how a plugin can be created. You have a config.js file which specifies the binary location and a few things about the plugin. To install the plugin, place it in a folder in your /sharedPlugins/ folder (this repository is not compiled, do that first)

Multiplatform is up to the modder to do. The config.js file is ran in a full v8 nodejs thingy, so you can use process.platform to check for OS etc and change the specified binary based upon that. All plugins are expected to work on windows and ubuntuu.

When clusterio loads a server, it starts by firing up all the plugin processes. Their stdin will recieve the contents of the file specified in their config.js. Everything they print to stdout will be sent to the server console by rcon. Avoid messages larger than 4k characters as they break frequently.
