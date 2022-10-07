package model

type receive_packet struct {
	Key                   string
	Cache_Control         string
	Content_Type          string
	Content_MD5           string
	x_amz_checksum_crc32  string
	x_amz_checksum_crc32c string
	x_amz_checksum_sha1   string
}

type send_packet struct {
	StatusCode HTTPStatusCode
}

type HTTPStatusCode int
