#if !ZEN_NOT_UNITY3D

using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using ModestTree.Util;
using UnityEditor;
using UnityEngine;
using ModestTree;

#if UNITY_5_3
using UnityEditor.SceneManagement;
#endif

namespace Zenject
{
    public static class ZenEditorUtil
    {
        public static IEnumerable<ZenjectResolveException> ValidateInstallers(SceneCompositionRoot compRoot)
        {
            return ValidateInstallers(compRoot, null);
        }

        public static IEnumerable<ZenjectResolveException> ValidateInstallers(SceneCompositionRoot compRoot, GameObject rootGameObject)
        {
            var globalContainer = GlobalCompositionRoot.CreateContainer(true, null);
            var container = compRoot.CreateContainer(true, globalContainer, new List<IInstaller>());

            foreach (var error in container.ValidateResolve(
                new InjectContext(container, typeof(IFacade), null)))
            {
                yield return error;
            }

            var injectedGameObjects = rootGameObject != null ? rootGameObject.GetComponentsInChildren<Transform>() : GameObject.FindObjectsOfType<Transform>();

            // Also make sure we can fill in all the dependencies in the built-in scene
            foreach (var curTransform in injectedGameObjects)
            {
                foreach (var monoBehaviour in curTransform.GetComponents<MonoBehaviour>())
                {
                    if (monoBehaviour == null)
                    {
                        // fiBackupSceneStorage shows up sometimes for reasons I don't understand
                        // but it's normal so ignore
                        if (curTransform.name != "fiBackupSceneStorage")
                        {
                            Log.Warn("Found null MonoBehaviour on " + curTransform.name);
                        }

                        continue;
                    }

                    foreach (var error in container.ValidateObjectGraph(monoBehaviour.GetType()))
                    {
                        yield return error;
                    }
                }
            }

            foreach (var installer in globalContainer.InstalledInstallers.Concat(container.InstalledInstallers).OfType<IValidatable>())
            {
                foreach (var error in installer.Validate())
                {
                    yield return error;
                }
            }

            foreach (var error in container.ValidateValidatables())
            {
                yield return error;
            }
        }

        public static void ValidateCurrentSceneThenPlay()
        {
            if (ValidateCurrentScene())
            {
                EditorApplication.isPlaying = true;
            }
        }

        public static SceneDecoratorCompositionRoot TryGetSceneDecoratorCompositionRoot()
        {
            return GameObject.FindObjectsOfType<SceneDecoratorCompositionRoot>().OnlyOrDefault();
        }

        public static SceneCompositionRoot TryGetSceneCompositionRoot()
        {
            return GameObject.FindObjectsOfType<SceneCompositionRoot>().OnlyOrDefault();
        }

        // Returns true if we should continue
        static bool CheckForExistingCompositionRoot()
        {
            if (TryGetSceneCompositionRoot() != null)
            {
                var shouldContinue = EditorUtility.DisplayDialog("Error", "There already exists a SceneCompositionRoot in the scene.  Are you sure you want to add another?", "Yes", "Cancel");
                return shouldContinue;
            }

            return true;
        }

        [MenuItem("GameObject/Zenject/Scene Composition Root", false, 9)]
        public static void CreateSceneCompositionRoot(MenuCommand menuCommand)
        {
            if (CheckForExistingCompositionRoot())
            {
                var root = new GameObject("CompositionRoot").AddComponent<SceneCompositionRoot>();
                Selection.activeGameObject = root.gameObject;
            }
        }

        [MenuItem("GameObject/Zenject/Decorator Composition Root", false, 9)]
        public static void CreateDecoratorCompositionRoot(MenuCommand menuCommand)
        {
            if (CheckForExistingCompositionRoot())
            {
                var root = new GameObject("DecoratorCompositionRoot").AddComponent<SceneDecoratorCompositionRoot>();
                Selection.activeGameObject = root.gameObject;
            }
        }

