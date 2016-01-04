using System;
using System.Collections;
using System.Collections.Generic;
using System.Linq;
using ModestTree;
using ModestTree.Util;

namespace Zenject.Commands
{
    public static class SignalExtensions
    {
        // Zero Parameters
        public static BindingConditionSetter BindSignalTrigger<TTrigger, TSignal>(this IBinder binder, string identifier)
            where TSignal : Signal
            where TTrigger : Signal.TriggerBase
        {
            var container = (DiContainer)binder;
            container.Bind<TSignal>().ToSingle();
            container.Bind<Signal>().ToSingle<TSignal>().WhenInjectedInto<TTrigger>();
            return container.Bind<TTrigger>().ToSingle();
        }

        public static BindingConditionSetter BindSignalTrigger<TTrigger, TSignal>(this IBinder container)
            where TSignal : Signal
            where TTrigger : Signal.TriggerBase
        {
            return BindSignalTrigger<TTrigger, TSignal>((DiContainer)container, null);
        }

        // One Parameter
        public static BindingConditionSetter BindSignalTrigger<TTrigger, TSignal, TParam1>(this IBinder binder, string identifier)
            where TSignal : Signal<TParam1>
            where TTrigger : Signal<TParam1>.TriggerBase
        {
            var container = (DiContainer)binder;
            container.Bind<TSignal>().ToSingle();
            container.Bind<Signal<TParam1>>().ToSingle<TSignal>().WhenInjectedInto<TTrigger>();
            return container.Bind<TTrigger>().ToSingle();
        }

        public static BindingConditionSetter BindSignalTrigger<TTrigger, TSignal, TParam1>(this IBinder container)
            where TSignal : Signal<TParam1>
            where TTrigger : Signal<TParam1>.TriggerBase
        {
            return BindSignalTrigger<TTrigger, TSignal, TParam1>((DiContainer)container, null);
        }

        // Two Parameters
        public static BindingConditionSetter BindSignalTrigger<TTrigger, TSignal, TParam1, TParam2>(this IBinder binder, string identifier)
            where TSignal : Signal<TParam1, TParam2>
            where TTrigger : Signal<TParam1, TParam2>.TriggerBase
        {
            var container = (DiContainer)binder;
            container.Bind<TSignal>().ToSingle();
            container.Bind<Signal<TParam1, TParam2>>().ToSingle<TSignal>().WhenInjectedInto<TTrigger>();
            return container.Bind<TTrigger>().ToSingle();
        }

        public static BindingConditionSetter BindSignalTrigger<TTrigger, TSignal, TParam1, TParam2>(this IBinder container)
            where TSignal : Signal<TParam1, TParam2>
            where TTrigger : Signal<TParam1, TParam2>.TriggerBase
        {
            return BindSignalTrigger<TTrigger, TSignal, TParam1, TParam2>((DiContainer)container, null);
        }

        // Three Parameters
        public static BindingConditionSetter BindSignalTrigger<TTrigger, TSignal, TParam1, TParam2, TParam3>(this IBinder binder, string identifier)
            where TSignal : Signal<TParam1, TParam2, TParam3>
            where TTrigger : Signal<TParam1, TParam2, TParam3>.TriggerBase
        {
            var container = (DiContainer)binder;
            container.Bind<TSignal>().ToSingle();
            container.Bind<Signal<TParam1, TParam2, TParam3>>().ToSingle<TSignal>().WhenInjectedInto<TTrigger>();
            return container.Bind<TTrigger>().ToSingle();
        }

        public static BindingConditionSetter BindSignalTrigger<TTrigger, TSignal, TParam1, TParam2, TParam3>(this IBinder container)
            where TSignal : Signal<TParam1, TParam2, TParam3>
            where TTrigger : Signal<TParam1, TParam2, TParam3>.TriggerBase
        {
            return BindSignalTrigger<TTrigger, TSignal, TParam1, TParam2, TParam3>((DiContainer)container, null);
        }

