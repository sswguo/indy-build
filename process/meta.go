package process

const (
	DEFAULT_TARGET      = "pnc-builds"
	DEFAULT_MVN_CENTRAL = "central"
	DEFAULT_NPM_CENTRAL = "npmjs"
	TYPE_MVN            = "maven"
	TYPE_NPM            = "npm"
)

type BuildMetadata struct {
	buildType     string
	centralName   string
	promoteTarget string
}

func decideMeta(buildType string) *BuildMetadata {
	if buildType == TYPE_MVN {
		return &BuildMetadata{
			buildType:     buildType,
			centralName:   DEFAULT_MVN_CENTRAL,
			promoteTarget: DEFAULT_TARGET,
		}
	} else if buildType == TYPE_NPM {
		return &BuildMetadata{
			buildType:     buildType,
			centralName:   DEFAULT_NPM_CENTRAL,
			promoteTarget: DEFAULT_TARGET,
		}
	}
	return nil
}