        [MenuItem("Assets/Create/Zenject/Global Installers Asset")]
        public static void AddGlobalInstallers()
        {
            var dir = UnityEditorUtil.TryGetCurrentDirectoryInProjectsTab();

            var assetName = GlobalCompositionRoot.GlobalInstallersResourceName + ".asset";

            if (dir == null)
            {
                EditorUtility.DisplayDialog("Error",
                    "Could not find directory to place the {0} asset.  Please try again by right clicking in the desired folder within the projects pane.".Fmt(assetName), "Ok");
                return;
            }

            var parentFolderName = Path.GetFileName(dir);

            if (parentFolderName != "Resources")
            {
                EditorUtility.DisplayDialog("Error", "{0} must be placed inside a directory named 'Resources'.  Please try again by right clicking within the Project pane in a valid Resources folder.".Fmt(assetName), "Ok");
                return;
            }

            var asset = ScriptableObject.CreateInstance<GlobalInstallerConfig>();

            string assetPath = AssetDatabase.GenerateUniqueAssetPath(Path.Combine(dir, assetName));
            AssetDatabase.CreateAsset(asset, assetPath);
            AssetDatabase.Refresh();

            Debug.Log("Created new asset at '{0}'".Fmt(assetPath));
        }

        // This can be called by build scripts using batch mode unity for continuous integration testing
        // This will exit with an error code for whether validation passed or not
        public static void ValidateAllScenesFromScript()
        {
            ValidateAllActiveScenes(true);
        }

        [MenuItem("Edit/Zenject/Validate All Active Scenes")]
        public static bool ValidateAllActiveScenes()
        {
            return ValidateAllActiveScenes(false);
        }

        public static bool ValidateAllActiveScenes(bool exitAfter)
        {
            return ValidateScenes(UnityEditorUtil.GetAllActiveScenePaths(), exitAfter);
        }

        static string GetActiveScene()
        {
#if UNITY_5_3
            return EditorSceneManager.GetActiveScene().path;
#else
            return EditorApplication.currentScene;
#endif
        }

        public static void OpenScene(string scenePath)
        {
#if UNITY_5_3
            EditorSceneManager.OpenScene(scenePath, OpenSceneMode.Single);
#else
            EditorApplication.OpenScene(scenePath);
#endif
        }

        public static void OpenSceneAdditive(string scenePath)
        {
#if UNITY_5_3
            EditorSceneManager.OpenScene(scenePath, OpenSceneMode.Additive);
#else
            EditorApplication.OpenSceneAdditive(scenePath);
#endif
        }

        public static bool ValidateScenes(List<string> scenePaths, bool exitAfter)
        {
            var startScene = GetActiveScene();

            var failedScenes = new List<string>();

            foreach (var scenePath in scenePaths)
            {
                var sceneName = Path.GetFileNameWithoutExtension(scenePath);

                Log.Trace("Validating scene '{0}'...", sceneName);

                OpenScene(scenePath);

                if (!ValidateCurrentScene())
                {
                    Log.Error("Failed to validate scene '{0}'", sceneName);
                    failedScenes.Add(sceneName);
                }
            }

            OpenScene(startScene);

            if (failedScenes.IsEmpty())
            {
                Log.Trace("Successfully validated all {0} scenes", scenePaths.Count);

                if (exitAfter)
                {
                    EditorApplication.Exit(0);
                }

                return true;
            }
            else
            {
                Log.Error("Validated {0}/{1} scenes. Failed to validate the following: {2}",
                    scenePaths.Count - failedScenes.Count, scenePaths.Count, failedScenes.Join(", "));

                if (exitAfter)
                {
                    EditorApplication.Exit(1);
                }

                return false;
            }
        }

        [MenuItem("Edit/Zenject/Validate Current Scene #%v")]
        public static bool ValidateCurrentScene()
        {
            var startTime = DateTime.Now;
            // Only show a few to avoid spamming the log too much
            var resolveErrors = GetCurrentSceneValidationErrors(10).ToList();

            foreach (var error in resolveErrors)
            {
                Log.ErrorException(error);
            }

            var secondsElapsed = (DateTime.Now - startTime).Milliseconds / 1000.0f;

            if (resolveErrors.Any())
            {
                Log.Error("Validation Completed With Errors, Took {0:0.00} Seconds.", secondsElapsed);
                return false;
            }

            Log.Info("Validation Completed Successfully, Took {0:0.00} Seconds.", secondsElapsed);
            return true;
        }

