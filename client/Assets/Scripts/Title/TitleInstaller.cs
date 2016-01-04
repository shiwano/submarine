using UnityEngine;
using System;
using Zenject;
using UnityEngine.UI;

namespace Submarine
{
    public class TitleInstaller : MonoInstaller
    {
        [SerializeField]
        TitleView view;

        public override void InstallBindings()
        {
            Container.Bind<TitleView>().ToSingleInstance(view);

            Container.Bind<IInitializable>().ToSingle<TitleMediator>();
            Container.Bind<TitleMediator>().ToSingle();
        }
    }
}
