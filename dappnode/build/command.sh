#!/usr/bin/expect
set command $argv;
set timeout -1
# Default testnet seeds
set seeds "610CF8A6E8CEFBADED845F1C1DC3B10A670BE26B@node1.testnet.pokt.network:26656, E6946760D9833F49DA39AAE9500537BEF6F33A7A@node2.testnet.pokt.network:26656, 7674A47CC977326F1DF6CB92C7B5A2AD36557EA2@node3.testnet.pokt.network:26656, C7B7B7665D20A7172D0C0AA58237E425F333560A@node4.testnet.pokt.network:26656, F6DC0B244C93232283CD1D8443363946D0A3D77A@node5.testnet.pokt.network:26656, 86209713BEFECA0807714BCDD5B79E81073FAF8F@node6.testnet.pokt.network:26656, 915A58AE437D2C2D6F35AC11B79F42972267700D@node7.testnet.pokt.network:26656, B3D86CD8AB4AA0CB9861CB795D8D154E685A94CF@node8.testnet.pokt.network:26656, 17CA63E4FF7535A40512C550DD0267E519CAFC1A@node9.testnet.pokt.network:26656, F99386C6D7CD42A486C63CCD80F5FBEA68759CD7@node10.testnet.pokt.network:26656"

if { [file exists /home/app/.pocket/chains.json] } {
    
} else {
    puts [open /home/app/.pocket/chains.json w] {{"0de3141aec1e69aea9d45d9156269b81a3ab4ead314fbf45a8007063879e743b":{"addr":"0de3141aec1e69aea9d45d9156269b81a3ab4ead314fbf45a8007063879e743b","url":"http://my.ethchain.dnp.dappnode.eth:8545"}}}
}

# checks if a command was passed
if { $command eq "" } {
    
    if  { [info exists env(POCKET_CORE_SEEDS)] } {
        
        if { [info exists env(POCKET_CORE_KEY)] } {
            # Import account using private key
            spawn pocket accounts import-raw $env(POCKET_CORE_KEY)
            sleep 1
            # Send private key decryption passphrase
            send -- "$env(POCKET_CORE_PASSPHRASE)\n"
            expect eof
            # Start pocket
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
            spawn pocket start --seeds $seeds
        } else {
            
            spawn pocket start --seeds $seeds
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
