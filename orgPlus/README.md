

Tested against memory:


```
orgcreated: org115975	 writeDelta: 2.594708ms, readDelat: 913.095µs
orgcreated: org115976	 writeDelta: 2.358454ms, readDelat: 876.684µs
orgcreated: org115977	 writeDelta: 3.347771ms, readDelat: 1.15197ms
orgcreated: org115978	 writeDelta: 2.570379ms, readDelat: 962.041µs
orgcreated: org115979	 writeDelta: 2.945391ms, readDelat: 883.664µs
orgcreated: org115980	 writeDelta: 5.369162ms, readDelat: 1.189241ms
orgcreated: org115981	 writeDelta: 4.895446ms, readDelat: 1.17267ms
orgcreated: org115982	 writeDelta: 2.344278ms, readDelat: 789.322µs
orgcreated: org115983	 writeDelta: 2.361629ms, readDelat: 904.679µs
```

first Failure:
```
orgcreated: org140817	 writeDelta: 2.031701ms, readDelat: 707.37µs
orgcreated: org140818	 writeDelta: 2.078206ms, readDelat: 756.251µs
panic: rpc error: code = Unknown desc = sqlite create: database table is locked: keto_relation_tuples
```

Postgres:

early:
```
orgcreated: org3458	 writeDelta: 24.763296ms, readDelat: 11.767181ms
orgcreated: org3459	 writeDelta: 19.947469ms, readDelat: 3.746926ms
orgcreated: org3460	 writeDelta: 23.829157ms, readDelat: 2.437934ms
orgcreated: org3461	 writeDelta: 21.942835ms, readDelat: 13.993682ms
orgcreated: org3462	 writeDelta: 16.456217ms, readDelat: 3.552686ms
orgcreated: org3463	 writeDelta: 18.853988ms, readDelat: 11.575864ms
orgcreated: org3464	 writeDelta: 16.815647ms, readDelat: 3.212493ms
orgcreated: org3465	 writeDelta: 26.463736ms, readDelat: 2.99829ms
orgcreated: org3466	 writeDelta: 29.908497ms, readDelat: 4.177437ms
orgcreated: org3467	 writeDelta: 18.14862ms, readDelat: 1.805817ms
orgcreated: org3468	 writeDelta: 26.443033ms, readDelat: 3.656936ms
orgcreated: org3469	 writeDelta: 19.653273ms, readDelat: 1.992305ms
orgcreated: org3470	 writeDelta: 17.936319ms, readDelat: 2.261635ms
orgcreated: org3471	 writeDelta: 14.756718ms, readDelat: 2.858929ms
```

some time later:
```
orgcreated: org163652	 writeDelta: 7.596275ms, readDelat: 1.155133ms
orgcreated: org163653	 writeDelta: 7.506345ms, readDelat: 1.131867ms
orgcreated: org163654	 writeDelta: 9.620313ms, readDelat: 1.064613ms
orgcreated: org163655	 writeDelta: 6.612847ms, readDelat: 1.620109ms
orgcreated: org163656	 writeDelta: 7.633135ms, readDelat: 1.305302ms
orgcreated: org163657	 writeDelta: 7.190014ms, readDelat: 1.109845ms
orgcreated: org163658	 writeDelta: 7.473842ms, readDelat: 1.006916ms
orgcreated: org163659	 writeDelta: 6.643777ms, readDelat: 1.503432ms
orgcreated: org163660	 writeDelta: 7.686414ms, readDelat: 1.105809ms
orgcreated: org163661	 writeDelta: 6.673226ms, readDelat: 1.083432ms
```

Postgres got better over time?


more time later:
```
orgcreated: org367738	 writeDelta: 18.424723ms, readDelat: 1.745016ms
orgcreated: org367739	 writeDelta: 11.230903ms, readDelat: 1.712498ms
orgcreated: org367740	 writeDelta: 9.176896ms, readDelat: 2.289539ms
orgcreated: org367741	 writeDelta: 17.47037ms, readDelat: 1.27514ms
orgcreated: org367742	 writeDelta: 10.614369ms, readDelat: 4.957197ms
orgcreated: org367743	 writeDelta: 10.989694ms, readDelat: 1.454222ms
orgcreated: org367744	 writeDelta: 11.213739ms, readDelat: 1.803383ms
orgcreated: org367745	 writeDelta: 9.825083ms, readDelat: 1.433434ms
orgcreated: org367746	 writeDelta: 10.880328ms, readDelat: 1.998395ms
orgcreated: org367747	 writeDelta: 9.350848ms, readDelat: 1.600485ms
orgcreated: org367748	 writeDelta: 10.424988ms, readDelat: 1.628545ms
orgcreated: org367749	 writeDelta: 9.291328ms, readDelat: 1.248032ms
orgcreated: org367750	 writeDelta: 11.040853ms, readDelat: 3.588557ms
orgcreated: org367751	 writeDelta: 10.491353ms, readDelat: 1.24829ms
orgcreated: org367752	 writeDelta: 10.890535ms, readDelat: 2.887754ms
orgcreated: org367753	 writeDelta: 10.410584ms, readDelat: 1.353068ms
orgcreated: org367754	 writeDelta: 11.530038ms, readDelat: 2.182554ms
orgcreated: org367755	 writeDelta: 11.034049ms, readDelat: 3.155937ms
orgcreated: org367756	 writeDelta: 14.114018ms, readDelat: 2.206708ms
```
