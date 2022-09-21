# ChatPuppy Api Document

## Auth
- Add headers `Authorization Bearer {token}`
- Add headers `X-Member-Address ${address}`

## Signature
1. Just post request
2. Encrypt raw post body string to `MD5`, getting string `bodyMD5`
3. Concatenate `timestamp` (from 1970, 10 digits) string end of step 2, getting string `s = bodyMD5 + timestamp`
4. Sign `s` getting `signature` string
5. Add headers `X-Signature ${signature}`
6. Add headers `X-Timestamp ${timestamp}`

## WebSocket
> `/message`

### Operates
- 1: register
- 2: set group key
- 3: new group message
