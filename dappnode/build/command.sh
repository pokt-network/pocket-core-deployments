#!/usr/bin/expect

set command $argv;
set timeout -1

if { $command eq "" } {
    
    if  { [info exists env(POCKET_CORE_SEEDS)] } {
        
        if { [info exists env(POCKET_CORE_KEY)] } {
            
            spawn pocket-core accounts import-raw $env(POCKET_CORE_KEY)
            sleep 1
            send -- "$env(POCKET_CORE_PASSPHRASE)\n"
            expect eof
            spawn pocket-core start --seeds $env(POCKET_CORE_SEEDS)
        } else {
            
            spawn pocket-core start --seeds $env(POCKET_CORE_SEEDS)
        } 
    } else {
        
        if { [info exists env(POCKET_CORE_KEY)] } {
            
            spawn pocket-core accounts import-raw $env(POCKET_CORE_KEY)
            sleep 1
            send -- "$env(POCKET_CORE_PASSPHRASE)\n"
            expect eof
            spawn pocket-core start 
        } else {
            
            spawn pocket-core start 
            sleep 1
            send -- "$env(POCKET_CORE_PASSPHRASE)\n"
        }
    } 
} elseif { $command ne "" } {
    
    if  { [info exists env(POCKET_CORE_KEY)] } {
        
        spawn pocket-core accounts import-raw $env(POCKET_CORE_KEY)
        sleep 1
        send -- "$env(POCKET_CORE_PASSPHRASE)\n"
        expect eof
        spawn sh -c "$command"
    } else {
        
        spawn sh -c "$command"
    }

}

sleep 1
send -- "$env(POCKET_CORE_PASSPHRASE)\n"

expect eof

exit
