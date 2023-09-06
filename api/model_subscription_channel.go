// Copyright 2023 GitBook, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gitbook

import (
	"encoding/json"
	"fmt"
)

// SubscriptionChannel - Channel to subscribe to for API updates.
type SubscriptionChannel struct {
	BackofficeUserInfoChannel         *BackofficeUserInfoChannel
	ChangeRequestReviewsChannel       *ChangeRequestReviewsChannel
	OrganizationCustomFieldsChannel   *OrganizationCustomFieldsChannel
	OrganizationEntitiesChannel       *OrganizationEntitiesChannel
	OrganizationEnvironmentsChannel   *OrganizationEnvironmentsChannel
	OrganizationMemberChannel         *OrganizationMemberChannel
	OrganizationMembersChannel        *OrganizationMembersChannel
	OrganizationRecordingsChannel     *OrganizationRecordingsChannel
	OrganizationSchemasChannel        *OrganizationSchemasChannel
	OrganizationSpaceRelationsChannel *OrganizationSpaceRelationsChannel
	OrganizationSpacesChannel         *OrganizationSpacesChannel
	OrganizationTeamChannel           *OrganizationTeamChannel
	OrganizationTeamMemberChannel     *OrganizationTeamMemberChannel
	OrganizationTeamMembersChannel    *OrganizationTeamMembersChannel
	OrganizationTeamsChannel          *OrganizationTeamsChannel
	SpaceCustomFieldsChannel          *SpaceCustomFieldsChannel
	SpaceEntitiesChannel              *SpaceEntitiesChannel
	SpaceGitInfoChannel               *SpaceGitInfoChannel
	SpaceInfoChannel                  *SpaceInfoChannel
	SpaceIntegrationsChannel          *SpaceIntegrationsChannel
	SpacePublishingAuthChannel        *SpacePublishingAuthChannel
	SpaceRelationsChannel             *SpaceRelationsChannel
	SubscriptionChannelOneOf          *SubscriptionChannelOneOf
	SubscriptionChannelOneOf1         *SubscriptionChannelOneOf1
	SubscriptionChannelOneOf2         *SubscriptionChannelOneOf2
	SubscriptionChannelOneOf3         *SubscriptionChannelOneOf3
	UserAPITokensChannel              *UserAPITokensChannel
}

// BackofficeUserInfoChannelAsSubscriptionChannel is a convenience function that returns BackofficeUserInfoChannel wrapped in SubscriptionChannel
func BackofficeUserInfoChannelAsSubscriptionChannel(v *BackofficeUserInfoChannel) SubscriptionChannel {
	return SubscriptionChannel{
		BackofficeUserInfoChannel: v,
	}
}

// ChangeRequestReviewsChannelAsSubscriptionChannel is a convenience function that returns ChangeRequestReviewsChannel wrapped in SubscriptionChannel
func ChangeRequestReviewsChannelAsSubscriptionChannel(v *ChangeRequestReviewsChannel) SubscriptionChannel {
	return SubscriptionChannel{
		ChangeRequestReviewsChannel: v,
	}
}

// OrganizationCustomFieldsChannelAsSubscriptionChannel is a convenience function that returns OrganizationCustomFieldsChannel wrapped in SubscriptionChannel
func OrganizationCustomFieldsChannelAsSubscriptionChannel(v *OrganizationCustomFieldsChannel) SubscriptionChannel {
	return SubscriptionChannel{
		OrganizationCustomFieldsChannel: v,
	}
}

// OrganizationEntitiesChannelAsSubscriptionChannel is a convenience function that returns OrganizationEntitiesChannel wrapped in SubscriptionChannel
func OrganizationEntitiesChannelAsSubscriptionChannel(v *OrganizationEntitiesChannel) SubscriptionChannel {
	return SubscriptionChannel{
		OrganizationEntitiesChannel: v,
	}
}

// OrganizationEnvironmentsChannelAsSubscriptionChannel is a convenience function that returns OrganizationEnvironmentsChannel wrapped in SubscriptionChannel
func OrganizationEnvironmentsChannelAsSubscriptionChannel(v *OrganizationEnvironmentsChannel) SubscriptionChannel {
	return SubscriptionChannel{
		OrganizationEnvironmentsChannel: v,
	}
}

// OrganizationMemberChannelAsSubscriptionChannel is a convenience function that returns OrganizationMemberChannel wrapped in SubscriptionChannel
func OrganizationMemberChannelAsSubscriptionChannel(v *OrganizationMemberChannel) SubscriptionChannel {
	return SubscriptionChannel{
		OrganizationMemberChannel: v,
	}
}

