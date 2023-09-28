<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# mageutil

```go
import "github.com/elisasre/mageutil"
```

package mageutil provides util functions for [Magefile](<https://magefile.org/>). For usage please refer to [documentation](<https://magefile.org/importing/>) provided by Magefile. For autocompletions see [completions](<https://github.com/elisasre/mageutil/tree/main/completions>). Since this package is private it is recommended to set GOPRIVATE env variable:

go env \-w GOPRIVATE=github.com/elisasre

With GOPRIVATE beeing set you can update mageutils by running:

go get github.com/elisasre/mageutil@main

## Index

- [Constants](<#constants>)
- [func BinDir\(\) \(string, error\)](<#BinDir>)
- [func Build\(ctx context.Context, name string\) error](<#Build>)
- [func BuildAll\(ctx context.Context\) error](<#BuildAll>)
- [func BuildFor\(ctx context.Context, goos, goarch, name string\) error](<#BuildFor>)
- [func BuildForArmMac\(ctx context.Context, name string\)](<#BuildForArmMac>)
- [func BuildForLinux\(ctx context.Context, name string\)](<#BuildForLinux>)
- [func BuildForMac\(ctx context.Context, name string\)](<#BuildForMac>)
- [func BuildWithSHA\(ctx context.Context, goos, goarch, name string\)](<#BuildWithSHA>)
- [func CGO\(enabled bool\)](<#CGO>)
- [func Clean\(ctx context.Context\) error](<#Clean>)
- [func CoverInfo\(ctx context.Context, profile string\) error](<#CoverInfo>)
- [func DefaultLabels\(imageName, url, desc string\) map\[string\]string](<#DefaultLabels>)
- [func Docker\(ctx context.Context, args ...string\) error](<#Docker>)
- [func DockerBuild\(ctx context.Context, platform, dockerfile, buildCtx string, tags \[\]string, extraCtx, labels map\[string\]string\) error](<#DockerBuild>)
- [func DockerBuildDefault\(ctx context.Context, imageName, url string\) error](<#DockerBuildDefault>)
- [func DockerPushAllTags\(ctx context.Context, imageName string\) error](<#DockerPushAllTags>)
- [func DockerTags\(imageName string, tags ...string\) \[\]string](<#DockerTags>)
- [func Git\(ctx context.Context, args ...string\) error](<#Git>)
- [func Go\(ctx context.Context, args ...string\) error](<#Go>)
- [func GoList\(ctx context.Context, target string\) \(\[\]string, error\)](<#GoList>)
- [func GoWith\(ctx context.Context, env map\[string\]string, args ...string\) error](<#GoWith>)
- [func GolangCILint\(ctx context.Context, args ...string\) error](<#GolangCILint>)
- [func IntegrationTest\(ctx context.Context, pkg string\) error](<#IntegrationTest>)
- [func LicenseCheck\(ctx context.Context, w io.Writer, targets ...string\) error](<#LicenseCheck>)
- [func LintAll\(ctx context.Context\) error](<#LintAll>)
- [func MergeCover\(ctx context.Context, coverFiles \[\]string, w io.Writer\) error](<#MergeCover>)
- [func MergeCoverProfiles\(ctx context.Context\) error](<#MergeCoverProfiles>)
- [func MustSetEnv\(k, v string\)](<#MustSetEnv>)
- [func Run\(ctx context.Context, name string, args ...string\) error](<#Run>)
- [func SHA256Sum\(ctx context.Context, name string\) error](<#SHA256Sum>)
- [func Targets\(ctx context.Context\) \(\[\]string, error\)](<#Targets>)
- [func UnitTest\(ctx context.Context\) error](<#UnitTest>)
- [func Verbose\(enabled bool\)](<#Verbose>)
- [func VulnCheck\(ctx context.Context, args ...string\) error](<#VulnCheck>)
- [func VulnCheckAll\(ctx context.Context\) error](<#VulnCheckAll>)


## Constants

<a name="OCILabelTitle"></a>

```go
const (
    OCILabelTitle       = "org.opencontainers.image.title"
    OCILabelURL         = "org.opencontainers.image.url"
    OCILabelVersion     = "org.opencontainers.image.version"
    OCILabelDescription = "org.opencontainers.image.description"
    OCILabelCreated     = "org.opencontainers.image.created"
    OCILabelSource      = "org.opencontainers.image.source"
    OCILabelLicenses    = "org.opencontainers.image.licenses"
    OCILabelAuthors     = "org.opencontainers.image.authors"
    OCILabelVendor      = "org.opencontainers.image.vendor"
    OCILabelRevision    = "org.opencontainers.image.revision"
)
```

<a name="DefaultPlatform"></a>

```go
const (
    DefaultPlatform   = "linux/amd64"
    DefaultDockerfile = "Dockerfile"
    DefaultBuildCtx   = "."
    DefaultExtraCtx   = TargetDir + "bin/" + "linux/amd64/"
)
```

<a name="CmdDir"></a>

```go
const (
    CmdDir     = "./cmd/"
    TargetDir  = "./target/"
    ReportsDir = TargetDir + "reports/"
)
```

<a name="UnitCoverProfile"></a>

```go
const (
    UnitCoverProfile        = ReportsDir + "unit-test-coverage.out"
    IntegrationCoverProfile = ReportsDir + "integration-test-coverage.out"
    MergedCoverProfile      = ReportsDir + "merged-test-coverage.out"
)
```

<a name="UNKNOWN"></a>

```go
const (
    UNKNOWN = "Unknown"
)
```

<a name="BinDir"></a>
## func [BinDir](<https://github.com/elisasre/mageutil/blob/main/go.go#L144>)

```go
func BinDir() (string, error)
```

BinDir returns path in format of target/bin/\{GOOS\}/\{GOARCH\}

<a name="Build"></a>
## func [Build](<https://github.com/elisasre/mageutil/blob/main/go.go#L76>)

```go
func Build(ctx context.Context, name string) error
```

Build binary using settings from system env.

<a name="BuildAll"></a>
## func [BuildAll](<https://github.com/elisasre/mageutil/blob/main/go.go#L62>)

```go
func BuildAll(ctx context.Context) error
```

BuildAll binaries for targets returned by utils.Targets using utils.Build.

<a name="BuildFor"></a>
## func [BuildFor](<https://github.com/elisasre/mageutil/blob/main/go.go#L97>)

```go
func BuildFor(ctx context.Context, goos, goarch, name string) error
```

BuildDefault binary using settings from system env.

<a name="BuildForArmMac"></a>
## func [BuildForArmMac](<https://github.com/elisasre/mageutil/blob/main/go.go#L118>)

```go
func BuildForArmMac(ctx context.Context, name string)
```

BuildForArmMac builds binary for arm64 based mac systems.

<a name="BuildForLinux"></a>
## func [BuildForLinux](<https://github.com/elisasre/mageutil/blob/main/go.go#L108>)

```go
func BuildForLinux(ctx context.Context, name string)
```

BuildForLinux builds binary for amd64 based linux systems.

<a name="BuildForMac"></a>
## func [BuildForMac](<https://github.com/elisasre/mageutil/blob/main/go.go#L113>)

```go
func BuildForMac(ctx context.Context, name string)
```

BuildForMac builds binary for amd64 based mac systems.

<a name="BuildWithSHA"></a>
## func [BuildWithSHA](<https://github.com/elisasre/mageutil/blob/main/go.go#L91>)

```go
func BuildWithSHA(ctx context.Context, goos, goarch, name string)
```

BuildDefault binary and SHA256 sum using settings from system env

<a name="CGO"></a>
## func [CGO](<https://github.com/elisasre/mageutil/blob/main/env.go#L23>)

```go
func CGO(enabled bool)
```

CGO can be used to enable of disable CGO. By default this package will disable CGO.

<a name="Clean"></a>
## func [Clean](<https://github.com/elisasre/mageutil/blob/main/git.go#L15>)

```go
func Clean(ctx context.Context) error
```

Clean removes all files ignored by git.

<a name="CoverInfo"></a>
## func [CoverInfo](<https://github.com/elisasre/mageutil/blob/main/testing.go#L88>)

```go
func CoverInfo(ctx context.Context, profile string) error
```

CoverInfo prints function level cover stats from given profile.

<a name="DefaultLabels"></a>
## func [DefaultLabels](<https://github.com/elisasre/mageutil/blob/main/docker.go#L89>)

```go
func DefaultLabels(imageName, url, desc string) map[string]string
```

DefaultLabels provides labels for Elisa SoSe/SRE organization.

<a name="Docker"></a>
## func [Docker](<https://github.com/elisasre/mageutil/blob/main/docker.go#L35>)

```go
func Docker(ctx context.Context, args ...string) error
```

Docker runs systems docker cmd with given args.

<a name="DockerBuild"></a>
## func [DockerBuild](<https://github.com/elisasre/mageutil/blob/main/docker.go#L53>)

```go
func DockerBuild(ctx context.Context, platform, dockerfile, buildCtx string, tags []string, extraCtx, labels map[string]string) error
```

DockerBuild is a short hand for docker buildx build with saine default flags

<a name="DockerBuildDefault"></a>
## func [DockerBuildDefault](<https://github.com/elisasre/mageutil/blob/main/docker.go#L45>)

```go
func DockerBuildDefault(ctx context.Context, imageName, url string) error
```

DockerBuildDefault build image with sane defaults.

<a name="DockerPushAllTags"></a>
## func [DockerPushAllTags](<https://github.com/elisasre/mageutil/blob/main/docker.go#L40>)

```go
func DockerPushAllTags(ctx context.Context, imageName string) error
```

DockerPushAllTags push all tags for given image.

<a name="DockerTags"></a>
## func [DockerTags](<https://github.com/elisasre/mageutil/blob/main/docker.go#L70>)

```go
func DockerTags(imageName string, tags ...string) []string
```

DockerTags creates slice of tags usign \`tags\` variable and DOCKER\_IMAGE\_TAGS env var.

<a name="Git"></a>
## func [Git](<https://github.com/elisasre/mageutil/blob/main/git.go#L10>)

```go
func Git(ctx context.Context, args ...string) error
```

Git is shorthand for git executable provided by system.

<a name="Go"></a>
## func [Go](<https://github.com/elisasre/mageutil/blob/main/go.go#L35>)

```go
func Go(ctx context.Context, args ...string) error
```

Go is shorthand for go executable provided by system.

<a name="GoList"></a>
## func [GoList](<https://github.com/elisasre/mageutil/blob/main/go.go#L134>)

```go
func GoList(ctx context.Context, target string) ([]string, error)
```

GoList lists all packages in given target.

<a name="GoWith"></a>
## func [GoWith](<https://github.com/elisasre/mageutil/blob/main/go.go#L40>)

```go
func GoWith(ctx context.Context, env map[string]string, args ...string) error
```

GoWith is shorthand for go executable provided by system.

<a name="GolangCILint"></a>
## func [GolangCILint](<https://github.com/elisasre/mageutil/blob/main/lint.go#L18>)

```go
func GolangCILint(ctx context.Context, args ...string) error
```

LintNative imports golangci\-lint and runs it as a helper library.

<a name="IntegrationTest"></a>
## func [IntegrationTest](<https://github.com/elisasre/mageutil/blob/main/testing.go#L34>)

```go
func IntegrationTest(ctx context.Context, pkg string) error
```

IntegrationTest executes all tests in given pkg with default flags.

<a name="LicenseCheck"></a>
## func [LicenseCheck](<https://github.com/elisasre/mageutil/blob/main/licenses.go#L20>)

```go
func LicenseCheck(ctx context.Context, w io.Writer, targets ...string) error
```

LicenseCheck runs github.com/google/go\-licenses/licenses for given targets and writes toe output into w.

<a name="LintAll"></a>
## func [LintAll](<https://github.com/elisasre/mageutil/blob/main/lint.go#L13>)

```go
func LintAll(ctx context.Context) error
```

LintAll uses golangci\-lint library to lint all go files.

<a name="MergeCover"></a>
## func [MergeCover](<https://github.com/elisasre/mageutil/blob/main/testing.go#L52>)

```go
func MergeCover(ctx context.Context, coverFiles []string, w io.Writer) error
```

MergeCover merges multiple go test \-cover profiles and writes the combined coverage into out.

coverage merging is adapted from https://github.com/wadey/gocovmerge

<a name="MergeCoverProfiles"></a>
## func [MergeCoverProfiles](<https://github.com/elisasre/mageutil/blob/main/testing.go#L72>)

```go
func MergeCoverProfiles(ctx context.Context) error
```

MergeCoverProfiles merges default unit and integration cover profile.

<a name="MustSetEnv"></a>
## func [MustSetEnv](<https://github.com/elisasre/mageutil/blob/main/env.go#L31>)

```go
func MustSetEnv(k, v string)
```



<a name="Run"></a>
## func [Run](<https://github.com/elisasre/mageutil/blob/main/go.go#L123>)

```go
func Run(ctx context.Context, name string, args ...string) error
```

Run executes app binary from default path.

<a name="SHA256Sum"></a>
## func [SHA256Sum](<https://github.com/elisasre/mageutil/blob/main/sha256.go#L14>)

```go
func SHA256Sum(ctx context.Context, name string) error
```

SHA256Sum calculates sum for single file and stores it in file. Output should be compatible with sha256sum program.

<a name="Targets"></a>
## func [Targets](<https://github.com/elisasre/mageutil/blob/main/go.go#L46>)

```go
func Targets(ctx context.Context) ([]string, error)
```

Targets returns list of main pkgs under utils.CmdDir.

<a name="UnitTest"></a>
## func [UnitTest](<https://github.com/elisasre/mageutil/blob/main/testing.go#L23>)

```go
func UnitTest(ctx context.Context) error
```

UnitTest executes all unit tests with default flags.

<a name="Verbose"></a>
## func [Verbose](<https://github.com/elisasre/mageutil/blob/main/env.go#L18>)

```go
func Verbose(enabled bool)
```

Verbose can be used to control mage's verbose state. By default this package will set mage in verbose state.

<a name="VulnCheck"></a>
## func [VulnCheck](<https://github.com/elisasre/mageutil/blob/main/vulncheck.go#L10>)

```go
func VulnCheck(ctx context.Context, args ...string) error
```

VulnChek runs golang.org/x/vuln/scan with given args.

<a name="VulnCheckAll"></a>
## func [VulnCheckAll](<https://github.com/elisasre/mageutil/blob/main/vulncheck.go#L20>)

```go
func VulnCheckAll(ctx context.Context) error
```

VulnChek runs golang.org/x/vuln/scan for all packages.

Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)
