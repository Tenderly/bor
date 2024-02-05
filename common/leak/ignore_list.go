package leak

import "go.uber.org/goleak"

func IgnoreList() []goleak.Option {
	return []goleak.Option{
		// a list of goroutne leaks that hard to fix due to external dependencies or too big refactoring needed
		goleak.IgnoreTopFunction("github.com/tenderly/bor/core.(*txSenderCacher).cache"),
		goleak.IgnoreTopFunction("github.com/rjeczalik/notify.(*recursiveTree).dispatch"),
		goleak.IgnoreTopFunction("github.com/rjeczalik/notify.(*recursiveTree).internal"),
		goleak.IgnoreTopFunction("github.com/rjeczalik/notify.(*nonrecursiveTree).dispatch"),
		goleak.IgnoreTopFunction("github.com/rjeczalik/notify.(*nonrecursiveTree).internal"),
		goleak.IgnoreTopFunction("github.com/rjeczalik/notify._Cfunc_CFRunLoopRun"),

		// todo: this leaks should be fixed
		goleak.IgnoreTopFunction("github.com/tenderly/bor/accounts/abi/bind/backends.nullSubscription.func1"),
		goleak.IgnoreTopFunction("github.com/tenderly/bor/accounts/abi/bind/backends.(*filterBackend).SubscribeNewTxsEvent.func1"),
		goleak.IgnoreTopFunction("github.com/tenderly/bor/accounts/abi/bind/backends.(*filterBackend).SubscribePendingLogsEvent.func1"),
		goleak.IgnoreTopFunction("github.com/tenderly/bor/consensus/ethash.(*remoteSealer).loop"),
		goleak.IgnoreTopFunction("github.com/tenderly/bor/core.(*BlockChain).updateFutureBlocks"),
		goleak.IgnoreTopFunction("github.com/tenderly/bor/core/state/snapshot.(*diskLayer).generate"),
		goleak.IgnoreTopFunction("github.com/tenderly/bor/eth/filters.(*EventSystem).eventLoop"),
		goleak.IgnoreTopFunction("github.com/tenderly/bor/event.NewSubscription.func1"),
		goleak.IgnoreTopFunction("github.com/tenderly/bor/event.NewSubscription"),
		goleak.IgnoreTopFunction("github.com/tenderly/bor/metrics.(*meterArbiter).tick"),
	}
}
