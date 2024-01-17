package hashing

const MD5 = "md5"
const SHA1 = "sha1"
const SHA256 = "sha256"
const SHA512 = "sha512"

func Build(algorithm string) Hasher {
	switch algorithm {
	case MD5:
		return MD5Hasher{}
	case SHA1:
		return SHA1Hasher{}
	case SHA256:
		return SHA256Hasher{}
	case SHA512:
		return SHA512Hasher{}
	default:
		return nil
	}
}
