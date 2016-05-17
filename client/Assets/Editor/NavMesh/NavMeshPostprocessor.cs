using UnityEngine;
using UnityEditor;
using System.Linq;
using System.IO;
using System.Text.RegularExpressions;
using Type = TyphenApi.Type.Submarine;

public class NavMeshPostprocessor : AssetPostprocessor
{
    const string NavMeshFbxPathPattern = @"^Assets/Art/Maps/\d+/NavMesh.fbx";

    bool IsTargetAsset
    {
        get { return Regex.Match(assetPath, NavMeshFbxPathPattern).Success; }
    }

    public void OnPostprocessModel(GameObject go)
    {
        if (!IsTargetAsset) return;

        var mesh = go.GetComponent<MeshFilter>().sharedMesh;
        var meshData = ExtractMeshData(mesh);
        var outputPath = Regex.Replace(assetPath, "NavMesh.fbx", "NavMesh.json");
        WriteTextToFile(outputPath, meshData.ToJSON());
    }

    Type.Battle.MeshData ExtractMeshData(Mesh mesh)
    {
        return new Type.Battle.MeshData()
        {
            Vertices = mesh.vertices.Select(v => new Type.Battle.Point() { X = v.x, Y = v.z }).ToList(),
            Triangles = mesh.triangles.Select(i => (long)i).ToList(),
        };
    }

    void WriteTextToFile(string outputPath, string contents)
    {
        File.WriteAllText(outputPath, contents);
        AssetDatabase.Refresh();
        Debug.Log("Generated " + outputPath);
    }

    string CombinePaths(string path1, string path2)
    {
        return Path.Combine(path1, path2).Replace("\\", "/");
    }
}