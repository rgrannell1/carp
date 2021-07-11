package main

func TestDependency(carpfile map[string]Group, tgt Dependency) (bool, []string) {
	switch id := tgt["id"]; {
	case id == "core/file":
		return TestFileDependency(tgt)
	case id == "core/apt":
		return TestAptDependency(tgt)
	case id == "core/folder":
		return TestFolderDependency(tgt)
	case id == "core/envvar":
		return TestEnvVarDependency(tgt)
	case id == "core/carpgroup":
		return TestCarpGroupDependency(carpfile, tgt)
	case id == "core/snap":
		return TestSnapDependency(tgt)
	case id == "core/command":
		return TestCommand(tgt)
	default:
		return false, []string{"invalid dependency."}
	}
}
