set DIR=%cd%
set TITLE="SongCF MillionHero: https://github.com/SongCF/ToolMillionHero"
set ADB="adb.exe"
set BIN="main.exe"
set FILE=/data/local/tmp/screenshot.png

%ADB% shell /system/bin/screencap -p %FILE%
%ADB% pull %FILE% ./
%BIN%
