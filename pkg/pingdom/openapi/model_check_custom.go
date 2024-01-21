package openapi

func (check *CheckWithoutID) AsPut() *CheckWithoutIDPUT{
	return &CheckWithoutIDPUT{
        Active: check.Active,
        ContactIds: check.ContactIds,
        CustomMessage: check.CustomMessage,
        IntegrationIds: convertInt64ArrayToInt32Array(check.IntegrationIds),
        Interval: check.Interval,
        Metadata: check.Metadata,
        Name: &check.Name,
        Region: check.Region,
        SeverityLevel: check.SeverityLevel,
        SendNotificationWhenDown: check.SendNotificationWhenDown,
        Steps: check.Steps,
        Tags: check.Tags,
        TeamIds: check.TeamIds,
    }
}