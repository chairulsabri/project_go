package campaign

import "strings"

type CampaignFormatter struct {
	ID               int    `json:"id"`
	UserID           int    `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageURL         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
	Slug             string `json:"slug"`
}

func FormatCampaign(campaign Campaign) CampaignFormatter {
	campaignFormatter := CampaignFormatter{}
	campaignFormatter.ID = campaign.ID
	campaignFormatter.UserID = campaign.UserID
	campaignFormatter.Name = campaign.Name
	campaignFormatter.ShortDescription = campaign.Description
	// campaignFormatter.ImageURL = campaign.ImageURL
	campaignFormatter.GoalAmount = campaign.GoalAmount
	campaignFormatter.CurrentAmount = campaign.CurrentAmount
	campaignFormatter.Slug = campaign.Slug
	campaignFormatter.ImageURL = ""
	if len(campaign.CampaignImages) > 0 {
		campaignFormatter.ImageURL = campaign.CampaignImages[0].FileName
	}

	return campaignFormatter
}

func FormatCampaings(campaigns []Campaign) []CampaignFormatter {
	// if len(campaigns) == 0 {
	// 	return []CampaignFormatter{}
	// }
	campaignsFormatter := []CampaignFormatter{}

	for _, campaign := range campaigns {
		campaignFormatter := FormatCampaign(campaign)
		campaignsFormatter = append(campaignsFormatter, campaignFormatter)
	}
	return campaignsFormatter
}

type CampaingnDetailFormatter struct {
	ID               int                      `json:"id"`
	Name             string                   `json:"name"`
	ShortDescription string                   `json:"short_desription"`
	Description      string                   `json:"description"`
	ImageURL         string                   `json:"image_url"`
	GoalAmount       int                      `json:"goal_amount"`
	CurrentAmount    int                      `json:"current_amount"`
	UserID           int                      `json:"user_id"`
	Slug             string                   `json:"slug"`
	Perks            []string                 `'json:"perks"`
	User             CampaignUserFormatter    `json:"user"`
	Images           []CampaignImageFormatter `json:"images"`
}

type CampaignUserFormatter struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

type CampaignImageFormatter struct {
	ImageURL  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
}

func FormatCampaignDetail(campaign Campaign) CampaingnDetailFormatter {
	campaingnDetailFormatter := CampaingnDetailFormatter{}
	campaingnDetailFormatter.ID = campaign.ID
	campaingnDetailFormatter.Name = campaign.Name
	campaingnDetailFormatter.ShortDescription = campaign.ShortDescription
	campaingnDetailFormatter.GoalAmount = campaign.GoalAmount
	campaingnDetailFormatter.CurrentAmount = campaign.CurrentAmount
	campaingnDetailFormatter.UserID = campaign.UserID
	campaingnDetailFormatter.Slug = campaign.Slug
	campaingnDetailFormatter.ImageURL = ""

	if len(campaign.CampaignImages) > 0 {
		campaingnDetailFormatter.ImageURL = campaign.CampaignImages[0].FileName
	}

	var perks []string

	for _, perk := range strings.Split(campaign.Perks, ",") {
		perks = append(perks, strings.TrimSpace(perk))
	}
	campaingnDetailFormatter.Perks = perks

	user := campaign.User

	campaignUserFormatter := CampaignUserFormatter{}
	campaignUserFormatter.Name = user.Name
	// campaignUserFormatter.ImageURL = user.AvatarFileName

	campaingnDetailFormatter.User = campaignUserFormatter

	images := []CampaignImageFormatter{}

	for _, image := range campaign.CampaignImages {
		campaignImageFormatter := CampaignImageFormatter{}
		campaignImageFormatter.ImageURL = image.FileName

		isPrimary := false

		if image.IsPrimary == 1 {
			isPrimary = true
		}
		campaignImageFormatter.IsPrimary = isPrimary

		images = append(images, campaignImageFormatter)
	}

	campaingnDetailFormatter.Images = images

	return campaingnDetailFormatter

}