// OrganizationMembersChannelAsSubscriptionChannel is a convenience function that returns OrganizationMembersChannel wrapped in SubscriptionChannel
func OrganizationMembersChannelAsSubscriptionChannel(v *OrganizationMembersChannel) SubscriptionChannel {
	return SubscriptionChannel{
		OrganizationMembersChannel: v,
	}
}

// OrganizationRecordingsChannelAsSubscriptionChannel is a convenience function that returns OrganizationRecordingsChannel wrapped in SubscriptionChannel
func OrganizationRecordingsChannelAsSubscriptionChannel(v *OrganizationRecordingsChannel) SubscriptionChannel {
	return SubscriptionChannel{
		OrganizationRecordingsChannel: v,
	}
}

// OrganizationSchemasChannelAsSubscriptionChannel is a convenience function that returns OrganizationSchemasChannel wrapped in SubscriptionChannel
func OrganizationSchemasChannelAsSubscriptionChannel(v *OrganizationSchemasChannel) SubscriptionChannel {
	return SubscriptionChannel{
		OrganizationSchemasChannel: v,
	}
}

// OrganizationSpaceRelationsChannelAsSubscriptionChannel is a convenience function that returns OrganizationSpaceRelationsChannel wrapped in SubscriptionChannel
func OrganizationSpaceRelationsChannelAsSubscriptionChannel(v *OrganizationSpaceRelationsChannel) SubscriptionChannel {
	return SubscriptionChannel{
		OrganizationSpaceRelationsChannel: v,
	}
}

// OrganizationSpacesChannelAsSubscriptionChannel is a convenience function that returns OrganizationSpacesChannel wrapped in SubscriptionChannel
func OrganizationSpacesChannelAsSubscriptionChannel(v *OrganizationSpacesChannel) SubscriptionChannel {
	return SubscriptionChannel{
		OrganizationSpacesChannel: v,
	}
}

// OrganizationTeamChannelAsSubscriptionChannel is a convenience function that returns OrganizationTeamChannel wrapped in SubscriptionChannel
func OrganizationTeamChannelAsSubscriptionChannel(v *OrganizationTeamChannel) SubscriptionChannel {
	return SubscriptionChannel{
		OrganizationTeamChannel: v,
	}
}

// OrganizationTeamMemberChannelAsSubscriptionChannel is a convenience function that returns OrganizationTeamMemberChannel wrapped in SubscriptionChannel
func OrganizationTeamMemberChannelAsSubscriptionChannel(v *OrganizationTeamMemberChannel) SubscriptionChannel {
	return SubscriptionChannel{
		OrganizationTeamMemberChannel: v,
	}
}

// OrganizationTeamMembersChannelAsSubscriptionChannel is a convenience function that returns OrganizationTeamMembersChannel wrapped in SubscriptionChannel
func OrganizationTeamMembersChannelAsSubscriptionChannel(v *OrganizationTeamMembersChannel) SubscriptionChannel {
	return SubscriptionChannel{
		OrganizationTeamMembersChannel: v,
	}
}

// OrganizationTeamsChannelAsSubscriptionChannel is a convenience function that returns OrganizationTeamsChannel wrapped in SubscriptionChannel
func OrganizationTeamsChannelAsSubscriptionChannel(v *OrganizationTeamsChannel) SubscriptionChannel {
	return SubscriptionChannel{
		OrganizationTeamsChannel: v,
	}
}

// SpaceCustomFieldsChannelAsSubscriptionChannel is a convenience function that returns SpaceCustomFieldsChannel wrapped in SubscriptionChannel
func SpaceCustomFieldsChannelAsSubscriptionChannel(v *SpaceCustomFieldsChannel) SubscriptionChannel {
	return SubscriptionChannel{
		SpaceCustomFieldsChannel: v,
	}
}

// SpaceEntitiesChannelAsSubscriptionChannel is a convenience function that returns SpaceEntitiesChannel wrapped in SubscriptionChannel
func SpaceEntitiesChannelAsSubscriptionChannel(v *SpaceEntitiesChannel) SubscriptionChannel {
	return SubscriptionChannel{
		SpaceEntitiesChannel: v,
	}
}

