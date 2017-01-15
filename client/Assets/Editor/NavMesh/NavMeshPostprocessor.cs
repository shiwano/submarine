using UnityEngine;
using UnityEditor;
using System.Linq;
using System.IO;
using System.Text;
using System.Text.RegularExpressions;

public class NavMeshPostprocessor : AssetPostprocessor
{
    const string NavMeshFbxPathPattern = @"^Assets/Art/Stages/(\d+)/NavMesh.fbx";

    bool IsTargetAsset
    {
        get { return Regex.Match(assetPath, NavMeshFbxPathPattern).Success; }
    }

    public void OnPostprocessModel(GameObject go)
    {
        var match = Regex.Match(assetPath, NavMeshFbxPathPattern);
        if (!match.Success) return;

        var mesh = go.GetComponent<MeshFilter>().sharedMesh;
        var meshData = ConvertMeshToJSON(mesh);
        var stageIdAsString = match.Groups[1];
        var outputPath = Regex.Replace(assetPath.Replace('\\', '/'), "NavMesh.fbx", "../../../../../server/battle/assets/stages/" + stageIdAsString + "/mesh.json");
        WriteTextToFile(outputPath, meshData);
    }

    string ConvertMeshToJSON(Mesh mesh)
    {
        var builder = new StringBuilder("{\"vertices\":[");

        for (var i = 0; i < mesh.vertices.Length; i++)
        {
            var vertex = mesh.vertices[i];
            builder.Append(string.Format("{{\"x\":{0},\"y\":{1}}}", vertex.x, vertex.z));
            if (i < mesh.vertices.Length - 1)
            {
                builder.Append(",");
            }
        }

        builder.Append("],\"triangles\":[");

        var triangles = mesh.triangles
            .Select((v, i) => new { Index = i, Value = v })
            .GroupBy(p => p.Index / 3)
            .Select(g => g.Select(p => p.Value).ToArray())
            .ToArray();
        for (var i = 0; i < triangles.Length; i++)
        {
            var triangle = triangles[i];
            builder.Append(string.Format("[{0},{1},{2}]", triangle[0], triangle[1], triangle[2]));
            if (i < triangles.Length - 1)
            {
                builder.Append(",");
            }
        }

        builder.Append("]}");
        return builder.ToString();
    }

    void WriteTextToFile(string outputPath, string contents)
    {
        var absoluteOutputPath = Path.GetFullPath(outputPath);
        Directory.CreateDirectory(Path.GetDirectoryName(absoluteOutputPath));
        File.WriteAllText(absoluteOutputPath, contents);
        Debug.Log("Generated " + absoluteOutputPath);
    }
}