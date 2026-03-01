// Package algochat implements the AlgoChat protocol for encrypted messaging on Algorand.
//
// AlgoChat provides end-to-end encrypted messaging using X25519 key exchange,
// ChaCha20-Poly1305 AEAD encryption, and HKDF-SHA256 key derivation. Messages are
// transmitted via Algorand transaction notes.
//
// # Protocol Features
//
//   - Forward secrecy via ephemeral key pairs per message
//   - Bidirectional decryption (sender can decrypt their own messages)
//   - Key-publish payload filtering
//   - Reply threading with context
//   - Ed25519 signature verification for key announcements
//   - In-memory caching for messages, keys, and public key discovery
//   - Offline message queue with retry logic
//
// # Quick Start
//
//	// Create accounts from Algorand seeds
//	sender, _ := algochat.NewChatAccountFromSeed("SENDER_ADDR", senderSeed, senderEd25519Pub)
//	recipient, _ := algochat.NewChatAccountFromSeed("RECIPIENT_ADDR", recipientSeed, recipientEd25519Pub)
//
//	// Encrypt a message
//	envelope, _ := algochat.EncryptMessage("Hello!", sender.EncryptionPublicKey, recipient.EncryptionPublicKey)
//
//	// Encode for transaction note
//	noteBytes := algochat.EncodeEnvelope(envelope)
//
//	// Decode and decrypt
//	decoded, _ := algochat.DecodeEnvelope(noteBytes)
//	content, _ := algochat.DecryptMessage(decoded, recipient.EncryptionPrivateKey, recipient.EncryptionPublicKey)
//	// content.Text == "Hello!"
//
// # Wire Format (v1.0)
//
// Standard envelope: 126-byte header + variable ciphertext
//
//	Byte 0:      Version (0x01)
//	Byte 1:      Protocol ID (0x01 = standard, 0x02 = PSK)
//	Bytes 2-33:  Sender's X25519 public key (32 bytes)
//	Bytes 34-65: Ephemeral X25519 public key (32 bytes)
//	Bytes 66-77: Nonce (12 bytes)
//	Bytes 78-125: Encrypted sender key (48 bytes = 32 key + 16 auth tag)
//	Bytes 126+:  Ciphertext (variable, max 882 bytes plaintext)
package algochat
