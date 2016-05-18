using UnityEngine;
using UnityEditor;
using System.IO;
using System.Text;
using System.Text.RegularExpressions;

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
        var meshData = ConvertMeshToJSON(mesh);
        var outputPath = Regex.Replace(assetPath, "NavMesh.fbx", "NavMesh.json");
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

        for (var i = 0; i < mesh.triangles.Length; i++)
        {
            builder.Append(mesh.triangles[i].ToString());
            if (i < mesh.triangles.Length - 1)
            {
                builder.Append(",");
            }
        }

        builder.Append("]}");
        return builder.ToString();
    }

    void WriteTextToFile(string outputPath, string contents)
    {
        File.WriteAllText(outputPath, contents);
        AssetDatabase.Refresh();
        Debug.Log("Generated " + outputPath);
    }
}