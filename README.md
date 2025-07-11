Inspired from [explainshell.com](explainshell.com) but here I explain UDS message.

# MVP Scope

1. Custom SQLite database
2. Only UDSonCAN
3. Only support Single CAN message with structures:
 - [SID] [SUB-FUCNTION] [PARAMS]
 - [SID] [PARAMS]
 - [SID]
 - [7F] [SID] [NRC]
4. Support as both cli tool and web interface
5. Test case
6. Package CLI binary
7. Dockerize web

# Architecture

## Version 0.1

- Backend: `gin`
- Frontend: `html/template`
- Database: homegrown `SQLite`
- UDS Parser: homegrown tool
- CLI Syntax:
  - default: `explainuds "SID SUB-FUNCTION PARAMS"`
  - example: `explainuds 0x22 0x01 0x01`

# UDS Parser

