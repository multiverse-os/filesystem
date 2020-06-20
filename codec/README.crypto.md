# Filesystem Codec: Cryptography

## Approaches to authenticated encryption 

#### Encrypt-then-MAC (EtM)
The plaintext is first encrypted, then a MAC is produced based on the resulting
ciphertext. The ciphertext and its MAC are sent together.

  **EtM Protocols** IPSec


#### Encrypt-and-MAC (E&M)
A MAC is produced based on the plaintext, and the plaintext is encrypted without
the MAC. The plaintext's MAC and the ciphertext are sent together.

  **E&M Protocols** SSH

#### MAC-then-Encrypt (MtE)
A MAC is produced based on the plaintext, then the plaintext and MAC are
together encrypted to produce a ciphertext based on both. The ciphertext
(containing an encrypted MAC) is sent. 

  **MtE Protocols** SSL/TLS

Even though the MtE approach has not been proven to be strongly unforgeable in 
itself, the SSL/TLS implementation has been proven to be strongly unforgeable by
Krawczyk who showed that SSL/TLS was, in fact, secure because of the encoding
used alongside the MtE mechanism. 

#### What replaces RC4? 
Google has selected Poly1305 along with Bernstein's ChaCha20 symmetric cipher 
as a replacement for RC4 in TLS/SSL, which is used for Internet security. 
Google's initial implementation is used in securing https (TLS/SSL) traffic 
between the Chrome browser on Android phones and Google's websites. Use of 
ChaCha20/Poly1305 has been standardized in RFC 7905.

Shortly after Google's adoption for use in TLS, both ChaCha20 and Poly1305 
support was added to OpenSSH via the chacha20-poly1305@openssh.com 
authenticated encryption cipher. Subsequently, this made it possible for 
OpenSSH to remove its dependency on OpenSSL through a compile-time option.

