# go-gitbook

Go client for the GitBook API.


## Installation

```shell
go get github.com/GitbookIO/go-gitbook
```

## Usage

Import the package:

```golang
import gitbook "github.com/GitbookIO/go-gitbook/api"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), gitbook.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), gitbook.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), gitbook.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), gitbook.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.gitbook.com/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AnalyticsApi* | [**GetContentAnalyticsForSpaceById**](docs/AnalyticsApi.md#getcontentanalyticsforspacebyid) | **Get** /spaces/{spaceId}/insights/content | Get content analytics for a given space.
*AnalyticsApi* | [**GetSearchAnalyticsForSpaceById**](docs/AnalyticsApi.md#getsearchanalyticsforspacebyid) | **Get** /spaces/{spaceId}/insights/search | Get an overview of the top search queries in a space.
*AnalyticsApi* | [**GetTrafficAnalyticsForSpaceById**](docs/AnalyticsApi.md#gettrafficanalyticsforspacebyid) | **Get** /spaces/{spaceId}/insights/traffic | Get traffic page views for a given space
*AnalyticsApi* | [**TrackViewInSpaceById**](docs/AnalyticsApi.md#trackviewinspacebyid) | **Post** /spaces/{spaceId}/insights/track_view | 
*ApiApi* | [**GetApiInformation**](docs/ApiApi.md#getapiinformation) | **Get** / | Get information about the state of the GitBook API
*CollectionsApi* | [**GetCollectionById**](docs/CollectionsApi.md#getcollectionbyid) | **Get** /collections/{collectionId} | Get the details about a collection using its ID
*CollectionsApi* | [**GetCollectionPublishingAuthById**](docs/CollectionsApi.md#getcollectionpublishingauthbyid) | **Get** /collections/{collectionId}/publishing/auth | Get the publishing authentication configuration for a collection.
*CollectionsApi* | [**ListSpacesInCollectionById**](docs/CollectionsApi.md#listspacesincollectionbyid) | **Get** /collections/{collectionId}/spaces | List all the spaces in a collection by its ID.
*CollectionsApi* | [**UpdateCollectionById**](docs/CollectionsApi.md#updatecollectionbyid) | **Patch** /collections/{collectionId} | Update the details of a collection
*CollectionsApi* | [**UpdateCollectionPublishingAuthById**](docs/CollectionsApi.md#updatecollectionpublishingauthbyid) | **Post** /collections/{collectionId}/publishing/auth | Update the publishing authentication configuration for a collection.
*DefaultApi* | [**AskQuery**](docs/DefaultApi.md#askquery) | **Post** /search/ask | Ask a question to an AI across spaces that is accessible by the currently authenticated target.
*DefaultApi* | [**AskQueryInSpace**](docs/DefaultApi.md#askqueryinspace) | **Post** /spaces/{spaceId}/search/ask | Ask a question to an AI within the context of the space.
*DefaultApi* | [**AskQueryInSpaceWithGet**](docs/DefaultApi.md#askqueryinspacewithget) | **Get** /spaces/{spaceId}/search/ask | Ask a question to an AI within the context of the space.
*DefaultApi* | [**AskQueryWithGet**](docs/DefaultApi.md#askquerywithget) | **Get** /search/ask | Ask a question to an AI across spaces that is accessible by the currently authenticated target.
*DefaultApi* | [**GetCurrentRevision**](docs/DefaultApi.md#getcurrentrevision) | **Get** /spaces/{spaceId}/content | Get the current primary content revision for a space
*DefaultApi* | [**GetPageById**](docs/DefaultApi.md#getpagebyid) | **Get** /spaces/{spaceId}/content/page/{pageId} | Get a page by its ID in the primary content.
*DefaultApi* | [**GetPageByPath**](docs/DefaultApi.md#getpagebypath) | **Get** /spaces/{spaceId}/content/path/{pagePath} | Get a page by its path in the primary content.
*DefaultApi* | [**GetPageInChangeRequestById**](docs/DefaultApi.md#getpageinchangerequestbyid) | **Get** /spaces/{spaceId}/change-requests/{changeRequestId}/content/page/{pageId} | Get a page by its ID in a change request.
*DefaultApi* | [**GetPageInChangeRequestByPath**](docs/DefaultApi.md#getpageinchangerequestbypath) | **Get** /spaces/{spaceId}/change-requests/{changeRequestId}/content/path/{pagePath} | Get a page by its path in a change request.
*DefaultApi* | [**GetPageInRevisionById**](docs/DefaultApi.md#getpageinrevisionbyid) | **Get** /spaces/{spaceId}/revisions/{revisionId}/page/{pageId} | Get a page by its ID in a revision.
*DefaultApi* | [**GetPageInRevisionByPath**](docs/DefaultApi.md#getpageinrevisionbypath) | **Get** /spaces/{spaceId}/revisions/{revisionId}/path/{pagePath} | Get a page by its path in a revision.
*DefaultApi* | [**GetRecommendedQuestions**](docs/DefaultApi.md#getrecommendedquestions) | **Post** /search/questions | Get a list of questions recommended by AI for a list of content.
*DefaultApi* | [**GetRecommendedQuestionsInSpace**](docs/DefaultApi.md#getrecommendedquestionsinspace) | **Get** /spaces/{spaceId}/search/questions | Get a list of questions that can be asked in a space.
*DefaultApi* | [**GetRevisionById**](docs/DefaultApi.md#getrevisionbyid) | **Get** /spaces/{spaceId}/revisions/{revisionId} | Get a specific revision in a space
*DefaultApi* | [**GetRevisionOfChangeRequestById**](docs/DefaultApi.md#getrevisionofchangerequestbyid) | **Get** /spaces/{spaceId}/change-requests/{changeRequestId}/content | Get the latest content revision for a change request.
*DefaultApi* | [**ImportContent**](docs/DefaultApi.md#importcontent) | **Post** /spaces/{spaceId}/content/import | Import content in a space.
*DefaultApi* | [**ImportContentInChangeRequest**](docs/DefaultApi.md#importcontentinchangerequest) | **Post** /spaces/{spaceId}/change-requests/{changeRequestId}/content/import | Import content in a change request.
*DefaultApi* | [**ImportContentInChangeRequestPageById**](docs/DefaultApi.md#importcontentinchangerequestpagebyid) | **Post** /spaces/{spaceId}/change-requests/{changeRequestId}/content/page/{pageId}/import | Import external content into a page of a change-request by its ID.
*DefaultApi* | [**ImportContentInPageById**](docs/DefaultApi.md#importcontentinpagebyid) | **Post** /spaces/{spaceId}/content/page/{pageId}/import | Import external content into a page by its ID.
*DefaultApi* | [**InstallIntegrationOnSpace**](docs/DefaultApi.md#installintegrationonspace) | **Post** /integrations/{integrationName}/installations/{installationId}/spaces | Install integration on a space using an existing installation
*DefaultApi* | [**ListFiles**](docs/DefaultApi.md#listfiles) | **Get** /spaces/{spaceId}/content/files | List all files for the latest primary revision content for a space
*DefaultApi* | [**ListFilesInChangeRequestById**](docs/DefaultApi.md#listfilesinchangerequestbyid) | **Get** /spaces/{spaceId}/change-requests/{changeRequestId}/content/files | List all files in the latest content of the change-request
*DefaultApi* | [**ListFilesInRevisionById**](docs/DefaultApi.md#listfilesinrevisionbyid) | **Get** /spaces/{spaceId}/revisions/{revisionId}/files | List all files in a revision
*DefaultApi* | [**RenderIntegrationUIWithGet**](docs/DefaultApi.md#renderintegrationuiwithget) | **Get** /integrations/{integrationName}/render | Render an integration UI in the context of an installation.
*DefaultApi* | [**RenderIntegrationUIWithPost**](docs/DefaultApi.md#renderintegrationuiwithpost) | **Post** /integrations/{integrationName}/render | Render an integration UI in the context of an installation.
*DefaultApi* | [**SearchSpaceContent**](docs/DefaultApi.md#searchspacecontent) | **Get** /spaces/{spaceId}/search | Search content in a space
*DefaultApi* | [**UninstallIntegrationFromSpace**](docs/DefaultApi.md#uninstallintegrationfromspace) | **Delete** /integrations/{integrationName}/installations/{installationId}/spaces/{spaceId} | Uninstall the integration from a space
*DefaultApi* | [**UpdateIntegrationSpaceInstallation**](docs/DefaultApi.md#updateintegrationspaceinstallation) | **Patch** /integrations/{integrationName}/installations/{installationId}/spaces/{spaceId} | Update external IDs and configurations of an integration&#39;s installation for a space
*DsyncApi* | [**ListDirectorySyncGroups**](docs/DsyncApi.md#listdirectorysyncgroups) | **Get** /orgs/{organizationId}/dsync/groups | Lists the groups exposed to the synced Directory on an organization.
*DsyncApi* | [**SetupDirectorySync**](docs/DsyncApi.md#setupdirectorysync) | **Post** /orgs/{organizationId}/dsync | Set up Directory Sync in an organization.
*DsyncApi* | [**SyncDirectorySyncGroupsToTeams**](docs/DsyncApi.md#syncdirectorysyncgroupstoteams) | **Post** /orgs/{organizationId}/dsync/teams | Syncs a list of group/team unique identifiers pairs together.
*EnvironmentsApi* | [**CreateEnvironment**](docs/EnvironmentsApi.md#createenvironment) | **Post** /orgs/{organizationId}/environments | Create a new environment within an organization
*EnvironmentsApi* | [**DeleteEnvironment**](docs/EnvironmentsApi.md#deleteenvironment) | **Delete** /orgs/{organizationId}/environments/{environmentName} | Delete an environment in an organization
*EnvironmentsApi* | [**GetEnvironmentByName**](docs/EnvironmentsApi.md#getenvironmentbyname) | **Get** /orgs/{organizationId}/environments/{environmentName} | Get an environment by its name
*EnvironmentsApi* | [**ListEnvironments**](docs/EnvironmentsApi.md#listenvironments) | **Get** /orgs/{organizationId}/environments | Get the environments in an organization
*EnvironmentsApi* | [**UpdateEnvironment**](docs/EnvironmentsApi.md#updateenvironment) | **Patch** /orgs/{organizationId}/environments/{environmentName} | Update an existing environment within an organization
*HiveApi* | [**GenerateHiveAccessToken**](docs/HiveApi.md#generatehiveaccesstoken) | **Post** /internal/hive/token | Returns a token to authenticate with Hive.
*HiveApi* | [**GenerateSpaceHiveReadAccessToken**](docs/HiveApi.md#generatespacehivereadaccesstoken) | **Post** /spaces/{spaceId}/hive/token | Returns a token to authenticate with Hive to read content from a given space.
*IntegrationsApi* | [**CreateIntegrationInstallationToken**](docs/IntegrationsApi.md#createintegrationinstallationtoken) | **Post** /integrations/{integrationName}/installations/{installationId}/tokens | Create an integration installation API token
*IntegrationsApi* | [**GetIntegrationByName**](docs/IntegrationsApi.md#getintegrationbyname) | **Get** /integrations/{integrationName} | Get a specific integration by its name
*IntegrationsApi* | [**GetIntegrationEntities**](docs/IntegrationsApi.md#getintegrationentities) | **Get** /integrations/{integrationName}/installations/{installationId}/entities | List entities managed by an integration in an installation.
*IntegrationsApi* | [**GetIntegrationEvent**](docs/IntegrationsApi.md#getintegrationevent) | **Get** /integrations/{integrationName}/events/{eventId} | Get a specific integration event by its id
*IntegrationsApi* | [**GetIntegrationInstallationById**](docs/IntegrationsApi.md#getintegrationinstallationbyid) | **Get** /integrations/{integrationName}/installations/{installationId} | Get a specific integration&#39;s installation by its ID
*IntegrationsApi* | [**InstallIntegration**](docs/IntegrationsApi.md#installintegration) | **Post** /integrations/{integrationName}/installations | Install integration on a target organization
*IntegrationsApi* | [**ListIntegrationEvents**](docs/IntegrationsApi.md#listintegrationevents) | **Get** /integrations/{integrationName}/events | List all integration events
*IntegrationsApi* | [**ListIntegrationInstallations**](docs/IntegrationsApi.md#listintegrationinstallations) | **Get** /integrations/{integrationName}/installations | Fetch a list of installations of an integration
*IntegrationsApi* | [**ListIntegrationSpaceInstallations**](docs/IntegrationsApi.md#listintegrationspaceinstallations) | **Get** /integrations/{integrationName}/spaces | Fetch a list of space installations of an integration
*IntegrationsApi* | [**ListIntegrations**](docs/IntegrationsApi.md#listintegrations) | **Get** /integrations | List all public integrations
*IntegrationsApi* | [**ListSpaceIntegrationsBlocks**](docs/IntegrationsApi.md#listspaceintegrationsblocks) | **Get** /spaces/{spaceId}/integration-blocks | List integrations blocks for a space
*IntegrationsApi* | [**PublishIntegration**](docs/IntegrationsApi.md#publishintegration) | **Post** /integrations/{integrationName} | Publish an integration
*IntegrationsApi* | [**RemoveIntegrationDevSpace**](docs/IntegrationsApi.md#removeintegrationdevspace) | **Delete** /integrations/{integrationName}/spaces/{spaceId}/dev | Remove the development space for an integration
*IntegrationsApi* | [**SyncIntegrationEntities**](docs/IntegrationsApi.md#syncintegrationentities) | **Post** /integrations/{integrationName}/installations/{installationId}/entities | Update all entities for an integration installation. Entities will be created and updated, missing entities will be deleted.
*IntegrationsApi* | [**UninstallIntegration**](docs/IntegrationsApi.md#uninstallintegration) | **Delete** /integrations/{integrationName}/installations/{installationId} | Uninstall the integration from a target organization
*IntegrationsApi* | [**UnpublishIntegration**](docs/IntegrationsApi.md#unpublishintegration) | **Delete** /integrations/{integrationName} | Unpublish an integration
*IntegrationsApi* | [**UpdateIntegrationDevSpace**](docs/IntegrationsApi.md#updateintegrationdevspace) | **Put** /integrations/{integrationName}/spaces/{spaceId}/dev | Update the development space for an integration
*IntegrationsApi* | [**UpdateIntegrationInstallation**](docs/IntegrationsApi.md#updateintegrationinstallation) | **Patch** /integrations/{integrationName}/installations/{installationId} | Update external IDs and configurations of an integration&#39;s installation
*OrganizationsApi* | [**AddMemberToOrganizationTeamById**](docs/OrganizationsApi.md#addmembertoorganizationteambyid) | **Put** /orgs/{organizationId}/teams/{teamId}/members/{userId} | Add or update a team membership
*OrganizationsApi* | [**CreateEnvironment**](docs/OrganizationsApi.md#createenvironment) | **Post** /orgs/{organizationId}/environments | Create a new environment within an organization
*OrganizationsApi* | [**CreateOrganization**](docs/OrganizationsApi.md#createorganization) | **Post** /orgs | Create an organization
*OrganizationsApi* | [**CreateOrganizationCustomField**](docs/OrganizationsApi.md#createorganizationcustomfield) | **Post** /orgs/{organizationId}/custom-fields | Create a new custom field in an orgamization
*OrganizationsApi* | [**CreateOrganizationTeam**](docs/OrganizationsApi.md#createorganizationteam) | **Put** /orgs/{organizationId}/teams | Create organization team
*OrganizationsApi* | [**DeleteEnvironment**](docs/OrganizationsApi.md#deleteenvironment) | **Delete** /orgs/{organizationId}/environments/{environmentName} | Delete an environment in an organization
*OrganizationsApi* | [**DeleteMemberFromOrganizationTeamById**](docs/OrganizationsApi.md#deletememberfromorganizationteambyid) | **Delete** /orgs/{organizationId}/teams/{teamId}/members/{userId} | Delete members from a team
*OrganizationsApi* | [**DeleteOrganizationCustomField**](docs/OrganizationsApi.md#deleteorganizationcustomfield) | **Delete** /orgs/{organizationId}/custom-fields/{fieldName} | Delete a custom field in an organization
*OrganizationsApi* | [**GetEnvironmentByName**](docs/OrganizationsApi.md#getenvironmentbyname) | **Get** /orgs/{organizationId}/environments/{environmentName} | Get an environment by its name
*OrganizationsApi* | [**GetMemberInOrganizationById**](docs/OrganizationsApi.md#getmemberinorganizationbyid) | **Get** /orgs/{organizationId}/members/{userId} | Get specified organization member
*OrganizationsApi* | [**GetOrganizationBillingPortal**](docs/OrganizationsApi.md#getorganizationbillingportal) | **Get** /orgs/{organizationId}/billing | Get the billing portal for an organization
*OrganizationsApi* | [**GetOrganizationById**](docs/OrganizationsApi.md#getorganizationbyid) | **Get** /orgs/{organizationId} | Get an organization by its ID
*OrganizationsApi* | [**GetOrganizationCustomFieldByName**](docs/OrganizationsApi.md#getorganizationcustomfieldbyname) | **Get** /orgs/{organizationId}/custom-fields/{fieldName} | Get a custom field by its name
*OrganizationsApi* | [**GetTeamInOrganizationById**](docs/OrganizationsApi.md#getteaminorganizationbyid) | **Get** /orgs/{organizationId}/teams/{teamId} | Get specified organization team
*OrganizationsApi* | [**InviteUsersToOrganization**](docs/OrganizationsApi.md#inviteuserstoorganization) | **Post** /orgs/{organizationId}/invites | Invite users to a given organization based on a list of emails
*OrganizationsApi* | [**JoinOrganizationWithInvite**](docs/OrganizationsApi.md#joinorganizationwithinvite) | **Post** /orgs/{organizationId}/invites/{inviteId} | Use an invite to join an organization.
*OrganizationsApi* | [**ListCollectionsInOrganizationById**](docs/OrganizationsApi.md#listcollectionsinorganizationbyid) | **Get** /orgs/{organizationId}/collections | List organization collections
*OrganizationsApi* | [**ListDirectorySyncGroups**](docs/OrganizationsApi.md#listdirectorysyncgroups) | **Get** /orgs/{organizationId}/dsync/groups | Lists the groups exposed to the synced Directory on an organization.
*OrganizationsApi* | [**ListEnvironments**](docs/OrganizationsApi.md#listenvironments) | **Get** /orgs/{organizationId}/environments | Get the environments in an organization
*OrganizationsApi* | [**ListMembersInOrganizationById**](docs/OrganizationsApi.md#listmembersinorganizationbyid) | **Get** /orgs/{organizationId}/members | List organization members
*OrganizationsApi* | [**ListOrganizationCustomFields**](docs/OrganizationsApi.md#listorganizationcustomfields) | **Get** /orgs/{organizationId}/custom-fields | Get the custom fields for spaces in an organization
*OrganizationsApi* | [**ListOrganizationsForAuthenticatedUser**](docs/OrganizationsApi.md#listorganizationsforauthenticateduser) | **Get** /orgs | Get the list of organizations for the currently authenticated user
*OrganizationsApi* | [**ListSpacesInOrganizationById**](docs/OrganizationsApi.md#listspacesinorganizationbyid) | **Get** /orgs/{organizationId}/spaces | List organization spaces
*OrganizationsApi* | [**ListSpacesWithGitSyncInOrganizationById**](docs/OrganizationsApi.md#listspaceswithgitsyncinorganizationbyid) | **Get** /orgs/{organizationId}/spaces/gitsync | List organization spaces including Git sync metadata
*OrganizationsApi* | [**ListTeamMembersInOrganizationById**](docs/OrganizationsApi.md#listteammembersinorganizationbyid) | **Get** /orgs/{organizationId}/teams/{teamId}/members | List team members
*OrganizationsApi* | [**ListTeamsInOrganizationById**](docs/OrganizationsApi.md#listteamsinorganizationbyid) | **Get** /orgs/{organizationId}/teams | List organization teams
*OrganizationsApi* | [**RemoveMemberFromOrganizationById**](docs/OrganizationsApi.md#removememberfromorganizationbyid) | **Delete** /orgs/{organizationId}/members/{userId} | Delete a member from an organization
*OrganizationsApi* | [**RemoveTeamFromOrganizationById**](docs/OrganizationsApi.md#removeteamfromorganizationbyid) | **Delete** /orgs/{organizationId}/teams/{teamId} | Delete a team in an organization
*OrganizationsApi* | [**RequestOrganizationUpgrade**](docs/OrganizationsApi.md#requestorganizationupgrade) | **Post** /orgs/{organizationId}/request_upgrade | Send a request to ask the organization&#39;s admin to upgrade it.
*OrganizationsApi* | [**SearchOrganizationContent**](docs/OrganizationsApi.md#searchorganizationcontent) | **Get** /orgs/{organizationId}/search | Search content in an organization
*OrganizationsApi* | [**SetUserAsSSOMemberForOrganization**](docs/OrganizationsApi.md#setuserasssomemberfororganization) | **Post** /orgs/{organizationId}/members/{userId}/sso | Set a user as an SSO member of an organization
*OrganizationsApi* | [**SetupDirectorySync**](docs/OrganizationsApi.md#setupdirectorysync) | **Post** /orgs/{organizationId}/dsync | Set up Directory Sync in an organization.
*OrganizationsApi* | [**SyncDirectorySyncGroupsToTeams**](docs/OrganizationsApi.md#syncdirectorysyncgroupstoteams) | **Post** /orgs/{organizationId}/dsync/teams | Syncs a list of group/team unique identifiers pairs together.
*OrganizationsApi* | [**TransferOrganization**](docs/OrganizationsApi.md#transferorganization) | **Post** /orgs/{organizationId}/transfer | Transfer one organization (source) into another organization (target).
*OrganizationsApi* | [**UpdateEnvironment**](docs/OrganizationsApi.md#updateenvironment) | **Patch** /orgs/{organizationId}/environments/{environmentName} | Update an existing environment within an organization
*OrganizationsApi* | [**UpdateMemberInOrganizationById**](docs/OrganizationsApi.md#updatememberinorganizationbyid) | **Patch** /orgs/{organizationId}/members/{userId} | Update specified organization member
*OrganizationsApi* | [**UpdateMembersInOrganizationTeam**](docs/OrganizationsApi.md#updatemembersinorganizationteam) | **Put** /orgs/{organizationId}/teams/{teamId}/members | Updates members of an organization team
*OrganizationsApi* | [**UpdateOrganizationCustomField**](docs/OrganizationsApi.md#updateorganizationcustomfield) | **Patch** /orgs/{organizationId}/custom-fields/{fieldName} | Update a custom field in an organization
*OrganizationsApi* | [**UpdateOrganizationMemberLastSeenAt**](docs/OrganizationsApi.md#updateorganizationmemberlastseenat) | **Post** /orgs/{organizationId}/ping | Update organization member&#39;s \&quot;last seen at\&quot; timestamp.
*OrganizationsApi* | [**UpdateTeamInOrganizationById**](docs/OrganizationsApi.md#updateteaminorganizationbyid) | **Patch** /orgs/{organizationId}/teams/{teamId} | Update specified organization team
*OrganizationsApi* | [**UpgradeOrganizationPlan**](docs/OrganizationsApi.md#upgradeorganizationplan) | **Post** /orgs/{organizationId}/billing | Upgrade an organization&#39;s billing plan
*PermissionsApi* | [**ListPermissionsAggregateInCollection**](docs/PermissionsApi.md#listpermissionsaggregateincollection) | **Get** /collections/{collectionId}/permissions/aggregate | List permissions for all users in a collection.
*PermissionsApi* | [**ListPermissionsAggregateInSpace**](docs/PermissionsApi.md#listpermissionsaggregateinspace) | **Get** /spaces/{spaceId}/permissions/aggregate | List permissions for all users in a space.
*PermissionsApi* | [**ListSpacesForOrganizationMember**](docs/PermissionsApi.md#listspacesfororganizationmember) | **Get** /orgs/{organizationId}/members/{userId}/spaces | List permissions accross all spaces for a member of an organization
*SearchApi* | [**SearchContent**](docs/SearchApi.md#searchcontent) | **Get** /search | Search content across spaces that is accessible by the currently authenticated target
*SpacesApi* | [**AddSpaceEntities**](docs/SpacesApi.md#addspaceentities) | **Post** /spaces/{spaceId}/entities | Link the space to entities
*SpacesApi* | [**CreateChangeRequest**](docs/SpacesApi.md#createchangerequest) | **Post** /spaces/{spaceId}/change-requests | Create a new change request for a space.
*SpacesApi* | [**CreateSpace**](docs/SpacesApi.md#createspace) | **Post** /orgs/{organizationId}/spaces | Create an organization space
*SpacesApi* | [**CreateSpaceRelation**](docs/SpacesApi.md#createspacerelation) | **Post** /spaces/{spaceId}/relations | Create a new relation between a source space and a target space
*SpacesApi* | [**DeleteCommentInChangeRequest**](docs/SpacesApi.md#deletecommentinchangerequest) | **Delete** /spaces/{spaceId}/change-requests/{changeRequestId}/comments/{commentId} | Delete a comment in a change request.
*SpacesApi* | [**DeleteCommentInSpace**](docs/SpacesApi.md#deletecommentinspace) | **Delete** /spaces/{spaceId}/comments/{commentId} | Delete a comment in a space.
*SpacesApi* | [**DeleteCommentReplyInChangeRequest**](docs/SpacesApi.md#deletecommentreplyinchangerequest) | **Delete** /spaces/{spaceId}/change-requests/{changeRequestId}/comments/{commentId}/replies/{commentReplyId} | Delete a comment reply in a change request.
*SpacesApi* | [**DeleteCommentReplyInSpace**](docs/SpacesApi.md#deletecommentreplyinspace) | **Delete** /spaces/{spaceId}/comments/{commentId}/replies/{commentReplyId} | Delete a comment reply in a space.
*SpacesApi* | [**DeleteSpaceEntity**](docs/SpacesApi.md#deletespaceentity) | **Delete** /spaces/{spaceId}/entities/{integrationName}/{entityId} | Delete a space entity
*SpacesApi* | [**DeleteSpaceRelation**](docs/SpacesApi.md#deletespacerelation) | **Delete** /spaces/{spaceId}/relations/{targetSpaceId} | Delete a relation between spaces
*SpacesApi* | [**DuplicateSpace**](docs/SpacesApi.md#duplicatespace) | **Post** /spaces/{spaceId}/duplicate | Create a duplicate of the space.
*SpacesApi* | [**ExportToGitRepository**](docs/SpacesApi.md#exporttogitrepository) | **Post** /spaces/{spaceId}/git/export | Export the space content to a Git repository.
*SpacesApi* | [**GetChangeRequestById**](docs/SpacesApi.md#getchangerequestbyid) | **Get** /spaces/{spaceId}/change-requests/{changeRequestId} | Get the change request with the given id.
*SpacesApi* | [**GetCommentInChangeRequest**](docs/SpacesApi.md#getcommentinchangerequest) | **Get** /spaces/{spaceId}/change-requests/{changeRequestId}/comments/{commentId} | Get a comment in a change request.
*SpacesApi* | [**GetCommentInSpace**](docs/SpacesApi.md#getcommentinspace) | **Get** /spaces/{spaceId}/comments/{commentId} | Get a comment in a space.
*SpacesApi* | [**GetCommentReplyInChangeRequest**](docs/SpacesApi.md#getcommentreplyinchangerequest) | **Get** /spaces/{spaceId}/change-requests/{changeRequestId}/comments/{commentId}/replies/{commentReplyId} | Get a comment reply in a change request.
*SpacesApi* | [**GetCommentReplyInSpace**](docs/SpacesApi.md#getcommentreplyinspace) | **Get** /spaces/{spaceId}/comments/{commentId}/replies/{commentReplyId} | Get a comment reply in a space.
*SpacesApi* | [**GetContentAnalyticsForSpaceById**](docs/SpacesApi.md#getcontentanalyticsforspacebyid) | **Get** /spaces/{spaceId}/insights/content | Get content analytics for a given space.
*SpacesApi* | [**GetContributorsByChangeRequestId**](docs/SpacesApi.md#getcontributorsbychangerequestid) | **Get** /spaces/{spaceId}/change-requests/{changeRequestId}/contributors | Get all contributors for the change request with the given id.
*SpacesApi* | [**GetRequestedReviewersByChangeRequestId**](docs/SpacesApi.md#getrequestedreviewersbychangerequestid) | **Get** /spaces/{spaceId}/change-requests/{changeRequestId}/requested-reviewers | Get all requested reviewers for a change request.
*SpacesApi* | [**GetReviewsByChangeRequestId**](docs/SpacesApi.md#getreviewsbychangerequestid) | **Get** /spaces/{spaceId}/change-requests/{changeRequestId}/reviews | Get all reviews for a change request.
*SpacesApi* | [**GetSearchAnalyticsForSpaceById**](docs/SpacesApi.md#getsearchanalyticsforspacebyid) | **Get** /spaces/{spaceId}/insights/search | Get an overview of the top search queries in a space.
*SpacesApi* | [**GetSpaceById**](docs/SpacesApi.md#getspacebyid) | **Get** /spaces/{spaceId} | Get the details about a space.
*SpacesApi* | [**GetSpaceCustomFields**](docs/SpacesApi.md#getspacecustomfields) | **Get** /spaces/{spaceId}/custom-fields | Get the values of the custom fields set on a space
*SpacesApi* | [**GetSpaceEntity**](docs/SpacesApi.md#getspaceentity) | **Get** /spaces/{spaceId}/entities/{integrationName}/{entityId} | Get a space entity
*SpacesApi* | [**GetSpaceGitInfo**](docs/SpacesApi.md#getspacegitinfo) | **Get** /spaces/{spaceId}/git/info | Get metadata about the Git Sync provider currently set up on the space
*SpacesApi* | [**GetSpacePublishingAuthById**](docs/SpacesApi.md#getspacepublishingauthbyid) | **Get** /spaces/{spaceId}/publishing/auth | Get the publishing authentication configuration for a space.
*SpacesApi* | [**GetSpaceRelation**](docs/SpacesApi.md#getspacerelation) | **Get** /spaces/{spaceId}/relations/{targetSpaceId} | Get the relation between 2 spaces.
*SpacesApi* | [**GetTrafficAnalyticsForSpaceById**](docs/SpacesApi.md#gettrafficanalyticsforspacebyid) | **Get** /spaces/{spaceId}/insights/traffic | Get traffic page views for a given space
*SpacesApi* | [**ImportGitRepository**](docs/SpacesApi.md#importgitrepository) | **Post** /spaces/{spaceId}/git/import | Import a Git repository
*SpacesApi* | [**ListChangeRequestsForSpace**](docs/SpacesApi.md#listchangerequestsforspace) | **Get** /spaces/{spaceId}/change-requests | List change requests for a space.
*SpacesApi* | [**ListCommentRepliesInChangeRequest**](docs/SpacesApi.md#listcommentrepliesinchangerequest) | **Get** /spaces/{spaceId}/change-requests/{changeRequestId}/comments/{commentId}/replies | List all the replies to a comment in a change request.
*SpacesApi* | [**ListCommentRepliesInSpace**](docs/SpacesApi.md#listcommentrepliesinspace) | **Get** /spaces/{spaceId}/comments/{commentId}/replies | List all the replies to a comment in a space.
*SpacesApi* | [**ListCommentsInChangeRequest**](docs/SpacesApi.md#listcommentsinchangerequest) | **Get** /spaces/{spaceId}/change-requests/{changeRequestId}/comments | List all the comments in a change request.
*SpacesApi* | [**ListCommentsInSpace**](docs/SpacesApi.md#listcommentsinspace) | **Get** /spaces/{spaceId}/comments | List all the comments in a space.
*SpacesApi* | [**ListEntitiesInOrganization**](docs/SpacesApi.md#listentitiesinorganization) | **Get** /orgs/{organizationId}/entities | List all entities in an organization.
*SpacesApi* | [**ListPermissionsAggregateInCollection**](docs/SpacesApi.md#listpermissionsaggregateincollection) | **Get** /collections/{collectionId}/permissions/aggregate | List permissions for all users in a collection.
*SpacesApi* | [**ListPermissionsAggregateInSpace**](docs/SpacesApi.md#listpermissionsaggregateinspace) | **Get** /spaces/{spaceId}/permissions/aggregate | List permissions for all users in a space.
*SpacesApi* | [**ListSpaceEntities**](docs/SpacesApi.md#listspaceentities) | **Get** /spaces/{spaceId}/entities | List all entities linked to a space
*SpacesApi* | [**ListSpaceRelations**](docs/SpacesApi.md#listspacerelations) | **Get** /spaces/{spaceId}/relations | List all relations for a space
*SpacesApi* | [**ListSpaceRelationsInOrganization**](docs/SpacesApi.md#listspacerelationsinorganization) | **Get** /orgs/{organizationId}/space-relations | List all relations between spaces in an organization
*SpacesApi* | [**ListSpacesForOrganizationMember**](docs/SpacesApi.md#listspacesfororganizationmember) | **Get** /orgs/{organizationId}/members/{userId}/spaces | List permissions accross all spaces for a member of an organization
*SpacesApi* | [**MergeChangeRequest**](docs/SpacesApi.md#mergechangerequest) | **Post** /spaces/{spaceId}/change-requests/{changeRequestId}/merge | Merge a change request in the primary content of a space.
*SpacesApi* | [**PostCommentInChangeRequest**](docs/SpacesApi.md#postcommentinchangerequest) | **Post** /spaces/{spaceId}/change-requests/{changeRequestId}/comments | Post a new comment in a change request.
*SpacesApi* | [**PostCommentInSpace**](docs/SpacesApi.md#postcommentinspace) | **Post** /spaces/{spaceId}/comments | Post a new comment in a space.
*SpacesApi* | [**PostCommentReplyInChangeRequest**](docs/SpacesApi.md#postcommentreplyinchangerequest) | **Post** /spaces/{spaceId}/change-requests/{changeRequestId}/comments/{commentId}/replies | Post a new reply to a comment in a change request.
*SpacesApi* | [**PostCommentReplyInSpace**](docs/SpacesApi.md#postcommentreplyinspace) | **Post** /spaces/{spaceId}/comments/{commentId}/replies | Post a new reply to a comment in a space.
*SpacesApi* | [**RequestReviewersForChangeRequest**](docs/SpacesApi.md#requestreviewersforchangerequest) | **Post** /spaces/{spaceId}/change-requests/{changeRequestId}/requested-reviewers | Request reviewers on a change request. Note that requesting a review from teams is not yet supported.
*SpacesApi* | [**SubmitChangeRequestReview**](docs/SpacesApi.md#submitchangerequestreview) | **Post** /spaces/{spaceId}/change-requests/{changeRequestId}/reviews | Submit a review for a change request.
*SpacesApi* | [**TrackViewInSpaceById**](docs/SpacesApi.md#trackviewinspacebyid) | **Post** /spaces/{spaceId}/insights/track_view | 
*SpacesApi* | [**UpdateChangeRequest**](docs/SpacesApi.md#updatechangerequest) | **Post** /spaces/{spaceId}/change-requests/{changeRequestId}/update | Update a change-request with changes from primary content.
*SpacesApi* | [**UpdateCommentInChangeRequest**](docs/SpacesApi.md#updatecommentinchangerequest) | **Put** /spaces/{spaceId}/change-requests/{changeRequestId}/comments/{commentId} | Update a comment in a change request.
*SpacesApi* | [**UpdateCommentInSpace**](docs/SpacesApi.md#updatecommentinspace) | **Put** /spaces/{spaceId}/comments/{commentId} | Update a comment in a space.
*SpacesApi* | [**UpdateCommentReplyInChangeRequest**](docs/SpacesApi.md#updatecommentreplyinchangerequest) | **Put** /spaces/{spaceId}/change-requests/{changeRequestId}/comments/{commentId}/replies/{commentReplyId} | Update a comment reply in a change request.
*SpacesApi* | [**UpdateCommentReplyInSpace**](docs/SpacesApi.md#updatecommentreplyinspace) | **Put** /spaces/{spaceId}/comments/{commentId}/replies/{commentReplyId} | Update a comment reply in a space.
*SpacesApi* | [**UpdateSpaceById**](docs/SpacesApi.md#updatespacebyid) | **Patch** /spaces/{spaceId} | Update the details of a space
*SpacesApi* | [**UpdateSpaceCustomFields**](docs/SpacesApi.md#updatespacecustomfields) | **Patch** /spaces/{spaceId}/custom-fields | Update the custom fields in a space
*SpacesApi* | [**UpdateSpaceEntity**](docs/SpacesApi.md#updatespaceentity) | **Patch** /spaces/{spaceId}/entities/{integrationName}/{entityId} | Update a space entity
*SpacesApi* | [**UpdateSpacePublishingAuthById**](docs/SpacesApi.md#updatespacepublishingauthbyid) | **Post** /spaces/{spaceId}/publishing/auth | Update the publishing authentication configuration for a space.
*TeamsApi* | [**ListTeamsForOrganizationMember**](docs/TeamsApi.md#listteamsfororganizationmember) | **Get** /orgs/{organizationId}/members/{userId}/teams | List all teams an organization member is part of
*UrlsApi* | [**GetContentByUrl**](docs/UrlsApi.md#getcontentbyurl) | **Get** /urls/content | Resolve a URL to a content (space, collection, page)
*UsersApi* | [**GetAuthenticatedUser**](docs/UsersApi.md#getauthenticateduser) | **Get** /user | Get profile of authenticated user
*UsersApi* | [**GetUserById**](docs/UsersApi.md#getuserbyid) | **Get** /users/{userId} | Get a user by its ID
*UsersApi* | [**ListSpacesForAuthenticatedUser**](docs/UsersApi.md#listspacesforauthenticateduser) | **Get** /user/spaces | List spaces for the authenticated user


## Documentation For Models

 - [APIIntegrationScope](docs/APIIntegrationScope.md)
 - [APIScope](docs/APIScope.md)
 - [APITemporaryToken](docs/APITemporaryToken.md)
 - [AddMemberToOrganizationTeamByIdRequest](docs/AddMemberToOrganizationTeamByIdRequest.md)
 - [AddSpaceEntitiesRequest](docs/AddSpaceEntitiesRequest.md)
 - [AddSpaceEntitiesRequestEntitiesInner](docs/AddSpaceEntitiesRequestEntitiesInner.md)
 - [AnalyticsContentPage](docs/AnalyticsContentPage.md)
 - [AnalyticsContentPageFeedbacks](docs/AnalyticsContentPageFeedbacks.md)
 - [AnalyticsContentPages](docs/AnalyticsContentPages.md)
 - [AnalyticsSearchPeriod](docs/AnalyticsSearchPeriod.md)
 - [AnalyticsSearchQuery](docs/AnalyticsSearchQuery.md)
 - [AnalyticsTopSearches](docs/AnalyticsTopSearches.md)
 - [AnalyticsTrafficInterval](docs/AnalyticsTrafficInterval.md)
 - [AnalyticsTrafficPageViews](docs/AnalyticsTrafficPageViews.md)
 - [AnalyticsTrafficPageViewsViewsInner](docs/AnalyticsTrafficPageViewsViewsInner.md)
 - [ApiInformation](docs/ApiInformation.md)
 - [AskQueryWithGet200Response](docs/AskQueryWithGet200Response.md)
 - [BackofficeUserInfoChannel](docs/BackofficeUserInfoChannel.md)
 - [BaseEvent](docs/BaseEvent.md)
 - [BillingInterval](docs/BillingInterval.md)
 - [BillingInvoicePreview](docs/BillingInvoicePreview.md)
 - [BillingInvoicePreviewLinesInner](docs/BillingInvoicePreviewLinesInner.md)
 - [BillingPortal](docs/BillingPortal.md)
 - [BillingProduct](docs/BillingProduct.md)
 - [BillingUpgrade](docs/BillingUpgrade.md)
 - [BillingUpgradeOneOf](docs/BillingUpgradeOneOf.md)
 - [BillingUpgradeOneOf1](docs/BillingUpgradeOneOf1.md)
 - [BillingUpgradeOneOf2](docs/BillingUpgradeOneOf2.md)
 - [BillingUpgradeOneOf3](docs/BillingUpgradeOneOf3.md)
 - [ChangeRequest](docs/ChangeRequest.md)
 - [ChangeRequestRequestedReviewer](docs/ChangeRequestRequestedReviewer.md)
 - [ChangeRequestRequestedReviewerAllOf](docs/ChangeRequestRequestedReviewerAllOf.md)
 - [ChangeRequestReview](docs/ChangeRequestReview.md)
 - [ChangeRequestReviewStatus](docs/ChangeRequestReviewStatus.md)
 - [ChangeRequestReviewsChannel](docs/ChangeRequestReviewsChannel.md)
 - [ChangeRequestStatus](docs/ChangeRequestStatus.md)
 - [ChangeRequestUrls](docs/ChangeRequestUrls.md)
 - [CloudflareHostnameStatus](docs/CloudflareHostnameStatus.md)
 - [CloudflareHostnameTLSCertificate](docs/CloudflareHostnameTLSCertificate.md)
 - [CloudflareHostnameTLSInfo](docs/CloudflareHostnameTLSInfo.md)
 - [CloudflareHostnameTLSStatus](docs/CloudflareHostnameTLSStatus.md)
 - [CloudflareHostnameTLSValidationError](docs/CloudflareHostnameTLSValidationError.md)
 - [CloudflareHostnameTLSValidationMethod](docs/CloudflareHostnameTLSValidationMethod.md)
 - [Collection](docs/Collection.md)
 - [Comment](docs/Comment.md)
 - [CommentAllOf](docs/CommentAllOf.md)
 - [CommentAllOfTarget](docs/CommentAllOfTarget.md)
 - [CommentAllOfTargetNode](docs/CommentAllOfTargetNode.md)
 - [CommentAllOfUrls](docs/CommentAllOfUrls.md)
 - [CommentReply](docs/CommentReply.md)
 - [CommentReplyUrls](docs/CommentReplyUrls.md)
 - [ContentKitAction](docs/ContentKitAction.md)
 - [ContentKitActionAnyOf](docs/ContentKitActionAnyOf.md)
 - [ContentKitBlock](docs/ContentKitBlock.md)
 - [ContentKitBlockControl](docs/ContentKitBlockControl.md)
 - [ContentKitBlockControlsInner](docs/ContentKitBlockControlsInner.md)
 - [ContentKitBox](docs/ContentKitBox.md)
 - [ContentKitButton](docs/ContentKitButton.md)
 - [ContentKitCard](docs/ContentKitCard.md)
 - [ContentKitCardHint](docs/ContentKitCardHint.md)
 - [ContentKitCardIcon](docs/ContentKitCardIcon.md)
 - [ContentKitCheckbox](docs/ContentKitCheckbox.md)
 - [ContentKitCheckboxValue](docs/ContentKitCheckboxValue.md)
 - [ContentKitCodeBlock](docs/ContentKitCodeBlock.md)
 - [ContentKitCodeBlockContent](docs/ContentKitCodeBlockContent.md)
 - [ContentKitCodeBlockLineNumbers](docs/ContentKitCodeBlockLineNumbers.md)
 - [ContentKitConfirm](docs/ContentKitConfirm.md)
 - [ContentKitContext](docs/ContentKitContext.md)
 - [ContentKitContextBase](docs/ContentKitContextBase.md)
 - [ContentKitContextConfigurationAccount](docs/ContentKitContextConfigurationAccount.md)
 - [ContentKitContextConfigurationAccountAllOf](docs/ContentKitContextConfigurationAccountAllOf.md)
 - [ContentKitContextConfigurationSpace](docs/ContentKitContextConfigurationSpace.md)
 - [ContentKitContextConfigurationSpaceAllOf](docs/ContentKitContextConfigurationSpaceAllOf.md)
 - [ContentKitContextDocument](docs/ContentKitContextDocument.md)
 - [ContentKitContextDocumentAllOf](docs/ContentKitContextDocumentAllOf.md)
 - [ContentKitDefaultAction](docs/ContentKitDefaultAction.md)
 - [ContentKitDefaultActionOneOf](docs/ContentKitDefaultActionOneOf.md)
 - [ContentKitDefaultActionOneOf1](docs/ContentKitDefaultActionOneOf1.md)
 - [ContentKitDefaultActionOneOf2](docs/ContentKitDefaultActionOneOf2.md)
 - [ContentKitDefaultActionOneOf3](docs/ContentKitDefaultActionOneOf3.md)
 - [ContentKitDefaultActionOneOf4](docs/ContentKitDefaultActionOneOf4.md)
 - [ContentKitDescendantElement](docs/ContentKitDescendantElement.md)
 - [ContentKitDivider](docs/ContentKitDivider.md)
 - [ContentKitDynamicBinding](docs/ContentKitDynamicBinding.md)
 - [ContentKitHStack](docs/ContentKitHStack.md)
 - [ContentKitIcon](docs/ContentKitIcon.md)
 - [ContentKitImage](docs/ContentKitImage.md)
 - [ContentKitImageSource](docs/ContentKitImageSource.md)
 - [ContentKitInlineElement](docs/ContentKitInlineElement.md)
 - [ContentKitInput](docs/ContentKitInput.md)
 - [ContentKitInputElement](docs/ContentKitInputElement.md)
 - [ContentKitInputHint](docs/ContentKitInputHint.md)
 - [ContentKitMarkdown](docs/ContentKitMarkdown.md)
 - [ContentKitMarkdownContent](docs/ContentKitMarkdownContent.md)
 - [ContentKitModal](docs/ContentKitModal.md)
 - [ContentKitRadio](docs/ContentKitRadio.md)
 - [ContentKitRadioValue](docs/ContentKitRadioValue.md)
 - [ContentKitRenderOutput](docs/ContentKitRenderOutput.md)
 - [ContentKitRootElement](docs/ContentKitRootElement.md)
 - [ContentKitSelect](docs/ContentKitSelect.md)
 - [ContentKitSelectInitialValue](docs/ContentKitSelectInitialValue.md)
 - [ContentKitSelectOption](docs/ContentKitSelectOption.md)
 - [ContentKitSelectOptions](docs/ContentKitSelectOptions.md)
 - [ContentKitSelectOptionsOneOf](docs/ContentKitSelectOptionsOneOf.md)
 - [ContentKitSwitch](docs/ContentKitSwitch.md)
 - [ContentKitText](docs/ContentKitText.md)
 - [ContentKitTextChildrenInner](docs/ContentKitTextChildrenInner.md)
 - [ContentKitTextInput](docs/ContentKitTextInput.md)
 - [ContentKitVStack](docs/ContentKitVStack.md)
 - [ContentKitWebFrame](docs/ContentKitWebFrame.md)
 - [ContentKitWebFrameDataValue](docs/ContentKitWebFrameDataValue.md)
 - [ContentKitWebFrameSource](docs/ContentKitWebFrameSource.md)
 - [ContentPublishingAuth](docs/ContentPublishingAuth.md)
 - [ContentVisibility](docs/ContentVisibility.md)
 - [CreateChangeRequest201Response](docs/CreateChangeRequest201Response.md)
 - [CreateChangeRequest201ResponseAllOf](docs/CreateChangeRequest201ResponseAllOf.md)
 - [CreateEnvironment](docs/CreateEnvironment.md)
 - [CreateOrganizationCustomFieldRequest](docs/CreateOrganizationCustomFieldRequest.md)
 - [CreateOrganizationTeamRequest](docs/CreateOrganizationTeamRequest.md)
 - [CreateSpaceRelationRequest](docs/CreateSpaceRelationRequest.md)
 - [CustomDomainInfo](docs/CustomDomainInfo.md)
 - [CustomField](docs/CustomField.md)
 - [CustomFieldType](docs/CustomFieldType.md)
 - [CustomFieldUrls](docs/CustomFieldUrls.md)
 - [CustomFieldValue](docs/CustomFieldValue.md)
 - [CustomFieldValuesInner](docs/CustomFieldValuesInner.md)
 - [Document](docs/Document.md)
 - [EmojiReaction](docs/EmojiReaction.md)
 - [EmojiReactionUsersInner](docs/EmojiReactionUsersInner.md)
 - [Environment](docs/Environment.md)
 - [EnvironmentUrls](docs/EnvironmentUrls.md)
 - [Error](docs/Error.md)
 - [ErrorError](docs/ErrorError.md)
 - [Event](docs/Event.md)
 - [FetchEvent](docs/FetchEvent.md)
 - [FetchEventAllOf](docs/FetchEventAllOf.md)
 - [FetchEventAllOfAuth](docs/FetchEventAllOfAuth.md)
 - [FetchPublishedScriptEvent](docs/FetchPublishedScriptEvent.md)
 - [FetchPublishedScriptEventAllOf](docs/FetchPublishedScriptEventAllOf.md)
 - [FetchRequest](docs/FetchRequest.md)
 - [FirebaseUserInfo](docs/FirebaseUserInfo.md)
 - [GenerateHiveAccessTokenRequest](docs/GenerateHiveAccessTokenRequest.md)
 - [GetChangeRequestByIdChangeRequestIdParameter](docs/GetChangeRequestByIdChangeRequestIdParameter.md)
 - [GetContentByUrl200Response](docs/GetContentByUrl200Response.md)
 - [GetContentByUrl200ResponseOneOf](docs/GetContentByUrl200ResponseOneOf.md)
 - [GetContentByUrl200ResponseOneOf1](docs/GetContentByUrl200ResponseOneOf1.md)
 - [GetContributorsByChangeRequestId200Response](docs/GetContributorsByChangeRequestId200Response.md)
 - [GetContributorsByChangeRequestId200ResponseAllOf](docs/GetContributorsByChangeRequestId200ResponseAllOf.md)
 - [GetIntegrationEntities200Response](docs/GetIntegrationEntities200Response.md)
 - [GetIntegrationEntities200ResponseAllOf](docs/GetIntegrationEntities200ResponseAllOf.md)
 - [GetIntegrationEvent200Response](docs/GetIntegrationEvent200Response.md)
 - [GetPageByPath200Response](docs/GetPageByPath200Response.md)
 - [GetRecommendedQuestionsRequest](docs/GetRecommendedQuestionsRequest.md)
 - [GetRequestedReviewersByChangeRequestId200Response](docs/GetRequestedReviewersByChangeRequestId200Response.md)
 - [GetRequestedReviewersByChangeRequestId200ResponseAllOf](docs/GetRequestedReviewersByChangeRequestId200ResponseAllOf.md)
 - [GetReviewsByChangeRequestId200Response](docs/GetReviewsByChangeRequestId200Response.md)
 - [GetReviewsByChangeRequestId200ResponseAllOf](docs/GetReviewsByChangeRequestId200ResponseAllOf.md)
 - [GetSpaceGitInfo404Response](docs/GetSpaceGitInfo404Response.md)
 - [GetSpaceGitInfo404ResponseError](docs/GetSpaceGitInfo404ResponseError.md)
 - [GetSpacePublishingAuthById400Response](docs/GetSpacePublishingAuthById400Response.md)
 - [GetSpacePublishingAuthById400ResponseError](docs/GetSpacePublishingAuthById400ResponseError.md)
 - [GitSyncOperation](docs/GitSyncOperation.md)
 - [GitSyncOperationState](docs/GitSyncOperationState.md)
 - [GitSyncState](docs/GitSyncState.md)
 - [HiveAccessToken](docs/HiveAccessToken.md)
 - [ImportContentResult](docs/ImportContentResult.md)
 - [ImportContentSource](docs/ImportContentSource.md)
 - [InstallIntegrationOnSpaceRequest](docs/InstallIntegrationOnSpaceRequest.md)
 - [InstallationEvent](docs/InstallationEvent.md)
 - [InstallationEventAllOf](docs/InstallationEventAllOf.md)
 - [InstallationSetupEvent](docs/InstallationSetupEvent.md)
 - [InstallationSetupEventAllOf](docs/InstallationSetupEventAllOf.md)
 - [InstallationSetupEventAllOfPrevious](docs/InstallationSetupEventAllOfPrevious.md)
 - [Integration](docs/Integration.md)
 - [IntegrationBlock](docs/IntegrationBlock.md)
 - [IntegrationBlockMarkdown](docs/IntegrationBlockMarkdown.md)
 - [IntegrationBlockMarkdownOneOf](docs/IntegrationBlockMarkdownOneOf.md)
 - [IntegrationCategory](docs/IntegrationCategory.md)
 - [IntegrationConfiguration](docs/IntegrationConfiguration.md)
 - [IntegrationConfigurationComponent](docs/IntegrationConfigurationComponent.md)
 - [IntegrationConfigurationSchema](docs/IntegrationConfigurationSchema.md)
 - [IntegrationConfigurationSchemaPropertiesValue](docs/IntegrationConfigurationSchemaPropertiesValue.md)
 - [IntegrationConfigurationSchemaPropertiesValueAllOf](docs/IntegrationConfigurationSchemaPropertiesValueAllOf.md)
 - [IntegrationConfigurationSchemaPropertiesValueAllOf1](docs/IntegrationConfigurationSchemaPropertiesValueAllOf1.md)
 - [IntegrationConfigurationSchemaPropertiesValueAllOf1OneOf](docs/IntegrationConfigurationSchemaPropertiesValueAllOf1OneOf.md)
 - [IntegrationConfigurationSchemaPropertiesValueAllOf1OneOf1](docs/IntegrationConfigurationSchemaPropertiesValueAllOf1OneOf1.md)
 - [IntegrationConfigurationSchemaPropertiesValueAllOf1OneOf2](docs/IntegrationConfigurationSchemaPropertiesValueAllOf1OneOf2.md)
 - [IntegrationConfigurationSchemaPropertiesValueAllOf1OneOf3](docs/IntegrationConfigurationSchemaPropertiesValueAllOf1OneOf3.md)
 - [IntegrationConfigurations](docs/IntegrationConfigurations.md)
 - [IntegrationContentSecurityPolicy](docs/IntegrationContentSecurityPolicy.md)
 - [IntegrationContentSecurityPolicyOneOf](docs/IntegrationContentSecurityPolicyOneOf.md)
 - [IntegrationEntity](docs/IntegrationEntity.md)
 - [IntegrationEntityAllOf](docs/IntegrationEntityAllOf.md)
 - [IntegrationEnvironment](docs/IntegrationEnvironment.md)
 - [IntegrationEvent](docs/IntegrationEvent.md)
 - [IntegrationEventLog](docs/IntegrationEventLog.md)
 - [IntegrationEventTrace](docs/IntegrationEventTrace.md)
 - [IntegrationExternalLinksInner](docs/IntegrationExternalLinksInner.md)
 - [IntegrationInstallation](docs/IntegrationInstallation.md)
 - [IntegrationInstallationSpaceSelection](docs/IntegrationInstallationSpaceSelection.md)
 - [IntegrationInstallationStatus](docs/IntegrationInstallationStatus.md)
 - [IntegrationInstallationTarget](docs/IntegrationInstallationTarget.md)
 - [IntegrationInstallationUrls](docs/IntegrationInstallationUrls.md)
 - [IntegrationScope](docs/IntegrationScope.md)
 - [IntegrationSpaceInstallation](docs/IntegrationSpaceInstallation.md)
 - [IntegrationSpaceInstallationUrls](docs/IntegrationSpaceInstallationUrls.md)
 - [IntegrationUrls](docs/IntegrationUrls.md)
 - [IntegrationVisibility](docs/IntegrationVisibility.md)
 - [InviteUsersToOrganization200Response](docs/InviteUsersToOrganization200Response.md)
 - [JSONDocument](docs/JSONDocument.md)
 - [JSONDocumentDocument](docs/JSONDocumentDocument.md)
 - [List](docs/List.md)
 - [ListChangeRequestsForSpace200Response](docs/ListChangeRequestsForSpace200Response.md)
 - [ListChangeRequestsForSpace200ResponseAllOf](docs/ListChangeRequestsForSpace200ResponseAllOf.md)
 - [ListCollectionsInOrganizationById200Response](docs/ListCollectionsInOrganizationById200Response.md)
 - [ListCollectionsInOrganizationById200ResponseAllOf](docs/ListCollectionsInOrganizationById200ResponseAllOf.md)
 - [ListCommentRepliesInChangeRequest200Response](docs/ListCommentRepliesInChangeRequest200Response.md)
 - [ListCommentRepliesInChangeRequest200ResponseAllOf](docs/ListCommentRepliesInChangeRequest200ResponseAllOf.md)
 - [ListCommentsInChangeRequest200Response](docs/ListCommentsInChangeRequest200Response.md)
 - [ListCommentsInChangeRequest200ResponseAllOf](docs/ListCommentsInChangeRequest200ResponseAllOf.md)
 - [ListDirectorySyncGroups200Response](docs/ListDirectorySyncGroups200Response.md)
 - [ListEnvironments200Response](docs/ListEnvironments200Response.md)
 - [ListEnvironments200ResponseAllOf](docs/ListEnvironments200ResponseAllOf.md)
 - [ListFiles200Response](docs/ListFiles200Response.md)
 - [ListFiles200ResponseAllOf](docs/ListFiles200ResponseAllOf.md)
 - [ListIntegrationEvents200Response](docs/ListIntegrationEvents200Response.md)
 - [ListIntegrationEvents200ResponseAllOf](docs/ListIntegrationEvents200ResponseAllOf.md)
 - [ListIntegrationInstallations200Response](docs/ListIntegrationInstallations200Response.md)
 - [ListIntegrationInstallations200ResponseAllOf](docs/ListIntegrationInstallations200ResponseAllOf.md)
 - [ListIntegrationSpaceInstallations200Response](docs/ListIntegrationSpaceInstallations200Response.md)
 - [ListIntegrationSpaceInstallations200ResponseAllOf](docs/ListIntegrationSpaceInstallations200ResponseAllOf.md)
 - [ListIntegrations200Response](docs/ListIntegrations200Response.md)
 - [ListIntegrations200ResponseAllOf](docs/ListIntegrations200ResponseAllOf.md)
 - [ListMembersInOrganizationById200Response](docs/ListMembersInOrganizationById200Response.md)
 - [ListMembersInOrganizationById200ResponseAllOf](docs/ListMembersInOrganizationById200ResponseAllOf.md)
 - [ListMembersInOrganizationByIdRoleParameter](docs/ListMembersInOrganizationByIdRoleParameter.md)
 - [ListNext](docs/ListNext.md)
 - [ListOrganizationCustomFields200Response](docs/ListOrganizationCustomFields200Response.md)
 - [ListOrganizationCustomFields200ResponseAllOf](docs/ListOrganizationCustomFields200ResponseAllOf.md)
 - [ListOrganizationsForAuthenticatedUser200Response](docs/ListOrganizationsForAuthenticatedUser200Response.md)
 - [ListOrganizationsForAuthenticatedUser200ResponseAllOf](docs/ListOrganizationsForAuthenticatedUser200ResponseAllOf.md)
 - [ListPermissionsAggregateInSpace200Response](docs/ListPermissionsAggregateInSpace200Response.md)
 - [ListPermissionsAggregateInSpace200ResponseAllOf](docs/ListPermissionsAggregateInSpace200ResponseAllOf.md)
 - [ListSpaceEntities200Response](docs/ListSpaceEntities200Response.md)
 - [ListSpaceEntities200ResponseAllOf](docs/ListSpaceEntities200ResponseAllOf.md)
 - [ListSpaceRelations200Response](docs/ListSpaceRelations200Response.md)
 - [ListSpaceRelations200ResponseAllOf](docs/ListSpaceRelations200ResponseAllOf.md)
 - [ListSpaceRelationsInOrganization200Response](docs/ListSpaceRelationsInOrganization200Response.md)
 - [ListSpaceRelationsInOrganization200ResponseAllOf](docs/ListSpaceRelationsInOrganization200ResponseAllOf.md)
 - [ListSpacesForAuthenticatedUser200Response](docs/ListSpacesForAuthenticatedUser200Response.md)
 - [ListSpacesForAuthenticatedUser200ResponseAllOf](docs/ListSpacesForAuthenticatedUser200ResponseAllOf.md)
 - [ListSpacesForOrganizationMember200Response](docs/ListSpacesForOrganizationMember200Response.md)
 - [ListSpacesForOrganizationMember200ResponseAllOf](docs/ListSpacesForOrganizationMember200ResponseAllOf.md)
 - [ListSpacesWithGitSyncInOrganizationById200Response](docs/ListSpacesWithGitSyncInOrganizationById200Response.md)
 - [ListSpacesWithGitSyncInOrganizationById200ResponseAllOf](docs/ListSpacesWithGitSyncInOrganizationById200ResponseAllOf.md)
 - [ListSpacesWithGitSyncInOrganizationById200ResponseAllOfItemsInner](docs/ListSpacesWithGitSyncInOrganizationById200ResponseAllOfItemsInner.md)
 - [ListTeamsForOrganizationMember200Response](docs/ListTeamsForOrganizationMember200Response.md)
 - [ListTeamsForOrganizationMember200ResponseAllOf](docs/ListTeamsForOrganizationMember200ResponseAllOf.md)
 - [ListTeamsForOrganizationMember200ResponseAllOfItemsInner](docs/ListTeamsForOrganizationMember200ResponseAllOfItemsInner.md)
 - [ListTeamsInOrganizationById200Response](docs/ListTeamsInOrganizationById200Response.md)
 - [ListTeamsInOrganizationById200ResponseAllOf](docs/ListTeamsInOrganizationById200ResponseAllOf.md)
 - [MarkdownDocument](docs/MarkdownDocument.md)
 - [MemberContentPermission](docs/MemberContentPermission.md)
 - [MemberRole](docs/MemberRole.md)
 - [MemberRoleOrGuest](docs/MemberRoleOrGuest.md)
 - [MergeChangeRequest200Response](docs/MergeChangeRequest200Response.md)
 - [Organization](docs/Organization.md)
 - [OrganizationCommunityType](docs/OrganizationCommunityType.md)
 - [OrganizationCustomFieldsChannel](docs/OrganizationCustomFieldsChannel.md)
 - [OrganizationDirectorySyncGroup](docs/OrganizationDirectorySyncGroup.md)
 - [OrganizationDirectorySyncGroupTeamStatus](docs/OrganizationDirectorySyncGroupTeamStatus.md)
 - [OrganizationDirectorySyncGroupTeamStatusAnyOf](docs/OrganizationDirectorySyncGroupTeamStatusAnyOf.md)
 - [OrganizationDirectorySyncGroupTeamStatusAnyOf1](docs/OrganizationDirectorySyncGroupTeamStatusAnyOf1.md)
 - [OrganizationEntitiesChannel](docs/OrganizationEntitiesChannel.md)
 - [OrganizationEnvironmentsChannel](docs/OrganizationEnvironmentsChannel.md)
 - [OrganizationMember](docs/OrganizationMember.md)
 - [OrganizationMemberChannel](docs/OrganizationMemberChannel.md)
 - [OrganizationMembersChannel](docs/OrganizationMembersChannel.md)
 - [OrganizationSpaceRelationsChannel](docs/OrganizationSpaceRelationsChannel.md)
 - [OrganizationSpacesChannel](docs/OrganizationSpacesChannel.md)
 - [OrganizationTarget](docs/OrganizationTarget.md)
 - [OrganizationTeam](docs/OrganizationTeam.md)
 - [OrganizationTeamChannel](docs/OrganizationTeamChannel.md)
 - [OrganizationTeamMemberChannel](docs/OrganizationTeamMemberChannel.md)
 - [OrganizationTeamMembersChannel](docs/OrganizationTeamMembersChannel.md)
 - [OrganizationTeamsChannel](docs/OrganizationTeamsChannel.md)
 - [OrganizationTransferResponse](docs/OrganizationTransferResponse.md)
 - [OrganizationType](docs/OrganizationType.md)
 - [OrganizationUrls](docs/OrganizationUrls.md)
 - [OrganizationUseCase](docs/OrganizationUseCase.md)
 - [PostCommentReplySchema](docs/PostCommentReplySchema.md)
 - [PostCommentSchema](docs/PostCommentSchema.md)
 - [PurgeCDNCacheContextType](docs/PurgeCDNCacheContextType.md)
 - [RequestBlockUserContext](docs/RequestBlockUserContext.md)
 - [RequestCreateChangeRequest](docs/RequestCreateChangeRequest.md)
 - [RequestCreateOrganization](docs/RequestCreateOrganization.md)
 - [RequestCreateSpace](docs/RequestCreateSpace.md)
 - [RequestExportToGitRepository](docs/RequestExportToGitRepository.md)
 - [RequestImportContent](docs/RequestImportContent.md)
 - [RequestImportGitRepository](docs/RequestImportGitRepository.md)
 - [RequestInviteUsersToOrganization](docs/RequestInviteUsersToOrganization.md)
 - [RequestInviteUsersToOrganizationEmailsInner](docs/RequestInviteUsersToOrganizationEmailsInner.md)
 - [RequestInviteUsersToOrganizationEmailsInnerOneOf](docs/RequestInviteUsersToOrganizationEmailsInnerOneOf.md)
 - [RequestPublishIntegration](docs/RequestPublishIntegration.md)
 - [RequestPurgeCDNCacheContext](docs/RequestPurgeCDNCacheContext.md)
 - [RequestRenderIntegrationUI](docs/RequestRenderIntegrationUI.md)
 - [RequestReviewersForChangeRequest200Response](docs/RequestReviewersForChangeRequest200Response.md)
 - [RequestReviewersForChangeRequestRequest](docs/RequestReviewersForChangeRequestRequest.md)
 - [RequestSpaceTrackPageView](docs/RequestSpaceTrackPageView.md)
 - [RequestSpaceTrackPageViewVisitor](docs/RequestSpaceTrackPageViewVisitor.md)
 - [RequestUpdateContentPublishingAuth](docs/RequestUpdateContentPublishingAuth.md)
 - [RequestUpdateIntegrationInstallation](docs/RequestUpdateIntegrationInstallation.md)
 - [RequestUpdateSpaceGitInfo](docs/RequestUpdateSpaceGitInfo.md)
 - [RequestUpgradeOrganizationBilling](docs/RequestUpgradeOrganizationBilling.md)
 - [Revision](docs/Revision.md)
 - [RevisionFile](docs/RevisionFile.md)
 - [RevisionGit](docs/RevisionGit.md)
 - [RevisionPage](docs/RevisionPage.md)
 - [RevisionPageBase](docs/RevisionPageBase.md)
 - [RevisionPageDocument](docs/RevisionPageDocument.md)
 - [RevisionPageDocumentAllOf](docs/RevisionPageDocumentAllOf.md)
 - [RevisionPageDocumentAllOfPagesInner](docs/RevisionPageDocumentAllOfPagesInner.md)
 - [RevisionPageGroup](docs/RevisionPageGroup.md)
 - [RevisionPageGroupAllOf](docs/RevisionPageGroupAllOf.md)
 - [RevisionPageLink](docs/RevisionPageLink.md)
 - [RevisionPageLinkAllOf](docs/RevisionPageLinkAllOf.md)
 - [RevisionUrls](docs/RevisionUrls.md)
 - [SearchAIAnswer](docs/SearchAIAnswer.md)
 - [SearchAIAnswerPagesInner](docs/SearchAIAnswerPagesInner.md)
 - [SearchAIQuery](docs/SearchAIQuery.md)
 - [SearchAIRecommendedQuestions](docs/SearchAIRecommendedQuestions.md)
 - [SearchContent200Response](docs/SearchContent200Response.md)
 - [SearchContent200ResponseAllOf](docs/SearchContent200ResponseAllOf.md)
 - [SearchPageResult](docs/SearchPageResult.md)
 - [SearchPageResultUrls](docs/SearchPageResultUrls.md)
 - [SearchSectionResult](docs/SearchSectionResult.md)
 - [SearchSectionResultUrls](docs/SearchSectionResultUrls.md)
 - [SearchSpaceContent200Response](docs/SearchSpaceContent200Response.md)
 - [SearchSpaceContent200ResponseAllOf](docs/SearchSpaceContent200ResponseAllOf.md)
 - [SearchSpaceResult](docs/SearchSpaceResult.md)
 - [SetupDirectorySync200Response](docs/SetupDirectorySync200Response.md)
 - [Space](docs/Space.md)
 - [SpaceContentUpdatedEvent](docs/SpaceContentUpdatedEvent.md)
 - [SpaceContentUpdatedEventAllOf](docs/SpaceContentUpdatedEventAllOf.md)
 - [SpaceCustomFieldsChannel](docs/SpaceCustomFieldsChannel.md)
 - [SpaceEntitiesChannel](docs/SpaceEntitiesChannel.md)
 - [SpaceEntity](docs/SpaceEntity.md)
 - [SpaceEntityUrls](docs/SpaceEntityUrls.md)
 - [SpaceEvent](docs/SpaceEvent.md)
 - [SpaceEventAllOf](docs/SpaceEventAllOf.md)
 - [SpaceGitInfoChannel](docs/SpaceGitInfoChannel.md)
 - [SpaceGitSyncCompletedEvent](docs/SpaceGitSyncCompletedEvent.md)
 - [SpaceGitSyncCompletedEventAllOf](docs/SpaceGitSyncCompletedEventAllOf.md)
 - [SpaceGitSyncStartedEvent](docs/SpaceGitSyncStartedEvent.md)
 - [SpaceGitSyncStartedEventAllOf](docs/SpaceGitSyncStartedEventAllOf.md)
 - [SpaceInfoChannel](docs/SpaceInfoChannel.md)
 - [SpaceInstallationSetupEvent](docs/SpaceInstallationSetupEvent.md)
 - [SpaceInstallationSetupEventAllOf](docs/SpaceInstallationSetupEventAllOf.md)
 - [SpaceInstallationSetupEventAllOfPrevious](docs/SpaceInstallationSetupEventAllOfPrevious.md)
 - [SpaceIntegrationBlocksInner](docs/SpaceIntegrationBlocksInner.md)
 - [SpaceIntegrationsChannel](docs/SpaceIntegrationsChannel.md)
 - [SpacePublishingAuthChannel](docs/SpacePublishingAuthChannel.md)
 - [SpaceRelation](docs/SpaceRelation.md)
 - [SpaceRelationEdge](docs/SpaceRelationEdge.md)
 - [SpaceRelationPair](docs/SpaceRelationPair.md)
 - [SpaceRelationPairAllOf](docs/SpaceRelationPairAllOf.md)
 - [SpaceRelationTarget](docs/SpaceRelationTarget.md)
 - [SpaceRelationType](docs/SpaceRelationType.md)
 - [SpaceRelationUrls](docs/SpaceRelationUrls.md)
 - [SpaceRelationsChannel](docs/SpaceRelationsChannel.md)
 - [SpaceType](docs/SpaceType.md)
 - [SpaceUrls](docs/SpaceUrls.md)
 - [SpaceViewEvent](docs/SpaceViewEvent.md)
 - [SpaceViewEventAllOf](docs/SpaceViewEventAllOf.md)
 - [SpaceViewEventAllOfVisitor](docs/SpaceViewEventAllOfVisitor.md)
 - [SpaceVisibilityUpdatedEvent](docs/SpaceVisibilityUpdatedEvent.md)
 - [SpaceVisibilityUpdatedEventAllOf](docs/SpaceVisibilityUpdatedEventAllOf.md)
 - [StaffUserInfo](docs/StaffUserInfo.md)
 - [SubmitChangeRequestReviewRequest](docs/SubmitChangeRequestReviewRequest.md)
 - [SubscriptionChannel](docs/SubscriptionChannel.md)
 - [SubscriptionChannelOneOf](docs/SubscriptionChannelOneOf.md)
 - [SubscriptionChannelOneOf1](docs/SubscriptionChannelOneOf1.md)
 - [SubscriptionChannelOneOf2](docs/SubscriptionChannelOneOf2.md)
 - [SubscriptionChannelOneOf3](docs/SubscriptionChannelOneOf3.md)
 - [SyncDirectorySyncGroupsToTeams200Response](docs/SyncDirectorySyncGroupsToTeams200Response.md)
 - [SyncDirectorySyncGroupsToTeamsRequest](docs/SyncDirectorySyncGroupsToTeamsRequest.md)
 - [SyncDirectorySyncGroupsToTeamsRequestToSyncInner](docs/SyncDirectorySyncGroupsToTeamsRequestToSyncInner.md)
 - [SyncIntegrationEntitiesRequest](docs/SyncIntegrationEntitiesRequest.md)
 - [Team](docs/Team.md)
 - [TeamMember](docs/TeamMember.md)
 - [TeamMemberRole](docs/TeamMemberRole.md)
 - [TransferOrganization200Response](docs/TransferOrganization200Response.md)
 - [TransferOrganizationRequest](docs/TransferOrganizationRequest.md)
 - [TransferOrganizationRequestDefaultOrgRole](docs/TransferOrganizationRequestDefaultOrgRole.md)
 - [TriggerContentIndexingContext](docs/TriggerContentIndexingContext.md)
 - [UIRenderEvent](docs/UIRenderEvent.md)
 - [UIRenderEventAllOf](docs/UIRenderEventAllOf.md)
 - [UnsplashImage](docs/UnsplashImage.md)
 - [UnsplashImageAuthor](docs/UnsplashImageAuthor.md)
 - [UnsplashImageUrls](docs/UnsplashImageUrls.md)
 - [UpdateChangeRequest200Response](docs/UpdateChangeRequest200Response.md)
 - [UpdateCollectionByIdRequest](docs/UpdateCollectionByIdRequest.md)
 - [UpdateCommentReplySchema](docs/UpdateCommentReplySchema.md)
 - [UpdateCommentSchema](docs/UpdateCommentSchema.md)
 - [UpdateCustomFieldValues](docs/UpdateCustomFieldValues.md)
 - [UpdateCustomFieldValuesValuesValue](docs/UpdateCustomFieldValuesValuesValue.md)
 - [UpdateEnvironment](docs/UpdateEnvironment.md)
 - [UpdateIntegrationDevSpaceRequest](docs/UpdateIntegrationDevSpaceRequest.md)
 - [UpdateMemberInOrganizationByIdRequest](docs/UpdateMemberInOrganizationByIdRequest.md)
 - [UpdateMembersInOrganizationTeamRequest](docs/UpdateMembersInOrganizationTeamRequest.md)
 - [UpdateOrganizationCustomFieldRequest](docs/UpdateOrganizationCustomFieldRequest.md)
 - [UpdateSpaceByIdRequest](docs/UpdateSpaceByIdRequest.md)
 - [UpdateSpaceEntityRequest](docs/UpdateSpaceEntityRequest.md)
 - [UpdateTeamInOrganizationByIdRequest](docs/UpdateTeamInOrganizationByIdRequest.md)
 - [UpsertIntegrationEntity](docs/UpsertIntegrationEntity.md)
 - [UpsertIntegrationEntityMetadataValue](docs/UpsertIntegrationEntityMetadataValue.md)
 - [User](docs/User.md)
 - [UserAPIToken](docs/UserAPIToken.md)
 - [UserAPITokenExtended](docs/UserAPITokenExtended.md)
 - [UserAPITokenExtendedAllOf](docs/UserAPITokenExtendedAllOf.md)
 - [UserAPITokensChannel](docs/UserAPITokensChannel.md)
 - [UserBackOfficeInfo](docs/UserBackOfficeInfo.md)
 - [UserContentPermission](docs/UserContentPermission.md)
 - [UserContributor](docs/UserContributor.md)
 - [UserImpersonation](docs/UserImpersonation.md)
 - [UserImpersonationAllOf](docs/UserImpersonationAllOf.md)
 - [UserImpersonationInfo](docs/UserImpersonationInfo.md)
 - [UserRiskEvaluation](docs/UserRiskEvaluation.md)
 - [UserTarget](docs/UserTarget.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### user

- **Type**: HTTP Bearer token authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "BEARER_TOKEN_STRING")
r, err := client.Service.Operation(auth, args)
```

### user-internal

- **Type**: HTTP Bearer token authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "BEARER_TOKEN_STRING")
r, err := client.Service.Operation(auth, args)
```

### user-staff

- **Type**: HTTP Bearer token authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "BEARER_TOKEN_STRING")
r, err := client.Service.Operation(auth, args)
```

### integration

- **Type**: HTTP Bearer token authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "BEARER_TOKEN_STRING")
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

---

© 2023 GitBook, Inc.
