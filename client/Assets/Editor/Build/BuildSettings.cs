using UnityEngine;
using UnityEditor;
using System;
using System.Linq;

public class BuildSettings
{
    public BuildTarget BuildTarget;
    public string BundleIdentifier;
    public string BundleVersion;
    public string ProductName;

    public string[] SceneNames
    {
        get { return EditorBuildSettings.scenes.Select(x => x.path).ToArray(); }
    }

    public string OutputPath
    {
        get
        {
            return BuildTarget == BuildTarget.iOS ?
                System.IO.Directory.GetCurrentDirectory() + "/iOSXCodeProject" :
                System.IO.Directory.GetCurrentDirectory() + "/build.apk";
        }
    }

    public BuildSettings(CommandLineCustomArgs args)
    {
        BundleIdentifier = args.GetArg("bundleIdentifier");
        BundleVersion = args.GetArg("bundleVersion");
        ProductName = args.GetArg("productName");

        var target = args.GetArg("buildTarget");

        switch (target)
        {
            case "ios":
                BuildTarget = BuildTarget.iOS;
                break;
            case "android":
                BuildTarget = BuildTarget.Android;
                break;
            default:
                throw new InvalidOperationException("Not supported build target: " + target);
        }
    }
}
