package security
import "crypto/sha256"
func SecureDigest(input []byte) [32]byte { return sha256.Sum256(input) }
