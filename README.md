# sunlight 

> thanks to https://github.com/luyuhuang/subsocks

```
                            InitState
                                |
                                v
                      ————  CommandState ————————————————
                     /       |             |             \
                    /        |             |              \
           CmdConnect   CmdProxy         CmdUDP           CmdUDPoverTCP
               |           |               |                 |
               v           v               v                 v
      TCPConnectState  TCPActiveState    UDPConnectState  UDPActiveState
               |           |               |                 |
               v           v               v                 v
              TCPTransferState           UDPRelayState    UDPTransferState
               |           |               |                 |
               v           v               v                 v
                ——————————————  TerminateState ——————————————
```
