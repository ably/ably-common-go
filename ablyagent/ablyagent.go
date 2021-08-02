// Code generated by github.com/ably/ably-common/go/cmd/ablyagent; DO NOT EDIT.

package ablyagent

// Agent is an agent identifier that can appear in the Ably-Agent HTTP header
// and be tracked as a metric.
//
// See protocol/agents.json.
type Agent struct {
	Identifier string
	Type       string
	Versioned  bool
}

// AblyLibMapping is an agent identifier that can appear in the X-Ably-Lib HTTP
// header and be converted to its equivalent Ably-Agent value.
//
// See protocol/agents.json.
type AblyLibMapping struct {
	Lib   string
	Agent string
}

// Agents is the "agents" list from protocol/agents.json.
var Agents = []*Agent{
	{
		Identifier: "ably-js",
		Type:       "sdk",
		Versioned:  true,
	},
	{
		Identifier: "nodejs",
		Type:       "runtime",
		Versioned:  true,
	},
	{
		Identifier: "browser",
		Type:       "runtime",
		Versioned:  false,
	},
	{
		Identifier: "nativescript",
		Type:       "wrapper",
		Versioned:  true,
	},
	{
		Identifier: "react-native",
		Type:       "wrapper",
		Versioned:  false,
	},
	{
		Identifier: "ably-cocoa",
		Type:       "sdk",
		Versioned:  true,
	},
	{
		Identifier: "iOS",
		Type:       "os",
		Versioned:  true,
	},
	{
		Identifier: "tvOS",
		Type:       "os",
		Versioned:  true,
	},
	{
		Identifier: "watchOS",
		Type:       "os",
		Versioned:  true,
	},
	{
		Identifier: "macOS",
		Type:       "os",
		Versioned:  true,
	},
	{
		Identifier: "windows",
		Type:       "os",
		Versioned:  true,
	},
	{
		Identifier: "linux",
		Type:       "os",
		Versioned:  true,
	},
	{
		Identifier: "ably-java",
		Type:       "sdk",
		Versioned:  true,
	},
	{
		Identifier: "jre",
		Type:       "runtime",
		Versioned:  true,
	},
	{
		Identifier: "android",
		Type:       "os",
		Versioned:  true,
	},
	{
		Identifier: "ably-ruby",
		Type:       "sdk",
		Versioned:  true,
	},
	{
		Identifier: "ruby",
		Type:       "runtime",
		Versioned:  true,
	},
	{
		Identifier: "ably-ruby-rest",
		Type:       "wrapper",
		Versioned:  false,
	},
	{
		Identifier: "ably-dotnet",
		Type:       "sdk",
		Versioned:  true,
	},
	{
		Identifier: "dotnet-framework",
		Type:       "runtime",
		Versioned:  true,
	},
	{
		Identifier: "dotnet-standard",
		Type:       "runtime",
		Versioned:  true,
	},
	{
		Identifier: "uwp",
		Type:       "runtime",
		Versioned:  false,
	},
	{
		Identifier: "xamarin",
		Type:       "runtime",
		Versioned:  false,
	},
	{
		Identifier: "unity",
		Type:       "runtime",
		Versioned:  true,
	},
	{
		Identifier: "python",
		Type:       "runtime",
		Versioned:  true,
	},
	{
		Identifier: "ably-python",
		Type:       "sdk",
		Versioned:  true,
	},
	{
		Identifier: "ably-php",
		Type:       "sdk",
		Versioned:  true,
	},
	{
		Identifier: "laravel",
		Type:       "wrapper",
		Versioned:  false,
	},
	{
		Identifier: "ably-go",
		Type:       "sdk",
		Versioned:  true,
	},
	{
		Identifier: "go",
		Type:       "runtime",
		Versioned:  true,
	},
}

// AblyLibMappings is the "ablyLibMappings" list from protocol/agents.json.
var AblyLibMappings = []*AblyLibMapping{
	{
		Lib:   "android",
		Agent: "ably-java/${version} android/0.0.0",
	},
	{
		Lib:   "cocoa.ios",
		Agent: "ably-cocoa/${version} iOS/0.0.0",
	},
	{
		Lib:   "cocoa.macos",
		Agent: "ably-cocoa/${version} macOS/0.0.0",
	},
	{
		Lib:   "cocoa.tvos",
		Agent: "ably-cocoa/${version} tvOS/0.0.0",
	},
	{
		Lib:   "dotnet",
		Agent: "ably-dotnet/${version}",
	},
	{
		Lib:   "dotnet.framework",
		Agent: "ably-dotnet/${version} dotnet-framework/0.0.0",
	},
	{
		Lib:   "dotnet.netstandard20",
		Agent: "ably-dotnet/${version} dotnet-standard/0.0.0",
	},
	{
		Lib:   "dotnet.uwp",
		Agent: "ably-dotnet/${version} uwp",
	},
	{
		Lib:   "dotnet.xamarin-android",
		Agent: "ably-dotnet/${version} xamarin android/0.0.0",
	},
	{
		Lib:   "dotnet.xamarin-ios",
		Agent: "ably-dotnet/${version} xamarin iOS/0.0.0",
	},
	{
		Lib:   "go",
		Agent: "ably-go/${version}",
	},
	{
		Lib:   "java",
		Agent: "ably-java/${version}",
	},
	{
		Lib:   "js-node",
		Agent: "ably-js/${version} nodejs/0.0.0",
	},
	{
		Lib:   "js-ns",
		Agent: "ably-js/${version} nativescript",
	},
	{
		Lib:   "js-rn",
		Agent: "ably-js/${version} react-native",
	},
	{
		Lib:   "js-web",
		Agent: "ably-js/${version} browser",
	},
	{
		Lib:   "php",
		Agent: "ably-php/${version}",
	},
	{
		Lib:   "php-laravel",
		Agent: "ably-php/${version} laravel",
	},
	{
		Lib:   "python",
		Agent: "ably-python/${version}",
	},
	{
		Lib:   "ruby",
		Agent: "ably-ruby/${version}",
	},
	{
		Lib:   "ruby-rest",
		Agent: "ably-ruby/${version} ably-ruby-rest",
	},
}
