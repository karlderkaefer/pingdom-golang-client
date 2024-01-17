package openapi

// object for workaround wrong API spec
type TMSCheckResponse struct {
    Check *CheckWithoutIDGET `json:"check"`
}

type CheckSimpleResponse struct {
	Check *CheckSimple `json:"check"`
}

// CheckWithID - new struct with an additional ID field for transaction checks
type CheckWithID struct {
    CheckWithoutID
    ID int64 `json:"id,omitempty"`
}

// CheckGetWithID - new struct with an additional ID field for transaction checks
type CheckGetWithID struct {
    CheckWithoutIDGET
    ID int64 `json:"id,omitempty"`
}