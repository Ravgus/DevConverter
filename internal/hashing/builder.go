package hashing

const (
	Md5    = "md5"
	Sha1   = "sha1"
	Sha156 = "sha256"
	Sha512 = "sha512"
)

func Build(algorithm string) Hasher {
	switch algorithm {
	case Md5:
		return MD5Hasher{}
	case Sha1:
		return SHA1Hasher{}
	case Sha156:
		return SHA256Hasher{}
	case Sha512:
		return SHA512Hasher{}
	default:
		return nil
	}
}
