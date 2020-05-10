#!/usr/bin/expect
set command $argv;
set timeout -1
# Default testnet seeds


# checks if a command was passed
if { $command eq "" } {
    
    if { [info exists env(POCKET_CORE_KEY)] } {
        spawn sh -c "cp /tmp/*.json /home/app/.pocket/config/"

        if { [file exists /home/app/.pocket/chains.json] } {
        } else {
            puts [open /home/app/.pocket/chains.json w] {[{"id":"0021","url":"http://my.ethchain.dnp.dappnode.eth:8545"}]}
        }
        spawn pocket accounts import-raw $env(POCKET_CORE_KEY)
        sleep 1
        send -- "$env(POCKET_CORE_PASSPHRASE)\n"
        expect eof
        spawn sh -c "pocket accounts set-validator `pocket accounts list | cut -d' ' -f2- `"
        sleep 1
        send -- "$env(POCKET_CORE_PASSPHRASE)\n"
        expect eof
        spawn pocket start
    } else {
        spawn pocket start 
    } 

} elseif { $command ne "" } {
    
    if  { [info exists env(POCKET_CORE_KEY)] } {
        spawn pocket accounts import-raw $env(POCKET_CORE_KEY)
        sleep 1
        send -- "$env(POCKET_CORE_PASSPHRASE)\n"
        expect eof
        spawn sh -c "pocket accounts set-validator `pocket accounts list | cut -d' ' -f2- `"
        sleep 1
        send -- "$env(POCKET_CORE_PASSPHRASE)\n"
        expect eof
        # executes command passed by user 
        spawn sh -c "$command"
    } else {
        
        spawn sh -c "$command"
    }
}

sleep 1
send -- "$env(POCKET_CORE_PASSPHRASE)\n"

expect eof

exit
