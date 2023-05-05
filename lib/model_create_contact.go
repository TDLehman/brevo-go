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

type CreateContact struct {
	// Email address of the user. Mandatory if \"SMS\" field is not passed in \"attributes\" parameter. Mobile Number in \"SMS\" field should be passed with proper country code. For example {\"SMS\":\"+91xxxxxxxxxx\"} or {\"SMS\":\"0091xxxxxxxxxx\"}
	Email string `json:"email,omitempty"`
	// Pass your own Id to create a contact.
	ExtId string `json:"ext_id,omitempty"`
	// Pass the set of attributes and their values. The attribute's parameter should be passed in capital letter while creating a contact. These attributes must be present in your Brevo account. For eg. {\"FNAME\":\"Elly\", \"LNAME\":\"Roger\"}
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	// Set this field to blacklist the contact for emails (emailBlacklisted = true)
	EmailBlacklisted bool `json:"emailBlacklisted,omitempty"`
	// Set this field to blacklist the contact for SMS (smsBlacklisted = true)
	SmsBlacklisted bool `json:"smsBlacklisted,omitempty"`
	// Ids of the lists to add the contact to
	ListIds []int64 `json:"listIds,omitempty"`
	// Facilitate to update the existing contact in the same request (updateEnabled = true)
	UpdateEnabled bool `json:"updateEnabled,omitempty"`
	// transactional email forbidden sender for contact. Use only for email Contact ( only available if updateEnabled = true )
	SmtpBlacklistSender []string `json:"smtpBlacklistSender,omitempty"`
}
