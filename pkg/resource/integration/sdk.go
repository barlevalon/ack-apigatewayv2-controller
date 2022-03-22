// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package integration

import (
	"context"
	"reflect"
	"strings"

	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	ackcondition "github.com/aws-controllers-k8s/runtime/pkg/condition"
	ackerr "github.com/aws-controllers-k8s/runtime/pkg/errors"
	ackrtlog "github.com/aws-controllers-k8s/runtime/pkg/runtime/log"
	"github.com/aws/aws-sdk-go/aws"
	svcsdk "github.com/aws/aws-sdk-go/service/apigatewayv2"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	svcapitypes "github.com/aws-controllers-k8s/apigatewayv2-controller/apis/v1alpha1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
	_ = strings.ToLower("")
	_ = &aws.JSONValue{}
	_ = &svcsdk.ApiGatewayV2{}
	_ = &svcapitypes.Integration{}
	_ = ackv1alpha1.AWSAccountID("")
	_ = &ackerr.NotFound
	_ = &ackcondition.NotManagedMessage
	_ = &reflect.Value{}
)

// sdkFind returns SDK-specific information about a supplied resource
func (rm *resourceManager) sdkFind(
	ctx context.Context,
	r *resource,
) (latest *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkFind")
	defer exit(err)
	// If any required fields in the input shape are missing, AWS resource is
	// not created yet. Return NotFound here to indicate to callers that the
	// resource isn't yet created.
	if rm.requiredFieldsMissingFromReadOneInput(r) {
		return nil, ackerr.NotFound
	}

	input, err := rm.newDescribeRequestPayload(r)
	if err != nil {
		return nil, err
	}

	var resp *svcsdk.GetIntegrationOutput
	resp, err = rm.sdkapi.GetIntegrationWithContext(ctx, input)
	rm.metrics.RecordAPICall("READ_ONE", "GetIntegration", err)
	if err != nil {
		if awsErr, ok := ackerr.AWSError(err); ok && awsErr.Code() == "NotFoundException" {
			return nil, ackerr.NotFound
		}
		return nil, err
	}

	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	if resp.ApiGatewayManaged != nil {
		ko.Status.APIGatewayManaged = resp.ApiGatewayManaged
	} else {
		ko.Status.APIGatewayManaged = nil
	}
	if resp.ConnectionId != nil {
		ko.Spec.ConnectionID = resp.ConnectionId
	} else {
		ko.Spec.ConnectionID = nil
	}
	if resp.ConnectionType != nil {
		ko.Spec.ConnectionType = resp.ConnectionType
	} else {
		ko.Spec.ConnectionType = nil
	}
	if resp.ContentHandlingStrategy != nil {
		ko.Spec.ContentHandlingStrategy = resp.ContentHandlingStrategy
	} else {
		ko.Spec.ContentHandlingStrategy = nil
	}
	if resp.CredentialsArn != nil {
		ko.Spec.CredentialsARN = resp.CredentialsArn
	} else {
		ko.Spec.CredentialsARN = nil
	}
	if resp.Description != nil {
		ko.Spec.Description = resp.Description
	} else {
		ko.Spec.Description = nil
	}
	if resp.IntegrationId != nil {
		ko.Status.IntegrationID = resp.IntegrationId
	} else {
		ko.Status.IntegrationID = nil
	}
	if resp.IntegrationMethod != nil {
		ko.Spec.IntegrationMethod = resp.IntegrationMethod
	} else {
		ko.Spec.IntegrationMethod = nil
	}
	if resp.IntegrationResponseSelectionExpression != nil {
		ko.Status.IntegrationResponseSelectionExpression = resp.IntegrationResponseSelectionExpression
	} else {
		ko.Status.IntegrationResponseSelectionExpression = nil
	}
	if resp.IntegrationSubtype != nil {
		ko.Spec.IntegrationSubtype = resp.IntegrationSubtype
	} else {
		ko.Spec.IntegrationSubtype = nil
	}
	if resp.IntegrationType != nil {
		ko.Spec.IntegrationType = resp.IntegrationType
	} else {
		ko.Spec.IntegrationType = nil
	}
	if resp.IntegrationUri != nil {
		ko.Spec.IntegrationURI = resp.IntegrationUri
	} else {
		ko.Spec.IntegrationURI = nil
	}
	if resp.PassthroughBehavior != nil {
		ko.Spec.PassthroughBehavior = resp.PassthroughBehavior
	} else {
		ko.Spec.PassthroughBehavior = nil
	}
	if resp.PayloadFormatVersion != nil {
		ko.Spec.PayloadFormatVersion = resp.PayloadFormatVersion
	} else {
		ko.Spec.PayloadFormatVersion = nil
	}
	if resp.RequestParameters != nil {
		f14 := map[string]*string{}
		for f14key, f14valiter := range resp.RequestParameters {
			var f14val string
			f14val = *f14valiter
			f14[f14key] = &f14val
		}
		ko.Spec.RequestParameters = f14
	} else {
		ko.Spec.RequestParameters = nil
	}
	if resp.RequestTemplates != nil {
		f15 := map[string]*string{}
		for f15key, f15valiter := range resp.RequestTemplates {
			var f15val string
			f15val = *f15valiter
			f15[f15key] = &f15val
		}
		ko.Spec.RequestTemplates = f15
	} else {
		ko.Spec.RequestTemplates = nil
	}
	if resp.ResponseParameters != nil {
		f16 := map[string]map[string]*string{}
		for f16key, f16valiter := range resp.ResponseParameters {
			f16val := map[string]*string{}
			for f16valkey, f16valvaliter := range f16valiter {
				var f16valval string
				f16valval = *f16valvaliter
				f16val[f16valkey] = &f16valval
			}
			f16[f16key] = f16val
		}
		ko.Spec.ResponseParameters = f16
	} else {
		ko.Spec.ResponseParameters = nil
	}
	if resp.TemplateSelectionExpression != nil {
		ko.Spec.TemplateSelectionExpression = resp.TemplateSelectionExpression
	} else {
		ko.Spec.TemplateSelectionExpression = nil
	}
	if resp.TimeoutInMillis != nil {
		ko.Spec.TimeoutInMillis = resp.TimeoutInMillis
	} else {
		ko.Spec.TimeoutInMillis = nil
	}
	if resp.TlsConfig != nil {
		f19 := &svcapitypes.TLSConfigInput{}
		if resp.TlsConfig.ServerNameToVerify != nil {
			f19.ServerNameToVerify = resp.TlsConfig.ServerNameToVerify
		}
		ko.Spec.TLSConfig = f19
	} else {
		ko.Spec.TLSConfig = nil
	}

	rm.setStatusDefaults(ko)
	return &resource{ko}, nil
}

