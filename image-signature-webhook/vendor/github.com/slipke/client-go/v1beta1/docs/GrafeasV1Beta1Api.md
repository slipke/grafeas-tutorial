# \GrafeasV1Beta1Api

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchCreateNotes**](GrafeasV1Beta1Api.md#BatchCreateNotes) | **Post** /v1beta1/{parent}/notes:batchCreate | Creates new notes in batch.
[**BatchCreateOccurrences**](GrafeasV1Beta1Api.md#BatchCreateOccurrences) | **Post** /v1beta1/{parent}/occurrences:batchCreate | Creates new occurrences in batch.
[**CreateNote**](GrafeasV1Beta1Api.md#CreateNote) | **Post** /v1beta1/{parent}/notes | Creates a new note.
[**CreateOccurrence**](GrafeasV1Beta1Api.md#CreateOccurrence) | **Post** /v1beta1/{parent}/occurrences | Creates a new occurrence.
[**DeleteNote**](GrafeasV1Beta1Api.md#DeleteNote) | **Delete** /v1beta1/{name} | Deletes the specified note.
[**GetNote**](GrafeasV1Beta1Api.md#GetNote) | **Get** /v1beta1/{name} | Gets the specified note.
[**GetOccurrenceNote**](GrafeasV1Beta1Api.md#GetOccurrenceNote) | **Get** /v1beta1/{name}/notes | Gets the note attached to the specified occurrence. Consumer projects can use this method to get a note that belongs to a provider project.
[**GetVulnerabilityOccurrencesSummary**](GrafeasV1Beta1Api.md#GetVulnerabilityOccurrencesSummary) | **Get** /v1beta1/{parent}/occurrences:vulnerabilitySummary | Gets a summary of the number and severity of occurrences.
[**ListNoteOccurrences**](GrafeasV1Beta1Api.md#ListNoteOccurrences) | **Get** /v1beta1/{name}/occurrences | Lists occurrences referencing the specified note. Provider projects can use this method to get all occurrences across consumer projects referencing the specified note.
[**ListNotes**](GrafeasV1Beta1Api.md#ListNotes) | **Get** /v1beta1/{parent}/notes | Lists notes for the specified project.
[**ListOccurrences**](GrafeasV1Beta1Api.md#ListOccurrences) | **Get** /v1beta1/{parent}/occurrences | Lists occurrences for the specified project.
[**UpdateNote**](GrafeasV1Beta1Api.md#UpdateNote) | **Patch** /v1beta1/{name} | Updates the specified note.


# **BatchCreateNotes**
> V1beta1BatchCreateNotesResponse BatchCreateNotes(ctx, parent, body)
Creates new notes in batch.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **parent** | **string**| The name of the project in the form of &#x60;projects/[PROJECT_ID]&#x60;, under which the notes are to be created. | 
  **body** | [**V1beta1BatchCreateNotesRequest**](V1beta1BatchCreateNotesRequest.md)|  | 

### Return type

[**V1beta1BatchCreateNotesResponse**](v1beta1BatchCreateNotesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BatchCreateOccurrences**
> V1beta1BatchCreateOccurrencesResponse BatchCreateOccurrences(ctx, parent, body)
Creates new occurrences in batch.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **parent** | **string**| The name of the project in the form of &#x60;projects/[PROJECT_ID]&#x60;, under which the occurrences are to be created. | 
  **body** | [**V1beta1BatchCreateOccurrencesRequest**](V1beta1BatchCreateOccurrencesRequest.md)|  | 

### Return type

[**V1beta1BatchCreateOccurrencesResponse**](v1beta1BatchCreateOccurrencesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateNote**
> V1beta1Note CreateNote(ctx, parent, body)
Creates a new note.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **parent** | **string**| The name of the project in the form of &#x60;projects/[PROJECT_ID]&#x60;, under which the note is to be created. | 
  **body** | [**V1beta1Note**](V1beta1Note.md)| The note to create. | 

### Return type

[**V1beta1Note**](v1beta1Note.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOccurrence**
> V1beta1Occurrence CreateOccurrence(ctx, parent, body)
Creates a new occurrence.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **parent** | **string**| The name of the project in the form of &#x60;projects/[PROJECT_ID]&#x60;, under which the occurrence is to be created. | 
  **body** | [**V1beta1Occurrence**](V1beta1Occurrence.md)| The occurrence to create. | 

### Return type

[**V1beta1Occurrence**](v1beta1Occurrence.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteNote**
> ProtobufEmpty DeleteNote(ctx, name)
Deletes the specified note.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The name of the note in the form of &#x60;projects/[PROVIDER_ID]/notes/[NOTE_ID]&#x60;. | 

### Return type

[**ProtobufEmpty**](protobufEmpty.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNote**
> V1beta1Note GetNote(ctx, name)
Gets the specified note.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The name of the note in the form of &#x60;projects/[PROVIDER_ID]/notes/[NOTE_ID]&#x60;. | 

### Return type

[**V1beta1Note**](v1beta1Note.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOccurrenceNote**
> V1beta1Note GetOccurrenceNote(ctx, name)
Gets the note attached to the specified occurrence. Consumer projects can use this method to get a note that belongs to a provider project.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The name of the occurrence in the form of &#x60;projects/[PROJECT_ID]/occurrences/[OCCURRENCE_ID]&#x60;. | 

### Return type

[**V1beta1Note**](v1beta1Note.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVulnerabilityOccurrencesSummary**
> V1beta1VulnerabilityOccurrencesSummary GetVulnerabilityOccurrencesSummary(ctx, parent, optional)
Gets a summary of the number and severity of occurrences.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **parent** | **string**| The name of the project to get a vulnerability summary for in the form of &#x60;projects/[PROJECT_ID]&#x60;. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **parent** | **string**| The name of the project to get a vulnerability summary for in the form of &#x60;projects/[PROJECT_ID]&#x60;. | 
 **filter** | **string**| The filter expression. | 

### Return type

[**V1beta1VulnerabilityOccurrencesSummary**](v1beta1VulnerabilityOccurrencesSummary.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListNoteOccurrences**
> V1beta1ListNoteOccurrencesResponse ListNoteOccurrences(ctx, name, optional)
Lists occurrences referencing the specified note. Provider projects can use this method to get all occurrences across consumer projects referencing the specified note.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The name of the note to list occurrences for in the form of &#x60;projects/[PROVIDER_ID]/notes/[NOTE_ID]&#x60;. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The name of the note to list occurrences for in the form of &#x60;projects/[PROVIDER_ID]/notes/[NOTE_ID]&#x60;. | 
 **filter** | **string**| The filter expression. | 
 **pageSize** | **int32**| Number of occurrences to return in the list. | 
 **pageToken** | **string**| Token to provide to skip to a particular spot in the list. | 

### Return type

[**V1beta1ListNoteOccurrencesResponse**](v1beta1ListNoteOccurrencesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListNotes**
> V1beta1ListNotesResponse ListNotes(ctx, parent, optional)
Lists notes for the specified project.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **parent** | **string**| The name of the project to list notes for in the form of &#x60;projects/[PROJECT_ID]&#x60;. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **parent** | **string**| The name of the project to list notes for in the form of &#x60;projects/[PROJECT_ID]&#x60;. | 
 **filter** | **string**| The filter expression. | 
 **pageSize** | **int32**| Number of notes to return in the list. Must be positive. Max allowed page size is 1000. If not specified, page size defaults to 20. | 
 **pageToken** | **string**| Token to provide to skip to a particular spot in the list. | 

### Return type

[**V1beta1ListNotesResponse**](v1beta1ListNotesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOccurrences**
> V1beta1ListOccurrencesResponse ListOccurrences(ctx, parent, optional)
Lists occurrences for the specified project.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **parent** | **string**| The name of the project to list occurrences for in the form of &#x60;projects/[PROJECT_ID]&#x60;. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **parent** | **string**| The name of the project to list occurrences for in the form of &#x60;projects/[PROJECT_ID]&#x60;. | 
 **filter** | **string**| The filter expression. | 
 **pageSize** | **int32**| Number of occurrences to return in the list. Must be positive. Max allowed page size is 1000. If not specified, page size defaults to 20. | 
 **pageToken** | **string**| Token to provide to skip to a particular spot in the list. | 

### Return type

[**V1beta1ListOccurrencesResponse**](v1beta1ListOccurrencesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNote**
> V1beta1Note UpdateNote(ctx, name, body)
Updates the specified note.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The name of the note in the form of &#x60;projects/[PROVIDER_ID]/notes/[NOTE_ID]&#x60;. | 
  **body** | [**V1beta1Note**](V1beta1Note.md)| The updated note. | 

### Return type

[**V1beta1Note**](v1beta1Note.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