// SpaceGitInfoChannelAsSubscriptionChannel is a convenience function that returns SpaceGitInfoChannel wrapped in SubscriptionChannel
func SpaceGitInfoChannelAsSubscriptionChannel(v *SpaceGitInfoChannel) SubscriptionChannel {
	return SubscriptionChannel{
		SpaceGitInfoChannel: v,
	}
}

// SpaceInfoChannelAsSubscriptionChannel is a convenience function that returns SpaceInfoChannel wrapped in SubscriptionChannel
func SpaceInfoChannelAsSubscriptionChannel(v *SpaceInfoChannel) SubscriptionChannel {
	return SubscriptionChannel{
		SpaceInfoChannel: v,
	}
}

// SpaceIntegrationsChannelAsSubscriptionChannel is a convenience function that returns SpaceIntegrationsChannel wrapped in SubscriptionChannel
func SpaceIntegrationsChannelAsSubscriptionChannel(v *SpaceIntegrationsChannel) SubscriptionChannel {
	return SubscriptionChannel{
		SpaceIntegrationsChannel: v,
	}
}

// SpacePublishingAuthChannelAsSubscriptionChannel is a convenience function that returns SpacePublishingAuthChannel wrapped in SubscriptionChannel
func SpacePublishingAuthChannelAsSubscriptionChannel(v *SpacePublishingAuthChannel) SubscriptionChannel {
	return SubscriptionChannel{
		SpacePublishingAuthChannel: v,
	}
}

// SpaceRelationsChannelAsSubscriptionChannel is a convenience function that returns SpaceRelationsChannel wrapped in SubscriptionChannel
func SpaceRelationsChannelAsSubscriptionChannel(v *SpaceRelationsChannel) SubscriptionChannel {
	return SubscriptionChannel{
		SpaceRelationsChannel: v,
	}
}

// SubscriptionChannelOneOfAsSubscriptionChannel is a convenience function that returns SubscriptionChannelOneOf wrapped in SubscriptionChannel
func SubscriptionChannelOneOfAsSubscriptionChannel(v *SubscriptionChannelOneOf) SubscriptionChannel {
	return SubscriptionChannel{
		SubscriptionChannelOneOf: v,
	}
}

// SubscriptionChannelOneOf1AsSubscriptionChannel is a convenience function that returns SubscriptionChannelOneOf1 wrapped in SubscriptionChannel
func SubscriptionChannelOneOf1AsSubscriptionChannel(v *SubscriptionChannelOneOf1) SubscriptionChannel {
	return SubscriptionChannel{
		SubscriptionChannelOneOf1: v,
	}
}

// SubscriptionChannelOneOf2AsSubscriptionChannel is a convenience function that returns SubscriptionChannelOneOf2 wrapped in SubscriptionChannel
func SubscriptionChannelOneOf2AsSubscriptionChannel(v *SubscriptionChannelOneOf2) SubscriptionChannel {
	return SubscriptionChannel{
		SubscriptionChannelOneOf2: v,
	}
}

// SubscriptionChannelOneOf3AsSubscriptionChannel is a convenience function that returns SubscriptionChannelOneOf3 wrapped in SubscriptionChannel
func SubscriptionChannelOneOf3AsSubscriptionChannel(v *SubscriptionChannelOneOf3) SubscriptionChannel {
	return SubscriptionChannel{
		SubscriptionChannelOneOf3: v,
	}
}