// requiredFieldsMissingFromReadOneInput returns true if there are any fields
// for the ReadOne Input shape that are required but not present in the
// resource's Spec or Status
func (rm *resourceManager) requiredFieldsMissingFromReadOneInput(
	r *resource,
) bool {
	return r.ko.Spec.APIID == nil || r.ko.Status.IntegrationID == nil

}

// newDescribeRequestPayload returns SDK-specific struct for the HTTP request
// payload of the Describe API call for the resource
func (rm *resourceManager) newDescribeRequestPayload(
	r *resource,
) (*svcsdk.GetIntegrationInput, error) {
	res := &svcsdk.GetIntegrationInput{}

	if r.ko.Spec.APIID != nil {
		res.SetApiId(*r.ko.Spec.APIID)
	}
	if r.ko.Status.IntegrationID != nil {
		res.SetIntegrationId(*r.ko.Status.IntegrationID)
	}

	return res, nil
}

// sdkCreate creates the supplied resource in the backend AWS service API and
// returns a copy of the resource with resource fields (in both Spec and
// Status) filled in with values from the CREATE API operation's Output shape.
func (rm *resourceManager) sdkCreate(
	ctx context.Context,
	desired *resource,
) (created *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkCreate")
	defer exit(err)
	input, err := rm.newCreateRequestPayload(ctx, desired)
	if err != nil {
		return nil, err
	}

	var resp *svcsdk.CreateIntegrationOutput
	_ = resp
	resp, err = rm.sdkapi.CreateIntegrationWithContext(ctx, input)
	rm.metrics.RecordAPICall("CREATE", "CreateIntegration", err)
	if err != nil {
		return nil, err
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := desired.ko.DeepCopy()

	if resp.ApiGatewayManaged != nil {
		ko.Status.APIGatewayManaged = resp.ApiGatewayManaged
	} else {
		ko.Status.APIGatewayManaged = nil
	}
	if resp.ConnectionId != nil {
		ko.Spec.ConnectionID = resp.ConnectionId
	} else {
		ko.Spec.ConnectionID = nil
	}
	if resp.ConnectionType != nil {
		ko.Spec.ConnectionType = resp.ConnectionType
	} else {
		ko.Spec.ConnectionType = nil
	}
	if resp.ContentHandlingStrategy != nil {
		ko.Spec.ContentHandlingStrategy = resp.ContentHandlingStrategy
	} else {
		ko.Spec.ContentHandlingStrategy = nil
	}
	if resp.CredentialsArn != nil {
		ko.Spec.CredentialsARN = resp.CredentialsArn
	} else {
		ko.Spec.CredentialsARN = nil
	}
	if resp.Description != nil {
		ko.Spec.Description = resp.Description
	} else {
		ko.Spec.Description = nil
	}
	if resp.IntegrationId != nil {
		ko.Status.IntegrationID = resp.IntegrationId
	} else {
		ko.Status.IntegrationID = nil
	}
	if resp.IntegrationMethod != nil {
		ko.Spec.IntegrationMethod = resp.IntegrationMethod
	} else {
		ko.Spec.IntegrationMethod = nil
	}
	if resp.IntegrationResponseSelectionExpression != nil {
		ko.Status.IntegrationResponseSelectionExpression = resp.IntegrationResponseSelectionExpression
	} else {
		ko.Status.IntegrationResponseSelectionExpression = nil
	}
	if resp.IntegrationSubtype != nil {
		ko.Spec.IntegrationSubtype = resp.IntegrationSubtype
	} else {
		ko.Spec.IntegrationSubtype = nil
	}
	if resp.IntegrationType != nil {
		ko.Spec.IntegrationType = resp.IntegrationType
	} else {
		ko.Spec.IntegrationType = nil
	}
	if resp.IntegrationUri != nil {
		ko.Spec.IntegrationURI = resp.IntegrationUri
	} else {
		ko.Spec.IntegrationURI = nil
	}
	if resp.PassthroughBehavior != nil {
		ko.Spec.PassthroughBehavior = resp.PassthroughBehavior
	} else {
		ko.Spec.PassthroughBehavior = nil
	}
	if resp.PayloadFormatVersion != nil {
		ko.Spec.PayloadFormatVersion = resp.PayloadFormatVersion
	} else {
		ko.Spec.PayloadFormatVersion = nil
	}
	if resp.RequestParameters != nil {
		f14 := map[string]*string{}
		for f14key, f14valiter := range resp.RequestParameters {
			var f14val string
			f14val = *f14valiter
			f14[f14key] = &f14val
		}
		ko.Spec.RequestParameters = f14
	} else {
		ko.Spec.RequestParameters = nil
	}
	if resp.RequestTemplates != nil {
		f15 := map[string]*string{}
		for f15key, f15valiter := range resp.RequestTemplates {
			var f15val string
			f15val = *f15valiter
			f15[f15key] = &f15val
		}
		ko.Spec.RequestTemplates = f15
	} else {
		ko.Spec.RequestTemplates = nil
	}
	if resp.ResponseParameters != nil {
		f16 := map[string]map[string]*string{}
		for f16key, f16valiter := range resp.ResponseParameters {
			f16val := map[string]*string{}
			for f16valkey, f16valvaliter := range f16valiter {
				var f16valval string
				f16valval = *f16valvaliter
				f16val[f16valkey] = &f16valval
			}
			f16[f16key] = f16val
		}
		ko.Spec.ResponseParameters = f16
	} else {
		ko.Spec.ResponseParameters = nil
	}
	if resp.TemplateSelectionExpression != nil {
		ko.Spec.TemplateSelectionExpression = resp.TemplateSelectionExpression
	} else {
		ko.Spec.TemplateSelectionExpression = nil
	}
	if resp.TimeoutInMillis != nil {
		ko.Spec.TimeoutInMillis = resp.TimeoutInMillis
	} else {
		ko.Spec.TimeoutInMillis = nil
	}
	if resp.TlsConfig != nil {
		f19 := &svcapitypes.TLSConfigInput{}
		if resp.TlsConfig.ServerNameToVerify != nil {
			f19.ServerNameToVerify = resp.TlsConfig.ServerNameToVerify
		}
		ko.Spec.TLSConfig = f19
	} else {
		ko.Spec.TLSConfig = nil
	}

	rm.setStatusDefaults(ko)
	return &resource{ko}, nil
}

// newCreateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Create API call for the resource
func (rm *resourceManager) newCreateRequestPayload(
	ctx context.Context,
	r *resource,
) (*svcsdk.CreateIntegrationInput, error) {
	res := &svcsdk.CreateIntegrationInput{}

	if r.ko.Spec.APIID != nil {
		res.SetApiId(*r.ko.Spec.APIID)
	}
	if r.ko.Spec.ConnectionID != nil {
		res.SetConnectionId(*r.ko.Spec.ConnectionID)
	}
	if r.ko.Spec.ConnectionType != nil {
		res.SetConnectionType(*r.ko.Spec.ConnectionType)
	}
	if r.ko.Spec.ContentHandlingStrategy != nil {
		res.SetContentHandlingStrategy(*r.ko.Spec.ContentHandlingStrategy)
	}
	if r.ko.Spec.CredentialsARN != nil {
		res.SetCredentialsArn(*r.ko.Spec.CredentialsARN)
	}
	if r.ko.Spec.Description != nil {
		res.SetDescription(*r.ko.Spec.Description)
	}
	if r.ko.Spec.IntegrationMethod != nil {
		res.SetIntegrationMethod(*r.ko.Spec.IntegrationMethod)
	}
	if r.ko.Spec.IntegrationSubtype != nil {
		res.SetIntegrationSubtype(*r.ko.Spec.IntegrationSubtype)
	}
	if r.ko.Spec.IntegrationType != nil {
		res.SetIntegrationType(*r.ko.Spec.IntegrationType)
	}
	if r.ko.Spec.IntegrationURI != nil {
		res.SetIntegrationUri(*r.ko.Spec.IntegrationURI)
	}
	if r.ko.Spec.PassthroughBehavior != nil {
		res.SetPassthroughBehavior(*r.ko.Spec.PassthroughBehavior)
	}
	if r.ko.Spec.PayloadFormatVersion != nil {
		res.SetPayloadFormatVersion(*r.ko.Spec.PayloadFormatVersion)
	}
	if r.ko.Spec.RequestParameters != nil {
		f12 := map[string]*string{}
		for f12key, f12valiter := range r.ko.Spec.RequestParameters {
			var f12val string
			f12val = *f12valiter
			f12[f12key] = &f12val
		}
		res.SetRequestParameters(f12)
	}
	if r.ko.Spec.RequestTemplates != nil {
		f13 := map[string]*string{}
		for f13key, f13valiter := range r.ko.Spec.RequestTemplates {
			var f13val string
			f13val = *f13valiter
			f13[f13key] = &f13val
		}
		res.SetRequestTemplates(f13)
	}
	if r.ko.Spec.ResponseParameters != nil {
		f14 := map[string]map[string]*string{}
		for f14key, f14valiter := range r.ko.Spec.ResponseParameters {
			f14val := map[string]*string{}
			for f14valkey, f14valvaliter := range f14valiter {
				var f14valval string
				f14valval = *f14valvaliter
				f14val[f14valkey] = &f14valval
			}
			f14[f14key] = f14val
		}
		res.SetResponseParameters(f14)
	}
	if r.ko.Spec.TemplateSelectionExpression != nil {
		res.SetTemplateSelectionExpression(*r.ko.Spec.TemplateSelectionExpression)
	}
	if r.ko.Spec.TimeoutInMillis != nil {
		res.SetTimeoutInMillis(*r.ko.Spec.TimeoutInMillis)
	}
	if r.ko.Spec.TLSConfig != nil {
		f17 := &svcsdk.TlsConfigInput{}
		if r.ko.Spec.TLSConfig.ServerNameToVerify != nil {
			f17.SetServerNameToVerify(*r.ko.Spec.TLSConfig.ServerNameToVerify)
		}
		res.SetTlsConfig(f17)
	}

	return res, nil
}

// sdkUpdate patches the supplied resource in the backend AWS service API and
// returns a new resource with updated fields.
func (rm *resourceManager) sdkUpdate(
	ctx context.Context,
	desired *resource,
	latest *resource,
	delta *ackcompare.Delta,
) (updated *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkUpdate")
	defer exit(err)
	input, err := rm.newUpdateRequestPayload(ctx, desired)
	if err != nil {
		return nil, err
	}

	var resp *svcsdk.UpdateIntegrationOutput
	_ = resp
	resp, err = rm.sdkapi.UpdateIntegrationWithContext(ctx, input)
	rm.metrics.RecordAPICall("UPDATE", "UpdateIntegration", err)
	if err != nil {
		return nil, err
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := desired.ko.DeepCopy()

	if resp.ApiGatewayManaged != nil {
		ko.Status.APIGatewayManaged = resp.ApiGatewayManaged
	} else {
		ko.Status.APIGatewayManaged = nil
	}
	if resp.ConnectionId != nil {
		ko.Spec.ConnectionID = resp.ConnectionId
	} else {
		ko.Spec.ConnectionID = nil
	}
	if resp.ConnectionType != nil {
		ko.Spec.ConnectionType = resp.ConnectionType
	} else {
		ko.Spec.ConnectionType = nil
	}
	if resp.ContentHandlingStrategy != nil {
		ko.Spec.ContentHandlingStrategy = resp.ContentHandlingStrategy
	} else {
		ko.Spec.ContentHandlingStrategy = nil
	}
	if resp.CredentialsArn != nil {
		ko.Spec.CredentialsARN = resp.CredentialsArn
	} else {
		ko.Spec.CredentialsARN = nil
	}
	if resp.Description != nil {
		ko.Spec.Description = resp.Description
	} else {
		ko.Spec.Description = nil
	}
	if resp.IntegrationId != nil {
		ko.Status.IntegrationID = resp.IntegrationId
	} else {
		ko.Status.IntegrationID = nil
	}
	if resp.IntegrationMethod != nil {
		ko.Spec.IntegrationMethod = resp.IntegrationMethod
	} else {
		ko.Spec.IntegrationMethod = nil
	}
	if resp.IntegrationResponseSelectionExpression != nil {
		ko.Status.IntegrationResponseSelectionExpression = resp.IntegrationResponseSelectionExpression
	} else {
		ko.Status.IntegrationResponseSelectionExpression = nil
	}
	if resp.IntegrationSubtype != nil {
		ko.Spec.IntegrationSubtype = resp.IntegrationSubtype
	} else {
		ko.Spec.IntegrationSubtype = nil
	}
	if resp.IntegrationType != nil {
		ko.Spec.IntegrationType = resp.IntegrationType
	} else {
		ko.Spec.IntegrationType = nil
	}
	if resp.IntegrationUri != nil {
		ko.Spec.IntegrationURI = resp.IntegrationUri
	} else {
		ko.Spec.IntegrationURI = nil
	}
	if resp.PassthroughBehavior != nil {
		ko.Spec.PassthroughBehavior = resp.PassthroughBehavior
	} else {
		ko.Spec.PassthroughBehavior = nil
	}
	if resp.PayloadFormatVersion != nil {
		ko.Spec.PayloadFormatVersion = resp.PayloadFormatVersion
	} else {
		ko.Spec.PayloadFormatVersion = nil
	}
	if resp.RequestParameters != nil {
		f14 := map[string]*string{}
		for f14key, f14valiter := range resp.RequestParameters {
			var f14val string
			f14val = *f14valiter
			f14[f14key] = &f14val
		}
		ko.Spec.RequestParameters = f14
	} else {
		ko.Spec.RequestParameters = nil
	}
	if resp.RequestTemplates != nil {
		f15 := map[string]*string{}
		for f15key, f15valiter := range resp.RequestTemplates {
			var f15val string
			f15val = *f15valiter
			f15[f15key] = &f15val
		}
		ko.Spec.RequestTemplates = f15
	} else {
		ko.Spec.RequestTemplates = nil
	}
	if resp.ResponseParameters != nil {
		f16 := map[string]map[string]*string{}
		for f16key, f16valiter := range resp.ResponseParameters {
			f16val := map[string]*string{}
			for f16valkey, f16valvaliter := range f16valiter {
				var f16valval string
				f16valval = *f16valvaliter
				f16val[f16valkey] = &f16valval
			}
			f16[f16key] = f16val
		}
		ko.Spec.ResponseParameters = f16
	} else {
		ko.Spec.ResponseParameters = nil
	}
	if resp.TemplateSelectionExpression != nil {
		ko.Spec.TemplateSelectionExpression = resp.TemplateSelectionExpression
	} else {
		ko.Spec.TemplateSelectionExpression = nil
	}
	if resp.TimeoutInMillis != nil {
		ko.Spec.TimeoutInMillis = resp.TimeoutInMillis
	} else {
		ko.Spec.TimeoutInMillis = nil
	}
	if resp.TlsConfig != nil {
		f19 := &svcapitypes.TLSConfigInput{}
		if resp.TlsConfig.ServerNameToVerify != nil {
			f19.ServerNameToVerify = resp.TlsConfig.ServerNameToVerify
		}
		ko.Spec.TLSConfig = f19
	} else {
		ko.Spec.TLSConfig = nil
	}

	rm.setStatusDefaults(ko)
	return &resource{ko}, nil
}

// newUpdateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Update API call for the resource
func (rm *resourceManager) newUpdateRequestPayload(
	ctx context.Context,
	r *resource,
) (*svcsdk.UpdateIntegrationInput, error) {
	res := &svcsdk.UpdateIntegrationInput{}

	if r.ko.Spec.APIID != nil {
		res.SetApiId(*r.ko.Spec.APIID)
	}
	if r.ko.Spec.ConnectionID != nil {
		res.SetConnectionId(*r.ko.Spec.ConnectionID)
	}
	if r.ko.Spec.ConnectionType != nil {
		res.SetConnectionType(*r.ko.Spec.ConnectionType)
	}
	if r.ko.Spec.ContentHandlingStrategy != nil {
		res.SetContentHandlingStrategy(*r.ko.Spec.ContentHandlingStrategy)
	}
	if r.ko.Spec.CredentialsARN != nil {
		res.SetCredentialsArn(*r.ko.Spec.CredentialsARN)
	}
	if r.ko.Spec.Description != nil {
		res.SetDescription(*r.ko.Spec.Description)
	}
	if r.ko.Status.IntegrationID != nil {
		res.SetIntegrationId(*r.ko.Status.IntegrationID)
	}
	if r.ko.Spec.IntegrationMethod != nil {
		res.SetIntegrationMethod(*r.ko.Spec.IntegrationMethod)
	}
	if r.ko.Spec.IntegrationSubtype != nil {
		res.SetIntegrationSubtype(*r.ko.Spec.IntegrationSubtype)
	}
	if r.ko.Spec.IntegrationType != nil {
		res.SetIntegrationType(*r.ko.Spec.IntegrationType)
	}
	if r.ko.Spec.IntegrationURI != nil {
		res.SetIntegrationUri(*r.ko.Spec.IntegrationURI)
	}
	if r.ko.Spec.PassthroughBehavior != nil {
		res.SetPassthroughBehavior(*r.ko.Spec.PassthroughBehavior)
	}
	if r.ko.Spec.PayloadFormatVersion != nil {
		res.SetPayloadFormatVersion(*r.ko.Spec.PayloadFormatVersion)
	}
	if r.ko.Spec.RequestParameters != nil {
		f13 := map[string]*string{}
		for f13key, f13valiter := range r.ko.Spec.RequestParameters {
			var f13val string
			f13val = *f13valiter
			f13[f13key] = &f13val
		}
		res.SetRequestParameters(f13)
	}
	if r.ko.Spec.RequestTemplates != nil {
		f14 := map[string]*string{}
		for f14key, f14valiter := range r.ko.Spec.RequestTemplates {
			var f14val string
			f14val = *f14valiter
			f14[f14key] = &f14val
		}
		res.SetRequestTemplates(f14)
	}
	if r.ko.Spec.ResponseParameters != nil {
		f15 := map[string]map[string]*string{}
		for f15key, f15valiter := range r.ko.Spec.ResponseParameters {
			f15val := map[string]*string{}
			for f15valkey, f15valvaliter := range f15valiter {
				var f15valval string
				f15valval = *f15valvaliter
				f15val[f15valkey] = &f15valval
			}
			f15[f15key] = f15val
		}
		res.SetResponseParameters(f15)
	}
	if r.ko.Spec.TemplateSelectionExpression != nil {
		res.SetTemplateSelectionExpression(*r.ko.Spec.TemplateSelectionExpression)
	}
	if r.ko.Spec.TimeoutInMillis != nil {
		res.SetTimeoutInMillis(*r.ko.Spec.TimeoutInMillis)
	}
	if r.ko.Spec.TLSConfig != nil {
		f18 := &svcsdk.TlsConfigInput{}
		if r.ko.Spec.TLSConfig.ServerNameToVerify != nil {
			f18.SetServerNameToVerify(*r.ko.Spec.TLSConfig.ServerNameToVerify)
		}
		res.SetTlsConfig(f18)
	}

	return res, nil
}

// sdkDelete deletes the supplied resource in the backend AWS service API
func (rm *resourceManager) sdkDelete(
	ctx context.Context,
	r *resource,
) (latest *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkDelete")
	defer exit(err)
	input, err := rm.newDeleteRequestPayload(r)
	if err != nil {
		return nil, err
	}
	var resp *svcsdk.DeleteIntegrationOutput
	_ = resp
	resp, err = rm.sdkapi.DeleteIntegrationWithContext(ctx, input)
	rm.metrics.RecordAPICall("DELETE", "DeleteIntegration", err)
	return nil, err
}

// newDeleteRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Delete API call for the resource
func (rm *resourceManager) newDeleteRequestPayload(
	r *resource,
) (*svcsdk.DeleteIntegrationInput, error) {
	res := &svcsdk.DeleteIntegrationInput{}

	if r.ko.Spec.APIID != nil {
		res.SetApiId(*r.ko.Spec.APIID)
	}
	if r.ko.Status.IntegrationID != nil {
		res.SetIntegrationId(*r.ko.Status.IntegrationID)
	}

	return res, nil
}

// setStatusDefaults sets default properties into supplied custom resource
func (rm *resourceManager) setStatusDefaults(
	ko *svcapitypes.Integration,
) {
	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if ko.Status.ACKResourceMetadata.Region == nil {
		ko.Status.ACKResourceMetadata.Region = &rm.awsRegion
	}
	if ko.Status.ACKResourceMetadata.OwnerAccountID == nil {
		ko.Status.ACKResourceMetadata.OwnerAccountID = &rm.awsAccountID
	}
	if ko.Status.Conditions == nil {
		ko.Status.Conditions = []*ackv1alpha1.Condition{}
	}
}

// updateConditions returns updated resource, true; if conditions were updated
// else it returns nil, false
func (rm *resourceManager) updateConditions(
	r *resource,
	onSuccess bool,
	err error,
) (*resource, bool) {
	ko := r.ko.DeepCopy()
	rm.setStatusDefaults(ko)

	// Terminal condition
	var terminalCondition *ackv1alpha1.Condition = nil
	var recoverableCondition *ackv1alpha1.Condition = nil
	var syncCondition *ackv1alpha1.Condition = nil
	for _, condition := range ko.Status.Conditions {
		if condition.Type == ackv1alpha1.ConditionTypeTerminal {
			terminalCondition = condition
		}
		if condition.Type == ackv1alpha1.ConditionTypeRecoverable {
			recoverableCondition = condition
		}
		if condition.Type == ackv1alpha1.ConditionTypeResourceSynced {
			syncCondition = condition
		}
	}

	if rm.terminalAWSError(err) || err == ackerr.SecretTypeNotSupported || err == ackerr.SecretNotFound {
		if terminalCondition == nil {
			terminalCondition = &ackv1alpha1.Condition{
				Type: ackv1alpha1.ConditionTypeTerminal,
			}
			ko.Status.Conditions = append(ko.Status.Conditions, terminalCondition)
		}
		var errorMessage = ""
		if err == ackerr.SecretTypeNotSupported || err == ackerr.SecretNotFound {
			errorMessage = err.Error()
		} else {
			awsErr, _ := ackerr.AWSError(err)
			errorMessage = awsErr.Error()
		}
		terminalCondition.Status = corev1.ConditionTrue
		terminalCondition.Message = &errorMessage
	} else {
		// Clear the terminal condition if no longer present
		if terminalCondition != nil {
			terminalCondition.Status = corev1.ConditionFalse
			terminalCondition.Message = nil
		}
		// Handling Recoverable Conditions
		if err != nil {
			if recoverableCondition == nil {
				// Add a new Condition containing a non-terminal error
				recoverableCondition = &ackv1alpha1.Condition{
					Type: ackv1alpha1.ConditionTypeRecoverable,
				}
				ko.Status.Conditions = append(ko.Status.Conditions, recoverableCondition)
			}
			recoverableCondition.Status = corev1.ConditionTrue
			awsErr, _ := ackerr.AWSError(err)
			errorMessage := err.Error()
			if awsErr != nil {
				errorMessage = awsErr.Error()
			}
			recoverableCondition.Message = &errorMessage
		} else if recoverableCondition != nil {
			recoverableCondition.Status = corev1.ConditionFalse
			recoverableCondition.Message = nil
		}
	}
	// Required to avoid the "declared but not used" error in the default case
	_ = syncCondition
	if terminalCondition != nil || recoverableCondition != nil || syncCondition != nil {
		return &resource{ko}, true // updated
	}
	return nil, false // not updated
}

// terminalAWSError returns awserr, true; if the supplied error is an aws Error type
// and if the exception indicates that it is a Terminal exception
// 'Terminal' exception are specified in generator configuration
func (rm *resourceManager) terminalAWSError(err error) bool {
	// No terminal_errors specified for this resource in generator config
	return false
}