        // Four Parameters
        public static BindingConditionSetter BindSignalTrigger<TTrigger, TSignal, TParam1, TParam2, TParam3, TParam4>(this IBinder binder, string identifier)
            where TSignal : Signal<TParam1, TParam2, TParam3, TParam4>
            where TTrigger : Signal<TParam1, TParam2, TParam3, TParam4>.TriggerBase
        {
            var container = (DiContainer)binder;
            container.Bind<TSignal>().ToSingle();
            container.Bind<Signal<TParam1, TParam2, TParam3, TParam4>>().ToSingle<TSignal>().WhenInjectedInto<TTrigger>();
            return container.Bind<TTrigger>().ToSingle();
        }

        public static BindingConditionSetter BindSignalTrigger<TTrigger, TSignal, TParam1, TParam2, TParam3, TParam4>(this IBinder container)
            where TSignal : Signal<TParam1, TParam2, TParam3, TParam4>
            where TTrigger : Signal<TParam1, TParam2, TParam3, TParam4>.TriggerBase
        {
            return BindSignalTrigger<TTrigger, TSignal, TParam1, TParam2, TParam3, TParam4>((DiContainer)container, null);
        }

        // Five Parameters
        public static BindingConditionSetter BindSignalTrigger<TTrigger, TSignal, TParam1, TParam2, TParam3, TParam4, TParam5>(this IBinder binder, string identifier)
            where TSignal : Signal<TParam1, TParam2, TParam3, TParam4, TParam5>
            where TTrigger : Signal<TParam1, TParam2, TParam3, TParam4, TParam5>.TriggerBase
        {
            var container = (DiContainer)binder;
            container.Bind<TSignal>().ToSingle();
            container.Bind<Signal<TParam1, TParam2, TParam3, TParam4, TParam5>>().ToSingle<TSignal>().WhenInjectedInto<TTrigger>();
            return container.Bind<TTrigger>().ToSingle();
        }

        public static BindingConditionSetter BindSignalTrigger<TTrigger, TSignal, TParam1, TParam2, TParam3, TParam4, TParam5>(this IBinder container)
            where TSignal : Signal<TParam1, TParam2, TParam3, TParam4, TParam5>
            where TTrigger : Signal<TParam1, TParam2, TParam3, TParam4, TParam5>.TriggerBase
        {
            return BindSignalTrigger<TTrigger, TSignal, TParam1, TParam2, TParam3, TParam4, TParam5>((DiContainer)container, null);
        }

        // Six Parameters
        public static BindingConditionSetter BindSignalTrigger<TTrigger, TSignal, TParam1, TParam2, TParam3, TParam4, TParam5, TParam6>(this IBinder binder, string identifier)
            where TSignal : Signal<TParam1, TParam2, TParam3, TParam4, TParam5, TParam6>
            where TTrigger : Signal<TParam1, TParam2, TParam3, TParam4, TParam5, TParam6>.TriggerBase
        {
            var container = (DiContainer)binder;
            container.Bind<TSignal>().ToSingle();
            container.Bind<Signal<TParam1, TParam2, TParam3, TParam4, TParam5, TParam6>>().ToSingle<TSignal>().WhenInjectedInto<TTrigger>();
            return container.Bind<TTrigger>().ToSingle();
        }

        public static BindingConditionSetter BindSignalTrigger<TTrigger, TSignal, TParam1, TParam2, TParam3, TParam4, TParam5, TParam6>(this IBinder container)
            where TSignal : Signal<TParam1, TParam2, TParam3, TParam4, TParam5, TParam6>
            where TTrigger : Signal<TParam1, TParam2, TParam3, TParam4, TParam5, TParam6>.TriggerBase
        {
            return BindSignalTrigger<TTrigger, TSignal, TParam1, TParam2, TParam3, TParam4, TParam5, TParam6>((DiContainer)container, null);
        }
    }
}
