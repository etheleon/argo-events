/*
Copyright 2018 KompiTech GmbH

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package github

import (
	"github.com/ghodss/yaml"
	"github.com/google/go-github/github"
	"github.com/rs/zerolog"
	"k8s.io/client-go/kubernetes"
)

// GithubEventSourceExecutor implements ConfigExecutor
type GithubEventSourceExecutor struct {
	Log zerolog.Logger
	// GitlabClient is client for gitlab api
	GithubClient *github.Client
	// Clientset is kubernetes client
	Clientset kubernetes.Interface
	// Namespace where gateway is deployed
	Namespace string
}

// GithubConfig contains information to setup a github project integration
// +k8s:openapi-gen=true
type GithubConfig struct {
	// GitHub owner name i.e. argoproj
	Owner string `json:"owner"`

	// GitHub repo name i.e. argo-events
	Repository string `json:"repository"`

	// Github events to subscribe to which the gateway will subscribe
	Events []string `json:"events"`

	// External URL for hooks
	URL string `json:"url"`

	// K8s secret containing github api token
	APIToken *GithubSecret `json:"apiToken"`

	// K8s secret containing WebHook Secret
	WebHookSecret *GithubSecret `json:"webHookSecret"`

	// Insecure tls verification
	Insecure bool `json:"insecure"`

	// Active
	Active bool `json:"active"`

	// ContentType json or form
	ContentType string `json:"contentType"`
}

// GithubSecret contains information of k8 secret which holds the github api access token key
// +k8s:openapi-gen=true
type GithubSecret struct {
	// Name of k8 secret containing api token or webhook secret
	Name string
	// Key for api token/webhook secret
	Key string
}

// cred stores the api access token or webhook secret
type cred struct {
	secret string
}

// parseEventSource parses a configuration of gateway
func parseEventSource(config string) (*GithubConfig, error) {
	var g *GithubConfig
	err := yaml.Unmarshal([]byte(config), &g)
	if err != nil {
		return nil, err
	}
	return g, err
}
