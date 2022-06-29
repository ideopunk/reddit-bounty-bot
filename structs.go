package main

type Campaign struct {
	Description             string  `json:"description"`               // // Campaign Description
	CurrencyId              int     `json:"currency_id"`               //	Campaign currency used ... string??
	Ends                    string  `json:"ends"`                      //	Campaign end date
	Blurb                   string  `json:"blurb"`                     //	Campaign blurb
	Hidden                  bool    `json:"hidden"`                    //	Whether Campaign is hidden
	EntryId                 int     `json:"entry_id"`                  //	Campaign ID  ... string??
	DurationTypeId          int     `json:"duration_type_id"`          //	Campaign type ... string??
	FundedAmount            float32 `json:"funded_amount"`             // Current amount funded
	PledgeProcessingStarted string  `json:"pledge_processing_started"` //	Pledge Processing start date
	Starts                  string  `json:"starts"`                    // Campaign start date
	Modified                string  `json:"modified"`                  // Campaign last modified date
	MaximumAllowedPledge    int     `json:"maximum_allowed_pledge"`    // Max amount allowed per pledge
	DisplayPriority         int     `json:"display_priority"`          // Order campaign will show up, only used when Featured ...string??
	Created                 string  `json:"created"`                   // Campaign creation date
	Disabled                bool    `json:"disabled"`                  // Whether Campaign is disabled
	RuntimeDays             int     `json:"runtime_days"`              // Days Campaign has been running
	ProfileTypeId           int     `json:"profile_type_id"`           // Whether Campaign is showing user profile or business profile. 1 = User Profile 2 = Business Profile
	Feature                 bool    `json:"featured"`                  // Whether Campaign is Featured
	EntryStatusID           int     `json:"entry_status_id"`           // Campaign status ID
	// 1 = Created / Being Edited
	// 2 = Approved / Running
	// 3 = Not Approved
	// 4 = Paused
	// 5 = Finished – Processing Pre Authorization
	// 6 = Finished – Accepted For Processing Capture
	// 7 = Finished – Declined For Processing Capture
	// 8 = Finished – Processing Capture
	// 9 = Finished – Capture Processing Complete
	// 10 = Sent For Review
	// 11 = Cancelled
	// 12 = Being Reviewed
	// 13 = Not Funded
	MaximumAllowedFundsRaised int    `json:"maximum_allowed_funds_raised"` // Max limit campaign can raise
	EverPublished             bool   `json:"ever_published"`               // If Campaign has ever been published
	Name                      string `json:"name"`                         // Campaign Name
	FundingGoal               int    `json:"funding_goal"`                 // Campaign Funding Goal
	RaiseModeID               int    `json:"raise_mode_id"`
}

type Post struct {
	Title string
	URL   string
}
