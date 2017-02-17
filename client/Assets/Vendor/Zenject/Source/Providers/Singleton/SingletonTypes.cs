namespace Zenject
{
    public enum SingletonTypes
    {
        FromNew,
        FromMethod,
        FromSubContainerMethod,
        FromSubContainerInstaller,
        FromInstance,
        FromPrefab,
        FromPrefabResource,
        FromFactory,
        FromGameObject,
        FromComponent,
        FromComponentGameObject,
        FromGetter,
        FromResolve,
        FromResource,
        FromScriptableObjectResource,
        FromSubContainerPrefab,
        FromSubContainerPrefabResource,
    }
}
