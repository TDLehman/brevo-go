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

// Sender details including id or email and name (optional). Only one of either Sender's email or Sender's ID shall be passed in one request at a time. For example `{\"name\":\"xyz\", \"email\":\"example@abc.com\"}` , `{\"name\":\"xyz\", \"id\":123}`
type UpdateSmtpTemplateSender struct {
	// Name of the sender
	Name string `json:"name,omitempty"`
	// Email of the sender
	Email string `json:"email,omitempty"`
	// Select the sender for the template on the basis of sender id. In order to select a sender with specific pool of IP’s, dedicated ip users shall pass id (instead of email).
	Id int64 `json:"id,omitempty"`
}
