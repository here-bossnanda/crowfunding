package campaign

import "time"

type Campaign struct {
	ID  						 int
	userId 					 int
	ShortDescription string
	Description 		 string
	Perks 					 string
	BackerCount 		 int
	GoalAmount 			 int
	CurrentAmount 	 int
	Slug 						 string
	CampaignImages 	 []CampaignImage
	CreatedAt 			 time.Time
	UpdatedAt 			 time.Time
}

type CampaignImage struct {
	ID 					int
	CampaignId  int
	FileName 		string
	IsThumbnail int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}