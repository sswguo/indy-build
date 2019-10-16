package process

type StoreKey struct {
	PackageType string `json:"packageType"`
	StoreType   string `json:"type"`
	Name        string `json:"name"`
}

type TrackingKey struct {
	Id string `json:"id"`
}

type TrackingContent struct {
	TrackingKey
	Uploads   []TrackingContentEntry `json:"uploads"`
	Downloads []TrackingContentEntry `json:"downloads"`
}

type TrackingContentEntry struct {
	StoreKey
	AccessChannel string `json:"accessChannel"`
	Path          string `json:"path"`
	OriginUrl     string `json:"originUrl"`
	LocalUrl      string `json:"localUrl"`
	MD5           string `json:"md5"`
	SHA256        string `json:"sha256"`
	SHA1          string `json:"sha1"`
	Size          int64  `json:"size"`
}
