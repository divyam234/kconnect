/*
Copyright 2020 The kconnect Authors.

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

package aws

// ResolveConfiguration will resolve the values for the AWS specific config items that have no value.
// It will query AWS and interactively ask the user for selections.
// func (p *ServiceProvider) ResolveConfiguration(cfg config.ConfigurationSet) error {
// 	p.ensureLogger()
// 	p.logger.Debug("resolving AWS identity configuration items")

// 	// NOTE: resolution is only needed for required fields
// 	if err := p.resolveIdpProvider("idp-provider", cfg); err != nil {
// 		return fmt.Errorf("resolving idp-provider: %w", err)
// 	}
// 	if err := p.resolveIdpEndpoint("idp-endpoint", cfg); err != nil {
// 		return fmt.Errorf("resolving idp-endpoint: %w", err)
// 	}

// 	if err := kaws.ResolvePartition(cfg); err != nil {
// 		return fmt.Errorf("resolving partition: %w", err)
// 	}
// 	if err := kaws.ResolveRegion(cfg); err != nil {
// 		return fmt.Errorf("resolving region: %w", err)
// 	}
// 	if err := resolve.Password(cfg); err != nil {
// 		return fmt.Errorf("resolving password: %w", err)
// 	}

// 	if err := resolve.Username(cfg); err != nil {
// 		return fmt.Errorf("resolving username: %w", err)
// 	}

// 	return nil
// }

// func (p *ServiceProvider) resolveIdpEndpoint(name string, cfg config.ConfigurationSet) error {
// 	if cfg.ExistsWithValue(name) {
// 		return nil
// 	}

// 	prompt := "Endpoint for the IdP"
// 	if cfg.ValueIsList(name) {
// 		listName := cfg.ValueString(name)
// 		return resolve.ChooseFromList(cfg, name, prompt, true, listName)
// 	} else {
// 		return resolve.Input(cfg, name, prompt, true)
// 	}
// }
