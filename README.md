LDA implementd using Golang laguage


```sh
./
├── Corpus.go
├── GibbsSampler.go
├── README.md
├── Util.go
├── Vocabulary.go
└── example
    ├── data.txt
    └── example.go
```

### Usage
`cd example`
`go run example.go`
or you can use other k values by run
`go run example.go -k int`

### Result
```
i= 2, LogLikelihood: -430779.972246 time 13.975162 s
i= 3, LogLikelihood: -424804.304056 time 16.269073 s
i= 4, LogLikelihood: -420607.585012 time 17.692962 s
i= 5, LogLikelihood: -418258.669527 time 20.116897 s
i= 6, LogLikelihood: -418185.871020 time 20.992553 s
i= 7, LogLikelihood: -417054.111322 time 23.988074 s
i= 8, LogLikelihood: -416256.521763 time 23.576564 s
i= 9, LogLikelihood: -416308.108587 time 25.713618 s
i= 10, LogLikelihood: -416423.259269 time 27.902169 s
i= 11, LogLikelihood: -416528.240322 time 29.264331 s
i= 12, LogLikelihood: -416734.268834 time 29.420496 s
i= 13, LogLikelihood: -417044.568530 time 31.671346 s
i= 14, LogLikelihood: -418060.253056 time 33.419062 s
i= 15, LogLikelihood: -418203.565668 time 34.985071 s
i= 16, LogLikelihood: -418764.573385 time 36.513028 s
i= 17, LogLikelihood: -419390.658564 time 37.995627 s
i= 18, LogLikelihood: -419439.018298 time 38.983892 s
i= 19, LogLikelihood: -420691.716497 time 40.828173 s
i= 20, LogLikelihood: -421379.494022 time 41.777093 s
```

### Explain
```
➜  example git:(master) ✗ go run example.go -k 8
topic 0 :
knowledge=0.0587025869
logic=0.0351517719
reasoning=0.0320365848
programming=0.0229402382
discovery=0.0153391816
rules=0.0134700693
description=0.0129716393
theory=0.0127224244
databases=0.0123486019
ontologies=0.0120993869

topic 1 :
language=0.0282442375
selection=0.0275109384
search=0.0257999071
feature=0.0221334116
model=0.0201779473
models=0.0184669160
based=0.0170003178
query=0.0168781012
probabilistic=0.0162670187
evaluation=0.0130893892

topic 2 :
support=0.0370168819
networks=0.0271005660
vector=0.0252866057
bayesian=0.0210540318
machines=0.0165795966
classification=0.0165795966
classifiers=0.0137981909
methods=0.0133144681
algorithm=0.0125888841
large=0.0123470227

topic 3 :
systems=0.0557493212
agent=0.0364340743
multi-agent=0.0283768570
system=0.0225270965
agents=0.0203196397
design=0.0193262842
user=0.0167877089
management=0.0150217434
software=0.0141387607
framework=0.0128142867

topic 4 :
learning=0.1402145542
approach=0.0209645363
decision=0.0188094155
machine=0.0186896865
reinforcement=0.0150978185
efficient=0.0128229688
relational=0.0124637820
models=0.0118651373
hierarchical=0.0110270348
induction=0.0107875769

topic 5 :
data=0.0735759300
analysis=0.0353228219
text=0.0293151222
mining=0.0289473039
classification=0.0255143327
clustering=0.0224491798
study=0.0134989333
functions=0.0117824477
techniques=0.0114146294
case=0.0111694171

topic 6 :
planning=0.0593898288
agents=0.0273409356
dynamic=0.0205088072
control=0.0136766788
distributed=0.0129313557
architecture=0.0126829147
environments=0.0118133711
mobile=0.0115649301
adaptive=0.0114407095
interactive=0.0103227249

topic 7 :
web=0.1010863835
information=0.0665925603
semantic=0.0544452339
extraction=0.0286608146
services=0.0160550984
integration=0.0147945268
approach=0.0145653320
service=0.0133047604
documents=0.0130755655
ontology=0.0127317733

k = 8, LogLikelihood: -416528.485967 time 25.262444 s

```
