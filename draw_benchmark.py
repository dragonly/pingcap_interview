import matplotlib.pyplot as plt


with open('benchmark_local.txt', 'rt') as fd:
    lines = fd.readlines()

data = {}
for line in lines:
    line = line.strip()
    if not line.endswith('ns/op'):
        continue
    split = line.split()
    benchmark_name, record_len = split[0].split('-')[:2]
    benchmark_name = benchmark_name.split('/')[1]
    ns_per_op = int(split[2])
    if benchmark_name not in data:
        data[benchmark_name] = {
            'record_len': [],
            'ns_per_op': [],
        }
    data[benchmark_name]['record_len'].append(record_len)
    data[benchmark_name]['ns_per_op'].append(ns_per_op)

benchmarks = sorted(list(data.keys()))
print(benchmarks)
for benchmark in benchmarks:
    value = data[benchmark]
    record_len = [l[:-3] + 'k' for l in value['record_len']]
    ns_per_op = value['ns_per_op']
    plt.plot(record_len, ns_per_op, '-o', label=benchmark)

plt.grid(True)
plt.xlabel('records number')
plt.ylabel('ns/op')
plt.yscale('log')
plt.legend()
plt.title('In-memory TopN Benchmark')
plt.show()
