# Pure Go GOST cryptographic functions library.

## GOGOST v 5.10
Fork of the library GoGost. The fork was made due to a crash in panic of services under heavy load based on version 5.8 of the library.

### GOST is GOvernment STandard of Russian Federation (and Soviet Union).

* GOST 28147-89 (RFC 5830) block cipher with ECB, CNT (CTR), CFB, MAC
  CBC (RFC 4357) modes of operation
* various 28147-89-related S-boxes included
* GOST R 34.11-94 hash function (RFC 5831)
* GOST R 34.11-2012 Стрибог (Streebog) hash function (RFC 6986)
* GOST R 34.10-2001 (RFC 5832) public key signature function
* GOST R 34.10-2012 (RFC 7091) public key signature function
* various 34.10 curve parameters included
* Coordinates conversion from twisted Edwards to Weierstrass form and
  vice versa
* VKO GOST R 34.10-2001 key agreement function (RFC 4357)
* VKO GOST R 34.10-2012 key agreement function (RFC 7836)
* KDF_GOSTR3411_2012_256 KDF function (RFC 7836)
* GOST R 34.12-2015 128-bit block cipher Кузнечик (Kuznechik) (RFC 7801)
* GOST R 34.12-2015 64-bit block cipher Магма (Magma)
* GOST R 34.13-2015 padding methods
* MGM AEAD mode for 64 and 128 bit ciphers (RFC 9058)
* TLSTREE keyscheduling function
* ESPTREE/IKETREE (IKE* is the same as ESP*) keyscheduling function
* PRF_IPSEC_PRFPLUS_GOSTR3411_2012_{256,512} and generic prf+ functions
  (Р 50.1.111-2016 with IKEv2 RFC 7296)

Probably you could be interested in
Go's support of GOST TLS 1.3 (http://www.gostls13.cypherpunks.ru/).

### Known problems:

* intermediate calculation values are not zeroed
* 34.10 is not time constant and slow

GoGOST is free software: see the file COPYING for copying conditions.

GoGOST'es home page is: http://www.gogost.cypherpunks.ru/
You can read about GOST algorithms more: http://www.gost.cypherpunks.ru/

Please send questions, bug reports and patches to
http://lists.cypherpunks.ru/gost.html mailing list.
Announcements also go to this mailing list.

Development Git source code repository currently is located here:
http://www.git.cypherpunks.ru/?p=gogost.git;a=summary
