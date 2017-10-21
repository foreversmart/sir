package errors

var (
	BadRequest                  = Error{400, "BadRequest", "Request is incorrect"}
	AccountProblem              = Error{403, "AccountProblem", "There is a problem with your Qiniu account that prevents the operation from completing successfully. Please use Contact Us."}
	InvalidUserType             = Error{403, "InvalidUserType", "Your account has no right to access EVM service. Please use Contact Us."}
	ShortBlocked                = Error{409, "UserShortBlocked", "Your account is not allowed to login too much time in a short time"}
	AccessDenied                = Error{403, "AccessDenied", "Access Denied"}
	AuthFailure                 = Error{403, "AuthFailure", "The provided credentials could not be validated. You may not be authorized to carry out the request; for example, associating an Elastic IP address that is not yours, or trying to use an AMI for which you do not have permissions. Ensure that your account is authorized to use the Qiniu Evm service, that your credit card details are correct, and that you are using the correct access keys."}
	Blocked                     = Error{403, "Blocked", "Your account is currently blocked. Contact aws-verification@amazon.com if you have questions."}
	InvalidTaskConfig           = Error{400, "InvalidTaskConfig", "The TaskConfig is invalid"}
	InvalidParameter            = Error{400, "InvalidParameter", "A parameter specified in a request is not valid, is unsupported, or cannot be used. The returned message provides an explanation of the error value. For example, if you are launching an instance, you can't specify a security group and subnet that are in different VPCs."}
	InvalidParameterCombination = Error{400, "InvalidParameterCombination", "Indicates an incorrect combination of parameters, or a missing parameter. For example, trying to terminate an instance without specifying the instance ID."}
	InvalidParameterValue       = Error{400, "InvalidParameterValue", "A value specified in a parameter is not valid, is unsupported, or cannot be used. Ensure that you specify a resource by using its full ID. The returned message provides an explanation of the error value."}
	InvalidQueryParameter       = Error{400, "InvalidQueryParameter", "The QINIU_EVM query string is malformed or does not adhere to QINIU_EVM standards."}
	MalformedQueryString        = Error{400, "MalformedQueryString", "The query string contains a syntax error."}
	MissingAction               = Error{400, "MissingAction", "The request is missing an action or a required parameter."}
	MissingAuthenticationToken  = Error{400, "MissingAuthenticationToken", "The request must contain either a valid (registered) QINIU_EVM access key ID or X.509 certificate."}
	MissingSecurityHeader       = Error{400, "MissingSecurityHeader", "Your request is missing a required header."}
	MissingParameter            = Error{400, "MissingParameter", "The request is missing a required parameter. Ensure that you have supplied all the required parameters for the request; for example, the resource ID."}
	OptInRequired               = Error{400, "OptInRequired", "You are not authorized to use the requested service. Ensure that you have subscribed to the service you are trying to use. If you are new to QINIU_EVM, your account might take some time to be activated while your credit card details are being verified."}
	PendingVerification         = Error{400, "PendingVerification", "Your account is pending verification. Until the verification process is complete, you may not be able to carry out requests with this account. If you have questions, contact QINIU_EVM Support."}
	RequestExpired              = Error{400, "RequestExpired", "The request reached the service more than 15 minutes after the date stamp on the request or more than 15 minutes after the request expiration date (such as for pre-signed URLs), or the date stamp on the request is more than 15 minutes in the future. If you're using temporary security credentials, this error can also occur if the credentials have expired. For more information, see Temporary Security Credentials in the IAM User Guide."}
	UnauthorizedOperation       = Error{400, "UnauthorizedOperation", "You are not authorized to perform this operation. Check your IAM policies, and ensure that you are using the correct access keys. For more information, see Controlling Access. If the returned message is encoded, you can decode it using the DecodeAuthorizationMessage action. For more information, see DecodeAuthorizationMessage in the QINIU_EVM Security Token Service API Reference."}
	UnknownParameter            = Error{400, "UnknownParameter", "An unknown or unrecognized parameter was supplied. Requests that could cause this error include supplying a misspelled parameter or a parameter that is not supported for the specified API version."}
	UnsupportedProtocol         = Error{400, "UnsupportedProtocol", "SOAP has been deprecated and is no longer supported. For more information, see SOAP Requests."}
	ValidationError             = Error{400, "ValidationError", "The input fails to satisfy the constraints specified by an QINIU_EVM service."}
	BundlingInProgress          = Error{400, "BundlingInProgress", "The specified instance already has a bundling task in progress."}
	CannotDelete                = Error{400, "CannotDelete", "You cannot delete the 'default' security group in your VPC, but you can change its rules. For more information, see Qiniu Evm Security Groups."}
	IncorrectState              = Error{400, "IncorrectState", "The resource is in an incorrect state for the request. This error can occur if you are trying to attach a volume that is still being created. Ensure that the volume is in the 'available' state. If you are creating a snapshot, ensure that the previous request to create a snapshot on the same volume has completed. If you are deleting a virtual private gateway, ensure that it's detached from the VPC."}
	InvalidFormat               = Error{400, "InvalidFormat", "The specified disk format (for the instance or volume import) is not valid."}
	InvalidID                   = Error{400, "InvalidID", "The specified ID for the resource you are trying to tag is not valid. Ensure that you provide the full resource ID; for example, ami-2bb65342 for an AMI. If you're using the command line tools on a Windows system, you might need to use quotation marks for the key-value pair; for example, Name=TestTag."}
	InvalidInput                = Error{400, "InvalidInput", "An input parameter in the request is not valid; for example, if you specified an incorrect Reserved instance listing ID in the request."}
	InvalidOptionConflict       = Error{400, "InvalidOption.Conflict", "A VPN connection between the virtual private gateway and the customer gateway already exists."}
	InvalidPermissionDuplicate  = Error{400, "InvalidPermission.Duplicate", "The specified inbound or outbound rule already exists for that security group."}
	InvalidPermissionMalformed  = Error{400, "InvalidPermission.Malformed", "The specified security group rule is malformed. If you are specifying an IP address range, ensure that you use CIDR notation; for example, 203.0.113.0/24."}
	InvalidPermissionNotFound   = Error{400, "InvalidPermission.NotFound", "The specified rule does not exist in this security group."}
	InvalidRequest              = Error{400, "InvalidRequest", "The request is not valid. The returned message provides details about the nature of the error."}
	InvalidState                = Error{400, "InvalidState", "The specified resource is not in the correct state for the request; for example, if you are trying to enable monitoring on a recently terminated instance, or if you are trying to create a snapshot when a previous identical request has not yet completed."}
	InvalidZoneNotFound         = Error{400, "InvalidZone.NotFound", "The specified Availability Zone does not exist, or is not available for you to use. Use the DescribeAvailabilityZones request to list the Availability Zones that are currently available to you. Ensure that you have indicated the region for the Availability Zone in the request, if it's not in the default region. Specify the full name of the Availability Zone: for example, us-east-1a."}
	OperationNotPermitted       = Error{400, "OperationNotPermitted", "The specified operation is not allowed. This error can occur for a number of reasons; for example, you might be trying to terminate an instance that has termination protection enabled, or trying to detach the primary network interface (eth0) from an instance."}
	Unsupported                 = Error{400, "Unsupported", "The specified request is unsupported. For example, you might be trying to launch an instance in an Availability Zone that currently has constraints on that instance type. The returned message provides details of the unsupported request."}
	UnsupportedOperation        = Error{400, "UnsupportedOperation", "The specified request includes an unsupported operation. For example, you can't stop an instance that's instance store-backed. Or you might be trying to launch an instance type that is not supported by the specified AMI. The returned message provides details of the unsupported operation."}
	InternalFailure             = Error{400, "InternalFailure", "The request processing has failed because of an unknown error, exception or failure."}
	ServiceUnavailable          = Error{400, "ServiceUnavailable", "The request has failed due to a temporary failure of the server."}
	Unavailable                 = Error{400, "Unavailable", "The server is overloaded and can't handle the request."}
	SignatureDoesNotMatch       = Error{403, "SignatureDoesNotMatch", "The request signature we calculated does not match the signature you provided."}
	InternalError               = Error{500, "InternalError", "An internal error has occurred. Retry your request, but if the problem persists, contact us with details by posting a message on the QINIU_EVM forums."}
)

type Error struct {
	Code    int
	Name    string
	Message string
}

func IsEmptyError(err Error) bool {
	return err.Code == 0 && err.Name == "" && err.Message == ""
}

func (e Error) Error() string {
	return e.Message
}

func (e Error) String() string {
	return e.Name + ": " + e.Message
}

type ErrorResponse struct {
	statusCode int    `json:"-"`
	Code       string `json:"code"`
	Message    string `json:"message"`
	Resource   string `json:"resource,omitempty"`
	RequestId  string `json:"request_id"`
	Error      error  `json:"error"`
}

func NewErrorResponse(requestId, resource string, err Error) *ErrorResponse {
	return &ErrorResponse{
		statusCode: err.Code,
		Code:       err.Name,
		Message:    err.Message,
		Resource:   resource,
		RequestId:  requestId,
		Error:      err,
	}
}

// implements gogo.StatusCoder interface
func (er *ErrorResponse) StatusCode() int {
	return er.statusCode
}
