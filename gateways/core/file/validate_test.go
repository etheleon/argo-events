/*
Copyright 2018 BlackRock, Inc.

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

package file

import (
	"context"
	"testing"

	"github.com/argoproj/argo-events/gateways"
	"github.com/smartystreets/goconvey/convey"
)

var (
	configKey   = "testConfig"
	configId    = "1234"
	configValue = `
directory: "/bin/"
type: CREATE
path: x.txt
`
)

func TestValidateFileEventSource(t *testing.T) {
	convey.Convey("Given a valid file event source spec, parse it and make sure no error occurs", t, func() {
		ese := &FileEventSourceExecutor{}
		valid, err := ese.ValidateEventSource(context.Background(), &gateways.EventSource{
			Name: configKey,
			Id:   configId,
			Data: configValue,
		})
		convey.So(err, convey.ShouldBeNil)
		convey.So(valid, convey.ShouldNotBeNil)
		convey.So(valid.IsValid, convey.ShouldBeTrue)
	})

	convey.Convey("Given an invalid file event source spec, parse it and make sure error occurs", t, func() {
		ese := &FileEventSourceExecutor{}
		invalidConfig := `
type: CREATE
path: x.txt`
		valid, err := ese.ValidateEventSource(context.Background(), &gateways.EventSource{
			Data: invalidConfig,
			Id:   configId,
			Name: configKey,
		})
		convey.So(err, convey.ShouldNotBeNil)
		convey.So(valid, convey.ShouldNotBeNil)
		convey.So(valid.IsValid, convey.ShouldBeFalse)
		convey.So(valid.Reason, convey.ShouldNotBeEmpty)
	})
}
