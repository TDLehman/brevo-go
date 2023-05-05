/*
 * Brevo API
 *
 * Brevo provide a RESTFul API that can be used with any languages. With this API, you will be able to :   - Manage your campaigns and get the statistics   - Manage your contacts   - Send transactional Emails and SMS   - and much more...  You can download our wrappers at https://github.com/orgs/brevo  **Possible responses**   | Code | Message |   | :-------------: | ------------- |   | 200  | OK. Successful Request  |   | 201  | OK. Successful Creation |   | 202  | OK. Request accepted |   | 204  | OK. Successful Update/Deletion  |   | 400  | Error. Bad Request  |   | 401  | Error. Authentication Needed  |   | 402  | Error. Not enough credit, plan upgrade needed  |   | 403  | Error. Permission denied  |   | 404  | Error. Object does not exist |   | 405  | Error. Method not allowed  |   | 406  | Error. Not Acceptable  |
 *
 * API version: 3.0.0
 * Contact: contact@brevo.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type GetExtendedList struct {
	// ID of the list
	Id int64 `json:"id"`
	// Name of the list
	Name string `json:"name"`
	// Number of blacklisted contacts in the list
	TotalBlacklisted int64 `json:"totalBlacklisted"`
	// Number of contacts in the list
	TotalSubscribers int64 `json:"totalSubscribers"`
	// Number of unique contacts in the list
	UniqueSubscribers int64 `json:"uniqueSubscribers"`
	// ID of the folder
	FolderId int64 `json:"folderId"`
	// Creation UTC date-time of the list (YYYY-MM-DDTHH:mm:ss.SSSZ)
	CreatedAt     string                         `json:"createdAt"`
	CampaignStats []GetExtendedListCampaignStats `json:"campaignStats,omitempty"`
	// Status telling if the list is dynamic or not (true=dynamic, false=not dynamic)
	DynamicList bool `json:"dynamicList,omitempty"`
}
