using System;
using System.Collections.Generic;

public class CommandLineCustomArgs
{
    readonly Dictionary<string, string> customArgs = new Dictionary<string, string>();

    public string[] CommandLineArgs
    {
        get { return Environment.GetCommandLineArgs(); }
    }

    public CommandLineCustomArgs()
    {
        foreach (var arg in CommandLineArgs)
        {
            var keyAndValue = arg.Split('=');

            if (keyAndValue.Length == 2)
            {
                customArgs.Add(keyAndValue[0], keyAndValue[1]);
            }
        }
    }

    public string GetArg(string argumentName)
    {
        return customArgs[argumentName];
    }

    public string GetArg(string argumentName, string defaultValue)
    {
        string value;
        return customArgs.TryGetValue(argumentName, out value) ? value : defaultValue;
    }
}