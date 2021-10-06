

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
