#!/usr/bin/expect

set command $argv;
set timeout -1

if { [file exists /home/app/.pocket/chains.json] } {
    
} else {
    puts [open /home/app/.pocket/chains.json w] {{"0de3141aec1e69aea9d45d9156269b81a3ab4ead314fbf45a8007063879e743b":{"addr":"0de3141aec1e69aea9d45d9156269b81a3ab4ead314fbf45a8007063879e743b","url":"http://my.ethchain.dnp.dappnode.eth:8545"}}}
}



if { $command eq "" } {
    
    if  { [info exists env(POCKET_CORE_SEEDS)] } {
        
        if { [info exists env(POCKET_CORE_KEY)] } {
            
            spawn pocket accounts import-raw $env(POCKET_CORE_KEY)
            sleep 1
            send -- "$env(POCKET_CORE_PASSPHRASE)\n"
            expect eof
            spawn pocket start --seeds $env(POCKET_CORE_SEEDS)
        } else {
            
            spawn pocket start --seeds $env(POCKET_CORE_SEEDS)
        } 
    } else {
        
        if { [info exists env(POCKET_CORE_KEY)] } {
            
            spawn pocket accounts import-raw $env(POCKET_CORE_KEY)
            sleep 1
            send -- "$env(POCKET_CORE_PASSPHRASE)\n"
            expect eof
            spawn pocket start 
        } else {
            
            spawn pocket start 
            sleep 1
            send -- "$env(POCKET_CORE_PASSPHRASE)\n"
        }
    } 
} elseif { $command ne "" } {
    
    if  { [info exists env(POCKET_CORE_KEY)] } {
        
        spawn pocket accounts import-raw $env(POCKET_CORE_KEY)
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
