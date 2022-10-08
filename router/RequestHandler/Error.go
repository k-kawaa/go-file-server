package RequestHandler

const (
	XML_HEADER           = "<?xml version=\"1.0\" encoding=\"UTF-8\"?>"
	ERROR_HEADER_TOP     = "<Error>"
	ERROR_HEADER_END     = "</Error>"
	MESSAGE_HEADER_TOP   = "<Message>"
	MESSAGE_HEADER_END   = "</Message>"
	RESOURCE_HEADER_TOP  = "<Resource>"
	RESOURCE_HEADER_END  = "</Resource>"
	REQUESTID_HEADER_TOP = "<Request>"
	REQUESTID_HEADER_END = "</Request>"
)

const (
	ERROR_MESSAGE_NOT_ACL_SUPPOTED                 = "AccessControlListNotSupported"
	ERROR_MESSAGE_ACCESS_DENIED                    = "AccessDenied"
	ERROR_MESSAGE_ACCESSPOINT_ALREADY_OWNED_BY_YOU = "AccessPointAlreadyOwnedByYou"
	ERROR_MESSAGE_ACCOUNT_PROBLEM                  = "AccountProblem"
	ERROR_MESSAGE_ALLACCESS_DISABLED               = "AllAccessDisabled"
	ERROR_MESSAGE_AMBIGUOUS_GRANT_BY_EMAILADDRESS  = "AmbiguousGrantByEmailAddress"
	ERROR_MESSAGE_AUTHORIZATION_HEADER_MALFORMED   = "AuthorizationHeaderMalformed"
	ERROR_MESSAGE_BADDIGEST                        = "BadDigest"
	ERROR_MESSAGE_BUCKET_ALREADY_EXISTS            = "BucketAlreadyExists"
)

func ErrorHandle() {

}
