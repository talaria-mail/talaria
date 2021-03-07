# Encryption

Talaria is built with privacy and security as a forethought. This document
describes the encryption systems in place. All users have a Curve25519 key-pair
generated on creation

```golang
type User struct {
    // Other fields elided
    

    ContentKey struct {
        Public []byte   // 32 byte Curve25519 public key
        Private []byte  // 32 byte encrypted Curve25519 private key
    }
}
```

All inbound message for this user are encrypted using the public key. The
public key encryption uses [ECIES][eceis] where `chacha20poly1305` is its
AEAD (we use the extended nonce version of chacha20 called xchacha20).

[ecies]: https://en.wikipedia.org/wiki/Integrated_Encryption_Scheme

The private key, which can be used to decrypt this message, is stored
encrypted. That private key can only be decrypted with a key generated from
the user password. When a user signs in we exchange their password for this
decryption key and recover the content private key. This means message are
decrypted only for authenticated users.

## Key-value store

Since all data storage is done with the key-value store, we can realize
data encryption by adding it as a middleware in the key-value store. We store
public keys in the context for `Put` calls to encrypt data,

```golang
ctx := context.Background()
ctx = encryption.WithKey(ctx, pubkey)

store := kv.NewStore()
store = encryption.KVMiddleware{
    Next: store,
}

err := store.Put(ctx, `key`, []byte(`value`)) // value will be encrypted
```

Conversely, we must add the private key for `Get` and `Delete` calls. We add
the pivate key for `Delete` calls because we want deletion to be authenicated.

```golang
ctx = encryption.WithKey(ctx, privkey)

value, err = store.Get(ctx, `key`)
```