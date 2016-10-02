# blinkstick-command-line
a simple go command line tool for the blinkstick device

## Requirements and Installation
Required the golang package led by boombuler
https://github.com/boombuler/led

* go get github.com/boombuler/led
* go build blinkstick.go

## Supported OS
Tested under Linux, but should also run under Windows and MacOS

## Usage
blinkstick -color [colorname]  
blinkstick -color [colorname] -lighttype blink  
  
Example:  
blinkstick -color blue -lighttype blink -duration 100 -times 10  

* color string  
color (for example, red, lime, white, etc. or off (default "black")  
* duration int  
time between blinks (default 300)  
* lighttype string  
lighttype (static, blink or pulse) (default "static")  
* steps int  
steps between pulse colors (default 15)  
* times int  
how many times it should blink/pulse (default 5)

