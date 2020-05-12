# Using

**- You need an User API Key for your HUE Bridge**

**- You need the ID of your lightbulb**


## create config

File ~/.hue/config_go

    {"ip":"BRIDGEIP","key":"APIKEY"}

## add module to your polybar setting

    [module/hue]
    type = custom/script
    exec = ~/.config/polybar/scripts/hue 1
    tail = true
    label = "HUE: %output%"

    click-left = kill -USR1 %pid%

Note: Set in the exec command your bulb ID. Mine is 1

don't forget to add "hue" to your modules-[left|center|right] entry

## Build
(install all dependencies (shown in main.go) )

    make build

## Install
    make install

Restart your Polybar - have fun



