using System;
using System.Collections.Generic;
using Zenject;
using ModestTree;
using ModestTree.Util;

namespace Zenject
{
    public class ExecutionOrderInstaller : Installer
    {
        List<Type> _typeOrder;

        public ExecutionOrderInstaller(List<Type> typeOrder)
        {
            _typeOrder = typeOrder;
        }

        public override void InstallBindings()
        {
            // All tickables without explicit priorities assigned are given priority of zero,
            // so put all of these before that (ie. negative)
            int priorityCount = -1 * _typeOrder.Count;

            foreach (var type in _typeOrder)
            {
                BindPriority(Container, type, priorityCount);
                priorityCount++;
            }
        }

        public static void BindPriority<T>(
            DiContainer container, int priorityCount)
        {
            BindPriority(container, typeof(T), priorityCount);
        }

        public static void BindPriority(
            DiContainer container, Type type, int priority)
        {
            Assert.That(type.DerivesFrom<ITickable>() || type.DerivesFrom<IInitializable>() || type.DerivesFrom<IDisposable>(),
                "Expected type '{0}' to derive from ITickable, IInitializable, or IDisposable", type.Name());

            if (type.DerivesFrom<ITickable>())
            {
                container.Bind<ModestTree.Util.Tuple<Type, int>>().ToInstance(
                    ModestTree.Util.Tuple.New(type, priority)).WhenInjectedInto<TickableManager>();
            }

            if (type.DerivesFrom<IInitializable>())
            {
                container.Bind<ModestTree.Util.Tuple<Type, int>>().ToInstance(
                    ModestTree.Util.Tuple.New(type, priority)).WhenInjectedInto<InitializableManager>();
            }

            if (type.DerivesFrom<IDisposable>())
            {
                container.Bind<ModestTree.Util.Tuple<Type, int>>().ToInstance(
                    ModestTree.Util.Tuple.New(type, priority)).WhenInjectedInto<DisposableManager>();
            }
        }
    }
}

