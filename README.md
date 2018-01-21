# dastard

A small tool to access the AccountsService over dbus.

It used to be easier for a user to change options such as keyboard layouts and background
images, by editing a config file.  But now these settings are managed by the AccountsService,
and the config file is not writable by the user. Instead, the poor user must now call the
AccountsService api.

The dbus-send tool can be used to change the BackgroundFile, which is just a string.  But it
can't be used to change the InputSources because that property is a much more (needlessly?)
complex data structure than a string.

So, this small program will let you change the BackgroundFile and InputSources properties.

For InputSources, I have assumed that the single key in each of the maps is "xkb", as I have not
seen any other possible values, and can't find any documentation.