        static List<ZenjectResolveException> GetCurrentSceneValidationErrors(int maxErrors)
        {
            var compRoot = GameObject.FindObjectsOfType<SceneCompositionRoot>().OnlyOrDefault();

            if (compRoot != null)
            {
                return ValidateCompRoot(compRoot, maxErrors);
            }

            var decoratorCompRoot = GameObject.FindObjectsOfType<SceneDecoratorCompositionRoot>().OnlyOrDefault();

            if (decoratorCompRoot != null)
            {
                return ValidateDecoratorCompRoot(decoratorCompRoot, maxErrors);
            }

            return new List<ZenjectResolveException>()
            {
                new ZenjectResolveException("Unable to find unique composition root in current scene"),
            };
        }

        static List<ZenjectResolveException> ValidateDecoratorCompRoot(SceneDecoratorCompositionRoot decoratorCompRoot, int maxErrors)
        {
            var sceneName = decoratorCompRoot.SceneName;
            var scenePath = UnityEditorUtil.GetScenePath(sceneName);

            if (scenePath == null)
            {
                return new List<ZenjectResolveException>()
                {
                    new ZenjectResolveException(
                        "Could not find scene path for decorated scene '{0}'".Fmt(sceneName)),
                };
            }

            var rootObjectsBefore = UnityUtil.GetRootGameObjects();

            OpenSceneAdditive(scenePath);

            var newRootObjects = UnityUtil.GetRootGameObjects().Except(rootObjectsBefore);

            // Use finally to ensure we clean up the data added from OpenSceneAdditive
            try
            {
                var previousBeforeInstallHook = SceneCompositionRoot.BeforeInstallHooks;
                SceneCompositionRoot.BeforeInstallHooks = (container) =>
                {
                    if (previousBeforeInstallHook != null)
                    {
                        previousBeforeInstallHook(container);
                    }

                    decoratorCompRoot.AddPreBindings(container);
                };

                var previousAfterInstallHook = SceneCompositionRoot.AfterInstallHooks;
                SceneCompositionRoot.AfterInstallHooks = (container) =>
                {
                    decoratorCompRoot.AddPostBindings(container);

                    if (previousAfterInstallHook != null)
                    {
                        previousAfterInstallHook(container);
                    }
                };

                var compRoot = newRootObjects.SelectMany(x => x.GetComponentsInChildren<SceneCompositionRoot>()).OnlyOrDefault();

                if (compRoot != null)
                {
                    return ValidateCompRoot(compRoot, maxErrors);
                }

                var newDecoratorCompRoot = newRootObjects.SelectMany(x => x.GetComponentsInChildren<SceneDecoratorCompositionRoot>()).OnlyOrDefault();

                if (newDecoratorCompRoot != null)
                {
                    return ValidateDecoratorCompRoot(newDecoratorCompRoot, maxErrors);
                }

                return new List<ZenjectResolveException>()
                {
                    new ZenjectResolveException(
                        "Could not find composition root for decorated scene '{0}'".Fmt(sceneName)),
                };
            }
            finally
            {
#if UNITY_5_3
                EditorSceneManager.CloseScene(EditorSceneManager.GetSceneByPath(scenePath), true);
#else
                foreach (var newObject in newRootObjects)
                {
                    GameObject.DestroyImmediate(newObject);
                }
#endif
            }
        }

        static List<ZenjectResolveException> ValidateCompRoot(SceneCompositionRoot compRoot, int maxErrors)
        {
            if (compRoot.Installers.IsEmpty())
            {
                return new List<ZenjectResolveException>()
                {
                    new ZenjectResolveException("Could not find installers while validating current scene"),
                };
            }

            return ValidateInstallers(compRoot).Take(maxErrors).ToList();
        }
    }
}
#endif
