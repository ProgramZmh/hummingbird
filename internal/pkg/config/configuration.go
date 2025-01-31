/*******************************************************************************
 * Copyright 2019 Dell Inc.
 * Copyright 2020 Intel Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 *******************************************************************************/

package config

// Configuration interface provides an abstraction around a configuration struct.
type Configuration interface {
	// UpdateFromRaw converts configuration received from the registry to a service-specific configuration struct which is
	// then used to overwrite the service's existing configuration struct.
	UpdateFromRaw(rawConfig interface{}) bool

	// EmptyWritablePtr returns a pointer to a service-specific empty WritableInfo struct.  It is used by the bootstrap to
	// provide the appropriate structure to registry.Client's WatchForChanges().
	EmptyWritablePtr() interface{}

	// UpdateWritableFromRaw converts configuration received from the registry to a service-specific WritableInfo struct
	// which is then used to overwrite the service's existing configuration's WritableInfo struct.
	UpdateWritableFromRaw(rawWritable interface{}) bool

	// GetBootstrap returns the configuration elements required by the bootstrap.
	GetBootstrap() BootstrapConfiguration

	// GetLogLevel returns the current ConfigurationStruct's log level.
	GetLogLevel() string

	// GetLogPath returns the current ConfigurationStruct's log path.
	GetLogPath() string

	// GetRegistryInfo gets the config.RegistryInfo field from the ConfigurationStruct.
	GetRegistryInfo() RegistryInfo

	// GetInsecureSecrets gets the config.InsecureSecrets field from the ConfigurationStruct.
	GetInsecureSecrets() InsecureSecrets
}