// UserAPITokensChannelAsSubscriptionChannel is a convenience function that returns UserAPITokensChannel wrapped in SubscriptionChannel
func UserAPITokensChannelAsSubscriptionChannel(v *UserAPITokensChannel) SubscriptionChannel {
	return SubscriptionChannel{
		UserAPITokensChannel: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *SubscriptionChannel) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into BackofficeUserInfoChannel
	err = newStrictDecoder(data).Decode(&dst.BackofficeUserInfoChannel)
	if err == nil {
		jsonBackofficeUserInfoChannel, _ := json.Marshal(dst.BackofficeUserInfoChannel)
		if string(jsonBackofficeUserInfoChannel) == "{}" { // empty struct
			dst.BackofficeUserInfoChannel = nil
		} else {
			match++
		}
	} else {
		dst.BackofficeUserInfoChannel = nil
	}

	// try to unmarshal data into ChangeRequestReviewsChannel
	err = newStrictDecoder(data).Decode(&dst.ChangeRequestReviewsChannel)
	if err == nil {
		jsonChangeRequestReviewsChannel, _ := json.Marshal(dst.ChangeRequestReviewsChannel)
		if string(jsonChangeRequestReviewsChannel) == "{}" { // empty struct
			dst.ChangeRequestReviewsChannel = nil
		} else {
			match++
		}
	} else {
		dst.ChangeRequestReviewsChannel = nil
	}

	// try to unmarshal data into OrganizationCustomFieldsChannel
	err = newStrictDecoder(data).Decode(&dst.OrganizationCustomFieldsChannel)
	if err == nil {
		jsonOrganizationCustomFieldsChannel, _ := json.Marshal(dst.OrganizationCustomFieldsChannel)
		if string(jsonOrganizationCustomFieldsChannel) == "{}" { // empty struct
			dst.OrganizationCustomFieldsChannel = nil
		} else {
			match++
		}
	} else {
		dst.OrganizationCustomFieldsChannel = nil
	}

	// try to unmarshal data into OrganizationEntitiesChannel
	err = newStrictDecoder(data).Decode(&dst.OrganizationEntitiesChannel)
	if err == nil {
		jsonOrganizationEntitiesChannel, _ := json.Marshal(dst.OrganizationEntitiesChannel)
		if string(jsonOrganizationEntitiesChannel) == "{}" { // empty struct
			dst.OrganizationEntitiesChannel = nil
		} else {
			match++
		}
	} else {
		dst.OrganizationEntitiesChannel = nil
	}

	// try to unmarshal data into OrganizationEnvironmentsChannel
	err = newStrictDecoder(data).Decode(&dst.OrganizationEnvironmentsChannel)
	if err == nil {
		jsonOrganizationEnvironmentsChannel, _ := json.Marshal(dst.OrganizationEnvironmentsChannel)
		if string(jsonOrganizationEnvironmentsChannel) == "{}" { // empty struct
			dst.OrganizationEnvironmentsChannel = nil
		} else {
			match++
		}
	} else {
		dst.OrganizationEnvironmentsChannel = nil
	}

	// try to unmarshal data into OrganizationMemberChannel
	err = newStrictDecoder(data).Decode(&dst.OrganizationMemberChannel)
	if err == nil {
		jsonOrganizationMemberChannel, _ := json.Marshal(dst.OrganizationMemberChannel)
		if string(jsonOrganizationMemberChannel) == "{}" { // empty struct
			dst.OrganizationMemberChannel = nil
		} else {
			match++
		}
	} else {
		dst.OrganizationMemberChannel = nil
	}

	// try to unmarshal data into OrganizationMembersChannel
	err = newStrictDecoder(data).Decode(&dst.OrganizationMembersChannel)
	if err == nil {
		jsonOrganizationMembersChannel, _ := json.Marshal(dst.OrganizationMembersChannel)
		if string(jsonOrganizationMembersChannel) == "{}" { // empty struct
			dst.OrganizationMembersChannel = nil
		} else {
			match++
		}
	} else {
		dst.OrganizationMembersChannel = nil
	}

	// try to unmarshal data into OrganizationRecordingsChannel
	err = newStrictDecoder(data).Decode(&dst.OrganizationRecordingsChannel)
	if err == nil {
		jsonOrganizationRecordingsChannel, _ := json.Marshal(dst.OrganizationRecordingsChannel)
		if string(jsonOrganizationRecordingsChannel) == "{}" { // empty struct
			dst.OrganizationRecordingsChannel = nil
		} else {
			match++
		}
	} else {
		dst.OrganizationRecordingsChannel = nil
	}

	// try to unmarshal data into OrganizationSchemasChannel
	err = newStrictDecoder(data).Decode(&dst.OrganizationSchemasChannel)
	if err == nil {
		jsonOrganizationSchemasChannel, _ := json.Marshal(dst.OrganizationSchemasChannel)
		if string(jsonOrganizationSchemasChannel) == "{}" { // empty struct
			dst.OrganizationSchemasChannel = nil
		} else {
			match++
		}
	} else {
		dst.OrganizationSchemasChannel = nil
	}

	// try to unmarshal data into OrganizationSpaceRelationsChannel
	err = newStrictDecoder(data).Decode(&dst.OrganizationSpaceRelationsChannel)
	if err == nil {
		jsonOrganizationSpaceRelationsChannel, _ := json.Marshal(dst.OrganizationSpaceRelationsChannel)
		if string(jsonOrganizationSpaceRelationsChannel) == "{}" { // empty struct
			dst.OrganizationSpaceRelationsChannel = nil
		} else {
			match++
		}
	} else {
		dst.OrganizationSpaceRelationsChannel = nil
	}

	// try to unmarshal data into OrganizationSpacesChannel
	err = newStrictDecoder(data).Decode(&dst.OrganizationSpacesChannel)
	if err == nil {
		jsonOrganizationSpacesChannel, _ := json.Marshal(dst.OrganizationSpacesChannel)
		if string(jsonOrganizationSpacesChannel) == "{}" { // empty struct
			dst.OrganizationSpacesChannel = nil
		} else {
			match++
		}
	} else {
		dst.OrganizationSpacesChannel = nil
	}

	// try to unmarshal data into OrganizationTeamChannel
	err = newStrictDecoder(data).Decode(&dst.OrganizationTeamChannel)
	if err == nil {
		jsonOrganizationTeamChannel, _ := json.Marshal(dst.OrganizationTeamChannel)
		if string(jsonOrganizationTeamChannel) == "{}" { // empty struct
			dst.OrganizationTeamChannel = nil
		} else {
			match++
		}
	} else {
		dst.OrganizationTeamChannel = nil
	}

	// try to unmarshal data into OrganizationTeamMemberChannel
	err = newStrictDecoder(data).Decode(&dst.OrganizationTeamMemberChannel)
	if err == nil {
		jsonOrganizationTeamMemberChannel, _ := json.Marshal(dst.OrganizationTeamMemberChannel)
		if string(jsonOrganizationTeamMemberChannel) == "{}" { // empty struct
			dst.OrganizationTeamMemberChannel = nil
		} else {
			match++
		}
	} else {
		dst.OrganizationTeamMemberChannel = nil
	}

	// try to unmarshal data into OrganizationTeamMembersChannel
	err = newStrictDecoder(data).Decode(&dst.OrganizationTeamMembersChannel)
	if err == nil {
		jsonOrganizationTeamMembersChannel, _ := json.Marshal(dst.OrganizationTeamMembersChannel)
		if string(jsonOrganizationTeamMembersChannel) == "{}" { // empty struct
			dst.OrganizationTeamMembersChannel = nil
		} else {
			match++
		}
	} else {
		dst.OrganizationTeamMembersChannel = nil
	}

	// try to unmarshal data into OrganizationTeamsChannel
	err = newStrictDecoder(data).Decode(&dst.OrganizationTeamsChannel)
	if err == nil {
		jsonOrganizationTeamsChannel, _ := json.Marshal(dst.OrganizationTeamsChannel)
		if string(jsonOrganizationTeamsChannel) == "{}" { // empty struct
			dst.OrganizationTeamsChannel = nil
		} else {
			match++
		}
	} else {
		dst.OrganizationTeamsChannel = nil
	}

	// try to unmarshal data into SpaceCustomFieldsChannel
	err = newStrictDecoder(data).Decode(&dst.SpaceCustomFieldsChannel)
	if err == nil {
		jsonSpaceCustomFieldsChannel, _ := json.Marshal(dst.SpaceCustomFieldsChannel)
		if string(jsonSpaceCustomFieldsChannel) == "{}" { // empty struct
			dst.SpaceCustomFieldsChannel = nil
		} else {
			match++
		}
	} else {
		dst.SpaceCustomFieldsChannel = nil
	}

	// try to unmarshal data into SpaceEntitiesChannel
	err = newStrictDecoder(data).Decode(&dst.SpaceEntitiesChannel)
	if err == nil {
		jsonSpaceEntitiesChannel, _ := json.Marshal(dst.SpaceEntitiesChannel)
		if string(jsonSpaceEntitiesChannel) == "{}" { // empty struct
			dst.SpaceEntitiesChannel = nil
		} else {
			match++
		}
	} else {
		dst.SpaceEntitiesChannel = nil
	}

	// try to unmarshal data into SpaceGitInfoChannel
	err = newStrictDecoder(data).Decode(&dst.SpaceGitInfoChannel)
	if err == nil {
		jsonSpaceGitInfoChannel, _ := json.Marshal(dst.SpaceGitInfoChannel)
		if string(jsonSpaceGitInfoChannel) == "{}" { // empty struct
			dst.SpaceGitInfoChannel = nil
		} else {
			match++
		}
	} else {
		dst.SpaceGitInfoChannel = nil
	}

	// try to unmarshal data into SpaceInfoChannel
	err = newStrictDecoder(data).Decode(&dst.SpaceInfoChannel)
	if err == nil {
		jsonSpaceInfoChannel, _ := json.Marshal(dst.SpaceInfoChannel)
		if string(jsonSpaceInfoChannel) == "{}" { // empty struct
			dst.SpaceInfoChannel = nil
		} else {
			match++
		}
	} else {
		dst.SpaceInfoChannel = nil
	}

	// try to unmarshal data into SpaceIntegrationsChannel
	err = newStrictDecoder(data).Decode(&dst.SpaceIntegrationsChannel)
	if err == nil {
		jsonSpaceIntegrationsChannel, _ := json.Marshal(dst.SpaceIntegrationsChannel)
		if string(jsonSpaceIntegrationsChannel) == "{}" { // empty struct
			dst.SpaceIntegrationsChannel = nil
		} else {
			match++
		}
	} else {
		dst.SpaceIntegrationsChannel = nil
	}

	// try to unmarshal data into SpacePublishingAuthChannel
	err = newStrictDecoder(data).Decode(&dst.SpacePublishingAuthChannel)
	if err == nil {
		jsonSpacePublishingAuthChannel, _ := json.Marshal(dst.SpacePublishingAuthChannel)
		if string(jsonSpacePublishingAuthChannel) == "{}" { // empty struct
			dst.SpacePublishingAuthChannel = nil
		} else {
			match++
		}
	} else {
		dst.SpacePublishingAuthChannel = nil
	}

	// try to unmarshal data into SpaceRelationsChannel
	err = newStrictDecoder(data).Decode(&dst.SpaceRelationsChannel)
	if err == nil {
		jsonSpaceRelationsChannel, _ := json.Marshal(dst.SpaceRelationsChannel)
		if string(jsonSpaceRelationsChannel) == "{}" { // empty struct
			dst.SpaceRelationsChannel = nil
		} else {
			match++
		}
	} else {
		dst.SpaceRelationsChannel = nil
	}

	// try to unmarshal data into SubscriptionChannelOneOf
	err = newStrictDecoder(data).Decode(&dst.SubscriptionChannelOneOf)
	if err == nil {
		jsonSubscriptionChannelOneOf, _ := json.Marshal(dst.SubscriptionChannelOneOf)
		if string(jsonSubscriptionChannelOneOf) == "{}" { // empty struct
			dst.SubscriptionChannelOneOf = nil
		} else {
			match++
		}
	} else {
		dst.SubscriptionChannelOneOf = nil
	}

	// try to unmarshal data into SubscriptionChannelOneOf1
	err = newStrictDecoder(data).Decode(&dst.SubscriptionChannelOneOf1)
	if err == nil {
		jsonSubscriptionChannelOneOf1, _ := json.Marshal(dst.SubscriptionChannelOneOf1)
		if string(jsonSubscriptionChannelOneOf1) == "{}" { // empty struct
			dst.SubscriptionChannelOneOf1 = nil
		} else {
			match++
		}
	} else {
		dst.SubscriptionChannelOneOf1 = nil
	}

	// try to unmarshal data into SubscriptionChannelOneOf2
	err = newStrictDecoder(data).Decode(&dst.SubscriptionChannelOneOf2)
	if err == nil {
		jsonSubscriptionChannelOneOf2, _ := json.Marshal(dst.SubscriptionChannelOneOf2)
		if string(jsonSubscriptionChannelOneOf2) == "{}" { // empty struct
			dst.SubscriptionChannelOneOf2 = nil
		} else {
			match++
		}
	} else {
		dst.SubscriptionChannelOneOf2 = nil
	}

	// try to unmarshal data into SubscriptionChannelOneOf3
	err = newStrictDecoder(data).Decode(&dst.SubscriptionChannelOneOf3)
	if err == nil {
		jsonSubscriptionChannelOneOf3, _ := json.Marshal(dst.SubscriptionChannelOneOf3)
		if string(jsonSubscriptionChannelOneOf3) == "{}" { // empty struct
			dst.SubscriptionChannelOneOf3 = nil
		} else {
			match++
		}
	} else {
		dst.SubscriptionChannelOneOf3 = nil
	}

	// try to unmarshal data into UserAPITokensChannel
	err = newStrictDecoder(data).Decode(&dst.UserAPITokensChannel)
	if err == nil {
		jsonUserAPITokensChannel, _ := json.Marshal(dst.UserAPITokensChannel)
		if string(jsonUserAPITokensChannel) == "{}" { // empty struct
			dst.UserAPITokensChannel = nil
		} else {
			match++
		}
	} else {
		dst.UserAPITokensChannel = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.BackofficeUserInfoChannel = nil
		dst.ChangeRequestReviewsChannel = nil
		dst.OrganizationCustomFieldsChannel = nil
		dst.OrganizationEntitiesChannel = nil
		dst.OrganizationEnvironmentsChannel = nil
		dst.OrganizationMemberChannel = nil
		dst.OrganizationMembersChannel = nil
		dst.OrganizationRecordingsChannel = nil
		dst.OrganizationSchemasChannel = nil
		dst.OrganizationSpaceRelationsChannel = nil
		dst.OrganizationSpacesChannel = nil
		dst.OrganizationTeamChannel = nil
		dst.OrganizationTeamMemberChannel = nil
		dst.OrganizationTeamMembersChannel = nil
		dst.OrganizationTeamsChannel = nil
		dst.SpaceCustomFieldsChannel = nil
		dst.SpaceEntitiesChannel = nil
		dst.SpaceGitInfoChannel = nil
		dst.SpaceInfoChannel = nil
		dst.SpaceIntegrationsChannel = nil
		dst.SpacePublishingAuthChannel = nil
		dst.SpaceRelationsChannel = nil
		dst.SubscriptionChannelOneOf = nil
		dst.SubscriptionChannelOneOf1 = nil
		dst.SubscriptionChannelOneOf2 = nil
		dst.SubscriptionChannelOneOf3 = nil
		dst.UserAPITokensChannel = nil

		return fmt.Errorf("data matches more than one schema in oneOf(SubscriptionChannel)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(SubscriptionChannel)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SubscriptionChannel) MarshalJSON() ([]byte, error) {
	if src.BackofficeUserInfoChannel != nil {
		return json.Marshal(&src.BackofficeUserInfoChannel)
	}

	if src.ChangeRequestReviewsChannel != nil {
		return json.Marshal(&src.ChangeRequestReviewsChannel)
	}

	if src.OrganizationCustomFieldsChannel != nil {
		return json.Marshal(&src.OrganizationCustomFieldsChannel)
	}

	if src.OrganizationEntitiesChannel != nil {
		return json.Marshal(&src.OrganizationEntitiesChannel)
	}

	if src.OrganizationEnvironmentsChannel != nil {
		return json.Marshal(&src.OrganizationEnvironmentsChannel)
	}

	if src.OrganizationMemberChannel != nil {
		return json.Marshal(&src.OrganizationMemberChannel)
	}

	if src.OrganizationMembersChannel != nil {
		return json.Marshal(&src.OrganizationMembersChannel)
	}

	if src.OrganizationRecordingsChannel != nil {
		return json.Marshal(&src.OrganizationRecordingsChannel)
	}

	if src.OrganizationSchemasChannel != nil {
		return json.Marshal(&src.OrganizationSchemasChannel)
	}

	if src.OrganizationSpaceRelationsChannel != nil {
		return json.Marshal(&src.OrganizationSpaceRelationsChannel)
	}

	if src.OrganizationSpacesChannel != nil {
		return json.Marshal(&src.OrganizationSpacesChannel)
	}

	if src.OrganizationTeamChannel != nil {
		return json.Marshal(&src.OrganizationTeamChannel)
	}

	if src.OrganizationTeamMemberChannel != nil {
		return json.Marshal(&src.OrganizationTeamMemberChannel)
	}

	if src.OrganizationTeamMembersChannel != nil {
		return json.Marshal(&src.OrganizationTeamMembersChannel)
	}

	if src.OrganizationTeamsChannel != nil {
		return json.Marshal(&src.OrganizationTeamsChannel)
	}

	if src.SpaceCustomFieldsChannel != nil {
		return json.Marshal(&src.SpaceCustomFieldsChannel)
	}

	if src.SpaceEntitiesChannel != nil {
		return json.Marshal(&src.SpaceEntitiesChannel)
	}

	if src.SpaceGitInfoChannel != nil {
		return json.Marshal(&src.SpaceGitInfoChannel)
	}

	if src.SpaceInfoChannel != nil {
		return json.Marshal(&src.SpaceInfoChannel)
	}

	if src.SpaceIntegrationsChannel != nil {
		return json.Marshal(&src.SpaceIntegrationsChannel)
	}

	if src.SpacePublishingAuthChannel != nil {
		return json.Marshal(&src.SpacePublishingAuthChannel)
	}

	if src.SpaceRelationsChannel != nil {
		return json.Marshal(&src.SpaceRelationsChannel)
	}

	if src.SubscriptionChannelOneOf != nil {
		return json.Marshal(&src.SubscriptionChannelOneOf)
	}

	if src.SubscriptionChannelOneOf1 != nil {
		return json.Marshal(&src.SubscriptionChannelOneOf1)
	}

	if src.SubscriptionChannelOneOf2 != nil {
		return json.Marshal(&src.SubscriptionChannelOneOf2)
	}

	if src.SubscriptionChannelOneOf3 != nil {
		return json.Marshal(&src.SubscriptionChannelOneOf3)
	}

	if src.UserAPITokensChannel != nil {
		return json.Marshal(&src.UserAPITokensChannel)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SubscriptionChannel) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.BackofficeUserInfoChannel != nil {
		return obj.BackofficeUserInfoChannel
	}

	if obj.ChangeRequestReviewsChannel != nil {
		return obj.ChangeRequestReviewsChannel
	}

	if obj.OrganizationCustomFieldsChannel != nil {
		return obj.OrganizationCustomFieldsChannel
	}

	if obj.OrganizationEntitiesChannel != nil {
		return obj.OrganizationEntitiesChannel
	}

	if obj.OrganizationEnvironmentsChannel != nil {
		return obj.OrganizationEnvironmentsChannel
	}

	if obj.OrganizationMemberChannel != nil {
		return obj.OrganizationMemberChannel
	}

	if obj.OrganizationMembersChannel != nil {
		return obj.OrganizationMembersChannel
	}

	if obj.OrganizationRecordingsChannel != nil {
		return obj.OrganizationRecordingsChannel
	}

	if obj.OrganizationSchemasChannel != nil {
		return obj.OrganizationSchemasChannel
	}

	if obj.OrganizationSpaceRelationsChannel != nil {
		return obj.OrganizationSpaceRelationsChannel
	}

	if obj.OrganizationSpacesChannel != nil {
		return obj.OrganizationSpacesChannel
	}

	if obj.OrganizationTeamChannel != nil {
		return obj.OrganizationTeamChannel
	}

	if obj.OrganizationTeamMemberChannel != nil {
		return obj.OrganizationTeamMemberChannel
	}

	if obj.OrganizationTeamMembersChannel != nil {
		return obj.OrganizationTeamMembersChannel
	}

	if obj.OrganizationTeamsChannel != nil {
		return obj.OrganizationTeamsChannel
	}

	if obj.SpaceCustomFieldsChannel != nil {
		return obj.SpaceCustomFieldsChannel
	}

	if obj.SpaceEntitiesChannel != nil {
		return obj.SpaceEntitiesChannel
	}

	if obj.SpaceGitInfoChannel != nil {
		return obj.SpaceGitInfoChannel
	}

	if obj.SpaceInfoChannel != nil {
		return obj.SpaceInfoChannel
	}

	if obj.SpaceIntegrationsChannel != nil {
		return obj.SpaceIntegrationsChannel
	}

	if obj.SpacePublishingAuthChannel != nil {
		return obj.SpacePublishingAuthChannel
	}

	if obj.SpaceRelationsChannel != nil {
		return obj.SpaceRelationsChannel
	}

	if obj.SubscriptionChannelOneOf != nil {
		return obj.SubscriptionChannelOneOf
	}

	if obj.SubscriptionChannelOneOf1 != nil {
		return obj.SubscriptionChannelOneOf1
	}

	if obj.SubscriptionChannelOneOf2 != nil {
		return obj.SubscriptionChannelOneOf2
	}

	if obj.SubscriptionChannelOneOf3 != nil {
		return obj.SubscriptionChannelOneOf3
	}

	if obj.UserAPITokensChannel != nil {
		return obj.UserAPITokensChannel
	}

	// all schemas are nil
	return nil
}

type NullableSubscriptionChannel struct {
	value *SubscriptionChannel
	isSet bool
}

func (v NullableSubscriptionChannel) Get() *SubscriptionChannel {
	return v.value
}

func (v *NullableSubscriptionChannel) Set(val *SubscriptionChannel) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionChannel) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionChannel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionChannel(val *SubscriptionChannel) *NullableSubscriptionChannel {
	return &NullableSubscriptionChannel{value: val, isSet: true}
}

func (v NullableSubscriptionChannel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionChannel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
